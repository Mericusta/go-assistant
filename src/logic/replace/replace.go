package replace

func ReplaceCode(argIdentType, argIdentNames, argRegexp, arg string) {
	// if len(argIdentType) == 0 || (len(argIdentNames) == 0 && len(argRegexp) == 0) || len(arg) == 0 {
	// 	fmt.Printf("not enough options, file %v, meta type %v, meta ident names %v, meta regexp %v\n", arg, argIdentType, argIdentNames, argRegexp)
	// 	return
	// }

	// argIdentNameSlice := strings.Split(argIdentNames, ",")

	// argSlice := strings.Split(arg, ",")
	// if len(argSlice) < 2 {
	// 	fmt.Println("not enough args")
	// 	return
	// }

	// fromFile, toFile := argSlice[0], argSlice[1]

	// fromFileMeta := utility.HandleFileMeta(fromFile)
	// if fromFileMeta == nil {
	// 	return
	// }

	// toFileMeta := utility.HandleFileMeta(toFile)
	// if toFileMeta == nil {
	// 	return
	// }

	// var fromMeta extractor.Meta
	// var toMeta extractor.Meta
	// switch argIdentType {
	// case "func":
	// 	fromMeta = extractor.SearchGoFunctionMeta(fromFileMeta, argIdentNameSlice[0])
	// 	toMeta = extractor.SearchGoFunctionMeta(toFileMeta, argIdentNameSlice[0])
	// case "struct":
	// 	fromMeta = extractor.SearchGoStructMeta(fromFileMeta, argIdentNameSlice[0])
	// 	toMeta = extractor.SearchGoFunctionMeta(toFileMeta, argIdentNameSlice[0])
	// case "interface":
	// 	fromMeta = extractor.SearchGoInterfaceMeta(fromFileMeta, argIdentNameSlice[0])
	// 	toMeta = extractor.SearchGoFunctionMeta(toFileMeta, argIdentNameSlice[0])
	// case "method":
	// 	fromMeta = extractor.SearchGoMethodMeta(fromFileMeta, argIdentNameSlice[0], argIdentNameSlice[1])
	// 	if reflect.ValueOf(fromMeta).IsNil() {
	// 		_meta := extractor.SearchGoInterfaceMeta(fromFileMeta, argIdentNameSlice[0])
	// 		if _meta != nil {
	// 			fromMeta = _meta.SearchMethodDecl(argIdentNameSlice[1])
	// 		}
	// 	}
	// 	toMeta = extractor.SearchGoFunctionMeta(toFileMeta, argIdentNameSlice[0])
	// 	if reflect.ValueOf(toMeta).IsNil() {
	// 		_meta := extractor.SearchGoInterfaceMeta(toFileMeta, argIdentNameSlice[0])
	// 		if _meta != nil {
	// 			toMeta = _meta.SearchMethodDecl(argIdentNameSlice[1])
	// 		}
	// 	}
	// default:
	// }

	// // read old content
	// toReplaceFileContent, err := os.ReadFile(toFile)
	// if err != nil {
	// 	fmt.Printf("read to replace file %v content occurs error: %v\n", toFile, err.Error())
	// 	return
	// }

	// // replace content
	// fromExpression, toExpression := fromMeta.Expression(), toMeta.Expression()
	// replacedFileContent := strings.Replace(string(toReplaceFileContent), toExpression, fromExpression, -1)

	// // write replaced content
	// toReplaceFile, err := os.OpenFile(toFile, os.O_TRUNC|os.O_RDWR, 0644)
	// if err != nil {
	// 	fmt.Printf("open to replace file %v occurs error: %v\n", toFile, err.Error())
	// 	return
	// }
	// _, err = toReplaceFile.WriteString(replacedFileContent)
	// if err != nil {
	// 	fmt.Printf("write to replace file %v occurs error: %v\n", toFile, err.Error())
	// 	return
	// }
}
