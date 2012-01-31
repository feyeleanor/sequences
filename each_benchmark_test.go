package sequences

import R "reflect"
import "testing"

var (
	IS = indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	MM = mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	MSM = mappable_string_map{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	SI = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	SUDT = []UDT{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	MII = map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	MSI = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	CI = func() (c chan int) {
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

	F1 = func(v interface{}) {}
	F2 = func(i int, v interface{}) {}
	F3 = func(i, v interface{}) {}
	F4 = func(v ...interface{}) {}
	F5 = func(v R.Value) {}
	F6 = func(i int, v R.Value) {}
	F7 = func(i interface{}, v R.Value) {}
	F8 = func(i, v R.Value) {}
	F9 = func(v ...R.Value) {}
	F10 = func(v ...int) {}
	F10u = func(v ...UDT) {}
	F11 = func(v int) {}
	F11u = func(v UDT) {}
	F12 = func(i, v int) {}
	F12u = func(i int, v UDT) {}
	F13 = func(s string, v interface{}) {}
	F14 = func(s string, v R.Value) {}
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

func BenchmarkRangeSliceF4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := []interface{}{}
		for _, v := range SI {
			p = append(p, v)
		}
		F4(p...)
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

func BenchmarkRangeSliceF9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := []R.Value{}
		for _, v := range SI {
			p = append(p, R.ValueOf(v))
		}
		F9(p...)
	}
}

func BenchmarkRangeSliceF10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		F10(SI...)
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

func BenchmarkEachBuiltingF4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F4)
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

func BenchmarkEachBuiltingF9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F9)
	}
}

func BenchmarkEachBuiltingF10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F10)
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

func BenchmarkEachIndexableF4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F4)
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

func BenchmarkEachIndexableF9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F9)
	}
}

func BenchmarkEachIndexableF10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(IS, F10)
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

func BenchmarkEachMappableF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F1)
	}
}

func BenchmarkEachMappableF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F2)
	}
}

func BenchmarkEachMappableF3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F3)
	}
}

func BenchmarkEachMappableF4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F4)
	}
}

func BenchmarkEachMappableF5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F5)
	}
}

func BenchmarkEachMappableF6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F6)
	}
}

func BenchmarkEachMappableF7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F7)
	}
}

func BenchmarkEachMappableF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F8)
	}
}

func BenchmarkEachMappableF9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F9)
	}
}

func BenchmarkEachMappableF10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F10)
	}
}

func BenchmarkEachMappableF11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F11)
	}
}

func BenchmarkEachMappableF12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MM, F12)
	}
}

func BenchmarkEachMappableF13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MSM, F13)
	}
}

func BenchmarkEachMappableF14(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MSM, F14)
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

func BenchmarkEachSliceF4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F4)
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

func BenchmarkEachSliceF9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F9)
	}
}

func BenchmarkEachSliceF10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(SUDT, F10u)
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

func BenchmarkEachMapF4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F4)
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

func BenchmarkEachMapF9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F9)
	}
}

func BenchmarkEachMapF10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MII, F10)
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

func BenchmarkEachMapF13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MSI, F13)
	}
}

func BenchmarkEachMapF14(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(MSI, F14)
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

func BenchmarkEachChannelF4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F4)
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

func BenchmarkEachChannelF9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F9)
	}
}

func BenchmarkEachChannelF10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(CI, F10)
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

func BenchmarkEachFunctionF4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F4)
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

func BenchmarkEachFunctionF9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F9)
	}
}

func BenchmarkEachFunctionF10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Each(FI, F10)
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