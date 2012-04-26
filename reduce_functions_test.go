package sequences

import (
	R "reflect"
	"testing"
)

func TestReduceIndexableFunction(t *testing.T) {
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

	ConfirmReduce(indexable_function(func(i int) interface{} { return i }), 0, 55)
	ConfirmReduce(indexable_function(func(i int) interface{} { return i }), 10, 65)
}

func TestReduceBoolFunction(t *testing.T) {
	f := func(s ...bool) func(int) bool {
		return func(index int) bool {
			return s[index]
		} 
	}
	
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

	ConfirmReduce(f(), false, false)
	ConfirmReduce(f(), true, true)
	ConfirmReduce(f(false), false, false)
	ConfirmReduce(f(false), true, false)
	ConfirmReduce(f(true), false, false)
	ConfirmReduce(f(true), true, true)
	ConfirmReduce(f(false, true), true, false)
	ConfirmReduce(f(true, false), true, false)
	ConfirmReduce(f(true, true), true, true)
	ConfirmReduce(f(true, true, true), true, true)
	ConfirmReduce(f(true, true, false, true), true, false)
}


func TestReduceComplex64Function(t *testing.T) {
	f := func(s ...complex64) func(int) complex64 {
		return func(index int) complex64 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 1, 1)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 1, 1)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 1, 2)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceComplex128Function(t *testing.T) {
	f := func(s ...complex128) func(int) complex128 {
		return func(index int) complex128 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 1, 1)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 1, 1)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 1, 2)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceErrorFunction(t *testing.T) {
	f := func(s ...error) func(int) error {
		return func(index int) error {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), Error(0), Error(0))
	ConfirmReduce(f(), Error(8), Error(8))
	ConfirmReduce(f(Error(0)), Error(0), Error(0))
	ConfirmReduce(f(Error(0)), Error(8), Error(8))
	ConfirmReduce(f(Error(1)), Error(0), Error(1))
	ConfirmReduce(f(Error(1)), Error(8), Error(9))
	ConfirmReduce(f(Error(0), Error(1)), Error(0), Error(1))
	ConfirmReduce(f(Error(0), Error(1)), Error(8), Error(9))
	ConfirmReduce(f(Error(0), Error(1), Error(2)), Error(0), Error(3))
	ConfirmReduce(f(Error(0), Error(1), Error(2)), Error(8), Error(11))
	ConfirmReduce(f(Error(0), Error(1), Error(2), Error(4)), Error(0), Error(7))
	ConfirmReduce(f(Error(0), Error(1), Error(2), Error(4)), Error(8), Error(15))
}

func TestReduceFloat32Function(t *testing.T) {
	f := func(s ...float32) func(int) float32 {
		return func(index int) float32 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 1, 1)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 1, 1)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 1, 2)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceFloat64Function(t *testing.T) {
	f := func(s ...float64) func(int) float64 {
		return func(index int) float64 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 1, 1)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 1, 1)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 1, 2)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceIntFunction(t *testing.T) {
	f := func(s ...int) func(int) int {
		return func(index int) int {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceInt8Function(t *testing.T) {
	f := func(s ...int8) func(int) int8 {
		return func(index int) int8 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceInt16Function(t *testing.T) {
	f := func(s ...int16) func(int) int16 {
		return func(index int) int16 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceInt32Function(t *testing.T) {
	f := func(s ...int32) func(int) int32 {
		return func(index int) int32 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceInt64Function(t *testing.T) {
	f := func(s ...int64) func(int) int64 {
		return func(index int) int64 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceInterfaceFunction(t *testing.T) {
	f := func(s ...interface{}) func(int) interface{} {
		return func(index int) interface{} {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)

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
	
	ConfirmReduce(f(), false, false)
	ConfirmReduce(f(), true, true)
	ConfirmReduce(f(false), false, false)
	ConfirmReduce(f(false), true, false)
	ConfirmReduce(f(true), false, false)
	ConfirmReduce(f(true), true, true)
	ConfirmReduce(f(false, true), true, false)
	ConfirmReduce(f(true, false), true, false)
	ConfirmReduce(f(true, true), true, true)
	ConfirmReduce(f(true, true, true), true, true)
	ConfirmReduce(f(true, true, false, true), true, false)
}

func TestReduceStringFunction(t *testing.T) {
	f := func(s ...string) func(int) string {
		return func(index int) string {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), "", "")
	ConfirmReduce(f(), "Z", "Z")
	ConfirmReduce(f(""), "", "")
	ConfirmReduce(f(""), "Z", "Z")
	ConfirmReduce(f("A"), "", "A")
	ConfirmReduce(f("A"), "Z", "ZA")
	ConfirmReduce(f("A", "B"), "", "AB")
	ConfirmReduce(f("A", "B"), "Z", "ZAB")
	ConfirmReduce(f("A", "B", "C"), "", "ABC")
	ConfirmReduce(f("A", "B", "C"), "Z", "ZABC")
	ConfirmReduce(f("A", "B", "C", "D"), "", "ABCD")
	ConfirmReduce(f("A", "B", "C", "D"), "Z", "ZABCD")
}

func TestReduceUintFunction(t *testing.T) {
	f := func(s ...uint) func(int) uint {
		return func(index int) uint {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceUint8Function(t *testing.T) {
	f := func(s ...uint8) func(int) uint8 {
		return func(index int) uint8 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceUint16Function(t *testing.T) {
	f := func(s ...uint16) func(int) uint16 {
		return func(index int) uint16 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceUint32Function(t *testing.T) {
	f := func(s ...uint32) func(int) uint32 {
		return func(index int) uint32 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceUint64Function(t *testing.T) {
	f := func(s ...uint64) func(int) uint64 {
		return func(index int) uint64 {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceUintptrFunction(t *testing.T) {
	f := func(s ...uintptr) func(int) uintptr {
		return func(index int) uintptr {
			return s[index]
		} 
	}

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

	ConfirmReduce(f(), 0, 0)
	ConfirmReduce(f(), 10, 10)
	ConfirmReduce(f(0), 0, 0)
	ConfirmReduce(f(0), 10, 10)
	ConfirmReduce(f(1), 0, 1)
	ConfirmReduce(f(1), 10, 11)
	ConfirmReduce(f(1, 2), 0, 3)
	ConfirmReduce(f(1, 2), 10, 13)
	ConfirmReduce(f(1, 2, 3), 0, 6)
	ConfirmReduce(f(1, 2, 3), 10, 16)
	ConfirmReduce(f(1, 2, 3, 4), 0, 10)
	ConfirmReduce(f(1, 2, 3, 4), 10, 20)
}

func TestReduceRValueFunction(t *testing.T) {
	f := func(s ...interface{}) func(int) R.Value {
		return func(index int) R.Value {
			return R.ValueOf(s[index])
		} 
	}

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

	ConfirmReduceValueSeed := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, R.ValueOf(seed), f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed(f(), 0, 0)
	ConfirmReduceInterfaceSeed(f(), 10, 10)
	ConfirmReduceInterfaceSeed(f(0), 0, 0)
	ConfirmReduceInterfaceSeed(f(0), 10, 10)
	ConfirmReduceInterfaceSeed(f(1), 0, 1)
	ConfirmReduceInterfaceSeed(f(1), 10, 11)
	ConfirmReduceInterfaceSeed(f(1, 2), 0, 3)
	ConfirmReduceInterfaceSeed(f(1, 2), 10, 13)
	ConfirmReduceInterfaceSeed(f(1, 2, 3), 0, 6)
	ConfirmReduceInterfaceSeed(f(1, 2, 3), 10, 16)
	ConfirmReduceInterfaceSeed(f(1, 2, 3, 4), 0, 10)
	ConfirmReduceInterfaceSeed(f(1, 2, 3, 4), 10, 20)

	ConfirmReduceValueSeed(f(), 0, 0)
	ConfirmReduceValueSeed(f(), 10, 10)
	ConfirmReduceValueSeed(f(0), 0, 0)
	ConfirmReduceValueSeed(f(0), 10, 10)
	ConfirmReduceValueSeed(f(1), 0, 1)
	ConfirmReduceValueSeed(f(1), 10, 11)
	ConfirmReduceValueSeed(f(1, 2), 0, 3)
	ConfirmReduceValueSeed(f(1, 2), 10, 13)
	ConfirmReduceValueSeed(f(1, 2, 3), 0, 6)
	ConfirmReduceValueSeed(f(1, 2, 3), 10, 16)
	ConfirmReduceValueSeed(f(1, 2, 3, 4), 0, 10)
	ConfirmReduceValueSeed(f(1, 2, 3, 4), 10, 20)

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

	ConfirmReduceInterfaceSeed(f(), false, false)
	ConfirmReduceInterfaceSeed(f(), true, true)
	ConfirmReduceInterfaceSeed(f(false), false, false)
	ConfirmReduceInterfaceSeed(f(false), true, false)
	ConfirmReduceInterfaceSeed(f(true), false, false)
	ConfirmReduceInterfaceSeed(f(true), true, true)
	ConfirmReduceInterfaceSeed(f(false, true), true, false)
	ConfirmReduceInterfaceSeed(f(true, false), true, false)
	ConfirmReduceInterfaceSeed(f(true, true), true, true)
	ConfirmReduceInterfaceSeed(f(true, true, true), true, true)
	ConfirmReduceInterfaceSeed(f(true, true, false, true), true, false)
	
	ConfirmReduceValueSeed(f(), false, false)
	ConfirmReduceValueSeed(f(), true, true)
	ConfirmReduceValueSeed(f(false), false, false)
	ConfirmReduceValueSeed(f(false), true, false)
	ConfirmReduceValueSeed(f(true), false, false)
	ConfirmReduceValueSeed(f(true), true, true)
	ConfirmReduceValueSeed(f(false, true), true, false)
	ConfirmReduceValueSeed(f(true, false), true, false)
	ConfirmReduceValueSeed(f(true, true), true, true)
	ConfirmReduceValueSeed(f(true, true, true), true, true)
	ConfirmReduceValueSeed(f(true, true, false, true), true, false)
}

func TestReduceFunction(t *testing.T) {
	f := func(s ...UDT) func(int) UDT {
		return func(index int) UDT {
			return s[index]
		} 
	}

	iterators := []interface{}{
		func(seed, v interface{}) interface{} { return seed.(UDT) + v.(UDT) },
		func(seed interface{}, index int, v interface{}) interface{} { return seed.(UDT) + v.(UDT) },
		func(seed, index, v interface{}) interface{} { return seed.(UDT) + v.(UDT) },
		func(seed, v UDT) UDT { return seed + v },
		func(seed UDT, index int, v UDT) UDT { return seed + v },
		func(seed UDT, index interface{}, v UDT) UDT { return seed + v },
		func(seed, v R.Value) R.Value { return R.ValueOf(UDT(seed.Int() + v.Int())) },
		func(seed R.Value, index int, v R.Value) R.Value { return R.ValueOf(UDT(seed.Int() + v.Int())) },
		func(seed R.Value, index interface{}, v R.Value) R.Value { return R.ValueOf(UDT(seed.Int() + v.Int())) },
		func(seed R.Value, index, v R.Value) R.Value { return R.ValueOf(UDT(seed.Int() + v.Int())) },
	}

	ConfirmReduceInterfaceSeed := func(o interface{}, seed, result UDT) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(o func(int) UDT, seed, result UDT) {
		defer reportReductionError(t, o, seed)
		ov := R.ValueOf(o)
		for n, f := range iterators {
			if v := Reduce(ov, R.ValueOf(seed), f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", ov, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed(f(), 0, 0)
	ConfirmReduceInterfaceSeed(f(), 10, 10)
	ConfirmReduceInterfaceSeed(f(0), 0, 0)
	ConfirmReduceInterfaceSeed(f(0), 10, 10)
	ConfirmReduceInterfaceSeed(f(1), 0, 1)
	ConfirmReduceInterfaceSeed(f(1), 10, 11)
	ConfirmReduceInterfaceSeed(f(1, 2), 0, 3)
	ConfirmReduceInterfaceSeed(f(1, 2), 10, 13)
	ConfirmReduceInterfaceSeed(f(1, 2, 3), 0, 6)
	ConfirmReduceInterfaceSeed(f(1, 2, 3), 10, 16)
	ConfirmReduceInterfaceSeed(f(1, 2, 3, 4), 0, 10)
	ConfirmReduceInterfaceSeed(f(1, 2, 3, 4), 10, 20)

	ConfirmReduceValueSeed(f(), 0, 0)
	ConfirmReduceValueSeed(f(), 10, 10)
	ConfirmReduceValueSeed(f(0), 0, 0)
	ConfirmReduceValueSeed(f(0), 10, 10)
	ConfirmReduceValueSeed(f(1), 0, 1)
	ConfirmReduceValueSeed(f(1), 10, 11)
	ConfirmReduceValueSeed(f(1, 2), 0, 3)
	ConfirmReduceValueSeed(f(1, 2), 10, 13)
	ConfirmReduceValueSeed(f(1, 2, 3), 0, 6)
	ConfirmReduceValueSeed(f(1, 2, 3), 10, 16)
	ConfirmReduceValueSeed(f(1, 2, 3, 4), 0, 10)
	ConfirmReduceValueSeed(f(1, 2, 3, 4), 10, 20)
}