package sequences

import R "reflect"

func eachBoolSlice(enum *Enumerator, seq []bool) (i int) {
	switch f := enum.f.(type) {
	case func(bool):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(bool) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, bool):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, bool) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, bool):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, bool) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan bool:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan bool:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachComplex64Slice(enum *Enumerator, seq []complex64) (i int) {
	switch f := enum.f.(type) {
	case func(complex64):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(complex64) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, complex64):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, complex64) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, complex64):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, complex64) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan complex64:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan complex64:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachComplex128Slice(enum *Enumerator, seq []complex128) (i int) {
	switch f := enum.f.(type) {
	case func(complex128):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(complex128) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, complex128):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, complex128) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, complex128):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, complex128) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan complex128:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan complex128:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachErrorSlice(enum *Enumerator, seq []error) (i int) {
	switch f := enum.f.(type) {
	case func(error):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(error) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, error):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, error) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, error):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, error) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan error:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan error:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachFloat32Slice(enum *Enumerator, seq []float32) (i int) {
	switch f := enum.f.(type) {
	case func(float32):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(float32) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, float32):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, float32) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, float32):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, float32) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan float32:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan float32:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachFloat64Slice(enum *Enumerator, seq []float64) (i int) {
	switch f := enum.f.(type) {
	case func(float64):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(float64) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, float64):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, float64) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, float64):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, float64) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan float64:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan float64:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachIntSlice(enum *Enumerator, seq []int) (i int) {
	switch f := enum.f.(type) {
	case func(int):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(int) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt8Slice(enum *Enumerator, seq []int8) (i int) {
	switch f := enum.f.(type) {
	case func(int8):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(int8) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int8):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int8) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int8):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int8) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int8:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int8:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt16Slice(enum *Enumerator, seq []int16) (i int) {
	switch f := enum.f.(type) {
	case func(int16):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(int16) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int16):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int16) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int16):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int16) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int16:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int16:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt32Slice(enum *Enumerator, seq []int32) (i int) {
	switch f := enum.f.(type) {
	case func(int32):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(int32) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int32):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int32) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int32):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int32) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int32:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int32:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt64Slice(enum *Enumerator, seq []int64) (i int) {
	switch f := enum.f.(type) {
	case func(int64):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(int64) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int64):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int64) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int64):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int64) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int64:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int64:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInterfaceSlice(enum *Enumerator, seq []interface{}) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachStringSlice(enum *Enumerator, seq []string) (i int) {
	switch f := enum.f.(type) {
	case func(string):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(string) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, string):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, string) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, string):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, string) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan string:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan string:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUintSlice(enum *Enumerator, seq []uint) (i int) {
	switch f := enum.f.(type) {
	case func(uint):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint8Slice(enum *Enumerator, seq []uint8) (i int) {
	switch f := enum.f.(type) {
	case func(uint8):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint8) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint8):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint8) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint8):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint8) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint8:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint8:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint16Slice(enum *Enumerator, seq []uint16) (i int) {
	switch f := enum.f.(type) {
	case func(uint16):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint16) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint16):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint16) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint16):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint16) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint16:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint16:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint32Slice(enum *Enumerator, seq []uint32) (i int) {
	switch f := enum.f.(type) {
	case func(uint32):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint32) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint32):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint32) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint32):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint32) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint32:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint32:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint64Slice(enum *Enumerator, seq []uint64) (i int) {
	switch f := enum.f.(type) {
	case func(uint64):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint64) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint64):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint64) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint64):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint64) bool :
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint64:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint64:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUintptrSlice(enum *Enumerator, seq []uintptr) (i int) {
	switch f := enum.f.(type) {
	case func(uintptr):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(uintptr) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uintptr):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uintptr) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uintptr):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uintptr) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uintptr:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uintptr:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachRValueSlice(enum *Enumerator, seq []R.Value) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(seq[cursor].Interface())
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor].Interface())
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor].Interface())
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor].Interface())
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor].Interface())
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor].Interface())
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(seq[cursor])
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), seq[cursor])
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), seq[cursor])
		})
	case chan interface{}:
		i = enum.each(func(cursor int) {
			f <- seq[cursor].Interface()
		})
	case chan R.Value:
		i = enum.each(func(cursor int) {
			f <- seq[cursor]
		})
	case []chan interface{}:
		i = enum.each(func(cursor int) {
			v := seq[cursor].Interface()
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.each(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachSlice(enum *Enumerator, s R.Value) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(s.Index(cursor).Interface())
		})
	case func(interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(s.Index(cursor).Interface())
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, s.Index(cursor).Interface())
		})
	case func(int, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, s.Index(cursor).Interface())
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, s.Index(cursor).Interface())
		})
	case func(interface{}, interface{}) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, s.Index(cursor).Interface())
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(s.Index(cursor))
		})
	case func(R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(s.Index(cursor))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, s.Index(cursor))
		})
	case func(int, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, s.Index(cursor))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, s.Index(cursor))
		})
	case func(interface{}, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(cursor, s.Index(cursor))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), s.Index(cursor))
		})
	case func(R.Value, R.Value) bool:
		i = enum.each(func(cursor int) bool {
			return f(R.ValueOf(cursor), s.Index(cursor))
		})
	default:
		switch f := R.ValueOf(f); f.Kind() {
		case R.Func:
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 1:				//	f(v)
					if t.NumOut() == 0 {
						p := []R.Value{}
						i = enum.each(func(cursor int) {
							p[0] = s.Index(cursor)
							f.Call(p)
						})
					} else {
						p := []R.Value{}
						i = enum.each(func(cursor int) bool {
							p[0] = s.Index(cursor)
							return f.Call(p)[0].Bool()
						})
					}
				case 2:				//	f(i, v)
					if t.NumOut() == 0 {
						p := []R.Value{}
						i = enum.each(func(cursor int) {
							p[0], p[1] = R.ValueOf(cursor), s.Index(cursor)
							f.Call(p)
						})
					} else {
						p := []R.Value{}
						i = enum.each(func(cursor int) bool {
							p[0], p[1] = R.ValueOf(cursor), s.Index(cursor)
							return f.Call(p)[0].Bool()
						})
					}
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		case R.Chan:
			i = enum.each(func(cursor int) {
				f.Send(s.Index(cursor))
			})
		case R.Slice:
			if f.Type().Elem().Kind() == R.Chan {
				n := f.Len()
				i = enum.each(func(cursor int) {
					for i := n; i > 0; i-- {
						f.Index(i).Send(s.Index(cursor))
					}
				})
			} else {
				panic(UNHANDLED_ITERATOR)
			}
		default:
			panic(UNHANDLED_ITERATOR)
		}
	}
	return
}