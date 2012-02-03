package sequences

import R "reflect"

func stepBool(container []bool, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(bool):						for _, v := range container { f(v) }
												ok = true

		case func(int, bool):					for i, v := range container { f(i, v) }
												ok = true

		case func(...bool):						f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, bool):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...bool):						s := make([]bool, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepComplex64(container []complex64, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(complex64):					for _, v := range container { f(v) }
												ok = true

		case func(int, complex64):				for i, v := range container { f(i, v) }
												ok = true

		case func(...complex64):				f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, complex64):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...complex64):				s := make([]complex64, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepComplex128(container []complex128, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(complex128):					for _, v := range container { f(v) }
												ok = true

		case func(int, complex128):				for i, v := range container { f(i, v) }
												ok = true

		case func(...complex128):				f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, complex128):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...complex128):				s := make([]complex128, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepError(container []error, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(error):						for _, v := range container { f(v) }
												ok = true

		case func(int, error):					for i, v := range container { f(i, v) }
												ok = true

		case func(...error):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, error):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...error):					s := make([]error, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepFloat32(container []float32, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(float32):						for _, v := range container { f(v) }
												ok = true

		case func(int, float32):				for i, v := range container { f(i, v) }
												ok = true

		case func(...float32):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, float32):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...float32):					s := make([]float32, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepFloat64(container []float64, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(float64):						for _, v := range container { f(v) }
												ok = true

		case func(int, float64):				for i, v := range container { f(i, v) }
												ok = true

		case func(...float64):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, float64):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...float64):					s := make([]float64, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepInt(container []int, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(int):							for _, v := range container { f(v) }
												ok = true

		case func(int, int):					for i, v := range container { f(i, v) }
												ok = true

		case func(...int):						f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, int):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...int):						s := make([]int, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepInt8(container []int8, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(int8):						for _, v := range container { f(v) }
												ok = true

		case func(int, int8):					for i, v := range container { f(i, v) }
												ok = true

		case func(...int8):						f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, int8):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...int8):						s := make([]int8, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepInt16(container []int16, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(int16):						for _, v := range container { f(v) }
												ok = true

		case func(int, int16):					for i, v := range container { f(i, v) }
												ok = true

		case func(...int16):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, int16):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...int16):					s := make([]int16, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepInt32(container []int32, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(int32):						for _, v := range container { f(v) }
												ok = true

		case func(int, int32):					for i, v := range container { f(i, v) }
												ok = true

		case func(...int32):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, int32):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...int32):					s := make([]int32, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepInt64(container []int64, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(int64):						for _, v := range container { f(v) }
												ok = true

		case func(int, int64):					for i, v := range container { f(i, v) }
												ok = true

		case func(...int64):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, int64):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...int64):					s := make([]int64, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepInterface(container []interface{}, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				f(container...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepString(container []string, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(string):						for _, v := range container { f(v) }
												ok = true

		case func(int, string):					for i, v := range container { f(i, v) }
												ok = true

		case func(...string):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, string):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...string):					s := make([]string, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepUint(container []uint, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(uint):						for _, v := range container { f(v) }
												ok = true

		case func(int, uint):					for i, v := range container { f(i, v) }
												ok = true

		case func(...uint):						f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, uint):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...uint):						s := make([]uint, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepUint8(container []uint8, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(uint8):						for _, v := range container { f(v) }
												ok = true

		case func(int, uint8):					for i, v := range container { f(i, v) }
												ok = true

		case func(...uint8):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, uint8):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...uint8):					s := make([]uint8, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepUint16(container []uint16, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(uint16):						for _, v := range container { f(v) }
												ok = true

		case func(int, uint16):					for i, v := range container { f(i, v) }
												ok = true

		case func(...uint16):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, uint16):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...uint16):					s := make([]uint16, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepUint32(container []uint32, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(uint32):						for _, v := range container { f(v) }
												ok = true

		case func(int, uint32):					for i, v := range container { f(i, v) }
												ok = true

		case func(...uint32):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, uint32):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...uint32):					s := make([]uint32, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepUint64(container []uint64, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(uint64):						for _, v := range container { f(v) }
												ok = true

		case func(int, uint64):					for i, v := range container { f(i, v) }
												ok = true

		case func(...uint64):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, uint64):					for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...uint64):					s := make([]uint64, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepUintptr(container []uintptr, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(uintptr):						for _, v := range container { f(v) }
												ok = true

		case func(int, uintptr):				for i, v := range container { f(i, v) }
												ok = true

		case func(...uintptr):					f(container...)
												ok = true

		case func(interface{}):					for _, v := range container { f(v) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(R.ValueOf(v)) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, R.ValueOf(v)) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), R.ValueOf(v)) }
												ok = true

		case func(...R.Value):					s := make([]R.Value, len(container), len(container))
												for i, v := range container { s[i] = R.ValueOf(v) }
												f(s...)
												ok = true
		}
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
												}
												ok = true

		case func(int, uintptr):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...uintptr):					s := make([]uintptr, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(interface{}):					for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(container[offset]))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, R.ValueOf(container[offset]))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), R.ValueOf(container[offset]))
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, R.ValueOf(container[offset]))
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepRValueSlice(container []R.Value, step int, f interface{}) (ok bool) {
	if step == 1 {
		switch f := f.(type) {
		case func(interface{}):					for _, v := range container { f(v.Interface()) }
												ok = true

		case func(int, interface{}):			for i, v := range container { f(i, v.Interface()) }
												ok = true

		case func(interface{}, interface{}):	for i, v := range container { f(i, v.Interface()) }
												ok = true

		case func(...interface{}):				s := make([]interface{}, len(container), len(container))
												for i, v := range container { s[i] = v.Interface() }
												f(s...)
												ok = true

		case func(R.Value):						for _, v := range container { f(v) }
												ok = true

		case func(int, R.Value):				for i, v := range container { f(i, v) }
												ok = true

		case func(interface{}, R.Value):		for i, v := range container { f(i, v) }
												ok = true

		case func(R.Value, R.Value):			for i, v := range container { f(R.ValueOf(i), v) }
												ok = true

		case func(...R.Value):					f(container...)
												ok = true
		}
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
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset].Interface())
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset].Interface())
												}
												ok = true

		case func(...interface{}):				s := make([]interface{}, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset].Interface())
												}
												f(s...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(container[offset])
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, container[offset])
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), container[offset])
												}
												ok = true

		case func(...R.Value):					s := make([]R.Value, 0, steps)
												for ; steps > 0; steps-- {
													offset = offset + step
													s = append(s, container[offset])
												}
												f(s...)
												ok = true
		}
	}
	return
}

func stepIndexable(container Indexable, step int, f interface{}) (ok bool) {
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
											}
											ok = true

	case func(int, interface{}):			for ; steps > 0; steps-- {
												offset = offset + step
												f(offset, container.AtOffset(offset))
											}
											ok = true

	case func(interface{}, interface{}):	for ; steps > 0; steps-- {
												offset = offset + step
												f(offset, container.AtOffset(offset))
											}
											ok = true

	case func(...interface{}):				s := make([]interface{}, steps, steps)
											for i := 0; steps > 0; steps-- {
												offset = offset + step
												s[i] = container.AtOffset(offset)
												i++
											}
											f(s...)
											ok = true

	case func(R.Value):						for ; steps > 0; steps-- {
												offset = offset + step
												f(R.ValueOf(container.AtOffset(offset)))
											}
											ok = true

	case func(int, R.Value):				for ; steps > 0; steps-- {
												offset = offset + step
												f(offset, R.ValueOf(container.AtOffset(offset)))
											}
											ok = true

	case func(interface{}, R.Value):		for ; steps > 0; steps-- {
												offset = offset + step
												f(offset, R.ValueOf(container.AtOffset(offset)))
											}
											ok = true

	case func(R.Value, R.Value):			for ; steps > 0; steps-- {
												offset = offset + step
												f(R.ValueOf(offset), R.ValueOf(container.AtOffset(offset)))
											}
											ok = true

	case func(...R.Value):					s := make([]R.Value, steps, steps)
											for i := 0; steps > 0; steps-- {
												offset = offset + step
												s[i] = R.ValueOf(container.AtOffset(offset))
												i++
											}
											f(s...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													// f(...V)
													p := make([]R.Value, steps, steps)
													for i := 0; steps > 0; steps-- {
														p[i] = R.ValueOf(container.AtOffset(i))
														i++
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for ; steps > 0; steps-- {
																			offset = offset + step
																			p[0] = R.ValueOf(container.AtOffset(offset))
																			f.Call(p)
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		for ; steps > 0; steps-- {
																			offset = offset + step
																			p[0], p[1] = R.ValueOf(offset), R.ValueOf(container.AtOffset(offset))
																			f.Call(p)
																		}
																		ok = true
													}
												}
											}
	}
	return
}

func stepSlice(s R.Value, step int, f interface{}) (ok bool) {
	l := Len(s)
	if step == 1 {
		switch f := f.(type) {
		case func(interface{}):					for i := 0; i < l; i++ {
													f(s.Index(i).Interface())
												}
												ok = true

		case func(int, interface{}):			for i := 0; i < l; i++ {
													f(i, s.Index(i).Interface())
												}
												ok = true

		case func(interface{}, interface{}):	for i := 0; i < l; i++ {
													f(i, s.Index(i).Interface())
												}
												ok = true

		case func(...interface{}):				p := make([]interface{}, l, l)
												for i := 0; i < l; i++ {
													p[i] = s.Index(i).Interface()
												}
												f(p...)
												ok = true

		case func(R.Value):						for i := 0; i < l; i++ {
													f(s.Index(i))
												}
												ok = true

		case func(int, R.Value):				for i := 0; i < l; i++ {
													f(i, s.Index(i))
												}
												ok = true

		case func(interface{}, R.Value):		for i := 0; i < l; i++ {
													f(i, s.Index(i))
												}
												ok = true

		case func(R.Value, R.Value):			for i := 0; i < l; i++ {
													f(R.ValueOf(i), s.Index(i))
												}
												ok = true

		case func(...R.Value):					p := make([]R.Value, l, l)
												for i := 0; i < l; i++ {
													p[i] = s.Index(i)
												}
												f(p...)
												ok = true

		default:								if f := R.ValueOf(f); f.Kind() == R.Func {
													if t := f.Type(); t.IsVariadic() {
														//	f(...v)
														p := make([]R.Value, l, l)
														for i := 0; i < l; i++ {
															p[i] = s.Index(i)
														}
														f.Call(p)
														ok = true
													} else {
														switch t.NumIn() {
														case 1:				//	f(v)
																			p := make([]R.Value, 1, 1)
																			for i := 0; i < l; i++ {
																				p[0] = s.Index(i)
																				f.Call(p)
																			}
																			ok = true

														case 2:				//	f(i, v)
																			p := make([]R.Value, 2, 2)
																			for i := 0; i < l; i++ {
																				p[0], p[1] = R.ValueOf(i), s.Index(i)
																				f.Call(p)
																			}
																			ok = true
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
												}
												ok = true

		case func(int, interface{}):			for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, s.Index(offset).Interface())
												}
												ok = true

		case func(interface{}, interface{}):	for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, s.Index(offset).Interface())
												}
												ok = true

		case func(...interface{}):				p := make([]interface{}, steps, steps)
												i := 0
												for ; steps > 0; steps-- {
													offset = offset + step
													p[i] = s.Index(offset).Interface()
												}
												f(p...)
												ok = true

		case func(R.Value):						for ; steps > 0; steps-- {
													offset = offset + step
													f(s.Index(offset))
												}
												ok = true

		case func(int, R.Value):				for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, s.Index(offset))
												}
												ok = true

		case func(interface{}, R.Value):		for ; steps > 0; steps-- {
													offset = offset + step
													f(offset, s.Index(offset))
												}
												ok = true

		case func(R.Value, R.Value):			for ; steps > 0; steps-- {
													offset = offset + step
													f(R.ValueOf(offset), s.Index(offset))
												}
												ok = true

		case func(...R.Value):					p := make([]R.Value, steps, steps)
												i := 0
												for ; steps > 0; steps-- {
													offset = offset + step
													p[i] = s.Index(offset)
												}
												f(p...)
												ok = true

		default:								if f := R.ValueOf(f); f.Kind() == R.Func {
													if t := f.Type(); t.IsVariadic() {
														//	f(...v)
														p := make([]R.Value, steps, steps)
														for i := 0; steps > 0; steps-- {
															offset = offset + step
															p[i] = s.Index(offset)
															i++
														}
														f.Call(p)
														ok = true
													} else {
														switch t.NumIn() {
														case 1:				//	f(v)
																			p := make([]R.Value, 1, 1)
																			for ; steps > 0; steps-- {
																				offset = offset + step
																				p[0] = s.Index(offset)
																				f.Call(p)
																			}
																			ok = true

														case 2:				//	f(i, v)
																			p := make([]R.Value, 2, 2)
																			for ; steps > 0; steps-- {
																				offset = offset + step
																				p[0], p[1] = R.ValueOf(offset), s.Index(offset)
																				f.Call(p)
																			}
																			ok = true
														}
													}
												}
		}
	}
	return
}

func stepChannel(c R.Value, step int, f interface{}) (ok bool) {
	switch {
	case step < 0:		return
	case step == 1:		switch f := f.(type) {
						case func(interface{}):					for v, open := c.Recv(); open; {
																	f(v.Interface())
																	v, open = c.Recv()
																}
																ok = true

						case func(int, interface{}):			i := 0
																for v, open := c.Recv(); open; i++ {
																	f(i, v.Interface())
																	v, open = c.Recv()
																}
																ok = true

						case func(interface{}, interface{}):	i := 0
																for v, open := c.Recv(); open; i++ {
																	f(i, v.Interface())
																	v, open = c.Recv()
																}
																ok = true

						case func(...interface{}):				p := make([]interface{}, 0, 4)
																for v, open := c.Recv(); open; {
																	p = append(p, v.Interface())
																	v, open = c.Recv()
																}
																f(p...)
																ok = true

						case func(R.Value):						for v, open := c.Recv(); open; {
																	f(v)
																	v, open = c.Recv()
																}
																ok = true

						case func(int, R.Value):				i := 0
																for v, open := c.Recv(); open; i++ {
																	f(i, v)
																	v, open = c.Recv()
																}
																ok = true

						case func(interface{}, R.Value):		i := 0
																for v, open := c.Recv(); open; i++ {
																	f(i, v)
																	v, open = c.Recv()
																}
																ok = true

						case func(R.Value, R.Value):			i := 0
																for v, open := c.Recv(); open; i++ {
																	f(R.ValueOf(i), v)
																	v, open = c.Recv()
																}
																ok = true

						case func(...R.Value):					p := make([]R.Value, 0, 4)
																for v, open := c.Recv(); open; {
																	p = append(p, v)
																	v, open = c.Recv()
																}
																f(p...)
																ok = true

						default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																	if t := f.Type(); t.IsVariadic() {
																		//	f(...v)
																		p := make([]R.Value, 0, 4)
																		for v, open := c.Recv(); open; {
																			p = append(p, v)
																			v, open = c.Recv()
																		}
																		f.Call(p)
																		ok = true
																	} else {
																		switch t.NumIn() {
																		case 1:				//	f(v)
																							p := make([]R.Value, 1, 1)
																							for v, open := c.Recv(); open; {
																								p[0] = v
																								f.Call(p)
																								v, open = c.Recv()
																							}
																							ok = true

																		case 2:				//	f(i, v)
																							p := make([]R.Value, 2, 2)
																							i := 0
																							for v, open := c.Recv(); open; i++ {
																								p[0], p[1] = R.ValueOf(i), v
																								f.Call(p)
																								v, open = c.Recv()
																							}
																							ok = true
																		}
																	}
																}
						}
						return

	default:			offset := step
						switch f := f.(type) {
						case func(interface{}):					for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(v.Interface())
																		offset = step
																	}
																	v, open = c.Recv()
																}
																ok = true

						case func(int, interface{}):			i := 0
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(i, v.Interface())
																		offset = step
																	}
																	v, open = c.Recv()
																	i++
																}
																ok = true

						case func(interface{}, interface{}):	i := 0
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(i, v.Interface())
																		offset = step
																	}
																	v, open = c.Recv()
																	i++
																}
																ok = true

						case func(...interface{}):				p := make([]interface{}, 0, 4)
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		p = append(p, v.Interface())
																		offset = step
																	}
																	v, open = c.Recv()
																}
																f(p...)
																ok = true

						case func(R.Value):						for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(v)
																		offset = step
																	}
																	v, open = c.Recv()
																}
																ok = true

						case func(int, R.Value):				i := 0
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(i, v)
																		offset = step
																	}
																	v, open = c.Recv()
																	i++
																}
																ok = true

						case func(interface{}, R.Value):		i := 0
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(i, v)
																		offset = step
																	}
																	v, open = c.Recv()
																	i++
																}
																ok = true

						case func(R.Value, R.Value):			i := 0
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		f(R.ValueOf(i), v)
																		offset = step
																	}
																	v, open = c.Recv()
																	i++
																}
																ok = true

						case func(...R.Value):					p := make([]R.Value, 0, 4)
																for v, open := c.Recv(); open; offset-- {
																	if offset == 0 {
																		p = append(p, v)
																		offset = step
																	}
																	v, open = c.Recv()
																}
																f(p...)
																ok = true

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
																		ok = true
																	} else {
																		switch t.NumIn() {
																		case 1:				//	f(v)
																							p := make([]R.Value, 1, 1)
																							for v, open := c.Recv(); open; offset-- {
																								if offset == 0 {
																									p[0] = v
																									f.Call(p)
																									offset = step
																								}
																								v, open = c.Recv()
																							}
																							ok = true

																		case 2:				//	f(i, v)
																							p := make([]R.Value, 2, 2)
																							i := 0
																							for v, open := c.Recv(); open; offset-- {
																								if offset == 0 {
																									p[0], p[1] = R.ValueOf(i), v
																									f.Call(p)
																									offset = step
																								}
																								v, open = c.Recv()
																								i++
																							}
																							ok = true
																		}
																	}
																}
						}
		}
	return
}

func stepFunction(g R.Value, step int, f interface{}) (ok bool) {
	if step > 0 {
		if tg := g.Type(); tg.NumOut() == 2 {
			offset := step
			switch tg.NumIn() {
			case 0:			p := []R.Value{}
							switch f := f.(type) {
							case func(interface{}):					for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(v[0].Interface())
																			offset = step
																		}
																	}
																	ok = true

							case func(int, interface{}):			for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(i, v[0].Interface())
																			offset = step
																		}
																		i++
																	} 
																	ok = true

							case func(interface{}, interface{}):	for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(i, v[0].Interface())
																			offset = step
																		}
																		i++
																	}
																	ok = true

							case func(...interface{}):				pf := make([]interface{}, 0, 4)
																	for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			pf = append(pf, v[0].Interface())
																			offset = step
																		}
																	}
																	f(pf...)
																	ok = true

							case func(R.Value):						for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(v[0])
																			offset = step
																		}
																	}
																	ok = true

							case func(int, R.Value):				for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(i, v[0])
																			offset = step
																		}
																		i++
																	}
																	ok = true

							case func(interface{}, R.Value):		for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(i, v[0])
																			offset = step
																		}
																		i++
																	}
																	ok = true

							case func(R.Value, R.Value):			for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(R.ValueOf(i), v[0])
																			offset = step
																		}
																		i++
																	}
																	ok = true

							case func(...R.Value):					pf := make([]R.Value, 0, 4)
																	for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			pf = append(pf, v[0])
																			offset = step
																		}
																	}
																	f(pf...)
																	ok = true

							default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																		if tf := f.Type(); tf.IsVariadic() {
																			//	f(...v)
																			i := 0
																			pf := make([]R.Value, 0, 4)
																			for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																				offset--
																				if offset == 0 {
																					pf = append(pf, v[0])
																					offset = step
																				}
																				i++
																			}
																			f.Call(pf)
																			ok = true
																		} else {
																			switch tf.NumIn() {
																			case 1:		//	f(v)
																						pf := make([]R.Value, 1, 1)
																						for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																							offset--
																							if offset == 0 {
																								pf[0] = v[0]
																								f.Call(pf)
																								offset = step
																							}
																						}
																						ok = true

																			case 2:		//	f(i, v)
																						i := 0
																						pf := make([]R.Value, 2, 2)
																						for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																							offset--
																							if offset == 0 {
																								pf[0], pf[1] = R.ValueOf(i), v[0]
																								f.Call(pf)
																								offset = step
																							}
																							i++
																						}
																						ok = true
																			}
																		}
																	}
							}

			case 1:			var i	int
							p := []R.Value{ R.ValueOf(0) }
							switch f := f.(type) {
							case func(interface{}):					for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(v[0].Interface())
																			i++
																			p[0] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	ok = true

							case func(int, interface{}):			for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(i, v[0].Interface())
																			i++
																			p[0] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	ok = true

							case func(interface{}, interface{}):	for v := g.Call(p); !v[0].IsValid(); v = g.Call(p) {
																		offset--
																		if offset == step {
																			f(i, v[0].Interface())
																			i++
																			p[0] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	ok = true

							case func(...interface{}):				pf := make([]interface{}, 0, 4)
																	for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == step {
																			pf = append(pf, v[0].Interface())
																			i++
																			p[0] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	f(pf...)
																	ok = true

							case func(R.Value):						for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(v[0])
																			i++
																			p[0] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	ok = true

							case func(int, R.Value):				for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(i, v[0])
																			i++
																			p[0] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	ok = true

							case func(interface{}, R.Value):		for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(i, v[0])
																			i++
																			p[0] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	ok = true

							case func(R.Value, R.Value):			for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			f(p[0], v[0])
																			i++
																			p[0] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	ok = true

							case func(...R.Value):					pf := make([]R.Value, 0, 4)
																	for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																		offset--
																		if offset == 0 {
																			pf = append(pf, v[0])
																			i++
																			pf[i] = R.ValueOf(i)
																			offset = step
																		}
																	}
																	f(pf...)
																	ok = true

							default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																		if tf := f.Type(); tf.IsVariadic() {
																			//	f(...v)
																			pf := make([]R.Value, 0, 4)
																			for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																				offset--
																				if offset == 0 {
																					pf = append(pf, v[0])
																					i++
																					p[0] = R.ValueOf(i)
																					offset = step
																				}
																			}
																			f.Call(pf)
																			ok = true
																		} else {
																			switch tf.NumIn() {
																			case 1:		//	f(v)
																						for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																							offset--
																							if offset == 0 {
																								p[0] = v[0]
																								f.Call(p)
																								i++
																								p[0] = R.ValueOf(i)
																								offset = step
																							}
																						}
																						ok = true

																			case 2:		//	f(i, v)
																						pf := make([]R.Value, 2, 2)
																						for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																							offset--
																							if offset == 0 {
																								pf[0], pf[1] = p[0], v[0]
																								f.Call(pf)
																								i++
																								p[0] = R.ValueOf(i)
																								offset = step
																							}
																						}
																						ok = true
																			}
																		}
																	}
							}
			}
		}
	}
	return
}