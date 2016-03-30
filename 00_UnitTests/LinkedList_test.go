package linkedlist

import (
	"testing"
)

func TestPushFront(t *testing.T) {
	var l LinkedList

	l.PushFront(2)

	t.Log(l)

	l.PushFront(3)

	t.Log(l)

}
