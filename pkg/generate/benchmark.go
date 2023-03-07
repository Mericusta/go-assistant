package generate

import (
	"fmt"
	"strings"

	"github.com/Mericusta/go-assistant/pkg/utility"
	"github.com/Mericusta/go-extractor"
)

func GenerateBenchmark(argFilepath, argFuncName, argMode, argTypeArgs string) {
	if len(argFilepath) == 0 || len(argFuncName) == 0 {
		fmt.Printf("not enough options, file %v, func %v\n", argFilepath, argFuncName)
		return
	}

	handleFileMeta := utility.HandleFileMeta(argFilepath)
	if handleFileMeta == nil {
		return
	}
	handleFuncMeta := extractor.SearchGoFunctionMeta(handleFileMeta, argFuncName)
	if handleFuncMeta == nil {
		fmt.Printf("can not find func meta\n")
		return
	}

	var argTypes []string
	if len(argTypeArgs) != 0 {
		argTypes = strings.Split(argTypeArgs, ",")
	}
	benchmarkFuncName, benchmarkFuncByte := handleFuncMeta.MakeBenchmark(argTypes)
	benchmarkFilepath := fmt.Sprintf("%v_benchmark_test.go", strings.Trim(handleFileMeta.Path(), ".go"))

	handleOutput(benchmarkFilepath, handleFileMeta.PkgName(), benchmarkFuncName, benchmarkFuncByte, argMode)
}
