package sequences

import R "reflect"

func stepBool(container []bool, step int, f interface{}) (i int) {
	if step == 1 {
		var v	bool
		switch f := f.(type) {
		case func(bool):						for i, v = range container { f(v) }
		case func(int, bool):					for i, v = range container { f(i, v) }
		case func(...bool):						f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(bool):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, bool):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...bool):						s := make([]bool, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepComplex64(container []complex64, step int, f interface{}) (i int) {
	if step == 1 {
		var v	complex64
		switch f := f.(type) {
		case func(complex64):					for i, v = range container { f(v) }
		case func(int, complex64):				for i, v = range container { f(i, v) }
		case func(...complex64):				f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(complex64):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, complex64):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...complex64):				s := make([]complex64, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepComplex128(container []complex128, step int, f interface{}) (i int) {
	if step == 1 {
		var v	complex128
		switch f := f.(type) {
		case func(complex128):					for i, v = range container { f(v) }
		case func(int, complex128):				for i, v = range container { f(i, v) }
		case func(...complex128):				f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(complex128):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, complex128):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...complex128):				s := make([]complex128, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepError(container []error, step int, f interface{}) (i int) {
	if step == 1 {
		var v	error
		switch f := f.(type) {
		case func(error):						for i, v = range container { f(v) }
		case func(int, error):					for i, v = range container { f(i, v) }
		case func(...error):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(error):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, error):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...error):					s := make([]error, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepFloat32(container []float32, step int, f interface{}) (i int) {
	if step == 1 {
		var v	float32
		switch f := f.(type) {
		case func(float32):						for i, v = range container { f(v) }
		case func(int, float32):				for i, v = range container { f(i, v) }
		case func(...float32):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(float32):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, float32):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...float32):					s := make([]float32, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepFloat64(container []float64, step int, f interface{}) (i int) {
	if step == 1 {
		var v	float64
		switch f := f.(type) {
		case func(float64):						for i, v = range container { f(v) }
		case func(int, float64):				for i, v = range container { f(i, v) }
		case func(...float64):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(float64):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, float64):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...float64):					s := make([]float64, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepInt(container []int, step int, f interface{}) (i int) {
	if step == 1 {
		var v	int
		switch f := f.(type) {
		case func(int):							for i, v = range container { f(v) }
		case func(int, int):					for i, v = range container { f(i, v) }
		case func(...int):						f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(int):							for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, int):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...int):						s := make([]int, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepInt8(container []int8, step int, f interface{}) (i int) {
	if step == 1 {
		var v	int8
		switch f := f.(type) {
		case func(int8):						for i, v = range container { f(v) }
		case func(int, int8):					for i, v = range container { f(i, v) }
		case func(...int8):						f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(int8):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, int8):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...int8):						s := make([]int8, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepInt16(container []int16, step int, f interface{}) (i int) {
	if step == 1 {
		var v	int16
		switch f := f.(type) {
		case func(int16):						for i, v = range container { f(v) }
		case func(int, int16):					for i, v = range container { f(i, v) }
		case func(...int16):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(int16):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, int16):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...int16):					s := make([]int16, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepInt32(container []int32, step int, f interface{}) (i int) {
	if step == 1 {
		var v	int32
		switch f := f.(type) {
		case func(int32):						for i, v = range container { f(v) }
		case func(int, int32):					for i, v = range container { f(i, v) }
		case func(...int32):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(int32):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, int32):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...int32):					s := make([]int32, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepInt64(container []int64, step int, f interface{}) (i int) {
	if step == 1 {
		var v	int64
		switch f := f.(type) {
		case func(int64):						for i, v = range container { f(v) }
		case func(int, int64):					for i, v = range container { f(i, v) }
		case func(...int64):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(int64):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, int64):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...int64):					s := make([]int64, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepInterface(container []interface{}, step int, f interface{}) (i int) {
	if step == 1 {
		var v	interface{}
		switch f := f.(type) {
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				f(container...)
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepString(container []string, step int, f interface{}) (i int) {
	if step == 1 {
		var v	string
		switch f := f.(type) {
		case func(string):						for i, v = range container { f(v) }
		case func(int, string):					for i, v = range container { f(i, v) }
		case func(...string):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(string):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, string):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...string):					s := make([]string, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepUint(container []uint, step int, f interface{}) (i int) {
	if step == 1 {
		var v	uint
		switch f := f.(type) {
		case func(uint):						for i, v = range container { f(v) }
		case func(int, uint):					for i, v = range container { f(i, v) }
		case func(...uint):						f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(uint):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, uint):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...uint):						s := make([]uint, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}
	
		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepUint8(container []uint8, step int, f interface{}) (i int) {
	if step == 1 {
		var v	uint8
		switch f := f.(type) {
		case func(uint8):						for i, v = range container { f(v) }
		case func(int, uint8):					for i, v = range container { f(i, v) }
		case func(...uint8):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(uint8):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, uint8):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...uint8):					s := make([]uint8, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepUint16(container []uint16, step int, f interface{}) (i int) {
	if step == 1 {
		var v	uint16
		switch f := f.(type) {
		case func(uint16):						for i, v = range container { f(v) }
		case func(int, uint16):					for i, v = range container { f(i, v) }
		case func(...uint16):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(uint16):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, uint16):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...uint16):					s := make([]uint16, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepUint32(container []uint32, step int, f interface{}) (i int) {
	if step == 1 {
		var v	uint32
		switch f := f.(type) {
		case func(uint32):						for i, v = range container { f(v) }
		case func(int, uint32):					for i, v = range container { f(i, v) }
		case func(...uint32):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(uint32):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, uint32):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...uint32):					s := make([]uint32, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepUint64(container []uint64, step int, f interface{}) (i int) {
	if step == 1 {
		var v	uint64
		switch f := f.(type) {
		case func(uint64):						for i, v = range container { f(v) }
		case func(int, uint64):					for i, v = range container { f(i, v) }
		case func(...uint64):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(uint64):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, uint64):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...uint64):					s := make([]uint64, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepUintptr(container []uintptr, step int, f interface{}) (i int) {
	if step == 1 {
		var v	uintptr
		switch f := f.(type) {
		case func(uintptr):						for i, v = range container { f(v) }
		case func(int, uintptr):				for i, v = range container { f(i, v) }
		case func(...uintptr):					f(container...)
		case func(interface{}):					for i, v = range container { f(v) }
		case func(int, interface{}):			for i, v = range container { f(i, v) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(R.ValueOf(v)) }
		case func(int, R.Value):				for i, v = range container { f(i, R.ValueOf(v)) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, R.ValueOf(v)) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), R.ValueOf(v)) }
		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v = range container { s[i] = R.ValueOf(v) }
												f(s...)
												i = 0
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(uintptr):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, uintptr):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...uintptr):					s := make([]uintptr, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepRValueSlice(container []R.Value, step int, f interface{}) (i int) {
	if step == 1 {
		var v	R.Value
		switch f := f.(type) {
		case func(interface{}):					for i, v = range container { f(v.Interface()) }
		case func(int, interface{}):			for i, v = range container { f(i, v.Interface()) }
		case func(interface{}, interface{}):	for i, v = range container { f(i, v.Interface()) }
		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v = range container { s[i] = v.Interface() }
												f(s...)
												i = 0
		case func(R.Value):						for i, v = range container { f(v) }
		case func(int, R.Value):				for i, v = range container { f(i, v) }
		case func(interface{}, R.Value):		for i, v = range container { f(i, v) }
		case func(R.Value, R.Value):			for i, v = range container { f(R.ValueOf(i), v) }
		case func(...R.Value):					f(container...)
		}
		i++
	} else {
		l := Len(container)
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}

		switch f := f.(type) {
		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset].Interface())
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset].Interface())
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset].Interface())
													i++
												}

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset].Interface())
												}
												f(s...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), container[offset])
													i++
												}

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												i = 1
		}
	}
	return
}

func stepIndexable(container Indexable, step int, f interface{}) (i int) {
	l := Len(container)
	steps := l / step
	offset := -1

	switch {
	case step == 0:							return
	case step > l:							return
	case step < 0:							offset = l
											steps = -steps
	}

	switch f := f.(type) {
	case func(interface{}):					for ; steps > 0; steps-- {
												offset = offset + step
												f(container.AtOffset(offset))
												i++
											}

	case func(int, interface{}):			for ; steps > 0; steps-- {
												offset = offset + step
												f(offset, container.AtOffset(offset))
												i++
											}

	case func(interface{}, interface{}):	for ; steps > 0; steps-- {
												offset = offset + step
												f(offset, container.AtOffset(offset))
												i++
											}

	case func(...interface{}):				s := make([]interface{}, steps, steps)
											for ; steps > 0; steps-- {
												offset = offset + step
												s[i] = container.AtOffset(offset)
												i++
											}
											f(s...)
											i = 1

	case func(R.Value):						for ; steps > 0; steps-- {
												offset = offset + step
												f(R.ValueOf(container.AtOffset(offset)))
												i++
											}

	case func(int, R.Value):				for ; steps > 0; steps-- {
												offset = offset + step
												f(offset, R.ValueOf(container.AtOffset(offset)))
												i++
											}

	case func(interface{}, R.Value):		for ; steps > 0; steps-- {
												offset = offset + step
												f(offset, R.ValueOf(container.AtOffset(offset)))
												i++
											}

	case func(R.Value, R.Value):			for ; steps > 0; steps-- {
												offset = offset + step
												f(R.ValueOf(offset), R.ValueOf(container.AtOffset(offset)))
												i++
											}

	case func(...R.Value):					s := make([]R.Value, steps, steps)
											for ; steps > 0; steps-- {
												offset = offset + step
												s[i] = R.ValueOf(container.AtOffset(offset))
												i++
											}
											f(s...)
											i = 1

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													// f(...V)
													p := make([]R.Value, steps, steps)
													for i := 0; steps > 0; steps-- {
														p[i] = R.ValueOf(container.AtOffset(i))
														i++
													}
													f.Call(p)
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for ; steps > 0; steps-- {
																			offset = offset + step
																			p[0] = R.ValueOf(container.AtOffset(offset))
																			f.Call(p)
																			i++
																		}

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		for ; steps > 0; steps-- {
																			offset = offset + step
																			p[0], p[1] = R.ValueOf(offset), R.ValueOf(container.AtOffset(offset))
																			f.Call(p)
																			i++
																		}
													}
												}
											}
	}
	return
}

func stepSlice(s R.Value, step int, f interface{}) (i int) {
	l := Len(s)
	if step == 1 {
		switch f := f.(type) {
		case func(interface{}):					for ; i < l; i++ {
													f(s.Index(i).Interface())
												}

		case func(int, interface{}):			for ; i < l; i++ {
													f(i, s.Index(i).Interface())
												}

		case func(interface{}, interface{}):	for ; i < l; i++ {
													f(i, s.Index(i).Interface())
												}

		case func(...interface{}):				p := make([]interface{}, l, l)
												for ; i < l; i++ {
													p[i] = s.Index(i).Interface()
												}
												f(p...)
												i = 1

		case func(R.Value):						for ; i < l; i++ {
													f(s.Index(i))
												}

		case func(int, R.Value):				for ; i < l; i++ {
													f(i, s.Index(i))
												}

		case func(interface{}, R.Value):		for ; i < l; i++ {
													f(i, s.Index(i))
												}

		case func(R.Value, R.Value):			for ; i < l; i++ {
													f(R.ValueOf(i), s.Index(i))
												}

		case func(...R.Value):					p := make([]R.Value, l, l)
												for ; i < l; i++ {
													p[i] = s.Index(i)
												}
												f(p...)
												i = 1

		default:								if f := R.ValueOf(f); f.Kind() == R.Func {
													if t := f.Type(); t.IsVariadic() {
														//	f(...v)
														p := make([]R.Value, l, l)
														for ; i < l; i++ {
															p[i] = s.Index(i)
														}
														f.Call(p)
														i = 1
													} else {
														switch t.NumIn() {
														case 1:				//	f(v)
																			p := make([]R.Value, 1, 1)
																			for ; i < l; i++ {
																				p[0] = s.Index(i)
																				f.Call(p)
																			}

														case 2:				//	f(i, v)
																			p := make([]R.Value, 2, 2)
																			for ; i < l; i++ {
																				p[0], p[1] = R.ValueOf(i), s.Index(i)
																				f.Call(p)
																			}
														}
													}
												}
		}
	} else {
		steps := l / step
		offset := -1

		switch {
		case step == 0:							return
		case step > l:							return
		case step < 0:							offset = l
												steps = -steps
		}
	
		switch f := f.(type) {
		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(s.Index(offset).Interface())
													i++
												}

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, s.Index(offset).Interface())
													i++
												}

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, s.Index(offset).Interface())
													i++
												}

		case func(...interface{}):				p := make([]interface{}, steps, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													p[i] = s.Index(offset).Interface()
													i++
												}
												f(p...)
												i = 1

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(s.Index(offset))
													i++
												}

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, s.Index(offset))
													i++
												}

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, s.Index(offset))
													i++
												}

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), s.Index(offset))
													i++
												}

		case func(...R.Value):					p := make([]R.Value, steps, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													p[i] = s.Index(offset)
													i++
												}
												f(p...)
												i = 1

		default:								if f := R.ValueOf(f); f.Kind() == R.Func {
													if t := f.Type(); t.IsVariadic() {
														//	f(...v)
														p := make([]R.Value, steps, steps)
														for ; steps > 0; steps-- {
															offset = offset + step
															p[i] = s.Index(offset)
															i++
														}
														f.Call(p)
														i = 1
													} else {
														switch t.NumIn() {
														case 1:				//	f(v)
																			p := make([]R.Value, 1, 1)
																			for ; steps > 0; steps-- {
																				offset = offset + step
																				p[0] = s.Index(offset)
																				f.Call(p)
																				i++
																			}

														case 2:				//	f(i, v)
																			p := make([]R.Value, 2, 2)
																			for ; steps > 0; steps-- {
																				offset = offset + step
																				p[0], p[1] = R.ValueOf(offset), s.Index(offset)
																				f.Call(p)
																				i++
																			}
														}
													}
												}
		}
	}
	return
}

func stepChannel(c R.Value, step int, f interface{}) (i int) {
	switch {
	case step < 0:		return
	case step == 1:		switch f := f.(type) {
						case func(interface{}):					for v, open := c.Recv(); open; i++ {
																	f(v.Interface())
																	v, open = c.Recv()
																}

						case func(int, interface{}):			for v, open := c.Recv(); open; i++ {
																	f(i, v.Interface())
																	v, open = c.Recv()
																}

						case func(interface{}, interface{}):	for v, open := c.Recv(); open; i++ {
																	f(i, v.Interface())
																	v, open = c.Recv()
																}

						case func(...interface{}):				p := make([]interface{}, 0, 4)
																for v, open := c.Recv(); open; {
																	p = append(p, v.Interface())
																	v, open = c.Recv()
																}
																f(p...)
																i = 1

						case func(R.Value):						for v, open := c.Recv(); open; i++ {
																	f(v)
																	v, open = c.Recv()
																}

						case func(int, R.Value):				for v, open := c.Recv(); open; i++ {
																	f(i, v)
																	v, open = c.Recv()
																}

						case func(interface{}, R.Value):		for v, open := c.Recv(); open; i++ {
																	f(i, v)
																	v, open = c.Recv()
																}

						case func(R.Value, R.Value):			for v, open := c.Recv(); open; i++ {
																	f(R.ValueOf(i), v)
																	v, open = c.Recv()
																}

						case func(...R.Value):					p := make([]R.Value, 0, 4)
																for v, open := c.Recv(); open; {
																	p = append(p, v)
																	v, open = c.Recv()
																}
																f(p...)
																i = 1

						default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																	if t := f.Type(); t.IsVariadic() {
																		//	f(...v)
																		p := make([]R.Value, 0, 4)
																		for v, open := c.Recv(); open; {
																			p = append(p, v)
																			v, open = c.Recv()
																		}
																		f.Call(p)
																		i = 1
																	} else {
																		switch t.NumIn() {
																		case 1:				//	f(v)
																							p := make([]R.Value, 1, 1)
																							for v, open := c.Recv(); open; i++ {
																								p[0] = v
																								f.Call(p)
																								v, open = c.Recv()
																							}

																		case 2:				//	f(i, v)
																							p := make([]R.Value, 2, 2)
																							for v, open := c.Recv(); open; i++ {
																								p[0], p[1] = R.ValueOf(i), v
																								f.Call(p)
																								v, open = c.Recv()
																							}
																		}
																	}
																}
						}
						return

	default:			offset := step
						n := 0
						switch f := f.(type) {
						case func(interface{}):					for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(v.Interface())
																		offset = step
																		i++
																	}
																	v, open = c.Recv()
																}

						case func(int, interface{}):			for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(n, v.Interface())
																		offset = step
																		i++
																	}
																	v, open = c.Recv()
																	n++
																}

						case func(interface{}, interface{}):	for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(n, v.Interface())
																		offset = step
																		i++
																	}
																	v, open = c.Recv()
																	n++
																}

						case func(...interface{}):				p := make([]interface{}, 0, 4)
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		p = append(p, v.Interface())
																		offset = step
																	}
																	v, open = c.Recv()
																}
																f(p...)
																i = 1

						case func(R.Value):						for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(v)
																		offset = step
																		i++
																	}
																	v, open = c.Recv()
																}

						case func(int, R.Value):				for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(n, v)
																		offset = step
																		i++
																	}
																	v, open = c.Recv()
																	n++
																}

						case func(interface{}, R.Value):		for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(n, v)
																		offset = step
																		i++
																	}
																	v, open = c.Recv()
																	n++
																}

						case func(R.Value, R.Value):			for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(R.ValueOf(n), v)
																		offset = step
																		i++
																	}
																	v, open = c.Recv()
																	n++
																}

						case func(...R.Value):					p := make([]R.Value, 0, 4)
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		p = append(p, v)
																		offset = step
																	}
																	v, open = c.Recv()
																}
																f(p...)
																i = 1

						default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																	if t := f.Type(); t.IsVariadic() {
																		//	f(...v)
																		p := make([]R.Value, 0, 4)
																		for v, open := c.Recv(); open; offset-- {
																			if offset == 0 {
																				p = append(p, v)
																				offset = step
																			}
																			v, open = c.Recv()
																		}
																		f.Call(p)
																		i = 1
																	} else {
																		switch t.NumIn() {
																		case 1:				//	f(v)
																							p := make([]R.Value, 1, 1)
																							for v, open := c.Recv(); open; offset-- {
																								if offset == 0 {
																									p[0] = v
																									f.Call(p)
																									offset = step
																									i++
																								}
																								v, open = c.Recv()
																							}

																		case 2:				//	f(i, v)
																							p := make([]R.Value, 2, 2)
																							for v, open := c.Recv(); open; offset-- {
																								if offset == 0 {
																									p[0], p[1] = R.ValueOf(n), v
																									f.Call(p)
																									offset = step
																									i++
																								}
																								v, open = c.Recv()
																								n++
																							}
																		}
																	}
																}
						}
		}
	return
}

func stepFunction(g R.Value, step int, f interface{}) (i int) {
	if step > 0 {
		offset := step
		if tg := g.Type(); tg.NumOut() == 2 {
			p := make([]R.Value, 0, 0)
			switch tg.NumIn() {
			case 0:			switch f := f.(type) {
							case func(interface{}):					for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(v[0].Interface())
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(int, interface{}):			for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(n, v[0].Interface())
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(interface{}, interface{}):	for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(n, v[0].Interface())
																			offset = step
																			i++
																		}
																		n++
																	}


							case func(...interface{}):				pf := make([]interface{}, 0, 4)
																	for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			pf = append(pf, v[0].Interface())
																			offset = step
																		}
																		n++
																	}
																	f(pf...)
																	i = 1

							case func(R.Value):						for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(v[0])
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(int, R.Value):				for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(n, v[0])
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(interface{}, R.Value):		for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(n, v[0])
																			offset = step
																			i++
																		}
																		n++
																	}


							case func(R.Value, R.Value):			for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(R.ValueOf(n), v[0])
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(...R.Value):					pf := make([]R.Value, 0, 4)
																	for n := 0; ; n++ {
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			pf = append(pf, v[0])
																			offset = step
																		}
																		n++
																	}
																	f(pf...)
																	i = 1

							default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																		if tf := f.Type(); tf.IsVariadic() {
																			//	f(...v)
																			pf := make([]R.Value, 0, 4)
																			for n := 0; ; n++ {
																				offset--
																				if offset == 0 {
																					v := g.Call(p)
																					if v[1].Bool() {
																						break
																					}
																					pf = append(pf, v[0])
																					offset = step
																				}
																				n++
																			}
																			f.Call(pf)
																			i = 1

																		} else {
																			switch tf.NumIn() {
																			case 1:		//	f(v)
																						pf := make([]R.Value, 1, 1)
																						for n := 0; ; n++ {
																							offset--
																							if offset == 0 {
																								v := g.Call(p)
																								if v[1].Bool() {
																									break
																								}
																								pf[0] = v[0]
																								f.Call(pf)
																								offset = step
																								i++
																							}
																							n++
																						}

																			case 2:		//	f(i, v)
																						pf := make([]R.Value, 2, 2)
																						for n := 0; ; n++ {
																							offset--
																							if offset == 0 {
																								v := g.Call(p)
																								if v[1].Bool() {
																									break
																								}
																								pf[0], pf[1] = R.ValueOf(n), v[0]
																								f.Call(pf)
																								offset = step
																								i++
																							}
																							n++
																						}
																			}
																		}
																	}
							}

			case 1:			p := make([]R.Value, 1, 1)
							switch f := f.(type) {
							case func(interface{}):					for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(v[0].Interface())
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(int, interface{}):			for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(n, v[0].Interface())
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(interface{}, interface{}):	for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(n, v[0].Interface())
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(...interface{}):				pf := make([]interface{}, 0, 4)
																	for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			pf = append(pf, v[0].Interface())
																			offset = step
																		}
																		n++
																	}
																	f(pf...)
																	i = 1

							case func(R.Value):						for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(v[0])
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(int, R.Value):				for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(n, v[0])
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(interface{}, R.Value):		for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(n, v[0])
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(R.Value, R.Value):			for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			f(R.ValueOf(n), v[0])
																			offset = step
																			i++
																		}
																		n++
																	}

							case func(...R.Value):					pf := make([]R.Value, 0, 4)
																	for n := 0; ; n++ {
																		p[0] = R.ValueOf(n)
																		offset--
																		if offset == 0 {
																			v := g.Call(p)
																			if v[1].Bool() {
																				break
																			}
																			pf = append(pf, v[0])
																			offset = step
																		}
																		n++
																	}
																	f(pf...)
																	i = 1

							default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																		if tf := f.Type(); tf.IsVariadic() {
																			//	f(...v)
																			pf := make([]R.Value, 0, 4)
																			for n := 0; ; n++ {
																				p[0] = R.ValueOf(n)
																				offset--
																				if offset == 0 {
																					v := g.Call(p)
																					if v[1].Bool() {
																						break
																					}
																					pf = append(pf, v[0])
																					offset = step
																				}
																				n++
																			}
																			f.Call(pf)
																			i = 1

																		} else {
																			switch tf.NumIn() {
																			case 1:		//	f(v)
																						pf := make([]R.Value, 1, 1)
																						for n := 0; ; n++ {
																							p[0] = R.ValueOf(n)
																							offset--
																							if offset == 0 {
																								v := g.Call(p)
																								if v[1].Bool() {
																									break
																								}
																								pf[0] = v[0]
																								f.Call(pf)
																								offset = step
																								i++
																							}
																							n++
																						}

																			case 2:		//	f(i, v)
																						pf := make([]R.Value, 2, 2)
																						for n := 0; ; n++ {
																							p[0] = R.ValueOf(n)
																							offset--
																							if offset == 0 {
																								v := g.Call(p)
																								if v[1].Bool() {
																									break
																								}
																								pf[0], pf[1] = p[0], v[0]
																								f.Call(pf)
																								offset = step
																								i++
																							}
																							n++
																						}


																			}
																		}
																	}
							}
			}
		}
	}
	return
}