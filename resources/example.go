package resources

import (
	"net"
	"reflect"
)

func ExampleFunc1() {}

func ExampleFunc2(v1 map[string]string) map[string]string { return nil }

func ExampleFunc3(v1 reflect.SliceHeader) []int { return nil }

func OneTemplateFunc[T any](tv *T) *T { return nil }

func DoubleSameTemplateFunc[T1, T2 any](tv1 T1, tv2 T2) (*T1, *T2) { return nil, nil }

func DoubleDifferenceTemplateFunc[T1 any, T2 comparable](tv1 T1, tv2 T2) (*T1, *T2) { return nil, nil }

type TypeConstraints interface{ int8 | int16 | uint8 | uint16 }

func TypeConstraintsTemplateFunc[T TypeConstraints](tv T) *T { return nil }

type ExampleStruct1 struct {
	b    bool
	i64  int64
	f64  float64
	c128 complex128
}

func (s *ExampleStruct1) ExampleStruct1Method1(b bool, i64 int64, f64 float64, c128 complex128) (bool, int64, float64, complex128) {
	return s.b, s.i64, s.f64, s.c128
}

type ExampleStruct2 struct {
	sPtr *ExampleStruct1
	m    map[int]string
	f    func() int
	c    chan int
}

type ExampleStruct3 struct {
	*ExampleStruct2
	i     net.Listener
	sh    reflect.SliceHeader
	str   string
	slice []int
}

type NormalInterface interface {
	NormalInterfaceMethod(int8, int16, uint8, uint16) (int8, int16, uint8, uint16)
}

type GenericInterface interface {
	int8 | int16 | uint8 | uint16
	GenericInterfaceMethod(int8, int16, uint8, uint16) (int8, int16, uint8, uint16)
}

type ExampleTemplateInterface[T any] interface {
	// This is ExampleFunc Doc
	ExampleFunc(T)
	// This is AnotherExampleFunc Doc
	AnotherExampleFunc(T, []T) (T, []T)
}

type OneTypeTemplateStruct[T any] struct {
	v T
}

func (t *OneTypeTemplateStruct[T]) V() T {
	return t.v
}

type TwoTypeTemplateStruct[K TypeConstraints, V any] struct {
	v map[K]V
}

func (t *TwoTypeTemplateStruct[K, V]) KVSlice() ([]K, []V) {
	ks := make([]K, 0, len(t.v))
	vs := make([]V, 0, len(t.v))
	for k, v := range t.v {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return ks, vs
}
