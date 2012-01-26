package sequences

import R "reflect"

func Len(container interface{}) (l int) {
	switch container := container.(type) {
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
	case []complex64:		l = len(container)
	case []complex128:		l = len(container)
	case []float32:			l = len(container)
	case []float64:			l = len(container)
	case []string:			l = len(container)
	case []interface{}:		l = len(container)
	case R.Value:			switch container.Kind() {
							case R.Slice:	fallthrough
							case R.Map:		l = container.Len()
							default:		l = 1
							}
	case Measurable:		l = container.Len()
	default:				switch container := R.ValueOf(container); container.Kind() {
							case R.Slice:	fallthrough
							case R.Map:		l = container.Len()
							default:		l = 1
							}
	}
	return
}

func Cap(container interface{}) (l int) {
	switch container := container.(type) {
	case []int:				l = cap(container)
	case []int8:			l = cap(container)
	case []int16:			l = cap(container)
	case []int32:			l = cap(container)
	case []int64:			l = cap(container)
	case []uint:			l = cap(container)
	case []uint8:			l = cap(container)
	case []uint16:			l = cap(container)
	case []uint32:			l = cap(container)
	case []uint64:			l = cap(container)
	case []complex64:		l = cap(container)
	case []complex128:		l = cap(container)
	case []float32:			l = cap(container)
	case []float64:			l = cap(container)
	case []string:			l = cap(container)
	case []interface{}:		l = cap(container)
	case R.Value:			switch container.Kind() {
							case R.Slice:	l = container.Cap()
							case R.Map:		l = -1
							default:		l = 1
							}
	case Confined:			l = container.Cap()
	default:				switch container := R.ValueOf(container); container.Kind() {
							case R.Slice:	l = container.Cap()
							case R.Map:		l = -1
							default:		l = 1
							}
	}
	return
}