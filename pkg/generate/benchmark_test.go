package generate

import (
	"reflect"
	"testing"
)

func Test_SliceGenerateBenchmark(t *testing.T) {
	type args struct {
		arg1 string
		arg2 string
	}
	tests := []struct {
		name  string
		args  args
		want0 *generateBenchmarkArgs
		want1 []*generateBenchmarkArgs
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0, got1 := SliceGenerateBenchmark(tt.args.arg1, tt.args.arg2)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("SliceGenerateBenchmark() got0 = %v, want0 %v", got0, tt.want0)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SliceGenerateBenchmark() got1 = %v, want1 %v", got1, tt.want1)
			}
		})
	}
}
