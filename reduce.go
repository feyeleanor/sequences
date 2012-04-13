package sequences

import "reflect"

func reduceIndexable(enum *Enumerator, seq Indexable) (r interface{}) {
	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		r = enum.seed
		enum.reduce(func(cursor int) {
			r = f(r, seq.AtOffset(cursor))
		})
	case func(interface{}, int, interface{}) interface{}:
		r = enum.seed
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq.AtOffset(cursor))
		})
	case func(interface{}, interface{}, interface{}) interface{}:
		r = enum.seed
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq.AtOffset(cursor))
		})
	default:
		if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
			switch f.Type().NumIn() {
			case 2:
				p := make([]reflect.Value, 2, 2)
				p[0] = reflect.ValueOf(enum.seed)
				enum.reduce(func(cursor int) {
					p[1] = reflect.ValueOf(seq.AtOffset(cursor))
					p[0] = f.Call(p)[0]
				})
				r = p[0].Interface()
			case 3:
				p := make([]reflect.Value, 3, 3)
				p[0] = reflect.ValueOf(enum.seed)
				enum.reduce(func(cursor int) {
					p[1], p[2] = reflect.ValueOf(cursor), reflect.ValueOf(seq.AtOffset(cursor))
					p[0] = f.Call(p)[0]
				})
				r = p[0].Interface()
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

func reduceMappable(enum *Enumerator, function interface{}) (r interface{}) {
	panic(UNHANDLED_ITERATOR)
}

func reduceEnumerable(enum *Enumerator, f interface{}) (r interface{}) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:
		r = enum.seed
		enum.Each(func(x interface{}) {
			r = f(r, x)
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

/*
func reduceSlice(s reflect.Value, seed, f interface{}) (r interface{}, e error) {
	if l := Len(s); l > 0 {
		switch f := f.(type) {
		case func(interface{}, interface{}) interface{}:
			v := reflect.New(s.Type().Elem()).Elem()
			v.Set(reflect.ValueOf(seed))
			for i := 0; i < l; i++ {
				v = reflect.ValueOf(f(v.Interface(), s.Index(i).Interface()))
			}
			r = v.Interface()
		default:
			if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
				v := reflect.New(s.Type().Elem()).Elem()
				v.Set(reflect.ValueOf(seed))
				switch f.Type().NumIn() {
				case 2:
					for i := 0; i < l; i++ {
						v = f.Call([]reflect.Value{ v, s.Index(i) })[0]
					}
					r = v.Interface()
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		}
	}
	return
}

func reduceMap(m reflect.Value, seed, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:
		v := reflect.New(m.Type().Elem()).Elem()
		v.Set(reflect.ValueOf(seed))
		for _, key := range m.MapKeys() {
			v = reflect.ValueOf(f(v.Interface(), m.MapIndex(key).Interface()))
		}
		r = v.Interface()
	default:
		if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
			v := reflect.New(m.Type().Elem()).Elem()
			v.Set(reflect.ValueOf(seed))
			switch f.Type().NumIn() {
			case 2:
				for _, key := range m.MapKeys() {
					v = f.Call([]reflect.Value{ v, m.MapIndex(key) })[0]
				}
				r = v.Interface()
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

func reduceChan(c reflect.Value, seed, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:
		v := reflect.New(c.Type().Elem()).Elem()
		v.Set(reflect.ValueOf(seed))
		for {
			if x, open := c.Recv(); open {
				v = reflect.ValueOf(f(v.Interface(), x))
			} else {
				break
			}
		}
		r = v.Interface()
	default:
		if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
			v := reflect.New(c.Type().Elem()).Elem()
			v.Set(reflect.ValueOf(seed))
			switch f.Type().NumIn() {
			case 2:
				for {
					if x, open := c.Recv(); open {
						v = f.Call([]reflect.Value{ v, x })[0]
					} else {
						break
					}
				}
				r = v.Interface()
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

func reduceFunction(g reflect.Value, seed, f interface{}) (r interface{}, e error) {
	if t := g.Type(); t.NumOut() == 2 {
		switch t.NumIn() {
		case 0:
			switch f := f.(type) {
			case func(interface{}) interface{}:
				v := reflect.New(g.Type().Out(0)).Elem()
				v.Set(reflect.ValueOf(seed))
				for x := g.Call([]reflect.Value{}); !x[1].Bool(); x = g.Call([]reflect.Value{}) {
					v = reflect.ValueOf(f(x[0].Interface()))
				}
				r = v.Interface()
			default:
				panic(UNHANDLED_ITERATOR)
			}
		case 1:
			count := 0
			switch f := f.(type) {
			case func(interface{}) interface{}:
				v := reflect.New(g.Type().Out(0)).Elem()
				v.Set(reflect.ValueOf(seed))
				p := []reflect.Value{ reflect.ValueOf(count) }
				for x := g.Call(p); !x[1].Bool(); x = g.Call(p) {
					v = reflect.ValueOf(f(x[0].Interface()))
					count++
					p[0] = reflect.ValueOf(count)
				}
				r = v.Interface()
			default:
				panic(UNHANDLED_ITERATOR)
			}
		default:
			panic(UNHANDLED_ITERATOR)
		}
	}
	return
}
*/