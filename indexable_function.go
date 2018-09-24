package sequences

import R "reflect"

func atOffsetBoolFunction(seq func(int) bool, i int, f interface{}) {
	switch f := f.(type) {
	case func(bool):
		f(seq(i))
	case chan bool:
		f <- seq(i)
	case []chan bool:
		writeBoolChannels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetComplex64Function(seq func(int) complex64, i int, f interface{}) {
	switch f := f.(type) {
	case func(complex64):
		f(seq(i))
	case chan complex64:
		f <- seq(i)
	case []chan complex64:
		writeComplex64Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetComplex128Function(seq func(int) complex128, i int, f interface{}) {
	switch f := f.(type) {
	case func(complex128):
		f(seq(i))
	case chan complex128:
		f <- seq(i)
	case []chan complex128:
		writeComplex128Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetErrorFunction(seq func(int) error, i int, f interface{}) {
	switch f := f.(type) {
	case func(error):
		f(seq(i))
	case chan error:
		f <- seq(i)
	case []chan error:
		writeErrorChannels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetFloat32Function(seq func(int) float32, i int, f interface{}) {
	switch f := f.(type) {
	case func(float32):
		f(seq(i))
	case chan float32:
		f <- seq(i)
	case []chan float32:
		writeFloat32Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetFloat64Function(seq func(int) float64, i int, f interface{}) {
	switch f := f.(type) {
	case func(float64):
		f(seq(i))
	case chan float64:
		f <- seq(i)
	case []chan float64:
		writeFloat64Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetIntFunction(seq func(int) int, i int, f interface{}) {
	switch f := f.(type) {
	case func(int):
		f(seq(i))
	case chan int:
		f <- seq(i)
	case []chan int:
		writeIntChannels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetInt8Function(seq func(int) int8, i int, f interface{}) {
	switch f := f.(type) {
	case func(int8):
		f(seq(i))
	case chan int8:
		f <- seq(i)
	case []chan int8:
		writeInt8Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetInt16Function(seq func(int) int16, i int, f interface{}) {
	switch f := f.(type) {
	case func(int16):
		f(seq(i))
	case chan int16:
		f <- seq(i)
	case []chan int16:
		writeInt16Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetInt32Function(seq func(int) int32, i int, f interface{}) {
	switch f := f.(type) {
	case func(int32):
		f(seq(i))
	case chan int32:
		f <- seq(i)
	case []chan int32:
		writeInt32Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetInt64Function(seq func(int) int64, i int, f interface{}) {
	switch f := f.(type) {
	case func(int64):
		f(seq(i))
	case chan int64:
		f <- seq(i)
	case []chan int64:
		writeInt64Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetStringFunction(seq func(int) string, i int, f interface{}) {
	switch f := f.(type) {
	case func(string):
		f(seq(i))
	case chan string:
		f <- seq(i)
	case []chan string:
		writeStringChannels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetUintFunction(seq func(int) uint, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint):
		f(seq(i))
	case chan uint:
		f <- seq(i)
	case []chan uint:
		writeUintChannels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetUint8Function(seq func(int) uint8, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint8):
		f(seq(i))
	case chan uint8:
		f <- seq(i)
	case []chan uint8:
		writeUint8Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetUint16Function(seq func(int) uint16, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint16):
		f(seq(i))
	case chan uint16:
		f <- seq(i)
	case []chan uint16:
		writeUint16Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetUint32Function(seq func(int) uint32, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint32):
		f(seq(i))
	case chan uint32:
		f <- seq(i)
	case []chan uint32:
		writeUint32Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetUint64Function(seq func(int) uint64, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint64):
		f(seq(i))
	case chan uint64:
		f <- seq(i)
	case []chan uint64:
		writeUint64Channels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetUintptrFunction(sdeq func(int) uintptr, i int, f interface{}) {
	switch f := f.(type) {
	case func(uintptr):
		f(seq(i))
	case chan uintptr:
		f <- seq(i)
	case []chan uintptr:
		writeUintptrChannels(f, seq(i))

	case func(interface{}):
		f(seq(i))
	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan interface{}:
		f <- seq(i)

	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetInterfaceFunction(seq func(int) interface{}, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		f(seq(i))
	case chan interface{}:
		f <- seq(i)
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))

	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetRValueFunction(seq func(int) R.Value, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		f(seq(i))
	case chan interface{}:
		f <- seq(i)
	case []chan interface{}:
		writeInterfaceChannels(f, seq(i))

	case func(R.Value):
		f(R.ValueOf(seq(i)))
	case chan R.Value:
		f <- R.ValueOf(seq(i))
	case []chan R.Value:
		writeRValueChannels(f, R.ValueOf(seq(i)))
	}
}

func atOffsetFunction(seq R.Value, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		f(readFunctionInterface(seq, i))
	case chan interface{}:
		f <- readFunctionInterface(seq, i)
	case []chan interface{}:
		writeInterfaceChannels(f, readFunctionInterface(seq, i))

	case func(R.Value):
		f(readFunction(seq, i))
	case chan R.Value:
		f <- readFunction(seq, i)
	case []chan R.Value:
		writeRValueChannels(f, readFunction(seq, i))
	}
}
