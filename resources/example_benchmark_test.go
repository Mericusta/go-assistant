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
