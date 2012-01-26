package sequences

import R "reflect"

func Len(container interface{}) (l int) {
	switch container := container.(type) {
	case nil:				l = 0
	case []bool:			l = len(container)
	case []int:				l = len(container)
	case []int8:			l = len(container)
	case []int16:			l = len(container)
	case []int32:			l = len(container)
	case []int64:			l = len(container)
	case []uint:			l = len(container)
	case []uint8:			l = len(container)
	case []uint16:			l = len(container)
	case []uint32:			l = len(container)
	case []uint64:			l = len(container)
	case []uintptr:			l = len(container)
	case []complex64:		l = len(container)
	case []complex128:		l = len(container)
	case []float32:			l = len(container)
	case []float64:			l = len(container)
	case []string:			l = len(container)
	case []interface{}:		l = len(container)
	case R.Value:			switch container.Kind() {
							case R.Slice:	fallthrough
							case R.Map:		l = container.Len()
							case R.Invalid:	l = 0
							default:		l = 1
							}
	case Measurable:		l = container.Len()
	default:				switch container := R.ValueOf(container); container.Kind() {
							case R.Slice:	fallthrough
							case R.Map:		l = container.Len()
							case R.Invalid:	l = 0
							default:		l = 1
							}
	}
	return
}

func Cap(container interface{}) (c int) {
	switch container := container.(type) {
	case nil:				c = 0
	case []bool:			c = cap(container)
	case []int:				c = cap(container)
	case []int8:			c = cap(container)
	case []int16:			c = cap(container)
	case []int32:			c = cap(container)
	case []int64:			c = cap(container)
	case []uint:			c = cap(container)
	case []uint8:			c = cap(container)
	case []uint16:			c = cap(container)
	case []uint32:			c = cap(container)
	case []uint64:			c = cap(container)
	case []uintptr:			c = cap(container)
	case []complex64:		c = cap(container)
	case []complex128:		c = cap(container)
	case []float32:			c = cap(container)
	case []float64:			c = cap(container)
	case []string:			c = cap(container)
	case []interface{}:		c = cap(container)
	case R.Value:			switch container.Kind() {
							case R.Slice:	c = container.Cap()
							case R.Map:		c = -1
							case R.Invalid:	c = 0
							default:		c = 1
							}
	case Confined:			c = container.Cap()
	default:				switch container := R.ValueOf(container); container.Kind() {
							case R.Slice:	c = container.Cap()
							case R.Map:		c = -1
							case R.Invalid:	c = 0
							default:		c = 1
							}
	}
	return
}