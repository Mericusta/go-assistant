package infer

import (
	"testing"
)

func Test_InferTheOptimalLayoutOfStructMemory(t *testing.T) {
	type args struct {
		argPlatform          int
		argFilepath          string
		argStructName        string
		argAllocationPreview bool
		argProcess           bool
	}
	tests := []struct {
		name string
		args args
	}{
		// {
		// 	"test case 1",
		// 	args{
		// 		argPlatform:          64,
		// 		argFilepath:          "../../resources/example.go",
		// 		argStructName:        "ExampleStruct1",
		// 		argAllocationPreview: true,
		// 	},
		// },
		// {
		// 	"test case 2",
		// 	args{
		// 		argPlatform:          64,
		// 		argFilepath:          "../../resources/example.go",
		// 		argStructName:        "ExampleStruct2",
		// 		argAllocationPreview: true,
		// 	},
		// },
		{
			"test case 3",
			args{
				argPlatform:          64,
				argFilepath:          "../../resources/example.go",
				argStructName:        "ExampleStruct3",
				argProcess:           false,
				argAllocationPreview: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InferTheOptimalLayoutOfStructMemory(tt.args.argPlatform, tt.args.argFilepath, tt.args.argStructName, tt.args.argAllocationPreview, tt.args.argProcess)
		})
	}
}
