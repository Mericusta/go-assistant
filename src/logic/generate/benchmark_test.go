package generate

import (
	"testing"
)

func Test_GenerateBenchmark(t *testing.T) {
	type args struct {
		argFilepath      string
		argFuncName      string
		argStructName    string
		argInterfaceName string
		argTypeArgs      string
		argMode          string
		arg              string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
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
		{
			"test case 8",
			args{
				argFilepath:   "../../resources/example.go",
				argFuncName:   "ExampleStruct1Method1",
				argStructName: "ExampleStruct1",
				argMode:       "replace",
			},
		},
		{
			"test case 9",
			args{
				argFilepath:      "../../resources/example.go",
				argFuncName:      "NormalInterfaceMethod",
				argInterfaceName: "NormalInterface",
				argMode:          "replace",
			},
		},
		{
			"test case 10",
			args{
				argFilepath:      "../../resources/example.go",
				argFuncName:      "ExampleFunc",
				argInterfaceName: "ExampleTemplateInterface",
				argTypeArgs:      "string",
				argMode:          "replace",
			},
		},
		{
			"test case 11",
			args{
				argFilepath:      "../../resources/example.go",
				argFuncName:      "AnotherExampleFunc",
				argInterfaceName: "ExampleTemplateInterface",
				argTypeArgs:      "string",
				argMode:          "replace",
			},
		},
		{
			"test case 12",
			args{
				argFilepath:   "../../resources/example.go",
				argFuncName:   "V",
				argStructName: "OneTypeTemplateStruct",
				argTypeArgs:   "string",
				argMode:       "replace",
			},
		},
		{
			"test case 13",
			args{
				argFilepath:   "../../resources/example.go",
				argFuncName:   "KVSlice",
				argStructName: "TwoTypeTemplateStruct",
				argTypeArgs:   "int8,reflect.StringHeader",
				argMode:       "replace",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateBenchmark(tt.args.argFilepath, tt.args.argFuncName, tt.args.argStructName, tt.args.argInterfaceName, tt.args.argTypeArgs, tt.args.argMode, tt.args.arg)
		})
	}
}
