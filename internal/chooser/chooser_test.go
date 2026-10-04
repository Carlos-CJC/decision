package chooser

import (
	"errors"
	"testing"
)

func TestPickReturnsMember(t *testing.T) {
	opts := []string{"a", "b", "c"}
	for range 100 {
		got, err := Pick(opts)
		if err != nil {
			t.Fatalf("Pick: %v", err)
		}
		if got != "a" && got != "b" && got != "c" {
			t.Fatalf("Pick = %q, not one of %v", got, opts)
		}
	}
}

func TestPickSingle(t *testing.T) {
	got, err := Pick([]string{"only"})
	if err != nil {
		t.Fatalf("Pick: %v", err)
	}
	if got != "only" {
		t.Errorf("Pick = %q, want %q", got, "only")
	}
}

func TestPickEmpty(t *testing.T) {
	if _, err := Pick(nil); !errors.Is(err, ErrNoOptions) {
		t.Errorf("Pick(nil) err = %v, want ErrNoOptions", err)
	}
}
