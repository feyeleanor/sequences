package sequences

import (
	"fmt"
	R "reflect"
)

type UDT int

func isPositive(i interface{}) bool {
	if i, ok := i.(int); ok {
		return i > 0
	}
	return false
}


type Error int

func (e Error) Error() string {
	return fmt.Sprintf("Error Code: %v\n", e)
}


type enumerable_slice []interface{}

func (s enumerable_slice) Len() int {
	return len(s)
}

func (s enumerable_slice) Each(enum *Enumerator) (i int) {
	switch f := enum.F().(type) {
	case func(interface{}):
		i = enum.each(func(cursor int) {
			f(s[cursor])
		})
	case func(int, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, s[cursor])
		})
	case func(interface{}, interface{}):
		i = enum.each(func(cursor int) {
			f(cursor, s[cursor])
		})
	case func(R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(s[cursor]))
		})
	case func(int, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(s[cursor]))
		})
	case func(interface{}, R.Value):
		i = enum.each(func(cursor int) {
			f(cursor, R.ValueOf(s[cursor]))
		})
	case func(R.Value, R.Value):
		i = enum.each(func(cursor int) {
			f(R.ValueOf(cursor), R.ValueOf(s[cursor]))
		})
	}
	return
}


type partially_enumerable_slice []interface{}
func (s partially_enumerable_slice) While(r bool, f interface{}) (count int) {
	if len(s) > 0 {
		switch f := f.(type) {
		case func(interface{}) bool:
			for _, v := range s {
				if f(v) != r {
					break 
				}
				count++
			}
		case func(int, interface{}) bool:
			for i, v := range s {
				if f(i, v) != r {
					break
				}
				count++
			}
		case func(interface{}, interface{}) bool:
			for i, v := range s {
				if f(i, v) != r {
					break
				}
				count++
			}
		}
	}
	return
}


type indexable_slice []interface{}

func (i indexable_slice) Len() int {
	return len(i)
}

func (i indexable_slice) AtOffset(x int) interface{} {
	return i[x]
}


type indexable_function func(i int) interface{}

func (f indexable_function) Len() int {
	return 10
}

func (f indexable_function) AtOffset(x int) interface{} {
	if x > f.Len() {
		PanicWithIndex(x)
	}
	return f(x)
}


type mappable_slice []int

func (m mappable_slice) Len() int {
	return len(m)
}

func (m mappable_slice) StoredAs(key interface{}) interface{} {
	return m[key.(int)]
}

func (m mappable_slice) Keys() interface{} {
	l := len(m)
	r := make(enumerable_slice, l, l)
	for i := l - 1; i > -1; i-- {
		r[i] = i
	}
	return r
}


type mappable_map map[int]int

func (m mappable_map) Len() int {
	return len(m)
}

func (m mappable_map) StoredAs(key interface{}) interface{} {
	return m[key.(int)]
}

func (m mappable_map) Keys() interface{} {
	r := make(enumerable_slice, len(m), len(m))
	i := 0
	for k, _ := range m {
		r[i] = k
		i++
	}
	return r
}


type mappable_string_map map[string]int

func (m mappable_string_map) Len() int {
	return len(m)
}

func (m mappable_string_map) StoredAs(key interface{}) interface{} {
	return m[key.(string)]
}

func (m mappable_string_map) Keys() interface{} {
	r := make(enumerable_slice, len(m), len(m))
	i := 0
	for k, _ := range m {
		r[i] = k
		i++
	}
	return r
}


type mappable_function func(i int) interface{}

func (f mappable_function) Len() int {
	return 10
}

func (f mappable_function) StoredAs(x interface{}) interface{} {
	return f(x.(int))
}

func (m mappable_function) Keys() (r interface{}) {
	if l := Len(m); l > 0 {
		s := make(enumerable_slice, l, l)
		for i := l - 1; i > -1; i-- {
			s[i] = i
		}
		r = s
	}
	return
}