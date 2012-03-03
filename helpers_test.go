package sequences

import "fmt"

type UDT int

type Error int

func (e Error) Error() string {
	return fmt.Sprintf("Error Code: %v\n", e)
}


type enumerable_slice []interface{}

func (s enumerable_slice) Step(start, span, steps int, f interface{}) (ok bool) {
	var offset	int

	switch l := len(s); {
	case span == 0:
		return
	case span > l:
		return
	case span == 1:
		switch f := f.(type) {
		case func(interface{}):
			for _, v := range s[start:start + steps] { f(v) }
			ok = true
		case func(int, interface{}):
			for i, v := range s[start:start + steps] { f(i, v) }
			ok = true
		case func(interface{}, interface{}):
			for i, v := range s[start:start + steps] { f(i, v) }
			ok = true
		case func(...interface{}):
			f(s[start:start + steps]...)
			ok = true
		}
		return
	case span > 1:
		offset = start
	}

	switch f := f.(type) {
	case func(interface{}):
		for ; steps > 0; steps-- {
			f(s[offset])
			offset += span
		}
		ok = true
	case func(int, interface{}):
		for ; steps > 0; steps-- {
			f(offset, s[offset])
			offset += span
		}
		ok = true
	case func(interface{}, interface{}):
		for ; steps > 0; steps-- {
			f(offset, s[offset])
			offset += span
		}
		ok = true
	case func(...interface{}):
		n := make([]interface{}, 0, steps)
		for ; steps > 0; steps-- {
			n = append(n, s[offset])
			offset += span
		}
		f(s...)
		ok = true
	}
	return
}


type iterable_slice []interface{}

func (s iterable_slice) Each(f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):
		for _, v := range s { f(v) }
		ok = true
	case func(int, interface{}):
		for i, v := range s { f(i, v) }
		ok = true
	case func(interface{}, interface{}):
		for i, v := range s { f(i, v) }
		ok = true
	}
	return
}


type partially_iterable_slice []interface{}
func (s partially_iterable_slice) While(r bool, f interface{}) (count int) {
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
	r := make(iterable_slice, l, l)
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
	r := make(iterable_slice, len(m), len(m))
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
	r := make(iterable_slice, len(m), len(m))
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
	if l, e := Len(m); e == nil {
		s := make(iterable_slice, l, l)
		for i := l - 1; i > -1; i-- {
			s[i] = i
		}
		r = s
	}
	return
}