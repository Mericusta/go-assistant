package operate

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Mericusta/go-assistant/pkg/utility"
	"github.com/Mericusta/go-stp"
)

const (
	MARKDOWN_TABLE_HEAD_LINE_INDEX = iota
	MARKDOWN_TABLE_SPLITTER_LINE_INDEX
)

type markdownTableRow struct {
	line    int
	content string
	values  []string
}

type markdownTable struct {
	head                    *markdownTableRow
	splitter                *markdownTableRow
	keywordColumnIndexSlice []int
	rows                    []*markdownTableRow
}

func (mt *markdownTable) format() {
	fmt.Printf("|%v|\n", strings.Join(mt.head.values, "|"))
	fmt.Printf("|%v|\n", strings.Join(mt.splitter.values, "|"))
	for _, row := range mt.rows {
		fmt.Printf("|%v|\n", strings.Join(row.values, "|"))
	}
}

func (mt *markdownTable) clear() {
	for _, row := range mt.rows {
		keywordColumnIndexArray := stp.NewArray(mt.keywordColumnIndexSlice)
		for columnIndex := range row.values {
			if keywordColumnIndexArray.Includes(columnIndex) {
				continue
			}
			row.values[columnIndex] = ""
		}
	}
}

func OperateMarkdownTable(argOption, argFilepath, args string) {
	markdownFile, err := os.Open(argFilepath)
	utility.PanicError(err)

	argSlice := strings.Split(args, ",")
	if len(argSlice) < 2 {
		utility.PanicError(fmt.Errorf("no table content line"))
	}

	tableBeginLine, err := strconv.Atoi(argSlice[0])
	utility.PanicError(err)
	tableEndLine, err := strconv.Atoi(argSlice[1])
	utility.PanicError(err)
	keywordColumnIndexSlice := make([]int, 0, 4)
	for index := 2; index < len(argSlice); index++ {
		keywordColumn, err := strconv.Atoi(argSlice[index])
		utility.PanicError(err)
		switch {
		case keywordColumn > 0:
			keywordColumnIndexSlice = append(keywordColumnIndexSlice, keywordColumn-1)
		case keywordColumn < 0:
			keywordColumnIndexSlice = append(keywordColumnIndexSlice, keywordColumn)
		}
	}

	table := &markdownTable{}
	table.head = &markdownTableRow{}
	table.splitter = &markdownTableRow{}
	table.rows = make([]*markdownTableRow, 0, tableEndLine-tableBeginLine+1)
	table.keywordColumnIndexSlice = keywordColumnIndexSlice
	err = stp.ReadContentLineOneByOne(markdownFile, func(s string, i int) bool {
		switch {
		case i+1 < tableBeginLine:
			return true
		case i+1 > tableEndLine:
			return false
		}
		tableRow := &markdownTableRow{}
		tableRow.content = s
		tableRow.line = i + 1
		tableRow.values = stp.NewArray(
			strings.Split(strings.TrimSpace(s)[1:len(s)-1], "|"),
		).Map(
			func(v string, i int) string { return strings.TrimSpace(v) },
		).Slice()
		switch i + 1 {
		case tableBeginLine + MARKDOWN_TABLE_HEAD_LINE_INDEX:
			table.head = tableRow
		case tableBeginLine + MARKDOWN_TABLE_SPLITTER_LINE_INDEX:
			table.splitter = tableRow
		default:
			table.rows = append(table.rows, tableRow)
		}
		return true
	})
	utility.PanicError(err)

	columnCount := len(table.head.values)
	stp.NewArray(table.keywordColumnIndexSlice).ForEach(func(v, i int) {
		if v < 0 {
			table.keywordColumnIndexSlice[i] = columnCount + v
		}
	})

	switch argOption {
	case "clear":
		table.clear()
	}

	table.format()
}
