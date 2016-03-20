# Table Driven Tests:
- DRY (Don't Repeat Yourself) also applies to Tests, so instead of writing multiple tests for the same function, write a table

### CODE:
```go
var addTable = []struct{
  in []int
  out int
} {
  {[]int{1,2},3},
  {[]int{1,2,3,4},10},
  {[]int{},0},
  {[]int{-2,-3}, -5},
}
```

```go
for _,entry := range addTable{ //range over the table
  result := Add(entry.in...) //spread the input
  if result != entry.out{
    t.Error("Failed to add numbers, Expected: ",entry.out , " Got: ", result)
  }
}
```


### RUN:

```bash
$ go test -v
=== RUN   TestCanAddNumbers
--- PASS: TestCanAddNumbers (0.00s)
PASS
ok  	github.com/gunjan5/go-test-driven-development/01_TableDrivenTests	0.005s
```
