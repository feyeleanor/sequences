package sequences

import "R" reflect

func readBoolFunction(seq func(int) (bool, bool), i int) (r bool) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readBoolFunctionRValue(seq func(int) (bool, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readBoolFunctionKeyed(seq func(int) (bool, bool)) (r, ok bool) {
	r, ok = seq(i)
	return
}

func readBoolFunctionRValueKeyed(seq func(int) (bool, bool)) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readComplex64Function(seq func(int) (complex64, bool), i int) (r complex64) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readComplex64FunctionRValue(seq func(int) (complex64, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readComplex64FunctionKeyed(seq func(int) (complex64, bool)) (r complex64, ok bool) {
	r, ok = seq(i)
	return
}

func readComplex64FunctionRValueKeyed(seq func(int) (complex64, bool)) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readErrorFunction(seq func(int) (error, bool), i int) (r error) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readErrorFunctionRValue(seq func(int) (error, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readErrorFunctionKeyed(seq func(int) (error, bool) , i int) (error) {
	r, ok = seq(i)
	return
}

func readErrorFunctionRValueKeyed(seq func(int) (error, bool) , i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readFloat32Function(seq func(int) (float32, bool), i int) (r float32) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readFloat32FunctionRValue(seq func(int) (float32, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readFloat32FunctionKeyed(seq func(int) (float32, bool), i int) (r float32, ok bool) {
	r, ok = seq(i)
	return
}

func readFloat32FunctionRValueKeyed(seq func(int) (float32, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readFloat64Function(seq func(int) (float64, bool), i int) (r float64) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readFloat64FunctionRValue(seq func(int) (float64, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readFloat64FunctionKeyed(seq func(int) (float64, bool), i int) (r float64, ok bool) {
	r, ok = seq(i)
	return
}

func readFloat64FunctionRValueKeyed(seq func(int) (float64, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readIntFunction(seq func(int) (int, bool), i int) (r int) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readIntFunctionRValue(seq func(int) (int, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readIntFunctionKeyed(seq func(int) (int, bool), i int) (r int, ok bool) {
	r, ok = seq(i)
	return
}

func readIntFunctionRValueKeyed(seq func(int) (int, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readInt8Function(seq func(int) (int8, bool), i int, i int) (r int8) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readInt8FunctionRValue(seq func(int) (int8, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readInt8FunctionKeyed(seq func(int) (int8, bool), i int) (r int8, ok bool) {
	r, ok = seq(i)
	return
}

func readInt8FunctionRValueKeyed(seq func(int) (int8, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readInt16Function(seq func(int) (int16, bool), i int) (r int16) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readInt16FunctionRValue(seq func(int) (int16, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readInt16FunctionKeyed(seq func(int) (int16, bool), i int) (r int16, ok bool) {
	r, ok = seq(i)
	return
}

func readInt16FunctionRValueKeyed(seq func(int) (int16, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readInt32Function(seq func(int) (int32, bool), i int) (r int32) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readInt32FunctionRValue(seq func(int) (int32, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readInt32FunctionKeyed(seq func(int) (int32, bool), i int) (r int32, ok bool) {
	r, ok = seq(i)
	return
}

func readInt32FunctionRValueKeyed(seq func(int) (int32, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readInt64Function(seq func(int) (int64, bool), i int) (r int64) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readInt64FunctionRValue(seq func(int) (int64, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readInt64FunctionKeyed(seq func(int) (int64, bool), i int) (r int64, ok bool) {
	r, ok = seq(i)
	return
}

func readInt64FunctionRValueKeyed(seq func(int) (int64, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readInterfaceFunction(seq func(int) (interface{}, bool), i int) (r interface{}) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readInterfaceFunctionRValue(seq func(int) (interface{}, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readInterfaceFunctionKeyed(seq func(int) (interface{}, bool), i int) (r interface{}, ok bool) {
	r, ok = seq(i)
	return
}

func readInterfaceFunctionRValueKeyed(seq func(int) (interface{}, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readStringFunction(seq func(int) (string, bool), i int) (r string) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readStringFunctionRValue(seq func(int) (string, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readStringFunctionKeyed(seq func(int) (string, bool), i int) (r string, ok bool) {
	r, ok = seq(i)
	return
}

func readStringFunctionRValueKeyed(seq func(int) (string, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readUintFunction(seq func(int) (uint, bool), i int) (r uint) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readUintFunctionRValue(seq func(int) (uint, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readUintFunctionKeyed(seq func(int) (uint, bool), i int) (r uint, ok bool) {
	r, ok = seq(i)
	return
}

func readUintFunctionRValueKeyed(seq func(int) (uint, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readUint8Function(seq func(int) (uint8, bool), i int) (r uint8) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readUint8FunctionRValue(seq func(int) (uint8, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readUint8FunctionKeyed(seq func(int) (uint8, bool), i int) (r uint8, ok bool) {
	r, ok = seq(i)
	return
}

func readUint8FunctionRValueKeyed(seq func(int) (uint8, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readUint16Function(seq func(int) (uint16, bool), i int) (r uint16) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readUint16FunctionRValue(seq func(int) (uint16, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readUint16FunctionKeyed(seq func(int) (uint16, bool), i int) (r uint16, ok bool) {
	r, ok = seq(i)
	return
}

func readUint16FunctionRValueKeyed(seq func(int) (uint16, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readUint32Function(seq func(int) (uint32, bool), i int) (r uint32) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readUint32FunctionRValue(seq func(int) (uint32, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readUint32FunctionKeyed(seq func(int) (uint32, bool), i int) (r uint32, ok bool) {
	r, ok = seq(i)
	return
}

func readUint32FunctionRValueKeyed(seq func(int) (uint32, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readUint64Function(seq func(int) (uint64, bool), i int) (r uint64) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readUint64FunctionRValue(seq func(int) (uint64, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readUint64FunctionKeyed(seq func(int) (uint64, bool), i int) (r uint64, ok bool) {
	r, ok = seq(i)
	return
}

func readUint64FunctionRValueKeyed(seq func(int) (uint64, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readUintptrFunction(seq func(int) (uintptr, bool), i int) (r uintptr) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readUintptrFunctionRValue(seq func(int) (uintptr, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = R.ValueOf(v)
	} else {
		r = R.ValueOf(false)
	}
	return
}

func readUintptrFunctionKeyed(seq func(int) (uintptr, bool), i int) (r uintptr, ok bool) {
	r, ok = seq(i)
	return
}

func readUintptrFunctionRValueKeyed(seq func(int) (uintptr, bool), i int) (r R.Value, ok bool) {
	var v bool
	r, ok = seq(i)
	r = R.ValueOf(v)
	return
}

func readRValueFunction(seq func(int) (R.Value, bool), i int) (r R.Value) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = false
	}
	return
}

func readRValueFunctionInterface(seq func(int) (R.Value, bool), i int) (r interface{}) {
	if v, ok := seq(i); ok {
		r = v
	} else {
		r = nil
	}
	return
}

func readRValueFunctionKeyed(seq func(int) (R.Value, bool), i int) (r R.Value, ok bool) {
	r, ok = seq(i)
	return
}

func readRValueFunctionKeyedInterface(seq func(int) (R.Value, bool), i int) (r interface{}, ok bool) {
	var v R.Value
	r, ok = seq(i)
	if ok {
		r = v.Interface()
	}
	return
}

// MERGE THESE TOGETHER
func readFunction(seq R.Value, i int) (r R.Value) {
	p := []R.Value{R.ValueOf(i)}
	return seq.Call(p)[0]
}

func readFunctionInterface(seq R.Value, i int) (r interface{}) {
	return readFunction(seq, i).Interface()
}

func readFunctionKeyed(seq R.Value, i int) (R.Value, bool) {
	p := []R.Value{R.ValueOf(i)}
	v := seq.Call(p)
	return v[0], v[1]
}

func readFunctionInterface(seq R.Value, i int) (r interface{}) {
	return readFunction(seq, i).Interface()
}

func readFunctionInterfaceKeyed(seq R.Value, i int) (interface{}, bool) {
	v, ok := readFunction(seq, i)
	return v.Interface(), ok
}
