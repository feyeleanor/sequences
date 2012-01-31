package sequences

import R "reflect"

func eachBool(container []bool, f interface{}) (ok bool) {
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
	return
}

func eachComplex64(container []complex64, f interface{}) (ok bool) {
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
	return
}

func eachComplex128(container []complex128, f interface{}) (ok bool) {
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
	return
}

func eachError(container []error, f interface{}) (ok bool) {
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
	return
}

func eachFloat32(container []float32, f interface{}) (ok bool) {
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
	return
}

func eachFloat64(container []float64, f interface{}) (ok bool) {
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
	return
}

func eachInt(container []int, f interface{}) (ok bool) {
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
	return
}

func eachInt8(container []int8, f interface{}) (ok bool) {
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
	return
}

func eachInt16(container []int16, f interface{}) (ok bool) {
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
	return
}

func eachInt32(container []int32, f interface{}) (ok bool) {
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
	return
}

func eachInt64(container []int64, f interface{}) (ok bool) {
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
	return
}

func eachInterface(container []interface{}, f interface{}) (ok bool) {
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
	return
}

func eachString(container []string, f interface{}) (ok bool) {
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
	return
}

func eachUint(container []uint, f interface{}) (ok bool) {
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
	return
}

func eachUint8(container []uint8, f interface{}) (ok bool) {
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
	return
}

func eachUint16(container []uint16, f interface{}) (ok bool) {
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
	return
}

func eachUint32(container []uint32, f interface{}) (ok bool) {
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
	return
}

func eachUint64(container []uint64, f interface{}) (ok bool) {
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
	return
}

func eachUintptr(container []uintptr, f interface{}) (ok bool) {
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
	return
}

func eachRValueSlice(container []R.Value, f interface{}) (ok bool) {
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
	return
}

func eachIndexable(container Indexable, f interface{}) (ok bool) {
	end := Len(container)
	switch f := f.(type) {
	case func(interface{}):					for i := 0; i < end; i++ {
												f(container.AtOffset(i))
											}
											ok = true

	case func(int, interface{}):			for i := 0; i < end; i++ {
												f(i, container.AtOffset(i))
											}
											ok = true

	case func(interface{}, interface{}):	for i := 0; i < end; i++ {
												f(i, container.AtOffset(i))
											}
											ok = true

	case func(...interface{}):				p := make([]interface{}, end, end)
											for i := 0; i < end; i++ {
												p[i] = container.AtOffset(i)
											}
											f(p...)
											ok = true

	case func(R.Value):						for i := 0; i < end; i++ {
												f(R.ValueOf(container.AtOffset(i)))
											}
											ok = true

	case func(int, R.Value):				for i := 0; i < end; i++ {
												f(i, R.ValueOf(container.AtOffset(i)))
											}
											ok = true

	case func(interface{}, R.Value):		for i := 0; i < end; i++ {
												f(i, R.ValueOf(container.AtOffset(i)))
											}
											ok = true

	case func(R.Value, R.Value):			for i := 0; i < end; i++ {
												f(R.ValueOf(i), R.ValueOf(container.AtOffset(i)))
											}
											ok = true


	case func(...R.Value):					p := make([]R.Value, end, end)
											for i := 0; i < end; i++ {
												p[i] = R.ValueOf(container.AtOffset(i))
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, end, end)
													for i := 0; i < end; i++ {
														p[i] = R.ValueOf(container.AtOffset(i))
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for i := 0; i < end; i++ {
																			p[0] = R.ValueOf(container.AtOffset(i))
																			f.Call(p)
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		for i := 0; i < end; i++ {
																			p[0], p[1] = R.ValueOf(i), R.ValueOf(container.AtOffset(i))
																			f.Call(p)
																		}
																		ok = true
													}
												}
											}
	}
	return
}

func eachMappable(container Mappable, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					Each(container.Keys(), func(v interface{}) {
												f(container.StoredAs(v))
												
											})
											ok = true

	case func(int, interface{}):			Each(container.Keys(), func(v interface{}) {
												f(v.(int), container.StoredAs(v))
											})
											ok = true

	case func(string, interface{}):			Each(container.Keys(), func(v interface{}) {
												f(v.(string), container.StoredAs(v))
											})
											ok = true

	case func(interface{}, interface{}):	Each(container.Keys(), func(v interface{}) {
												f(v, container.StoredAs(v))
											})
											ok = true

	case func(...interface{}):				l := Len(container)
											p := make([]interface{}, l, l)
											Each(container.Keys(), func(i int, v interface{}) {
												p[i] = container.StoredAs(v)
											})
											f(p...)
											ok = true

	case func(R.Value):						Each(container.Keys(), func(v interface{}) {
												f(R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(int, R.Value):				Each(container.Keys(), func(v interface{}) {
												f(v.(int), R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(string, R.Value):				Each(container.Keys(), func(v interface{}) {
												f(v.(string), R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(interface{}, R.Value):		Each(container.Keys(), func(v interface{}) {
												f(v, R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(R.Value, R.Value):			Each(container.Keys(), func(v interface{}) {
												f(R.ValueOf(v), R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(...R.Value):					l := Len(container)
											p := make([]R.Value, l, l)
											Each(container.Keys(), func(i int, v interface{}) {
												p[i] = R.ValueOf(container.StoredAs(v))
											})
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													// f(...v)
													l := Len(container)
													p := make([]R.Value, l, l)
													Each(container.Keys(), func(i int, v interface{}) {
														p[i] = R.ValueOf(container.StoredAs(v))
													})
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		Each(container.Keys(), func(v interface{}) {
																			p[0] = R.ValueOf(container.StoredAs(v))
																			f.Call(p)																			
																		})
																		ok = true

													case 2:				// f(i, v)
																		p := make([]R.Value, 2, 2)
																		Each(container.Keys(), func(v interface{}) {
																			p[0], p[1] = R.ValueOf(v), R.ValueOf(container.StoredAs(v))
																			f.Call(p)
																		})
																		ok = true
													}
												}
											}
	}
	return
}

func eachSlice(s R.Value, f interface{}) (ok bool) {
	end := Len(s)
	switch f := f.(type) {
	case func(interface{}):					for i := 0; i < end; i++ {
												f(s.Index(i).Interface())
											}
											ok = true

	case func(int, interface{}):			for i := 0; i < end; i++ {
												f(i, s.Index(i).Interface())
											}
											ok = true

	case func(interface{}, interface{}):	for i := 0; i < end; i++ {
												f(i, s.Index(i).Interface())
											}
											ok = true

	case func(...interface{}):				p := make([]interface{}, end, end)
											for i := 0; i < end; i++ {
												p[i] = s.Index(i).Interface()
											}
											f(p...)
											ok = true

	case func(R.Value):						for i := 0; i < end; i++ {
												f(s.Index(i))
											}
											ok = true

	case func(int, R.Value):				for i := 0; i < end; i++ {
												f(i, s.Index(i))
											}
											ok = true

	case func(interface{}, R.Value):		for i := 0; i < end; i++ {
												f(i, s.Index(i))
											}
											ok = true

	case func(R.Value, R.Value):			for i := 0; i < end; i++ {
												f(R.ValueOf(i), s.Index(i))
											}
											ok = true

	case func(...R.Value):					p := make([]R.Value, end, end)
											for i := 0; i < end; i++ {
												p[i] = s.Index(i)
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												end := Len(s)
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, end, end)
													for i := 0; i < end; i++ {
														p[i] = s.Index(i)
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for i := 0; i < end; i++ {
																			p[0] = s.Index(i)
																			f.Call(p)
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		for i := 0; i < end; i++ {
																			p[0], p[1] = R.ValueOf(i), s.Index(i)
																			f.Call(p)
																		}
																		ok = true
													}
												}
											}
	}
	return
}

func eachMap(m R.Value, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					for _, key := range m.MapKeys() {
												f(m.MapIndex(key).Interface())
											}
											ok = true

	case func(int, interface{}):			for _, key := range m.MapKeys() {
												f(int(key.Int()), m.MapIndex(key).Interface())
											}
											ok = true

	case func(string, interface{}):			for _, key := range m.MapKeys() {
												f(key.String(), m.MapIndex(key).Interface())
											}
											ok = true

	case func(interface{}, interface{}):	for _, key := range m.MapKeys() {
												f(key.Interface(), m.MapIndex(key).Interface())
											}
											ok = true

	case func(...interface{}):				l := Len(m)
											p := make([]interface{}, l, l)
											for i, key := range m.MapKeys() {
												p[i] = m.MapIndex(key).Interface()
											}
											f(p...)
											ok = true

	case func(R.Value):						for _, key := range m.MapKeys() {
												f(m.MapIndex(key))
											}
											ok = true

	case func(interface{}, R.Value):		for _, key := range m.MapKeys() {
												f(key.Interface(), m.MapIndex(key))
											}
											ok = true

	case func(int, R.Value):				for _, key := range m.MapKeys() {
												f(int(key.Int()), m.MapIndex(key))
											}
											ok = true

	case func(string, R.Value):				for _, key := range m.MapKeys() {
												f(key.String(), m.MapIndex(key))
											}
											ok = true

	case func(R.Value, R.Value):			for _, key := range m.MapKeys() {
												f(key, m.MapIndex(key))
											}
											ok = true

	case func(...R.Value):					l := Len(m)
											p := make([]R.Value, l, l)
											for i, key := range m.MapKeys() {
												p[i] = m.MapIndex(key)
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													l := Len(m)
													p := make([]R.Value, l, l)
													for i, key := range m.MapKeys() {
														p[i] = m.MapIndex(key)
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:			//	f(v)
																	p := make([]R.Value, 1, 1)
																	for _, key := range m.MapKeys() {
																		p[0] = m.MapIndex(key)
																		f.Call(p)
																	}
																	ok = true

													case 2:			//	f(i, v)
																	p := make([]R.Value, 2, 2)
																	for _, key := range m.MapKeys() {
																		p[0], p[1] = key, m.MapIndex(key)
																		f.Call(p)
																	}
																	ok = true
													}
												}
											}
	}
	return
}

func eachChannel(c R.Value, f interface{}) (ok bool) {
	switch f := f.(type) {
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
}

func eachFunction(g R.Value, f interface{}) (ok bool) {
	if tg := g.Type(); tg.NumOut() == 2 {
		switch tg.NumIn() {
		case 0:			switch f := f.(type) {
						case func(interface{}):					p := []R.Value{}
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0].Interface())
																}
																ok = true

						case func(int, interface{}):			p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																} 
																ok = true

						case func(interface{}, interface{}):	p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																}
																ok = true

						case func(...interface{}):				pg := []R.Value{}
																pf := make([]interface{}, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0].Interface())
																}
																f(pf...)
																ok = true

						case func(R.Value):						p := []R.Value{}
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0])
																}
																ok = true

						case func(int, R.Value):				p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																}
																ok = true

						case func(interface{}, R.Value):		p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																}
																ok = true

						case func(R.Value, R.Value):			p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(R.ValueOf(i), v[0])
																	i++
																}
																ok = true

						case func(...R.Value):					pg := []R.Value{}
																pf := make([]R.Value, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0])
																}
																f(pf...)
																ok = true

						default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																	if tf := f.Type(); tf.IsVariadic() {
																		//	f(...v)
																		i := 0
																		pg := []R.Value{}
																		pf := make([]R.Value, 0, 4)
																		for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																			pf = append(pf, v[0])
																			i++
																		}
																		f.Call(pf)
																		ok = true
																	} else {
																		switch tf.NumIn() {
																		case 1:		//	f(v)
																					pg := []R.Value{}
																					pf := make([]R.Value, 1, 1)
																					for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																						pf[0] = v[0]
																						f.Call(pf)
																					}
																					ok = true

																		case 2:		//	f(i, v)
																					i := 0
																					pg := []R.Value{}
																					pf := make([]R.Value, 2, 2)
																					for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																						pf[0], pf[1] = R.ValueOf(i), v[0]
																						f.Call(pf)
																						i++
																					}
																					ok = true
																		}
																	}
																}
						}

		case 1:			switch f := f.(type) {
						case func(interface{}):					i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(int, interface{}):			i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(interface{}, interface{}):	i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[0].IsValid(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(...interface{}):				i := 0
																pg := []R.Value{ R.ValueOf(0) }
																pf := make([]interface{}, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0].Interface())
																	i++
																	pg[i] = R.ValueOf(i)
																}
																f(pf...)
																ok = true

						case func(R.Value):						i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(int, R.Value):				i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(interface{}, R.Value):		i := 0
																p := []R.Value{ R.ValueOf(i) }
																for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true


						case func(R.Value, R.Value):			i := 0
																p := []R.Value{ R.ValueOf(i) }
																for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																	f(p[0], v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(...R.Value):					i := 0
																pg := []R.Value{ R.ValueOf(i) }
																pf := make([]R.Value, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0])
																	i++
																	pf[i] = R.ValueOf(i)
																}
																f(pf...)
																ok = true

						default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																	if tf := f.Type(); tf.IsVariadic() {
																		//	f(...v)
																		i := 0
																		pg := []R.Value{ R.ValueOf(0) }
																		pf := make([]R.Value, 0, 4)
																		for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																			pf = append(pf, v[0])
																			i++
																			pg[0] = R.ValueOf(i)
																		}
																		f.Call(pf)
																		ok = true
																	} else {
																		switch tf.NumIn() {
																		case 1:		//	f(v)
																					i := 0
																					p := []R.Value{ R.ValueOf(0) }
																					for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																						p[0] = v[0]
																						f.Call(p)
																						i++
																						p[0] = R.ValueOf(i)
																					}
																					ok = true

																		case 2:		//	f(i, v)
																					i := 0
																					pg := []R.Value{ R.ValueOf(0) }
																					pf := make([]R.Value, 2, 2)
																					for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																						pf[0], pf[1] = pg[0], v[0]
																						f.Call(pf)
																						i++
																						pg[0] = R.ValueOf(i)
																					}
																					ok = true
																		}
																	}
																}
						}
		}
	}
	return
}