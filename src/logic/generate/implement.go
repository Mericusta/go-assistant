package generate

import "fmt"

func GenerateImplement(argFilepath, argInterfaceName, argMetaIdent, argTypeArgs, args string) {
	if len(argFilepath) == 0 || len(argInterfaceName) == 0 || len(argMetaIdent) == 0 || len(argTypeArgs) == 0 || len(args) == 0 {
		fmt.Printf("not enough option, file %v, interface %v, receiver ident %v, receiver type %v, generate file %v", argFilepath, argInterfaceName, argMetaIdent, argTypeArgs, args)
		return
	}

	// 	interfaceFilePath, interfaceName := argFilepath, argInterfaceName
	// 	implementReceiverIdent, implementReceiverType := argMetaIdent, argTypeArgs
	// 	implementOutputFile := args

	// 	handleFileMeta := utility.HandleFileMeta(interfaceFilePath)
	// 	if handleFileMeta == nil {
	// 		return
	// 	}

	// 	gim := extractor.SearchGoInterfaceMeta(handleFileMeta, interfaceName)
	// 	if gim == nil {
	// 		fmt.Printf("can not find interface %v\n", argInterfaceName)
	// 		return
	// 	}

	// 	toImplementMethodDeclSlice := make([]*ast.Field, 0)
	// 	gim.ForeachMethodDecl(func(f *ast.Field) bool {
	// 		toImplementMethodDeclSlice = append(toImplementMethodDeclSlice, f)
	// 		return true
	// 	})
	// 	if len(toImplementMethodDeclSlice) == 0 {
	// 		fmt.Printf("no method needs to implement")
	// 		return
	// 	}
	// 	toImplementMethodMetaSlice := make([]*extractor.GoInterfaceMethodMeta, 0, len(toImplementMethodDeclSlice))
	// 	for _, toImplementMethodDecl := range toImplementMethodDeclSlice {
	// 		gimm := gim.SearchMethodDecl(toImplementMethodDecl.Names[0].Name)
	// 		if gimm == nil {
	// 			fmt.Printf("can not find go interface method %v meta\n", toImplementMethodDecl.Names[0].Name)
	// 			continue
	// 		}
	// 		toImplementMethodMetaSlice = append(toImplementMethodMetaSlice, gimm)
	// 	}
	// 	implementMethodNameSlice := make([]string, 0, len(toImplementMethodMetaSlice))
	// 	implementMethodMetaSlice := make([]*extractor.GoFunctionMeta, 0, len(toImplementMethodMetaSlice))
	// 	for _, toImplementMethodMeta := range toImplementMethodMetaSlice {
	// 		implementMethodName, implementMethodMeta := toImplementMethodMeta.MakeImplementMethodMeta(implementReceiverIdent, implementReceiverType)
	// 		implementMethodNameSlice = append(implementMethodNameSlice, implementMethodName)
	// 		implementMethodMetaSlice = append(implementMethodMetaSlice, implementMethodMeta)
	// 	}

	// handleOutputImplement(implementOutputFile, implementReceiverType, implementMethodNameSlice, implementMethodMetaSlice)
}

// func handleOutputImplement(implementOutputFile, implementReceiverType string, implementMethodNameSlice []string, implementMethodMetaSlice []*extractor.GoFunctionMeta) {
// 	implementOutputFileState, err := os.Stat(implementOutputFile)
// 	if implementOutputFileState == nil && !errors.Is(err, fs.ErrNotExist) {
// 		fmt.Printf("get test file %v stat occurs error: %v\n", implementOutputFile, err)
// 		return
// 	}

// 	implementOutputFileHandler, err := os.OpenFile(implementOutputFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
// 	if err != nil {
// 		fmt.Printf("create test file %v occurs error: %v\n", implementOutputFile, err)
// 		return
// 	}
// 	defer implementOutputFileHandler.Close()

// 	if implementOutputFileState == nil {
// 		for index := range implementMethodNameSlice {
// 			implementOutputFileHandler.WriteString(implementMethodMetaSlice[index].Format())
// 			implementOutputFileHandler.WriteString("\n")
// 		}
// 	} else {
// 		implementOutputFileMeta := utility.HandleFileMeta(implementOutputFile)
// 		if implementOutputFileMeta != nil {
// 			implementOutputFileContentBytes, err := os.ReadFile(implementOutputFile)
// 			if err != nil {
// 				panic(err)
// 			}
// 			implementOutputFileContent := string(implementOutputFileContentBytes)
// 			implementOutputMethodExpressionMap := make(map[string]string)
// 			implementOutputMethodMetaMap := make(map[string]*extractor.GoMethodMeta)
// 			for _, implementMethodName := range implementMethodNameSlice {
// 				gmm := extractor.SearchGoMethodMeta(implementOutputFileMeta, implementReceiverType, implementMethodName)
// 				if gmm == nil {
// 					continue
// 				}

// 				implementOutputMethodExpressionMap[implementMethodName] = gmm.Expression()
// 				implementOutputMethodMetaMap[implementMethodName] = gmm
// 			}
// 			for index, implementMethodName := range implementMethodNameSlice {
// 				gmm := implementOutputMethodMetaMap[implementMethodName]
// 				implementMethodMeta := implementMethodMetaSlice[index]
// 				// fmt.Printf("method %v\n", implementMethodName)
// 				// fmt.Printf("implement \n|%v|\n", implementMethodMeta.Format())
// 				if gmm == nil {
// 					// fmt.Printf("append %v\n", implementMethodName)
// 					implementOutputFileContent += "\n\n" + implementMethodMeta.Format()
// 					continue
// 				} else {
// 					// fmt.Printf("replaced \n|%v|\n", implementOutputFileContent)
// 					implementMethodMeta.ReplaceBody(gmm.Body())
// 					implementOutputFileContent = strings.ReplaceAll(implementOutputFileContent, implementOutputMethodExpressionMap[implementMethodName], implementMethodMeta.Format())
// 				}

// 			}
// 			stp.WriteFileByOverwriting(implementOutputFile, func(b []byte) ([]byte, error) {
// 				return []byte(implementOutputFileContent), nil
// 			})
// 			extractor.GoFmtFile(implementOutputFile)
// 		} else {
// 			panic("implement output file meta is nil")
// 		}
// 	}

// }
