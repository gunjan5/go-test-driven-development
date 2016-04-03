package linkedlist

import (
	"fmt"
)

type LinkedList struct {
	Head, Tail *Node
	Count      int
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func NewList() *LinkedList {
	return &LinkedList{}

}

func (l *LinkedList) String() string {
	s := ""
	for n := l.Head; n != nil; n = n.Next {
		s += fmt.Sprintf("\n%d", n.Value)
	}

	s += fmt.Sprintf("\nCount: %d", l.Count)

	return s

}

func (l *LinkedList) PushFront(v int) *Node {
	//_ = "breakpoint"
	n := &Node{v, nil, nil}

	if l.Head == l.Tail && l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head.Prev = n
		l.Head = n
	}

	l.Count++

	return n

}

func (l *LinkedList) PushBack(v int) *Node {
	n := &Node{v, nil, nil}

	if l.Head == l.Tail && l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		n.Prev = l.Tail
		l.Tail = n
	}
	l.Count++
	return n

}

func (l *LinkedList) InsertBefore(v int, e *Node) *Node {
	n := &Node{v, nil, nil}

	if l.Head == e {
		return l.PushFront(v)
	} else {
		e.Prev.Next = n
		n.Prev = e.Prev
		n.Next = e
		e.Prev = n

		l.Count++
	}

	return n

}

func (l *LinkedList) InsertAfter(v int, e *Node) *Node {
	n := &Node{v, nil, nil}

	if l.Tail == e {
		return l.PushBack(v)
	} else {
		e.Next.Prev = n
		n.Next = e.Next
		e.Next = n
		n.Prev = e

		l.Count++
	}

	return n

}

func (l *LinkedList) Search(v int) bool {

	for n := l.Head; n != nil; n = n.Next {
		if n.Value == v {
			return true
		}
	}

	return false

}

func (l *LinkedList) DeleteValue(v int) {
	for n := l.Head; n != nil; n = n.Next {
		if n.Value == v {
			n.Prev.Next = n.Next
			n.Next.Prev = n.Prev

			l.Count--
			break
		}
	}

}

func (l *LinkedList) DeleteElement(e *Node) {
	_ = "breakpoint"
	if l.Head == l.Tail && l.Head == e {
		l.Head = nil
		l.Tail = nil
	} else if l.Head == e {
		l.Head = e.Next
		e.Next.Prev = nil
	} else if l.Tail == e {
		l.Tail = e.Prev
		e.Prev.Next = nil
	} else {
		e.Next.Prev = e.Prev
		e.Prev.Next = e.Next
	}
	l.Count--

}

func (l *LinkedList) Size() int {
	return l.Count
}
