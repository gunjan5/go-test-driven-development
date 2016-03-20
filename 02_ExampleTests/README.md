##INPUT:
```go
fmt.Println("Gopher")
fmt.Println("Go")
```

## RUN1 with FAIL:
```bash
//Output:
//NotGopher
```
```bash
$ go test
--- FAIL: ExampleTest (0.00s)
got:
Gopher
Go
want:
Hello
FAIL
exit status 1
FAIL	github.com/gunjan5/go-test-driven-development/ExampleTests	0.006s
```

## RUN2 with PASS:
```bash
//Output:
//Gopher
//Go
```
```bash
$ go test
PASS
ok  	github.com/gunjan5/go-test-driven-development/ExampleTests	0.006s
```
