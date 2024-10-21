package services

import (
	"fmt"
	"strings"
	"time"

	game "github.com/FlyDragonGO/DoubleMerge_Proto/AutoProtoOutput/go/game"
	"github.com/FlyDragonGO/DoubleMerge_Proto/AutoProtoOutput/go/models"
	"github.com/FlyDragonGO/DoubleMerge_Server/common/logger"
	"github.com/FlyDragonGO/DoubleMerge_Server/framework"
	"github.com/FlyDragonGO/DoubleMerge_Server/internal/game/event"
	playerMgr "github.com/FlyDragonGO/DoubleMerge_Server/internal/game/logic/player_mgr"
	obj "github.com/FlyDragonGO/DoubleMerge_Server/internal/game/player/object"
	"github.com/Mericusta/go-stp"
	"golang.org/x/net/context"
)

func (s *Service) ChangePlayerName(p0 Context, p1 *CChangePlayerName) (*SChangePlayerName, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("ChangePlayerName on goroutine %v\n", goroutineID)
	s.player.IdentifyComponent().SetName(req.GetName())
	res = &game.SChangePlayerName{}
	res.Result = 0
	res.Name = s.player.IdentifyComponent().GetName()
	return
}

func (s *Service) CheckTime(p0 Context, p1 *CCheckTime) (*SCheckTime, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("CheckTime on goroutine %v\n", goroutineID)
	res = &game.SCheckTime{}
	res.Timestamp = time.Now().UnixMilli()
	res.Timezone = 0
	return res, nil
}

func (s *Service) GetCumulativeLoginDays(p0 Context, p1 *CGetCumulativeLoginDays) (*SGetCumulativeLoginDays, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetCumulativeLoginDays on goroutine %v\n", goroutineID)
	res = &game.SGetCumulativeLoginDays{}
	res.CumulativeLoginDays = s.player.StateComponent().GetCumulativeLoginDays()
	return res, nil
}

func (s *Service) GetDiamondCount(p0 Context, p1 *CGetDiamondCount) (*SGetDiamondCount, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetDiamondCount on goroutine %v\n", goroutineID)
	res = &game.SGetDiamondCount{}
	res.DiamondCount = s.player.ResourceComponent().GetDiamondCount()
	return res, nil
}

func (s *Service) GetEnergyData(p0 Context, p1 *CGetEnergyData) (*SGetEnergyData, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetEnergyData on goroutine %v\n", goroutineID)
	res = &game.SGetEnergyData{}
	res.EnergyCount = s.player.ResourceComponent().GetEnergyCount()
	res.NextUpdateTs = s.player.ResourceComponent().GetEnergyNextUpdateTimestamp()
	return res, nil
}

func (s *Service) GetGoldCount(p0 Context, p1 *CGetGoldCount) (*SGetGoldCount, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetGoldCount on goroutine %v\n", goroutineID)
	res = &game.SGetGoldCount{}
	res.GoldCount = s.player.ResourceComponent().GetGoldCount()
	return res, nil
}

func (s *Service) GetLevelExp(p0 Context, p1 *CGetLevelExp) (*SGetLevelExp, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetLevelExp on goroutine %v\n", goroutineID)
	res = &game.SGetLevelExp{}
	res.Level, res.Exp = s.player.BasicComponent().GetLevel(), s.player.BasicComponent().GetExp()
	return res, nil
}

func (s *Service) GetMaterials(p0 Context, p1 *CGetMaterials) (*SGetMaterials, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetMaterials on goroutine %v\n", goroutineID)
	res = &game.SGetMaterials{}
	res.Materials = make(map[int64]int64)
	s.player.ResourceComponent().GetMaterials()
	return res, nil
}

func (s *Service) GetOrderList(p0 Context, p1 *CGetOrderList) (*SGetOrderList, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetOrderList on goroutine %v\n", goroutineID)
	res = &game.SGetOrderList{}
	res.OrderList = make(map[int64]*models.PBSingleOrder)
	return res, nil
}

func (s *Service) GetPlayerName(p0 Context, p1 *CGetPlayerName) (*SGetPlayerName, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetPlayerName on goroutine %v\n", goroutineID)
	res = &game.SGetPlayerName{}
	res.Name = s.player.IdentifyComponent().GetName()
	return
}

func (s *Service) GetProfile(p0 Context, p1 *CGetProfile) (*SGetProfile, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetProfile on goroutine %v\n", goroutineID)
	res = &game.SGetProfile{}
	res.ProfilePicture, res.ProfilePictureFrame = s.player.BasicComponent().GetProfilePicture(), s.player.BasicComponent().GetProfilePictureFrame()
	return res, nil
}

func (s *Service) GetResults(p0 Context, p1 *CGetResults) (*SGetResults, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetResults on goroutine %v\n", goroutineID)
	res = &game.SGetResults{}
	res.Results = s.player.ResourceComponent().GetResults()
	return
}

func (s *Service) Login(p0 Context, p1 *CLogin) (*SLogin, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("Login on goroutine %v\n", goroutineID)
	ctx := c.(framework.IContext)
	p, err := playerMgr.Instance().GetPlayer(ctx.GetPlayerId())
	if err != nil {
		return nil, err
	}
	if p == nil {
		p, err = obj.NewPlayer(ctx.GetPlayerId())
		if err != nil {
			return nil, err
		}
		if err = playerMgr.Instance().PutPlayer(p); err != nil {
			return nil, err
		}
		err = p.SavePlayerFull()
		if err != nil {
			return nil, err
		}
	}
	s.player = p
	p.GetModelPlayerInfo().SetLastLoginTime(time.Now().UnixMilli())
	ctx.PublishEventLocal(&event.LoginEvent{Pid: s.playerId})
	res := &game.SLogin{}
	res.Model = s.player.ToPB()
	logger.Info("player login finished！", "playerId", ctx.GetPlayerId())
	logger.Debugf("player login at Goroutine %v", func() uint64 {
		gID, _ := stp.GoroutineID()
		return gID
	}())
	return res, nil
}

func (s *Service) Logout(p0 Context, p1 *CLogout) (*SLogout, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("Logout on goroutine %v\n", goroutineID)
	return nil, nil
}

func (s *Service) MergeGenerator(p0 Context, p1 *CMergeGenerator) (*SMergeGenerator, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("MergeGenerator on goroutine %v\n", goroutineID)
	return nil, nil
}

func (s *Service) PlayerOffline(p0 Context, p1 *CPlayerOffline) (*SPlayerOffline, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("PlayerOffline on goroutine %v\n", goroutineID)
	ctx := c.(framework.IContext)
	_, err := playerMgr.Instance().GetPlayer(ctx.GetPlayerId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Service) SyncModel(p0 Context, p1 *CSyncModel) (*SSyncModel, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("SyncModel on goroutine %v\n", goroutineID)
	for _, syncItem := range req.GetItems() {
		err = s.player.SyncFromClient(strings.Split(syncItem.Id, "."), 0, syncItem)
		if err != nil {
			fmt.Printf("player %v SyncFromClient occurs error %v", s.player, err.Error())
			continue
		}
	}
	s.player.SavePlayerFull()
	res = &game.SSyncModel{}
	res.TestRes = "this is a test response"
	return
}

func (s *Service) UnlockMaterialSlot(p0 Context, p1 *CUnlockMaterialSlot) (*SUnlockMaterialSlot, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("UnlockMaterialSlot on goroutine %v\n", goroutineID)
	return nil, nil
}

func (s *Service) mustEmbedUnimplementedGameServiceServer() {
}
