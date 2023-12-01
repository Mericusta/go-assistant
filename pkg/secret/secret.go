package secret

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Mericusta/go-stp"
)

const uint8Max = 255

// func MakeASecret(operation, inputFilePath, outputFilePath string) {
// 	if inputFilePath == "" {
// 		fmt.Println("ERROR: input file required")
// 		return
// 	}
// 	if outputFilePath == "" {
// 		fmt.Println("ERROR: output file required")
// 		return
// 	}

// 	inputFile, osOpenError := os.Open(inputFilePath)
// 	if osOpenError != nil {
// 		fmt.Printf("open file %v occurs error, %v\n", inputFile, osOpenError)
// 		return
// 	}

// 	outputFile, osCreateError := os.Create(outputFilePath)
// 	if osCreateError != nil {
// 		fmt.Printf("create %v error, %v", outputFilePath, osCreateError)
// 		return
// 	}

// 	operationResult := true

// 	if operation == "encode" {
// 		// operationResult = encode(inputFile, outputFile)
// 	} else if operation == "decode" {
// 		// operationResult = decode(inputFile, outputFile)
// 	} else {
// 		fmt.Printf("unknown operation %v\n", operation)
// 		return
// 	}

// 	if operationResult {
// 		fmt.Printf("NOTE: %v successfully\n", operation)
// 	} else {
// 		fmt.Printf("NOTE: %v failed\n", operation)
// 	}
// }

func Encode(inputPath, outputPath, outputMode, args, argRegexpExpression string) {
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
		// if err != nil && ! {
		// 	return
		// }
		if outputPathStat.IsDir() {
			fmt.Println("ERROR: output path is directory")
			return
		}
	}

	inputPathStat, err := os.Stat(inputPath)
	if err != nil {
		fmt.Println("ERROR: get input path stat occurs error,", err.Error())
		return
	}

	toEncodeFilePathSlice := make([]string, 0, 8)
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
			fmt.Println("TODO: to encode file", name)
			toEncodeFilePath := filepath.Join(inputPath, name)
			toEncodeFilePathSlice = append(toEncodeFilePathSlice, toEncodeFilePath)
		}
	} else {

	}

	handleReaderMap := make(map[string]io.Reader)
	handleWriterMap := make(map[string]io.Writer)
	for _, toEncodeFilePath := range toEncodeFilePathSlice {
		f, err := os.Open(toEncodeFilePath)
		if f == nil || err != nil {
			fmt.Println("ERROR: open file", toEncodeFilePath, "occurs error", err.Error())
			continue
		}
		handleReaderMap[toEncodeFilePath] = f
		// handleWriterMap[toEncodeFilePath] = bytes.NewBuffer(make([]byte, 0, 1024))
		handleWriterMap[toEncodeFilePath] = os.Stdout
	}

	// for toEncodeFilePath, toEncodeFileReader := range handleReaderMap {
	// 	toEncodeFileWriter := handleWriterMap[toEncodeFilePath]
	// 	err = secretHandler(toEncodeFileReader, toEncodeFileWriter)
	// 	if err != nil {
	// 		fmt.Println("ERROR: encode file", toEncodeFilePath, "occurs error", err.Error())
	// 		continue
	// 	}
	// }

	// switch outputMode {
	// case "replace", "append":

	// default:
	// 	for toEncodeFilePath, toEncodeFileWriter := range handleWriterMap {
	// 		fmt.Println("TODO: file", toEncodeFilePath, toEncodeFileWriter.(*bytes.Buffer).Bytes())
	// 	}
	// }
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
