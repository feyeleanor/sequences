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


func BenchmarkEachMappable1(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable2(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i int, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable3(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable4(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v ...interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable5(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable6(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i int, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable7(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i interface{}, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable8(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable9(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v ...R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable10(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v ...int) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable11(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v int) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable12(b *testing.B) {
	M := mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i, v int) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable13(b *testing.B) {
	M := mappable_string_map{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	F := func(s string, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMappable14(b *testing.B) {
	M := mappable_string_map{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	F := func(s string, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
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

func BenchmarkEachMap1(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap2(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i int, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap3(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap4(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v ...interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap5(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap6(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i int, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap7(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i interface{}, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap8(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap9(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v ...R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap10(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v ...int) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap11(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(v int) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap12(b *testing.B) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	F := func(i, v int) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap13(b *testing.B) {
	M := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	F := func(s string, v interface{}) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}

func BenchmarkEachMap14(b *testing.B) {
	M := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	F := func(s string, v R.Value) {}
	for i := 0; i < b.N; i++ {
		Each(M, F)
	}
}