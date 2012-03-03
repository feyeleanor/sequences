package sequences

import R "reflect"

type Indexable interface {
	AtOffset(index int) interface{}
}

func AtOffset(container interface{}, index int) (r interface{}) {
	switch container := container.(type) {
	case Indexable:
		r = container.AtOffset(index)
	case []bool:
		r = container[index]
	case []interface{}:
		r = container[index]
	case []int:
		r = container[index]
	case []int8:
		r = container[index]
	case []int16:
		r = container[index]
	case []int32:
		r = container[index]
	case []int64:
		r = container[index]
	case []uint:
		r = container[index]
	case []uint8:
		r = container[index]
	case []uint16:
		r = container[index]
	case []uint32:
		r = container[index]
	case []uint64:
		r = container[index]
	case []uintptr:
		r = container[index]
	case []float32:
		r = container[index]
	case []float64:
		r = container[index]
	case []complex64:
		r = container[index]
	case []complex128:
		r = container[index]
	case []string:
		r = container[index]
	case []error:
		r = container[index]
	case []R.Value:
		r = container[index]
	case R.Value:
		switch container.Kind() {
		case R.Slice:
			r = container.Index(index).Interface()
		case R.Map:
			r = container.MapIndex(R.ValueOf(index)).Interface()
		case R.Func:
			r = container.Call([]R.Value{R.ValueOf(index)})[0]
		}
	default:
		switch v := R.ValueOf(container); v.Kind() {
		case R.Slice:
			r = v.Index(index).Interface()
		case R.Map:
			r = v.MapIndex(R.ValueOf(index)).Interface()
		case R.Func:
			r = v.Call([]R.Value{R.ValueOf(index)})[0]
		}
	}
	return
}

func ValuesAtOffsets(container interface{}, offsets ...int) (r interface{}) {
	switch container := container.(type) {
	case Indexable:
		n := make([]interface{}, len(offsets))
		for i, _ := range n {
			n[i] = container.AtOffset(offsets[i])
		}
		r = n
	case []bool:
		n := make([]bool, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []interface{}:
		n := make([]interface{}, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []int:
		n := make([]int, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []int8:
		n := make([]int8, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []int16:
		n := make([]int16, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []int32:
		n := make([]int32, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []int64:
		n := make([]int64, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []uint:
		n := make([]uint, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []uint8:
		n := make([]uint8, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []uint16:
		n := make([]uint16, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []uint32:
		n := make([]uint32, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []uint64:
		n := make([]uint64, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []uintptr:
		n := make([]uintptr, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []float32:
		n := make([]float32, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []float64:
		n := make([]float64, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []complex64:
		n := make([]complex64, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []complex128:
		n := make([]complex128, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []string:
		n := make([]string, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []error:
		n := make([]error, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case []R.Value:
		n := make([]R.Value, len(offsets))
		for i, _ := range n {
			n[i] = container[offsets[i]]
		}
		r = n
	case R.Value:
		switch container.Kind() {
		case R.Slice:
			panic("implement for arbitrary slices")
		case R.Map:
			panic("implement for arbitrary maps")
		case R.Func:
			panic("implement for arbitrary functions")
		case R.Chan:
			panic("Channels not yet supported")
		default:
			panic("this object cannot be indexed")
		}
	default:
		switch container := R.ValueOf(container); container.Kind() {
		case R.Slice:
			panic("implement for arbitrary slices")
		case R.Map:
			panic("implement for arbitrary maps")
		case R.Func:
			panic("implement for arbitrary functions")
		case R.Chan:
			panic("Channels not yet supported")
		default:
			panic("this object cannot be indexed")
		}
	}
	return
}