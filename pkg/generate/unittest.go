package generate

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mericusta/go-extractor"
	stpconvert "github.com/Mericusta/go-stp/convert"
)

type _generateUnittestArg struct {
	handlePath     string
	handleFunc     string
	handleTypeArgs string
}

var (
	generateUnittestArg *_generateUnittestArg
)

func GenerateUnittest(args string) {
	generateUnittestArg = stpconvert.ConvertStringToStringStruct[_generateUnittestArg](args, ";")

	handlePathAbs, err := filepath.Abs(generateUnittestArg.handlePath)
	if err != nil {
		fmt.Printf("get file abs path occurs error: %v\n", err)
		return
	}

	handlePathStat, err := os.Stat(handlePathAbs)
	if err != nil {
		fmt.Printf("get file stat occurs error: %v\n", err)
		return
	}

	if handlePathStat == nil {
		fmt.Printf("file not exist\n")
		return
	}

	if len(generateUnittestArg.handleFunc) == 0 {
		fmt.Printf("need specify a func\n")
		return
	}

	if handlePathStat != nil && handlePathStat.IsDir() {
		fmt.Printf("not support dir\n")
		return
	}

	gfm, err := extractor.ExtractGoFileMeta(handlePathAbs)
	if gfm == nil || err != nil {
		fmt.Printf("extract file meta occurs error: %v\n", err)
		return
	}

	unittestFuncByte := handleSingleFunc(handlePathAbs, generateUnittestArg.handleFunc, strings.Split(generateUnittestArg.handleTypeArgs, ","))

	unittestFilePath := fmt.Sprintf("%v_test.go", strings.Trim(handlePathAbs, ".go"))
	unittestFileStat, err := os.Stat(unittestFilePath)
	if unittestFileStat == nil && !errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("get unit test file %v stat occurs error: %v\n", unittestFilePath, err)
		return
	}

	unittestFileHandler, err := os.OpenFile(unittestFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("create unit test file %v occurs error: %v\n", unittestFilePath, err)
		return
	}
	defer unittestFileHandler.Close()

	if unittestFileStat == nil {
		unittestFileHandler.Write(extractor.MakeUnitTestFile(gfm.PkgName(), nil))
	}

	unittestFileHandler.WriteString("\n")
	unittestFileHandler.Write(unittestFuncByte)

	extractor.GoFmtFile(unittestFilePath)
}

// func handleSingleFile(handlePath string) {
// 	extractor.ExtractGoFileMeta(handlePath)
// }

func handleSingleFunc(handlePath, handleFunc string, handleTypeArgs []string) []byte {
	// gfm, _ := extractor.ExtractGoFileMeta(handlePath)
	// gfm.PrintAST()
	// return nil

	gfm, err := extractor.ExtractGoFunctionMeta(handlePath, handleFunc)
	if err != nil {
		fmt.Printf("extract go function meta occurs error: %v\n", err)
		return nil
	}

	if len(gfm.TypeParams()) > 0 {
		return extractor.MakeUnitTestWithTypeArgs(gfm, handleTypeArgs)
	}

	return extractor.MakeUnitTest(gfm)
}
