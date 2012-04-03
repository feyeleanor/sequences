package sequences

import "testing"
import R "reflect"

func TestIsFunctionSequence(t *testing.T) {
	Valid := func(f interface{}) bool {
		return isFunctionSequence(R.ValueOf(f))
	}

	Invalid := func(f interface{}) bool {
		return !isFunctionSequence(R.ValueOf(f))
	}

	if Valid(17) {
		t.Fatalf("17 should not be a valid function sequence")
	}

	if Invalid(func(i int) int { return i }) {
		t.Fatalf("func(int) int{} should be a valid function sequence")
	}
}