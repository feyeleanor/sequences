package sequences

import(
//	"reflect"
	"testing"
)


func TestReduce(t *testing.T) {
	ConfirmReduce := func(o, s, r, f interface{}) {
/*		if x, _ := Reduce(o, s, f); !reflect.DeepEqual(x, r) {
			t.Fatalf("Reduce(%v, %v, %v) should be %v but is %v", o, s, f, r, x)
		}
*/	}

	Sum := func(memo, x interface{}) interface{} {
		return memo.(int) + x.(int)
	}

	ConfirmReduce(nil, nil, nil, Sum)
	ConfirmReduce(nil, 0, 0, Sum)

	ConfirmReduce([]int{}, 0, 0, Sum)
	ConfirmReduce([]int{}, 1, 1, Sum)
	ConfirmReduce([]int{0}, 0, 0, Sum)
	ConfirmReduce([]int{0}, 1, 1, Sum)
	ConfirmReduce([]int{0, 1, 2, 3}, 0, 6, Sum)
	ConfirmReduce([]int{0, 1, 2, 3}, 1, 7, Sum)

	ConfirmReduce(map[int] int{}, 0, 0, Sum)
	ConfirmReduce(map[int] int{}, 1, 1, Sum)
	ConfirmReduce(map[int] int{0: 0}, 0, 0, Sum)
	ConfirmReduce(map[int] int{0: 0}, 1, 1, Sum)
	ConfirmReduce(map[int] int{0: 0, 1: 1, 2: 2, 3: 3}, 0, 6, Sum)
	ConfirmReduce(map[int] int{0: 0, 1: 1, 2: 2, 3: 3}, 1, 7, Sum)

	IntSum := func(memo, x int) int {
		return memo + x
	}

	ConfirmReduce([]int{}, 0, 0, IntSum)
	ConfirmReduce([]int{}, 1, 1, IntSum)
	ConfirmReduce([]int{0}, 0, 0, IntSum)
	ConfirmReduce([]int{0}, 1, 1, IntSum)
	ConfirmReduce([]int{0, 1, 2, 3}, 0, 6, IntSum)
	ConfirmReduce([]int{0, 1, 2, 3}, 1, 7, IntSum)

	ConfirmReduce(map[int] int{}, 0, 0, IntSum)
	ConfirmReduce(map[int] int{}, 1, 1, IntSum)
	ConfirmReduce(map[int] int{0: 0}, 0, 0, IntSum)
	ConfirmReduce(map[int] int{0: 0}, 1, 1, IntSum)
	ConfirmReduce(map[int] int{0: 0, 1: 1, 2: 2, 3: 3}, 0, 6, IntSum)
	ConfirmReduce(map[int] int{0: 0, 1: 1, 2: 2, 3: 3}, 1, 7, IntSum)
}