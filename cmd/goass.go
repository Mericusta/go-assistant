package main

import (
	"flag"

	"github.com/Mericusta/go-assistant/pkg/generate"
)

// command flag var
var (
	command = flag.String("cmd", "", "command")
	option  = flag.String("opt", "", "command option")
	args    = flag.String("arg", "", "command args")
)

// common var
var (
	argFilepath = flag.String("file", "", "handle file path")
	argMode     = flag.String("mode", "replace", "replace or append or preview in stdout")
)

// unittest flag var
var (
	argFuncName = flag.String("func", "", "unit test function name")
	argTypeArgs = flag.String("types", "", "unit test generic func type arguments")
)

// benchmark flag var
var (
	argFromUnittest = flag.Bool("fromUnittest", true, "generate benchmark by unittest cases")
)

func init() {
	flag.Parse()
}

func main() {
	switch {
	case *command == "generate" && *option == "unittest":
		generate.GenerateUnittest(*argFilepath, *argFuncName, *argTypeArgs, *argMode)
	case *command == "generate" && *option == "benchmark":
		generate.GenerateBenchmark(*argFilepath, *argFuncName, *argMode, *argFromUnittest)
	}
}
