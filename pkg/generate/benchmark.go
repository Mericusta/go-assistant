package generate

import (
	"github.com/Mericusta/go-extractor"
)

func GenerateBenchmark(argFilepath, argFuncName, argStructName, argInterfaceName, argTypeArgs, argMode, args string) {
	handleGenerateTest(argFilepath, argFuncName, argStructName, argInterfaceName, argTypeArgs, argMode, args, "benchmark_test", extractor.GoTestMaker.MakeBenchmark)
}
