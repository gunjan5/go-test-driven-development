package linkedlist

import (
	//"fmt"
	"testing"
)

func TestPushFront(t *testing.T) {
	l := NewList()

	l.PushFront(-2)
	t.Log(l.String())

	l.PushFront(-1)
	t.Log(l.String())

	for i := 0; i < 10; i++ {
		l.PushFront(i)
	}
	t.Log(l.String())

}

func TestPushBack(t *testing.T) {
	l := NewList()

	l.PushBack(-2)
	t.Log(l.String())

	l.PushBack(-1)
	t.Log(l.String())

	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	t.Log(l.String())

}

func TestInsertBefore(t *testing.T) {
	l := NewList()

	e1 := l.PushFront(1)
	e3 := l.PushBack(3)
	t.Log(l.String())

	l.InsertBefore(2, e3)
	t.Log(l.String())

	l.InsertBefore(0, e1)
	t.Log(l.String())

}

func TestInsertAfter(t *testing.T) {
	l := NewList()

	e1 := l.PushFront(1)
	e3 := l.PushBack(3)
	t.Log(l.String())

	l.InsertAfter(2, e1)
	t.Log(l.String())

	l.InsertAfter(4, e3)
	t.Log(l.String())

}

func TestSearch(t *testing.T) {
	l := NewList()

	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	if l.Search(4) != true {
		t.Error("Failed to search")
	}

	if l.Search(99) != false {
		t.Error("Failed to search")
	}
}

func TestDeleteValue(t *testing.T) {
	l := NewList()

	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	t.Log(l.String())

	l.DeleteValue(4)
	t.Log(l.String())

	l.DeleteValue(99)
	t.Log(l.String())

}

func TestDeleteElement(t *testing.T) {
	l := NewList()

	e1 := l.PushFront(1)
	e2 := l.PushBack(2)
	e3 := l.PushBack(3)
	t.Log(l.String())

	l.DeleteElement(e2)
	t.Log(l.String())
	l.DeleteElement(e1)
	t.Log(l.String())
	e0 := l.PushFront(0)
	l.DeleteElement(e3)
	t.Log(l.String())
	l.DeleteElement(e0)
	t.Log(l.String())

}

func TestSize(t *testing.T) {
	l := NewList()

	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	if l.Size() != 10 {
		t.Error("Failed to get the correct size")
	}

}
