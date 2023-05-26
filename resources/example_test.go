package resources

import (
	"reflect"
	"testing"
)

func Test_ExampleFunc1(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleFunc1()
		})
	}
}

func Test_ExampleFunc2(t *testing.T) {
	type args struct {
		v1 map[string]string
	}
	tests := []struct {
		name  string
		args  args
		want0 map[string]string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ExampleFunc2(tt.args.v1)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ExampleFunc2() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_ExampleFunc3(t *testing.T) {
	type args struct {
		v1 reflect.SliceHeader
	}
	tests := []struct {
		name  string
		args  args
		want0 []int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ExampleFunc3(tt.args.v1)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ExampleFunc3() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_OneTemplateFunc_6d6d5833d0a567609a6a6de032241d35(t *testing.T) {
	type args struct {
		tv *string
	}
	tests := []struct {
		name  string
		args  args
		want0 *string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := OneTemplateFunc(tt.args.tv)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("OneTemplateFunc() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_DoubleDifferenceTemplateFunc_ced46be80777d0e8dc06e94a495ea5b6(t *testing.T) {
	type args struct {
		tv1 string
		tv2 reflect.StringHeader
	}
	tests := []struct {
		name  string
		args  args
		want0 *string
		want1 *reflect.StringHeader
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0, got1 := DoubleDifferenceTemplateFunc(tt.args.tv1, tt.args.tv2)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("DoubleDifferenceTemplateFunc() got0 = %v, want0 %v", got0, tt.want0)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DoubleDifferenceTemplateFunc() got1 = %v, want1 %v", got1, tt.want1)
			}
		})
	}
}

func Test_TypeConstraintsTemplateFunc_e9568baeb7bb7af5d4168fa861979e8c(t *testing.T) {
	type args struct {
		tv int8
	}
	tests := []struct {
		name  string
		args  args
		want0 *int8
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := TypeConstraintsTemplateFunc(tt.args.tv)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("TypeConstraintsTemplateFunc() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_ExampleStruct1_ExampleStruct1Method1(t *testing.T) {
	type args struct {
		b    bool
		i64  int64
		f64  float64
		c128 complex128
	}
	tests := []struct {
		name  string
		s     *ExampleStruct1
		args  args
		want0 bool
		want1 int64
		want2 float64
		want3 complex128
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0, got1, got2, got3 := tt.s.ExampleStruct1Method1(tt.args.b, tt.args.i64, tt.args.f64, tt.args.c128)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ExampleStruct1Method1() got0 = %v, want0 %v", got0, tt.want0)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ExampleStruct1Method1() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ExampleStruct1Method1() got2 = %v, want2 %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("ExampleStruct1Method1() got3 = %v, want3 %v", got3, tt.want3)
			}
		})
	}
}

func Test_NormalInterface_NormalInterfaceMethod(t *testing.T) {
	type args struct {
		p0 int8
		p1 int16
		p2 uint8
		p3 uint16
	}
	tests := []struct {
		name  string
		i     NormalInterface
		args  args
		want0 int8
		want1 int16
		want2 uint8
		want3 uint16
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0, got1, got2, got3 := tt.i.NormalInterfaceMethod(tt.args.p0, tt.args.p1, tt.args.p2, tt.args.p3)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("NormalInterfaceMethod() got0 = %v, want0 %v", got0, tt.want0)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NormalInterfaceMethod() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("NormalInterfaceMethod() got2 = %v, want2 %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("NormalInterfaceMethod() got3 = %v, want3 %v", got3, tt.want3)
			}
		})
	}
}

func Test_ExampleTemplateInterface_ExampleFunc_6d6d5833d0a567609a6a6de032241d35(t *testing.T) {
	type args struct {
		p0 string
	}
	tests := []struct {
		name string
		i    ExampleTemplateInterface[string]
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ExampleFunc(tt.args.p0)
		})
	}
}

func Test_OneTypeTemplateStruct_V_6d6d5833d0a567609a6a6de032241d35(t *testing.T) {
	tests := []struct {
		name  string
		s     *OneTypeTemplateStruct[int]
		want0 int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := tt.s.V()
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("V() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}
