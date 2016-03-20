package table

import (
  "testing"
)

var addTable = []struct{
  in []int
  out int
} {
  {[]int{1,2},3},
  {[]int{1,2,3,4},10},
  {[]int{},0},
  {[]int{-2,-3}, -5},
}

func TestCanAddNumbers(t *testing.T)  {
  t.Parallel()
  for _,entry := range addTable{
    result := Add(entry.in...)
    if result != entry.out{
      t.Error("Failed to add numbers, Expected: ",entry.out , " Got: ", result)
    }
  }

}
