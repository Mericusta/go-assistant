package search

func SearchFuncByRegexp(argFilepath, argMode, argRegexp string) {
	// if len(argRegexp) == 0 {
	// 	fmt.Printf("empty func regexp\n")
	// 	return
	// }

	// funcRegexp := regexp.MustCompile(argRegexp)
	// if funcRegexp == nil {
	// 	fmt.Printf("compile func regexp failed\n")
	// 	return
	// }

	// submatchKeyIndex := funcRegexp.SubexpIndex("KEY")
	// if submatchKeyIndex == -1 {
	// 	fmt.Printf("regexp not has submatch KEY")
	// 	return
	// }

	// fileContent, err := os.ReadFile(argFilepath)
	// if err != nil {
	// 	panic(err)
	// }

	// handleFileMeta := utility.HandleFileMeta(argFilepath)
	// if handleFileMeta == nil {
	// 	return
	// }

	// pkg := "game"
	// switch argFilepath {
	// case "cross_parking.go", "cross.go", "cross_gm.go":
	// 	pkg = "cross"
	// case "chat.go":
	// 	pkg = "chat"
	// case "qualifying.go":
	// 	pkg = "qualifying"
	// case "battle.go", "clan_war.go", "clan_war_gm.go":
	// 	pkg = "clan_war"
	// case "friend.go":
	// 	pkg = "friend"
	// case "replay.go":
	// 	pkg = "replay"
	// case "cache.go", "cache_clan.go", "cache_player.go":
	// 	pkg = "cache"
	// case "leaderboard.go":
	// 	pkg = "leaderboard"
	// case "showdown.go":
	// 	pkg = "showdown"
	// default:
	// }

	// for _, matchFunc := range funcRegexp.FindAllSubmatch(fileContent, -1) {
	// 	handleFuncMeta := extractor.SearchGoFunctionMeta(handleFileMeta, string(matchFunc[submatchKeyIndex]))
	// 	funcExpression := handleFuncMeta.Expression()
	// 	lineSlice := strings.Split(funcExpression, "\n")
	// 	body := handleFuncMeta.Body().Expression()
	// 	if argMode == "s2c_empty" {
	// 		if len(lineSlice) == 8 && strings.TrimSpace(lineSlice[5]) == "_ = pbmsg" && strings.TrimSpace(lineSlice[6]) == "return nil" {
	// 			body = "{ return nil }"
	// 		} else {
	// 			continue
	// 		}
	// 	} else if argMode == "s2c_body" {
	// 		if len(lineSlice) == 8 && strings.TrimSpace(lineSlice[5]) == "_ = pbmsg" && strings.TrimSpace(lineSlice[6]) == "return nil" {
	// 			continue
	// 		}
	// 	} else if argMode == "c2s_empty" {
	// 		if len(lineSlice) == 9 && strings.TrimSpace(lineSlice[7]) == "return nil" {
	// 			body = fmt.Sprintf(`{
	// 				c2sMsg := &%v.%v{}
	// 				c.Robot().SendMsg(c2sMsg)
	// 				return nil
	// 			}`, pkg, handleFuncMeta.FunctionName())
	// 		} else {
	// 			continue
	// 		}
	// 	} else if argMode == "c2s_body" {
	// 		if len(lineSlice) == 9 && strings.TrimSpace(lineSlice[7]) == "return nil" {
	// 			continue
	// 		}
	// 	}
	// 	if len(argMode) >= 3 {
	// 		if argMode[:3] == "s2c" {
	// 			fmt.Printf("func %v(c inObj.IRobotContext, s2cMsg *%v.%v) error %v\n", handleFuncMeta.FunctionName(), pkg, handleFuncMeta.FunctionName(), body)
	// 		} else if argMode[:3] == "c2s" {
	// 			fmt.Printf("func %v(c inObj.IRobotContext, args ...any) error %v\n", handleFuncMeta.FunctionName(), body)
	// 		}
	// 	}
	// }
}
