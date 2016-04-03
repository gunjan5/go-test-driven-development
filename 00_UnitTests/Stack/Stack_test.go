package stack

import (
	"testing"
)

func TestStackIsEmpty(t *testing.T) {
	s := NewStack()

	if s.isEmpty() != true {
		t.Error("Stack is empty")
	}

}

func TestStackPush(t *testing.T) {
	s := NewStack()

	if s.isEmpty() != true {
		t.Error("Stack is empty, but it's returning not-empty")
	}

	s.Push("abc")
	s.Push("xyz")
	if s.isEmpty() != false {
		t.Error("Stack is not empty, but it's returning empty")
	}
	if testing.Verbose() {
		t.Log(s.Print())
	}

}

func TestStackSize(t *testing.T) {
	s := NewStack()

	result := s.Size()
	if result != 0 {
		t.Errorf("Size should be 0, Got: %d", result)
	}

	s.Push("aaa")
	result2 := s.Size()
	if result2 != 1 {
		t.Errorf("Size should be 1, Got: %d", result2)
	}

}

func TestStackPop(t *testing.T) {
	s := NewStack()

	s.Push("aaa")

	if testing.Verbose() {
		t.Log(s.Size())
		t.Log(s.Pop())
		t.Log(s.Size())
	}
	s.Pop()

}
