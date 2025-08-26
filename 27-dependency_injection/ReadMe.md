- Test files should be named `*_test.go`
- Test functions should be named `Test*`
- Test files should be in the same package as the code being tested
- Test functions should be defined in a separate file from the code being tested
- Test functions should have a signature of `func TestXxx(t *testing.T)`
- Test functions should use the `t.Error` or `t.Errorf` function to report errors


- Test functions should use the `t.Log` or `t.Logf` function to report informational messages
- Test functions should use the `t.Fail` or `t.FailNow` function to report failures
- Test functions should use the `t.Skip` or `t.SkipNow` function to skip tests
- Test functions should use the `t.Parallel` function to run tests in parallel


- To run all tests in a package, use the `go test` command in the package directory
- To run all tests in the project, use the `go test ./...` command in the project directory
- To run a specific test function, use the `-run` flag with a regular expression that matches the test function name
- To run tests with verbose output, use the `-v` flag
- To run tests with coverage analysis, use the `-cover` flag


- Example functions should be defined in the test file and should have a signature of `func ExampleXxx()`
- Example functions can have an `Output: ` comment that specifies the expected output and triggers execution of the example function.
- Example functions should be used to demonstrate how to use the code being tested
- Example functions should not take any arguments and should not return any values