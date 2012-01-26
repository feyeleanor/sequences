package sequences

import R "reflect"

func eachIndexable(container Indexable, f interface{}) (ok bool) {
	end := Len(container)
	switch f := f.(type) {
	case func(interface{}):					for i := 0; i < end; i++ {
												f(container.At(i))
											}
											ok = true

	case func(int, interface{}):			for i := 0; i < end; i++ {
												f(i, container.At(i))
											}
											ok = true

	case func(interface{}, interface{}):	for i := 0; i < end; i++ {
												f(i, container.At(i))
											}
											ok = true

	case func(...interface{}):				p := make([]interface{}, end, end)
											for i := 0; i < end; i++ {
												p[i] = container.At(i)
											}
											f(p...)
											ok = true

	case func(R.Value):						for i := 0; i < end; i++ {
												f(R.ValueOf(container.At(i)))
											}
											ok = true

	case func(int, R.Value):				for i := 0; i < end; i++ {
												f(i, R.ValueOf(container.At(i)))
											}
											ok = true

	case func(interface{}, R.Value):		for i := 0; i < end; i++ {
												f(i, R.ValueOf(container.At(i)))
											}
											ok = true

	case func(R.Value, R.Value):			for i := 0; i < end; i++ {
												f(R.ValueOf(i), R.ValueOf(container.At(i)))
											}
											ok = true


	case func(...R.Value):					p := make([]R.Value, end, end)
											for i := 0; i < end; i++ {
												p[i] = R.ValueOf(container.At(i))
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, end, end)
													for i := 0; i < end; i++ {
														p[i] = R.ValueOf(container.At(i))
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for i := 0; i < end; i++ {
																			p[0] = R.ValueOf(container.At(i))
																			f.Call(p)
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		for i := 0; i < end; i++ {
																			p[0], p[1] = R.ValueOf(i), R.ValueOf(container.At(i))
																			f.Call(p)
																		}
																		ok = true
													}
												}
											}
	}
	return
}

func eachMappable(container Mappable, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					Each(container.Keys(), func(v interface{}) {
												f(container.Lookup(v))
												
											})
											ok = true

	case func(int, interface{}):			Each(container.Keys(), func(v interface{}) {
												f(v.(int), container.Lookup(v))
											})
											ok = true

	case func(string, interface{}):			Each(container.Keys(), func(v interface{}) {
												f(v.(string), container.Lookup(v))
											})
											ok = true

	case func(interface{}, interface{}):	Each(container.Keys(), func(v interface{}) {
												f(v, container.Lookup(v))
											})
											ok = true

	case func(...interface{}):				l := Len(container)
											p := make([]interface{}, l, l)
											Each(container.Keys(), func(i int, v interface{}) {
												p[i] = container.Lookup(v)
											})
											f(p...)
											ok = true

	case func(R.Value):						Each(container.Keys(), func(v interface{}) {
												f(R.ValueOf(container.Lookup(v)))
											})
											ok = true

	case func(int, R.Value):				Each(container.Keys(), func(v interface{}) {
												f(v.(int), R.ValueOf(container.Lookup(v)))
											})
											ok = true

	case func(string, R.Value):				Each(container.Keys(), func(v interface{}) {
												f(v.(string), R.ValueOf(container.Lookup(v)))
											})
											ok = true

	case func(interface{}, R.Value):		Each(container.Keys(), func(v interface{}) {
												f(v, R.ValueOf(container.Lookup(v)))
											})
											ok = true

	case func(R.Value, R.Value):			Each(container.Keys(), func(v interface{}) {
												f(R.ValueOf(v), R.ValueOf(container.Lookup(v)))
											})
											ok = true

	case func(...R.Value):					l := Len(container)
											p := make([]R.Value, l, l)
											Each(container.Keys(), func(i int, v interface{}) {
												p[i] = R.ValueOf(container.Lookup(v))
											})
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													// f(...v)
													l := Len(container)
													p := make([]R.Value, l, l)
													Each(container.Keys(), func(i int, v interface{}) {
														p[i] = R.ValueOf(container.Lookup(v))
													})
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		Each(container.Keys(), func(v interface{}) {
																			p[0] = R.ValueOf(container.Lookup(v))
																			f.Call(p)																			
																		})
																		ok = true

													case 2:				// f(i, v)
																		p := make([]R.Value, 2, 2)
																		Each(container.Keys(), func(v interface{}) {
																			p[0], p[1] = R.ValueOf(v), R.ValueOf(container.Lookup(v))
																			f.Call(p)
																		})
																		ok = true
													}
												}
											}
	}
	return
}

func eachSlice(s R.Value, f interface{}) (ok bool) {
	end := Len(s)
	switch f := f.(type) {
	case func(interface{}):					for i := 0; i < end; i++ {
												f(s.Index(i).Interface())
											}
											ok = true

	case func(int, interface{}):			for i := 0; i < end; i++ {
												f(i, s.Index(i).Interface())
											}
											ok = true

	case func(interface{}, interface{}):	for i := 0; i < end; i++ {
												f(i, s.Index(i).Interface())
											}
											ok = true

	case func(...interface{}):				p := make([]interface{}, end, end)
											for i := 0; i < end; i++ {
												p[i] = s.Index(i).Interface()
											}
											f(p...)
											ok = true

	case func(R.Value):						for i := 0; i < end; i++ {
												f(s.Index(i))
											}
											ok = true

	case func(int, R.Value):				for i := 0; i < end; i++ {
												f(i, s.Index(i))
											}
											ok = true

	case func(interface{}, R.Value):		for i := 0; i < end; i++ {
												f(i, s.Index(i))
											}
											ok = true

	case func(R.Value, R.Value):			for i := 0; i < end; i++ {
												f(R.ValueOf(i), s.Index(i))
											}
											ok = true

	case func(...R.Value):					p := make([]R.Value, end, end)
											for i := 0; i < end; i++ {
												p[i] = s.Index(i)
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												end := Len(s)
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, end, end)
													for i := 0; i < end; i++ {
														p[i] = s.Index(i)
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for i := 0; i < end; i++ {
																			p[0] = s.Index(i)
																			f.Call(p)
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		for i := 0; i < end; i++ {
																			p[0], p[1] = R.ValueOf(i), s.Index(i)
																			f.Call(p)
																		}
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

func eachChannel(c R.Value, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					for v, open := c.Recv(); open; {
												f(v.Interface())
												v, open = c.Recv()
											}
											ok = true

	case func(int, interface{}):			i := 0
											for v, open := c.Recv(); open; i++ {
												f(i, v.Interface())
												v, open = c.Recv()
											}
											ok = true

	case func(interface{}, interface{}):	i := 0
											for v, open := c.Recv(); open; i++ {
												f(i, v.Interface())
												v, open = c.Recv()
											}
											ok = true

	case func(...interface{}):				p := make([]interface{}, 0, 4)
											for v, open := c.Recv(); open; {
												p = append(p, v.Interface())
												v, open = c.Recv()
											}
											f(p...)
											ok = true

	case func(R.Value):						for v, open := c.Recv(); open; {
												f(v)
												v, open = c.Recv()
											}
											ok = true

	case func(int, R.Value):				i := 0
											for v, open := c.Recv(); open; i++ {
												f(i, v)
												v, open = c.Recv()
											}
											ok = true

	case func(interface{}, R.Value):		i := 0
											for v, open := c.Recv(); open; i++ {
												f(i, v)
												v, open = c.Recv()
											}
											ok = true

	case func(R.Value, R.Value):			i := 0
											for v, open := c.Recv(); open; i++ {
												f(R.ValueOf(i), v)
												v, open = c.Recv()
											}
											ok = true

	case func(...R.Value):					p := make([]R.Value, 0, 4)
											for v, open := c.Recv(); open; {
												p = append(p, v)
												v, open = c.Recv()
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, 0, 4)
													for v, open := c.Recv(); open; {
														p = append(p, v)
														v, open = c.Recv()
													}
													f.Call(p)
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for v, open := c.Recv(); open; {
																			p[0] = v
																			f.Call(p)
																			v, open = c.Recv()
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		i := 0
																		for v, open := c.Recv(); open; i++ {
																			p[0], p[1] = R.ValueOf(i), v
																			f.Call(p)
																			v, open = c.Recv()
																		}
																		ok = true
													}
												}
											}
	}
	return
}

func eachFunction(g R.Value, f interface{}) (ok bool) {
	if tg := g.Type(); tg.NumOut() == 2 {
		switch tg.NumIn() {
		case 0:			switch f := f.(type) {
						case func(interface{}):					p := []R.Value{}
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0].Interface())
																}
																ok = true

						case func(int, interface{}):			p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																} 
																ok = true

						case func(interface{}, interface{}):	p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																}
																ok = true

						case func(...interface{}):				pg := []R.Value{}
																pf := make([]interface{}, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0].Interface())
																}
																f(pf...)
																ok = true

						case func(R.Value):						p := []R.Value{}
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0])
																}
																ok = true

						case func(int, R.Value):				p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																}
																ok = true

						case func(interface{}, R.Value):		p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																}
																ok = true

						case func(R.Value, R.Value):			p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(R.ValueOf(i), v[0])
																	i++
																}
																ok = true

						case func(...R.Value):					pg := []R.Value{}
																pf := make([]R.Value, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0])
																}
																f(pf...)
																ok = true

						default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																	if tf := f.Type(); tf.IsVariadic() {
																		//	f(...v)
																		i := 0
																		pg := []R.Value{}
																		pf := make([]R.Value, 0, 4)
																		for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																			pf = append(pf, v[0])
																			i++
																		}
																		f.Call(pf)
																		ok = true
																	} else {
																		switch tf.NumIn() {
																		case 1:		//	f(v)
																					pg := []R.Value{}
																					pf := make([]R.Value, 1, 1)
																					for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																						pf[0] = v[0]
																						f.Call(pf)
																					}
																					ok = true

																		case 2:		//	f(i, v)
																					i := 0
																					pg := []R.Value{}
																					pf := make([]R.Value, 2, 2)
																					for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																						pf[0], pf[1] = R.ValueOf(i), v[0]
																						f.Call(pf)
																						i++
																					}
																					ok = true
																		}
																	}
																}
						}

		case 1:			switch f := f.(type) {
						case func(interface{}):					i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(int, interface{}):			i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(interface{}, interface{}):	i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(...interface{}):				i := 0
																pg := []R.Value{ R.ValueOf(0) }
																pf := make([]interface{}, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0].Interface())
																	i++
																	pg[i] = R.ValueOf(i)
																}
																f(pf...)
																ok = true

						case func(R.Value):						i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(int, R.Value):				i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(interface{}, R.Value):		i := 0
																p := []R.Value{ R.ValueOf(i) }
																for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true


						case func(R.Value, R.Value):			i := 0
																p := []R.Value{ R.ValueOf(i) }
																for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																	f(p[0], v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(...R.Value):					i := 0
																pg := []R.Value{ R.ValueOf(i) }
																pf := make([]R.Value, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0])
																	i++
																	pf[i] = R.ValueOf(i)
																}
																f(pf...)
																ok = true

						default:								if f := R.ValueOf(f); f.Kind() == R.Func {
																	if tf := f.Type(); tf.IsVariadic() {
																		//	f(...v)
																		i := 0
																		pg := []R.Value{ R.ValueOf(0) }
																		pf := make([]R.Value, 0, 4)
																		for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																			pf = append(pf, v[0])
																			i++
																			pg[0] = R.ValueOf(i)
																		}
																		f.Call(pf)
																		ok = true
																	} else {
																		switch tf.NumIn() {
																		case 1:		//	f(v)
																					i := 0
																					p := []R.Value{ R.ValueOf(0) }
																					for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																						p[0] = v[0]
																						f.Call(p)
																						i++
																						p[0] = R.ValueOf(i)
																					}
																					ok = true

																		case 2:		//	f(i, v)
																					i := 0
																					pg := []R.Value{ R.ValueOf(0) }
																					pf := make([]R.Value, 2, 2)
																					for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																						pf[0], pf[1] = pg[0], v[0]
																						f.Call(pf)
																						i++
																						pg[0] = R.ValueOf(i)
																					}
																					ok = true
																		}
																	}
																}
						}
		}
	}
	return
}