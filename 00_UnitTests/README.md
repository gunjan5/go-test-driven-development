## Coverage:
```bash
$ go test -v -cover
=== RUN   TestCanAddNumbers
--- PASS: TestCanAddNumbers (0.00s)
=== RUN   TestCanAddMultipleNumbers
--- PASS: TestCanAddMultipleNumbers (0.00s)
PASS
coverage: 71.4% of statements
ok  	github.com/gunjan5/go-test-driven-development/BasicTests	0.005s
```

## HTML Cool output from `coverage.out` file: (`go tool cover -html=coverage.out`)
![Code coverage html image](https://raw.githubusercontent.com/gunjan5/go-test-driven-development/master/coverage_html.png?token=AFsMeNKyNVWefbQsy2IorN14XmkzgnUSks5W94VSwA%3D%3D)

## Coverage.out: (`go test -coverageprofile=coverage.out`)
```bash
mode: set
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:4.24,7.17 2 1
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:11.2,11.22 1 1
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:14.2,14.12 1 1
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:7.17,10.3 2 0
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:11.22,13.3 1 1
```

## Short:
```go
if testing.Short() {
  t.Skip("Skipping long test because short flag is enabled")
}
```
```bash
$ go test -v -short
=== RUN   TestCanAddNumbers
--- PASS: TestCanAddNumbers (0.00s)
=== RUN   TestCanAddMultipleNumbers
--- SKIP: TestCanAddMultipleNumbers (0.00s)
	math_test.go:20: Skipping long test because short flag is enabled
=== RUN   TestCanSubNumbers
--- PASS: TestCanSubNumbers (0.00s)
	math_test.go:35: so you want me to talk more???...
PASS
ok  	github.com/gunjan5/go-test-driven-development/BasicTests	0.006s
```
## Non-parallel (notice the total time & sequence):
```bash
$ go test -v
=== RUN   TestCanAddNumbers
--- PASS: TestCanAddNumbers (1.00s)
=== RUN   TestCanAddMultipleNumbers
--- PASS: TestCanAddMultipleNumbers (1.00s)
=== RUN   TestCanSubNumbers
--- PASS: TestCanSubNumbers (1.00s)
	math_test.go:38: so you want me to talk more???...
PASS
ok  	github.com/gunjan5/go-test-driven-development/BasicTests	3.009s
```

## Parallel (notice the total time & sequence):
- `t.Parallel()` //to indicate that this test can be run in parallel with other tests that indicate the same
```bash
$ go test -v
=== RUN   TestCanAddNumbers
=== RUN   TestCanAddMultipleNumbers
=== RUN   TestCanSubNumbers
--- PASS: TestCanAddMultipleNumbers (1.00s)
--- PASS: TestCanAddNumbers (1.00s)
--- PASS: TestCanSubNumbers (1.00s)
	math_test.go:41: so you want me to talk more???...
PASS
ok  	github.com/gunjan5/go-test-driven-development/BasicTests	1.008s
```

## Custom Test runner (notice setup and teardown):
```go
func TestMain(m *testing.M)  {
	println("...setup goes here...")
	result:= m.Run()
	println("...teardown goes here...")
	os.Exit(result)
}
```
```bash
$ go test -v
...setup goes here...
=== RUN   TestCanAddNumbers
=== RUN   TestCanAddMultipleNumbers
=== RUN   TestCanSubNumbers
--- PASS: TestCanAddMultipleNumbers (1.00s)
--- PASS: TestCanAddNumbers (1.00s)
--- PASS: TestCanSubNumbers (1.00s)
	math_test.go:52: so you want me to talk more???...
PASS
...teardown goes here...
ok  	github.com/gunjan5/go-test-driven-development/BasicTests	1.007s
```
