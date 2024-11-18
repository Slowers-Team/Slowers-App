# Database function tests

## Adding tests to existing test files

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
	// The s.Equal() tests for the other fields (AddedTime, Grower...) are omitted,
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
