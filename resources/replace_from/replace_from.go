package resources

import (
	"context"
	"fmt"
)

func ReplaceFunc(arg1 context.Context, arg2 map[int64]struct{}, arg3 ...interface{}) {
	fmt.Println("ReplaceFunc From")
}

func OtherNoReplaceFromFunc() {
	fmt.Println("OtherNoReplaceFromFunc")
}
