package pack

import (
	"testing"
)

func TestCanAddNumbers(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Log("Failed to add")
		t.Fail()
	}

}

func TestCanAddMultipleNumbers(t *testing.T) {
	result := Add(1, 2, 3, 4)

	if result != 10 {
		t.Error("Failed to add multiple numbers")
	}
}
