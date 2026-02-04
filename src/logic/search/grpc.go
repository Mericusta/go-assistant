package search

import (
	"fmt"

	"github.com/Mericusta/go-extractor"
)

func SearchGRPCServiceMethods(argFilepath, argInterfaceIdent string) {
	goProjectMeta, err := extractor.ExtractGoProjectMeta(argFilepath, nil)
	if goProjectMeta == nil || err != nil {
		panic(err)
	}

	for _, gpm := range goProjectMeta.PackageMap() {
		gim := gpm.SearchInterfaceMeta(argInterfaceIdent)
		if gim == nil {
			fmt.Printf("can not find interface %v\n", argInterfaceIdent)
			return
		}
		fmt.Printf("gim.Expression = |\n%v\n|\n", gim.Expression())
	}
}
