package utility

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mericusta/go-extractor"
)

func HandleFileMeta(argFilepath string) *extractor.GoFileMeta {
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

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
