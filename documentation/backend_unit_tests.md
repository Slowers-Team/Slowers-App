# Backend unit tests

- Implemented with:
  - Go standard library: `testing` package ([API reference here](https://pkg.go.dev/testing), though for the most cases `testify` library's documentation is more useful)
  - `testify`, a testing library
    - [README](https://github.com/stretchr/testify#testify---thou-shalt-write-tests) (quick overview of all testify's features and some example code, a good first stop when you need help)
    - [API reference](https://pkg.go.dev/github.com/stretchr/testify) (has more example code)
    - (`testify` provides additional features such as assertions, suites and mocks to Go's standard `testing` package)
- Test data in `backend/testdata` (Usage examples in later sections)
- Helper functions for testing in `backend/testutils` (Usage examples in later sections)

## Database functions

- Tests are located in `backend/database/tests`
- More specifically, the tests for the database functions in `backend/database/<filename_here>.go` are located in the file `backend/database/tests/<filename_here>_test.go`
  - So for example the tests for `backend/database/flower.go` are located in the file `backend/database/tests/flower_test.go`

### Adding tests to existing test files

Let's look at an example test that tests `GetFlower()` function in `backend/database/flower.go` and go through it bit by bit. First here is the whole code:
```golang
func (s *DbFlowerTestSuite) TestAddAndGetFlower() {
	flowers := testdata.GetTestFlowers()
	flower := flowers[0]
	flower.ID = database.NilObjectID

	s.Db.AddFlower(context.Background(), flower)

	fetchedFlowers, err := s.Db.GetFlowers(context.Background())

	s.Require().NoError(
		err,
		"GetFlowers() should not return an error",
	)
	s.Require().Len(
		fetchedFlowers,
		1,
		"GetFlowers() should return a slice of length 1",
	)
	s.NotZero(
		fetchedFlowers[0].ID,
		"fetched flower should have non-zero ID",
	)
	s.Equal(
		flower.Name,
		fetchedFlowers[0].Name,
		"wrong Name for the flower returned from GetFlowers()",
	)
	s.Equal(
		flower.LatinName,
		fetchedFlowers[0].LatinName,
		"wrong LatinName for the flower returned from GetFlowers()",
	)
	// I left out the s.Equal() tests for the other fields (AddedTime, Grower...),
	// because they work similarly as the tests for Name and LatinName above
}
```

1. All tests for flowers begin with something like this:
    ```golang
    func (s *DbFlowerTestSuite) TestInsertSomeTestNameHere() {
    ```
    - The test function's name (in our full example `TestAddAndGetFlower`) has to always start with `Test`, but you can freely choose the rest of it
    - (The `(s *DbFlowerTestSuite)` part will be explained in the next section, but it is the same for all tests in the same file)

2. Next we get flower test data from our `testdata` package:
    ```golang
    flowers := testdata.GetTestFlowers()
    flower := flowers[0]
    flower.ID = database.NilObjectID
    ```
    - First we save the test flowers to the `flowers` variable
    - Next we save the first one of the test flowers to the `flower` variable for ease of use
    - Then we set its object ID to nil ID, because in this case when the flower is added to the database, the database will pick a suitable ID for it

3. To be able to test getting flowers from the database, we need to have some flowers in the database, so we add the flower from the previous step there by calling the database function `AddFlower()`:
    ```golang
    s.Db.AddFlower(context.Background(), flower)
    ```
    - Normally outside the tests we would call it like `db.AddFlower(c.Context(), flower)`, but:
      - Here we don't want to mess up the actual database, so we will be using a separate testing database, hence we have `s.Db.AddFlower` instead of `db.AddFlower`
      - We are not inside a HTTP request handler, so we can't use its context `c.Context()`, so instead we are using a default context provided by Go's standard library: `context.Background()` (More info about Go contexts [here](https://pkg.go.dev/context), though you don't really have to understand them to be able to write unit tests)

4. Now we get to call the function `GetFlowers()` that we are actually testing here:
    ```golang
    fetchedFlowers, err := s.Db.GetFlowers(context.Background())
    ```
    - The returned flowers are saved to the variable `fetchedFlowers` and the possible error to the variable `err`

5. In the last part of the test we check that the function `GetFlowers()` returned what it should have. First we check that there was no error:
    ```golang
    s.Require().NoError(
    	err,
    	"GetFlowers() should not return an error",
    )
    ```
    - The checks (i.e. _assertions_) all start with either "s." or "s.Require().". When using just "s." and the check fails the failure will be reported in the end but otherwise the testing continues normally to the next check. When using "s.Require()." and the check fails, all following checks for this test will be skipped. The point why we want to use "s.Require()." here is that if `GetFlowers()` returned an error (i.e. the check fails) then `fetchedFlowers` won't have any meaningful data in it and it doesn't make sense to do any further checks for it, so we want to skip them.
    - `NoError()` checks that the error variable given to it (here `err`) doesn't contain an error.
    - The `"GetFlowers() should not return an error"` given as the second argument for `NoError()` contains a message shown if the check fails. The message is always contained in the last argument of the check.

6. Next we check that we fetched exactly one flower (because earlier we added to the database exactly one flower) and use the `Require()` in the check so that all following checks are skipped if this one fails:
    ```golang
    s.Require().Len(
    	fetchedFlowers,
    	1,
    	"GetFlowers() should return a slice of length 1",
    )
    ```

7. Finally we check that the one flower we fetched (i.e. `fetchedFlowers[0]`) contains the data it should. First we check that its object ID is not nil (the ID was chosen by the database and varies from test run to test run, so this is the only thing we can check about it):
    ```golang
    s.NotZero(
    	fetchedFlowers[0].ID,
    	"fetched flower should have non-zero ID",
    )
    ```

8. Then we check that all other fields of the fetched flower match with the fields of the original flower:
    ```golang
    s.Equal(
    	flower.Name,
    	fetchedFlowers[0].Name,
    	"wrong Name for the flower returned from GetFlowers()",
    )
    ```
    - The first argument in `Equal()` contains the expected value (here `flower.Name`) and the second argument the actual value (here `fetchedFlowers[0].Name`)
- A list of all possible assertion functions in `testify` (s.NoError(), s.Len(), s.Equal()...) can be found [here](https://pkg.go.dev/github.com/stretchr/testify/assert#pkg-functions) (Though in our tests we leave out the first argument "t TestingT" because we are using testify's `suite` feature)

### Adding a new test file

Let's look at the structure of the test file `backend/database/tests/flower_test.go` and go through it bit by bit. First here is the whole file (with some parts omitted):
```golang
package tests

import (
	...
)

type DbFlowerTestSuite struct {
	suite.Suite
	Db          database.Database
	TestFlowers []database.Flower
}

func (s *DbFlowerTestSuite) SetupSuite() {
	s.Db = testutils.ConnectDB()
	s.Db.Clear()
	s.TestFlowers = testdata.GetTestFlowers()
}

func (s *DbFlowerTestSuite) TestAddFlower() {
	...
}

func (s *DbFlowerTestSuite) TestAddAndGetFlower() {
	...
}

...

func (s *DbFlowerTestSuite) TearDownTest() {
	s.Db.Clear()
}

func (s *DbFlowerTestSuite) TearDownSuite() {
	testutils.DisconnectDB(s.Db)
}

func TestDbFlowerTestSuite(t *testing.T) {
	suite.Run(t, new(DbFlowerTestSuite))
}
```

1. First we always have the `package` line, because all the test files are part of the `tests` package, and the imports:
```golang
package tests

import (
	...
)
```

2. Then we define the testing suite struct:
    ```golang
    type DbFlowerTestSuite struct {
    	suite.Suite
    	Db          database.Database
    	TestFlowers []database.Flower
    }
    ```
    - A suite is a collection of tests inside one file.
    - You can pick the name freely (here `DbFlowerTestSuite`), but you will need to use it later in the beginning of each test function.
    - We need to always have `suite.Suite` in the beginning of the struct, so that Go knows we are defining a suite using testify's suite feature here.
    - After that we always have `Db database.Database`. This struct field will contain the testing database we will use in the tests.
    - Then we can freely define any other struct fields we want to use. For example here we will save the test flowers to the struct field `TestFlowers []databa.Flower` for ease of use.

3. Next we have the `SetupSuite()` function. It is called automatically once before running any tests from this file.
    ```golang
    func (s *DbFlowerTestSuite) SetupSuite() {
    	s.Db = testutils.ConnectDB()
    	s.Db.Clear()
    	s.TestFlowers = testdata.GetTestFlowers()
    }
    ```
    - First we use the `ConnectDB` function from our `testutils` to initialize the testing database and save it to the `Db` struct field we defined earlier. (The variable `s` contains the testing suite we are currently using. This is because of the "(s *DbFlowerTestSuite)" on the line "func (s *DbFlowerTestSuite) SetupSuite() {".)
    - Then we empty the database with `s.Db.Clear()` to make sure that it will contain only the entries that we add there in the tests.
    - Lastly we get the flowers test data from our `testdata` package and save them to the struct field `TestFlowers` we defined earlier.

4. In this next section we have all the test functions.

5. `TearDownTest()` is called automatically once after each test function:
    ```golang
    func (s *DbFlowerTestSuite) TearDownTest() {
    	s.Db.Clear()
    }
    ```
    - We also empty the database here, so when running the next test we can be sure there are only the entries we added in that test.

6. `TearDownSuite()` is called automatically once after all the tests have been run:
    ```golang
    func (s *DbFlowerTestSuite) TearDownSuite() {
    	testutils.DisconnectDB(s.Db)
    }
    ```
    - Here we disconnect from the database by calling `DisconnectDB()`

7. Lastly we have a test function without the `"(s *DbFlowerTestSuite)"` and with the argument `"t *testing.T"`. The earlier test functions that are part of the test suite are in the wrong format for command `go test` to run them automatically. So we need one test function in the correct format that runs all the tests in the suite:
    ```golang
    func TestDbFlowerTestSuite(t *testing.T) {
    	suite.Run(t, new(DbFlowerTestSuite))
    }
    ```
    - It uses the function `Run` in testify's `suite` package and takes as arguments the testing variable `t` from the arguments of this current function and a newly created `DbFlowerTestSuite` struct.
