package generate

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Mericusta/go-extractor"
)

func handleFileMeta(argFilepath string) *extractor.GoFileMeta {
	handlePathAbs, err := filepath.Abs(argFilepath)
	if err != nil {
		fmt.Printf("get file abs path occurs error: %v\n", err)
		return nil
	}

	handlePathStat, err := os.Stat(handlePathAbs)
	if err != nil {
		fmt.Printf("get file stat occurs error: %v\n", err)
		return nil
	}

	if handlePathStat == nil {
		fmt.Printf("file not exist\n")
		return nil
	}

	if handlePathStat != nil && handlePathStat.IsDir() {
		fmt.Printf("not support dir\n")
		return nil
	}

	gfm, err := extractor.ExtractGoFileMeta(handlePathAbs)
	if gfm == nil || err != nil {
		fmt.Printf("extract file meta occurs error: %v\n", err)
		return nil
	}

	return gfm
}

func handleOutput(unittestFilepath, unittestFilePkg, unittestFuncName string, unittestFuncByte []byte, argMode string) {
	if argMode == "preview" {
		os.Stdout.Write(unittestFuncByte)
		return
	}

	unittestFileStat, err := os.Stat(unittestFilepath)
	if unittestFileStat == nil && !errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("get unit test file %v stat occurs error: %v\n", unittestFilepath, err)
		return
	}

	unittestFileHandler, err := os.OpenFile(unittestFilepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("create unit test file %v occurs error: %v\n", unittestFilepath, err)
		return
	}
	defer unittestFileHandler.Close()

	if unittestFileStat == nil {
		unittestFileHandler.Write(extractor.MakeUnitTestFile(unittestFilePkg, nil))
	}

	switch argMode {
	case "append":
		unittestFileHandler.WriteString("\n")
		unittestFileHandler.Write(unittestFuncByte)
	case "replace":
		unittestFileMeta, err := extractor.ExtractGoFileMeta(unittestFilepath)
		if err != nil {
			fmt.Printf("extract unittest file meta occurs error: %v\n", err)
			return
		}
		oldUnittestFuncMeta := extractor.SearchGoFunctionMeta(unittestFileMeta, unittestFuncName)
		if oldUnittestFuncMeta == nil {
			unittestFileHandler.WriteString("\n")
			unittestFileHandler.Write(unittestFuncByte)
		} else {
			oldUnittestFuncContent := oldUnittestFuncMeta.Expression()
			unittestFileContent, err := io.ReadAll(unittestFileHandler)
			if err != nil {
				fmt.Printf("read unittest file content occurs error: %v\n", err)
				return
			}
			unittestFileHandler.Close()

			unittestFileHandler, err := os.OpenFile(unittestFilepath, os.O_TRUNC|os.O_RDWR, 0644)
			if err != nil {
				fmt.Printf("create unit test file %v occurs error: %v\n", unittestFilepath, err)
				return
			}
			replacedFileContent := strings.ReplaceAll(
				strings.ReplaceAll(string(unittestFileContent), "\r", ""),
				strings.ReplaceAll(oldUnittestFuncContent, "\r", ""),
				string(unittestFuncByte),
			)
			unittestFileHandler.WriteString(replacedFileContent)
			unittestFileHandler.Close()
		}
	}

	extractor.GoFmtFile(unittestFilepath)
}

func GenerateUnittest(argFilepath, argFuncName, argTypeArgs, argMode string) {
	if len(argFilepath) == 0 || len(argFuncName) == 0 {
		fmt.Printf("not enough options, file %v, func %v\n", argFilepath, argFuncName)
		return
	}

	handleFileMeta := handleFileMeta(argFilepath)
	handleFuncMeta := extractor.SearchGoFunctionMeta(handleFileMeta, argFuncName)
	if handleFuncMeta == nil {
		fmt.Printf("can not find func meta\n")
		return
	}

	unittestFuncName, unittestFuncByte := handleFuncMeta.MakeUnitTest(strings.Split(argTypeArgs, ","))
	unittestFilepath := fmt.Sprintf("%v_test.go", strings.Trim(handleFileMeta.Path(), ".go"))

	handleOutput(unittestFilepath, handleFileMeta.PkgName(), unittestFuncName, unittestFuncByte, argMode)
}
