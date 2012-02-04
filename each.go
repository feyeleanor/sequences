package sequences

import R "reflect"

func eachMappable(container Mappable, f interface{}) (i int) {
	switch f := f.(type) {
	case func(interface{}):					i = Each(container.Keys(), func(v interface{}) {
												f(container.StoredAs(v))
												
											})

	case func(int, interface{}):			i = Each(container.Keys(), func(v interface{}) {
												f(v.(int), container.StoredAs(v))
											})

	case func(string, interface{}):			i = Each(container.Keys(), func(v interface{}) {
												f(v.(string), container.StoredAs(v))
											})

	case func(interface{}, interface{}):	i = Each(container.Keys(), func(v interface{}) {
												f(v, container.StoredAs(v))
											})

	case func(...interface{}):				l := Len(container)
											p := make([]interface{}, l, l)
											Each(container.Keys(), func(i int, v interface{}) {
												p[i] = container.StoredAs(v)
											})
											f(p...)
											i = 1

	case func(R.Value):						i = Each(container.Keys(), func(v interface{}) {
												f(R.ValueOf(container.StoredAs(v)))
											})

	case func(int, R.Value):				i = Each(container.Keys(), func(v interface{}) {
												f(v.(int), R.ValueOf(container.StoredAs(v)))
											})

	case func(string, R.Value):				i = Each(container.Keys(), func(v interface{}) {
												f(v.(string), R.ValueOf(container.StoredAs(v)))
											})

	case func(interface{}, R.Value):		i = Each(container.Keys(), func(v interface{}) {
												f(v, R.ValueOf(container.StoredAs(v)))
											})

	case func(R.Value, R.Value):			i = Each(container.Keys(), func(v interface{}) {
												f(R.ValueOf(v), R.ValueOf(container.StoredAs(v)))
											})

	case func(...R.Value):					l := Len(container)
											p := make([]R.Value, l, l)
											Each(container.Keys(), func(i int, v interface{}) {
												p[i] = R.ValueOf(container.StoredAs(v))
											})
											f(p...)
											i = 1

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													// f(...v)
													l := Len(container)
													p := make([]R.Value, l, l)
													i = Each(container.Keys(), func(i int, v interface{}) {
														p[i] = R.ValueOf(container.StoredAs(v))
													})
													f.Call(p)
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		i = Each(container.Keys(), func(v interface{}) {
																			p[0] = R.ValueOf(container.StoredAs(v))
																			f.Call(p)																			
																		})

													case 2:				// f(i, v)
																		p := make([]R.Value, 2, 2)
																		i = Each(container.Keys(), func(v interface{}) {
																			p[0], p[1] = R.ValueOf(v), R.ValueOf(container.StoredAs(v))
																			f.Call(p)
																		})
													}
												}
											}
	}
	return
}

func eachMap(m R.Value, f interface{}) (i int) {
	var key	R.Value
	switch f := f.(type) {
	case func(interface{}):					for i, key = range m.MapKeys() {
												f(m.MapIndex(key).Interface())
											}
											i++

	case func(int, interface{}):			for i, key = range m.MapKeys() {
												f(int(key.Int()), m.MapIndex(key).Interface())
											}
											i++

	case func(string, interface{}):			for i, key = range m.MapKeys() {
												f(key.String(), m.MapIndex(key).Interface())
											}
											i++

	case func(interface{}, interface{}):	for i, key = range m.MapKeys() {
												f(key.Interface(), m.MapIndex(key).Interface())
											}
											i++

	case func(...interface{}):				l := Len(m)
											p := make([]interface{}, l, l)
											for i, key = range m.MapKeys() {
												p[i] = m.MapIndex(key).Interface()
											}
											f(p...)
											i = 1

	case func(R.Value):						for i, key = range m.MapKeys() {
												f(m.MapIndex(key))
											}
											i++

	case func(interface{}, R.Value):		for i, key = range m.MapKeys() {
												f(key.Interface(), m.MapIndex(key))
											}
											i++

	case func(int, R.Value):				for i, key = range m.MapKeys() {
												f(int(key.Int()), m.MapIndex(key))
											}
											i++

	case func(string, R.Value):				for i, key = range m.MapKeys() {
												f(key.String(), m.MapIndex(key))
											}
											i++

	case func(R.Value, R.Value):			for i, key = range m.MapKeys() {
												f(key, m.MapIndex(key))
											}
											i++

	case func(...R.Value):					l := Len(m)
											p := make([]R.Value, l, l)
											for i, key = range m.MapKeys() {
												p[i] = m.MapIndex(key)
											}
											f(p...)
											i = 1

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													l := Len(m)
													p := make([]R.Value, l, l)
													for i, key = range m.MapKeys() {
														p[i] = m.MapIndex(key)
													}
													f.Call(p)
													i = 1

												} else {
													switch t.NumIn() {
													case 1:			//	f(v)
																	p := make([]R.Value, 1, 1)
																	for i, key = range m.MapKeys() {
																		p[0] = m.MapIndex(key)
																		f.Call(p)
																	}
																	i++

													case 2:			//	f(i, v)
																	p := make([]R.Value, 2, 2)
																	for i, key = range m.MapKeys() {
																		p[0], p[1] = key, m.MapIndex(key)
																		f.Call(p)
																	}
																	i++
													}
												}
											}
	}
	return
}