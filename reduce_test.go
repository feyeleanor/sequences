package sequences

import (
	//	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	ConfirmReduce := func(o, s, r, f interface{}) {
		/*		if x, _ := Reduce(o, s, f); !reflect.DeepEqual(x, r) {
					t.Fatalf("Reduce(%v, %v, %v) should be %v but is %v", o, s, f, r, x)
				}
		*/
	}

	Sum := func(memo, x interface{}) interface{} {
		return memo.(int) + x.(int)
	}

	ConfirmReduce(nil, nil, nil, Sum)
	ConfirmReduce(nil, 0, 0, Sum)

	ConfirmReduce([]int{}, 0, 0, Sum)
	ConfirmReduce([]int{}, 1, 1, Sum)
	ConfirmReduce([]int{0}, 0, 0, Sum)
	ConfirmReduce([]int{0}, 1, 1, Sum)
	ConfirmReduce([]int{0, 1, 2, 3}, 0, 6, Sum)
	ConfirmReduce([]int{0, 1, 2, 3}, 1, 7, Sum)

	ConfirmReduce(map[int]int{}, 0, 0, Sum)
	ConfirmReduce(map[int]int{}, 1, 1, Sum)
	ConfirmReduce(map[int]int{0: 0}, 0, 0, Sum)
	ConfirmReduce(map[int]int{0: 0}, 1, 1, Sum)
	ConfirmReduce(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}, 0, 6, Sum)
	ConfirmReduce(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}, 1, 7, Sum)

	IntSum := func(memo, x int) int {
		return memo + x
	}

	ConfirmReduce([]int{}, 0, 0, IntSum)
	ConfirmReduce([]int{}, 1, 1, IntSum)
	ConfirmReduce([]int{0}, 0, 0, IntSum)
	ConfirmReduce([]int{0}, 1, 1, IntSum)
	ConfirmReduce([]int{0, 1, 2, 3}, 0, 6, IntSum)
	ConfirmReduce([]int{0, 1, 2, 3}, 1, 7, IntSum)

	ConfirmReduce(map[int]int{}, 0, 0, IntSum)
	ConfirmReduce(map[int]int{}, 1, 1, IntSum)
	ConfirmReduce(map[int]int{0: 0}, 0, 0, IntSum)
	ConfirmReduce(map[int]int{0: 0}, 1, 1, IntSum)
	ConfirmReduce(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}, 0, 6, IntSum)
	ConfirmReduce(map[int]int{0: 0, 1: 1, 2: 2, 3: 3}, 1, 7, IntSum)
}

func TestReduceIndexableFunction(t *testing.T) {
	iterators := []interface{}{
		func(seed, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Int() + v.Int())
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Int() + v.Int())
		},
	}

	ConfirmReduce := func(o interface{}, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(indexable_function(func(i int) interface{} { return i }), 0, 55)
	ConfirmReduce(indexable_function(func(i int) interface{} { return i }), 10, 65)
}

func TestReduceIndexableSlice(t *testing.T) {
	s := func(v ...interface{}) indexable_slice {
		return indexable_slice(v)
	}

	iterators := []interface{}{
		func(seed, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Int() + v.Int())
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Int() + v.Int())
		},
	}

	ConfirmReduce := func(o interface{}, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(s(), 0, 0)
	ConfirmReduce(s(), 10, 10)
	ConfirmReduce(s(0), 0, 0)
	ConfirmReduce(s(0), 10, 10)
	ConfirmReduce(s(1), 0, 1)
	ConfirmReduce(s(1), 10, 11)
	ConfirmReduce(s(1, 2), 0, 3)
	ConfirmReduce(s(1, 2), 10, 13)
	ConfirmReduce(s(1, 2, 3), 0, 6)
	ConfirmReduce(s(1, 2, 3), 10, 16)
	ConfirmReduce(s(1, 2, 3, 4), 0, 10)
	ConfirmReduce(s(1, 2, 3, 4), 10, 20)

	iterators = []interface{}{
		func(seed, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
	}

	ConfirmReduce(s(), false, false)
	ConfirmReduce(s(), true, true)
	ConfirmReduce(s(false), false, false)
	ConfirmReduce(s(false), true, false)
	ConfirmReduce(s(true), false, false)
	ConfirmReduce(s(true), true, true)
	ConfirmReduce(s(false, true), true, false)
	ConfirmReduce(s(true, false), true, false)
	ConfirmReduce(s(true, true), true, true)
	ConfirmReduce(s(true, true, true), true, true)
	ConfirmReduce(s(true, true, false, true), true, false)
}
