package gosequence

type indexable_slice	[]interface{}

func (i indexable_slice) Len() int {
	return len(i)
}

func (i indexable_slice) At(x int) interface{} {
	return i[x]
}


type indexable_function	func(i int) interface{}
func (f indexable_function) Len() int {
	return 10
}

func (f indexable_function) At(x int) interface{} {
	return f(x)
}


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