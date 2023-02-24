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

func handleOutput(testFilepath, testFilePkg, testFuncName string, testFuncByte []byte, argMode string) {
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
