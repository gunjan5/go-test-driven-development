# go-test-driven-development
Test Driven Development with Golang

Reference repo for Go TDD


### Naming conventions:
- Test file: xxx_test.go
- Test func: TestXxx
- Benchmark: BenchmarkXxx
- Example  : ExampleXxx

I will be following TDD approach, where I write the test first, and then write the main code to make it pass eventually

### Tips:
- Follow the AAA rule for EACH test
- [x] Arrange
- [x] Act
- [x] Assert (assert only one set of actions)

- `t.Error()` = `t.Log()`+`t.Fail()`
- `go test` //for normal testing in the current dir
- `go test <package_name>` // to test the package
- `go test -cover` //to see test coverage
- `go test -coverageprofile=coverage.out; go tool cover -html=coverage.out` //this will open a html in browser with showing which lines of code are covered in green (and not covered with red)
