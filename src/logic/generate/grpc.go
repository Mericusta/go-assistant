package generate

import (
	"fmt"
	"go/ast"
	"log"

	"github.com/Mericusta/go-extractor"
)

func GenerateGRPC_CS(argFilePath, argInterfaceIdent string) {
	goPackageMeta, err := extractor.ExtractGoProjectMeta(argFilePath, nil)
	if err != nil {
		log.Panic(err)
	}

	for _, gpm := range goPackageMeta.PackageMap() {
		gim := gpm.SearchInterfaceMeta(argInterfaceIdent)
		if gim == nil {
			log.Panicf("can not find %v meta", argInterfaceIdent)
		}

		for method, gimm := range gim.MethodMetaMap() {
			params := gimm.Params()
			if len(params) < 2 {
				continue
			}
			paramC2SMsg := params[1]
			c2sFuncMeta := extractor.MakeUpFuncMeta(
				paramC2SMsg.TypeIdent(),
				[]*extractor.GoVarMeta[*ast.Field]{
					extractor.MakeUpVarMeta("ctx", "inObj.IRobotContext"),
					extractor.MakeUpVarMeta("args", "[]any"),
				},
				nil,
			)
			fmt.Printf("\nc2s method: %v\n|\n%v\n|\n", method, c2sFuncMeta.Format())

			returns := gimm.Returns()
			if len(returns) < 2 {
				continue
			}
			returnS2CMsg := returns[0]
			s2cFuncMeta := extractor.MakeUpFuncMeta(
				returnS2CMsg.TypeIdent(),
				[]*extractor.GoVarMeta[*ast.Field]{
					extractor.MakeUpVarMeta("ctx", "inObj.IRobotContext"),
					extractor.MakeUpVarMeta("s2cMsg", returnS2CMsg.TypeExpression()),
				},
				[]*extractor.GoVarMeta[*ast.Field]{
					extractor.MakeUpVarMeta("", "error"),
				},
			)
			fmt.Printf("\ns2c method: %v\n|\n%v\n|\n", method, s2cFuncMeta.Format())

			// func CCrossParkGetBriefInfo(c inObj.IRobotContext, args ...any) {
			// 	c2sMsg := &cross.CCrossParkGetBriefInfo{}
			// 	c.Robot().SetTickTreeBehaviorAtUUID(c.Robot().SendMsg(c2sMsg))
			// }

			// func SCrossParkJoinDefend(c inObj.IRobotContext, s2cMsg *cross.SCrossParkJoinDefend) error {
			// 	return nil
			// }

		}
	}
}
