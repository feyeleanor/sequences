package sequences

import(
	R "reflect"
	"testing"
)

func TestEachPrimitive(t *testing.T) {
	ConfirmEach := func(s, f interface{}) {
		if count, e := Each(s, f); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		} else if l, _ := Len(s); count != l {
			t.Fatalf("total iterations should be %v but are %v", l, count)
		}
	}

	ConfirmVariadicEach := func(s, f interface{}) {
		switch count, e := Each(s, f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachBuiltin := func(s interface{}) {
		ConfirmEach(s, func(v interface{}) {})
		ConfirmEach(s, func(i int, v interface{}) {})
		ConfirmEach(s, func(k, v interface{}) {})
		ConfirmVariadicEach(s, func(v ...interface{}) {})
		ConfirmEach(s, func(v R.Value) {})
		ConfirmEach(s, func(i int, v R.Value) {})
		ConfirmEach(s, func(k interface{}, v R.Value) {})
		ConfirmEach(s, func(k, v R.Value) {})
		ConfirmVariadicEach(s, func(v ...R.Value) {})
	}

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...bool) {})
		ConfirmEach(s, func(v bool) {})
		ConfirmEach(s, func(i int, v bool) {})
	}([]bool{true, false, false, true, true})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...complex64) {})
		ConfirmEach(s, func(v complex64) {})
		ConfirmEach(s, func(i int, v complex64) {})
	}([]complex64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...complex128) {})
		ConfirmEach(s, func(v complex128) {})
		ConfirmEach(s, func(i int, v complex128) {})
	}([]complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...error) {})
		ConfirmEach(s, func(v error) {})
		ConfirmEach(s, func(i int, v error) {})
	}([]error{Error(0), Error(1), Error(2), Error(3), Error(4), Error(5), Error(6), Error(7), Error(8), Error(9)})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...float32) {})
		ConfirmEach(s, func(v float32) {})
		ConfirmEach(s, func(i int, v float32) {})
	}([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...float64) {})
		ConfirmEach(s, func(v float64) {})
		ConfirmEach(s, func(i int, v float64) {})
	}([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int) {})
		ConfirmEach(s, func(v int) {})
		ConfirmEach(s, func(i, v int) {})
	}([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int8) {})
		ConfirmEach(s, func(v int8) {})
		ConfirmEach(s, func(i int, v int8) {})
	}([]int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int16) {})
		ConfirmEach(s, func(v int16) {})
		ConfirmEach(s, func(i int, v int16) {})
	}([]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int32) {})
		ConfirmEach(s, func(v int32) {})
		ConfirmEach(s, func(i int, v int32) {})
	}([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int64) {})
		ConfirmEach(s, func(v int64) {})
		ConfirmEach(s, func(i int, v int64) {})
	}([]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	ConfirmEachBuiltin([]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...string) {})
		ConfirmEach(s, func(v string) {})
		ConfirmEach(s, func(i int, v string) {})
	}([]string{"A", "B", "C", "D", "E", "F"})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint) {})
		ConfirmEach(s, func(v uint) {})
		ConfirmEach(s, func(i int, v uint) {})
	}([]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint8) {})
		ConfirmEach(s, func(v uint8) {})
		ConfirmEach(s, func(i int, v uint8) {})
	}([]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint16) {})
		ConfirmEach(s, func(v uint16) {})
		ConfirmEach(s, func(i int, v uint16) {})
	}([]uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint32) {})
		ConfirmEach(s, func(v uint32) {})
		ConfirmEach(s, func(i int, v uint32) {})
	}([]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint64) {})
		ConfirmEach(s, func(v uint64) {})
		ConfirmEach(s, func(i int, v uint64) {})
	}([]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uintptr) {})
		ConfirmEach(s, func(v uintptr) {})
		ConfirmEach(s, func(i int, v uintptr) {})
	}([]uintptr{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s []R.Value) {
		ConfirmEach(s, func(v interface{}) {})
		ConfirmEach(s, func(i int, v interface{}) {})
		ConfirmEach(s, func(k, v interface{}) {})
		ConfirmVariadicEach(s, func(v ...interface{}) {})
		ConfirmEach(s, func(v R.Value) {})
		ConfirmEach(s, func(i int, v R.Value) {})
		ConfirmEach(s, func(k interface{}, v R.Value) {})
		ConfirmEach(s, func(k, v R.Value) {})
		ConfirmVariadicEach(s, func(v ...R.Value) {})
	}([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4), R.ValueOf(5), R.ValueOf(6), R.ValueOf(7), R.ValueOf(8), R.ValueOf(9)})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int) {})
		ConfirmEach(s, func(v int) {})
		ConfirmEach(s, func(i, v int) {})
	}(R.ValueOf([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func TestEachIterable(t *testing.T) {
	ConfirmEach := func(s iterable_slice, f interface{}) {
		switch count, e := Each(s, f); {
		case e != nil:
			t.Fatalf("iteration failed with error %v", e)
		case count != len(s):
			t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	ConfirmVariadicEach := func(s iterable_slice, f interface{}) {
		switch count, e := Each(s, f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	S := iterable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmEach(S, func(i interface{}) {})
	ConfirmEach(S, func(i int, v interface{}) {})
	ConfirmEach(S, func(k, v interface{}) {})
	ConfirmVariadicEach(S, func(i ...interface{}) {})
	ConfirmEach(S, func(i R.Value) {})
	ConfirmEach(S, func(i int, v R.Value) {})
	ConfirmEach(S, func(k interface{}, v R.Value) {})
	ConfirmEach(S, func(k, v R.Value) {})
	ConfirmVariadicEach(S, func(v ...R.Value) {})
}

func TestEachEnumerable(t *testing.T) {
	ConfirmEach := func(s enumerable_slice, f interface{}) {
		switch count, e := Each(s, f); {
		case e != nil:
			t.Fatalf("iteration failed with error %v", e)
		case count != len(s):
			t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	ConfirmVariadicEach := func(s enumerable_slice, f interface{}) {
		switch count, e := Each(s, f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	S := enumerable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmEach(S, func(i interface{}) {})
	ConfirmEach(S, func(i int, v interface{}) {})
	ConfirmEach(S, func(k, v interface{}) {})
	ConfirmVariadicEach(S, func(i ...interface{}) {})
	ConfirmEach(S, func(i R.Value) {})
	ConfirmEach(S, func(i int, v R.Value) {})
	ConfirmEach(S, func(k interface{}, v R.Value) {})
	ConfirmEach(S, func(k, v R.Value) {})
	ConfirmVariadicEach(S, func(v ...R.Value) {})
}

func TestEachIndexable(t *testing.T) {
	ConfirmEach := func(s Indexable, f interface{}) {
		if count, e := Each(s, f); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		} else if l, _ := Len(s); count != l {
			t.Fatalf("total iterations should be %v but are %v", l, count)
		}
	}

	ConfirmVariadicEach := func(s Indexable, f interface{}) {
		switch count, e := Each(s, f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachIndexable := func(S Indexable) {
		ConfirmEach(S, func(i interface{}) {})
		ConfirmEach(S, func(i int, v interface{}) {})
		ConfirmEach(S, func(k, v interface{}) {})
		ConfirmVariadicEach(S, func(i ...interface{}) {})
		ConfirmEach(S, func(i R.Value) {})
		ConfirmEach(S, func(i int, v R.Value) {})
		ConfirmEach(S, func(k interface{}, v R.Value) {})
		ConfirmEach(S, func(k, v R.Value) {})
		ConfirmVariadicEach(S, func(v ...R.Value) {})
	}

	ConfirmEachIndexable(indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmEachIndexable(indexable_function(func(i int) interface{} { return i }))
}

func TestEachMappable(t *testing.T) {
	t.Fatalf("Implement support for Mappables")
	ConfirmEach := func(m Mappable, f interface{}) {
		if count, e := Each(m, f); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		} else if l, _ := Len(m); count != l {
			t.Fatalf("total iterations should be %v but are %v", l, count)
		}
	}

	ConfirmVariadicEach := func(m Mappable, f interface{}) {
		switch count, e := Each(m, f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachMappable := func(M Mappable) {
		ConfirmEach(M, func(i interface{}) {})
		ConfirmEach(M, func(i int, v interface{}) {})
		ConfirmEach(M, func(k, v interface{}) {})
		ConfirmVariadicEach(M, func(i ...interface{}) {})
		ConfirmEach(M, func(i R.Value) {})
		ConfirmEach(M, func(i int, v R.Value) {})
		ConfirmEach(M, func(k interface{}, v R.Value) {})
		ConfirmEach(M, func(k, v R.Value) {})
		ConfirmVariadicEach(M, func(v ...R.Value) {})
	}

	ConfirmEachMappable(mappable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmEachMappable(mappable_function(func(i int) interface{} { return i }))
	ConfirmEachMappable(mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9})

	m := mappable_string_map{"0": 0, "1": 1, "2": 2, "3": 3}
	ConfirmEach(m, func(k string, v interface{}) {})
	ConfirmEach(m, func(k string, v R.Value) {})
}

func TestEachSlice(t *testing.T) {
	ConfirmEach := func(s []int, f interface{}) {
		switch count, e := Each(s, f); {
		case e != nil:
			t.Fatalf("iteration failed with error %v", e)
		case count != len(s):
			t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	ConfirmVariadicEach := func(s []int, f interface{}) {
		switch count, e := Each(s, f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachSlice := func(s []int) {
		ConfirmEach(s, func(v interface{}) {})
		ConfirmEach(s, func(i int, v interface{}) {})
		ConfirmEach(s, func(k, v interface{}) {})
		ConfirmVariadicEach(s, func(v ...interface{}) {})
		ConfirmEach(s, func(v R.Value) {})
		ConfirmEach(s, func(i int, v R.Value) {})
		ConfirmEach(s, func(k interface{}, v R.Value) {})
		ConfirmEach(s, func(k, v R.Value) {})
		ConfirmVariadicEach(s, func(v ...R.Value) {})
		ConfirmVariadicEach(s, func(v ...int) {})
		ConfirmEach(s, func(v int) {})
		ConfirmEach(s, func(i, v int) {})
	}
	ConfirmEachSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestEachMap(t *testing.T) {
	t.Fatalf("Implement support for Maps")
	ConfirmEach := func(m, f interface{}) {
		if count, e := Each(m, f); e != nil {
			t.Fatalf("iteration failed with error %v", e)
		} else if l, _ := Len(m); count != l {
			t.Fatalf("total iterations should be %v but are %v", l, count)
		}
	}

	ConfirmVariadicEach := func(m, f interface{}) {
		switch count, e := Each(m, f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	ConfirmEach(M, func(v interface{}) {})
	ConfirmEach(M, func(i int, v interface{}) {})
	ConfirmEach(M, func(k, v interface{}) {})
	ConfirmVariadicEach(M, func(v ...interface{}) {})
	ConfirmEach(M, func(v R.Value) {})
	ConfirmEach(M, func(i int, v R.Value) {})
	ConfirmEach(M, func(k interface{}, v R.Value) {})
	ConfirmEach(M, func(k, v R.Value) {})
	ConfirmVariadicEach(M, func(v ...R.Value) {})
	ConfirmVariadicEach(M, func(v ...int) {})
	ConfirmEach(M, func(v int) {})
	ConfirmEach(M, func(i, v int) {})

	MS := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	ConfirmEach(MS, func(k string, v interface{}) {})
	ConfirmEach(MS, func(k string, v R.Value) {})
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

	ConfirmEach := func(s []int, f interface{}) {
		switch count, e := Each(Generate(s), f); {
		case e != nil:
			t.Fatalf("iteration failed with error %v", e)
		case count != len(s):
			t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	ConfirmVariadicEach := func(s []int, f interface{}) {
		switch count, e := Each(Generate(s), f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	S := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	ConfirmEach(S, func(v interface{}) {})
	ConfirmEach(S, func(i int, v interface{}) {})
	ConfirmEach(S, func(i, v interface{}) {})
	ConfirmVariadicEach(S, func(v ...interface{}) {})
	ConfirmEach(S, func(v R.Value) {})
	ConfirmEach(S, func(i int, v R.Value) {})
	ConfirmEach(S, func(i interface{}, v R.Value) {})
	ConfirmEach(S, func(i, v R.Value) {})
	ConfirmVariadicEach(S, func(v ...R.Value) {})
	ConfirmVariadicEach(S, func(v ...int) {})
	ConfirmEach(S, func(v int) {})
	ConfirmEach(S, func(i, v int) {})
}

func TestEachFunction(t *testing.T) {
	limit := 10
	value := 0
	ConfirmEach := func(F, f interface{}) {
		value = 0
		switch i, e := Each(F, f); {
		case e != nil:
			t.Fatalf("iteration failed with error %v", e)
		case i != limit:
			t.Fatalf("total iterations should be %v but are %v", limit, i)
		}
	}

	ConfirmVariadicEach := func(F, f interface{}) {
		value = 0
		switch i, e := Each(F, f); {
		case e != nil:
			t.Fatalf("variadic iteration failed with error %v", e)
		case i != 1:
			t.Fatalf("total iterations should be 1 but are %v", i)
		}
	}

	ConfirmEachFunction := func(F interface{}) {
		ConfirmEach(F, func(v interface{}) {})
		ConfirmEach(F, func(i int, v interface{}) {})
		ConfirmEach(F, func(i, v interface{}) {})
		ConfirmVariadicEach(F, func(v ...interface{}) {})
		ConfirmEach(F, func(v R.Value) {})
		ConfirmEach(F, func(i int, v R.Value) {})
		ConfirmEach(F, func(i interface{}, v R.Value) {})
		ConfirmEach(F, func(i, v R.Value) {})
		ConfirmVariadicEach(F, func(v ...R.Value) {})
	}

	ConfirmEachIntFunction := func(F interface{}) {
		ConfirmEachFunction(F)
		ConfirmVariadicEach(F, func(v ...int) {})
		ConfirmEach(F, func(v int) {})
		ConfirmEach(F, func(i, v int) {})
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