package generate

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/Mericusta/go-assistant/pkg/utility"
	"github.com/Mericusta/go-extractor"
)

func handleGenerateTest(argFilepath, argFuncName, argStructName, argInterfaceName, argTypeArgs, argMode, argTestFilePath, suffix string, maker func(extractor.GoTestMaker, []string) (string, []byte)) {
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
			fmt.Printf("can not find interface %v\n", argInterfaceName)
			return
		}
		handleMeta = gim.SearchMethodDecl(argFuncName)
	default:
		handleMeta = extractor.SearchGoFunctionMeta(handleFileMeta, argFuncName)
	}
	if handleMeta == nil {
		fmt.Printf("can not find meta %v\n", argFuncName)
		return
	}

	var argTypes []string
	if len(argTypeArgs) != 0 {
		argTypes = strings.Split(argTypeArgs, ",")
	}

	testFuncName, testFuncByte := maker(handleMeta, argTypes)
	handleFilePath := handleFileMeta.Path()

	testFilepath := argTestFilePath
	if len(testFilepath) == 0 {
		testFilepath = fmt.Sprintf("%v_%v.go", handleFilePath[:strings.LastIndex(handleFilePath, ".go")], suffix)
	}

	handleOutputTest(testFilepath, handleFileMeta.PkgName(), testFuncName, testFuncByte, argMode)
}

func handleOutputTest(testFilepath, testFilePkg, testFuncName string, testFuncByte []byte, argMode string) {
	if argMode == "preview" {
		os.Stdout.Write(testFuncByte)
		return
	}

	testFileStat, err := os.Stat(testFilepath)
	if testFileStat == nil && !errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("get test file %v stat occurs error: %v\n", testFilepath, err)
		return
	}

	testFileHandler, err := os.OpenFile(testFilepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("create test file %v occurs error: %v\n", testFilepath, err)
		return
	}
	defer testFileHandler.Close()

	if testFileStat == nil {
		testFileHandler.Write(extractor.MakeTestFile(testFilePkg, nil))
	}

	switch argMode {
	case "append":
		testFileHandler.WriteString("\n")
		testFileHandler.Write(testFuncByte)
	case "replace":
		testFileMeta, err := extractor.ExtractGoFileMeta(testFilepath)
		if err != nil {
			fmt.Printf("extract test file meta occurs error: %v\n", err)
			return
		}
		oldTestFuncMeta := extractor.SearchGoFunctionMeta(testFileMeta, testFuncName)
		if oldTestFuncMeta == nil {
			testFileHandler.WriteString("\n")
			testFileHandler.Write(testFuncByte)
		} else {
			oldUnittestFuncContent := oldTestFuncMeta.Expression()
			unittestFileContent, err := io.ReadAll(testFileHandler)
			if err != nil {
				fmt.Printf("read test file content occurs error: %v\n", err)
				return
			}
			testFileHandler.Close()

			testFileHandler, err := os.OpenFile(testFilepath, os.O_TRUNC|os.O_RDWR, 0644)
			if err != nil {
				fmt.Printf("create test file %v occurs error: %v\n", testFilepath, err)
				return
			}
			replacedFileContent := strings.ReplaceAll(
				strings.ReplaceAll(string(unittestFileContent), "\r", ""),
				strings.ReplaceAll(oldUnittestFuncContent, "\r", ""),
				string(testFuncByte),
			)
			testFileHandler.WriteString(replacedFileContent)
			testFileHandler.Close()
		}
	}

	extractor.GoFmtFile(testFilepath)
}
