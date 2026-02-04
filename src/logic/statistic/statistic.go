package statistic

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/Mericusta/go-stp"
)

func StatisticByKey(argFilepath, argSplitRegexp, argStatisticsRegexp, args string) {
	splitKeyRegexp := regexp.MustCompile(argSplitRegexp)
	if splitKeyRegexp == nil {
		fmt.Printf("compile split key regexp failed\n")
		return
	}
	statisticsKeyRegexp := regexp.MustCompile(argStatisticsRegexp)
	if statisticsKeyRegexp == nil {
		fmt.Printf("compile statistics key regexp failed\n")
		return
	}
	submatchSplitKeyIndex := splitKeyRegexp.SubexpIndex("KEY")
	if submatchSplitKeyIndex == -1 {
		fmt.Printf("split key regexp not has submatch KEY")
		return
	}
	submatchStatisticsIndex := statisticsKeyRegexp.SubexpIndex("KEY")
	if submatchStatisticsIndex == -1 {
		fmt.Printf("statistics regexp not has submatch KEY")
		return
	}

	type keyCount struct {
		splitKey        string
		statisticsCount int
	}

	keyMap := make(map[string]*keyCount)
	stp.ReadFileLineOneByOne(argFilepath, func(s string, l int) bool {
		if splitKeyRegexp.MatchString(s) {
			splitKeySubmatchSlice := splitKeyRegexp.FindStringSubmatch(s)
			if submatchSplitKeyIndex >= len(splitKeySubmatchSlice) {
				return true
			}
			splitKey := splitKeySubmatchSlice[submatchSplitKeyIndex]
			if len(splitKey) == 0 {
				return true
			}
			statisticsKeySubmatchSlice := statisticsKeyRegexp.FindStringSubmatch(s)
			if submatchStatisticsIndex >= len(statisticsKeySubmatchSlice) {
				return true
			}
			statisticsKey := statisticsKeySubmatchSlice[submatchStatisticsIndex]
			if len(statisticsKey) == 0 {
				return true
			}
			if _, has := keyMap[splitKey]; !has {
				keyMap[splitKey] = &keyCount{
					splitKey: splitKey,
				}
			}
			keyMap[splitKey].statisticsCount++
		}
		return true
	})

	if args == "count" {
		fmt.Printf("- count:\n")
		total := 0
		for _, keyCount := range keyMap {
			// fmt.Printf("\t- key: %v, count: %v\n", keyCount.splitKey, keyCount.statisticsCount)
			total += keyCount.statisticsCount
		}
		fmt.Printf("\t- total count: %v\n", total)
	} else {
		topN, err := strconv.ParseInt(args, 10, 64)
		if err != nil {
			panic(err)
		}

		keyCountSlice := make([]*keyCount, 0, 64)
		for _, kc := range keyMap {
			keyCountSlice = append(keyCountSlice, kc)
		}
		sort.Slice(keyCountSlice, func(i, j int) bool {
			return keyCountSlice[i].statisticsCount < keyCountSlice[j].statisticsCount
		})
		keyCountLen := len(keyCountSlice)
		fmt.Printf("- top %v:\n", topN)
		for index := keyCountLen - 1; index >= keyCountLen-int(topN) && index >= 0; index-- {
			fmt.Printf("\t- key: %v, count: %v\n", keyCountSlice[index].splitKey, keyCountSlice[index].statisticsCount)
		}
	}
}
