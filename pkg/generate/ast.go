package generate

import (
	"fmt"
	"strings"

	"github.com/Mericusta/go-assistant/pkg/utility"
	"github.com/Mericusta/go-extractor"
)

func GenerateAST(argFilepath, argIdentType, argIdentNames string) {
	if len(argFilepath) == 0 || len(argIdentType) == 0 || len(argIdentNames) == 0 {
		fmt.Printf("not enough options, file %v, meta type %v, meta ident names %v\n", argFilepath, argIdentType, argIdentNames)
		return
	}

	argIdentNameSlice := strings.Split(argIdentNames, ",")

	handleFileMeta := utility.HandleFileMeta(argFilepath)
	if handleFileMeta == nil {
		return
	}

	var meta extractor.Meta
	switch argIdentType {
	case "func":
		meta = extractor.SearchGoFunctionMeta(handleFileMeta, argIdentNameSlice[0])
	case "method":
		meta = extractor.SearchGoMethodMeta(handleFileMeta, argIdentNameSlice[0], argIdentNameSlice[1])
	case "struct":
		meta = extractor.SearchGoStructMeta(handleFileMeta, argIdentNameSlice[0])
	case "interface":
		meta = extractor.SearchGoInterfaceMeta(handleFileMeta, argIdentNameSlice[0])
	default:
	}

	if meta == nil {
		fmt.Printf("can not find ident meta\n")
		return
	}

	meta.PrintAST()
}
