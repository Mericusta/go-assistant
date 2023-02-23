package generate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mericusta/go-extractor"
)

func GenerateBenchmark(argFilepath, argFuncName, argMode string, fromUnittest bool) {
	handlePathAbs, err := filepath.Abs(argFilepath)
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

	if handlePathStat != nil && handlePathStat.IsDir() {
		fmt.Printf("not support dir\n")
		return
	}

	gfm, err := extractor.ExtractGoFileMeta(handlePathAbs)
	if gfm == nil || err != nil {
		fmt.Printf("extract file meta occurs error: %v\n", err)
		return
	}

	gfm.PrintAST()
}
