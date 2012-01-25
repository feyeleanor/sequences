package gosequence

import R "reflect"
import "testing"

func BenchmarkEachIndexable1(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable2(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i int, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable3(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable4(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v ...interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable5(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable6(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i int, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable7(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i interface{}, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable8(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable9(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v ...R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable10(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v ...int) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable11(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v int) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachIndexable12(b *testing.B) {
	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i, v int) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}






func BenchmarkEachSlice1(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice2(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i int, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice3(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice4(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v ...interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice5(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice6(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i int, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice7(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i interface{}, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice8(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice9(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v ...R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice10(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v ...int) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice11(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(v int) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}

func BenchmarkEachSlice12(b *testing.B) {
	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	F := func(i, v int) {}
	for i := 0; i < b.N; i++ {
		Each(S, F)
	}
}
