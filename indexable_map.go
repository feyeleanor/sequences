package sequences

import R "reflect"

func atOffsetBoolMap(seq map[int]bool, i int, f interface{}) {
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

func atOffsetComplex64Map(seq map[int]complex64, i int, f interface{}) {
	switch f := f.(type) {
	case func(complex64):
		f(seq[i])
	case chan complex64:
		f <- seq[i]
	case []chan complex64:
		writeComplex64Channels(f, seq[i])

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

func atOffsetComplex128Map(g map[int]complex128, i int, f interface{}) {
	switch f := f.(type) {
	case func(complex128):
		f(seq[i])
	case chan complex128:
		f <- seq[i]
	case []chan complex128:
		writeComplex128Channels(f, seq[i])

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

func atOffsetErrorMap(seq map[int]error, i int, f interface{}) {
	switch f := f.(type) {
	case func(error):
		f(seq[i])
	case chan error:
		f <- seq[i]
	case []chan error:
		writeErrorChannels(f, seq[i])

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

func atOffsetFloat32Map(seq map[int]float32, i int, f interface{}) {
	switch f := f.(type) {
	case func(float32):
		f(seq[i])
	case chan float32:
		f <- seq[i]
	case []chan float32:
		writeFloat32Channels(f, seq[i])

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

func atOffsetFloat64Map(seq map[int]float64, i int, f interface{}) {
	switch f := f.(type) {
	case func(float64):
		f(seq[i])
	case chan float64:
		f <- seq[i]
	case []chan float64:
		writeFloat64Channels(f, seq[i])

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

func atOffsetIntMap(seq map[int]int, i int, f interface{}) {
	switch f := f.(type) {
	case func(int):
		f(seq[i])
	case chan int:
		f <- seq[i]
	case []chan int:
		writeIntChannels(f, seq[i])

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

func atOffsetInt8Map(seq map[int]int8, i int, f interface{}) {
	switch f := f.(type) {
	case func(int8):
		f(seq[i])
	case chan int8:
		f <- seq[i]
	case []chan int8:
		writeInt8Channels(f, seq[i])

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

func atOffsetInt16Map(seq map[int]int16, i int, f interface{}) {
	switch f := f.(type) {
	case func(int16):
		f(seq[i])
	case chan int16:
		f <- seq[i]
	case []chan int16:
		writeIntChannels(f, seq[i])

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

func atOffsetInt32Map(seq map[int]int32, i int, f interface{}) {
	switch f := f.(type) {
	case func(int32):
		f(seq[i])
	case chan int32:
		f <- seq[i]
	case []chan int32:
		writeIntChannels(f, seq[i])

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

func atOffsetInt64Map(seq map[int]int64, i int, f interface{}) {
	switch f := f.(type) {
	case func(int64):
		f(seq[i])
	case chan int64:
		f <- seq[i]
	case []chan int64:
		writeInt64Channels(f, seq[i])

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

func atOffsetStringMap(seq map[int]string, i int, f interface{}) {
	switch f := f.(type) {
	case func(string):
		f(seq[i])
	case chan string:
		f <- seq[i]
	case []chan string:
		writeStringChannels(f, seq[i])

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

func atOffsetUintMap(seq map[int]uint, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint):
		f(seq[i])
	case chan uint:
		f <- seq[i]
	case []chan uint:
		writeUintChannels(f, seq[i])

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

func atOffsetUint8Map(seq map[int]uint8, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint8):
		f(seq[i])
	case chan uint8:
		f <- seq[i]
	case []chan uint8:
		writeUint8Channels(f, seq[i])

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

func atOffsetUint16Map(seq map[int]uint16, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint16):
		f(seq[i])
	case chan uint16:
		f <- seq[i]
	case []chan uint16:
		writeUint16Channels(f, seq[i])

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

func atOffsetUint32Map(seq map[int]uint32, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint32):
		f(seq[i])
	case chan uint32:
		f <- seq[i]
	case []chan uint32:
		writeUint32Channels(f, seq[i])

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

func atOffsetUint64Map(seq map[int]uint64, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint64):
		f(seq[i])
	case chan uint64:
		f <- seq[i]
	case []chan uint64:
		writeUint64Channels(f, seq[i])

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

func atOffsetUintptrMap(seq map[int]uintptr, i int, f interface{}) {
	switch f := f.(type) {
	case func(uintptr):
		f(seq[i])
	case chan uintptr:
		f <- seq[i]
	case []chan uintptr:
		writeUintptrChannels(f, seq[i])

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

func atOffsetInterfaceMap(seq map[int]interface{}, i int, f interface{}) {
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

func atOffsetRValueMap(seq map[int]R.Value, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		f(R.ValueOf(seq[i]))
	case chan interface{}:
		f <- R.ValueOf(seq[i])
	case []chan interface{}:
		writeInterfaceChannels(f, R.ValueOf(seq[i]))

	case func(R.Value):
		f(seq[i])
	case chan R.Value:
		f <- seq[i]
	case []chan R.Value:
		writeRValueChannels(f, seq[i])
	}
}

func atOffsetMap(seq R.Value, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		f(readMapInterface(seq, i))
	case chan interface{}:
		f <- readMapInterface(seq, i)
	case []chan interface{}:
		writeInterfaceChannels(f, readMapInterface(seq, i))

	case func(R.Value):
		f(readMap(seq, i))
	case chan R.Value:
		f <- readMap(seq, i)
	case []chan R.Value:
		writeRValueChannels(f, readMap(seq, i))
	}
}
