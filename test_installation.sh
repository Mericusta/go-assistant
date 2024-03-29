#!/bin/bash

go install ./cmd/goass.go

rm ./resources/example_test.go
rm ./resources/example_benchmark_test.go

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleFunc1
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleFunc1

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleFunc2
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleFunc2

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleFunc3
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleFunc3

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=OneTemplateFunc -types=string
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=OneTemplateFunc -types=string

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=DoubleSameTemplateFunc -types=reflect.SliceHeader,reflect.SliceHeader
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=DoubleSameTemplateFunc -types=reflect.SliceHeader,reflect.SliceHeader

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=DoubleDifferenceTemplateFunc -types=string,reflect.StringHeader
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=DoubleDifferenceTemplateFunc -types=string,reflect.StringHeader

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=TypeConstraintsTemplateFunc -types=int8
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=TypeConstraintsTemplateFunc -types=int8

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleStruct1Method1 -struct=ExampleStruct1
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleStruct1Method1 -struct=ExampleStruct1

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=NormalInterfaceMethod -interface=NormalInterface
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=NormalInterfaceMethod -interface=NormalInterface

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleFunc -interface=ExampleTemplateInterface -types=string
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleFunc -interface=ExampleTemplateInterface -types=string

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=AnotherExampleFunc -interface=ExampleTemplateInterface -types=string
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=AnotherExampleFunc -interface=ExampleTemplateInterface -types=string

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=V -struct=OneTypeTemplateStruct -types=string
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=V -struct=OneTypeTemplateStruct -types=string

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=KVSlice -struct=TwoTypeTemplateStruct -types=int8,reflect.StringHeader
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=KVSlice -struct=TwoTypeTemplateStruct -types=int8,reflect.StringHeader

sed -i 's/testing\.T/T\.testing/g' ./resources/example_test.go
sed -i 's/testing\.B/B\.testing/g' ./resources/example_benchmark_test.go

grep 'T\.testing' ./resources/example_test.go | wc -l
grep 'B\.testing' ./resources/example_benchmark_test.go | wc -l

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleFunc1 -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleFunc1 -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleFunc2 -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleFunc2 -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleFunc3 -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleFunc3 -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=OneTemplateFunc -types=string -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=OneTemplateFunc -types=string -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=DoubleSameTemplateFunc -types=reflect.SliceHeader,reflect.SliceHeader -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=DoubleSameTemplateFunc -types=reflect.SliceHeader,reflect.SliceHeader -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=DoubleDifferenceTemplateFunc -types=string,reflect.StringHeader -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=DoubleDifferenceTemplateFunc -types=string,reflect.StringHeader -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=TypeConstraintsTemplateFunc -types=int8 -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=TypeConstraintsTemplateFunc -types=int8 -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleStruct1Method1 -struct=ExampleStruct1 -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleStruct1Method1 -struct=ExampleStruct1 -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=NormalInterfaceMethod -interface=NormalInterface -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=NormalInterfaceMethod -interface=NormalInterface -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=ExampleFunc -interface=ExampleTemplateInterface -types=string -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=ExampleFunc -interface=ExampleTemplateInterface -types=string -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=AnotherExampleFunc -interface=ExampleTemplateInterface -types=string -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=AnotherExampleFunc -interface=ExampleTemplateInterface -types=string -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=V -struct=OneTypeTemplateStruct -types=string -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=V -struct=OneTypeTemplateStruct -types=string -mode=replace

goass -cmd=generate -opt=unittest -file=./resources/example.go -func=KVSlice -struct=TwoTypeTemplateStruct -types=int8,reflect.StringHeader -mode=replace
goass -cmd=generate -opt=benchmark -file=./resources/example.go -func=KVSlice -struct=TwoTypeTemplateStruct -types=int8,reflect.StringHeader -mode=replace

grep 'testing\.T' ./resources/example_test.go | wc -l
grep 'testing\.B' ./resources/example_benchmark_test.go | wc -l

rm ./resources/ExampleFunc1.ast
rm ./resources/DoubleDifferenceTemplateFunc.ast
rm ./resources/TypeConstraints.ast
rm ./resources/ExampleTemplateInterface.ast
rm ./resources/ExampleTemplateInterface_AnotherExampleFunc.ast
rm ./resources/OneTypeTemplateStruct.ast
rm ./resources/TwoTypeTemplateStruct_KVSlice.ast

# func
echo "func ExampleFunc1"
goass -cmd=generate -opt=ast -file=./resources/example.go -meta=func -ident=ExampleFunc1 > ./resources/ExampleFunc1.ast
# generic func
echo "func DoubleDifferenceTemplateFunc"
goass -cmd=generate -opt=ast -file=./resources/example.go -meta=func -ident=DoubleDifferenceTemplateFunc > ./resources/DoubleDifferenceTemplateFunc.ast
# interface(type constraints) not support yet
echo "interface TypeConstraints not support yet"
goass -cmd=generate -opt=ast -file=./resources/example.go -meta=interface -ident=TypeConstraints > ./resources/TypeConstraints.ast
# interface
echo "interface ExampleTemplateInterface"
goass -cmd=generate -opt=ast -file=./resources/example.go -meta=interface -ident=ExampleTemplateInterface > ./resources/ExampleTemplateInterface.ast
# interface method
echo "method ExampleTemplateInterface.AnotherExampleFunc"
goass -cmd=generate -opt=ast -file=./resources/example.go -meta=method -ident=ExampleTemplateInterface,AnotherExampleFunc > ./resources/ExampleTemplateInterface_AnotherExampleFunc.ast
# struct
echo "struct OneTypeTemplateStruct"
goass -cmd=generate -opt=ast -file=./resources/example.go -meta=struct -ident=OneTypeTemplateStruct > ./resources/OneTypeTemplateStruct.ast
# struct method
echo "method TwoTypeTemplateStruct.KVSlice"
goass -cmd=generate -opt=ast -file=./resources/example.go -meta=method -ident=TwoTypeTemplateStruct,KVSlice > ./resources/TwoTypeTemplateStruct_KVSlice.ast

ls ./resources/*.ast | wc -l