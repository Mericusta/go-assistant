package generate

import (
	"github.com/Mericusta/go-extractor"
)

func GenerateUnittest(argFilepath, argFuncName, argStructName, argInterfaceName, argTypeArgs, argMode, arg string) {
	handleGenerateTest(argFilepath, argFuncName, argStructName, argInterfaceName, argTypeArgs, argMode, arg, "test", extractor.GoTestMaker.MakeUnitTest)
}
