package sequences

import R "reflect"

func atOffsetBoolSlice(seq []bool, i int, f interface{}) {
	switch f := f.(type) {
	case func(bool):
		f(seq[i])
	case chan bool:
		f <- f(seq[i])
	case []chan bool:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetComplex64Slice(seq []complex64, i int, f interface{}) {
	switch f := f.(type) {
	case func(complex64):
		f(seq[i])
	case chan complex64:
		f <- f(seq[i])
	case []chan complex64:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetComplex128Slice(seq []complex128, i int, f interface{}) {
	switch f := f.(type) {
	case func(complex128):
		f(seq[i])
	case chan complex128:
		f <- f(seq[i])
	case []chan complex128:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetErrorSlice(seq []error, i int, f interface{}) {
	switch f := f.(type) {
	case func(error):
		f(seq[i])
	case chan error:
		f <- f(seq[i])
	case []chan error:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetFloat32Slice(seq []float32, i int, f interface{}) {
	switch f := f.(type) {
	case func(float32):
		f(seq[i])
	case chan float32:
		f <- f(seq[i])
	case []chan float32:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetFloat64Slice(seq []float64, i int, f interface{}) {
	switch f := f.(type) {
	case func(float64):
		f(seq[i])
	case chan float64:
		f <- f(seq[i])
	case []chan float64:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetIntSlice(seq []int, i int, f interface{}) {
	switch f := f.(type) {
	case func(int):
		f(seq[i])
	case chan int:
		f <- f(seq[i])
	case []chan int:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetInt8Slice(seq []int8, i int, f interface{}) {
	switch f := f.(type) {
	case func(int8):
		f(seq[i])
	case chan int8:
		f <- f(seq[i])
	case []chan int8:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetInt16Slice(seq []int16, i int, f interface{}) {
	switch f := f.(type) {
	case func(int16):
		f(seq[i])
	case chan int16:
		f <- f(seq[i])
	case []chan int16:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetInt32Slice(seq []int32, i int, f interface{}) {
	switch f := f.(type) {
	case func(int32):
		f(seq[i])
	case chan int32:
		f <- f(seq[i])
	case []chan int32:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetInt64Slice(seq []int64, i int, f interface{}) {
	switch f := f.(type) {
	case func(int64):
		f(seq[i])
	case chan int64:
		f <- f(seq[i])
	case []chan int64:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetStringSlice(seq []string, i int, f interface{}) {
	switch f := f.(type) {
	case func(string):
		f(seq[i])
	case chan string:
		f <- f(seq[i])
	case []chan string:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetUintSlice(seq []uint, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint):
		f(seq[i])
	case chan uint:
		f <- f(seq[i])
	case []chan uint:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetUint8Slice(seq []uint8, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint8):
		f(seq[i])
	case chan uint8:
		f <- f(seq[i])
	case []chan uint8:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetUint16Slice(seq []uint16, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint16):
		f(seq[i])
	case chan uint16:
		f <- f(seq[i])
	case []chan uint16:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetUint32Slice(seq []uint32, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint32):
		f(seq[i])
	case chan uint32:
		f <- f(seq[i])
	case []chan uint32:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetUint64Slice(seq []uint64, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint64):
		f(seq[i])
	case chan uint64:
		f <- f(seq[i])
	case []chan uint64:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetUintptrSlice(seq []uintptr, i int, f interface{}) {
	switch f := f.(type) {
	case func(uintptr):
		f(seq[i])
	case chan uintptr:
		f <- f(seq[i])
	case []chan uintptr:
		writeBoolChannels(f, seq[i])

	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- f(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- f(R.ValueOf(seq[i]))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetInterfaceSlice(seq []interface{}, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		f(seq[i])
	case chan interface{}:
		f <- seq[i]
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i])

	case func(R.Value):
		f(R.ValueOf(seq[i]))
	case chan R.Value:
		f <- R.ValueOf(seq[i])
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq[i]))
	}
}

func atOffsetRValueSlice(seq []R.Value, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		f(seq[i].Interface())
	case chan interface{}:
		f <- seq[i].Interface()
	case []chan interface{}:
		writeInterfaceChannels(f, seq[i].Interface())

	case func(R.Value):
		f(seq[i])
	case chan R.Value:
		f <- seq[i]
	case []chan R.Value:
		writeRValueChannels(f, seq[i])
	}
}

func atOffsetSlice(seq R.Value, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		f(readSliceInterface(seq, i))
	case chan interface{}:
		f <- readSliceInterface(seq, i)
	case []chan interface{}:
		writeInterfaceChannels(f, readSliceInterface(seq, i))

	case func(R.Value):
		f(readSlice(seq, i))
	case chan R.Value:
		f <- readSlice(seq, i)
	case []chan R.Value:
		writeRValueChannels(f, readSlice(seq, i))
	}
}
