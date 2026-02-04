package secret

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Mericusta/go-stp"
)

const uint8Max = 255

func Secret(inputPath, outputPath, outputMode, args, argRegexpExpression string) {
	if len(inputPath) == 0 {
		fmt.Println("ERROR: input path required")
		return
	}
	switch outputMode {
	case "replace", "append":
		if len(outputPath) == 0 {
			fmt.Println("ERROR: output path required")
			return
		}
		outputPathStat, err := os.Stat(outputPath)
		notExist := os.IsNotExist(err)
		if !notExist && err != nil {
			fmt.Println("ERROR: get output path stat occurs error,", err.Error())
			return
		}
		if outputPathStat != nil && outputPathStat.IsDir() {
			fmt.Println("ERROR: output path is directory")
			return
		}
	}

	inputPathStat, err := os.Stat(inputPath)
	if err != nil {
		fmt.Println("ERROR: get input path stat occurs error,", err.Error())
		return
	}

	toEncodeContentBytesReader := bytes.NewBuffer(make([]byte, 0, 1024))
	if inputPathStat.IsDir() {
		if len(args) == 0 {
			fmt.Println("ERROR: input path is directory, need file extend name")
			return
		}
		var argRegexp *regexp.Regexp
		if len(argRegexpExpression) > 0 {
			argRegexp = regexp.MustCompile(argRegexpExpression)
			if argRegexp == nil {
				fmt.Println("ERROR: compile arg regexp occurs failed")
				return
			}
		}
		extendedSlice := stp.NewArray(strings.Split(args, ",")).
			Filter(func(v string, i int) bool { return len(v) > 0 }).
			Map(func(v string, i int) string {
				if v[0] == '.' {
					return v
				} else {
					return "." + v
				}
			})
		dirEntry, err := os.ReadDir(inputPath)
		if err != nil {
			fmt.Println("ERROR: read directory occurs error,", err.Error())
			return
		}
		toEncodeFileBytesMap := make(map[string][]byte)
		for _, fe := range dirEntry {
			name := fe.Name()
			if fe.IsDir() || !extendedSlice.Includes(filepath.Ext(name)) {
				continue
			}
			if argRegexp != nil {
				_name, _ := strings.CutSuffix(filepath.Base(name), filepath.Ext(name))
				if !argRegexp.MatchString(_name) {
					continue
				}
			}
			toEncodeFilePath := filepath.Join(inputPath, name)
			contentBytes, err := os.ReadFile(toEncodeFilePath)
			if err != nil {
				fmt.Println("ERROR: read file", toEncodeFilePath, "occurs error", err.Error())
				continue
			}
			if len(contentBytes) == 0 {
				continue
			}
			toEncodeFileBytesMap[toEncodeFilePath] = contentBytes
		}
		toEncodeDirectoryJSON, err := json.Marshal(toEncodeFileBytesMap)
		if err != nil {
			fmt.Println("ERROR: json marshal to encode directory", inputPath, "occurs error", err.Error())
			return
		}
		toEncodeContentBytesReader.Write(toEncodeDirectoryJSON)
	} else {
		contentBytes, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Println("ERROR: read file", inputPath, "occurs error", err.Error())
			return
		}
		if len(contentBytes) == 0 {
			fmt.Println("ERROR: read file", inputPath, "but get empty")
			return
		}
		toEncodeContentBytesReader.Write(contentBytes)
	}

	var toEncodeContentBytesWriter io.Writer
	switch outputMode {
	case "append":
		toEncodeContentBytesWriter, err = os.OpenFile(outputPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("ERROR: open file", outputPath, "occurs error,", err.Error())
			return
		}
	case "replace":
		toEncodeContentBytesWriter, err = os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("ERROR: open file", outputPath, "occurs error,", err.Error())
		}
	default:
		toEncodeContentBytesWriter = os.Stdout
	}

	err = secretHandler(toEncodeContentBytesReader, toEncodeContentBytesWriter)
	if err != nil {
		fmt.Println("ERROR: handle content bytes occurs error", err.Error())
		return
	}
}

func secretHandler(toHandleContentReader io.Reader, resultContentWriter io.Writer) (err error) {
	// read
	contentBytes, err := io.ReadAll(toHandleContentReader)
	if err != nil {
		return
	}
	l := len(contentBytes)
	if l == 0 {
		err = fmt.Errorf("to handle content is empty")
		return
	}

	// encode
	for index, toEncryptByte := range contentBytes {
		contentBytes[index] = toEncryptByte ^ uint8Max
	}

	// write
	writer := bufio.NewWriter(resultContentWriter)
	_, err = writer.Write(contentBytes)
	if err != nil {
		return
	}
	err = writer.Flush()

	return
}
