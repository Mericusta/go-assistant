package resources

import (
	"context"
	"fmt"
)

func ReplaceFunc(argA map[int64]struct{}, argB context.Context, argC ...interface{}) {
	fmt.Println("ReplaceFunc To")
}

func OtherNoReplaceToFunc() {
	fmt.Println("OtherNoReplaceToFunc")
}
