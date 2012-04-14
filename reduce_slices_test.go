package sequences

import (
	R "reflect"
	"testing"
)

func reportReductionError(t *testing.T, o, seed interface{}) {
	if e := recover(); e != nil {
		t.Fatalf("Reduce(%v, %v, f) iteration failed with error %v", o, seed, e)
	}
}

func TestReduceIndexable(t *testing.T) {
	iterators := []interface{}{
		func(seed, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, index, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
	}

	ConfirmReduce := func(o interface{}, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(indexable_slice{}, 0, 0)
	ConfirmReduce(indexable_slice{}, 10, 10)
	ConfirmReduce(indexable_slice{0}, 0, 0)
	ConfirmReduce(indexable_slice{0}, 10, 10)
	ConfirmReduce(indexable_slice{1}, 0, 1)
	ConfirmReduce(indexable_slice{1}, 10, 11)
	ConfirmReduce(indexable_slice{1, 2}, 0, 3)
	ConfirmReduce(indexable_slice{1, 2}, 10, 13)
	ConfirmReduce(indexable_slice{1, 2, 3}, 0, 6)
	ConfirmReduce(indexable_slice{1, 2, 3}, 10, 16)
	ConfirmReduce(indexable_slice{1, 2, 3, 4}, 0, 10)
	ConfirmReduce(indexable_slice{1, 2, 3, 4}, 10, 20)
	ConfirmReduce(indexable_function(func(i int) interface{} { return i }), 0, 55)
	ConfirmReduce(indexable_function(func(i int) interface{} { return i }), 10, 65)

	iterators = []interface{}{
		func(seed, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed, index, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed, v bool) bool { return seed && v },
		func(seed bool, index int, v bool) bool { return seed && v },
		func(seed bool, index interface{}, v bool) bool { return seed && v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
	}
	
	ConfirmReduce(indexable_slice{}, false, false)
	ConfirmReduce(indexable_slice{}, true, true)
	ConfirmReduce(indexable_slice{false}, false, false)
	ConfirmReduce(indexable_slice{false}, true, false)
	ConfirmReduce(indexable_slice{true}, false, false)
	ConfirmReduce(indexable_slice{true}, true, true)
	ConfirmReduce(indexable_slice{false, true}, true, false)
	ConfirmReduce(indexable_slice{true, false}, true, false)
	ConfirmReduce(indexable_slice{true, true}, true, true)
	ConfirmReduce(indexable_slice{true, true, true}, true, true)
	ConfirmReduce(indexable_slice{true, true, false, true}, true, false)
}

func TestReduceBoolSlice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed, index, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed, v bool) bool { return seed && v },
		func(seed bool, index int, v bool) bool { return seed && v },
		func(seed bool, index interface{}, v bool) bool { return seed && v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
	}

	ConfirmReduce := func(o interface{}, seed, result bool) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]bool{}, false, false)
	ConfirmReduce([]bool{}, true, true)
	ConfirmReduce([]bool{false}, false, false)
	ConfirmReduce([]bool{false}, true, false)
	ConfirmReduce([]bool{true}, false, false)
	ConfirmReduce([]bool{true}, true, true)
	ConfirmReduce([]bool{false, true}, true, false)
	ConfirmReduce([]bool{true, false}, true, false)
	ConfirmReduce([]bool{true, true}, true, true)
	ConfirmReduce([]bool{true, true, true}, true, true)
	ConfirmReduce([]bool{true, true, false, true}, true, false)
}

func TestReduceIntSlice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, index, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
	}

	ConfirmReduce := func(o interface{}, seed, result int) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int{}, 0, 0)
	ConfirmReduce([]int{}, 10, 10)
	ConfirmReduce([]int{0}, 0, 0)
	ConfirmReduce([]int{0}, 10, 10)
	ConfirmReduce([]int{1}, 0, 1)
	ConfirmReduce([]int{1}, 10, 11)
	ConfirmReduce([]int{1, 2}, 0, 3)
	ConfirmReduce([]int{1, 2}, 10, 13)
	ConfirmReduce([]int{1, 2, 3}, 0, 6)
	ConfirmReduce([]int{1, 2, 3}, 10, 16)
	ConfirmReduce([]int{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int{1, 2, 3, 4}, 10, 20)

	ConfirmReduce(indexable_slice{}, 0, 0)
	ConfirmReduce(indexable_slice{}, 10, 10)
	ConfirmReduce(indexable_slice{0}, 0, 0)
	ConfirmReduce(indexable_slice{0}, 10, 10)
	ConfirmReduce(indexable_slice{1}, 0, 1)
	ConfirmReduce(indexable_slice{1}, 10, 11)
	ConfirmReduce(indexable_slice{1, 2}, 0, 3)
	ConfirmReduce(indexable_slice{1, 2}, 10, 13)
	ConfirmReduce(indexable_slice{1, 2, 3}, 0, 6)
	ConfirmReduce(indexable_slice{1, 2, 3}, 10, 16)
	ConfirmReduce(indexable_slice{1, 2, 3, 4}, 0, 10)
	ConfirmReduce(indexable_slice{1, 2, 3, 4}, 10, 20)
	ConfirmReduce(indexable_function(func(i int) interface{} { return i }), 0, 55)
	ConfirmReduce(indexable_function(func(i int) interface{} { return i }), 10, 65)
}