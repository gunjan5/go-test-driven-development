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

## Notes:
- `t.Error()` = `t.Log()`+`t.Fail()` //t.Fail() fails the current test but continues with rest of the tests
- `t.Fatal()` = `t.Log()`+`t.FailNow()` // t.Fatal() and t.FailNow() fails & stops ALL the tests

### Commands:
- `go test` //for normal testing in the current dir
- `go test -v` //for verbose output
- `go test <package_name>` // to test the package

### Coverage:
- `go test -cover` //to see test coverage
- `go test -coverageprofile=coverage.out; go tool cover -html=coverage.out` //this will open a html in browser with showing which lines of code are covered in green (and not covered with red) (sample image below)

![Code coverage html image](https://raw.githubusercontent.com/gunjan5/go-test-driven-development/master/coverage_html.png?token=AFsMeNKyNVWefbQsy2IorN14XmkzgnUSks5W94VSwA%3D%3D)

### Benchmark:
- `go test -bench .` //run benchmark in the current dir
- `go test -v -bench .` //verbose
