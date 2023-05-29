package resources

import (
	"reflect"
	"testing"
	"time"
)

func Benchmark_ExampleFunc1(b *testing.B) {
	tests := []struct {
		name  string
		limit time.Duration
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ExampleFunc1()
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_ExampleFunc2(b *testing.B) {
	type args struct {
		v1 map[string]string
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ExampleFunc2(tt.args.v1)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_ExampleFunc3(b *testing.B) {
	type args struct {
		v1 reflect.SliceHeader
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ExampleFunc3(tt.args.v1)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_OneTemplateFunc_6d6d5833d0a567609a6a6de032241d35(b *testing.B) {
	type args struct {
		tv *string
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = OneTemplateFunc(tt.args.tv)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_DoubleDifferenceTemplateFunc_ced46be80777d0e8dc06e94a495ea5b6(b *testing.B) {
	type args struct {
		tv1 string
		tv2 reflect.StringHeader
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = DoubleDifferenceTemplateFunc(tt.args.tv1, tt.args.tv2)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_TypeConstraintsTemplateFunc_e9568baeb7bb7af5d4168fa861979e8c(b *testing.B) {
	type args struct {
		tv int8
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = TypeConstraintsTemplateFunc(tt.args.tv)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_ExampleStruct1_ExampleStruct1Method1(b *testing.B) {
	type args struct {
		b    bool
		i64  int64
		f64  float64
		c128 complex128
	}
	tests := []struct {
		name  string
		limit time.Duration
		s     *ExampleStruct1
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, _, _ = tt.s.ExampleStruct1Method1(tt.args.b, tt.args.i64, tt.args.f64, tt.args.c128)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_NormalInterface_NormalInterfaceMethod(b *testing.B) {
	type args struct {
		p0 int8
		p1 int16
		p2 uint8
		p3 uint16
	}
	tests := []struct {
		name  string
		limit time.Duration
		i     NormalInterface
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, _, _ = tt.i.NormalInterfaceMethod(tt.args.p0, tt.args.p1, tt.args.p2, tt.args.p3)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_ExampleTemplateInterface_ExampleFunc_6d6d5833d0a567609a6a6de032241d35(b *testing.B) {
	type args struct {
		p0 string
	}
	tests := []struct {
		name  string
		limit time.Duration
		i     ExampleTemplateInterface[string]
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tt.i.ExampleFunc(tt.args.p0)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_ExampleTemplateInterface_AnotherExampleFunc_6d6d5833d0a567609a6a6de032241d35(b *testing.B) {
	type args struct {
		p0 string
		p1 []string
	}
	tests := []struct {
		name  string
		limit time.Duration
		i     ExampleTemplateInterface[string]
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = tt.i.AnotherExampleFunc(tt.args.p0, tt.args.p1)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_OneTypeTemplateStruct_V_6d6d5833d0a567609a6a6de032241d35(b *testing.B) {
	tests := []struct {
		name  string
		limit time.Duration
		t     *OneTypeTemplateStruct[string]
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = tt.t.V()
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_TwoTypeTemplateStruct_KVSlice_f8d79beb910f9b7740f701743544b509(b *testing.B) {
	type args struct {
		k int8
		v reflect.StringHeader
	}
	tests := []struct {
		name  string
		limit time.Duration
		t     *TwoTypeTemplateStruct[int8, reflect.StringHeader]
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = tt.t.KVSlice(tt.args.k, tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_DoubleSameTemplateFunc_3b4f7099c4644d92f9dcc05adad31ba6(b *testing.B) {
	type args struct {
		tv1 reflect.SliceHeader
		tv2 reflect.SliceHeader
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = DoubleSameTemplateFunc(tt.args.tv1, tt.args.tv2)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}
