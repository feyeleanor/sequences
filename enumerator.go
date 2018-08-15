package sequences

import R "reflect"

type Enumerator struct {
	Sequence interface{}
	Span     int
	f        interface{}
	seed     interface{}
}

func NewEnumerator(seq interface{}) *Enumerator {
	return &Enumerator{Sequence: seq, Span: 1}
}

func (enum Enumerator) Each(f interface{}) {
	enum.f = f
	switch seq := enum.Sequence.(type) {
	case Enumerable:
		seq.Each(&enum)
	case Indexable:
		eachIndexable(&enum, seq)
	case []bool:
		eachBoolSlice(&enum, seq)
	case []complex64:
		eachComplex64Slice(&enum, seq)
	case []complex128:
		eachComplex128Slice(&enum, seq)
	case []error:
		eachErrorSlice(&enum, seq)
	case []float32:
		eachFloat32Slice(&enum, seq)
	case []float64:
		eachFloat64Slice(&enum, seq)
	case []int:
		eachIntSlice(&enum, seq)
	case []int8:
		eachInt8Slice(&enum, seq)
	case []int16:
		eachInt16Slice(&enum, seq)
	case []int32:
		eachInt32Slice(&enum, seq)
	case []int64:
		eachInt64Slice(&enum, seq)
	case []interface{}:
		eachInterfaceSlice(&enum, seq)
	case []string:
		eachStringSlice(&enum, seq)
	case []uint:
		eachUintSlice(&enum, seq)
	case []uint8:
		eachUint8Slice(&enum, seq)
	case []uint16:
		eachUint16Slice(&enum, seq)
	case []uint32:
		eachUint32Slice(&enum, seq)
	case []uint64:
		eachUint64Slice(&enum, seq)
	case []uintptr:
		eachUintptrSlice(&enum, seq)
	case []R.Value:
		eachRValueSlice(&enum, seq)
	case chan bool:
		eachBoolChannel(&enum, seq)
	case chan complex64:
		eachComplex64Channel(&enum, seq)
	case chan complex128:
		eachComplex128Channel(&enum, seq)
	case chan error:
		eachErrorChannel(&enum, seq)
	case chan float32:
		eachFloat32Channel(&enum, seq)
	case chan float64:
		eachFloat64Channel(&enum, seq)
	case chan int:
		eachIntChannel(&enum, seq)
	case chan int8:
		eachInt8Channel(&enum, seq)
	case chan int16:
		eachInt16Channel(&enum, seq)
	case chan int32:
		eachInt32Channel(&enum, seq)
	case chan int64:
		eachInt64Channel(&enum, seq)
	case chan interface{}:
		eachInterfaceChannel(&enum, seq)
	case chan string:
		eachStringChannel(&enum, seq)
	case chan uint:
		eachUintChannel(&enum, seq)
	case chan uint8:
		eachUint8Channel(&enum, seq)
	case chan uint16:
		eachUint16Channel(&enum, seq)
	case chan uint32:
		eachUint32Channel(&enum, seq)
	case chan uint64:
		eachUint64Channel(&enum, seq)
	case chan uintptr:
		eachUintptrChannel(&enum, seq)
	case chan R.Value:
		eachRValueChannel(&enum, seq)
	case func(int) bool:
		enum.eachboolFunction(seq)
	case func(int) complex64:
		enum.eachcomplex64Function(seq)
	case func(int) complex128:
		enum.eachcomplex128Function(seq)
	case func(int) error:
		enum.eacherrorFunction(seq)
	case func(int) float32:
		enum.eachfloat32Function(seq)
	case func(int) float64:
		enum.eachfloat64Function(seq)
	case func(int) int:
		enum.eachintFunction(seq)
	case func(int) int8:
		enum.eachint8Function(seq)
	case func(int) int16:
		enum.eachint16Function(seq)
	case func(int) int32:
		enum.eachint32Function(seq)
	case func(int) int64:
		enum.eachint64Function(seq)
	case func(int) interface{}:
		enum.eachinterfaceFunction(seq)
	case func(int) string:
		enum.eachstringFunction(seq)
	case func(int) uint:
		enum.eachuintFunction(seq)
	case func(int) uint8:
		enum.eachuint8Function(seq)
	case func(int) uint16:
		enum.eachuint16Function(seq)
	case func(int) uint32:
		enum.eachuint32Function(seq)
	case func(int) uint64:
		enum.eachuint64Function(seq)
	case func(int) uintptr:
		enum.eachuintptrFunction(seq)
	case func(int) R.Value:
		enum.eachRValueFunction(seq)
	case R.Value:
		switch seq.Kind() {
		case R.Slice:
			eachSlice(&enum, seq)
		case R.Chan:
			if enum.Span < 1 {
				panic(ASCENDING_SEQUENCE)
			} else {
				eachChannel(&enum, seq)
			}
		case R.Func:
			if isFunctionSequence(seq) {
				eachFunction(&enum, seq)
			} else {
				panic(NOT_A_SEQUENCE)
			}
		default:
			switch c := R.ValueOf(seq); c.Kind() {
			case R.Slice:
				eachSlice(&enum, c)
			case R.Chan:
				if enum.Span < 1 {
					panic(ASCENDING_SEQUENCE)
				} else {
					eachChannel(&enum, c)
				}
			case R.Func:
				if isFunctionSequence(seq) {
					eachFunction(&enum, c)
				} else {
					panic(NOT_A_SEQUENCE)
				}
			}
		}
	default:
		enum.Each(R.ValueOf(enum.Sequence))
	}
	return
}

func (enum Enumerator) each(f interface{}) {
	defer IgnoreIndexOutOfRange()
	var cursor int

	switch f := f.(type) {
	case func(int):
		for {
			f(cursor)
			cursor++
		}
	case func(int) bool:
		for {
			f(cursor)
			cursor++
		}
	default:
		panic(f)
	}
}

func (enum Enumerator) Reduce(seed, f interface{}) (r interface{}) {
	enum.f = f
	enum.seed = seed
	switch seq := enum.Sequence.(type) {
	case Reducible:
		r = seq.Reduce(&enum)
	case Indexable:
		r = reduceIndexable(&enum, seq)
	case []bool:
		r = reduceBoolSlice(&enum, seq)
	case []complex64:
		r = reduceComplex64Slice(&enum, seq)
	case []complex128:
		r = reduceComplex128Slice(&enum, seq)
	case []error:
		r = reduceErrorSlice(&enum, seq)
	case []float32:
		r = reduceFloat32Slice(&enum, seq)
	case []float64:
		r = reduceFloat64Slice(&enum, seq)
	case []int:
		r = reduceIntSlice(&enum, seq)
	case []int8:
		r = reduceInt8Slice(&enum, seq)
	case []int16:
		r = reduceInt16Slice(&enum, seq)
	case []int32:
		r = reduceInt32Slice(&enum, seq)
	case []int64:
		r = reduceInt64Slice(&enum, seq)
	case []interface{}:
		r = reduceInterfaceSlice(&enum, seq)
	case []string:
		r = reduceStringSlice(&enum, seq)
	case []uint:
		r = reduceUintSlice(&enum, seq)
	case []uint8:
		r = reduceUint8Slice(&enum, seq)
	case []uint16:
		r = reduceUint16Slice(&enum, seq)
	case []uint32:
		r = reduceUint32Slice(&enum, seq)
	case []uint64:
		r = reduceUint64Slice(&enum, seq)
	case []uintptr:
		r = reduceUintptrSlice(&enum, seq)
	case []R.Value:
		r = reduceRValueSlice(&enum, seq)
	case func(int) bool:
		r = reduceBoolFunction(&enum, seq)
	case func(int) complex64:
		r = reduceComplex64Function(&enum, seq)
	case func(int) complex128:
		r = reduceComplex128Function(&enum, seq)
	case func(int) error:
		r = reduceErrorFunction(&enum, seq)
	case func(int) float32:
		r = reduceFloat32Function(&enum, seq)
	case func(int) float64:
		r = reduceFloat64Function(&enum, seq)
	case func(int) int:
		r = reduceIntFunction(&enum, seq)
	case func(int) int8:
		r = reduceInt8Function(&enum, seq)
	case func(int) int16:
		r = reduceInt16Function(&enum, seq)
	case func(int) int32:
		r = reduceInt32Function(&enum, seq)
	case func(int) int64:
		r = reduceInt64Function(&enum, seq)
	case func(int) interface{}:
		r = reduceInterfaceFunction(&enum, seq)
	case func(int) string:
		r = reduceStringFunction(&enum, seq)
	case func(int) uint:
		r = reduceUintFunction(&enum, seq)
	case func(int) uint8:
		r = reduceUint8Function(&enum, seq)
	case func(int) uint16:
		r = reduceUint16Function(&enum, seq)
	case func(int) uint32:
		r = reduceUint32Function(&enum, seq)
	case func(int) uint64:
		r = reduceUint64Function(&enum, seq)
	case func(int) uintptr:
		r = reduceUintptrFunction(&enum, seq)
	case func(int) R.Value:
		r = reduceRValueFunction(&enum, seq)
	case chan bool:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceBoolChannel(&enum, seq)
		}
	case chan complex64:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceComplex64Channel(&enum, seq)
		}
	case chan complex128:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceComplex128Channel(&enum, seq)
		}
	case chan error:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceErrorChannel(&enum, seq)
		}
	case chan float32:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceFloat32Channel(&enum, seq)
		}
	case chan float64:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceFloat64Channel(&enum, seq)
		}
	case chan int:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceIntChannel(&enum, seq)
		}
	case chan int8:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceInt8Channel(&enum, seq)
		}
	case chan int16:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceInt16Channel(&enum, seq)
		}
	case chan int32:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceInt32Channel(&enum, seq)
		}
	case chan int64:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceInt64Channel(&enum, seq)
		}
	case chan interface{}:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceInterfaceChannel(&enum, seq)
		}
	case chan string:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceStringChannel(&enum, seq)
		}
	case chan uint:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceUintChannel(&enum, seq)
		}
	case chan uint8:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceUint8Channel(&enum, seq)
		}
	case chan uint16:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceUint16Channel(&enum, seq)
		}
	case chan uint32:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceUint32Channel(&enum, seq)
		}
	case chan uint64:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceUint64Channel(&enum, seq)
		}
	case chan uintptr:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceUintptrChannel(&enum, seq)
		}
	case chan R.Value:
		if enum.Span < 1 {
			panic(ASCENDING_SEQUENCE)
		} else {
			r = reduceRValueChannel(&enum, seq)
		}
	case R.Value:
		switch seq.Kind() {
		case R.Slice:
			r = reduceSlice(&enum, seq)
		case R.Func:
			r = reduceFunction(&enum, seq)
		case R.Chan:
			if enum.Span < 1 {
				panic(ASCENDING_SEQUENCE)
			} else {
				r = reduceChannel(&enum, seq)
			}
		}
	default:
		switch c := R.ValueOf(seq); c.Kind() {
		case R.Slice:
			r = reduceSlice(&enum, c)
		case R.Func:
			r = reduceFunction(&enum, c)
		case R.Chan:
			if enum.Span < 1 {
				panic(ASCENDING_SEQUENCE)
			} else {
				r = reduceChannel(&enum, c)
			}
		}
	}
	return
}

func (enum Enumerator) reduce(f func(int)) {
	defer IgnoreIndexOutOfRange()
	var cursor int
	for {
		f(cursor)
		cursor++
	}
	return
}

func (enum Enumerator) Copy() (r Enumerator) {
	r = Enumerator{
		Sequence: enum.Sequence,
		Span:     enum.Span,
		f:        enum.f,
	}
	return
}

func skipToChannelOffset(c R.Value, offset int) int {
	for open := true; open && offset > 0; offset-- {
		_, open = c.Recv()
	}
	return offset
}

func eachIndexable(enum *Enumerator, seq Indexable) (i int) {
	if l := Len(seq); l > 0 {
		switch f := enum.f.(type) {
		case func(interface{}):
			enum.each(func(cursor int) {
				f(seq.AtOffset(cursor))
			})
		case func(int, interface{}):
			enum.each(func(cursor int) {
				f(cursor, seq.AtOffset(cursor))
			})
		case func(interface{}, interface{}):
			enum.each(func(cursor int) {
				f(cursor, seq.AtOffset(cursor))
			})
		case func(R.Value):
			enum.each(func(cursor int) {
				f(R.ValueOf(seq.AtOffset(cursor)))
			})
		case func(int, R.Value):
			enum.each(func(cursor int) {
				f(cursor, R.ValueOf(seq.AtOffset(cursor)))
			})
		case func(interface{}, R.Value):
			enum.each(func(cursor int) {
				f(cursor, R.ValueOf(seq.AtOffset(cursor)))
			})
		case func(R.Value, R.Value):
			enum.each(func(cursor int) {
				f(R.ValueOf(cursor), R.ValueOf(seq.AtOffset(cursor)))
			})
		default:
			switch f := R.ValueOf(f); f.Kind() {
			case R.Func:
				if t := f.Type(); !t.IsVariadic() {
					switch t.NumIn() {
					case 1: //	f(v)
						p := make([]R.Value, 1, 1)
						enum.each(func(cursor int) {
							p[0] = R.ValueOf(seq.AtOffset(cursor))
							f.Call(p)
						})
					case 2: //	f(i, v)
						p := make([]R.Value, 2, 2)
						enum.each(func(cursor int) {
							p[0], p[1] = R.ValueOf(cursor), R.ValueOf(seq.AtOffset(cursor))
							f.Call(p)
						})
					default:
						panic(UNHANDLED_ITERATOR)
					}
				}
			case R.Chan:
				enum.each(func(cursor int) {
					f.Send(R.ValueOf(seq.AtOffset(cursor)))
				})
			case R.Slice:
				if f.Type().Elem().Kind() == R.Chan {
					n := f.Len()
					enum.each(func(cursor int) {
						v := R.ValueOf(seq.AtOffset(cursor))
						for i := 0; i < n; i++ {
							f.Index(i).Send(v)
						}
					})
				} else {
					panic(UNHANDLED_ITERATOR)
				}
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

/*
	Enumerators can automatically generate results from functions which take an integer input parameter
	and return an single output value of arbitrary type which should be dependent on this input index.
	Additionally the function may also return an error value which causes enumeration to terminate.

	The output value is passed as input to the enumerating function in the same way as with other
	enumerated types.

	Behaviour is undefined when the generator function manipulates variables over which it is closed,
	so avoid modifying such variables or incorrect results will occur.
*/
func eachFunction(enum *Enumerator, g R.Value) (i int) {
	if tg := g.Type(); tg.NumIn() == 1 && tg.NumOut() == 1 {
		p := make([]R.Value, 1, 1)
		switch f := enum.f.(type) {
		case func(interface{}):
			enum.each(func(cursor int) {
				p[0] = R.ValueOf(cursor)
				f(g.Call(p)[0].Interface())
			})
		case func(int, interface{}):
			enum.each(func(cursor int) {
				p[0] = R.ValueOf(cursor)
				f(cursor, g.Call(p)[0].Interface())
			})
		case func(interface{}, interface{}):
			enum.each(func(cursor int) {
				p[0] = R.ValueOf(cursor)
				f(cursor, g.Call(p)[0].Interface())
			})
		case func(R.Value):
			enum.each(func(cursor int) {
				p[0] = R.ValueOf(cursor)
				f(g.Call(p)[0])
			})
		case func(int, R.Value):
			enum.each(func(cursor int) {
				p[0] = R.ValueOf(cursor)
				f(cursor, g.Call(p)[0])
			})
		case func(interface{}, R.Value):
			enum.each(func(cursor int) {
				p[0] = R.ValueOf(cursor)
				f(cursor, g.Call(p)[0])
			})
		case func(R.Value, R.Value):
			enum.each(func(cursor int) {
				p[0] = R.ValueOf(cursor)
				f(R.ValueOf(cursor), g.Call(p)[0])
			})
		default:
			if f := R.ValueOf(f); f.Kind() == R.Func {
				if tf := f.Type(); !tf.IsVariadic() {
					switch tf.NumIn() {
					case 1: //	f(v)
						pf := make([]R.Value, 1, 1)
						enum.each(func(cursor int) {
							p[0] = R.ValueOf(cursor)
							pf[0] = g.Call(p)[0]
							f.Call(pf)
						})
					case 2: //	f(i, v)
						pf := make([]R.Value, 2, 2)
						enum.each(func(cursor int) {
							p[0] = R.ValueOf(cursor)
							pf[0], pf[1] = p[0], g.Call(p)[0]
							f.Call(pf)
						})
					default:
						panic(UNHANDLED_ITERATOR)
					}
				}
			}
		}
	}
	return
}
