package sequences

import (
	R "reflect"
	"testing"
)

func TestReduceBoolChannel(t *testing.T) {
	seq := func(s ...bool) func() chan bool {
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

	funcs := []func() chan bool{
		seq(),
		seq(false),
		seq(true),
		seq(false, true),
		seq(true, false),
		seq(true, true),
		seq(true, true, true),
		seq(true, false, true, false),
	}

	ConfirmReduce := func(i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceBoolIterators() {
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
	seq := func(s ...complex64) func() chan complex64 {
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

	funcs := []func() chan complex64{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result complex64) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceComplex64Iterators() {
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
	seq := func(s ...complex128) func() chan complex128 {
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

	funcs := []func() chan complex128{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result complex128) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceComplex128Iterators() {
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
	seq := func(s ...error) func() chan error {
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

	funcs := []func() chan error{
		seq(),
		seq(Error(0)),
		seq(Error(1)),
		seq(Error(0), Error(1)),
		seq(Error(0), Error(1), Error(2)),
		seq(Error(0), Error(1), Error(2), Error(4)),
	}

	ConfirmReduce := func(i int, seed, result error) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceErrorIterators() {
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
	seq := func(s ...float32) func() chan float32 {
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

	funcs := []func() chan float32{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result float32) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceFloat32Iterators() {
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
	seq := func(s ...float64) func() chan float64 {
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

	funcs := []func() chan float64{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result float64) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceFloat64Iterators() {
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
	seq := func(s ...int) func() chan int {
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

	funcs := []func() chan int{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceIntIterators() {
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
	seq := func(s ...int8) func() chan int8 {
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

	funcs := []func() chan int8{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result int8) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceInt8Iterators() {
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
	seq := func(s ...int16) func() chan int16 {
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

	funcs := []func() chan int16{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result int16) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceInt16Iterators() {
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
	seq := func(s ...int32) func() chan int32 {
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

	funcs := []func() chan int32{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result int32) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceInt32Iterators() {
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
	seq := func(s ...int64) func() chan int64 {
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

	funcs := []func() chan int64{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result int64) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceInt64Iterators() {
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
	seq := func(s ...interface{}) func() chan interface{} {
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

	funcs := []func() chan interface{}{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(iterators []interface{}, i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(ReduceInterfaceIteratorsInt(), 0, 0, 0)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 0, 10, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 1, 0, 0)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 1, 10, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 2, 0, 1)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 2, 10, 11)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 3, 0, 3)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 3, 10, 13)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 4, 0, 6)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 4, 10, 16)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 5, 0, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), 5, 10, 20)

	funcs = []func() chan interface{}{
		f(),
		f(false),
		f(true),
		f(false, true),
		f(true, false),
		f(true, true),
		f(true, true, true),
		f(true, true, false, true),
	}

	ConfirmReduce(ReduceInterfaceIteratorsBool(), 0, false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 0, true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 1, false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 1, true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 2, false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 2, true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 3, true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 4, true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 5, true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 6, true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), 7, true, false)
}

func TestReduceStringChannel(t *testing.T) {
	seq := func(s ...string) func() chan string {
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

	funcs := []func() chan string{
		seq(),
		seq(""),
		seq("A"),
		seq("A", "B"),
		seq("A", "B", "C"),
		seq("A", "B", "C", "D"),
	}

	ConfirmReduce := func(i int, seed, result string) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceStringIterators() {
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
	seq := func(s ...uint) func() chan uint {
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

	funcs := []func() chan uint{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceUintIterators() {
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
	seq := func(s ...uint8) func() chan uint8 {
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

	funcs := []func() chan uint8{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint8) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceUint8Iterators() {
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
	seq := func(s ...uint16) func() chan uint16 {
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

	funcs := []func() chan uint16{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint16) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceUint16Iterators() {
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
	seq := func(s ...uint32) func() chan uint32 {
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

	funcs := []func() chan uint32{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint32) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceUint32Iterators() {
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
	seq := func(s ...uint64) func() chan uint64 {
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

	funcs := []func() chan uint64{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uint64) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceUint64Iterators() {
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
	seq := func(s ...uintptr) func() chan uintptr {
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

	funcs := []func() chan uintptr{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduce := func(i int, seed, result uintptr) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceUintptrIterators() {
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
	seq := func(s ...interface{}) func() chan R.Value {
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

	funcs := []func() chan R.Value{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduceInterfaceSeed := func(iterators []interface{}, i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(iterators []interface{}, i int, seed, result interface{}) {
		defer reportReductionError(t, i, seed)
		for n, f := range iterators {
			if v := Reduce(funcs[i], R.ValueOf(seed), f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 0, 0, 0)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 0, 10, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 1, 0, 0)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 1, 10, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 2, 0, 1)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 2, 10, 11)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 3, 0, 3)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 3, 10, 13)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 4, 0, 6)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 4, 10, 16)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 5, 0, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), 5, 10, 20)

	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 0, 0, 0)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 0, 10, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 1, 0, 0)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 1, 10, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 2, 0, 1)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 2, 10, 11)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 3, 0, 3)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 3, 10, 13)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 4, 0, 6)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 4, 10, 16)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 5, 0, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), 5, 10, 20)

	funcs = []func() chan R.Value{
		seq(),
		seq(false),
		seq(true),
		seq(false, true),
		seq(true, false),
		seq(true, true, true),
		seq(true, true, false, true),
	}

	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 0, false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 0, true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 1, false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 1, true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 2, false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 2, true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 3, true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 3, true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 4, true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 5, true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), 6, true, false)

	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 0, false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 0, true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 1, false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 1, true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 2, false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 2, true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 3, true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 3, true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 4, true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 5, true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), 6, true, false)
}

func TestReduceChannel(t *testing.T) {
	seq := func(s ...UDT) func() chan UDT {
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

	funcs := []func() chan UDT{
		seq(),
		seq(0),
		seq(1),
		seq(1, 2),
		seq(1, 2, 3),
		seq(1, 2, 3, 4),
	}

	ConfirmReduceInterfaceSeed := func(i int, seed, result UDT) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceRValueChannel() {
			if v := Reduce(funcs[i], seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", i, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(i int, seed, result UDT) {
		defer reportReductionError(t, i, seed)
		for n, f := range ReduceRValueChannel() {
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
