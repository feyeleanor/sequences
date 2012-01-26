package gosequence

import (
	R "reflect"
	"testing"
)

func BenchmarkLen1(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		_ = len(S)
	}
}

func BenchmarkLen2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Len(0)
	}
}

func BenchmarkLen3(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Len(S)
	}
}

func BenchmarkLen4(b *testing.B) {
	S := R.ValueOf(0)
	for i := 0; i < b.N; i++ {
		Len(S)
	}
}

func BenchmarkLen5(b *testing.B) {
	S := R.ValueOf([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for i := 0; i < b.N; i++ {
		Len(S)
	}
}

func BenchmarkLen6(b *testing.B) {
	S := mappable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Len(S)
	}
}

func BenchmarkLen7(b *testing.B) {
	M := R.ValueOf(map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9})
	for i := 0; i < b.N; i++ {
		Len(M)
	}
}

func BenchmarkLen8(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	for i := 0; i < b.N; i++ {
		Len(M)
	}
}

func BenchmarkCap1(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		_ = cap(S)
	}
}

func BenchmarkCap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cap(0)
	}
}

func BenchmarkCap3(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Cap(S)
	}
}

func BenchmarkCap4(b *testing.B) {
	S := R.ValueOf(0)
	for i := 0; i < b.N; i++ {
		Cap(S)
	}
}

func BenchmarkCap5(b *testing.B) {
	S := R.ValueOf([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for i := 0; i < b.N; i++ {
		Cap(S)
	}
}

func BenchmarkCap6(b *testing.B) {
	S := mappable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Cap(S)
	}
}

func BenchmarkCap7(b *testing.B) {
	M := R.ValueOf(map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9})
	for i := 0; i < b.N; i++ {
		Cap(M)
	}
}

func BenchmarkCap8(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	for i := 0; i < b.N; i++ {
		Cap(M)
	}
}