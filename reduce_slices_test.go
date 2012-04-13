package sequences

import "testing"

func TestReduceSlice(t *testing.T) {
	var iterators	[]interface{}
	reportIterationError := func(o, seed interface{}) {
		if e := recover(); e != nil {
			t.Fatalf("Reduce(%v, %v, f) iteration failed with error %v", o, seed, e)
		}
	}

	ConfirmReduce := func(o interface{}, seed, result bool) {
		defer reportIterationError(o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	iterators = []interface{}{
		func(seed, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed, index, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed, v bool) bool { return seed && v },
		func(seed bool, index int, v bool) bool { return seed && v },
		func(seed bool, index interface{}, v bool) bool { return seed && v },
	}
	ConfirmReduce(indexable_slice{}, true, true)
	ConfirmReduce(indexable_slice{false}, true, false)
	ConfirmReduce(indexable_slice{true}, true, true)
	ConfirmReduce(indexable_slice{false, true}, true, false)
	ConfirmReduce(indexable_slice{true, false}, true, false)
	ConfirmReduce(indexable_slice{true, true}, true, true)
	ConfirmReduce(indexable_slice{true, true, true}, true, true)
	ConfirmReduce(indexable_slice{true, true, false, true}, true, false)
//	ConfirmReduce(indexable_function(func(i int) interface{} { return i > 2 }), true, false)
//	ConfirmReduce(indexable_function(func(i int) interface{} { return i < 4 }), true, false)
//	ConfirmReduce(indexable_function(func(i int) interface{} { return i < 4 }), true, false)

/*	iterators = append(
		iterators,
		func(seed, v reflect.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index int, v reflect.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index interface{}, v reflect.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index, v reflect.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
	)
*/
	ConfirmReduce([]bool{}, true, true)
	ConfirmReduce([]bool{false}, true, false)
	ConfirmReduce([]bool{true}, true, true)
	ConfirmReduce([]bool{false, true}, true, false)
	ConfirmReduce([]bool{true, false}, true, false)
	ConfirmReduce([]bool{true, true}, true, true)
	ConfirmReduce([]bool{true, true, true}, true, true)
	ConfirmReduce([]bool{true, true, false, true}, true, false)
}