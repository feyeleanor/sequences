package sequences

import (
	R "reflect"
	"testing"
)

func TestValuesAtOffsets(t *testing.T) {
	ConfirmValuesAtOffsets := func(c interface{}, o []int, r interface{}) {
		if x := ValuesAtOffsets(c, o...); !R.DeepEqual(x, r) {
			t.Fatalf("ValuesAtOffsets(%v, %v) should be %v but is %v", c, o, r, x)
		}
	}

	ConfirmValuesAtOffsets([]complex128{0, 1, 2, 3, 4, 5}, []int{}, []complex128{})
	ConfirmValuesAtOffsets([]complex128{0, 1, 2, 3, 4, 5}, []int{0, 1}, []complex128{0, 1})
	ConfirmValuesAtOffsets([]complex128{0, 1, 2, 3, 4, 5}, []int{0, 3}, []complex128{0, 3})
	ConfirmValuesAtOffsets([]complex128{0, 1, 2, 3, 4, 5}, []int{0, 3, 4, 3}, []complex128{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]complex64{0, 1, 2, 3, 4, 5}, []int{}, []complex64{})
	ConfirmValuesAtOffsets([]complex64{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []complex64{0, 1})
	ConfirmValuesAtOffsets([]complex64{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []complex64{0, 3})
	ConfirmValuesAtOffsets([]complex64{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []complex64{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]float32{0, 1, 2, 3, 4, 5}, []int{}, []float32{})
	ConfirmValuesAtOffsets([]float32{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []float32{0, 1})
	ConfirmValuesAtOffsets([]float32{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []float32{0, 3})
	ConfirmValuesAtOffsets([]float32{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []float32{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]float64{0, 1, 2, 3, 4, 5}, []int{}, []float64{})
	ConfirmValuesAtOffsets([]float64{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []float64{0, 1})
	ConfirmValuesAtOffsets([]float64{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []float64{0, 3})
	ConfirmValuesAtOffsets([]float64{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []float64{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]int16{0, 1, 2, 3, 4, 5}, []int{}, []int16{})
	ConfirmValuesAtOffsets([]int16{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []int16{0, 1})
	ConfirmValuesAtOffsets([]int16{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []int16{0, 3})
	ConfirmValuesAtOffsets([]int16{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []int16{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]int32{0, 1, 2, 3, 4, 5}, []int{}, []int32{})
	ConfirmValuesAtOffsets([]int32{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []int32{0, 1})
	ConfirmValuesAtOffsets([]int32{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []int32{0, 3})
	ConfirmValuesAtOffsets([]int32{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []int32{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]int64{0, 1, 2, 3, 4, 5}, []int{}, []int64{})
	ConfirmValuesAtOffsets([]int64{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []int64{0, 1})
	ConfirmValuesAtOffsets([]int64{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []int64{0, 3})
	ConfirmValuesAtOffsets([]int64{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []int64{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]int8{0, 1, 2, 3, 4, 5}, []int{}, []int8{})
	ConfirmValuesAtOffsets([]int8{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []int8{0, 1})
	ConfirmValuesAtOffsets([]int8{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []int8{0, 3})
	ConfirmValuesAtOffsets([]int8{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []int8{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]int{0, 1, 2, 3, 4, 5}, []int{}, []int{})
	ConfirmValuesAtOffsets([]int{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []int{0, 1})
	ConfirmValuesAtOffsets([]int{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []int{0, 3})
	ConfirmValuesAtOffsets([]int{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []int{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]uint16{0, 1, 2, 3, 4, 5}, []int{}, []uint16{})
	ConfirmValuesAtOffsets([]uint16{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []uint16{0, 1})
	ConfirmValuesAtOffsets([]uint16{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []uint16{0, 3})
	ConfirmValuesAtOffsets([]uint16{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []uint16{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]uint32{0, 1, 2, 3, 4, 5}, []int{}, []uint32{})
	ConfirmValuesAtOffsets([]uint32{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []uint32{0, 1})
	ConfirmValuesAtOffsets([]uint32{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []uint32{0, 3})
	ConfirmValuesAtOffsets([]uint32{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []uint32{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]uint64{0, 1, 2, 3, 4, 5}, []int{}, []uint64{})
	ConfirmValuesAtOffsets([]uint64{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []uint64{0, 1})
	ConfirmValuesAtOffsets([]uint64{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []uint64{0, 3})
	ConfirmValuesAtOffsets([]uint64{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []uint64{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]uint8{0, 1, 2, 3, 4, 5}, []int{}, []uint8{})
	ConfirmValuesAtOffsets([]uint8{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []uint8{0, 1})
	ConfirmValuesAtOffsets([]uint8{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []uint8{0, 3})
	ConfirmValuesAtOffsets([]uint8{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []uint8{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]uint{0, 1, 2, 3, 4, 5}, []int{}, []uint{})
	ConfirmValuesAtOffsets([]uint{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []uint{0, 1})
	ConfirmValuesAtOffsets([]uint{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []uint{0, 3})
	ConfirmValuesAtOffsets([]uint{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []uint{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]uintptr{0, 1, 2, 3, 4, 5}, []int{}, []uintptr{})
	ConfirmValuesAtOffsets([]uintptr{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []uintptr{0, 1})
	ConfirmValuesAtOffsets([]uintptr{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []uintptr{0, 3})
	ConfirmValuesAtOffsets([]uintptr{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []uintptr{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]interface{}{0, 1, 2, 3, 4, 5}, []int{}, []interface{}{})
	ConfirmValuesAtOffsets([]interface{}{0, 1, 2, 3, 4, 5}, []int{ 0, 1 }, []interface{}{0, 1})
	ConfirmValuesAtOffsets([]interface{}{0, 1, 2, 3, 4, 5}, []int{ 0, 3 }, []interface{}{0, 3})
	ConfirmValuesAtOffsets([]interface{}{0, 1, 2, 3, 4, 5}, []int{ 0, 3, 4, 3 }, []interface{}{0, 3, 4, 3})

	ConfirmValuesAtOffsets([]string{"A", "B", "A", "B", "A"}, []int{}, []string{})
	ConfirmValuesAtOffsets([]string{"A", "B", "A", "B", "A"}, []int{ 0, 1 }, []string{"A", "B"})
	ConfirmValuesAtOffsets([]string{"A", "B", "A", "B", "A"}, []int{ 0, 3 }, []string{"A", "B"})
	ConfirmValuesAtOffsets([]string{"A", "B", "A", "B", "A"}, []int{ 0, 3, 4, 3 }, []string{"A", "B", "A", "B"})

	ConfirmValuesAtOffsets([]rune{'A', 'B', 'A', 'B', 'A'}, []int{}, []rune{})
	ConfirmValuesAtOffsets([]rune{'A', 'B', 'A', 'B', 'A'}, []int{ 0, 1 }, []rune{'A', 'B'})
	ConfirmValuesAtOffsets([]rune{'A', 'B', 'A', 'B', 'A'}, []int{ 0, 3 }, []rune{'A', 'B'})
	ConfirmValuesAtOffsets([]rune{'A', 'B', 'A', 'B', 'A'}, []int{ 0, 3, 4, 3 }, []rune{'A', 'B', 'A', 'B'})
}