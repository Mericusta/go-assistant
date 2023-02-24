package generate

import (
	"testing"
)

func Test_GenerateBenchmark(t *testing.T) {
	type args struct {
		argFilepath  string
		argFuncName  string
		argMode      string
		argTypeArgs  string
		fromUnittest bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{
				argFilepath: "../../resources/example.go",
				argFuncName: "ExampleFunc1",
				argMode:     "replace",
			},
		},
		{
			"test case 2",
			args{
				argFilepath: "../../resources/example.go",
				argFuncName: "ExampleFunc2",
				argMode:     "replace",
			},
		},
		{
			"test case 3",
			args{
				argFilepath: "../../resources/example.go",
				argFuncName: "ExampleFunc3",
				argMode:     "replace",
			},
		},
		{
			"test case 4",
			args{
				argFilepath: "../../resources/example.go",
				argFuncName: "OneTemplateFunc",
				argTypeArgs: "string",
				argMode:     "replace",
			},
		},
		{
			"test case 5",
			args{
				argFilepath: "../../resources/example.go",
				argFuncName: "DoubleSameTemplateFunc",
				argTypeArgs: "reflect.SliceHeader",
				argMode:     "replace",
			},
		},
		{
			"test case 6",
			args{
				argFilepath: "../../resources/example.go",
				argFuncName: "DoubleDifferenceTemplateFunc",
				argTypeArgs: "string,reflect.StringHeader",
				argMode:     "replace",
			},
		},
		{
			"test case 7",
			args{
				argFilepath: "../../resources/example.go",
				argFuncName: "TypeConstraintsTemplateFunc",
				argTypeArgs: "int8",
				argMode:     "replace",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateBenchmark(tt.args.argFilepath, tt.args.argFuncName, tt.args.argMode, tt.args.argTypeArgs, tt.args.fromUnittest)
		})
	}
}
