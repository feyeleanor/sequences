package sequences

import R "reflect"

func forEachBoolReceived(enum *Enumerator) (r func(func(bool))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(bool)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(bool)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
	return
}

func eachBoolChannel(enum *Enumerator, c chan bool) {
	var cursor int

	forEachReceived := forEachBoolReceived(enum)
	switch f := enum.f.(type) {
	case func(bool):
		forEachReceived(func(v bool) {
			f(v)
		})
	case func(int, bool):
		forEachReceived(func(v bool) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, bool):
		forEachReceived(func(v bool) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v bool) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v bool) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v bool) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v bool) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v bool) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v bool) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v bool) {
			f(R.ValueOf(cursor), R.ValueOf(v))
		})
	case chan bool:
		forEachReceived(func(v bool) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v bool) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v bool) {
			f <- R.ValueOf(v)
		})
	case []chan bool:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachComplex64Received(enum *Enumerator) (r func(func(complex64))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		forEachReceived = func(f func(complex64)) {
			for x := range c {
				f(x)
			}
		}
	default:
		forEachReceived = func(f func(complex64)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
	return
}

func eachComplex64Channel(enum *Enumerator, c chan complex64) {
	var cursor int

	forEachReceived := forEachComplex64Received(enum)
	switch f := enum.f.(type) {
	case func(complex64):
		forEachReceived(func(v complex64) {
			f(v)
		})
	case func(int, complex64):
		forEachReceived(func(v complex64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, complex64):
		forEachReceived(func(v complex64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v complex64) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v complex64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v complex64) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v complex64) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v complex64) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v complex64) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v complex64) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan complex64:
		forEachReceived(func(v complex64) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v complex64) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v complex64) {
			f <- R.ValueOf(v)
		})
	case []chan complex64:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachComplex128Received(enum *Enumerator) (r func(func(complex128))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		forEachReceived = func(f func(complex128)) {
			for x := range c {
				f(x)
			}
		}
	default:
		forEachReceived = func(f func(complex128)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
	return
}

func eachComplex128Channel(enum *Enumerator, c chan complex128) {
	var cursor int

	forEachReceived := forEachComplex128Received(enum)
	switch f := enum.f.(type) {
	case func(complex128):
		forEachReceived(func(v complex128) {
			f(v)
		})
	case func(int, complex128):
		forEachReceived(func(v complex128) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, complex128):
		forEachReceived(func(v complex128) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v complex128) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v complex128) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v complex128) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v complex128) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v complex128) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v complex128) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v complex128) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan complex128:
		forEachReceived(func(v complex128) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v complex128) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v complex128) {
			f <- R.ValueOf(v)
		})
	case []chan complex128:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachErrorReceived(enum *Enumerator) (r func(func(error))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(error)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(error)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}

}

func eachErrorChannel(enum *Enumerator, c chan error) {
	var cursor int

	forEachReceived := forEachErrorReceived(enum)
	switch f := enum.f.(type) {
	case func(error):
		forEachReceived(func(v error) {
			f(v)
		})
	case func(int, error):
		forEachReceived(func(v error) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, error):
		forEachReceived(func(v error) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v error) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v error) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v error) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v error) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v error) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v error) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v error) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan error:
		forEachReceived(func(v error) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v error) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v error) {
			f <- R.ValueOf(v)
		})
	case []chan Error:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachFloat32Received(enum *Enumerator) (r func(func(float32))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(float32)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(float32)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}

}

func eachFloat32Channel(enum *Enumerator, c chan float32) {
	var cursor int

	forEachReceived := forEachFloat32Received(enum)
	switch f := enum.f.(type) {
	case func(float32):
		forEachReceived(func(v float32) {
			f(v)
		})
	case func(int, float32):
		forEachReceived(func(v float32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, float32):
		forEachReceived(func(v float32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v float32) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v float32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v float32) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v float32) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v float32) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v float32) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v float32) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan float32:
		forEachReceived(func(v float32) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v float32) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v float32) {
			f <- R.ValueOf(v)
		})
	case []chan float32:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachFloat64Received(enum *Enumerator) (r func(func(float64))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(float64)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(float64)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}

}

func eachFloat64Channel(enum *Enumerator, c chan float64) {
	var cursor int

	forEachReceived := forEachFloat64Received(enum)
	switch f := enum.f.(type) {
	case func(float64):
		forEachReceived(func(v float64) {
			f(v)
		})
	case func(int, float64):
		forEachReceived(func(v float64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, float64):
		forEachReceived(func(v float64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v float64) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v float64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v float64) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v float64) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v float64) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v float64) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v float64) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan float64:
		forEachReceived(func(v float64) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v float64) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v float64) {
			f <- R.ValueOf(v)
		})
	case []chan float64:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachIntReceived(enum *Enumerator) (r func(func(int))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(int)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(int)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachIntChannel(enum *Enumerator, c chan int) {
	var cursor int

	forEachReceived := forEachIntReceived(enum)
	switch f := enum.f.(type) {
	case func(int):
		forEachReceived(func(v int) {
			f(v)
		})
	case func(int, int):
		forEachReceived(func(v int) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, int):
		forEachReceived(func(v int) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v int) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v int) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan int:
		forEachReceived(func(v int) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v int) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v int) {
			f <- R.ValueOf(v)
		})
	case []chan int:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachInt8Received(enum *Enumerator) (r func(func(int8))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(int8)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(int8)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachInt8Channel(enum *Enumerator, c chan int8) {
	var cursor int

	forEachReceived := forEachInt8Received(enum)
	switch f := enum.f.(type) {
	case func(int8):
		forEachReceived(func(v int8) {
			f(v)
		})
	case func(int, int8):
		forEachReceived(func(v int8) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, int8):
		forEachReceived(func(v int8) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v int8) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int8) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int8) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v int8) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int8) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int8) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int8) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan int8:
		forEachReceived(func(v int8) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v int8) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v int8) {
			f <- R.ValueOf(v)
		})
	case []chan int8:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachInt16Received(enum *Enumerator) (r func(func(int16))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(int16)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(int16)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachInt16Channel(enum *Enumerator, c chan int16) {
	var cursor int

	forEachReceived := forEachInt16Received(enum)
	switch f := enum.f.(type) {
	case func(int16):
		forEachReceived(func(v int16) {
			f(v)
		})
	case func(int, int16):
		forEachReceived(func(v int16) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, int16):
		forEachReceived(func(v int16) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v int16) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int16) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int16) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v int16) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int16) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int16) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int16) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan int16:
		forEachReceived(func(v int16) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v int16) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v int16) {
			f <- R.ValueOf(v)
		})
	case []chan int16:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachInt32Received(enum *Enumerator) (r func(func(int32))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(int32)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(int32)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachInt32Channel(enum *Enumerator, c chan int32) {
	var cursor int

	forEachReceived := forEachInt32Received(enum)
	switch f := enum.f.(type) {
	case func(int32):
		forEachReceived(func(v int32) {
			f(v)
		})
	case func(int, int32):
		forEachReceived(func(v int32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, int32):
		forEachReceived(func(v int32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v int32) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int32) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v int32) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int32) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int32) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int32) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan int32:
		forEachReceived(func(v int32) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v int32) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v int32) {
			f <- R.ValueOf(v)
		})
	case []chan int32:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachInt64Received(enum *Enumerator) (r func(func(int64))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(int64)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(int64)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachInt64Channel(enum *Enumerator, c chan int64) {
	var cursor int

	forEachReceived := forEachInt64Received(enum)
	switch f := enum.f.(type) {
	case func(int64):
		forEachReceived(func(v int64) {
			f(v)
		})
	case func(int, int64):
		forEachReceived(func(v int64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, int64):
		forEachReceived(func(v int64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v int64) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int64) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v int64) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int64) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int64) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int64) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan int64:
		forEachReceived(func(v int64) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v int64) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v int64) {
			f <- R.ValueOf(v)
		})
	case []chan int64:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachInterfaceReceived(enum *Enumerator) (r func(func(interface{}))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(interface{})) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(interface{})) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachInterfaceChannel(enum *Enumerator, c chan interface{}) {
	var cursor int

	forEachReceived := forEachInterfaceReceived(enum)
	switch f := enum.f.(type) {
	case func(interface{}):
		forEachReceived(func(v interface{}) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v interface{}) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v interface{}) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v interface{}) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v interface{}) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v interface{}) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v interface{}) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan interface{}:
		forEachReceived(func(v interface{}) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v interface{}) {
			f <- R.ValueOf(v)
		})
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachStringReceived(enum *Enumerator) (r func(func(string))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(string)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(string)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachStringChannel(enum *Enumerator, c chan string) {
	var cursor int

	forEachReceived := forEachStringReceived(enum)
	switch f := enum.f.(type) {
	case func(string):
		forEachReceived(func(v string) {
			f(v)
		})
	case func(int, string):
		forEachReceived(func(v string) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, string):
		forEachReceived(func(v string) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v string) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v string) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v string) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v string) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v string) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v string) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v string) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan string:
		forEachReceived(func(v string) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v string) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v string) {
			f <- R.ValueOf(v)
		})
	case []chan string:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachUintReceived(enum *Enumerator) (r func(func(uint))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(uint)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(uint)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachUintChannel(enum *Enumerator, c chan uint) {
	var cursor int

	forEachReceived := forEachUintReceived(enum)
	switch f := enum.f.(type) {
	case func(uint):
		forEachReceived(func(v uint) {
			f(v)
		})
	case func(int, uint):
		forEachReceived(func(v uint) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, uint):
		forEachReceived(func(v uint) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v uint) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v uint) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan uint:
		forEachReceived(func(v uint) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v uint) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v uint) {
			f <- R.ValueOf(v)
		})
	case []chan uint:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachUint8Received(enum *Enumerator) (r func(func(uint8))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(uint8)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(uint8)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachUint8Channel(enum *Enumerator, c chan uint8) {
	var cursor int

	forEachReceived := forEachUint8Received(enum)
	switch f := enum.f.(type) {
	case func(uint8):
		forEachReceived(func(v uint8) {
			f(v)
		})
	case func(int, uint8):
		forEachReceived(func(v uint8) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, uint8):
		forEachReceived(func(v uint8) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v uint8) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint8) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint8) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v uint8) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint8) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint8) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint8) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan uint8:
		forEachReceived(func(v uint8) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v uint8) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v uint8) {
			f <- R.ValueOf(v)
		})
	case []chan uint8:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachUint16Received(enum *Enumerator) (r func(func(uint16))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(uint16)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(uint16)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachUint16Channel(enum *Enumerator, c chan uint16) {
	var cursor int

	forEachReceived := forEachUint16Received(enum)
	switch f := enum.f.(type) {
	case func(uint16):
		forEachReceived(func(v uint16) {
			f(v)
		})
	case func(int, uint16):
		forEachReceived(func(v uint16) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, uint16):
		forEachReceived(func(v uint16) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v uint16) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint16) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint16) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v uint16) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint16) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint16) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint16) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan uint16:
		forEachReceived(func(v uint16) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v uint16) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v uint16) {
			f <- R.ValueOf(v)
		})
	case []chan uint16:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachUint32Received(enum *Enumerator) (r func(func(uint32))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(uint32)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(uint32)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachUint32Channel(enum *Enumerator, c chan uint32) {
	var cursor int

	forEachReceived := forEachUint32Received(enum)
	switch f := enum.f.(type) {
	case func(uint32):
		forEachReceived(func(v uint32) {
			f(v)
		})
	case func(int, uint32):
		forEachReceived(func(v uint32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, uint32):
		forEachReceived(func(v uint32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v uint32) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint32) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint32) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v uint32) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint32) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint32) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint32) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan uint32:
		forEachReceived(func(v uint32) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v uint32) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v uint32) {
			f <- R.ValueOf(v)
		})
	case []chan uint32:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachUint64Received(enum *Enumerator) (r func(func(uint64))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(uint64)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(uint64)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachUint64Channel(enum *Enumerator, c chan uint64) {
	var cursor int

	forEachReceived := forEachUint64Received(enum)
	switch f := enum.f.(type) {
	case func(uint64):
		forEachReceived(func(v uint64) {
			f(v)
		})
	case func(int, uint64):
		forEachReceived(func(v uint64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, uint64):
		forEachReceived(func(v uint64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v uint64) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint64) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint64) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v uint64) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint64) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint64) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint64) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan uint64:
		forEachReceived(func(v uint64) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v uint64) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v uint64) {
			f <- R.ValueOf(v)
		})
	case []chan uint64:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachUintptrReceived(enum *Enumerator) (r func(func(uintptr))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(uintptr)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(uintptr)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachUintptrChannel(enum *Enumerator, c chan uintptr) {
	var cursor int

	forEachReceived := forEachUintptrReceived(enum)
	switch f := enum.f.(type) {
	case func(uintptr):
		forEachReceived(func(v uintptr) {
			f(v)
		})
	case func(int, uintptr):
		forEachReceived(func(v uintptr) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, uintptr):
		forEachReceived(func(v uintptr) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}):
		forEachReceived(func(v uintptr) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uintptr) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uintptr) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v uintptr) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uintptr) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uintptr) {
			f(cursor, R.ValueOf(v))
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uintptr) {
			f(R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
	case chan uintptr:
		forEachReceived(func(v uintptr) {
			f <- v
		})
	case chan interface{}:
		forEachReceived(func(v uintptr) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v uintptr) {
			f <- R.ValueOf(v)
		})
	case []chan uintptr:
		panic(UNHANDLED_ITERATOR)
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachRValueReceived(enum *Enumerator) (r func(func(R.Value))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(R.Value)) {
			for x := range c {
				f(x)
			}
		}
	default:
		r = func(f func(R.Value)) {
			offset := enum.Span
			for x := range c {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachRValueChannel(enum *Enumerator, c chan R.Value) {
	var cursor int

	forEachReceived := forEachRValueReceived(enum)
	switch f := enum.f.(type) {
	case func(interface{}):
		forEachReceived(func(v R.Value) {
			f(v.Interface())
		})
	case func(int, interface{}):
		forEachReceived(func(v R.Value) {
			f(cursor, v.Interface())
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v R.Value) {
			f(cursor, v.Interface())
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v R.Value) {
			f(v)
		})
	case func(int, R.Value):
		forEachReceived(func(v R.Value) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v R.Value) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v R.Value) {
			f(R.ValueOf(cursor), v)
			cursor++
		})
	case chan interface{}:
		forEachReceived(func(v R.Value) {
			f <- v.Interface()
		})
	case chan R.Value:
		forEachReceived(func(v R.Value) {
			f <- v
		})
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func forEachChannelReceived(enum *Enumerator) (r func(func(R.Value))) {
	switch {
	case enum.Span < 1:
		panic(ASCENDING_SEQUENCE)
	case enum.Span == 1:
		r = func(f func(R.Value)) {
			for x, open = c.Recv(); open; x, open = c.Recv() {
				f(x)
			}
		}
	default:
		r = func(f func(R.Value)) {
			offset := enum.Span
			for x, open = c.Recv(); open; x, open = c.Recv() {
				if offset == 1 {
					f(x)
					offset = enum.Span
				} else {
					offset--
				}
			}
		}
	}
}

func eachChannel(enum *Enumerator, c R.Value) {
	var cursor int

	forEachReceived := forEachChannelReceived(enum)
	switch f := enum.f.(type) {
	case func(interface{}):
		forEachReceived(func(v R.Value) {
			f(v.Interface())
		})
	case func(int, interface{}):
		forEachReceived(func(v R.Value) {
			f(cursor, v.Interface())
			cursor++
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v R.Value) {
			f(cursor, v.Interface())
			cursor++
		})
	case func(R.Value):
		forEachReceived(func(v R.Value) {
			f(v)
		})
	case func(int, R.Value):
		forEachReceived(func(v R.Value) {
			f(cursor, v)
			cursor++
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v R.Value) {
			f(cursor, v)
			cursor++
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v R.Value) {
			f(R.ValueOf(cursor), v)
			cursor++
		})
	case chan interface{}:
		forEachReceived(func(v R.Value) {
			f <- v.Interface()
		})
	case chan R.Value:
		forEachReceived(func(v R.Value) {
			f <- v
		})
	case []chan interface{}:
		panic(UNHANDLED_ITERATOR)
	case []chan R.Value:
		panic(UNHANDLED_ITERATOR)
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 1: //	f(v)
					p := make([]R.Value, 1, 1)
					forEachReceived(func(v R.Value) {
						p[0] = v
						f.Call(p)
					})
				case 2: //	f(i, v)
					p := make([]R.Value, 2, 2)
					forEachReceived(func(v R.Value) {
						p[0], p[1] = R.ValueOf(cursor), v
						f.Call(p)
						cursor++
					})
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		} else {
			panic(UNHANDLED_ITERATOR)
		}
	}
}
