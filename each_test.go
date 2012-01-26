package sequences

import(
	R "reflect"
	"strconv"
	"testing"
)

func TestEachIterable(t *testing.T) {
	var count	int

	ConfirmEach := func(s iterable_slice, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != len(s):	t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	S := iterable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})
}

func TestEachIndexable(t *testing.T) {
	var count	int

	ConfirmEach := func(s Indexable, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != Len(s):	t.Fatalf("total iterations should be %v but are %v", Len(s), count)
		}
	}

	ConfirmVariadicEach := func(s Indexable, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachIndexable := func(s Indexable) {
		ConfirmEach(s, func(v interface{}) {
			if v != s.At(count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v interface{}) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s.At(i):					t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v interface{}) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v != s.At(count):				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v R.Value) {
			if v.Interface() != s.At(count) {
				t.Fatalf("element %v erroneously reported as %v", count, v.Interface())
			}
			count++
		})

		ConfirmEach(s, func(i int, v R.Value) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != s.At(i):		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k interface{}, v R.Value) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s.At(count):	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v R.Value) {
			switch {
			case k.Interface() != count:		t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s.At(count):	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...R.Value) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...int) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int) {
			if v != s.At(count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i, v int) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s.At(i):					t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}

	ConfirmEachIndexable(indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmEachIndexable(indexable_function(func(i int) interface{} { return i }))
}

func TestEachMappable(t *testing.T) {
	var count	int

	ConfirmEach := func(m Mappable, f interface{}) {
		count = 0
		switch {
		case !Each(m, f):		t.Fatalf("failed to perform iteration %v over %v", f, m)
		case count != Len(m):	t.Fatalf("total iterations should be %v but are %v", Len(m), count)
		}
	}

	ConfirmVariadicEach := func(m Mappable, f interface{}) {
		count = 0
		switch {
		case !Each(m, f):		t.Fatalf("failed to perform iteration %v over %v", f, m)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachMappable := func(m Mappable) {
		ConfirmEach(m, func(v interface{}) {
			count++
		})

		ConfirmEach(m, func(i int, v interface{}) {
			if v != m.Lookup(i) {
				t.Fatalf("element %v erroneously reported as %v", i, v)
			}
			count++
		})

		ConfirmEach(m, func(k, v interface{}) {
			if v != m.Lookup(k) {
				t.Fatalf("element %v erroneously reported as %v", k, v)
			}
			count++
		})

		ConfirmVariadicEach(m, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(m):				t.Fatalf("variadic slice %v erroneously passed as %v", m, v)
			}
			count++
		})

		ConfirmEach(m, func(v R.Value) {
			count++
		})

		ConfirmEach(m, func(i int, v R.Value) {
			if v.Interface() != m.Lookup(i) {
				t.Fatalf("element %v erroneously reported as %v", i, v)
			}
			count++
		})

		ConfirmEach(m, func(k interface{}, v R.Value) {
			if v.Interface() != m.Lookup(k) {
				t.Fatalf("element %v erroneously reported as %v", k, v)
			}
			count++
		})

		ConfirmEach(m, func(k, v R.Value) {
			if v.Interface() != m.Lookup(k.Interface()) {
				t.Fatalf("element %v erroneously reported as %v", k.Interface(), v.Interface())
			}
			count++
		})

		ConfirmVariadicEach(m, func(v ...R.Value) {
			switch {
			case count == 1:					t.Fatalf("variadic function erroneously called more than once")
			case len(v) != Len(m):				t.Fatalf("variadic slice %v erroneously passed as %v", m, v)
			}
			count++
		})

		ConfirmVariadicEach(m, func(v ...int) {
			switch {
			case count == 1:					t.Fatalf("variadic function erroneously called more than once")
			case len(v) != Len(m):				t.Fatalf("variadic slice %v erroneously passed as %v", m, v)
			}
			count++
		})

		ConfirmEach(m, func(v int) {
			count++
		})

		ConfirmEach(m, func(i, v int) {
			if v != m.Lookup(i) {
				t.Fatalf("element %v erroneously reported as %v", i, v)
			}
			count++
		})
	}

	ConfirmEachMappable(mappable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmEachMappable(mappable_function(func(i int) interface{} { return i }))
	ConfirmEachMappable(mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9})

	m := mappable_string_map{"0": 0, "1": 1, "2": 2, "3": 3}
	ConfirmEach(m, func(k string, v interface{}) {
		if v != m.Lookup(k) {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmEach(m, func(k string, v R.Value) {
		if v.Interface() != m.Lookup(k) {
			t.Fatalf("element %v erroneously reported as %v", k, v.Interface())
		}
		count++
	})
}

func TestEachSlice(t *testing.T) {
	var count	int

	ConfirmEach := func(s []int, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != len(s):	t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	ConfirmVariadicEach := func(s []int, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}


	ConfirmEachSlice := func(s []int) {
		ConfirmEach(s, func(v interface{}) {
			if v != s[count] {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v interface{}) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s[i]:						t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v interface{}) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v != s[count]:					t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v R.Value) {
			if v.Interface() != s[count] {
				t.Fatalf("element %v erroneously reported as %v", count, v.Interface())
			}
			count++
		})

		ConfirmEach(s, func(i int, v R.Value) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != s[i]:			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k interface{}, v R.Value) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s[count]:		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v R.Value) {
			switch {
			case k.Interface() != count:		t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s[count]:		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...R.Value) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...int) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int) {
			if v != s[count] {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i, v int) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s[i]:						t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}

	ConfirmEachSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestEachMap(t *testing.T) {
	var count	int

	ConfirmEach := func(m, f interface{}) {
		count = 0
		switch {
		case !Each(m, f):		t.Fatalf("failed to perform iteration %v over %v", f, m)
		case count != Len(m):	t.Fatalf("total iterations should be %v but are %v", Len(m), count)
		}
	}

	ConfirmVariadicEach := func(m, f interface{}) {
		count = 0
		switch {
		case !Each(m, f):		t.Fatalf("failed to perform iteration %v over %v", f, m)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}

	ConfirmEach(M, func(v interface{}) {
		count++
	})

	ConfirmEach(M, func(i int, v interface{}) {
		if v != M[i] {
			t.Fatalf("element %v erroneously reported as %v", i, v)
		}
		count++
	})

	ConfirmEach(M, func(k, v interface{}) {
		if v != M[k.(int)] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmVariadicEach(M, func(v ...interface{}) {
		switch {
		case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
		case len(v) != Len(M):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", M, v)
		}
		count++
	})

	ConfirmEach(M, func(v R.Value) {
		count++
	})

	ConfirmEach(M, func(i int, v R.Value) {
		if v.Interface() != M[i] {
			t.Fatalf("element %v erroneously reported as %v", i, v)
		}
		count++
	})

	ConfirmEach(M, func(k interface{}, v R.Value) {
		if v.Interface() != M[k.(int)] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmEach(M, func(k, v R.Value) {
		if v.Interface() != M[int(k.Int())] {
			t.Fatalf("element %v erroneously reported as %v", k.Interface(), v.Interface())
		}
		count++
	})

	ConfirmVariadicEach(M, func(v ...R.Value) {
		switch {
		case count == 1:					t.Fatalf("variadic function erroneously called more than once")
		case len(v) != Len(M):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", M, v)
		}
		count++
	})

	ConfirmVariadicEach(M, func(v ...int) {
		switch {
		case count == 1:					t.Fatalf("variadic function erroneously called more than once")
		case len(v) != Len(M):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", M, v)
		}
		count++
	})

	ConfirmEach(M, func(v int) {
		count++
	})

	ConfirmEach(M, func(i, v int) {
		if v != M[i] {
			t.Fatalf("element %v erroneously reported as %v", i, v)
		}
		count++
	})

	MS := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

	ConfirmEach(MS, func(k string, v interface{}) {
		if v != MS[k] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmEach(MS, func(k, v interface{}) {
		if v != MS[k.(string)] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})


	ConfirmEach(MS, func(k string, v R.Value) {
		if v.Interface() != MS[k] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmEach(MS, func(k interface{}, v R.Value) {
		if v.Interface() != MS[k.(string)] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})
}

func TestEachChannel(t *testing.T) {
	var index	int

	ConfirmEach := func(s []int, f interface{}) {
		var v		int
		var count	int

		C := make(chan int)
		go func() {
			for count, v = range s {
				C <- v
			}
			close(C)
		}()

		index = 0
		Each(C, f)
		if count != len(s) - 1 {
			t.Fatalf("total iterations should be %v but is %v", len(s) - 1, count)
		}
	}

	S := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	ConfirmEach(S, func(v interface{}) {
		if v != S[index] {
			t.Fatalf("index %v erroneously reported as %v", index, v)
		}
		index++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != index:		t.Fatalf("index %v erroneously reported as %v", index, i)
		case v != S[i]:			t.Fatalf("value %v erroneously reported as %v", S[9], v)
		}
		index++
	})
}

func TestEachFunction(t *testing.T) {
	F1 := func(v interface{}) (r interface{}, ok bool) {
		if v.(int) < 10 {
			r = v
			ok = true
		}
		return
	}

	count := 0
	Each(F1, func(v interface{}) {
		if v != count {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	Each(F1, func(i interface{}, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F1, func(i int, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	F2 := func(v int) (r int, ok bool) {
		if v < 10 {
			r = v
			ok = true
		}
		return
	}

	count = 0
	Each(F2, func(v int) {
		if v != count {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(F2, func(v interface{}) {
		if v != count {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	Each(F2, func(i interface{}, v int) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F2, func(i int, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F2, func(i int, v int) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	F3 := func(v int) (r string, ok bool) {
		if v < 10 {
			r = strconv.Itoa(v)
			ok = true
		}
		return
	}
	
	count = 0
	Each(F3, func(v string) {
		if v != strconv.Itoa(count) {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(F3, func(v interface{}) {
		if v != strconv.Itoa(count) {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})
	
	Each(F3, func(i interface{}, v string) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F3, func(i int, v string) {
		if strconv.Itoa(i) != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})
}