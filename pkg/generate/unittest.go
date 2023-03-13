package generate

import (
	"fmt"
	"strings"

	"github.com/Mericusta/go-assistant/pkg/utility"
	"github.com/Mericusta/go-extractor"
)

func GenerateUnittest(argFilepath, argFuncName, argTypeArgs, argMode string) {
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
	unittestFuncName, unittestFuncByte := handleFuncMeta.MakeUnitTest(argTypes)
	handleFilePath := handleFileMeta.Path()
	unittestFilepath := fmt.Sprintf("%v_test.go", handleFilePath[:strings.LastIndex(handleFilePath, ".go")])

	handleOutput(unittestFilepath, handleFileMeta.PkgName(), unittestFuncName, unittestFuncByte, argMode)
}
