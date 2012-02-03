package sequences

import R "reflect"

func eachMappable(container Mappable, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					Each(container.Keys(), func(v interface{}) {
												f(container.StoredAs(v))
												
											})
											ok = true

	case func(int, interface{}):			Each(container.Keys(), func(v interface{}) {
												f(v.(int), container.StoredAs(v))
											})
											ok = true

	case func(string, interface{}):			Each(container.Keys(), func(v interface{}) {
												f(v.(string), container.StoredAs(v))
											})
											ok = true

	case func(interface{}, interface{}):	Each(container.Keys(), func(v interface{}) {
												f(v, container.StoredAs(v))
											})
											ok = true

	case func(...interface{}):				l := Len(container)
											p := make([]interface{}, l, l)
											Each(container.Keys(), func(i int, v interface{}) {
												p[i] = container.StoredAs(v)
											})
											f(p...)
											ok = true

	case func(R.Value):						Each(container.Keys(), func(v interface{}) {
												f(R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(int, R.Value):				Each(container.Keys(), func(v interface{}) {
												f(v.(int), R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(string, R.Value):				Each(container.Keys(), func(v interface{}) {
												f(v.(string), R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(interface{}, R.Value):		Each(container.Keys(), func(v interface{}) {
												f(v, R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(R.Value, R.Value):			Each(container.Keys(), func(v interface{}) {
												f(R.ValueOf(v), R.ValueOf(container.StoredAs(v)))
											})
											ok = true

	case func(...R.Value):					l := Len(container)
											p := make([]R.Value, l, l)
											Each(container.Keys(), func(i int, v interface{}) {
												p[i] = R.ValueOf(container.StoredAs(v))
											})
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													// f(...v)
													l := Len(container)
													p := make([]R.Value, l, l)
													Each(container.Keys(), func(i int, v interface{}) {
														p[i] = R.ValueOf(container.StoredAs(v))
													})
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		Each(container.Keys(), func(v interface{}) {
																			p[0] = R.ValueOf(container.StoredAs(v))
																			f.Call(p)																			
																		})
																		ok = true

													case 2:				// f(i, v)
																		p := make([]R.Value, 2, 2)
																		Each(container.Keys(), func(v interface{}) {
																			p[0], p[1] = R.ValueOf(v), R.ValueOf(container.StoredAs(v))
																			f.Call(p)
																		})
																		ok = true
													}
												}
											}
	}
	return
}

func eachMap(m R.Value, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					for _, key := range m.MapKeys() {
												f(m.MapIndex(key).Interface())
											}
											ok = true

	case func(int, interface{}):			for _, key := range m.MapKeys() {
												f(int(key.Int()), m.MapIndex(key).Interface())
											}
											ok = true

	case func(string, interface{}):			for _, key := range m.MapKeys() {
												f(key.String(), m.MapIndex(key).Interface())
											}
											ok = true

	case func(interface{}, interface{}):	for _, key := range m.MapKeys() {
												f(key.Interface(), m.MapIndex(key).Interface())
											}
											ok = true

	case func(...interface{}):				l := Len(m)
											p := make([]interface{}, l, l)
											for i, key := range m.MapKeys() {
												p[i] = m.MapIndex(key).Interface()
											}
											f(p...)
											ok = true

	case func(R.Value):						for _, key := range m.MapKeys() {
												f(m.MapIndex(key))
											}
											ok = true

	case func(interface{}, R.Value):		for _, key := range m.MapKeys() {
												f(key.Interface(), m.MapIndex(key))
											}
											ok = true

	case func(int, R.Value):				for _, key := range m.MapKeys() {
												f(int(key.Int()), m.MapIndex(key))
											}
											ok = true

	case func(string, R.Value):				for _, key := range m.MapKeys() {
												f(key.String(), m.MapIndex(key))
											}
											ok = true

	case func(R.Value, R.Value):			for _, key := range m.MapKeys() {
												f(key, m.MapIndex(key))
											}
											ok = true

	case func(...R.Value):					l := Len(m)
											p := make([]R.Value, l, l)
											for i, key := range m.MapKeys() {
												p[i] = m.MapIndex(key)
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													l := Len(m)
													p := make([]R.Value, l, l)
													for i, key := range m.MapKeys() {
														p[i] = m.MapIndex(key)
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:			//	f(v)
																	p := make([]R.Value, 1, 1)
																	for _, key := range m.MapKeys() {
																		p[0] = m.MapIndex(key)
																		f.Call(p)
																	}
																	ok = true

													case 2:			//	f(i, v)
																	p := make([]R.Value, 2, 2)
																	for _, key := range m.MapKeys() {
																		p[0], p[1] = key, m.MapIndex(key)
																		f.Call(p)
																	}
																	ok = true
													}
												}
											}
	}
	return
}