package sequences

import R "reflect"

func reduceBoolSlice(enum *Enumerator, seq []bool) (r bool) {
	switch f := enum.f.(type) {
	case func(bool, bool) bool:
		r = enum.seed.(bool)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(bool, int, bool) bool:
		r = enum.seed.(bool)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(bool, interface{}, bool) bool:
		r = enum.seed.(bool)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(bool)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(bool)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(bool)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = rv.Bool()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.Bool()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.Bool()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = rv.Bool()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

/*
func reduceComplex64Slice(enum *Enumerator, seq []complex64) (i int) {
	switch f := enum.f.(type) {
	case func(complex64):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(complex64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, complex64):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, complex64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, complex64):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, complex64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan complex64:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan complex64:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceComplex128Slice(enum *Enumerator, seq []complex128) (i int) {
	switch f := enum.f.(type) {
	case func(complex128):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(complex128) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, complex128):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, complex128) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, complex128):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, complex128) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan complex128:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan complex128:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceErrorSlice(enum *Enumerator, seq []error) (i int) {
	switch f := enum.f.(type) {
	case func(error):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(error) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, error):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, error) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, error):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, error) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan error:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan error:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceFloat32Slice(enum *Enumerator, seq []float32) (i int) {
	switch f := enum.f.(type) {
	case func(float32):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(float32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, float32):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, float32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, float32):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, float32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan float32:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan float32:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceFloat64Slice(enum *Enumerator, seq []float64) (i int) {
	switch f := enum.f.(type) {
	case func(float64):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(float64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, float64):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, float64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, float64):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, float64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan float64:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan float64:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceIntSlice(enum *Enumerator, seq []int) (i int) {
	switch f := enum.f.(type) {
	case func(int):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(int) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceInt8Slice(enum *Enumerator, seq []int8) (i int) {
	switch f := enum.f.(type) {
	case func(int8):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(int8) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int8):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int8) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int8):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int8) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int8:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int8:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceInt16Slice(enum *Enumerator, seq []int16) (i int) {
	switch f := enum.f.(type) {
	case func(int16):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(int16) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int16):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int16) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int16):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int16) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int16:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int16:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceInt32Slice(enum *Enumerator, seq []int32) (i int) {
	switch f := enum.f.(type) {
	case func(int32):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(int32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int32):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int32):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int32:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int32:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceInt64Slice(enum *Enumerator, seq []int64) (i int) {
	switch f := enum.f.(type) {
	case func(int64):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(int64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, int64):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, int64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, int64):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, int64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan int64:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan int64:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceInterfaceSlice(enum *Enumerator, seq []interface{}) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceStringSlice(enum *Enumerator, seq []string) (i int) {
	switch f := enum.f.(type) {
	case func(string):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(string) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, string):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, string) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, string):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, string) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan string:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan string:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceUintSlice(enum *Enumerator, seq []uint) (i int) {
	switch f := enum.f.(type) {
	case func(uint):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceUint8Slice(enum *Enumerator, seq []uint8) (i int) {
	switch f := enum.f.(type) {
	case func(uint8):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint8) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint8):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint8) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint8):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint8) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint8:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint8:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceUint16Slice(enum *Enumerator, seq []uint16) (i int) {
	switch f := enum.f.(type) {
	case func(uint16):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint16) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint16):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint16) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint16):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint16) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint16:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint16:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceUint32Slice(enum *Enumerator, seq []uint32) (i int) {
	switch f := enum.f.(type) {
	case func(uint32):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint32):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint32):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint32) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint32:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint32:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceUint64Slice(enum *Enumerator, seq []uint64) (i int) {
	switch f := enum.f.(type) {
	case func(uint64):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(uint64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uint64):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uint64) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uint64):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uint64) bool :
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uint64:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uint64:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceUintptrSlice(enum *Enumerator, seq []uintptr) (i int) {
	switch f := enum.f.(type) {
	case func(uintptr):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(uintptr) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, uintptr):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, uintptr) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, uintptr):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, uintptr) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(seq[cursor]))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
	case chan uintptr:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- R.ValueOf(seq[cursor])
		})
	case []chan uintptr:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor]
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceRValueSlice(enum *Enumerator, seq []R.Value) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor].Interface())
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor].Interface())
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor].Interface())
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor].Interface())
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor].Interface())
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor].Interface())
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(seq[cursor])
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(seq[cursor])
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, seq[cursor])
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, seq[cursor])
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), seq[cursor])
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), seq[cursor])
		})
	case chan interface{}:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor].Interface()
		})
	case chan R.Value:
		i = enum.reduce(func(cursor int) {
			f <- seq[cursor]
		})
	case []chan interface{}:
		i = enum.reduce(func(cursor int) {
			v := seq[cursor].Interface()
			for _, c := range f {
				c <- v
			}
		})
	case []chan R.Value:
		i = enum.reduce(func(cursor int) {
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

func reduceSlice(enum *Enumerator, s R.Value) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(s.Index(cursor).Interface())
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(s.Index(cursor).Interface())
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, s.Index(cursor).Interface())
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, s.Index(cursor).Interface())
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, s.Index(cursor).Interface())
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, s.Index(cursor).Interface())
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(s.Index(cursor))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(s.Index(cursor))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, s.Index(cursor))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, s.Index(cursor))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, s.Index(cursor))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, s.Index(cursor))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), s.Index(cursor))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
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
						i = enum.reduce(func(cursor int) {
							p[0] = s.Index(cursor)
							f.Call(p)
						})
					} else {
						p := []R.Value{}
						i = enum.reduce(func(cursor int) bool {
							p[0] = s.Index(cursor)
							return f.Call(p)[0].Bool()
						})
					}
				case 2:				//	f(i, v)
					if t.NumOut() == 0 {
						p := []R.Value{}
						i = enum.reduce(func(cursor int) {
							p[0], p[1] = R.ValueOf(cursor), s.Index(cursor)
							f.Call(p)
						})
					} else {
						p := []R.Value{}
						i = enum.reduce(func(cursor int) bool {
							p[0], p[1] = R.ValueOf(cursor), s.Index(cursor)
							return f.Call(p)[0].Bool()
						})
					}
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		case R.Chan:
			i = enum.reduce(func(cursor int) {
				f.Send(s.Index(cursor))
			})
		case R.Slice:
			if f.Type().Elem().Kind() == R.Chan {
				n := f.Len()
				i = enum.reduce(func(cursor int) {
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
*/