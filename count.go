package sequences

import R "reflect"

func countFunction(g R.Value, step int, f interface{}) (count int) {
	if tg := g.Type(); tg.NumOut() == 2 {
		offset := step
		switch tg.NumIn() {
		case 0:
			p := []R.Value{}
			switch f := f.(type) {
			case func(interface{}) bool:
				for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					if offset == 0 {
						f(v[0].Interface())
						offset = step
						count++
					}
				}
			case func(int, interface{}) bool:
				for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					if offset == 0 {
						f(i, v[0].Interface())
						offset = step
						count++
					}
					i++
				}
			case func(interface{}, interface{}) bool:
				for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					if offset == 0 {
						f(i, v[0].Interface())
						offset = step
						count++
					}
					i++
				}
			case func(R.Value):
				for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					if offset == 0 {
						f(v[0])
						offset = step
						count++
					}
				}
			case func(int, R.Value):
				for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					if offset == 0 {
						f(i, v[0])
						offset = step
						count++
					}
					i++
				}
			case func(interface{}, R.Value):
				for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					if offset == 0 {
						f(i, v[0])
						offset = step
						count++
					}
					i++
				}
			case func(R.Value, R.Value):
				for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					if offset == 0 {
						f(R.ValueOf(i), v[0])
						offset = step
						count++
					}
					i++
				}
			default:
				if f := R.ValueOf(f); f.Kind() == R.Func {
					if tf := f.Type(); !tf.IsVariadic() {
						switch tf.NumIn() {
						case 1:
							//	f(v)
							pf := make([]R.Value, 1, 1)
							for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
								offset--
								if offset == 0 {
									pf[0] = v[0]
									f.Call(pf)
									offset = step
									count++
								}
							}
						case 2:
							//	f(i, v)
							i := 0
							pf := make([]R.Value, 2, 2)
							for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
								offset--
								if offset == 0 {
									pf[0], pf[1] = R.ValueOf(i), v[0]
									f.Call(pf)
									offset = step
									count++
								}
								i++
							}
						}
					}
				}
			}
		case 1:
			var i int
			p := []R.Value{R.ValueOf(0)}
			switch f := f.(type) {
			case func(interface{}):
				for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					i++
					if offset == 0 {
						f(v[0].Interface())
						p[0] = R.ValueOf(i)
						offset = step
						count++
					}
				}
			case func(int, interface{}):
				for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					i++
					if offset == 0 {
						f(i, v[0].Interface())
						p[0] = R.ValueOf(i)
						offset = step
						count++
					}
				}
			case func(interface{}, interface{}):
				for v := g.Call(p); !v[0].IsValid(); v = g.Call(p) {
					offset--
					i++
					if offset == step {
						f(i, v[0].Interface())
						p[0] = R.ValueOf(i)
						offset = step
						count++
					}
				}
			case func(R.Value):
				for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					i++
					if offset == 0 {
						f(v[0])
						p[0] = R.ValueOf(i)
						offset = step
						count++
					}
				}
			case func(int, R.Value):
				for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
					offset--
					i++
					if offset == 0 {
						f(i, v[0])
						p[0] = R.ValueOf(i)
						offset = step
						count++
					}
				}
			case func(interface{}, R.Value):
				for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
					offset--
					i++
					if offset == 0 {
						f(i, v[0])
						p[0] = R.ValueOf(i)
						offset = step
						count++
					}
				}
			case func(R.Value, R.Value):
				for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
					offset--
					i++
					if offset == 0 {
						f(p[0], v[0])
						p[0] = R.ValueOf(i)
						offset = step
						count++
					}
				}
			default:
				if f := R.ValueOf(f); f.Kind() == R.Func {
					if tf := f.Type(); !tf.IsVariadic() {
					} else {
						switch tf.NumIn() {
						case 1:
							//	f(v)
							for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
								offset--
								i++
								if offset == 0 {
									p[0] = v[0]
									f.Call(p)
									p[0] = R.ValueOf(i)
									offset = step
									count++
								}
							}
						case 2:
							//	f(i, v)
							pf := make([]R.Value, 2, 2)
							for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
								offset--
								i++
								if offset == 0 {
									pf[0], pf[1] = p[0], v[0]
									f.Call(pf)
									p[0] = R.ValueOf(i)
									offset = step
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
