package sequences

import R "reflect"

func writeBoolChannels(seq []chan bool, v bool) {
	for _, c := range seq {
		c <- v
	}
}

func writeComplex64Channels(seq []chan complex64, v complex64) {
	for _, c := range seq {
		c <- v
	}
}

func writeComplex128Channels(seq []chan complex128, v complex128) {
	for _, c := range seq {
		c <- v
	}
}

func writeErrorChannels(seq []chan error, v error) {
	for _, c := range seq {
		c <- v
	}
}

func writeFloat32Channels(seq []chan float32, v float32) {
	for _, c := range seq {
		c <- v
	}
}

func writeFloat64Channels(seq []chan float64, v float64) {
	for _, c := range seq {
		c <- v
	}
}

func writeIntChannels(seq []chan int, v int) {
	for _, c := range seq {
		c <- v
	}
}

func writeInt8Channels(seq []chan int8, v int8) {
	for _, c := range seq {
		c <- v
	}
}

func writeInt16Channels(seq []chan int16, v int16) {
	for _, c := range seq {
		c <- v
	}
}

func writeInt32Channels(seq []chan int32, v int32) {
	for _, c := range seq {
		c <- v
	}
}

func writeInt64Channels(seq []chan int64, v int64) {
	for _, c := range seq {
		c <- v
	}
}

func writeInterfaceChannels(seq []chan interface{}, v interface{}) {
	for _, c := range seq {
		c <- v
	}
}

func writeStringChannels(seq []chan string, v string) {
	for _, c := range seq {
		c <- v
	}
}

func writeUintChannels(seq []chan uint, v uint) {
	for _, c := range seq {
		c <- v
	}
}

func writeUint8Channels(seq []chan uint8, v uint8) {
	for _, c := range seq {
		c <- v
	}
}

func writeUint16Channels(seq []chan uint16, v uint16) {
	for _, c := range seq {
		c <- v
	}
}

func writeUint32Channels(seq []chan uint32, v uint32) {
	for _, c := range seq {
		c <- v
	}
}

func writeUint64Channels(seq []chan uint64, v uint64) {
	for _, c := range seq {
		c <- v
	}
}

func writeUintptrChannels(seq []chan uintptr, v uintptr) {
	for _, c := range seq {
		c <- v
	}
}

func writeRValueChannels(seq []chan R.Value, v interface{}) {
	if v, ok := v.(R.Value); ok {
		for _, c := range seq {
			c <- v
		}
	} else {
		for _, c := range seq {
			c <- R.ValueOf(v)
		}
	}
}
