package search

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/Mericusta/go-stp"
)

func SplitLogByKey(argFilepath, argMode, argRegexp string) {
	keyRegexp := regexp.MustCompile(argRegexp)
	if keyRegexp == nil {
		fmt.Printf("compile key regexp failed\n")
		return
	}

	submatchKeyIndex := keyRegexp.SubexpIndex("KEY")
	if submatchKeyIndex == -1 {
		fmt.Printf("regexp not has submatch KEY")
		return
	}

	splitDir := fmt.Sprintf("tmp_split_%v", time.Now().Format("20060102_150405"))
	err := os.Mkdir(splitDir, os.ModePerm)
	if err != nil {
		fmt.Printf("create dir occurs error: %v\n", err)
		return
	}

	var wg sync.WaitGroup
	keyMap := make(map[string]chan string)
	stp.ReadFileLineOneByOne(argFilepath, func(s string, l int) bool {
		if keyRegexp.MatchString(s) {
			submatchSlice := keyRegexp.FindStringSubmatch(s)
			if submatchKeyIndex >= len(submatchSlice) {
				return true
			}
			key := submatchSlice[submatchKeyIndex]
			if _, has := keyMap[key]; !has {
				keyMap[key] = make(chan string, 1024)
				go keyLog(splitDir, key, keyMap[key], &wg)
			}
			keyMap[key] <- s
		}
		return true
	})

	keyCount := len(keyMap)
	fmt.Printf("key count: %v\n", keyCount)
	wg.Add(keyCount)
	for _, keyChan := range keyMap {
		close(keyChan)
	}
	wg.Wait()
}

func keyLog(splitDir, key string, keyChan chan string, wg *sync.WaitGroup) {
	keyFile, err := os.OpenFile(fmt.Sprintf("./%v/%v.log", splitDir, key), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if keyFile == nil || err != nil {
		for s := range keyChan {
			fmt.Printf("key %v: \n", s)
		}
	} else {
		for s := range keyChan {
			l, err := keyFile.WriteString(fmt.Sprintln(s))
			if len(s) != l-1 {
				fmt.Printf("key file %v write string length %v not match string length %v\n", keyFile.Name(), l, len(s))
			}
			if err != nil {
				fmt.Printf("key file %v write string occurs error: %v\n", keyFile.Name(), err)
			}
		}
		keyFile.Close()
		wg.Done()
	}
}
