package sequences

import R "reflect"

func eachBoolChannel(enum *Enumerator, seq chan bool) (i int) {
	var next	func() (bool, bool)

	if enum.Span == 1 {
		next = func() (v, open bool) {
			if v, open = <- seq; open {
				enum.cursor++
			}
			return
		}		
	} else {
		next = func() (v, open bool) {
			offset := enum.Span
			for open = true; open && offset > 0; offset-- {
				v, open = <- seq
				enum.cursor++
			}
			return
		}
	}

	forEachReceived := func(f func(v bool)) {
		steps := enum.steps
		for v, open := <- seq; open && steps > 0; v, open = next() {
			f(v)
			steps--
		}
		i = enum.steps - steps
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
		l := len(f) - 1
		forEachReceived(func(v bool) {
//			go func() {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
//			}()
		})
	case []chan interface{}:
		l := len(f) - 1
		forEachReceived(func(v bool) {
			for n := l; n > 0; n-- {
				f[n] <- v
			}
		})
	case []chan R.Value:
		l := len(f) - 1
		forEachReceived(func(v bool) {
			for n := l; n > 0; n-- {
				f[n] <- R.ValueOf(v)
			}
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachComplex64Channel(enum *Enumerator, seq chan complex64) (i int) {
	var v		complex64

	skipToOffset := func(c chan complex64, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(complex64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, complex64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, complex64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan complex64:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan complex64:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachComplex128Channel(enum *Enumerator, seq chan complex128) (i int) {
	var v		complex128

	skipToOffset := func(c chan complex128, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(complex128):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, complex128):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, complex128):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan complex128:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan complex128:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachErrorChannel(enum *Enumerator, seq chan error) (i int) {
	var v		error

	skipToOffset := func(c chan error, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(error):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, error):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, error):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan error:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan error:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachFloat32Channel(enum *Enumerator, seq chan float32) (i int) {
	var v		float32

	skipToOffset := func(c chan float32, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(float32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, float32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, float32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan float32:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan float32:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachFloat64Channel(enum *Enumerator, seq chan float64) (i int) {
	var v		float64

	skipToOffset := func(c chan float64, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(float64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, float64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, float64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan float64:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan float64:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachIntChannel(enum *Enumerator, seq chan int) (i int) {
	var v		int

	skipToOffset := func(c chan int, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(int):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, int):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, int):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan int:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan int:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt8Channel(enum *Enumerator, seq chan int8) (i int) {
	var v		int8

	skipToOffset := func(c chan int8, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(int8):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, int8):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, int8):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan int8:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan int8:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt16Channel(enum *Enumerator, seq chan int16) (i int) {
	var v		int16

	skipToOffset := func(c chan int16, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(int16):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, int16):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, int16):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan int16:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan int16:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt32Channel(enum *Enumerator, seq chan int32) (i int) {
	var v		int32

	skipToOffset := func(c chan int32, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(int32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, int32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, int32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan int32:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan int32:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInt64Channel(enum *Enumerator, seq chan int64) (i int) {
	var v		int64

	skipToOffset := func(c chan int64, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(int64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, int64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, int64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan int64:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan int64:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachInterfaceChannel(enum *Enumerator, seq chan interface{}) (i int) {
	var v		interface{}

	skipToOffset := func(c chan interface{}, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachStringChannel(enum *Enumerator, seq chan string) (i int) {
	var v		string

	skipToOffset := func(c chan string, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(string):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, string):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, string):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan string:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan string:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUintChannel(enum *Enumerator, seq chan uint) (i int) {
	var v		uint

	skipToOffset := func(c chan uint, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(uint):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, uint):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, uint):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan uint:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan uint:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint8Channel(enum *Enumerator, seq chan uint8) (i int) {
	var v		uint8

	skipToOffset := func(c chan uint8, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(uint8):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, uint8):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, uint8):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan uint8:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan uint8:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint16Channel(enum *Enumerator, seq chan uint16) (i int) {
	var v		uint16

	skipToOffset := func(c chan uint16, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(uint16):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, uint16):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, uint16):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan uint16:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan uint16:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint32Channel(enum *Enumerator, seq chan uint32) (i int) {
	var v		uint32

	skipToOffset := func(c chan uint32, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(uint32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, uint32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, uint32):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan uint32:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan uint32:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUint64Channel(enum *Enumerator, seq chan uint64) (i int) {
	var v		uint64

	skipToOffset := func(c chan uint64, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(uint64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, uint64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, uint64):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan uint64:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan uint64:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachUintptrChannel(enum *Enumerator, seq chan uintptr) (i int) {
	var v		uintptr

	skipToOffset := func(c chan uintptr, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(uintptr):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, uintptr):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, uintptr):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), R.ValueOf(v))
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan uintptr:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- R.ValueOf(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan uintptr:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- R.ValueOf(v)
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachRValueChannel(enum *Enumerator, seq chan R.Value) (i int) {
	var v		R.Value

	skipToOffset := func(c chan R.Value, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	steps := enum.steps
	switch f := enum.f.(type) {
	case func(interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v.Interface())
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v.Interface())
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, interface{}):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v.Interface())
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(int, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(interface{}, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(enum.cursor, v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case func(R.Value, R.Value):
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f(R.ValueOf(enum.cursor), v)
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan interface{}:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v.Interface()
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case chan R.Value:
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				f <- v
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan interface{}:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v.Interface()
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	case []chan R.Value:
		l := len(f) - 1
		for open := true; open && steps > 0; steps-- {
			if v, open = <- seq; open {
				for n := l; n > 0; n-- {
					f[n] <- v
				}
				enum.cursor++
				i++
				open = skipToOffset(seq, enum.Span)
			}
		}
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func eachChannel(enum *Enumerator, c R.Value) (i int) {
	steps := enum.steps
println("limit:", _MAXINT_)
println("enum.cursor:", enum.cursor)
println("enum.steps:", enum.steps)
	switch f := enum.f.(type) {
	case func(interface{}):
		for v, open := c.Recv(); open && steps > 0; steps-- {
			f(v.Interface())
			skipToChannelOffset(c, enum.Span)
			v, open = c.Recv()
		}
		i = enum.steps - steps
	case func(int, interface{}):
		for v, open := c.Recv(); open && steps > 0; steps-- {
			f(enum.cursor, v.Interface())
			skipToChannelOffset(c, enum.Span)
			enum.cursor += enum.Span
			v, open = c.Recv()
		}
		i = enum.steps - steps
	case func(interface{}, interface{}):
		for v, open := c.Recv(); open && steps > 0; steps-- {
			f(enum.cursor, v.Interface())
			skipToChannelOffset(c, enum.Span)
			enum.cursor += enum.Span
			i++
			v, open = c.Recv()
		}
		i = enum.steps - steps
	case func(R.Value):
		for v, open := c.Recv(); open && steps > 0; steps-- {
			f(v)
			skipToChannelOffset(c, enum.Span)
			enum.cursor += enum.Span
			i++
			v, open = c.Recv()
		}
		i = enum.steps - steps
	case func(int, R.Value):
		for v, open := c.Recv(); open && steps > 0; steps-- {
			f(enum.cursor, v)
			skipToChannelOffset(c, enum.Span)
			enum.cursor += enum.Span
			i++
			v, open = c.Recv()
		}
		i = enum.steps - steps
	case func(interface{}, R.Value):
		for v, open := c.Recv(); open && steps > 0; steps-- {
			f(enum.cursor, v)
			skipToChannelOffset(c, enum.Span)
			enum.cursor += enum.Span
			i++
			v, open = c.Recv()
		}
		i = enum.steps - steps
	case func(R.Value, R.Value):
		for v, open := c.Recv(); open && steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), v)
			skipToChannelOffset(c, enum.Span)
			enum.cursor += enum.Span
			i++
			v, open = c.Recv()
		}
		i = enum.steps - steps
	case []chan interface{}:
		panic("[]chan interface{} iterator not yet implemented")
	case []chan R.Value:
		panic("[]chan R.Value iterator not yet implemented")
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 1:				//	f(v)
					p := make([]R.Value, 1, 1)
					for v, open := c.Recv(); open && steps > 0; steps-- {
						p[0] = v
						f.Call(p)
						skipToChannelOffset(c, enum.Span)
						v, open = c.Recv()
					}
					i = enum.steps - steps
				case 2:				//	f(i, v)
					p := make([]R.Value, 2, 2)
					for v, open := c.Recv(); open && steps > 0; steps-- {
						p[0], p[1] = R.ValueOf(enum.cursor), v
						f.Call(p)
						skipToChannelOffset(c, enum.Span)
						enum.cursor += enum.Span
						v, open = c.Recv()
					}
					i = enum.steps - steps
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		}
	}
	return
}