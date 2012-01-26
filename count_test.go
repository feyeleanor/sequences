package sequences

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	ConfirmCount := func(o, r interface{}, f func(interface{}) bool) {
		if x := Count(o, f); !reflect.DeepEqual(r, x) {
			t.Fatalf("Count(%v, f) should be %v but is %v", o, r, x)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmCount([]int{}, 0, IsPositive)
	ConfirmCount([]interface{}{}, 0, IsPositive)

	ConfirmCount([]int{0}, 0, IsPositive)
	ConfirmCount([]int{1}, 1, IsPositive)
	ConfirmCount([]interface{}{"test"}, 0, IsPositive)
	ConfirmCount([]interface{}{1}, 1, IsPositive)


	ConfirmCount([]int{0, 1}, 1, IsPositive)
	ConfirmCount([]int{1, 2}, 2, IsPositive)
	ConfirmCount([]interface{}{"test", 1}, 1, IsPositive)
	ConfirmCount([]interface{}{1, 2}, 2, IsPositive)
}