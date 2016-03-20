package pack

import (
	"testing"
	"time"
	"os"
	//"fmt"
)


func TestMain(m *testing.M)  {
	println("...setup goes here...")
	result:= m.Run()
	println("...teardown goes here...")
	os.Exit(result)

}

func TestCanAddNumbers(t *testing.T) {
	t.Parallel() //to indicate that this test can be run in parallel with other tests that indicate the same
	result := Add(1, 2)
	time.Sleep(1 * time.Second)

	if result != 3 {
		t.Log("Failed to add two numbers")
		t.Fail()
	}

}

func TestCanAddMultipleNumbers(t *testing.T) {
	t.Parallel() //to indicate that this test can be run in parallel with other tests that indicate the same
	time.Sleep(1 * time.Second)
	if testing.Short() {
		t.Skip("Skipping long test because short flag is enabled")
	}
	result := Add(1, 2, 3, 4)

	if result != 10 {
		t.Error("Failed to add multiple numbers")
	}
}

func TestCanSubNumbers(t *testing.T) {
	t.Parallel() //to indicate that this test can be run in parallel with other tests that indicate the same
	time.Sleep(1 * time.Second)
	result := Sub(10, 5)
	if result != 5 {
		t.Error("Failed to subtract two numbers")
	}
	if testing.Verbose() {
		t.Log("so you want me to talk more???...")
	}
}
