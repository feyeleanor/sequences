package sequences

import R "reflect"

func eachMappable(container Mappable, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		Each(container.Keys(), func(v interface{}) {
			f(container.StoredAs(v))
		})
	case func(int, interface{}):
		Each(container.Keys(), func(v interface{}) {
			f(v.(int), container.StoredAs(v))
		})
	case func(string, interface{}):
		Each(container.Keys(), func(v interface{}) {
			f(v.(string), container.StoredAs(v))
		})
	case func(interface{}, interface{}):
		Each(container.Keys(), func(v interface{}) {
			f(v, container.StoredAs(v))
		})
	case func(R.Value):
		Each(container.Keys(), func(v interface{}) {
			f(R.ValueOf(container.StoredAs(v)))
		})
	case func(int, R.Value):
		Each(container.Keys(), func(v interface{}) {
			f(v.(int), R.ValueOf(container.StoredAs(v)))
		})
	case func(string, R.Value):
		Each(container.Keys(), func(v interface{}) {
			f(v.(string), R.ValueOf(container.StoredAs(v)))
		})
	case func(interface{}, R.Value):
		Each(container.Keys(), func(v interface{}) {
			f(v, R.ValueOf(container.StoredAs(v)))
		})
	case func(R.Value, R.Value):
		Each(container.Keys(), func(v interface{}) {
			f(R.ValueOf(v), R.ValueOf(container.StoredAs(v)))
		})
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 1: //	f(v)
					p := make([]R.Value, 1, 1)
					Each(container.Keys(), func(v interface{}) {
						p[0] = R.ValueOf(container.StoredAs(v))
						f.Call(p)
					})
				case 2: // f(i, v)
					p := make([]R.Value, 2, 2)
					Each(container.Keys(), func(v interface{}) {
						p[0], p[1] = R.ValueOf(v), R.ValueOf(container.StoredAs(v))
						f.Call(p)
					})
				}
			}
		}
	}
}

func eachMap(m R.Value, f interface{}) {
	var key R.Value
	switch f := f.(type) {
	case func(interface{}):
		for _, key = range m.MapKeys() {
			f(m.MapIndex(key).Interface())
		}
	case func(int, interface{}):
		for _, key = range m.MapKeys() {
			f(int(key.Int()), m.MapIndex(key).Interface())
		}
	case func(string, interface{}):
		for _, key = range m.MapKeys() {
			f(key.String(), m.MapIndex(key).Interface())
		}
	case func(interface{}, interface{}):
		for _, key = range m.MapKeys() {
			f(key.Interface(), m.MapIndex(key).Interface())
		}
	case func(R.Value):
		for _, key = range m.MapKeys() {
			f(m.MapIndex(key))
		}
	case func(interface{}, R.Value):
		for _, key = range m.MapKeys() {
			f(key.Interface(), m.MapIndex(key))
		}
	case func(int, R.Value):
		for _, key = range m.MapKeys() {
			f(int(key.Int()), m.MapIndex(key))
		}
	case func(string, R.Value):
		for _, key = range m.MapKeys() {
			f(key.String(), m.MapIndex(key))
		}
	case func(R.Value, R.Value):
		for _, key = range m.MapKeys() {
			f(key, m.MapIndex(key))
		}
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 1: //	f(v)
					p := make([]R.Value, 1, 1)
					for _, key = range m.MapKeys() {
						p[0] = m.MapIndex(key)
						f.Call(p)
					}
				case 2: //	f(i, v)
					p := make([]R.Value, 2, 2)
					for _, key = range m.MapKeys() {
						p[0], p[1] = key, m.MapIndex(key)
						f.Call(p)
					}
				}
			}
		}
	}
}
