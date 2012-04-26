package sequences

import(
	R "reflect"
	"testing"
)

func sliceOffset(s interface{}, count, span int) (r int) {
	switch {
	case span == 0:
		return 0
	case span < 0:
		r = Len(s) -1 + (span * count)
	default:
		r = span * count
	}
	return
}

func spannedLen(s interface{}, span int) (r int) {
	r, _ = CountBy(s, span, func(interface{}) bool { return true })
	return
}

func ConfirmSliceEachBy(t *testing.T, s interface{}, span int, f interface{}) {
	defer func() {
		if e := recover(); e != nil {
			panic(e)
			t.Fatalf("%T span %v: iteration failed with error %v", s, span, e)
		}
	}()
	if count := EachBy(s, span, f); count != spannedLen(s, span) {
		panic(span)
		t.Fatalf("%T (%v): total iterations should be %v but are %v", s, span, spannedLen(s, span), count)
	}
}

func TestEnumerateSlice(t *testing.T) {
	ConfirmEach := func(s interface{}, span int) {
		count := 0
		ConfirmSliceEachBy(t, s, span, func(v interface{}) {
			offset := sliceOffset(s, count, span)
			if x, ok := v.(R.Value); ok {
				v = x.Interface()
			}
			x := AtOffset(s, offset)
			if y, ok := x.(R.Value); ok {
				x = y.Interface()
			}
			if x != v {
				panic(offset)
			}
			count++
		})

		count = 0
		ConfirmSliceEachBy(t, s, span, func(i int, v interface{}) {
			if x, ok := v.(R.Value); ok {
				v = x.Interface()
			}
			offset := sliceOffset(s, count, span)
			x := AtOffset(s, offset)
			if y, ok := x.(R.Value); ok {
				x = y.Interface()
			}
			switch {
			case i != offset:
				t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
			case v != x:
				t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})

		count = 0
		ConfirmSliceEachBy(t, s, span, func(k, v interface{}) {
			if x, ok := v.(R.Value); ok {
				v = x.Interface()
			}
			offset := sliceOffset(s, count, span)
			x := AtOffset(s, offset)
			if y, ok := x.(R.Value); ok {
				x = y.Interface()
			}
			switch {
			case k != offset:
				t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, k)
			case v != x:
				t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})
		
		count = 0
		ConfirmSliceEachBy(t, s, span, func(v R.Value) {
			offset := sliceOffset(s, count, span)
			x := AtOffset(s, offset)
			if y, ok := x.(R.Value); ok {
				x = y.Interface()
			}
			if v.Interface() != x {
				t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v.Interface())
			}
			count++
		})

		count = 0
		ConfirmSliceEachBy(t, s, span, func(i int, v R.Value) {
			offset := sliceOffset(s, count, span)
			x := AtOffset(s, offset)
			if y, ok := x.(R.Value); ok {
				x = y.Interface()
			}
			switch {
			case i != offset:
				t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
			case v.Interface() != x:
				t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, count, v)
			}
			count++
		})

		count = 0
		ConfirmSliceEachBy(t, s, span, func(k interface{}, v R.Value) {
			offset := sliceOffset(s, count, span)
			x := AtOffset(s, offset)
			if y, ok := x.(R.Value); ok {
				x = y.Interface()
			}
			switch {
			case k != offset:
				t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, k)
			case v.Interface() != x:
				t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})

		count = 0
		ConfirmSliceEachBy(t, s, span, func(k, v R.Value) {
			offset := sliceOffset(s, count, span)
			x := AtOffset(s, offset)
			if y, ok := x.(R.Value); ok {
				x = y.Interface()
			}
			switch {
			case k.Interface() != offset:
				t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, k)
			case v.Interface() != x:
				t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})

		switch s := s.(type) {
		case []interface{}, []R.Value, Enumerable, Indexable:
			//	already covered by previous test cases
		case []bool:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v bool) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v bool) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []complex64:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v complex64) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v complex64) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []complex128:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v complex128) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v complex128) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []error:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v error) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v error) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []float32:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v float32) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v float32) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []float64:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v float64) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v float64) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []int:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i, v int) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []int8:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int8) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int8) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int) {
				offset := sliceOffset(s, count, span)
				if int8(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case int8(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []int16:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int16) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int16) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int) {
				offset := sliceOffset(s, count, span)
				if int16(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case int16(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []int32:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int32) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int32) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int) {
				offset := sliceOffset(s, count, span)
				if int32(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case int32(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []int64:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int64) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int64) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v int) {
				offset := sliceOffset(s, count, span)
				if int64(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case int64(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []string:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v string) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v string) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []uint:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []uint8:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint8) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint8) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint) {
				offset := sliceOffset(s, count, span)
				if uint8(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case uint8(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []uint16:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint16) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint16) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint) {
				offset := sliceOffset(s, count, span)
				if uint16(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case uint16(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []uint32:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint32) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint32) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint) {
				offset := sliceOffset(s, count, span)
				if uint32(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case uint32(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []uint64:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint64) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint64) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint) {
				offset := sliceOffset(s, count, span)
				if uint64(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case uint64(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case []uintptr:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uintptr) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uintptr) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v uint) {
				offset := sliceOffset(s, count, span)
				if uintptr(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v uint) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case uintptr(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		case R.Value:
		case []UDT:
			count = 0
			ConfirmSliceEachBy(t, s, span, func(v UDT) {
				offset := sliceOffset(s, count, span)
				if v != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v UDT) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case v != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(v UDT) {
				offset := sliceOffset(s, count, span)
				if uintptr(v) != AtOffset(s, offset) {
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, offset, v)
				}
				count++
			})

			count = 0
			ConfirmSliceEachBy(t, s, span, func(i int, v int) {
				switch offset := sliceOffset(s, count, span); {
				case i != offset:
					t.Fatalf("%T (%v): index %v erroneously reported as %v", s, span, offset, i)
				case uintptr(v) != AtOffset(s, i):
					t.Fatalf("%T (%v): element %v erroneously reported as %v", s, span, i, v)
				}
				count++
			})
		default:
			t.Fatalf("unknown sequence type %T", s)
		}
	}

	sequences := []interface{}{
		enumerable_slice{0, 1, 2, 3, 4},
		indexable_slice{0, 1, 2, 3, 4},
		indexable_function(func(i int) interface{} { return i }),
		[]bool{true, false, false, true, true},
		[]complex64{0, 1, 2, 3, 4},
		[]complex128{0, 1, 2, 3, 4},
		[]error{Error(0), Error(1), Error(2), Error(3), Error(4)},
		[]float32{0, 1, 2, 3, 4},
		[]float64{0, 1, 2, 3, 4},
		[]int{0, 1, 2, 3, 4},
		[]int8{0, 1, 2, 3, 4},
		[]int16{0, 1, 2, 3, 4},
		[]int32{0, 1, 2, 3, 4},
		[]int64{0, 1, 2, 3, 4},
		[]interface{}{0, 1, 2, 3, 4},
		[]string{"A", "B", "C", "D", "E"},
		[]uint{0, 1, 2, 3, 4},
		[]uint8{0, 1, 2, 3, 4},
		[]uint16{0, 1, 2, 3, 4},
		[]uint32{0, 1, 2, 3, 4},
		[]uint64{0, 1, 2, 3, 4},
		[]uintptr{0, 1, 2, 3, 4},
		[]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)},
		R.ValueOf([]int{0, 1, 2, 3, 4}),
		[]UDT{0, 1, 2, 3, 4},
	}

	//	Enumerate each item of the sequence
	for i, sequence := range sequences {
		t.Logf("%v: Write failing test case for span < 0", i)
		t.Logf("%v: Write test case for span 0", i)
		t.Logf("%v: Write test cases for channels", i)
		for j := Len(sequence) * 3; j > 0; j-- {
			ConfirmEach(sequence, j)
		}
	}
}