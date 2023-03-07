package main

import (
	"flag"

	"github.com/Mericusta/go-assistant/pkg/generate"
	"github.com/Mericusta/go-assistant/pkg/infer"
	"github.com/Mericusta/go-assistant/pkg/search"
)

var (
	// command var
	command = flag.String("cmd", "", "command")
	option  = flag.String("opt", "", "command option")
	args    = flag.String("arg", "", "command args")

	// common var
	argFilepath = flag.String("file", "", "handle file path")
	argMode     = flag.String("mode", "replace", "replace or append or preview in stdout")

	// generate var
	argFuncName   = flag.String("func", "", "generate test function name")
	argTypeArgs   = flag.String("types", "", "generate generic test function type arguments")
	argMetaType   = flag.String("meta", "", "search meta type: func, method, struct, interface")
	argMetaIdents = flag.String("idents", "", "search meta idents: ident1[,ident1]")

	// infer var
	argPlatform          = flag.Int("platform", 64, "32 or 64")
	argStructName        = flag.String("struct", "", "infer struct name")
	argProcess           = flag.Bool("process", false, "show calculate process or not")
	argAllocationPreview = flag.Bool("preview", false, "show allocation preview")

	// search var
	argRegexp = flag.String("regexp", "", "content regexp which contains search key, must like (?P<KEY>regexp)")
)

func init() {
	flag.Parse()
}

func main() {
	switch {
	case *command == "generate" && *option == "unittest":
		generate.GenerateUnittest(*argFilepath, *argFuncName, *argTypeArgs, *argMode)
	case *command == "generate" && *option == "benchmark":
		generate.GenerateBenchmark(*argFilepath, *argFuncName, *argMode, *argTypeArgs)
	case *command == "generate" && *option == "ast":
		generate.GenerateAST(*argFilepath, *argMetaType, *argMetaIdents)
	case *command == "infer" && *option == "allocation":
		infer.InferTheOptimalLayoutOfStructMemory(*argPlatform, *argFilepath, *argStructName, *argAllocationPreview, *argProcess)
	case *command == "search" && *option == "log":
		search.SplitLogByKey(*argFilepath, *argMode, *argRegexp)
	}
}
