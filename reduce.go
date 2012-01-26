package gosequence

import "reflect"

func reduceIndexable(c Indexable, seed, f interface{}) (r interface{}, ok bool) {
	end := Len(c)
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:	for i := 0; i < end; i++ {
															seed = f(seed, c.At(i))
														}
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															switch f.Type().NumIn() {
															case 2:				p := make([]reflect.Value, 2, 2)
																				for i := 0; i < end; i++ {
																					p[0], p[1] = reflect.ValueOf(i), reflect.ValueOf(c.At(i))
																					f.Call(p)
																				}
																				ok = true
															}
														}
	}
	return
}

func reduceMappable(c Mappable, seed, function interface{}) (r interface{}, ok bool) {
	return
}

func reduceIterable(container Iterable, seed, f interface{}) (r interface{}, ok bool) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:		r = seed
															Each(container, func(x interface{}) {
																r = f(r, x)
															})
															ok = true
	}
	return
}

func reduceSlice(s reflect.Value, seed, f interface{}) (r interface{}, ok bool) {
	end := Len(s)
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:	v := reflect.New(s.Type().Elem()).Elem()
														v.Set(reflect.ValueOf(seed))
														for i := 0; i < end; i++ {
															v = reflect.ValueOf(f(v.Interface(), s.Index(i).Interface()))
														}
														r = v.Interface()
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															v := reflect.New(s.Type().Elem()).Elem()
															v.Set(reflect.ValueOf(seed))
															switch f.Type().NumIn() {
															case 2:				for i, end := 0, Len(s); i < end; i++ {
																					v = f.Call([]reflect.Value{ v, s.Index(i) })[0]
																				}
																				ok = true
															}
															if ok {
																r = v.Interface()
															}
														}
	}
	return
}

func reduceMap(m reflect.Value, seed, f interface{}) (r interface{}, ok bool) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:	v := reflect.New(m.Type().Elem()).Elem()
														v.Set(reflect.ValueOf(seed))
														for _, key := range m.MapKeys() {
															v = reflect.ValueOf(f(v.Interface(), m.MapIndex(key).Interface()))
														}
														r = v.Interface()
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															v := reflect.New(m.Type().Elem()).Elem()
															v.Set(reflect.ValueOf(seed))
															switch f.Type().NumIn() {
															case 2:				for _, key := range m.MapKeys() {
																					v = f.Call([]reflect.Value{ v, m.MapIndex(key) })[0]
																				}
																				ok = true
															}
															if ok {
																r = v.Interface()
															}
														}
	}
	return
}

func reduceChan(c reflect.Value, seed, f interface{}) (r interface{}, ok bool) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:	v := reflect.New(c.Type().Elem()).Elem()
														v.Set(reflect.ValueOf(seed))
														for {
															if x, open := c.Recv(); open {
																v = reflect.ValueOf(f(v.Interface(), x))
															} else {
																break
															}
														}
														r = v.Interface()
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															v := reflect.New(c.Type().Elem()).Elem()
															v.Set(reflect.ValueOf(seed))
															switch f.Type().NumIn() {
															case 2:				for {
																					if x, open := c.Recv(); open {
																						v = f.Call([]reflect.Value{ v, x })[0]
																					} else {
																						break
																					}
																				}
																				ok = true
															}
															if ok {
																r = v.Interface()
															}
														}
	}
	return
}

func reduceFunction(g reflect.Value, seed, f interface{}) (r interface{}, ok bool) {
	if t := g.Type(); t.NumOut() == 2 {
		switch t.NumIn() {
		case 0:			switch f := f.(type) {
						case func(interface{}) interface{}:		v := reflect.New(g.Type().Out(0)).Elem()
																v.Set(reflect.ValueOf(seed))
																for x := g.Call([]reflect.Value{}); !x[1].Bool(); x = g.Call([]reflect.Value{}) {
																	v = reflect.ValueOf(f(x[0].Interface()))
																}
																r = v.Interface()
																ok = true
						}

		case 1:			count := 0
						switch f := f.(type) {
						case func(interface{}) interface{}:		v := reflect.New(g.Type().Out(0)).Elem()
																v.Set(reflect.ValueOf(seed))
																p := []reflect.Value{ reflect.ValueOf(count) }
																for x := g.Call(p); !x[1].Bool(); x = g.Call(p) {
																	v = reflect.ValueOf(f(x[0].Interface()))
																	count++
																	p[0] = reflect.ValueOf(count)
																}
																r = v.Interface()
																ok = true
						}
		}
	}
	return
}