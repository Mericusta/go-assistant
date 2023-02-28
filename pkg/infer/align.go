package infer

import (
	"fmt"
	"math"
	"strings"

	"github.com/Mericusta/go-assistant/pkg/utility"
	"github.com/Mericusta/go-extractor"
)

func InferTheOptimalLayoutOfStructMemory(argPlatform int, argFilepath, argStructName string, argAllocationPreview, argProcess bool) {
	gfm := utility.HandleFileMeta(argFilepath)
	if gfm == nil {
		return
	}

	gsm := extractor.SearchGoStructMeta(gfm, argStructName)
	if gsm == nil {
		fmt.Printf("can not find struct meta\n")
		return
	}

	gsm.PrintAST()

	unknownDescMemberTypeMap := make(map[string]int)
	originStructMemberDesc := make([]*structMemberSizeAlignTypeDesc, 0)
	for _, member := range gsm.Members() {
		gvm := gsm.SearchMemberMeta(member)
		if gvm == nil {
			continue
		}
		gvm.PrintAST()
		_, underLyingTypeStr, _ := gvm.Type()
		sm, has := x64TypeSizeMap[underLyingTypeStr]
		if has {
			originStructMemberDesc = append(originStructMemberDesc, sm)
		} else {
			unknownDescMemberTypeMap[underLyingTypeStr]++
		}
	}

	compilerDefaultAlign := 8
	if argPlatform == 32 {
		compilerDefaultAlign = 4
	}

	minWasting := math.MaxInt16
	allocation, wasting, _, allocationPreview := calculateStructMemoryAllocation(compilerDefaultAlign, originStructMemberDesc, argProcess)
	if minWasting < wasting {
		wasting = minWasting
	}
	if len(unknownDescMemberTypeMap) > 0 {
		fmt.Printf("the optimal layout of known struct allocation is: %v, memory utilization %.2f%%\n", allocation, float64(allocation-wasting)/float64(allocation)*100)
		for unknownDescMemberType := range unknownDescMemberTypeMap {
			fmt.Printf("unknown desc member type: %v\n", unknownDescMemberType)
		}
		if argAllocationPreview {
			fmt.Printf("allocation preview: %v\n", allocationPreview)
		}
	} else {
		fmt.Printf("the optimal layout of struct allocation is: %v, memory utilization %.2f%%\n", allocation, float64(allocation-wasting)/float64(allocation)*100)
		if argAllocationPreview {
			fmt.Printf("allocation preview: %v\n", allocationPreview)
		}
	}
}

var x64TypeSizeMap = map[string]*structMemberSizeAlignTypeDesc{
	"bool":       {size: 1, align: 1, desc: "bool"},
	"int8":       {size: 1, align: 1, desc: "int8"},
	"int16":      {size: 2, align: 2, desc: "int16"},
	"int32":      {size: 4, align: 4, desc: "int32"},
	"int64":      {size: 8, align: 8, desc: "int64"},
	"uint8":      {size: 1, align: 1, desc: "uint8"},
	"uint16":     {size: 2, align: 2, desc: "uint16"},
	"uint32":     {size: 4, align: 4, desc: "uint32"},
	"uint64":     {size: 8, align: 8, desc: "uint64"},
	"float32":    {size: 4, align: 4, desc: "float32"},
	"float64":    {size: 8, align: 8, desc: "float64"},
	"complex64":  {size: 8, align: 8, desc: "complex64"},
	"complex128": {size: 16, align: 8, desc: "complex128"},
	"int":        {size: 8, align: 8, desc: "int"},
	"uint":       {size: 8, align: 8, desc: "uint"},
	"uintptr":    {size: 8, align: 8, desc: "uintptr"},
	"pointer":    {size: 8, align: 8, desc: "pointer"},
	"map":        {size: 8, align: 8, desc: "map"},
	"func":       {size: 8, align: 8, desc: "func"},
	"chan":       {size: 8, align: 8, desc: "chan"},
	"interface":  {size: 16, align: 8, desc: "interface"},
	"string":     {size: 16, align: 8, desc: "string"},
	"array":      {size: 24, align: 8, desc: "array"},
}

// ┌────────────────────────────────────┬─────────────────┬──────────────────┬─────────────────┬──────────────────┐
// │type                                │x64 platform size│x64 platform align│x32 platform size│x32 platform align│
// ├────────────────────────────────────┼─────────────────┼──────────────────┼─────────────────┼──────────────────┤
// │bool                                │1 byte           │1 byte            │1 byte           │1 byte            │
// ├────────────────────────────────────┼─────────────────┼──────────────────┼─────────────────┼──────────────────┤
// │intN, uintN, floatN, complexN       │N/8 byte         │N/8 byte          │N/4 byte         │N/4 byte          │
// ├────────────────────────────────────┼─────────────────┼──────────────────┼─────────────────┼──────────────────┤
// │int, uint, uintptr                  │8 byte           │8 byte            │4 byte           │4 byte            │
// ├────────────────────────────────────┼─────────────────┼──────────────────┼─────────────────┼──────────────────┤
// │pointer, map, func, chan            │8 byte           │8 byte            │4 byte           │4 byte            │
// ├────────────────────────────────────┼─────────────────┼──────────────────┼─────────────────┼──────────────────┤
// │interface (tab, data or _type, data)│16 byte          │8 byte            │8 byte           │4 byte            │
// ├────────────────────────────────────┼─────────────────┼──────────────────┼─────────────────┼──────────────────┤
// │string (ptr, len)                   │16 byte          │8 byte            │8 byte           │4 byte            │
// ├────────────────────────────────────┼─────────────────┼──────────────────┼─────────────────┼──────────────────┤
// │array (ptr, len, cap)               │24 byte          │8 byte            │12 byte          │4 byte            │
// └────────────────────────────────────┴─────────────────┴──────────────────┴─────────────────┴──────────────────┘

type structMemberSizeAlignTypeDesc struct {
	size  int
	align int
	desc  string
}

func calculateStructMemoryAllocation(compilerDefaultAlign int, smDesc []*structMemberSizeAlignTypeDesc, process bool) (int, int, []int, string) {
	memberLen := len(smDesc)
	memberAllocation := make([]int, memberLen)
	b := strings.Builder{}
	allocation := 0
	wasting := 0
	maxMemberSize := 0

	// allocate begin
	b.WriteRune('[')
	b.WriteRune(' ')

	// according to rule 1
	for i, sm := range smDesc {
		if process {
			fmt.Printf("struct No.%v member, size %v, align %v, type %v\n", i+1, sm.size, sm.align, sm.desc)
		}

		min := int(math.Min(float64(sm.size), float64(sm.align)))
		if process {
			fmt.Printf("No.%v member, min(size, align) = %v\n", i, min)
		}
		if allocation%min == 0 {
			if process {
				fmt.Printf("allocate offset = %v, offset mod min(size, align) == 0, no need fill\n", allocation)
			}
		} else {
			fillByte := min - allocation
			if allocation > min {
				v := min
				for v <= allocation {
					v += min
				}
				fillByte = v - allocation
			}
			if process {
				fmt.Printf("allocate offset = %v, offset mod min(size, align) != 0, need fill %v byte\n", allocation, fillByte)
			}
			wasting += fillByte
			// allocate fill byte
			allocation += fillByte
			memberAllocation[i] += fillByte
			for i := 0; i < fillByte; i++ {
				b.WriteRune('_')
				b.WriteRune(' ')
			}
		}

		// allocate member byte
		allocation += sm.size
		memberAllocation[i] += sm.size
		for i := 0; i < sm.size; i++ {
			b.WriteRune('0')
			b.WriteRune(' ')
		}
		if maxMemberSize < sm.size {
			maxMemberSize = sm.size
		}

		// allocate end
		if i != memberLen-1 {
			b.WriteRune('|')
			b.WriteRune(' ')
		}

		// output allocation
		if process {
			fmt.Printf("struct No.%v member memory allocation: %v\n", i+1, b.String())
			fmt.Println()
		}
	}
	if process {
		fmt.Printf("according to rule 1, struct member memory allocation size: %v\n", allocation)
		fmt.Printf("allocation: %v\n", b.String())
	}

	// according to rule 2
	if process {
		fmt.Printf("struct max member size = %v\n", maxMemberSize)
		fmt.Printf("compiler default align = %v\n", compilerDefaultAlign)
	}
	min := int(math.Min(float64(maxMemberSize), float64(compilerDefaultAlign)))
	if process {
		fmt.Printf("min(maxMemberSize, compilerDefaultAlign) = %v\n", min)
	}
	if allocation%min == 0 {
		if process {
			fmt.Printf("struct no need fill\n")
		}
	} else {
		fillByte := min - allocation
		if allocation > min {
			v := min
			for v <= allocation {
				v += min
			}
			fillByte = v - allocation
		}
		if process {
			fmt.Printf("struct need fill %v byte\n", fillByte)
		}
		wasting += fillByte
		// allocate fill byte
		allocation += fillByte
		b.WriteRune('|')
		b.WriteRune(' ')
		for i := 0; i < fillByte; i++ {
			b.WriteRune('_')
			b.WriteRune(' ')
		}
	}

	// allocate end
	b.WriteRune(']')

	if process {
		fmt.Printf("according to rule 2, struct member memory allocation size: %v\n", allocation)
		fmt.Printf("allocation: %v\n", b.String())
		fmt.Printf("wasting offset: %v, memory utilization %.2f%%\n", wasting, float64(allocation-wasting)/float64(allocation)*100)
	}

	return allocation, wasting, memberAllocation, b.String()
}
