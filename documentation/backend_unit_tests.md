# Backend unit tests

- Implemented with:
  - Go standard library: `testing` package ([API reference here](https://pkg.go.dev/testing), though for the most cases `testify` library's documentation is more useful)
  - `testify`, a testing library
    - [README](https://github.com/stretchr/testify#testify---thou-shalt-write-tests) (quick overview of all testify's features and some example code, a good first stop when you need help)
    - [API reference](https://pkg.go.dev/github.com/stretchr/testify) (has more example code)
    - (`testify` provides additional features such as assertions, suites and mocks to Go's standard `testing` package)
- Test data in `backend/testdata` (Usage examples in later sections)
- Helper functions for testing in `backend/testutils` (Usage examples in later sections)

## API routes

- Tests are located in `backend/apitests`
- More specifically, the tests for the API routes whose handlers are in `backend/handlers/<filename_here>.go` are located in the file `backend/apitests/<filename_here>s_test.go`
  - So for example the tests for the API routes for flowers are located in the file `backend/apitests/flowers_test.go`

(Adding tests to existing test files)[backend_unit_tests/api_route_add_test.md]

(Adding a new test file)[backend_unit_tests/api_route_add_file.md]

## Database functions

- Tests are located in `backend/database/tests`
- More specifically, the tests for the database functions in `backend/database/<filename_here>.go` are located in the file `backend/database/tests/<filename_here>_test.go`
  - So for example the tests for `backend/database/flower.go` are located in the file `backend/database/tests/flower_test.go`

(Adding tests to existing test files)[backend_unit_tests/db_function_add_test.md]

(Adding a new test file)[backend_unit_tests/db_function_add_file.md]

