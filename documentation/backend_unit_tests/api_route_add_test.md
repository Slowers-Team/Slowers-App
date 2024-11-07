# API route tests

## Adding tests to existing test files

Let's look at an example test that tests the `GET /api/flowers` API route and go through it bit by bit. First here is the whole code:
```golang
func (s *FlowersAPITestSuite) TestListingFlowers() {
	testutils.RunTest(s.T(), testutils.TestCase{
		Description:  "GET /api/flowers",
		Route:        "/api/flowers",
		Method:       "GET",
		Body:         []byte{},
		ExpectedCode: 200,
		ExpectedBody: utils.FlowersToJSON(s.TestFlowers),
		SetupMocks: func(db *mocks.Database) {
			db.EXPECT().GetFlowers(
				mock.Anything,
			).Return(
				s.TestFlowers, nil,
			).Once()
		},
	})
}
```

1. All tests for flowers begin with something like this:
    ```golang
    func (s *FlowersAPITestSuite) TestListingFlowers() {
    ```
    - The test function's name (in our full example `TestListingFlowers`) has to always start with `Test`, but you can freely choose the rest of it
    - (The `(s *FlowersAPITestSuite)` part will be explained in the next section, but it is the same for all tests in the same file)

2. Inside the test function we run the test using the `RunTest()` function in our `testutils`:
    ```golang
    testutils.RunTest(s.T(), testutils.TestCase{
    ```
    - The first argument is always `s.T()`. (This passes a testing struct used by Go's standard `testing` library to the `RunTest` function, so that `RunTest()` is able to communicate the successes/failures of the tests to the library.)
    - The second argument is a `TestCase` struct (defined in our `testutils`), in which we will give all the details of the test. Next we will go through every field of this struct and explain their purpose.

3. A description text that you can freely choose for the test:
    ```golang
    Description:  "GET /api/flowers",
    ```
    - Will be shown when the test fails.
    - (For example the format "HTTPMETHODHERE /api/route/here" used here works fine.)

4. The API route you are testing:
    ```golang
    Route:        "/api/flowers",
    ```
    - You can find all this application's routes in `application/application.go`. For example the route we are testing in this example is defined on the line `app.Get("/api/flowers", handlers.GetFlowers)` there.
    - If you are testing a route with parameters such as `/api/flowers/:id`, you would have something like `Route: "/api/flowers/" + flower.ID.Hex(),` here. (The `Hex()` method gives the ObjectID as a string.)

5. The HTTP method your API route uses:
    ```golang
    Method:       "GET",
    ```
    - You can find the HTTP method in `application/application.go` as the method name (such as `app.Get(`).

6. The body of the request (usually the data of the flowers/site/user... to create/modify/... given as JSON, when the body is not empty) as a slice of bytes (an array/list of variable length is called a _slice_ in Go):
    ```golang
    Body:         []byte{},
    ```
    - This route does not use a request body, so we use just an empty slice of bytes `[]byte{}` here
    - If you want to give for example the string `"some_data_to_send_to_the_backend"` as the body, you would convert it to a slice of bytes by using the function `[]byte()`. So the complete line would be `Body: []byte("some_data_to_send_to_the_backend"),`.
    - (If the body is not empty, we need to have a content type field before this field. It would almost always be just `ContentType: "application/json",` because we are almost always using JSON data.)

7. The expected HTTP status code:
    ```golang
    ExpectedCode: 200,
    ```
    - Here we expect the listing operation to succeed, so the status code should be 200 (i.e. "OK").
    - You can find the status code an API route should return in the corresponding handler function (`GetFlowers()` in `handlers/flower.go` for this example). It could be given there like `return c.SendStatus(204)` or when there is no status code defined it is 200 by default.

8. The expected body of the response from the API route:
    ```golang
    ExpectedBody: utils.FlowersToJSON(s.TestFlowers),
    ```
    - The `GET /api/flowers` route should return all flowers in the database, so the response body should contain all flowers from the database as JSON.
    - Here the database will contain the flowers in the test data, so we use `s.TestFlowers`.
    - The file `backend/utils/json.go` contains conversion functions from different data types to JSON. Here we use `utils.FlowersToJSON()`, because we need to convert a slice of flowers to JSON (an array/list of variable length is called a _slice_ in Go).
    - (This `ExpectedBody` has also a variation called `ExpectedBodyFunc`. This let's you give instead of a `[]byte` a function of type `func(body []byte)` for testing the expected body. The function gets the response body as an argument automatically and then you would test the body manually inside the function by using `testify` assertions the same way as in the `Database functions` section of these instructions. You can find an example of how to use `ExpectedBodyFunc` in the `TestAddingFlower()` test function in `apitests/flowers_test.go`. `ExpectedBodyFunc` should be used instead of `ExpectedBody` when the expected body varies from test run to test run. This would happen for example when adding a new flower, because whenever adding a new flower, its AddedTime is set to the current time at that moment, so the AddedTime would vary depending on when the flower was added.)

9. We don't want the API route to fetch the flower data from an actual database, because we want to test the handler function code of the API route separately from the database code. For that reason we setup mocks for the database functions in this last field of the struct:
    ```golang
    SetupMocks: func(db *mocks.Database) {
    ```
    - The mock setup begins always with this line. We are saving in the `SetupMocks` field an anonymous function `func(db *mocks.Database)` in which the mocks are set up and this function is called automatically before the actual test is run.

10. Here we need to mock only one database function (i.e. `func GetFlowers(ctx context.Context) ([]Flower, error)`):
    ```golang
    db.EXPECT().GetFlowers(
    	mock.Anything,
    ).Return(
    	s.TestFlowers, nil,
    ).Once()
    ```
    - You can find out, which database functions you have to mock by checking, which database functions are called in the handler function for the API route (in this example the handler function `func GetFlowers(c *fiber.Ctx) error` is located in `handlers/flower.go` and only one database function is called there, that is `db.GetFlowers()` on the line `flowers, err := db.GetFlowers(c.Context())`).
    - A mock basically specifies, what value/values the function should return when certain arguments have been given to it.
    - Here normally we would call the function `GetFlowers` like `db.GetFlowers(c.Context())`, so the first line of the mock has `db.EXPECT().GetFlowers(`, that is `db.GetFlowers(` with ".EXPECT()" added in the middle. The idea is that we are telling the testing library which function we _expect_ to be called during the test.
    - The next line contains the arguments for which this mock is called. The function `GetFlowers()` gets `c.Context()` as an argument, but this `c.Context()` depends on the implementation of the API routes (i.e. whether they have been implemented with Fiber as we have or somehow else). We don't want the unit tests to be dependent on the implementation of the routes, so will just put `mock.Anything` there, which means that this mock is called whenever the function `db.GetFlowers()` is called with any one argument. (As another example, for the database function `func DeleteFlower(ctx context.Context, id ObjectID) (bool, error)` we would have something like `mock.Anything, flower.ID,` on this line, because besides the context it also takes as the second argument the ObjectID of the flower to delete.)
    - Next we call the `Return()` method and give it the return values that we want the function `db.GetFlowers()` to return when called during the test. You can check from the original `db.GetFlowers()` function in `database/flower.go`, what kind of values it should return. You will find there that it has return values of types `([]Flower, error)` (see the line `func (mDb MongoDatabase) GetFlowers(ctx context.Context) ([]Flower, error) {` in that file) and that normally it returns `flowers` and `nil` (see the line `return flowers, nil`), so now we are using `s.TestFlowers` as the first return value of this mock instead of `flowers` and `nil` as the error, because there should be no error.
    - Because of the `Once()` method in the last line, it will be automatically checked that this mocked function will be called exactly once during the test (as it should, because `db.GetFlowers()` is called exactly once in the handler function).
    - (In this example we define only one mock, but in some tests there may be more. Then you would just do this kind of mock definition multiple times inside the function given to the `SetupMocks` field, once for each mock. The order of these definitions won't matter.)
