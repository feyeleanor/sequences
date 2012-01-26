package sequences

type iterable_slice []interface{}

func (s iterable_slice) Each(f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):						for _, v := range s { f(v) }
												ok = true

	case func(int, interface{}):				for i, v := range s { f(i, v) }
												ok = true

	case func(interface{}, interface{}):		for i, v := range s { f(i, v) }
												ok = true
	}
	return
}


type partially_iterable_slice []interface{}
func (s partially_iterable_slice) While(r bool, f interface{}) (count int) {
	if len(s) > 0 {
		switch f := f.(type) {
		case func(interface{}) bool:					for _, v := range s {
															if f(v) != r {
																break 
															}
															count++
														}

		case func(int, interface{}) bool:				for i, v := range s {
															if f(i, v) != r {
																break
															}
															count++
														}

		case func(interface{}, interface{}) bool:		for i, v := range s {
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

func (i indexable_slice) At(x int) interface{} {
	return i[x]
}


type indexable_function func(i int) interface{}

func (f indexable_function) Len() int {
	return 10
}

func (f indexable_function) At(x int) interface{} {
	return f(x)
}


type mappable_slice []int

func (m mappable_slice) Len() int {
	return len(m)
}

func (m mappable_slice) Lookup(key interface{}) interface{} {
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

func (m mappable_map) Lookup(key interface{}) interface{} {
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

func (m mappable_string_map) Lookup(key interface{}) interface{} {
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

func (f mappable_function) Lookup(x interface{}) interface{} {
	return f(x.(int))
}

func (m mappable_function) Keys() interface{} {
	l := Len(m)
	r := make(iterable_slice, l, l)
	for i := l - 1; i > -1; i-- {
		r[i] = i
	}
	return r
}