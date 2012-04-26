package sequences

import R "reflect"

func reduceBoolChannel(enum *Enumerator, seq chan bool) (r bool) {
	var forEachReceived	func(func(bool))
	var x				bool
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(bool)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(bool)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(bool, bool) bool:
		r = enum.seed.(bool)
		forEachReceived(func(v bool) {
			r = f(r, v)
		})
	case func(bool, int, bool) bool:
		r = enum.seed.(bool)
		forEachReceived(func(v bool) {
			r = f(r, enum.cursor, v)
		})
	case func(bool, interface{}, bool) bool:
		r = enum.seed.(bool)
		forEachReceived(func(v bool) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v bool) {
			ri = f(ri, v)
		})
		r = ri.(bool)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v bool) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(bool)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v bool) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(bool)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v bool) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Bool()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v bool) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Bool()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v bool) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Bool()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v bool) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = rv.Bool()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}


func reduceComplex64Channel(enum *Enumerator, seq chan complex64) (r complex64) {
	var forEachReceived	func(func(complex64))
	var x				complex64
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(complex64)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(complex64)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(complex64, complex64) complex64:
		r = enum.seed.(complex64)
		forEachReceived(func(v complex64) {
			r = f(r, v)
		})
	case func(complex64, int, complex64) complex64:
		r = enum.seed.(complex64)
		forEachReceived(func(v complex64) {
			r = f(r, enum.cursor, v)
		})
	case func(complex64, interface{}, complex64) complex64:
		r = enum.seed.(complex64)
		forEachReceived(func(v complex64) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex64) {
			ri = f(ri, v)
		})
		r = ri.(complex64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex64) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(complex64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex64) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(complex64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v complex64) {
			rv = f(rv, R.ValueOf(v))
		})
		r = complex64(rv.Complex())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex64) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = complex64(rv.Complex())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex64) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = complex64(rv.Complex())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex64) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = complex64(rv.Complex())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceComplex128Channel(enum *Enumerator, seq chan complex128) (r complex128) {
	var forEachReceived	func(func(complex128))
	var x				complex128
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(complex128)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(complex128)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(complex128, complex128) complex128:
		r = enum.seed.(complex128)
		forEachReceived(func(v complex128) {
			r = f(r, v)
		})
	case func(complex128, int, complex128) complex128:
		r = enum.seed.(complex128)
		forEachReceived(func(v complex128) {
			r = f(r, enum.cursor, v)
		})
	case func(complex128, interface{}, complex128) complex128:
		r = enum.seed.(complex128)
		forEachReceived(func(v complex128) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex128) {
			ri = f(ri, v)
		})
		r = ri.(complex128)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex128) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(complex128)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex128) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(complex128)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v complex128) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Complex()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex128) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Complex()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex128) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Complex()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex128) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = rv.Complex()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}


func reduceErrorChannel(enum *Enumerator, seq chan error) (r error) {
	var forEachReceived	func(func(error))
	var x				error
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(error)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(error)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(error, error) error:
		r = enum.seed.(error)
		forEachReceived(func(v error) {
			r = f(r, v)
		})
	case func(error, int, error) error:
		r = enum.seed.(error)
		forEachReceived(func(v error) {
			r = f(r, enum.cursor, v)
		})
	case func(error, interface{}, error) error:
		r = enum.seed.(error)
		forEachReceived(func(v error) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v error) {
			ri = f(ri, v)
		})
		r = ri.(error)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v error) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(error)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v error) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(error)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v error) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Interface().(error)
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v error) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Interface().(error)
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v error) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Interface().(error)
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v error) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = rv.Interface().(error)
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceFloat32Channel(enum *Enumerator, seq chan float32) (r float32) {
	var forEachReceived	func(func(float32))
	var x				float32
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(float32)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(float32)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(float32, float32) float32:
		r = enum.seed.(float32)
		forEachReceived(func(v float32) {
			r = f(r, v)
		})
	case func(float32, int, float32) float32:
		r = enum.seed.(float32)
		forEachReceived(func(v float32) {
			r = f(r, enum.cursor, v)
		})
	case func(float32, interface{}, float32) float32:
		r = enum.seed.(float32)
		forEachReceived(func(v float32) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float32) {
			ri = f(ri, v)
		})
		r = ri.(float32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float32) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(float32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float32) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(float32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v float32) {
			rv = f(rv, R.ValueOf(v))
		})
		r = float32(rv.Float())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float32) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = float32(rv.Float())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float32) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = float32(rv.Float())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float32) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = float32(rv.Float())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceFloat64Channel(enum *Enumerator, seq chan float64) (r float64) {
	var forEachReceived	func(func(float64))
	var x				float64
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(float64)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(float64)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(float64, float64) float64:
		r = enum.seed.(float64)
		forEachReceived(func(v float64) {
			r = f(r, v)
		})
	case func(float64, int, float64) float64:
		r = enum.seed.(float64)
		forEachReceived(func(v float64) {
			r = f(r, enum.cursor, v)
		})
	case func(float64, interface{}, float64) float64:
		r = enum.seed.(float64)
		forEachReceived(func(v float64) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float64) {
			ri = f(ri, v)
		})
		r = ri.(float64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float64) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(float64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float64) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(float64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v float64) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Float()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float64) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Float()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float64) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Float()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float64) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = rv.Float()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceIntChannel(enum *Enumerator, seq chan int) (r int) {
	var forEachReceived	func(func(int))
	var x				int
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(int, int) int:
		r = enum.seed.(int)
		forEachReceived(func(v int) {
			r = f(r, v)
		})
	case func(int, int, int) int:
		r = enum.seed.(int)
		forEachReceived(func(v int) {
			r = f(r, enum.cursor, v)
		})
	case func(int, interface{}, int) int:
		r = enum.seed.(int)
		forEachReceived(func(v int) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int) {
			ri = f(ri, v)
		})
		r = ri.(int)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v int) {
			rv = f(rv, R.ValueOf(v))
		})
		r = int(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = int(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = int(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = int(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt8Channel(enum *Enumerator, seq chan int8) (r int8) {
	var forEachReceived	func(func(int8))
	var x				int8
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int8)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int8)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(int8, int8) int8:
		r = enum.seed.(int8)
		forEachReceived(func(v int8) {
			r = f(r, v)
		})
	case func(int8, int, int8) int8:
		r = enum.seed.(int8)
		forEachReceived(func(v int8) {
			r = f(r, enum.cursor, v)
		})
	case func(int8, interface{}, int8) int8:
		r = enum.seed.(int8)
		forEachReceived(func(v int8) {
			r = f(r, enum.cursor, v)
		})
	case func(int, int) int:
		ri := int(enum.seed.(int8))
		forEachReceived(func(v int8) {
			ri = f(ri, int(v))
		})
		r = int8(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int8))
		forEachReceived(func(v int8) {
			ri = f(ri, enum.cursor, int(v))
		})
		r = int8(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int8))
		forEachReceived(func(v int8) {
			ri = f(ri, enum.cursor, int(v))
		})
		r = int8(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int8) {
			ri = f(ri, v)
		})
		r = ri.(int8)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int8) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int8)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int8) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int8)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v int8) {
			rv = f(rv, R.ValueOf(v))
		})
		r = int8(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int8) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = int8(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int8) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = int8(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int8) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = int8(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt16Channel(enum *Enumerator, seq chan int16) (r int16) {
	var forEachReceived	func(func(int16))
	var x				int16
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int16)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int16)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(int16, int16) int16:
		r = enum.seed.(int16)
		forEachReceived(func(v int16) {
			r = f(r, v)
		})
	case func(int16, int, int16) int16:
		r = enum.seed.(int16)
		forEachReceived(func(v int16) {
			r = f(r, enum.cursor, v)
		})
	case func(int16, interface{}, int16) int16:
		r = enum.seed.(int16)
		forEachReceived(func(v int16) {
			r = f(r, enum.cursor, v)
		})
	case func(int, int) int:
		ri := int(enum.seed.(int16))
		forEachReceived(func(v int16) {
			ri = f(ri, int(v))
		})
		r = int16(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int16))
		forEachReceived(func(v int16) {
			ri = f(ri, enum.cursor, int(v))
		})
		r = int16(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int16))
		forEachReceived(func(v int16) {
			ri = f(ri, enum.cursor, int(v))
		})
		r = int16(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int16) {
			ri = f(ri, v)
		})
		r = ri.(int16)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int16) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int16)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int16) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int16)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v int16) {
			rv = f(rv, R.ValueOf(v))
		})
		r = int16(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int16) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = int16(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int16) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = int16(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int16) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = int16(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt32Channel(enum *Enumerator, seq chan int32) (r int32) {
	var forEachReceived	func(func(int32))
	var x				int32
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int32)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int32)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(int32, int32) int32:
		r = enum.seed.(int32)
		forEachReceived(func(v int32) {
			r = f(r, v)
		})
	case func(int32, int, int32) int32:
		r = enum.seed.(int32)
		forEachReceived(func(v int32) {
			r = f(r, enum.cursor, v)
		})
	case func(int32, interface{}, int32) int32:
		r = enum.seed.(int32)
		forEachReceived(func(v int32) {
			r = f(r, enum.cursor, v)
		})
	case func(int, int) int:
		ri := int(enum.seed.(int32))
		forEachReceived(func(v int32) {
			ri = f(ri, int(v))
		})
		r = int32(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int32))
		forEachReceived(func(v int32) {
			ri = f(ri, enum.cursor, int(v))
		})
		r = int32(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int32))
		forEachReceived(func(v int32) {
			ri = f(ri, enum.cursor, int(v))
		})
		r = int32(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int32) {
			ri = f(ri, v)
		})
		r = ri.(int32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int32) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int32) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v int32) {
			rv = f(rv, R.ValueOf(v))
		})
		r = int32(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int32) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = int32(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int32) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = int32(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int32) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = int32(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt64Channel(enum *Enumerator, seq chan int64) (r int64) {
	var forEachReceived	func(func(int64))
	var x				int64
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(int64)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(int64)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(int64, int64) int64:
		r = enum.seed.(int64)
		forEachReceived(func(v int64) {
			r = f(r, v)
		})
	case func(int64, int, int64) int64:
		r = enum.seed.(int64)
		forEachReceived(func(v int64) {
			r = f(r, enum.cursor, v)
		})
	case func(int64, interface{}, int64) int64:
		r = enum.seed.(int64)
		forEachReceived(func(v int64) {
			r = f(r, enum.cursor, v)
		})
	case func(int, int) int:
		ri := int(enum.seed.(int64))
		forEachReceived(func(v int64) {
			ri = f(ri, int(v))
		})
		r = int64(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int64))
		forEachReceived(func(v int64) {
			ri = f(ri, enum.cursor, int(v))
		})
		r = int64(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int64))
		forEachReceived(func(v int64) {
			ri = f(ri, enum.cursor, int(v))
		})
		r = int64(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int64) {
			ri = f(ri, v)
		})
		r = ri.(int64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int64) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int64) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(int64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v int64) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Int()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int64) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Int()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int64) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Int()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int64) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = rv.Int()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInterfaceChannel(enum *Enumerator, seq chan interface{}) (r interface{}) {
	var forEachReceived	func(func(interface{}))
	var x				interface{}
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(interface{})) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(interface{})) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}
	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		r = enum.seed
		forEachReceived(func(v interface{}) {
			r = f(r, v)
		})
	case func(interface{}, int, interface{}) interface{}:
		r = enum.seed
		forEachReceived(func(v interface{}) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}, interface{}) interface{}:
		r = enum.seed
		forEachReceived(func(v interface{}) {
			r = f(r, enum.cursor, v)
		})
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v interface{}) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Interface()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v interface{}) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Interface()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v interface{}) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Interface()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v interface{}) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = rv.Interface()
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			switch f.Type().NumIn() {
			case 2:
				p := make([]R.Value, 2, 2)
				p[0] = R.ValueOf(enum.seed)
				forEachReceived(func(v interface{}) {
					p[1] = R.ValueOf(v)
					p[0] = f.Call(p)[0]
				})
				r = p[0].Interface()
			case 3:
				p := make([]R.Value, 3, 3)
				p[0] = R.ValueOf(enum.seed)
				forEachReceived(func(v interface{}) {
					p[1], p[2] = R.ValueOf(enum.cursor), R.ValueOf(v)
					p[0] = f.Call(p)[0]
				})
				r = p[0].Interface()
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

func reduceStringChannel(enum *Enumerator, seq chan string) (r string) {
	var forEachReceived	func(func(string))
	var x				string
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(string)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(string)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(string, string) string:
		r = enum.seed.(string)
		forEachReceived(func(v string) {
			r = f(r, v)
		})
	case func(string, int, string) string:
		r = enum.seed.(string)
		forEachReceived(func(v string) {
			r = f(r, enum.cursor, v)
		})
	case func(string, interface{}, string) string:
		r = enum.seed.(string)
		forEachReceived(func(v string) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v string) {
			ri = f(ri, v)
		})
		r = ri.(string)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v string) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(string)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v string) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(string)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v string) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.String()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v string) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.String()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v string) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.String()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v string) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = rv.String()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUintChannel(enum *Enumerator, seq chan uint) (r uint) {
	var forEachReceived	func(func(uint))
	var x				uint
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(uint, uint) uint:
		r = enum.seed.(uint)
		forEachReceived(func(v uint) {
			r = f(r, v)
		})
	case func(uint, int, uint) uint:
		r = enum.seed.(uint)
		forEachReceived(func(v uint) {
			r = f(r, enum.cursor, v)
		})
	case func(uint, interface{}, uint) uint:
		r = enum.seed.(uint)
		forEachReceived(func(v uint) {
			r = f(r, enum.cursor, v)
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint) {
			ri = f(ri, v)
		})
		r = ri.(uint)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v uint) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uint(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uint(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uint(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = uint(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint8Channel(enum *Enumerator, seq chan uint8) (r uint8) {
	var forEachReceived	func(func(uint8))
	var x				uint8
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint8)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint8)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(uint8, uint8) uint8:
		r = enum.seed.(uint8)
		forEachReceived(func(v uint8) {
			r = f(r, v)
		})
	case func(uint8, int, uint8) uint8:
		r = enum.seed.(uint8)
		forEachReceived(func(v uint8) {
			r = f(r, enum.cursor, v)
		})
	case func(uint8, interface{}, uint8) uint8:
		r = enum.seed.(uint8)
		forEachReceived(func(v uint8) {
			r = f(r, enum.cursor, v)
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uint8))
		forEachReceived(func(v uint8) {
			ru = f(ru, uint(v))
		})
		r = uint8(ru)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uint8))
		forEachReceived(func(v uint8) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uint8(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uint8))
		forEachReceived(func(v uint8) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uint8(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint8) {
			ri = f(ri, v)
		})
		r = ri.(uint8)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint8) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint8)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint8) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint8)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v uint8) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uint8(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint8) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uint8(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint8) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uint8(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint8) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = uint8(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint16Channel(enum *Enumerator, seq chan uint16) (r uint16) {
	var forEachReceived	func(func(uint16))
	var x				uint16
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint16)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint16)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(uint16, uint16) uint16:
		r = enum.seed.(uint16)
		forEachReceived(func(v uint16) {
			r = f(r, v)
		})
	case func(uint16, int, uint16) uint16:
		r = enum.seed.(uint16)
		forEachReceived(func(v uint16) {
			r = f(r, enum.cursor, v)
		})
	case func(uint16, interface{}, uint16) uint16:
		r = enum.seed.(uint16)
		forEachReceived(func(v uint16) {
			r = f(r, enum.cursor, v)
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uint16))
		forEachReceived(func(v uint16) {
			ru = f(ru, uint(v))
		})
		r = uint16(ru)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uint16))
		forEachReceived(func(v uint16) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uint16(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uint16))
		forEachReceived(func(v uint16) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uint16(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint16) {
			ri = f(ri, v)
		})
		r = ri.(uint16)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint16) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint16)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint16) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint16)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v uint16) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uint16(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint16) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uint16(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint16) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uint16(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint16) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = uint16(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint32Channel(enum *Enumerator, seq chan uint32) (r uint32) {
	var forEachReceived	func(func(uint32))
	var x				uint32
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint32)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint32)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(uint32, uint32) uint32:
		r = enum.seed.(uint32)
		forEachReceived(func(v uint32) {
			r = f(r, v)
		})
	case func(uint32, int, uint32) uint32:
		r = enum.seed.(uint32)
		forEachReceived(func(v uint32) {
			r = f(r, enum.cursor, v)
		})
	case func(uint32, interface{}, uint32) uint32:
		r = enum.seed.(uint32)
		forEachReceived(func(v uint32) {
			r = f(r, enum.cursor, v)
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uint32))
		forEachReceived(func(v uint32) {
			ru = f(ru, uint(v))
		})
		r = uint32(ru)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uint32))
		forEachReceived(func(v uint32) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uint32(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uint32))
		forEachReceived(func(v uint32) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uint32(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint32) {
			ri = f(ri, v)
		})
		r = ri.(uint32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint32) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint32) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v uint32) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uint32(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint32) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uint32(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint32) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uint32(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint32) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = uint32(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint64Channel(enum *Enumerator, seq chan uint64) (r uint64) {
	var forEachReceived	func(func(uint64))
	var x				uint64
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uint64)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uint64)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(uint64, uint64) uint64:
		r = enum.seed.(uint64)
		forEachReceived(func(v uint64) {
			r = f(r, v)
		})
	case func(uint64, int, uint64) uint64:
		r = enum.seed.(uint64)
		forEachReceived(func(v uint64) {
			r = f(r, enum.cursor, v)
		})
	case func(uint64, interface{}, uint64) uint64:
		r = enum.seed.(uint64)
		forEachReceived(func(v uint64) {
			r = f(r, enum.cursor, v)
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uint64))
		forEachReceived(func(v uint64) {
			ru = f(ru, uint(v))
		})
		r = uint64(ru)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uint64))
		forEachReceived(func(v uint64) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uint64(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uint64))
		forEachReceived(func(v uint64) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uint64(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint64) {
			ri = f(ri, v)
		})
		r = ri.(uint64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint64) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint64) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uint64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v uint64) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Uint()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint64) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Uint()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint64) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = rv.Uint()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint64) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = rv.Uint()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUintptrChannel(enum *Enumerator, seq chan uintptr) (r uintptr) {
	var forEachReceived	func(func(uintptr))
	var x				uintptr
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(uintptr)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(uintptr)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(uintptr, uintptr) uintptr:
		r = enum.seed.(uintptr)
		forEachReceived(func(v uintptr) {
			r = f(r, v)
		})
	case func(uintptr, int, uintptr) uintptr:
		r = enum.seed.(uintptr)
		forEachReceived(func(v uintptr) {
			r = f(r, enum.cursor, v)
		})
	case func(uintptr, interface{}, uintptr) uintptr:
		r = enum.seed.(uintptr)
		forEachReceived(func(v uintptr) {
			r = f(r, enum.cursor, v)
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uintptr))
		forEachReceived(func(v uintptr) {
			ru = f(ru, uint(v))
		})
		r = uintptr(r)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uintptr))
		forEachReceived(func(v uintptr) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uintptr(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uintptr))
		forEachReceived(func(v uintptr) {
			ru = f(ru, enum.cursor, uint(v))
		})
		r = uintptr(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uintptr) {
			ri = f(ri, v)
		})
		r = ri.(uintptr)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uintptr) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uintptr)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uintptr) {
			ri = f(ri, enum.cursor, v)
		})
		r = ri.(uintptr)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		forEachReceived(func(v uintptr) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uintptr(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uintptr) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uintptr(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uintptr) {
			rv = f(rv, enum.cursor, R.ValueOf(v))
		})
		r = uintptr(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uintptr) {
			rv = f(rv, R.ValueOf(enum.cursor), R.ValueOf(v))
		})
		r = uintptr(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceRValueChannel(enum *Enumerator, seq chan R.Value) (r R.Value) {
	var forEachReceived	func(func(R.Value))
	var x				R.Value
	var open			bool

	if enum.Span == 1 {
		forEachReceived = func(f func(R.Value)) {
			for {
				if x, open = <- seq; !open {
					return
				}
				f(x)
			}
			enum.cursor++
		}
	} else {
		forEachReceived = func(f func(R.Value)) {
			var offset	int
			for {
				for offset = enum.Span; offset > 0; offset-- {
					if x, open = <- seq; !open {
						return
					}
				}
				f(x)
			}
			enum.cursor += enum.Span
		}
	}

	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		var ri	interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, v.Interface())
		})
		r = R.ValueOf(ri)
	case func(interface{}, int, interface{}) interface{}:
		var ri	interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, enum.cursor, v.Interface())
		})
		r = R.ValueOf(ri)
	case func(interface{}, interface{}, interface{}) interface{}:
		var ri	interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, enum.cursor, v.Interface())
		})
		r = R.ValueOf(ri)
	case func(R.Value, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, v)
		})
	case func(R.Value, int, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, enum.cursor, v)
		})
	case func(R.Value, interface{}, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, enum.cursor, v)
		})
	case func(R.Value, R.Value, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, R.ValueOf(enum.cursor), v)
		})
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			switch f.Type().NumIn() {
			case 2:
				p := make([]R.Value, 2, 2)
				if rv, ok := enum.seed.(R.Value); ok {
					p[0] = rv
				} else {
					p[0] = R.ValueOf(enum.seed)
				}
				forEachReceived(func(v R.Value) {
					p[1] = v
					p[0] = f.Call(p)[0]
				})
				r = p[0]
			case 3:
				p := make([]R.Value, 3, 3)
				if rv, ok := enum.seed.(R.Value); ok {
					p[0] = rv
				} else {
					p[0] = R.ValueOf(enum.seed)
				}
				forEachReceived(func(v R.Value) {
					p[1], p[2] = R.ValueOf(enum.cursor), v
					p[0] = f.Call(p)[0]
				})
				r = p[0]
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}
/*
func reduceChannel(enum *Enumerator, s R.Value) (r R.Value) {
	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		var ok		bool
		var v		R.Value
		var ri		interface{}
		if r, ok = enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		enum.reduce(func(cursor int) {
			if v, ok = s.Recv(); ok {
				ri = f(ri, v.Interface())
			} else {
				PanicWithIndex(cursor)
			}
		})
		r = R.ValueOf(ri)
	case func(interface{}, int, interface{}) interface{}:
		var ok		bool
		var v		R.Value
		var ri		interface{}
		if r, ok = enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		enum.reduce(func(cursor int) {
			if v, ok = s.Recv(); ok {
				ri = f(ri, cursor, v.Interface())
			} else {
				PanicWithIndex(cursor)
			}
		})
		r = R.ValueOf(ri)
	case func(interface{}, interface{}, interface{}) interface{}:
		var ok		bool
		var v		R.Value
		var ri		interface{}
		if r, ok = enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		enum.reduce(func(cursor int) {
			if v, ok = s.Recv(); ok {
				ri = f(ri, cursor, v.Interface())
			} else {
				PanicWithIndex(cursor)
			}
		})
		r = R.ValueOf(ri)
	case func(R.Value, R.Value) R.Value:
		var ok		bool
		var v		R.Value
		if v, ok = enum.seed.(R.Value); ok {
			r = v
		} else {
			r = R.ValueOf(enum.seed)
		}
		enum.reduce(func(cursor int) {
			if v, ok = s.Recv(); ok {
				r = f(r, v)
			} else {
				PanicWithIndex(cursor)
			}
		})
	case func(R.Value, int, R.Value) R.Value:
		var ok		bool
		var v		R.Value
		if v, ok = enum.seed.(R.Value); ok {
			r = v
		} else {
			r = R.ValueOf(enum.seed)
		}
		enum.reduce(func(cursor int) {
			if v, ok = s.Recv(); ok {
				r = f(r, cursor, v)
			} else {
				PanicWithIndex(cursor)
			}
		})
	case func(R.Value, interface{}, R.Value) R.Value:
		var ok		bool
		var v		R.Value
		if v, ok = enum.seed.(R.Value); ok {
			r = v
		} else {
			r = R.ValueOf(enum.seed)
		}
		enum.reduce(func(cursor int) {
			if v, ok = s.Recv(); ok {
				r = f(r, cursor, v)
			} else {
				PanicWithIndex(cursor)
			}
		})
	case func(R.Value, R.Value, R.Value) R.Value:
		var ok		bool
		var v		R.Value
		if v, ok = enum.seed.(R.Value); ok {
			r = v
		} else {
			r = R.ValueOf(enum.seed)
		}
		enum.reduce(func(cursor int) {
			if v, ok = s.Recv(); ok {
				r = f(r, R.ValueOf(cursor), v)
			} else {
				PanicWithIndex(cursor)
			}
		})
	default:
		switch f := R.ValueOf(f); f.Kind() {
		case R.Func:
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 2:
					var ok		bool
					var v		R.Value
					p := make([]R.Value, 2, 2)
					if r, ok = enum.seed.(R.Value); ok {
						p[0] = r
					} else {
						p[0] = R.ValueOf(enum.seed)
					}
					enum.reduce(func(cursor int) {
						if v, ok = s.Recv(); ok {
							p[1] = v
							p[0] = f.Call(p)[0]
						} else {
							PanicWithIndex(cursor)
						}
					})
					r = p[0]
				case 3:
					var ok		bool
					var v		R.Value
					p := make([]R.Value, 3, 3)
					if r, ok = enum.seed.(R.Value); ok {
						p[0] = r
					} else {
						p[0] = R.ValueOf(enum.seed)
					}
					enum.reduce(func(cursor int) {
						if v, ok = s.Recv(); ok {
							p[1], p[2] = R.ValueOf(cursor), v
							p[0] = f.Call(p)[0]
						} else {
							PanicWithIndex(cursor)
						}
					})
					r = p[0]
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		default:
			panic(UNHANDLED_ITERATOR)
		}
	}
	return
}
*/