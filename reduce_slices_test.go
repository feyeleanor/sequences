package sequences

import (
	R "reflect"
	"testing"
)

func reportReductionError(t *testing.T, o, seed interface{}) {
	if e := recover(); e != nil {
panic(e)
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
		func(seed, v bool) bool { return seed && v },
		func(seed bool, index int, v bool) bool { return seed && v },
		func(seed bool, index interface{}, v bool) bool { return seed && v },
		func(seed, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed, index, v interface{}) interface{} { return seed.(bool) && v.(bool) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Bool() && v.Bool()) },
	}

	ConfirmReduce := func(o, seed, result interface{}) {
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

func TestReduceComplex64Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v complex64) complex64 { return seed + v },
		func(seed complex64, index int, v complex64) complex64 { return seed + v },
		func(seed complex64, index interface{}, v complex64) complex64 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(complex64) + v.(complex64) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(complex64) + v.(complex64) },
		func(seed, index, v interface{}) interface{} { return seed.(complex64) + v.(complex64) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Complex() + v.Complex()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Complex() + v.Complex()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Complex() + v.Complex()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Complex() + v.Complex()) },
	}

	ConfirmReduce := func(o interface{}, seed, result complex64) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]complex64{}, 0, 0)
	ConfirmReduce([]complex64{}, 1, 1)
	ConfirmReduce([]complex64{0}, 0, 0)
	ConfirmReduce([]complex64{0}, 1, 1)
	ConfirmReduce([]complex64{1}, 0, 1)
	ConfirmReduce([]complex64{1}, 1, 2)
	ConfirmReduce([]complex64{1, 2}, 0, 3)
	ConfirmReduce([]complex64{1, 2}, 10, 13)
	ConfirmReduce([]complex64{1, 2, 3}, 0, 6)
	ConfirmReduce([]complex64{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]complex64{1, 2, 3, 4}, 10, 20)
}

func TestReduceComplex128Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v complex128) complex128 { return seed + v },
		func(seed complex128, index int, v complex128) complex128 { return seed + v },
		func(seed complex128, index interface{}, v complex128) complex128 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(complex128) + v.(complex128) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(complex128) + v.(complex128) },
		func(seed, index, v interface{}) interface{} { return seed.(complex128) + v.(complex128) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Complex() + v.Complex()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Complex() + v.Complex()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Complex() + v.Complex()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Complex() + v.Complex()) },
	}

	ConfirmReduce := func(o interface{}, seed, result complex128) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]complex128{}, 0, 0)
	ConfirmReduce([]complex128{}, 1, 1)
	ConfirmReduce([]complex128{0}, 0, 0)
	ConfirmReduce([]complex128{0}, 1, 1)
	ConfirmReduce([]complex128{1}, 0, 1)
	ConfirmReduce([]complex128{1}, 1, 2)
	ConfirmReduce([]complex128{1, 2}, 0, 3)
	ConfirmReduce([]complex128{1, 2}, 10, 13)
	ConfirmReduce([]complex128{1, 2, 3}, 0, 6)
	ConfirmReduce([]complex128{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]complex128{1, 2, 3, 4}, 10, 20)
}

func TestReduceErrorSlice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v error) error { return seed.(Error) | v.(Error) },
		func(seed error, index int, v error) error { return seed.(Error) | v.(Error) },
		func(seed error, index interface{}, v error) error { return seed.(Error) | v.(Error) },
		func(seed, v interface{}) interface{} { return seed.(Error) | v.(Error) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(Error) | v.(Error) },
		func(seed, index, v interface{}) interface{} { return seed.(Error) | v.(Error) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Interface().(Error) | v.Interface().(Error)) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Interface().(Error) | v.Interface().(Error)) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Interface().(Error) | v.Interface().(Error)) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Interface().(Error) | v.Interface().(Error)) },
	}

	ConfirmReduce := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]error{}, Error(0), Error(0))
	ConfirmReduce([]error{}, Error(8), Error(8))
	ConfirmReduce([]error{Error(0)}, Error(0), Error(0))
	ConfirmReduce([]error{Error(0)}, Error(8), Error(8))
	ConfirmReduce([]error{Error(1)}, Error(0), Error(1))
	ConfirmReduce([]error{Error(1)}, Error(8), Error(9))
	ConfirmReduce([]error{Error(0), Error(1)}, Error(0), Error(1))
	ConfirmReduce([]error{Error(0), Error(1)}, Error(8), Error(9))
	ConfirmReduce([]error{Error(0), Error(1), Error(2)}, Error(0), Error(3))
	ConfirmReduce([]error{Error(0), Error(1), Error(2)}, Error(8), Error(11))
	ConfirmReduce([]error{Error(0), Error(1), Error(2), Error(4)}, Error(0), Error(7))
	ConfirmReduce([]error{Error(0), Error(1), Error(2), Error(4)}, Error(8), Error(15))
}

func TestReduceFloat32Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v float32) float32 { return seed + v },
		func(seed float32, index int, v float32) float32 { return seed + v },
		func(seed float32, index interface{}, v float32) float32 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(float32) + v.(float32) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(float32) + v.(float32) },
		func(seed, index, v interface{}) interface{} { return seed.(float32) + v.(float32) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Float() + v.Float()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Float() + v.Float()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Float() + v.Float()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Float() + v.Float()) },
	}

	ConfirmReduce := func(o interface{}, seed, result float32) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]float32{}, 0, 0)
	ConfirmReduce([]float32{}, 1, 1)
	ConfirmReduce([]float32{0}, 0, 0)
	ConfirmReduce([]float32{0}, 1, 1)
	ConfirmReduce([]float32{1}, 0, 1)
	ConfirmReduce([]float32{1}, 1, 2)
	ConfirmReduce([]float32{1, 2}, 0, 3)
	ConfirmReduce([]float32{1, 2}, 10, 13)
	ConfirmReduce([]float32{1, 2, 3}, 0, 6)
	ConfirmReduce([]float32{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]float32{1, 2, 3, 4}, 10, 20)
}

func TestReduceFloat64Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v float64) float64 { return seed + v },
		func(seed float64, index int, v float64) float64 { return seed + v },
		func(seed float64, index interface{}, v float64) float64 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(float64) + v.(float64) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(float64) + v.(float64) },
		func(seed, index, v interface{}) interface{} { return seed.(float64) + v.(float64) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Float() + v.Float()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Float() + v.Float()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Float() + v.Float()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Float() + v.Float()) },
	}

	ConfirmReduce := func(o interface{}, seed, result float64) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]float64{}, 0, 0)
	ConfirmReduce([]float64{}, 1, 1)
	ConfirmReduce([]float64{0}, 0, 0)
	ConfirmReduce([]float64{0}, 1, 1)
	ConfirmReduce([]float64{1}, 0, 1)
	ConfirmReduce([]float64{1}, 1, 2)
	ConfirmReduce([]float64{1, 2}, 0, 3)
	ConfirmReduce([]float64{1, 2}, 10, 13)
	ConfirmReduce([]float64{1, 2, 3}, 0, 6)
	ConfirmReduce([]float64{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]float64{1, 2, 3, 4}, 10, 20)
}

func TestReduceIntSlice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, index, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
	}

	ConfirmReduce := func(o, seed, result interface{}) {
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
}

func TestReduceInt8Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v int8) int8 { return seed + v },
		func(seed int8, index int, v int8) int8 { return seed + v },
		func(seed int8, index interface{}, v int8) int8 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(int8) + v.(int8) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int8) + v.(int8) },
		func(seed, index, v interface{}) interface{} { return seed.(int8) + v.(int8) },
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
	}

	ConfirmReduce := func(o interface{}, seed, result int8) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int8{}, 0, 0)
	ConfirmReduce([]int8{}, 10, 10)
	ConfirmReduce([]int8{0}, 0, 0)
	ConfirmReduce([]int8{0}, 10, 10)
	ConfirmReduce([]int8{1}, 0, 1)
	ConfirmReduce([]int8{1}, 10, 11)
	ConfirmReduce([]int8{1, 2}, 0, 3)
	ConfirmReduce([]int8{1, 2}, 10, 13)
	ConfirmReduce([]int8{1, 2, 3}, 0, 6)
	ConfirmReduce([]int8{1, 2, 3}, 10, 16)
	ConfirmReduce([]int8{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int8{1, 2, 3, 4}, 10, 20)
}

func TestReduceInt16Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v int16) int16 { return seed + v },
		func(seed int16, index int, v int16) int16 { return seed + v },
		func(seed int16, index interface{}, v int16) int16 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(int16) + v.(int16) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int16) + v.(int16) },
		func(seed, index, v interface{}) interface{} { return seed.(int16) + v.(int16) },
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
	}

	ConfirmReduce := func(o interface{}, seed, result int16) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int16{}, 0, 0)
	ConfirmReduce([]int16{}, 10, 10)
	ConfirmReduce([]int16{0}, 0, 0)
	ConfirmReduce([]int16{0}, 10, 10)
	ConfirmReduce([]int16{1}, 0, 1)
	ConfirmReduce([]int16{1}, 10, 11)
	ConfirmReduce([]int16{1, 2}, 0, 3)
	ConfirmReduce([]int16{1, 2}, 10, 13)
	ConfirmReduce([]int16{1, 2, 3}, 0, 6)
	ConfirmReduce([]int16{1, 2, 3}, 10, 16)
	ConfirmReduce([]int16{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int16{1, 2, 3, 4}, 10, 20)
}

func TestReduceInt32Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v int32) int32 { return seed + v },
		func(seed int32, index int, v int32) int32 { return seed + v },
		func(seed int32, index interface{}, v int32) int32 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(int32) + v.(int32) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int32) + v.(int32) },
		func(seed, index, v interface{}) interface{} { return seed.(int32) + v.(int32) },
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
	}

	ConfirmReduce := func(o interface{}, seed, result int32) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int32{}, 0, 0)
	ConfirmReduce([]int32{}, 10, 10)
	ConfirmReduce([]int32{0}, 0, 0)
	ConfirmReduce([]int32{0}, 10, 10)
	ConfirmReduce([]int32{1}, 0, 1)
	ConfirmReduce([]int32{1}, 10, 11)
	ConfirmReduce([]int32{1, 2}, 0, 3)
	ConfirmReduce([]int32{1, 2}, 10, 13)
	ConfirmReduce([]int32{1, 2, 3}, 0, 6)
	ConfirmReduce([]int32{1, 2, 3}, 10, 16)
	ConfirmReduce([]int32{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int32{1, 2, 3, 4}, 10, 20)
}

func TestReduceInt64Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v int64) int64 { return seed + v },
		func(seed int64, index int, v int64) int64 { return seed + v },
		func(seed int64, index interface{}, v int64) int64 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(int64) + v.(int64) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int64) + v.(int64) },
		func(seed, index, v interface{}) interface{} { return seed.(int64) + v.(int64) },
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Int() + v.Int()) },
	}

	ConfirmReduce := func(o interface{}, seed, result int64) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int64{}, 0, 0)
	ConfirmReduce([]int64{}, 10, 10)
	ConfirmReduce([]int64{0}, 0, 0)
	ConfirmReduce([]int64{0}, 10, 10)
	ConfirmReduce([]int64{1}, 0, 1)
	ConfirmReduce([]int64{1}, 10, 11)
	ConfirmReduce([]int64{1, 2}, 0, 3)
	ConfirmReduce([]int64{1, 2}, 10, 13)
	ConfirmReduce([]int64{1, 2, 3}, 0, 6)
	ConfirmReduce([]int64{1, 2, 3}, 10, 16)
	ConfirmReduce([]int64{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int64{1, 2, 3, 4}, 10, 20)
}

func TestReduceInterface(t *testing.T) {
	iterators := []interface{}{
		func(seed, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, index, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(int(seed.Int() + v.Int())) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(int(seed.Int() + v.Int())) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(int(seed.Int() + v.Int())) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(int(seed.Int() + v.Int())) },
	}

	ConfirmReduce := func(o interface{}, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]interface{}{}, 0, 0)
	ConfirmReduce([]interface{}{}, 10, 10)
	ConfirmReduce([]interface{}{0}, 0, 0)
	ConfirmReduce([]interface{}{0}, 10, 10)
	ConfirmReduce([]interface{}{1}, 0, 1)
	ConfirmReduce([]interface{}{1}, 10, 11)
	ConfirmReduce([]interface{}{1, 2}, 0, 3)
	ConfirmReduce([]interface{}{1, 2}, 10, 13)
	ConfirmReduce([]interface{}{1, 2, 3}, 0, 6)
	ConfirmReduce([]interface{}{1, 2, 3}, 10, 16)
	ConfirmReduce([]interface{}{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]interface{}{1, 2, 3, 4}, 10, 20)

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
	
	ConfirmReduce([]interface{}{}, false, false)
	ConfirmReduce([]interface{}{}, true, true)
	ConfirmReduce([]interface{}{false}, false, false)
	ConfirmReduce([]interface{}{false}, true, false)
	ConfirmReduce([]interface{}{true}, false, false)
	ConfirmReduce([]interface{}{true}, true, true)
	ConfirmReduce([]interface{}{false, true}, true, false)
	ConfirmReduce([]interface{}{true, false}, true, false)
	ConfirmReduce([]interface{}{true, true}, true, true)
	ConfirmReduce([]interface{}{true, true, true}, true, true)
	ConfirmReduce([]interface{}{true, true, false, true}, true, false)
}

func TestReduceStringSlice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v string) string { return seed + v },
		func(seed string, index int, v string) string { return seed + v },
		func(seed string, index interface{}, v string) string { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(string) + v.(string) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(string) + v.(string) },
		func(seed, index, v interface{}) interface{} { return seed.(string) + v.(string) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.String() + v.String()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.String() + v.String()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.String() + v.String()) },
		func(seed, index, v R.Value) R.Value { return R.ValueOf(seed.String() + v.String()) },
	}

	ConfirmReduce := func(o interface{}, seed, result string) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]string{}, "", "")
	ConfirmReduce([]string{}, "Z", "Z")
	ConfirmReduce([]string{""}, "", "")
	ConfirmReduce([]string{""}, "Z", "Z")
	ConfirmReduce([]string{"A"}, "", "A")
	ConfirmReduce([]string{"A"}, "Z", "ZA")
	ConfirmReduce([]string{"A", "B"}, "", "AB")
	ConfirmReduce([]string{"A", "B"}, "Z", "ZAB")
	ConfirmReduce([]string{"A", "B", "C"}, "", "ABC")
	ConfirmReduce([]string{"A", "B", "C"}, "Z", "ZABC")
	ConfirmReduce([]string{"A", "B", "C", "D"}, "", "ABCD")
	ConfirmReduce([]string{"A", "B", "C", "D"}, "Z", "ZABCD")
}

func TestReduceUintSlice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v uint) uint { return seed + v },
		func(seed uint, index int, v uint) uint { return seed + v },
		func(seed uint, index interface{}, v uint) uint { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(uint) + v.(uint) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(uint) + v.(uint) },
		func(seed, index, v interface{}) interface{} { return seed.(uint) + v.(uint) },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
	}

	ConfirmReduce := func(o interface{}, seed, result uint) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint{}, 0, 0)
	ConfirmReduce([]uint{}, 10, 10)
	ConfirmReduce([]uint{0}, 0, 0)
	ConfirmReduce([]uint{0}, 10, 10)
	ConfirmReduce([]uint{1}, 0, 1)
	ConfirmReduce([]uint{1}, 10, 11)
	ConfirmReduce([]uint{1, 2}, 0, 3)
	ConfirmReduce([]uint{1, 2}, 10, 13)
	ConfirmReduce([]uint{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint{1, 2, 3, 4}, 10, 20)
}

func TestReduceUint8Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v uint8) uint8 { return seed + v },
		func(seed uint8, index int, v uint8) uint8 { return seed + v },
		func(seed uint8, index interface{}, v uint8) uint8 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(uint8) + v.(uint8) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(uint8) + v.(uint8) },
		func(seed, index, v interface{}) interface{} { return seed.(uint8) + v.(uint8) },
		func(seed, v uint) uint { return seed + v },
		func(seed uint, index int, v uint) uint { return seed + v },
		func(seed uint, index interface{}, v uint) uint { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
	}

	ConfirmReduce := func(o interface{}, seed, result uint8) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint8{}, 0, 0)
	ConfirmReduce([]uint8{}, 10, 10)
	ConfirmReduce([]uint8{0}, 0, 0)
	ConfirmReduce([]uint8{0}, 10, 10)
	ConfirmReduce([]uint8{1}, 0, 1)
	ConfirmReduce([]uint8{1}, 10, 11)
	ConfirmReduce([]uint8{1, 2}, 0, 3)
	ConfirmReduce([]uint8{1, 2}, 10, 13)
	ConfirmReduce([]uint8{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint8{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint8{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint8{1, 2, 3, 4}, 10, 20)
}

func TestReduceUint16Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v uint16) uint16 { return seed + v },
		func(seed uint16, index int, v uint16) uint16 { return seed + v },
		func(seed uint16, index interface{}, v uint16) uint16 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(uint16) + v.(uint16) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(uint16) + v.(uint16) },
		func(seed, index, v interface{}) interface{} { return seed.(uint16) + v.(uint16) },
		func(seed, v uint) uint { return seed + v },
		func(seed uint, index int, v uint) uint { return seed + v },
		func(seed uint, index interface{}, v uint) uint { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
	}

	ConfirmReduce := func(o interface{}, seed, result uint16) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint16{}, 0, 0)
	ConfirmReduce([]uint16{}, 10, 10)
	ConfirmReduce([]uint16{0}, 0, 0)
	ConfirmReduce([]uint16{0}, 10, 10)
	ConfirmReduce([]uint16{1}, 0, 1)
	ConfirmReduce([]uint16{1}, 10, 11)
	ConfirmReduce([]uint16{1, 2}, 0, 3)
	ConfirmReduce([]uint16{1, 2}, 10, 13)
	ConfirmReduce([]uint16{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint16{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint16{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint16{1, 2, 3, 4}, 10, 20)
}

func TestReduceUint32Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v uint32) uint32 { return seed + v },
		func(seed uint32, index int, v uint32) uint32 { return seed + v },
		func(seed uint32, index interface{}, v uint32) uint32 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(uint32) + v.(uint32) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(uint32) + v.(uint32) },
		func(seed, index, v interface{}) interface{} { return seed.(uint32) + v.(uint32) },
		func(seed, v uint) uint { return seed + v },
		func(seed uint, index int, v uint) uint { return seed + v },
		func(seed uint, index interface{}, v uint) uint { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
	}

	ConfirmReduce := func(o interface{}, seed, result uint32) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint32{}, 0, 0)
	ConfirmReduce([]uint32{}, 10, 10)
	ConfirmReduce([]uint32{0}, 0, 0)
	ConfirmReduce([]uint32{0}, 10, 10)
	ConfirmReduce([]uint32{1}, 0, 1)
	ConfirmReduce([]uint32{1}, 10, 11)
	ConfirmReduce([]uint32{1, 2}, 0, 3)
	ConfirmReduce([]uint32{1, 2}, 10, 13)
	ConfirmReduce([]uint32{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint32{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint32{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint32{1, 2, 3, 4}, 10, 20)
}

func TestReduceUint64Slice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v uint64) uint64 { return seed + v },
		func(seed uint64, index int, v uint64) uint64 { return seed + v },
		func(seed uint64, index interface{}, v uint64) uint64 { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(uint64) + v.(uint64) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(uint64) + v.(uint64) },
		func(seed, index, v interface{}) interface{} { return seed.(uint64) + v.(uint64) },
		func(seed, v uint) uint { return seed + v },
		func(seed uint, index int, v uint) uint { return seed + v },
		func(seed uint, index interface{}, v uint) uint { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
	}

	ConfirmReduce := func(o interface{}, seed, result uint64) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint64{}, 0, 0)
	ConfirmReduce([]uint64{}, 10, 10)
	ConfirmReduce([]uint64{0}, 0, 0)
	ConfirmReduce([]uint64{0}, 10, 10)
	ConfirmReduce([]uint64{1}, 0, 1)
	ConfirmReduce([]uint64{1}, 10, 11)
	ConfirmReduce([]uint64{1, 2}, 0, 3)
	ConfirmReduce([]uint64{1, 2}, 10, 13)
	ConfirmReduce([]uint64{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint64{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint64{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint64{1, 2, 3, 4}, 10, 20)
}

func TestReduceUintptrSlice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v uintptr) uintptr { return seed + v },
		func(seed uintptr, index int, v uintptr) uintptr { return seed + v },
		func(seed uintptr, index interface{}, v uintptr) uintptr { return seed + v },
		func(seed, v interface{}) interface{} { return seed.(uintptr) + v.(uintptr) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(uintptr) + v.(uintptr) },
		func(seed, index, v interface{}) interface{} { return seed.(uintptr) + v.(uintptr) },
		func(seed, v uintptr) uintptr { return seed + v },
		func(seed uintptr, index int, v uintptr) uintptr { return seed + v },
		func(seed uintptr, index interface{}, v uintptr) uintptr { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(seed.Uint() + v.Uint()) },
	}

	ConfirmReduce := func(o interface{}, seed, result uintptr) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uintptr{}, 0, 0)
	ConfirmReduce([]uintptr{}, 10, 10)
	ConfirmReduce([]uintptr{0}, 0, 0)
	ConfirmReduce([]uintptr{0}, 10, 10)
	ConfirmReduce([]uintptr{1}, 0, 1)
	ConfirmReduce([]uintptr{1}, 10, 11)
	ConfirmReduce([]uintptr{1, 2}, 0, 3)
	ConfirmReduce([]uintptr{1, 2}, 10, 13)
	ConfirmReduce([]uintptr{1, 2, 3}, 0, 6)
	ConfirmReduce([]uintptr{1, 2, 3}, 10, 16)
	ConfirmReduce([]uintptr{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uintptr{1, 2, 3, 4}, 10, 20)
}

func TestReduceRSlice(t *testing.T) {
	iterators := []interface{}{
		func(seed, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, index, v interface{}) interface{} { return seed.(int) + v.(int) },
		func(seed, v int) int { return seed + v },
		func(seed int, index int, v int) int { return seed + v },
		func(seed int, index interface{}, v int) int { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(int(seed.Int() + v.Int())) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(int(seed.Int() + v.Int())) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(int(seed.Int() + v.Int())) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(int(seed.Int() + v.Int())) },
	}

	ConfirmReduceInterfaceSeed := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(o interface{}, seed, result R.Value) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v.(R.Value).Interface() != result.Interface() {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result.Interface(), v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed([]R.Value{}, 0, 0)
	ConfirmReduceInterfaceSeed([]R.Value{}, 10, 10)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(0)}, 0, 0)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(0)}, 10, 10)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(1)}, 0, 1)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(1)}, 10, 11)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(1), R.ValueOf(2)}, 0, 3)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(1), R.ValueOf(2)}, 10, 13)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3)}, 0, 6)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3)}, 10, 16)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 0, 10)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 10, 20)

	ConfirmReduceValueSeed([]R.Value{}, R.ValueOf(0), R.ValueOf(0))
	ConfirmReduceValueSeed([]R.Value{}, R.ValueOf(10), R.ValueOf(10))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(0)}, R.ValueOf(0), R.ValueOf(0))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(0)}, R.ValueOf(10), R.ValueOf(10))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(1)}, R.ValueOf(0), R.ValueOf(1))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(1)}, R.ValueOf(10), R.ValueOf(11))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(1), R.ValueOf(2)}, R.ValueOf(0), R.ValueOf(3))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(1), R.ValueOf(2)}, R.ValueOf(10), R.ValueOf(13))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3)}, R.ValueOf(0), R.ValueOf(6))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3)}, R.ValueOf(10), R.ValueOf(16))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, R.ValueOf(0), R.ValueOf(10))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, R.ValueOf(10), R.ValueOf(20))

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

	ConfirmReduceInterfaceSeed([]R.Value{}, false, false)
	ConfirmReduceInterfaceSeed([]R.Value{}, true, true)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(false)}, false, false)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(false)}, true, false)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(true)}, false, false)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(true)}, true, true)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(false), R.ValueOf(true)}, true, false)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(true), R.ValueOf(false)}, true, false)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(true), R.ValueOf(true)}, true, true)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(true), R.ValueOf(true), R.ValueOf(true)}, true, true)
	ConfirmReduceInterfaceSeed([]R.Value{R.ValueOf(true), R.ValueOf(true), R.ValueOf(false), R.ValueOf(true)}, true, false)
	
	ConfirmReduceValueSeed([]R.Value{}, R.ValueOf(false), R.ValueOf(false))
	ConfirmReduceValueSeed([]R.Value{}, R.ValueOf(true), R.ValueOf(true))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(false)}, R.ValueOf(false), R.ValueOf(false))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(false)}, R.ValueOf(true), R.ValueOf(false))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(true)}, R.ValueOf(false), R.ValueOf(false))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(true)}, R.ValueOf(true), R.ValueOf(true))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(false), R.ValueOf(true)}, R.ValueOf(true), R.ValueOf(false))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(true), R.ValueOf(false)}, R.ValueOf(true), R.ValueOf(false))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(true), R.ValueOf(true)}, R.ValueOf(true), R.ValueOf(true))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(true), R.ValueOf(true), R.ValueOf(true)}, R.ValueOf(true), R.ValueOf(true))
	ConfirmReduceValueSeed([]R.Value{R.ValueOf(true), R.ValueOf(true), R.ValueOf(false), R.ValueOf(true)}, R.ValueOf(true), R.ValueOf(false))
}