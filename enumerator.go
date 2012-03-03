package sequences

import R "reflect"

type Bounds struct {
	Start		int
	Span		int
	Limit		int
	cursor		int
	steps		int
}

func (b Bounds) Cursor() int {
	return b.cursor
}


type process struct {
	f			interface{}
	error
}


type Enumerator struct {
	Sequence	interface{}
	Bounds
	process
}

func (enum Enumerator) Copy() (r *Enumerator) {
	r = &Enumerator{
		Bounds:		enum.Bounds,
		Sequence:	enum.Sequence,
		process:	enum.process,
	}
	return
}

func (enum *Enumerator) Bind(f interface{}) (e error) {
	switch fv := R.ValueOf(f); fv.Kind() {
	case R.Func, R.Chan:
		enum.f = f
	case R.Slice:
		switch fv.Type().Elem().Kind() {
		case R.Func, R.Chan:
			enum.f = f
		default:
			e = NOT_AN_ITERATOR
		}
	default:
		e = NOT_AN_ITERATOR
	}
	return
}

func (enum *Enumerator) SetBounds(length int) {
	switch {
	case enum.Span == 0:
		return
	case enum.Span < 0:
		enum.cursor = length - 1
	default:
		enum.cursor = enum.Start
	}

	switch {
	case enum.Limit == 0:
		enum.steps = 0
		return
	case enum.Limit < 0:
		enum.steps = length / enum.Span
		if enum.steps < 0 {
			enum.steps = -enum.steps
		}
		if length % enum.Span != 0 {
			enum.steps++
		}
	default:
		enum.steps = length / enum.Span
		if enum.steps < 0 {
			enum.steps = -enum.steps
		}
		if length % enum.Span != 0 {
			enum.steps++
		}
		if enum.steps > enum.Limit {
			enum.steps = enum.Limit
		}
	}
	return
}

func (enum Enumerator) Execute() (i int, e error) {
	switch seq := enum.Sequence.(type) {
	case Enumerable:
		i, e = seq.Step(enum.Start, enum.Span, enum.Limit, enum.f)
	case []bool:
		i, e = enum.eachBoolSlice(seq)
	case []complex64:
		i, e = enum.eachComplex64Slice(seq)
	case []complex128:
		i, e = enum.eachComplex128Slice(seq)
	case []error:
		i, e = enum.eachErrorSlice(seq)
	case []float32:
		i, e = enum.eachFloat32Slice(seq)
	case []float64:
		i, e = enum.eachFloat64Slice(seq)
	case []int:
		i, e = enum.eachIntSlice(seq)
	case []int8:
		i, e = enum.eachInt8Slice(seq)
	case []int16:
		i, e = enum.eachInt16Slice(seq)
	case []int32:
		i, e = enum.eachInt32Slice(seq)
	case []int64:
		i, e = enum.eachInt64Slice(seq)
	case []interface{}:
		i, e = enum.eachInterfaceSlice(seq)
	case []string:
		i, e = enum.eachStringSlice(seq)
	case []uint:
		i, e = enum.eachUintSlice(seq)
	case []uint8:
		i, e = enum.eachUint8Slice(seq)
	case []uint16:
		i, e = enum.eachUint16Slice(seq)
	case []uint32:
		i, e = enum.eachUint32Slice(seq)
	case []uint64:
		i, e = enum.eachUint64Slice(seq)
	case []uintptr:
		i, e = enum.eachUintptrSlice(seq)
	case []R.Value:
		i, e = enum.eachRValueSlice(seq)
	case chan bool:
		i, e = enum.eachBoolChannel(seq)
	case Indexable:
		i, e = enum.eachIndexable(seq)
	case R.Value:
		switch seq.Kind() {
		case R.Slice:
			i, e = enum.eachSlice(seq)				
		case R.Chan:
			if enum.Start < 0 {
				e = ASCENDING_SEQUENCE
			} else {
				i, e = enum.eachChannel(seq)
			}
		case R.Func:
			i, e = enum.eachFunction(seq)
		}
	default:
		switch c := R.ValueOf(seq); c.Kind() {
		case R.Slice:
			i, e = enum.eachSlice(c)				
		case R.Chan:
			if enum.Start < 0 {
				e = ASCENDING_SEQUENCE
			} else {
				i, e = enum.eachChannel(c)
			}
		case R.Func:
			i, e = enum.eachFunction(c)
		}
	}
	return
}

func (enum *Enumerator) Each(f interface{}) (i int, e error) {
	if e = enum.Bind(f); e == nil {
		i, e = enum.Execute()
	}
	return
}

func skipToChannelOffset(c R.Value, offset int) int {
	for open := true; open && offset > 0; offset-- {
		_, open = c.Recv()
	}
	return offset
}


func (enum Enumerator) eachIndexable(seq Indexable) (i int, e error) {
	var steps	int
	var l	int
	if l, e = Len(seq); e == nil {
		enum.SetBounds(l)
		switch f := enum.f.(type) {
		case func(interface{}):
			for steps = enum.steps; steps > 0; steps-- {
				f(seq.AtOffset(enum.cursor))
				enum.cursor += enum.Span
			}
			i = enum.steps
		case func(int, interface{}):
			for steps = enum.steps; steps > 0; steps-- {
				f(enum.cursor, seq.AtOffset(enum.cursor))
				enum.cursor += enum.Span
			}
			i = enum.steps
		case func(interface{}, interface{}):
			for steps = enum.steps; steps > 0; steps-- {
				f(enum.cursor, seq.AtOffset(enum.cursor))
				enum.cursor += enum.Span
			}
			i = enum.steps
		case func(...interface{}):
			s := make([]interface{}, enum.steps, enum.steps)
			for steps = enum.steps; steps > 0; steps-- {
				s[i] = seq.AtOffset(enum.cursor)
				enum.cursor += enum.Span
			}
			f(s...)
			i = 1
		case func(R.Value):
			for steps = enum.steps; steps > 0; steps-- {
				f(R.ValueOf(seq.AtOffset(enum.cursor)))
				enum.cursor += enum.Span
			}
			i = enum.steps
		case func(int, R.Value):
			for steps = enum.steps; steps > 0; steps-- {
				f(enum.cursor, R.ValueOf(seq.AtOffset(enum.cursor)))
				enum.cursor += enum.Span
			}
			i = enum.steps
		case func(interface{}, R.Value):
			for steps = enum.steps; steps > 0; steps-- {
				f(enum.cursor, R.ValueOf(seq.AtOffset(enum.cursor)))
				enum.cursor += enum.Span
			}
			i = enum.steps
		case func(R.Value, R.Value):
			for steps = enum.steps; steps > 0; steps-- {
				f(R.ValueOf(enum.cursor), R.ValueOf(seq.AtOffset(enum.cursor)))
				enum.cursor += enum.Span
			}
			i = enum.steps
		case func(...R.Value):
			s := make([]R.Value, enum.steps, enum.steps)
			for steps = enum.steps; steps > 0; steps-- {
				s[i] = R.ValueOf(seq.AtOffset(enum.cursor))
				enum.cursor += enum.Span
			}
			f(s...)
			i = 1
		default:
			switch f := R.ValueOf(f); f.Kind() {
			case R.Func:
				if t := f.Type(); t.IsVariadic() {
					// f(...V)
					p := make([]R.Value, enum.steps, enum.steps)
					for steps = enum.steps; steps > 0; steps-- {
						p[i] = R.ValueOf(seq.AtOffset(i))
						enum.cursor += enum.Span
					}
					f.Call(p)
					i = 1
				} else {
					switch t.NumIn() {
					case 1:				//	f(v)
						p := make([]R.Value, 1, 1)
						for steps = enum.steps; steps > 0; steps-- {
							p[0] = R.ValueOf(seq.AtOffset(enum.cursor))
							f.Call(p)
							enum.cursor += enum.Span
						}
						i = enum.steps
					case 2:				//	f(i, v)
						p := make([]R.Value, 2, 2)
						for steps = enum.steps; steps > 0; steps-- {
							p[0], p[1] = R.ValueOf(enum.cursor), R.ValueOf(seq.AtOffset(enum.cursor))
							f.Call(p)
							enum.cursor += enum.Span
						}
						i = enum.steps
					default:
						e = UNKNOWN_ITERATOR
					}
				}
			case R.Chan:
				for steps = enum.steps; steps > 0; steps-- {
					f.Send(R.ValueOf(seq.AtOffset(enum.cursor)))
					enum.cursor += enum.Span
				}
				i = enum.steps
			case R.Slice:
				if f.Type().Elem().Kind() == R.Chan {
					n := f.Len()
					for steps = enum.steps; steps > 0; steps-- {
						for i := n; i > 0; i-- {
							f.Index(i).Send(R.ValueOf(seq.AtOffset(enum.cursor)))
						}
						enum.cursor += enum.Span
					}
					i = enum.steps
				} else {
					e = UNKNOWN_ITERATOR
				}
			default:
				e = UNKNOWN_ITERATOR
			}
		}
	}
	return
}

func (enum Enumerator) eachBoolSlice(seq []bool) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(bool):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, bool):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, bool):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...bool):
		s := make([]bool, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan bool:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan bool:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachComplex64Slice(seq []complex64) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(complex64):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, complex64):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, complex64):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...complex64):
		s := make([]complex64, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan complex64:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan complex64:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachComplex128Slice(seq []complex128) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(complex128):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, complex128):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, complex128):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...complex128):
		s := make([]complex128, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan complex128:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan complex128:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachErrorSlice(seq []error) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(error):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, error):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, error):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...error):
		s := make([]error, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan error:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan error:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachFloat32Slice(seq []float32) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(float32):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, float32):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, float32):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...float32):
		s := make([]float32, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan float32:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan float32:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachFloat64Slice(seq []float64) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(float64):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, float64):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, float64):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...float64):
		s := make([]float64, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan float64:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan float64:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachIntSlice(seq []int) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(int):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, int):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, int):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...int):
		s := make([]int, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan int:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan int:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInt8Slice(seq []int8) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(int8):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, int8):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, int8):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...int8):
		s := make([]int8, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan int8:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan int8:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInt16Slice(seq []int16) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(int16):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, int16):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, int16):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...int16):
		s := make([]int16, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan int16:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan int16:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInt32Slice(seq []int32) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(int32):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, int32):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, int32):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...int32):
		s := make([]int32, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan int32:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan int32:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInt64Slice(seq []int64) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(int64):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, int64):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, int64):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...int64):
		s := make([]int64, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan int64:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan int64:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInterfaceSlice(seq []interface{}) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachStringSlice(seq []string) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(string):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, string):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, string):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...string):
		s := make([]string, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan string:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan string:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUintSlice(seq []uint) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(uint):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, uint):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, uint):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...uint):
		s := make([]uint, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan uint:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan uint:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUint8Slice(seq []uint8) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(uint8):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, uint8):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, uint8):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...uint8):
		s := make([]uint8, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan uint8:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan uint8:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUint16Slice(seq []uint16) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(uint16):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, uint16):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, uint16):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...uint16):
		s := make([]uint16, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan uint16:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan uint16:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUint32Slice(seq []uint32) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(uint32):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, uint32):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, uint32):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...uint32):
		s := make([]uint32, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan uint32:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan uint32:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUint64Slice(seq []uint64) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(uint64):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, uint64):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, uint64):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...uint64):
		s := make([]uint64, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan uint64:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan uint64:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUintptrSlice(seq []uintptr) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(uintptr):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, uintptr):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, uintptr):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...uintptr):
		s := make([]uintptr, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, R.ValueOf(seq[enum.cursor]))
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan uintptr:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- R.ValueOf(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan uintptr:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- R.ValueOf(v)
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachRValueSlice(seq []R.Value) (i int, e error) {
	var steps	int
	enum.SetBounds(len(seq))
	switch f := enum.f.(type) {
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor].Interface())
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor].Interface())
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor].Interface())
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		s := make([]interface{}, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), seq[enum.cursor])
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		s := make([]R.Value, 0, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			s = append(s, seq[enum.cursor])
			enum.cursor += enum.Span
		}
		f(s...)
		i = 1
	case chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor].Interface()
			enum.cursor += enum.Span
		}
		i = enum.steps
	case chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			f <- seq[enum.cursor]
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan interface{}:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor].Interface()
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	case []chan R.Value:
		for steps = enum.steps; steps > 0; steps-- {
			v := seq[enum.cursor]
			for _, c := range f {
				c <- v
			}
			enum.cursor += enum.Span
		}
		i = enum.steps
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachSlice(s R.Value) (i int, e error) {
	var steps	int
	enum.SetBounds(s.Len())
	switch f := enum.f.(type) {
	case func(interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(s.Index(enum.cursor).Interface())
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, s.Index(enum.cursor).Interface())
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, interface{}):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, s.Index(enum.cursor).Interface())
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...interface{}):
		p := make([]interface{}, enum.steps, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			p[i] = s.Index(enum.cursor).Interface()
			enum.cursor += enum.Span
		}
		f(p...)
		i = 1
	case func(R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(s.Index(enum.cursor))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(int, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, s.Index(enum.cursor))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(interface{}, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(enum.cursor, s.Index(enum.cursor))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(R.Value, R.Value):
		for steps = enum.steps; steps > 0; steps-- {
			f(R.ValueOf(enum.cursor), s.Index(enum.cursor))
			enum.cursor += enum.Span
		}
		i = enum.steps
	case func(...R.Value):
		p := make([]R.Value, enum.steps, enum.steps)
		for steps = enum.steps; steps > 0; steps-- {
			p[i] = s.Index(enum.cursor)
			enum.cursor += enum.Span
		}
		f(p...)
		i = 1
	default:
		switch f := R.ValueOf(f); f.Kind() {
		case R.Func:
			if t := f.Type(); t.IsVariadic() {
				//	f(...v)
				steps = enum.steps
				p := make([]R.Value, 0, steps)
				for ; steps > 0; steps-- {
					p = append(p, s.Index(enum.cursor))
					enum.cursor += enum.Span
				}
				f.Call(p)
				i = 1
			} else {
				switch t.NumIn() {
				case 1:				//	f(v)
					p := make([]R.Value, 1, 1)
					for steps = enum.steps; steps > 0; steps-- {
						p[0] = s.Index(enum.cursor)
						f.Call(p)
						enum.cursor += enum.Span
					}
					i = enum.steps
				case 2:				//	f(i, v)
					p := make([]R.Value, 2, 2)
					for steps = enum.steps; steps > 0; steps-- {
						p[0], p[1] = R.ValueOf(enum.cursor), s.Index(enum.cursor)
						f.Call(p)
						enum.cursor += enum.Span
					}
					i = enum.steps
				default:
					e = UNKNOWN_ITERATOR
				}
			}
		case R.Chan:
			for steps = enum.steps; steps > 0; steps-- {
				f.Send(s.Index(enum.cursor))
				enum.cursor += enum.Span
			}
			i = enum.steps
		case R.Slice:
			if f.Type().Elem().Kind() == R.Chan {
				n := f.Len()
				for steps = enum.steps; steps > 0; steps-- {
					for i := n; i > 0; i-- {
						f.Index(i).Send(s.Index(enum.cursor))
					}
					enum.cursor += enum.Span
				}
				i = enum.steps
			} else {
				e = UNKNOWN_ITERATOR
			}
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachBoolChannel(seq chan bool) (i int, e error) {
	var v		bool
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan bool, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(bool):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, bool):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, bool):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...bool):
		if skipToOffset(seq, enum.Start) {
			s := make([]bool, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan bool:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan bool:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachComplex64Channel(seq chan complex64) (i int, e error) {
	var v		complex64
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan complex64, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(complex64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, complex64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, complex64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...complex64):
		if skipToOffset(seq, enum.Start) {
			s := make([]complex64, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan complex64:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan complex64:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachComplex128Channel(seq chan complex128) (i int, e error) {
	var v		complex128
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan complex128, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(complex128):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, complex128):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, complex128):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...complex128):
		if skipToOffset(seq, enum.Start) {
			s := make([]complex128, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan complex128:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan complex128:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachErrorChannel(seq chan error) (i int, e error) {
	var v		error
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan error, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(error):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, error):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, error):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...error):
		if skipToOffset(seq, enum.Start) {
			s := make([]error, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan error:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan error:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachFloat32Channel(seq chan float32) (i int, e error) {
	var v		float32
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan float32, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(float32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, float32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, float32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...float32):
		if skipToOffset(seq, enum.Start) {
			s := make([]float32, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan float32:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan float32:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachFloat64Channel(seq chan float64) (i int, e error) {
	var v		float64
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan float64, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(float64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, float64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, float64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...float64):
		if skipToOffset(seq, enum.Start) {
			s := make([]float64, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan float64:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan float64:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachIntChannel(seq chan int) (i int, e error) {
	var v		int
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan int, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(int):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, int):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, int):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...int):
		if skipToOffset(seq, enum.Start) {
			s := make([]int, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan int:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan int:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInt8Channel(seq chan int8) (i int, e error) {
	var v		int8
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan int8, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(int8):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, int8):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, int8):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...int8):
		if skipToOffset(seq, enum.Start) {
			s := make([]int8, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan int8:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan int8:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInt16Channel(seq chan int16) (i int, e error) {
	var v		int16
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan int16, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(int16):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, int16):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, int16):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...int16):
		if skipToOffset(seq, enum.Start) {
			s := make([]int16, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan int16:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan int16:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInt32Channel(seq chan int32) (i int, e error) {
	var v		int32
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan int32, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(int32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, int32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, int32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...int32):
		if skipToOffset(seq, enum.Start) {
			s := make([]int32, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan int32:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan int32:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInt64Channel(seq chan int64) (i int, e error) {
	var v		int64
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan int64, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(int64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, int64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, int64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...int64):
		if skipToOffset(seq, enum.Start) {
			s := make([]int64, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan int64:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan int64:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachInterfaceChannel(seq chan interface{}) (i int, e error) {
	var v		interface{}
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan interface{}, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachStringChannel(seq chan string) (i int, e error) {
	var v		string
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan string, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(string):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, string):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, string):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...string):
		if skipToOffset(seq, enum.Start) {
			s := make([]string, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan string:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan string:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUintChannel(seq chan uint) (i int, e error) {
	var v		uint
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan uint, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(uint):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, uint):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, uint):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...uint):
		if skipToOffset(seq, enum.Start) {
			s := make([]uint, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan uint:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan uint:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUint8Channel(seq chan uint8) (i int, e error) {
	var v		uint8
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan uint8, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(uint8):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, uint8):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, uint8):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...uint8):
		if skipToOffset(seq, enum.Start) {
			s := make([]uint8, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan uint8:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan uint8:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUint16Channel(seq chan uint16) (i int, e error) {
	var v		uint16
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan uint16, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(uint16):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, uint16):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, uint16):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...uint16):
		if skipToOffset(seq, enum.Start) {
			s := make([]uint16, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan uint16:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan uint16:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUint32Channel(seq chan uint32) (i int, e error) {
	var v		uint32
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan uint32, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(uint32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, uint32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, uint32):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...uint32):
		if skipToOffset(seq, enum.Start) {
			s := make([]uint32, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan uint32:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan uint32:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUint64Channel(seq chan uint64) (i int, e error) {
	var v		uint64
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan uint64, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(uint64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, uint64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, uint64):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...uint64):
		if skipToOffset(seq, enum.Start) {
			s := make([]uint64, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan uint64:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan uint64:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachUintptrChannel(seq chan uintptr) (i int, e error) {
	var v		uintptr
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan uintptr, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(uintptr):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, uintptr):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, uintptr):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...uintptr):
		if skipToOffset(seq, enum.Start) {
			s := make([]uintptr, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), R.ValueOf(v))
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, R.ValueOf(v))
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan uintptr:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- R.ValueOf(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan uintptr:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachRValueChannel(seq chan R.Value) (i int, e error) {
	var v		R.Value
	var steps	int

	switch {
	case enum.Span < 1:
		e = ASCENDING_SEQUENCE
	case enum.Limit < 0:
		enum.SetBounds(_MAXINT_)
	default:
		enum.SetBounds(enum.Limit + 1)
	}

	skipToOffset := func(c chan R.Value, offset int) (open bool) {
		for open = true; open && offset > 1; offset-- {
			_, open = <- seq
			enum.cursor++
		}
		return
	}
	switch f := enum.f.(type) {
	case func(interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v.Interface())
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v.Interface())
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, interface{}):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v.Interface())
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...interface{}):
		if skipToOffset(seq, enum.Start) {
			s := make([]interface{}, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v.Interface())
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case func(R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(int, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(interface{}, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(enum.cursor, v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(R.Value, R.Value):
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f(R.ValueOf(enum.cursor), v)
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case func(...R.Value):
		if skipToOffset(seq, enum.Start) {
			s := make([]R.Value, 0, 4)
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					s = append(s, v)
					enum.cursor++
					open = skipToOffset(seq, enum.Span)
				}
			}
			f(s...)
			i = 1
		}
	case chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v.Interface()
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
			for open := true; open && steps > 0; steps-- {
				if v, open = <- seq; open {
					f <- v
					enum.cursor++
					i++
					open = skipToOffset(seq, enum.Span)
				}
			}
		}
	case []chan interface{}:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	case []chan R.Value:
		if skipToOffset(seq, enum.Start) {
			steps = enum.steps
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
		}
	default:
		e = UNKNOWN_ITERATOR
	}
	return
}

func (enum Enumerator) eachChannel(c R.Value) (i int, e error) {
	var steps	int
	if enum.Span < 0 {
		return 0, ASCENDING_SEQUENCE
	}
	enum.SetBounds(_MAXINT_)
println("limit:", _MAXINT_)
println("enum.cursor:", enum.cursor)
println("enum.steps:", enum.steps)
	switch f := enum.f.(type) {
	case func(interface{}):
		if enum.steps > 0 {
			steps = enum.steps
			for v, open := c.Recv(); open && steps > 0; steps-- {
				f(v.Interface())
				skipToChannelOffset(c, enum.Span)
				i++
				v, open = c.Recv()
			}
		} else {
			for v, open := c.Recv(); open; {
				f(v.Interface())
				skipToChannelOffset(c, enum.Span)
				i++
				v, open = c.Recv()
			}			
		}
	case func(int, interface{}):
		if enum.steps > 0 {
			steps = enum.steps
			for v, open := c.Recv(); open && steps > 0; steps-- {
				f(enum.cursor, v.Interface())
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
			i = enum.steps - steps
		} else {
			for v, open := c.Recv(); open; {
				f(enum.cursor, v.Interface())
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
		}
	case func(interface{}, interface{}):
		if enum.steps > 0 {
			steps = enum.steps
			for v, open := c.Recv(); open && steps > 0; steps-- {
				f(enum.cursor, v.Interface())
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
			i = enum.steps - steps
		} else {
			for v, open := c.Recv(); open; {
				f(enum.cursor, v.Interface())
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
		}
	case func(...interface{}):
		var p	[]interface{}
		if enum.steps > 0 {
			steps = enum.steps
			p = make([]interface{}, 0, steps)
			for v, open := c.Recv(); open && steps > 0; steps-- {
				p = append(p, v.Interface())
				skipToChannelOffset(c, enum.Span)
				v, open = c.Recv()
			}
			i = enum.steps - steps
		} else {
			p = make([]interface{}, 0, 4)
			for v, open := c.Recv(); open; {
				p = append(p, v.Interface())
				skipToChannelOffset(c, enum.Span)
				v, open = c.Recv()
			}
		}
		f(p...)
		i = 1
	case func(R.Value):
		if enum.steps > 0 {
			steps = enum.steps
			for v, open := c.Recv(); open && steps > 0; steps-- {
				f(v)
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
			i = enum.steps - steps
		} else {
			for v, open := c.Recv(); open; {
				f(v)
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
		}
	case func(int, R.Value):
		if enum.steps > 0 {
			steps = enum.steps
			for v, open := c.Recv(); open && steps > 0; steps-- {
				f(enum.cursor, v)
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
			i = enum.steps - steps
		} else {
			for v, open := c.Recv(); open; {
				f(enum.cursor, v)
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
		}
	case func(interface{}, R.Value):
		if enum.steps > 0 {
			steps = enum.steps
			for v, open := c.Recv(); open && steps > 0; steps-- {
				f(enum.cursor, v)
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
			i = enum.steps - steps
		} else {
			for v, open := c.Recv(); open; {
				f(enum.cursor, v)
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
		}
	case func(R.Value, R.Value):
		if enum.steps > 0 {
			steps = enum.steps
			for v, open := c.Recv(); open && steps > 0; steps-- {
				f(R.ValueOf(enum.cursor), v)
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
			i = enum.steps - steps
		} else {
			for v, open := c.Recv(); open; {
				f(R.ValueOf(enum.cursor), v)
				skipToChannelOffset(c, enum.Span)
				enum.cursor += enum.Span
				i++
				v, open = c.Recv()
			}
		}
	case func(...R.Value):
		var p	[]R.Value
		if enum.steps > 0 {
			steps = enum.steps
			p = make([]R.Value, 0, steps)
			for v, open := c.Recv(); open && steps > 0; steps-- {
				p = append(p, v)
				skipToChannelOffset(c, enum.Span)
				v, open = c.Recv()
			}
			i = enum.steps - steps
		} else {
			p = make([]R.Value, 0, 4)
			for v, open := c.Recv(); open; {
				p = append(p, v)
				skipToChannelOffset(c, enum.Span)
				v, open = c.Recv()
			}
		}
		f(p...)
		i = 1
	case []chan interface{}:
		panic("[]chan interface{} iterator not yet implemented")
	case []chan R.Value:
		panic("[]chan R.Value iterator not yet implemented")
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); t.IsVariadic() {
				//	f(...v)
				var p	[]R.Value
				if enum.steps > 0 {
					steps = enum.steps
					p = make([]R.Value, 0, enum.steps)
					steps = enum.steps
					for v, open := c.Recv(); open && steps > 0; steps-- {
						p = append(p, v)
						skipToChannelOffset(c, enum.Span)
						v, open = c.Recv()
					}
					i = enum.steps - steps
				} else {
					p = make([]R.Value, 0, 4)
					for v, open := c.Recv(); open; {
						p = append(p, v)
						skipToChannelOffset(c, enum.Span)
						v, open = c.Recv()
					}
				}
				f.Call(p)
				i = 1
			} else {
				switch t.NumIn() {
				case 1:				//	f(v)
					p := make([]R.Value, 1, 1)
					if enum.steps > 0 {
						steps = enum.steps
						for v, open := c.Recv(); open && steps > 0; steps-- {
							p[0] = v
							f.Call(p)
							skipToChannelOffset(c, enum.Span)
							i++
							v, open = c.Recv()
						}
					} else {
						for v, open := c.Recv(); open; {
							p[0] = v
							f.Call(p)
							skipToChannelOffset(c, enum.Span)
							i++
							v, open = c.Recv()
						}
					}					
				case 2:				//	f(i, v)
					p := make([]R.Value, 2, 2)
					if enum.steps > 0 {
						steps = enum.steps
						for v, open := c.Recv(); open && steps > 0; steps-- {
							p[0], p[1] = R.ValueOf(enum.cursor), v
							f.Call(p)
							skipToChannelOffset(c, enum.Span)
							enum.cursor += enum.Span
							i++
							v, open = c.Recv()
						}
					} else {
						for v, open := c.Recv(); open; {
							p[0], p[1] = R.ValueOf(enum.cursor), v
							f.Call(p)
							skipToChannelOffset(c, enum.Span)
							enum.cursor += enum.Span
							i++
							v, open = c.Recv()
						}
					}
				default:
					e = UNKNOWN_ITERATOR
				}
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
	so avoid modifying such variables or incorrect results will occur. This limitation is imposed to
	prevent the need to iterate over all indices prior to the start index.
*/
func (enum Enumerator) eachFunction(g R.Value) (i int, e error) {
	var steps	int
	if enum.Start > 0 {
		enum.SetBounds(_MAXINT_)
		if tg := g.Type(); tg.NumIn() == 1 {
			switch tg.NumOut() {
			case 1:		//	f(x) -> y
				p := make([]R.Value, 1, 1)
				switch f := enum.f.(type) {
				case func(interface{}):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						f(g.Call(p)[0].Interface())
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(int, interface{}):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						f(enum.cursor, g.Call(p)[0].Interface())
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(interface{}, interface{}):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						f(enum.cursor, g.Call(p)[0].Interface())
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(...interface{}):
					steps = enum.steps
					pf := make([]interface{}, 0, steps)
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						pf = append(pf, g.Call(p)[0].Interface())
						enum.cursor += enum.Span
					}
					f(pf...)
					i = 1
				case func(R.Value):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						f(g.Call(p)[0])
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(int, R.Value):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						f(enum.cursor, g.Call(p)[0])
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(interface{}, R.Value):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						f(enum.cursor, g.Call(p)[0])
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(R.Value, R.Value):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						f(R.ValueOf(enum.cursor), g.Call(p)[0])
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(...R.Value):
					steps = enum.steps
					pf := make([]R.Value, 0, steps)
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						pf = append(pf, g.Call(p)[0])
						enum.cursor += enum.Span
					}
					f(pf...)
					i = 1
				default:
					if f := R.ValueOf(f); f.Kind() == R.Func {
						if tf := f.Type(); tf.IsVariadic() {
							//	f(...v)
							steps = enum.steps
							pf := make([]R.Value, 0, steps)
							for ; steps > 0; steps-- {
								p[0] = R.ValueOf(enum.cursor)
								pf = append(pf, g.Call(p)[0])
								enum.cursor += enum.Span
							}
							f.Call(pf)
							i = 1
						} else {
							switch tf.NumIn() {
							case 1:		//	f(v)
								pf := make([]R.Value, 1, 1)
								steps = enum.steps
								for ; steps > 0; steps-- {
									p[0] = R.ValueOf(enum.cursor)
									pf[0] = g.Call(p)[0]
									f.Call(pf)
									enum.cursor += enum.Span
								}
								i = enum.steps
							case 2:		//	f(i, v)
								pf := make([]R.Value, 2, 2)
								steps = enum.steps
								for ; steps > 0; steps-- {
									p[0] = R.ValueOf(enum.cursor)
									pf[0], pf[1] = p[0], g.Call(p)[0]
									f.Call(pf)
									enum.cursor += enum.Span
								}
								i = enum.steps
							default:
								e = UNKNOWN_ITERATOR
							}
						}
					}
				}

			case 2:		//	f(x) -> y, error
				p := make([]R.Value, 1, 1)
				switch f := enum.f.(type) {
				case func(interface{}):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						f(v[0].Interface())
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(int, interface{}):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						f(enum.cursor, v[0].Interface())
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(interface{}, interface{}):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						f(enum.cursor, v[0].Interface())
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(...interface{}):
					pf := make([]interface{}, 0, enum.steps)
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						pf = append(pf, v[0].Interface())
					}
					f(pf...)
					i = 1
				case func(R.Value):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						f(v[0])
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(int, R.Value):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						f(enum.cursor, v[0])
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(interface{}, R.Value):
					for ; enum.steps > 0; enum.steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						f(enum.cursor, v[0])
						enum.cursor += enum.Span
					}
				case func(R.Value, R.Value):
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						f(R.ValueOf(enum.cursor), v[0])
						enum.cursor += enum.Span
					}
					i = enum.steps
				case func(...R.Value):
					pf := make([]R.Value, 0, enum.steps)
					steps = enum.steps
					for ; steps > 0; steps-- {
						p[0] = R.ValueOf(enum.cursor)
						v := g.Call(p)
						if v[1].Bool() {
							break
						}
						pf = append(pf, v[0])
						enum.cursor += enum.Span
					}
					f(pf...)
					i = 1
				default:
					if f := R.ValueOf(f); f.Kind() == R.Func {
						if tf := f.Type(); tf.IsVariadic() {
							//	f(...v)
							steps = enum.steps
							pf := make([]R.Value, 0, steps)
							for ; steps > 0; steps-- {
								p[0] = R.ValueOf(enum.cursor)
								v := g.Call(p)
								if v[1].Bool() {
									break
								}
								pf = append(pf, v[0])
								enum.cursor += enum.Span
							}
							f.Call(pf)
							i = 1
						} else {
							switch tf.NumIn() {
							case 1:		//	f(v)
								pf := make([]R.Value, 1, 1)
								steps = enum.steps
								for ; steps > 0; steps-- {
									p[0] = R.ValueOf(enum.cursor)
									v := g.Call(p)
									if v[1].Bool() {
										break
									}
									f.Call(pf)
									enum.cursor += enum.Span
								}
								i = enum.steps
							case 2:		//	f(i, v)
								pf := make([]R.Value, 2, 2)
								steps = enum.steps
								for ; steps > 0; steps-- {
									p[0] = R.ValueOf(enum.cursor)
									v := g.Call(p)
									if v[1].Bool() {
										break
									}
									pf[0], pf[1] = p[0], v[0]
									f.Call(pf)
									enum.cursor += enum.Span
								}
								i = enum.steps
							default:
								e = UNKNOWN_ITERATOR
							}
						}
					}
				}
			}
		}
	}
	return
}