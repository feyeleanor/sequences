package sequences

import R "reflect"

func (enum Enumerator) eachG1bool(g func(int) bool) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(bool):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, bool):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, bool):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...bool):
			pf := make([]bool, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1complex64(g func(int) complex64) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(complex64):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, complex64):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, complex64):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...complex64):
			pf := make([]complex64, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1complex128(g func(int) complex128) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(complex128):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, complex128):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, complex128):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...complex128):
			pf := make([]complex128, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1error(g func(int) error) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(error):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, error):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, error):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...error):
			pf := make([]error, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1float32(g func(int) float32) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(float32):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, float32):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, float32):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...float32):
			pf := make([]float32, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1float64(g func(int) float64) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(float64):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, float64):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, float64):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...float64):
			pf := make([]float64, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1int(g func(int) int) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, int):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, int):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...int):
			pf := make([]int, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1int8(g func(int) int8) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int8):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, int8):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, int8):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...int8):
			pf := make([]int8, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1int16(g func(int) int16) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int16):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, int16):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, int16):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...int16):
			pf := make([]int16, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1int32(g func(int) int32) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int32):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, int32):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, int32):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...int32):
			pf := make([]int32, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1int64(g func(int) int64) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int64):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, int64):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, int64):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...int64):
			pf := make([]int64, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1interface(g func(int) interface{}) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			//	Defer further processing to reflection
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1string(g func(int) string) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(string):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, string):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, string):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...string):
			pf := make([]string, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1uint(g func(int) uint) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, uint):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...uint):
			pf := make([]uint, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1uint8(g func(int) uint8) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint8):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, uint8):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint8):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...uint8):
			pf := make([]uint8, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1uint16(g func(int) uint16) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint16):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, uint16):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint16):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...uint16):
			pf := make([]uint16, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1uint32(g func(int) uint32) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint32):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, uint32):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint32):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...uint32):
			pf := make([]uint32, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1uint64(g func(int) uint64) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint64):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, uint64):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint64):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...uint64):
			pf := make([]uint64, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1uintptr(g func(int) uintptr) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uintptr):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, uintptr):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, uintptr):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...uintptr):
			pf := make([]uintptr, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, R.ValueOf(g(cursor)))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG1RValue(g func(int) R.Value) (i int, e error) {
	if enum.Start > 0 {
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(interface{}):
			for ; steps > 0; steps-- {
				f(g(cursor).Interface())
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor).Interface())
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor).Interface())
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor).Interface())
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				f(g(cursor))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				f(cursor, g(cursor))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				f(R.ValueOf(cursor), g(cursor))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				pf = append(pf, g(cursor))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2bool(g func(int) (bool, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	bool
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(bool):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, bool):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, bool):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...bool):
			pf := make([]bool, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2complex64(g func(int) (complex64, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	complex64
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(complex64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, complex64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, complex64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...complex64):
			pf := make([]complex64, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2complex128(g func(int) (complex128, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	complex128
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(complex128):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, complex128):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, complex128):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...complex128):
			pf := make([]complex128, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2error(g func(int) (error, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	error
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(error):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, error):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, error):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...error):
			pf := make([]error, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}
func (enum Enumerator) eachG2float32(g func(int) (float32, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	float32
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(float32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, float32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, float32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...float32):
			pf := make([]float32, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2float64(g func(int) (float64, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	float64
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(float64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, float64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, float64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...float64):
			pf := make([]float64, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2int(g func(int) (int, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	int
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, int):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, int):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...int):
			pf := make([]int, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2int8(g func(int) (int8, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	int8
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int8):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, int8):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, int8):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...int8):
			pf := make([]int8, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2int16(g func(int) (int16, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	int16
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int16):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, int16):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, int16):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...int16):
			pf := make([]int16, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2int32(g func(int) (int32, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	int32
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, int32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, int32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...int32):
			pf := make([]int32, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2int64(g func(int) (int64, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	int64
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(int64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, int64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, int64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...int64):
			pf := make([]int64, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2interface(g func(int) (interface{}, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	interface{}
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2string(g func(int) (string, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	string
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(string):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, string):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, string):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...string):
			pf := make([]string, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2uint(g func(int) (uint, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	uint
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, uint):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...uint):
			pf := make([]uint, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2uint8(g func(int) (uint8, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	uint8
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint8):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, uint8):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint8):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...uint8):
			pf := make([]uint8, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2uint16(g func(int) (uint16, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	uint16
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint16):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, uint16):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint16):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...uint16):
			pf := make([]uint16, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2uint32(g func(int) (uint32, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	uint32
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, uint32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint32):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...uint32):
			pf := make([]uint32, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2uint64(g func(int) (uint64, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	uint64
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uint64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, uint64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, uint64):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...uint64):
			pf := make([]uint64, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2uintptr(g func(int) (uintptr, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	uintptr
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(uintptr):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, uintptr):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, uintptr):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...uintptr):
			pf := make([]uintptr, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), R.ValueOf(v))
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, R.ValueOf(v))
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}

func (enum Enumerator) eachG2RValue(g func(int) (R.Value, error)) (i int, e error) {
	if enum.Start > 0 {
		var v	R.Value
		steps := enum.Limit
		cursor := enum.Start
		switch f := enum.f.(type) {
		case func(interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v.Interface())
				i++
				cursor += enum.Span
			}
		case func(int, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v.Interface())
				i++
				cursor += enum.Span
			}
		case func(interface{}, interface{}):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v.Interface())
				i++
				cursor += enum.Span
			}
		case func(...interface{}):
			pf := make([]interface{}, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v.Interface())
			}
			f(pf...)
			i = 1
		case func(R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(v)
				i++
				cursor += enum.Span
			}
		case func(int, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(interface{}, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(cursor, v)
				i++
				cursor += enum.Span
			}
		case func(R.Value, R.Value):
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				f(R.ValueOf(cursor), v)
				i++
				cursor += enum.Span
			}
		case func(...R.Value):
			pf := make([]R.Value, 0, steps)
			for ; steps > 0; steps-- {
				if v, e = g(cursor); e != nil {
					return
				}
				pf = append(pf, v)
				cursor += enum.Span
			}
			f(pf...)
			i = 1
		default:
			e = UNKNOWN_ITERATOR
		}
	}
	return
}