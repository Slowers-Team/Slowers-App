# API route tests

## Adding a new test file

Let's look at the structure of the test file `backend/apitests/flowers_test.go` and go through it bit by bit. First here is the whole file (with some parts omitted):
```golang
package apitests

import (
	...
)

type FlowersAPITestSuite struct {
	suite.Suite
	TestFlowers []database.Flower
}

func (s *FlowersAPITestSuite) SetupSuite() {
	s.TestFlowers = testdata.GetTestFlowers()
}

func (s *FlowersAPITestSuite) TestListingFlowersWithoutError() {
	...
}

func (s *FlowersAPITestSuite) TestListingFlowersWithError() {
	...
}

...

func TestFlowersAPITestSuite(t *testing.T) {
	suite.Run(t, new(FlowersAPITestSuite))
}
```

1. First we always have the `package` line, because all the test files are part of the `apitests` package, and then the imports:
```golang
package apitests

import (
	...
)
```

2. Then we define the testing suite struct:
    ```golang
    type FlowersAPITestSuite struct {
    	suite.Suite
    	TestFlowers []database.Flower
    }
    ```
    - A suite is a collection of tests inside one file.
    - You can pick the name freely (here `FlowersAPITestSuite`), but you will need to use it later in the beginning of each test function.
    - We need to always have `suite.Suite` in the beginning of the struct, so that Go knows we are defining a suite using testify's suite feature here.
    - Other than that we can freely define any other struct fields we want to use. For example here we will save the test flowers to the struct field `TestFlowers []databa.Flower` for ease of use.

3. Next we have the `SetupSuite()` function. It is called automatically once before running any tests from this file.
    ```golang
    func (s *FlowersAPITestSuite) SetupSuite() {
    	s.TestFlowers = testdata.GetTestFlowers()
    }
    ```
    - Here we get the test data from our `testdata` package and save it the fields we defined for the suite struct earlier. In the example code, we get the flowers test data with `testdata.GetTestFlowers()` and save it to the struct field `TestFlowers` we defined earlier.

4. In the next section we have all the test functions (see details in the earlier part of these instructions).

5. Lastly we have a test function without the `"(s *FlowersAPITestSuite)"` and with the argument `"t *testing.T"`. The earlier test functions that are part of the test suite are in the wrong format for command `go test` to run them automatically, so we need one test function in the correct format that runs all the tests in the suite:
    ```golang
    func TestFlowersAPITestSuite(t *testing.T) {
    	suite.Run(t, new(FlowersAPITestSuite))
    }
    ```
    - It uses the function `Run` in testify's `suite` package and takes as arguments the testing variable `t` from the arguments of this current function and a newly created `FlowersAPITestSuite` struct.
