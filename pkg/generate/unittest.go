package generate

import (
	"fmt"
	"strings"

	"github.com/Mericusta/go-extractor"
)

func GenerateUnittest(argFilepath, argFuncName, argTypeArgs, argMode string) {
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

	unittestFuncName, unittestFuncByte := handleFuncMeta.MakeUnitTest(strings.Split(argTypeArgs, ","))
	unittestFilepath := fmt.Sprintf("%v_test.go", strings.Trim(handleFileMeta.Path(), ".go"))

	handleOutput(unittestFilepath, handleFileMeta.PkgName(), unittestFuncName, unittestFuncByte, argMode)
}
