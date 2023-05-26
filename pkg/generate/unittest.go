package generate

import (
	"fmt"
	"strings"

	"github.com/Mericusta/go-assistant/pkg/utility"
	"github.com/Mericusta/go-extractor"
)

func GenerateUnittest(argFilepath, argFuncName, argStructName, argInterfaceName, argTypeArgs, argMode string) {
	if len(argFilepath) == 0 || len(argFuncName) == 0 {
		fmt.Printf("not enough options, file %v, func %v\n", argFilepath, argFuncName)
		return
	}

	handleFileMeta := utility.HandleFileMeta(argFilepath)
	if handleFileMeta == nil {
		return
	}

	var (
		handleMeta  extractor.GoTestMaker
		isMethod    bool = len(argStructName) != 0
		isInterface bool = len(argInterfaceName) != 0
	)

	switch {
	case isMethod:
		handleMeta = extractor.SearchGoMethodMeta(handleFileMeta, argStructName, argFuncName)
	case isInterface:
		gim := extractor.SearchGoInterfaceMeta(handleFileMeta, argInterfaceName)
		if gim == nil {
			fmt.Printf("can not find interface")
			return
		}
		handleMeta = gim.SearchMethodDecl(argFuncName)
	default:
		handleMeta = extractor.SearchGoFunctionMeta(handleFileMeta, argFuncName)
	}
	if handleMeta == nil {
		fmt.Printf("can not find meta\n")
		return
	}

	var argTypes []string
	if len(argTypeArgs) != 0 {
		argTypes = strings.Split(argTypeArgs, ",")
	}

	unittestFuncName, unittestFuncByte := handleMeta.MakeUnitTest(argTypes)
	handleFilePath := handleFileMeta.Path()
	unittestFilepath := fmt.Sprintf("%v_test.go", handleFilePath[:strings.LastIndex(handleFilePath, ".go")])

	handleOutput(unittestFilepath, handleFileMeta.PkgName(), unittestFuncName, unittestFuncByte, argMode)
}
