package generate

import (
	"fmt"
	"strings"

	"github.com/Mericusta/go-extractor"
)

func GenerateBenchmark(argFilepath, argFuncName, argMode, argTypeArgs string, fromUnittest bool) {
	if len(argFilepath) == 0 || len(argFuncName) == 0 {
		fmt.Printf("not enough options, file %v, func %v\n", argFilepath, argFuncName)
		return
	}

	handleFileMeta := handleFileMeta(argFilepath)
	handleFuncMeta := extractor.SearchGoFunctionMeta(handleFileMeta, argFuncName)
	if handleFuncMeta == nil {
		fmt.Printf("can not find func meta\n")
		return
	}

	benchmarkFuncName, benchmarkFuncByte := handleFuncMeta.MakeBenchmark(strings.Split(argTypeArgs, ","))
	benchmarkFilepath := fmt.Sprintf("%v_benchmark_test.go", strings.Trim(handleFileMeta.Path(), ".go"))

	handleOutput(benchmarkFilepath, handleFileMeta.PkgName(), benchmarkFuncName, benchmarkFuncByte, argMode)
}
