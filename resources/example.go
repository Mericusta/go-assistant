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
