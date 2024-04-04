package math

import (
	"os"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	// -1, 0, 1
	if Abs(-1) < 0 {
		t.Error("Negative value found in abs() with", -1)
	}

	if Abs(0) < 0 {
		t.Error("Negative value found in abs() with", 0)
	}

	if Abs(1) < 0 {
		t.Error("Negative value found in abs() with", 1)
	}
}

func TestAbsSub(t *testing.T) {
	t.Run("Positibe", func(t *testing.T) {
		if Abs(1) < 0 {
			t.Error("Negative value found in abs() with", 1)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		if Abs(0) < 0 {
			t.Error("Negative value found in abs() with", 0)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		if Abs(-1) < 0 {
			t.Error("Negative value found in abs() with", -1)
		}
	})
}

func TestSkip(t *testing.T) {
	if len(os.Getenv("GOPATH")) == 0 {
		t.Skip("Skipping test because GOPATH isn't set")
	}

	t.Log("Tested with GOPATH: ", os.Getenv("GOPATH"))
}

func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		t.Log("Cleanup")
	})

	t.Log("Running some test")
	t.Log("Running some test")
	t.Log("Running some test")
	t.Log("Running some test")
}

func TestParallelOne(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func TestParallelTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func TestParallelThree(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
