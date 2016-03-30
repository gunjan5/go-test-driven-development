# :wrench::nut_and_bolt::hammer::repeat:Test Driven Development (TDD) with Go :raised_hands:
I will be following TDD approach, where I write the test first, and then write the main code to make it pass eventually

### Naming conventions:
- Test file         : xxx_test.go
- Test func         : TestXxx
- Benchmark         : BenchmarkXxx
- Example           : ExampleXxx
- Custom Test runner: TestMain



### Tips:
- Follow the AAA rule for EACH test
  - [x] Arrange
  - [x] Act
  - [x] Assert (assert only one set of actions)
- DRY (Don't Repeat Yourself) also applies to Tests 

## Notes:
- `t.Error()` = `t.Log()`+`t.Fail()` //t.Fail() fails the current test but continues with rest of the tests
- `t.Fatal()` = `t.Log()`+`t.FailNow()` // t.Fatal() and t.FailNow() fails & stops ALL the tests

### Commands:
- `go test` //for normal testing in the current dir
- `go test -v` //for verbose output
- `go test <package_name>` // to test the package
- `go test -run <func_name>` //to run tests only on a specific function instead of the whole file
- `go test -timeout <duration>` //times out & panics the test if not finished in the specified timeout i.e. 2s, 1m, etc.
- `go test -v -short`//Skip some long running tests (Goes with `if testing.Short(){t.Skip("Skipping")}` in the Test func)

### Coverage:
- `go test -cover` //to see test coverage
- `go test -coverprofile=coverage.out; go tool cover -html=coverage.out` //this will open a html in browser with showing which lines of code are covered in green (and not covered with red) (sample image below)

![Code coverage html image](https://raw.githubusercontent.com/gunjan5/go-test-driven-development/master/coverage_html.png?token=AFsMeNKyNVWefbQsy2IorN14XmkzgnUSks5W94VSwA%3D%3D)

### Benchmark:
- `go test -bench .` //run benchmark in the current dir
- `go test -v -bench .` //verbose
