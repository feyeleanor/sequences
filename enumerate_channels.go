package sequences

import R "reflect"

func eachBoolChannel(enum *Enumerator, seq chan bool) (i int) {
	var forEachReceived	func(func(bool))
	var x				bool
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(bool)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(bool)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(bool):
		forEachReceived(func(v bool) {
			f(v)
		})
	case func(int, bool):
		forEachReceived(func(v bool) {
			f(enum.cursor, v)
		})
	case func(interface{}, bool):
		forEachReceived(func(v bool) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v bool) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v bool) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v bool) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v bool) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v bool) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v bool) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v bool) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
/*		l := len(f) - 1
		forEachReceived(func(v bool) {
//			go func() {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
//			}()
		})
*/	case []chan interface{}:
/*		l := len(f) - 1
		forEachReceived(func(v bool) {
			for n := l; n > 0; n-- {
				f[n] <- v
			}
		})
*/	case []chan R.Value:
/*		l := len(f) - 1
		forEachReceived(func(v bool) {
			for n := l; n > 0; n-- {
				f[n] <- R.ValueOf(v)
			}
		})
*/	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachComplex64Channel(enum *Enumerator, seq chan complex64) (i int) {
	var forEachReceived	func(func(complex64))
	var x				complex64
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(complex64)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(complex64)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(complex64):
		forEachReceived(func(v complex64) {
			f(v)
		})
	case func(int, complex64):
		forEachReceived(func(v complex64) {
			f(enum.cursor, v)
		})
	case func(interface{}, complex64):
		forEachReceived(func(v complex64) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v complex64) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v complex64) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v complex64) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v complex64) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v complex64) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v complex64) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v complex64) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan complex64:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachComplex128Channel(enum *Enumerator, seq chan complex128) (i int) {
	var forEachReceived	func(func(complex128))
	var x				complex128
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(complex128)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(complex128)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(complex128):
		forEachReceived(func(v complex128) {
			f(v)
		})
	case func(int, complex128):
		forEachReceived(func(v complex128) {
			f(enum.cursor, v)
		})
	case func(interface{}, complex128):
		forEachReceived(func(v complex128) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v complex128) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v complex128) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v complex128) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v complex128) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v complex128) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v complex128) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v complex128) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan complex128:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachErrorChannel(enum *Enumerator, seq chan error) (i int) {
	var forEachReceived	func(func(error))
	var x				error
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(error)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(error)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(error):
		forEachReceived(func(v error) {
			f(v)
		})
	case func(int, error):
		forEachReceived(func(v error) {
			f(enum.cursor, v)
		})
	case func(interface{}, error):
		forEachReceived(func(v error) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v error) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v error) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v error) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v error) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v error) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v error) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v error) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan error:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachFloat32Channel(enum *Enumerator, seq chan float32) (i int) {
	var forEachReceived	func(func(float32))
	var x				float32
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(float32)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(float32)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(float32):
		forEachReceived(func(v float32) {
			f(v)
		})
	case func(int, float32):
		forEachReceived(func(v float32) {
			f(enum.cursor, v)
		})
	case func(interface{}, float32):
		forEachReceived(func(v float32) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v float32) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v float32) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v float32) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v float32) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v float32) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v float32) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v float32) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan float32:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachFloat64Channel(enum *Enumerator, seq chan float64) (i int) {
	var forEachReceived	func(func(float64))
	var x				float64
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(float64)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(float64)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(float64):
		forEachReceived(func(v float64) {
			f(v)
		})
	case func(int, float64):
		forEachReceived(func(v float64) {
			f(enum.cursor, v)
		})
	case func(interface{}, float64):
		forEachReceived(func(v float64) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v float64) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v float64) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v float64) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v float64) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v float64) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v float64) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v float64) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan float64:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachIntChannel(enum *Enumerator, seq chan int) (i int) {
	var forEachReceived	func(func(int))
	var x				int
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(int):
		forEachReceived(func(v int) {
			f(v)
		})
	case func(int, int):
		forEachReceived(func(v int) {
			f(enum.cursor, v)
		})
	case func(interface{}, int):
		forEachReceived(func(v int) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v int) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v int) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan int:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt8Channel(enum *Enumerator, seq chan int8) (i int) {
	var forEachReceived	func(func(int8))
	var x				int8
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int8)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int8)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(int8):
		forEachReceived(func(v int8) {
			f(v)
		})
	case func(int, int8):
		forEachReceived(func(v int8) {
			f(enum.cursor, v)
		})
	case func(interface{}, int8):
		forEachReceived(func(v int8) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v int8) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int8) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int8) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v int8) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int8) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int8) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int8) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan int8:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt16Channel(enum *Enumerator, seq chan int16) (i int) {
	var forEachReceived	func(func(int16))
	var x				int16
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int16)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int16)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(int16):
		forEachReceived(func(v int16) {
			f(v)
		})
	case func(int, int16):
		forEachReceived(func(v int16) {
			f(enum.cursor, v)
		})
	case func(interface{}, int16):
		forEachReceived(func(v int16) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v int16) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int16) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int16) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v int16) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int16) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int16) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int16) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan int16:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt32Channel(enum *Enumerator, seq chan int32) (i int) {
	var forEachReceived	func(func(int32))
	var x				int32
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int32)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int32)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(int32):
		forEachReceived(func(v int32) {
			f(v)
		})
	case func(int, int32):
		forEachReceived(func(v int32) {
			f(enum.cursor, v)
		})
	case func(interface{}, int32):
		forEachReceived(func(v int32) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v int32) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int32) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int32) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v int32) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int32) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int32) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int32) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan int32:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt64Channel(enum *Enumerator, seq chan int64) (i int) {
	var forEachReceived	func(func(int64))
	var x				int64
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int64)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int64)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(int64):
		forEachReceived(func(v int64) {
			f(v)
		})
	case func(int, int64):
		forEachReceived(func(v int64) {
			f(enum.cursor, v)
		})
	case func(interface{}, int64):
		forEachReceived(func(v int64) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v int64) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v int64) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v int64) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v int64) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v int64) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v int64) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v int64) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan int64:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInterfaceChannel(enum *Enumerator, seq chan interface{}) (i int) {
	var forEachReceived	func(func(interface{}))
	var x				interface{}
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(interface{})) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(interface{})) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(interface{}):
		forEachReceived(func(v interface{}) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v interface{}) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v interface{}) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v interface{}) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v interface{}) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v interface{}) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v interface{}) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
		})
	case chan interface{}:
		forEachReceived(func(v interface{}) {
			f <- v
		})
	case chan R.Value:
		forEachReceived(func(v interface{}) {
			f <- R.ValueOf(v)
		})
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachStringChannel(enum *Enumerator, seq chan string) (i int) {
	var forEachReceived	func(func(string))
	var x				string
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(string)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(string)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(string):
		forEachReceived(func(v string) {
			f(v)
		})
	case func(int, string):
		forEachReceived(func(v string) {
			f(enum.cursor, v)
		})
	case func(interface{}, string):
		forEachReceived(func(v string) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v string) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v string) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v string) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v string) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v string) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v string) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v string) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan string:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUintChannel(enum *Enumerator, seq chan uint) (i int) {
	var forEachReceived	func(func(uint))
	var x				uint
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(uint):
		forEachReceived(func(v uint) {
			f(v)
		})
	case func(int, uint):
		forEachReceived(func(v uint) {
			f(enum.cursor, v)
		})
	case func(interface{}, uint):
		forEachReceived(func(v uint) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v uint) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v uint) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan uint:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint8Channel(enum *Enumerator, seq chan uint8) (i int) {
	var forEachReceived	func(func(uint8))
	var x				uint8
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint8)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint8)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(uint8):
		forEachReceived(func(v uint8) {
			f(v)
		})
	case func(int, uint8):
		forEachReceived(func(v uint8) {
			f(enum.cursor, v)
		})
	case func(interface{}, uint8):
		forEachReceived(func(v uint8) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v uint8) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint8) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint8) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v uint8) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint8) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint8) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint8) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan uint8:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint16Channel(enum *Enumerator, seq chan uint16) (i int) {
	var forEachReceived	func(func(uint16))
	var x				uint16
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint16)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint16)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(uint16):
		forEachReceived(func(v uint16) {
			f(v)
		})
	case func(int, uint16):
		forEachReceived(func(v uint16) {
			f(enum.cursor, v)
		})
	case func(interface{}, uint16):
		forEachReceived(func(v uint16) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v uint16) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint16) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint16) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v uint16) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint16) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint16) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint16) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan uint16:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint32Channel(enum *Enumerator, seq chan uint32) (i int) {
	var forEachReceived	func(func(uint32))
	var x				uint32
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint32)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint32)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(uint32):
		forEachReceived(func(v uint32) {
			f(v)
		})
	case func(int, uint32):
		forEachReceived(func(v uint32) {
			f(enum.cursor, v)
		})
	case func(interface{}, uint32):
		forEachReceived(func(v uint32) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v uint32) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint32) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint32) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v uint32) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint32) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint32) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint32) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan uint32:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint64Channel(enum *Enumerator, seq chan uint64) (i int) {
	var forEachReceived	func(func(uint64))
	var x				uint64
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint64)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint64)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(uint64):
		forEachReceived(func(v uint64) {
			f(v)
		})
	case func(int, uint64):
		forEachReceived(func(v uint64) {
			f(enum.cursor, v)
		})
	case func(interface{}, uint64):
		forEachReceived(func(v uint64) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v uint64) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uint64) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uint64) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v uint64) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uint64) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uint64) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uint64) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan uint64:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUintptrChannel(enum *Enumerator, seq chan uintptr) (i int) {
	var forEachReceived	func(func(uintptr))
	var x				uintptr
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uintptr)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uintptr)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(uintptr):
		forEachReceived(func(v uintptr) {
			f(v)
		})
	case func(int, uintptr):
		forEachReceived(func(v uintptr) {
			f(enum.cursor, v)
		})
	case func(interface{}, uintptr):
		forEachReceived(func(v uintptr) {
			f(enum.cursor, v)
		})
	case func(interface{}):
		forEachReceived(func(v uintptr) {
			f(v)
		})
	case func(int, interface{}):
		forEachReceived(func(v uintptr) {
			f(enum.cursor, v)
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v uintptr) {
			f(enum.cursor, v)
		})
	case func(R.Value):
		forEachReceived(func(v uintptr) {
			f(R.ValueOf(v))
		})
	case func(int, R.Value):
		forEachReceived(func(v uintptr) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v uintptr) {
			f(enum.cursor, R.ValueOf(v))
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v uintptr) {
			f(R.ValueOf(enum.cursor), R.ValueOf(v))
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
//	case []chan uintptr:
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachRValueChannel(enum *Enumerator, seq chan R.Value) (i int) {
	var forEachReceived	func(func(R.Value))
	var x				R.Value
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(R.Value)) {
			for {
				if x, open = <- seq; !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(R.Value)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(interface{}):
		forEachReceived(func(v R.Value) {
			f(v.Interface())
		})
	case func(int, interface{}):
		forEachReceived(func(v R.Value) {
			f(enum.cursor, v.Interface())
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v R.Value) {
			f(enum.cursor, v.Interface())
		})
	case func(R.Value):
		forEachReceived(func(v R.Value) {
			f(v)
		})
	case func(int, R.Value):
		forEachReceived(func(v R.Value) {
			f(enum.cursor, v)
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v R.Value) {
			f(enum.cursor, v)
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v R.Value) {
			f(R.ValueOf(enum.cursor), v)
		})
	case chan interface{}:
		forEachReceived(func(v R.Value) {
			f <- v.Interface()
		})
	case chan R.Value:
		forEachReceived(func(v R.Value) {
			f <- v
		})
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachChannel(enum *Enumerator, c R.Value) (i int) {
	var forEachReceived	func(func(R.Value))
	var x				R.Value
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(R.Value)) {
			for {
				if x, open = c.Recv(); !open {
					PanicWithIndex(enum.cursor)
				}
				f(x)
				i++
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(R.Value)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = c.Recv(); !open {
						PanicWithIndex(enum.cursor + enum.Span)
					}
				}
				f(x)
				i++
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(interface{}):
		forEachReceived(func(v R.Value) {
			f(v.Interface())
		})
	case func(int, interface{}):
		forEachReceived(func(v R.Value) {
			f(enum.cursor, v.Interface())
		})
	case func(interface{}, interface{}):
		forEachReceived(func(v R.Value) {
			f(enum.cursor, v.Interface())
		})
	case func(R.Value):
		forEachReceived(func(v R.Value) {
			f(v)
		})
	case func(int, R.Value):
		forEachReceived(func(v R.Value) {
			f(enum.cursor, v)
		})
	case func(interface{}, R.Value):
		forEachReceived(func(v R.Value) {
			f(enum.cursor, v)
		})
	case func(R.Value, R.Value):
		forEachReceived(func(v R.Value) {
			f(R.ValueOf(enum.cursor), v)
		})
	case chan interface{}:
		forEachReceived(func(v R.Value) {
			f <- v.Interface()
		})
	case chan R.Value:
		forEachReceived(func(v R.Value) {
			f <- v
		})
//	case []chan interface{}:
//	case []chan R.Value:
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 1:				//	f(v)
					p := make([]R.Value, 1, 1)
					forEachReceived(func(v R.Value) {
						p[0] = v
						f.Call(p)
					})
				case 2:				//	f(i, v)
					p := make([]R.Value, 2, 2)
					forEachReceived(func(v R.Value) {
						p[0], p[1] = R.ValueOf(enum.cursor), v
						f.Call(p)
					})
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		} else {
			panic(UNHANDLED_ITERATOR)
		}
	}
	return
}