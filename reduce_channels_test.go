package sequences

import (
	R "reflect"
	"testing"
)

func TestReduceBoolChannel(t *testing.T) {
	f := func(s ...bool) func() chan bool {
		return func() (r chan bool) {
			r = make(chan bool)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := [] func() chan bool {
		f(),
		f(false),
		f(true),
		f(false, true),
		f(true, false),
		f(true, true),
		f(true, true, true),
		f(true, false, true, false),
	}

	ConfirmReduce := func(i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, false, false)
	ConfirmReduce(0, true, true)
	ConfirmReduce(1, false, false)
	ConfirmReduce(1, true, false)
	ConfirmReduce(2, false, false)
	ConfirmReduce(2, true, true)
	ConfirmReduce(3, true, false)
	ConfirmReduce(4, true, false)
	ConfirmReduce(5, true, true)
	ConfirmReduce(6, true, true)
	ConfirmReduce(7, true, false)
}

func TestReduceComplex64Channel(t *testing.T) {
	f := func(s ...complex64) func() chan complex64 {
		return func() (r chan complex64) {
			r = make(chan complex64)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan complex64{
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result complex64) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 1, 1)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 1, 1)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 1, 2)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}


func TestReduceComplex128Channel(t *testing.T) {
	f := func(s ...complex128) func() chan complex128 {
		return func() (r chan complex128) {
			r = make(chan complex128)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan complex128 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3), f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result complex128) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 1, 1)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 1, 1)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 1, 2)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceErrorChannel(t *testing.T) {
	f := func(s ...error) func() chan error {
		return func() (r chan error) {
			r = make(chan error)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan error {
		f(),
		f(Error(0)),
		f(Error(1)),
		f(Error(0), Error(1)),
		f(Error(0), Error(1), Error(2)),
		f(Error(0), Error(1), Error(2), Error(4)),
	}

	ConfirmReduce := func(i int, seed, result error) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, Error(0), Error(0))
	ConfirmReduce(0, Error(8), Error(8))
	ConfirmReduce(1, Error(0), Error(0))
	ConfirmReduce(1, Error(8), Error(8))
	ConfirmReduce(2, Error(0), Error(1))
	ConfirmReduce(2, Error(8), Error(9))
	ConfirmReduce(3, Error(0), Error(1))
	ConfirmReduce(3, Error(8), Error(9))
	ConfirmReduce(4, Error(0), Error(3))
	ConfirmReduce(4, Error(8), Error(11))
	ConfirmReduce(5, Error(0), Error(7))
	ConfirmReduce(5, Error(8), Error(15))
}

func TestReduceFloat32Channel(t *testing.T) {
	f := func(s ...float32) func() chan float32 {
		return func() (r chan float32) {
			r = make(chan float32)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan float32 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result float32) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 1, 1)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 1, 1)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 1, 2)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceFloat64Channel(t *testing.T) {
	f := func(s ...float64) func() chan float64 {
		return func() (r chan float64) {
			r = make(chan float64)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan float64 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result float64) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 1, 1)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 1, 1)
	ConfirmReduce(1, 0, 1)
	ConfirmReduce(1, 1, 2)
	ConfirmReduce(2, 0, 3)
	ConfirmReduce(2, 10, 13)
	ConfirmReduce(3, 0, 6)
	ConfirmReduce(4, 0, 10)
	ConfirmReduce(4, 10, 20)
}

func TestReduceIntChannel(t *testing.T) {
	f := func(s ...int) func() chan int {
		return func() (r chan int) {
			r = make(chan int)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan int {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceInt8Channel(t *testing.T) {
	f := func(s ...int8) func() chan int8 {
		return func() (r chan int8) {
			r = make(chan int8)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan int8 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result int8) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceInt16Channel(t *testing.T) {
	f := func(s ...int16) func() chan int16 {
		return func() (r chan int16) {
			r = make(chan int16)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan int16 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result int16) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceInt32Channel(t *testing.T) {
	f := func(s ...int32) func() chan int32 {
		return func() (r chan int32) {
			r = make(chan int32)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan int32 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result int32) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceInt64Channel(t *testing.T) {
	f := func(s ...int64) func() chan int64 {
		return func() (r chan int64) {
			r = make(chan int64)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan int64 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result int64) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceInterfaceChannel(t *testing.T) {
	f := func(s ...interface{}) func() chan interface{} {
		return func() (r chan interface{}) {
			r = make(chan interface{})
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan interface{} {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
   	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)

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

	funcs = []func() chan interface{} {
		f(),
		f(false),
		f(true),
		f(false, true),
		f(true, false),
		f(true, true),
		f(true, true, true),
		f(true, true, false, true),
	}

	ConfirmReduce(0, false, false)
	ConfirmReduce(0, true, true)
	ConfirmReduce(1, false, false)
	ConfirmReduce(1, true, false)
	ConfirmReduce(2, false, false)
	ConfirmReduce(2, true, true)
	ConfirmReduce(3, true, false)
	ConfirmReduce(4, true, false)
	ConfirmReduce(5, true, true)
	ConfirmReduce(6, true, true)
	ConfirmReduce(7, true, false)
}

func TestReduceStringChannel(t *testing.T) {
	f := func(s ...string) func() chan string {
		return func() (r chan string) {
			r = make(chan string)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan string {
		f(),
		f(""),
		f("A"),
		f("A", "B"),
		f("A", "B", "C"),
		f("A", "B", "C", "D"),
	}

	ConfirmReduce := func(i int, seed, result string) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, "", "")
	ConfirmReduce(0, "Z", "Z")
	ConfirmReduce(1, "", "")
	ConfirmReduce(1, "Z", "Z")
	ConfirmReduce(2, "", "A")
	ConfirmReduce(2, "Z", "ZA")
	ConfirmReduce(3, "", "AB")
	ConfirmReduce(3, "Z", "ZAB")
	ConfirmReduce(4, "", "ABC")
	ConfirmReduce(4, "Z", "ZABC")
	ConfirmReduce(5, "", "ABCD")
	ConfirmReduce(5, "Z", "ZABCD")
}

func TestReduceUintChannel(t *testing.T) {
	f := func(s ...uint) func() chan uint {
		return func() (r chan uint) {
			r = make(chan uint)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan uint {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceUint8Channel(t *testing.T) {
	f := func(s ...uint8) func() chan uint8 {
		return func() (r chan uint8) {
			r = make(chan uint8)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan uint8 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}


	ConfirmReduce := func(i int, seed, result uint8) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceUint16Channel(t *testing.T) {
	f := func(s ...uint16) func() chan uint16 {
		return func() (r chan uint16) {
			r = make(chan uint16)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan uint16 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint16) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceUint32Channel(t *testing.T) {
	f := func(s ...uint32) func() chan uint32 {
		return func() (r chan uint32) {
			r = make(chan uint32)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan uint32 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint32) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceUint64Channel(t *testing.T) {
	f := func(s ...uint64) func() chan uint64 {
		return func() (r chan uint64) {
			r = make(chan uint64)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan uint64 {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint64) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceUintptrChannel(t *testing.T) {
	f := func(s ...uintptr) func() chan uintptr {
		return func() (r chan uintptr) {
			r = make(chan uintptr)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan uintptr {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uintptr) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(0, 0, 0)
	ConfirmReduce(0, 10, 10)
	ConfirmReduce(1, 0, 0)
	ConfirmReduce(1, 10, 10)
	ConfirmReduce(2, 0, 1)
	ConfirmReduce(2, 10, 11)
	ConfirmReduce(3, 0, 3)
	ConfirmReduce(3, 10, 13)
	ConfirmReduce(4, 0, 6)
	ConfirmReduce(4, 10, 16)
	ConfirmReduce(5, 0, 10)
	ConfirmReduce(5, 10, 20)
}

func TestReduceRValueChannel(t *testing.T) {
	f := func(s ...interface{}) func() chan R.Value {
		return func() (r chan R.Value) {
			r = make(chan R.Value)
			go func() {
				for _, v := range s {
					r <- R.ValueOf(v)
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan R.Value {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduceInterfaceSeed := func(i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], R.ValueOf(seed), f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed(0, 0, 0)
	ConfirmReduceInterfaceSeed(0, 10, 10)
	ConfirmReduceInterfaceSeed(1, 0, 0)
	ConfirmReduceInterfaceSeed(1, 10, 10)
	ConfirmReduceInterfaceSeed(2, 0, 1)
	ConfirmReduceInterfaceSeed(2, 10, 11)
	ConfirmReduceInterfaceSeed(3, 0, 3)
	ConfirmReduceInterfaceSeed(3, 10, 13)
	ConfirmReduceInterfaceSeed(4, 0, 6)
	ConfirmReduceInterfaceSeed(4, 10, 16)
	ConfirmReduceInterfaceSeed(5, 0, 10)
	ConfirmReduceInterfaceSeed(5, 10, 20)

	ConfirmReduceValueSeed(0, 0, 0)
	ConfirmReduceValueSeed(0, 10, 10)
	ConfirmReduceValueSeed(1, 0, 0)
	ConfirmReduceValueSeed(1, 10, 10)
	ConfirmReduceValueSeed(2, 0, 1)
	ConfirmReduceValueSeed(2, 10, 11)
	ConfirmReduceValueSeed(3, 0, 3)
	ConfirmReduceValueSeed(3, 10, 13)
	ConfirmReduceValueSeed(4, 0, 6)
	ConfirmReduceValueSeed(4, 10, 16)
	ConfirmReduceValueSeed(5, 0, 10)
	ConfirmReduceValueSeed(5, 10, 20)

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

	funcs = []func() chan R.Value {
		f(),
		f(false),
		f(true),
		f(false, true),
		f(true, false),
		f(true, true, true),
		f(true, true, false, true),
	}

	ConfirmReduceInterfaceSeed(0, false, false)
	ConfirmReduceInterfaceSeed(0, true, true)
	ConfirmReduceInterfaceSeed(1, false, false)
	ConfirmReduceInterfaceSeed(1, true, false)
	ConfirmReduceInterfaceSeed(2, false, false)
	ConfirmReduceInterfaceSeed(2, true, true)
	ConfirmReduceInterfaceSeed(3, true, false)
	ConfirmReduceInterfaceSeed(3, true, false)
	ConfirmReduceInterfaceSeed(4, true, true)
	ConfirmReduceInterfaceSeed(5, true, true)
	ConfirmReduceInterfaceSeed(6, true, false)

	ConfirmReduceValueSeed(0, false, false)
	ConfirmReduceValueSeed(0, true, true)
	ConfirmReduceValueSeed(1, false, false)
	ConfirmReduceValueSeed(1, true, false)
	ConfirmReduceValueSeed(2, false, false)
	ConfirmReduceValueSeed(2, true, true)
	ConfirmReduceValueSeed(3, true, false)
	ConfirmReduceValueSeed(3, true, false)
	ConfirmReduceValueSeed(4, true, true)
	ConfirmReduceValueSeed(5, true, true)
	ConfirmReduceValueSeed(6, true, false)
}

func TestReduceChannel(t *testing.T) {
	f := func(s ...UDT) func() chan UDT {
		return func() (r chan UDT) {
			r = make(chan UDT)
			go func() {
				for _, v := range s {
					r <- v
				}
				close(r)
			}()
			return r
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

	funcs := []func() chan UDT {
		f(),
		f(0),
		f(1),
		f(1, 2),
		f(1, 2, 3),
		f(1, 2, 3, 4),
	}

	ConfirmReduceInterfaceSeed := func(i int, seed, result UDT) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(i int, seed, result UDT) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(R.ValueOf(funcs[i]), R.ValueOf(seed), f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed(0, 0, 0)
	ConfirmReduceInterfaceSeed(0, 10, 10)
	ConfirmReduceInterfaceSeed(1, 0, 0)
	ConfirmReduceInterfaceSeed(1, 10, 10)
	ConfirmReduceInterfaceSeed(2, 0, 1)
	ConfirmReduceInterfaceSeed(2, 10, 11)
	ConfirmReduceInterfaceSeed(3, 0, 3)
	ConfirmReduceInterfaceSeed(3, 10, 13)
	ConfirmReduceInterfaceSeed(4, 0, 6)
	ConfirmReduceInterfaceSeed(4, 10, 16)
	ConfirmReduceInterfaceSeed(5, 0, 10)
	ConfirmReduceInterfaceSeed(5, 10, 20)

	ConfirmReduceValueSeed(0, 0, 0)
	ConfirmReduceValueSeed(0, 10, 10)
	ConfirmReduceValueSeed(1, 0, 0)
	ConfirmReduceValueSeed(1, 10, 10)
	ConfirmReduceValueSeed(2, 0, 1)
	ConfirmReduceValueSeed(2, 10, 11)
	ConfirmReduceValueSeed(3, 0, 3)
	ConfirmReduceValueSeed(3, 10, 13)
	ConfirmReduceValueSeed(4, 0, 6)
	ConfirmReduceValueSeed(4, 10, 16)
	ConfirmReduceValueSeed(5, 0, 10)
	ConfirmReduceValueSeed(5, 10, 20)
}
