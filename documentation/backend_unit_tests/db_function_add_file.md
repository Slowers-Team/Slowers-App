# Database function tests

## Adding a new test file

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

1. First we always have the `package` line, because all the test files are part of the `tests` package, and then the imports:
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
    - First we use the `ConnectDB` function from our `testutils` to initialize the testing database and save it to the `Db` struct field we defined earlier. (The variable `s` contains the testing suite we are currently using. This is because of the "(s *DbFlowerTestSuite)" on the first line "func (s *DbFlowerTestSuite) SetupSuite() {".)
    - Then we empty the database with `s.Db.Clear()` to make sure that it will contain only the entries that we add there in the tests.
    - Lastly we get the flowers test data from our `testdata` package and save them to the struct field `TestFlowers` we defined earlier.

4. In the next section we have all the test functions (see details in the earlier part of these instructions).

5. `TearDownTest()` is called automatically once after each test function:
    ```golang
    func (s *DbFlowerTestSuite) TearDownTest() {
    	s.Db.Clear()
    }
    ```
    - We also empty the database here, so when running the next test we can be sure there are only the entries we added in that test.

6. `TearDownSuite()` is called automatically once after all the tests in this file have been run:
    ```golang
    func (s *DbFlowerTestSuite) TearDownSuite() {
    	testutils.DisconnectDB(s.Db)
    }
    ```
    - Here we disconnect from the database by calling `DisconnectDB()`

7. Lastly we have a test function without the `"(s *DbFlowerTestSuite)"` and with the argument `"t *testing.T"`. The earlier test functions that are part of the test suite are in the wrong format for command `go test` to run them automatically, so we need one test function in the correct format that runs all the tests in the suite:
    ```golang
    func TestDbFlowerTestSuite(t *testing.T) {
    	suite.Run(t, new(DbFlowerTestSuite))
    }
    ```
    - It uses the function `Run` in testify's `suite` package and takes as arguments the testing variable `t` from the arguments of this current function and a newly created `DbFlowerTestSuite` struct.
