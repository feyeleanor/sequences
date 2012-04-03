package sequences

import "reflect"

func reduceIndexable(c Indexable, seed, f interface{}) (r interface{}, e error) {
	if l := Len(c); l > 0 {
		switch f := f.(type) {
		case func(interface{}, interface{}) interface{}:
			for i := 0; i < l; i++ {
				seed = f(seed, c.AtOffset(i))
			}
		default:
			if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
				switch f.Type().NumIn() {
				case 2:
					p := make([]reflect.Value, 2, 2)
					for i := 0; i < l; i++ {
						p[0], p[1] = reflect.ValueOf(i), reflect.ValueOf(c.AtOffset(i))
						f.Call(p)
					}
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		}
	}
	return
}

func reduceMappable(c Mappable, seed, function interface{}) (r interface{}, e error) {
	panic(UNHANDLED_ITERATOR)
}

func reduceEnumerable(container Enumerable, seed, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:
		r = seed
		Each(container, func(x interface{}) {
			r = f(r, x)
		})
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

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