package sequences

/*
import(
	R "reflect"
	"testing"
)

func TestEachPrimitive(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		}
	}()

	ConfirmEach := func(s, f interface{}) {
		Each(s, f)
	}

	ConfirmVariadicEach := func(s, f interface{}) {
		Each(s, f)
	}

	ConfirmEachBuiltin := func(s interface{}) {
		Each(s, func(v interface{}) {})
		Each(s, func(i int, v interface{}) {})
		Each(s, func(k, v interface{}) {})
		Each(s, func(v R.Value) {})
		Each(s, func(i int, v R.Value) {})
		Each(s, func(k interface{}, v R.Value) {})
		Each(s, func(k, v R.Value) {})
	}

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v bool) {})
		Each(s, func(i int, v bool) {})
	}([]bool{true, false, false, true, true})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v complex64) {})
		Each(s, func(i int, v complex64) {})
	}([]complex64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v complex128) {})
		Each(s, func(i int, v complex128) {})
	}([]complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v error) {})
		Each(s, func(i int, v error) {})
	}([]error{Error(0), Error(1), Error(2), Error(3), Error(4), Error(5), Error(6), Error(7), Error(8), Error(9)})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v float32) {})
		Each(s, func(i int, v float32) {})
	}([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v float64) {})
		Each(s, func(i int, v float64) {})
	}([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v int) {})
		Each(s, func(i, v int) {})
	}([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v int8) {})
		Each(s, func(i int, v int8) {})
	}([]int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v int16) {})
		Each(s, func(i int, v int16) {})
	}([]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v int32) {})
		Each(s, func(i int, v int32) {})
	}([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v int64) {})
		Each(s, func(i int, v int64) {})
	}([]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	ConfirmEachBuiltin([]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v string) {})
		Each(s, func(i int, v string) {})
	}([]string{"A", "B", "C", "D", "E", "F"})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v uint) {})
		Each(s, func(i int, v uint) {})
	}([]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v uint8) {})
		Each(s, func(i int, v uint8) {})
	}([]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v uint16) {})
		Each(s, func(i int, v uint16) {})
	}([]uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v uint32) {})
		Each(s, func(i int, v uint32) {})
	}([]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v uint64) {})
		Each(s, func(i int, v uint64) {})
	}([]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v uintptr) {})
		Each(s, func(i int, v uintptr) {})
	}([]uintptr{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s []R.Value) {
		Each(s, func(v interface{}) {})
		Each(s, func(i int, v interface{}) {})
		Each(s, func(k, v interface{}) {})
		Each(s, func(v R.Value) {})
		Each(s, func(i int, v R.Value) {})
		Each(s, func(k interface{}, v R.Value) {})
		Each(s, func(k, v R.Value) {})
	}([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4), R.ValueOf(5), R.ValueOf(6), R.ValueOf(7), R.ValueOf(8), R.ValueOf(9)})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		Each(s, func(v int) {})
		Each(s, func(i, v int) {})
	}(R.ValueOf([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func TestEachEnumerable(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		}
	}()

	S := enumerable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Each(S, func(i interface{}) {})
	Each(S, func(i int, v interface{}) {})
	Each(S, func(k, v interface{}) {})
	Each(S, func(i R.Value) {})
	Each(S, func(i int, v R.Value) {})
	Each(S, func(k interface{}, v R.Value) {})
	Each(S, func(k, v R.Value) {})
}

func TestEachEnumerable(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		}
	}()

	S := enumerable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Each(S, func(i interface{}) {})
	Each(S, func(i int, v interface{}) {})
	Each(S, func(k, v interface{}) {})
	Each(S, func(i R.Value) {})
	Each(S, func(i int, v R.Value) {})
	Each(S, func(k interface{}, v R.Value) {})
	Each(S, func(k, v R.Value) {})
}

func TestEachIndexable(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		}
	}()

	ConfirmEachIndexable := func(S Indexable) {
		Each(S, func(i interface{}) {})
		Each(S, func(i int, v interface{}) {})
		Each(S, func(k, v interface{}) {})
		Each(S, func(i R.Value) {})
		Each(S, func(i int, v R.Value) {})
		Each(S, func(k interface{}, v R.Value) {})
		Each(S, func(k, v R.Value) {})
	}

	ConfirmEachIndexable(indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmEachIndexable(indexable_function(func(i int) interface{} { return i }))
}

func TestEachSlice(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		}
	}()

	ConfirmEachSlice := func(s []int) {
		Each(s, func(v interface{}) {})
		Each(s, func(i int, v interface{}) {})
		Each(s, func(k, v interface{}) {})
		Each(s, func(v R.Value) {})
		Each(s, func(i int, v R.Value) {})
		Each(s, func(k interface{}, v R.Value) {})
		Each(s, func(k, v R.Value) {})
		Each(s, func(v int) {})
		Each(s, func(i, v int) {})
	}
	ConfirmEachSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestEachChannel(t *testing.T) {
	Generate := func(s []int) (c chan int) {
		c = make(chan int)
		go func() {
			for _, v := range s {
				c <- v
			}
			close(c)
		}()
		return c
	}

	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		}
	}()

	S := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	Each(S, func(v interface{}) {})
	Each(S, func(i int, v interface{}) {})
	Each(S, func(i, v interface{}) {})
	Each(S, func(v R.Value) {})
	Each(S, func(i int, v R.Value) {})
	Each(S, func(i interface{}, v R.Value) {})
	Each(S, func(i, v R.Value) {})
	Each(S, func(v int) {})
	Each(S, func(i, v int) {})
}

func TestEachFunction(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		}
	}()

	limit := 10
	value := 0
	ConfirmEachFunction := func(F interface{}) {
		Each(F, func(v interface{}) {})
		Each(F, func(i int, v interface{}) {})
		Each(F, func(i, v interface{}) {})
		Each(F, func(v R.Value) {})
		Each(F, func(i int, v R.Value) {})
		Each(F, func(i interface{}, v R.Value) {})
		Each(F, func(i, v R.Value) {})
	}

	ConfirmEachIntFunction := func(F interface{}) {
		EachFunction(F)
		Each(F, func(v int) {})
		Each(F, func(i, v int) {})
	}

	ConfirmEachFunction(func() (r interface{}, finished bool) {
		r = value
		if finished = value > 9; !finished {
			value++
		}
		return
	})

	ConfirmEachIntFunction(func() (r int, finished bool) {
		r = value
		if finished = value > 9; !finished {
			value++
		}
		return
	})
}
*/
