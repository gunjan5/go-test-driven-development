package stack

import (
	"fmt"
)

type Stack struct {
	data []string
}

func (s Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(in string) {
	s.data = append(s.data, in)

}

func (s Stack) Print() string {
	k := ""
	for i, v := range s.data {
		k += fmt.Sprintf("\n Stack I:%d, V:%v", i, v)
	}
	return k

}
func (s Stack) Size() int {
	k := 0
	for range s.data {
		k++
	}
	return k
}

func (s *Stack) Pop() string {
	if s.isEmpty() {
		return ""
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

// func main() {
// 	// s := Stack{}

// 	// s.Push("abc")
// 	// s.Push("xyz")
// 	// s.Push("qqq")
// 	// s.Push("www")

// 	// fmt.Println(s.isEmpty())
// 	// fmt.Println(s.Print())

// 	// s.Pop()
// 	// s.Pop()
// 	// fmt.Println(s.Print())

// 	// fmt.Println(s.Size())

// 	//fmt.Println(Balance("(})["))
// }
