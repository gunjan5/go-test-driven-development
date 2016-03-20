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

## Coverage.out: (`go test -coverageprofile=coverage.out`)
```bash
mode: set
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:4.24,7.17 2 1
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:11.2,11.22 1 1
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:14.2,14.12 1 1
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:7.17,10.3 2 0
github.com/gunjan5/go-test-driven-development/BasicTests/math.go:11.22,13.3 1 1
```

## HTML Cool output from `coverage.out` file: (`go tool cover -html=coverage.out`)
![Code coverage html image](https://raw.githubusercontent.com/gunjan5/go-test-driven-development/master/coverage_html.png?token=AFsMeNKyNVWefbQsy2IorN14XmkzgnUSks5W94VSwA%3D%3D)
