package main

import (
	"flag"

	"github.com/Mericusta/go-assistant/pkg/generate"
	"github.com/Mericusta/go-assistant/pkg/infer"
	"github.com/Mericusta/go-assistant/pkg/operate"
	"github.com/Mericusta/go-assistant/pkg/replace"
	"github.com/Mericusta/go-assistant/pkg/search"
	"github.com/Mericusta/go-assistant/pkg/secret"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	// temp var
	dev = flag.Bool("dev", false, "use dev UI assistant")

	// command var
	command = flag.String("cmd", "", "command")
	option  = flag.String("opt", "", "command option")
	args    = flag.String("arg", "", "command args")

	// common var
	argFilepath = flag.String("file", "", "handle file path")
	argMode     = flag.String("mode", "replace", "replace or append or preview in stdout")

	// generate var
	argFuncName      = flag.String("func", "", "generate test function name")
	argStructName    = flag.String("struct", "", "generate test struct name")
	argInterfaceName = flag.String("interface", "", "generate test interface name")
	argTypeArgs      = flag.String("types", "", "generate generic test function type arguments")
	argMetaType      = flag.String("meta", "", "search meta type: func, method, struct, interface")
	argMetaIdent     = flag.String("ident", "", "search meta ident: ident1[,ident1]")

	// infer var
	argPlatform          = flag.Int("platform", 64, "32 or 64")
	argProcess           = flag.Bool("process", false, "show calculate process or not")
	argAllocationPreview = flag.Bool("preview", false, "show allocation preview")

	// search var
	argRegexp = flag.String("regexp", "", "content regexp which contains search key, must like (?P<KEY>regexp)")

	// secret var
	argSecretInputFile  = flag.String("input", "", "secret input file")
	argSecretOutputFile = flag.String("output", "", "secret output file")

	// redis/mysql var
	argSource = flag.String("source", "", "operate source: redis, mysql")
	argURL    = flag.String("url", "", "url")
)

func init() {
	flag.Parse()
}

func main() {
	if *dev {
		mainForDevUI()
	} else {
		mainForCommand()
	}
}

func mainForCommand() {
	switch {
	case *command == "generate" && *option == "unittest":
		generate.GenerateUnittest(*argFilepath, *argFuncName, *argStructName, *argInterfaceName, *argTypeArgs, *argMode, *args)
	case *command == "generate" && *option == "benchmark":
		generate.GenerateBenchmark(*argFilepath, *argFuncName, *argStructName, *argInterfaceName, *argTypeArgs, *argMode, *args)
	case *command == "generate" && *option == "implement":
		generate.GenerateImplement(*argFilepath, *argInterfaceName, *argMetaIdent, *argTypeArgs, *args)
	case *command == "generate" && *option == "ast":
		generate.GenerateAST(*argFilepath, *argMetaType, *argMetaIdent)
	case *command == "infer" && *option == "allocation":
		infer.InferTheOptimalLayoutOfStructMemory(*argPlatform, *argFilepath, *argStructName, *argAllocationPreview, *argProcess)
	case *command == "search" && *option == "log":
		search.SplitLogByKey(*argFilepath, *argMode, *argRegexp)
	case *command == "generate" && *option == "secret":
		secret.Secret(*argSecretInputFile, *argSecretOutputFile, *argMode, *args, *argRegexp)
	case *command == "operate" && *argSource == "redis":
		operate.OperateRedis(*argURL, *option, *argRegexp)
	case *command == "operate" && *argSource == "mysql":
		operate.OperateMySQL(*argURL, *option, *argFilepath, *args)
	case *command == "operate" && *argSource == "markdown":
		operate.OperateMarkdownTable(*option, *argFilepath, *args)
	case *command == "replace":
		replace.ReplaceCode(*argMetaType, *argMetaIdent, *argRegexp, *args)
	}
}

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func mainForDevUI() {

}
