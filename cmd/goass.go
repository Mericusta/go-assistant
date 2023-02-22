package main

import (
	"flag"

	"github.com/Mericusta/go-assistant/pkg/generate"
)

// flag var
var (
	command = flag.String("cmd", "", "command")
	option  = flag.String("opt", "", "command option")
	args    = flag.String("arg", "", "command args")
)

func init() {
	flag.Parse()
}

func main() {
	switch {
	case *command == "generate" && *option == "unittest":
		generate.GenerateUnittest(*args)
	case *command == "generate" && *option == "benchmark":
	}
}
