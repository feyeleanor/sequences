package sequences

import R "reflect"

func (enum Enumerator) eachboolFunction(g func(int) bool) {
	switch f := enum.f.(type) {
	case func(bool):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, bool):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, bool):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachcomplex64Function(g func(int) complex64) {
	switch f := enum.f.(type) {
	case func(complex64):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, complex64):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, complex64):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachcomplex128Function(g func(int) complex128) {
	switch f := enum.f.(type) {
	case func(complex128):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, complex128):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, complex128):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eacherrorFunction(g func(int) error) {
	switch f := enum.f.(type) {
	case func(error):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, error):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, error):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachfloat32Function(g func(int) float32) {
	switch f := enum.f.(type) {
	case func(float32):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, float32):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, float32):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachfloat64Function(g func(int) float64) {
	switch f := enum.f.(type) {
	case func(float64):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, float64):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, float64):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachintFunction(g func(int) int) {
	switch f := enum.f.(type) {
	case func(int):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachint8Function(g func(int) int8) {
	switch f := enum.f.(type) {
	case func(int8):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int8):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int8):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachint16Function(g func(int) int16) {
	switch f := enum.f.(type) {
	case func(int16):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int16):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int16):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachint32Function(g func(int) int32) {
	switch f := enum.f.(type) {
	case func(int32):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int32):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int32):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachint64Function(g func(int) int64) {
	switch f := enum.f.(type) {
	case func(int64):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int64):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int64):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachinterfaceFunction(g func(int) interface{}) {
	switch f := enum.f.(type) {
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		//	Defer further processing to reflection
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachstringFunction(g func(int) string) {
	switch f := enum.f.(type) {
	case func(string):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, string):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, string):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachuintFunction(g func(int) uint) {
	switch f := enum.f.(type) {
	case func(uint):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachuint8Function(g func(int) uint8) {
	switch f := enum.f.(type) {
	case func(uint8):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint8):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint8):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachuint16Function(g func(int) uint16) {
	switch f := enum.f.(type) {
	case func(uint16):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint16):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint16):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachuint32Function(g func(int) uint32) {
	switch f := enum.f.(type) {
	case func(uint32):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint32):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint32):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachuint64Function(g func(int) uint64) {
	switch f := enum.f.(type) {
	case func(uint64):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint64):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint64):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachuintptrFunction(g func(int) uintptr) {
	switch f := enum.f.(type) {
	case func(uintptr):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uintptr):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uintptr):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}

func (enum Enumerator) eachRValueFunction(g func(int) R.Value) {
	switch f := enum.f.(type) {
	case func(interface{}):
		enum.each(func(cursor int) {
			f(g(cursor).Interface())
		})
	case func(int, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor).Interface())
		})
	case func(interface{}, interface{}):
		enum.each(func(cursor int) {
			f(cursor, g(cursor).Interface())
		})
	case func(R.Value):
		enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, R.Value):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, R.Value):
		enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value, R.Value):
		enum.each(func(cursor int) {
			f(R.ValueOf(cursor), g(cursor))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
}
