## OUTPUT:

### RUN1:
```bash
$ go test -v -bench .
testing: warning: no tests to run
PASS
BenchmarkSample-8	10000000	       202 ns/op
ok  	github.com/gunjan5/go-test-driven-development/Benchmark	2.232s
```

### RUN2:
```bash
$ go test -bench .
testing: warning: no tests to run
PASS
BenchmarkSample-8	10000000	       209 ns/op
ok  	github.com/gunjan5/go-test-driven-development/Benchmark	2.319s
```
