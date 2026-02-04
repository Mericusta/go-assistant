package main

import (
	"flag"

	goass "github.com/Mericusta/go-assistant/src"
	"github.com/Mericusta/go-assistant/src/logic/generate"
	"github.com/Mericusta/go-assistant/src/logic/infer"
	"github.com/Mericusta/go-assistant/src/logic/monitor"
	"github.com/Mericusta/go-assistant/src/logic/operate"
	"github.com/Mericusta/go-assistant/src/logic/replace"
	"github.com/Mericusta/go-assistant/src/logic/search"
	"github.com/Mericusta/go-assistant/src/logic/secret"
	"github.com/Mericusta/go-assistant/src/logic/statistic"
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
	argSplitRegexp      = flag.String("split_regexp", "", "content regexp which contains search split key, must like (?P<KEY>regexp)")
	argFuncRegexp       = flag.String("func_regexp", "", "content regexp which contains search func key, must like (?P<KEY>regexp)")
	argStatisticsRegexp = flag.String("statistics_regexp", "", "content  regexp which contains search statistics key, must like (?P<KEY>regexp)")

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
		search.SplitLogByKey(*argFilepath, *argMode, *argSplitRegexp)
	case *command == "split" && *option == "log":
		search.SplitLogByLine(*argFilepath, *argMode, *args)
	case *command == "search" && *option == "func":
		search.SearchFuncByRegexp(*argFilepath, *argMode, *argFuncRegexp)
	case *command == "statistic" && *option == "log":
		statistic.StatisticByKey(*argFilepath, *argSplitRegexp, *argStatisticsRegexp, *args)
	case *command == "generate" && *option == "secret":
		secret.Secret(*argSecretInputFile, *argSecretOutputFile, *argMode, *args, *argSplitRegexp)
	case *command == "operate" && *argSource == "redis":
		operate.OperateRedis(*argURL, *option, *argSplitRegexp)
	case *command == "operate" && *argSource == "mysql":
		operate.OperateMySQL(*argURL, *option, *argFilepath, *args)
	case *command == "operate" && *argSource == "markdown":
		operate.OperateMarkdownTable(*option, *argFilepath, *args)
	case *command == "replace":
		replace.ReplaceCode(*argMetaType, *argMetaIdent, *argSplitRegexp, *args)
	case *command == "monitor":
		monitor.MonitorProcess(*args)
	case *command == "search" && *option == "grpc":
		search.SearchGRPCServiceMethods(*argFilepath, *argInterfaceName)
	case *command == "generate" && *option == "grpc":
		generate.GenerateGRPC_CS(*argFilepath, *argInterfaceName)
	}
}

func mainForDevUI() {
	goass.GoAssistant()
}
