package sequences

import R "reflect"

func (enum Enumerator) eachboolFunction(g func(int) bool) (i int) {
	switch f := enum.f.(type) {
	case func(bool):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, bool):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, bool):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachcomplex64Function(g func(int) complex64) (i int) {
	switch f := enum.f.(type) {
	case func(complex64):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, complex64):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, complex64):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachcomplex128Function(g func(int) complex128) (i int) {
	switch f := enum.f.(type) {
	case func(complex128):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, complex128):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, complex128):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eacherrorFunction(g func(int) error) (i int) {
	switch f := enum.f.(type) {
	case func(error):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, error):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, error):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachfloat32Function(g func(int) float32) (i int) {
	switch f := enum.f.(type) {
	case func(float32):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, float32):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, float32):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachfloat64Function(g func(int) float64) (i int) {
	switch f := enum.f.(type) {
	case func(float64):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, float64):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, float64):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachintFunction(g func(int) int) (i int) {
	switch f := enum.f.(type) {
	case func(int):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachint8Function(g func(int) int8) (i int) {
	switch f := enum.f.(type) {
	case func(int8):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int8):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int8):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachint16Function(g func(int) int16) (i int) {
	switch f := enum.f.(type) {
	case func(int16):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int16):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int16):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachint32Function(g func(int) int32) (i int) {
	switch f := enum.f.(type) {
	case func(int32):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int32):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int32):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachint64Function(g func(int) int64) (i int) {
	switch f := enum.f.(type) {
	case func(int64):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, int64):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, int64):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachinterfaceFunction(g func(int) interface{}) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		//	Defer further processing to reflection
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachstringFunction(g func(int) string) (i int) {
	switch f := enum.f.(type) {
	case func(string):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, string):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, string):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachuintFunction(g func(int) uint) (i int) {
	switch f := enum.f.(type) {
	case func(uint):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachuint8Function(g func(int) uint8) (i int) {
	switch f := enum.f.(type) {
	case func(uint8):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint8):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint8):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachuint16Function(g func(int) uint16) (i int) {
	switch f := enum.f.(type) {
	case func(uint16):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint16):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint16):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachuint32Function(g func(int) uint32) (i int) {
	switch f := enum.f.(type) {
	case func(uint32):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint32):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint32):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachuint64Function(g func(int) uint64) (i int) {
	switch f := enum.f.(type) {
	case func(uint64):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uint64):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uint64):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachuintptrFunction(g func(int) uintptr) (i int) {
	switch f := enum.f.(type) {
	case func(uintptr):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, uintptr):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, uintptr):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(g(cursor)))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(g(cursor)))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(g(cursor)))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func (enum Enumerator) eachRValueFunction(g func(int) R.Value) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(g(cursor).Interface())
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor).Interface())
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor).Interface())
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(g(cursor))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, g(cursor))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), g(cursor))
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}