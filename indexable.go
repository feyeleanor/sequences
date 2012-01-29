package sequences

import R "reflect"

type Indexable interface {
	AtOffset(index int) interface{}
}

func AtOffset(container interface{}, index int) (r interface{}) {
	switch container := container.(type) {
	case Indexable:				r = container.AtOffset(index)
	case []bool:				r = container[index]
	case []interface{}:			r = container[index]
	case []int:					r = container[index]
	case []int8:				r = container[index]
	case []int16:				r = container[index]
	case []int32:				r = container[index]
	case []int64:				r = container[index]
	case []uint:				r = container[index]
	case []uint8:				r = container[index]
	case []uint16:				r = container[index]
	case []uint32:				r = container[index]
	case []uint64:				r = container[index]
	case []uintptr:				r = container[index]
	case []float32:				r = container[index]
	case []float64:				r = container[index]
	case []complex64:			r = container[index]
	case []complex128:			r = container[index]
	case []string:				r = container[index]
	case []error:				r = container[index]
	case []R.Value:				r = container[index]
	case R.Value:				switch container.Kind() {
								case R.Slice:		r = container.Index(index).Interface()
								case R.Map:			r = container.MapIndex(R.ValueOf(index)).Interface()
								case R.Func:		r = container.Call([]R.Value{R.ValueOf(index)})[0]
								}
	default:					switch v := R.ValueOf(container); v.Kind() {
								case R.Slice:		r = v.Index(index).Interface()
								case R.Map:			r = v.MapIndex(R.ValueOf(index)).Interface()
								case R.Func:		r = v.Call([]R.Value{R.ValueOf(index)})[0]
								}
	}
	return
}