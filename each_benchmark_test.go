package sequences

import R "reflect"
import "testing"

var (
	IS   = indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	SI   = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	SUDT = []UDT{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	MII  = map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	CI   = func() (c chan int) {
		c = make(chan int)
		go func() {
			for _, v := range SI {
				c <- v
			}
			close(c)
		}()
		return
	}
	FI = func(i int) (r int, finished bool) {
		r = i
		if i == 10 {
			finished = true
		}
		return
	}

	F1   = func(v interface{}) {}
	F2   = func(i int, v interface{}) {}
	F3   = func(i, v interface{}) {}
	F5   = func(v R.Value) {}
	F6   = func(i int, v R.Value) {}
	F7   = func(i interface{}, v R.Value) {}
	F8   = func(i, v R.Value) {}
	F11  = func(v int) {}
	F11u = func(v UDT) {}
	F12  = func(i, v int) {}
	F12u = func(i int, v UDT) {}
	F13  = func(s string, v interface{}) {}
	F14  = func(s string, v R.Value) {}
)

func BenchmarkRangeSliceF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range SI {
			F1(v)
		}
	}
}

func BenchmarkRangeSliceF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i, v := range SI {
			F2(i, v)
		}
	}
}

func BenchmarkRangeSliceF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i, v := range SI {
			F3(i, v)
		}
	}
}

func BenchmarkRangeSliceF5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range SI {
			F5(R.ValueOf(v))
		}
	}
}

func BenchmarkRangeSliceF6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i, v := range SI {
			F6(i, R.ValueOf(v))
		}
	}
}

func BenchmarkRangeSliceF7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i, v := range SI {
			F7(i, R.ValueOf(v))
		}
	}
}

func BenchmarkRangeSliceF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i, v := range SI {
			F8(R.ValueOf(i), R.ValueOf(v))
		}
	}
}

func BenchmarkRangeSliceF11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range SI {
			F11(v)
		}
	}
}

func BenchmarkRangeSliceF12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i, v := range SI {
			F12(i, v)
		}
	}
}

func BenchmarkEachBuiltingF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F1)
	}
}

func BenchmarkEachBuiltingF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F2)
	}
}

func BenchmarkEachBuiltingF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F3)
	}
}

func BenchmarkEachBuiltingF5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F5)
	}
}

func BenchmarkEachBuiltingF6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F6)
	}
}

func BenchmarkEachBuiltingF7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F7)
	}
}

func BenchmarkEachBuiltingF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F8)
	}
}

func BenchmarkEachBuiltingF11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F11)
	}
}

func BenchmarkEachBuiltingF12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F12)
	}
}

func BenchmarkEachIndexableF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F1)
	}
}

func BenchmarkEachIndexableF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F2)
	}
}

func BenchmarkEachIndexableF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F3)
	}
}

func BenchmarkEachIndexableF5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F5)
	}
}

func BenchmarkEachIndexableF6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F6)
	}
}

func BenchmarkEachIndexableF7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F7)
	}
}

func BenchmarkEachIndexableF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F8)
	}
}

func BenchmarkEachIndexableF11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F11)
	}
}

func BenchmarkEachIndexableF12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F12)
	}
}

func BenchmarkEachSliceF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F1)
	}
}

func BenchmarkEachSliceF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F2)
	}
}

func BenchmarkEachSliceF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F3)
	}
}

func BenchmarkEachSliceF5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F5)
	}
}

func BenchmarkEachSliceF6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F6)
	}
}

func BenchmarkEachSliceF7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F7)
	}
}

func BenchmarkEachSliceF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F8)
	}
}

func BenchmarkEachSliceF11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F11u)
	}
}

func BenchmarkEachSliceF12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F12u)
	}
}

func BenchmarkEachMapF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F1)
	}
}

func BenchmarkEachMapF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F2)
	}
}

func BenchmarkEachMapF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F3)
	}
}

func BenchmarkEachMapF5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F5)
	}
}

func BenchmarkEachMapF6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F6)
	}
}

func BenchmarkEachMapF7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F7)
	}
}

func BenchmarkEachMapF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F8)
	}
}

func BenchmarkEachMapF11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F11)
	}
}

func BenchmarkEachMapF12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F12)
	}
}

func BenchmarkEachChannelF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F1)
	}
}

func BenchmarkEachChannelF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F2)
	}
}

func BenchmarkEachChannelF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F3)
	}
}

func BenchmarkEachChannelF5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F5)
	}
}

func BenchmarkEachChannelF6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F6)
	}
}

func BenchmarkEachChannelF7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F7)
	}
}

func BenchmarkEachChannelF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F8)
	}
}

func BenchmarkEachChannelF11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F11)
	}
}

func BenchmarkEachChannelF12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F12)
	}
}

func BenchmarkEachFunctionF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F1)
	}
}

func BenchmarkEachFunctionF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F2)
	}
}

func BenchmarkEachFunctionF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F3)
	}
}

func BenchmarkEachFunctionF5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F5)
	}
}

func BenchmarkEachFunctionF6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F6)
	}
}

func BenchmarkEachFunctionF7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F7)
	}
}

func BenchmarkEachFunctionF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F8)
	}
}

func BenchmarkEachFunctionF11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F11)
	}
}

func BenchmarkEachFunctionF12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F12)
	}
}
