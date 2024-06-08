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

// TODO: Player 层的 handler
func (s *Service) ChangePlayerName(c context.Context, req *game.CChangePlayerName) (res *game.SChangePlayerName, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("ChangePlayerName on goroutine %v\n", goroutineID)
	s.player.IdentifyComponent().SetName(req.GetName())
	res = &game.SChangePlayerName{}
	res.Result = 0
	res.Name = s.player.IdentifyComponent().GetName()
	return
}

// TODO: Player 层的 handler
func (s *Service) CheckTime(iContext context.Context, req *game.CCheckTime) (res *game.SCheckTime, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("CheckTime on goroutine %v\n", goroutineID)
	res = &game.SCheckTime{}
	res.Timestamp = time.Now().UnixMilli()
	res.Timezone = 0
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetCumulativeLoginDays(iContext context.Context, req *game.CGetCumulativeLoginDays) (res *game.SGetCumulativeLoginDays, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetCumulativeLoginDays on goroutine %v\n", goroutineID)
	res = &game.SGetCumulativeLoginDays{}
	res.CumulativeLoginDays = s.player.StateComponent().GetCumulativeLoginDays()
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetDiamondCount(iContext context.Context, req *game.CGetDiamondCount) (res *game.SGetDiamondCount, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetDiamondCount on goroutine %v\n", goroutineID)
	res = &game.SGetDiamondCount{}
	res.DiamondCount = s.player.ResourceComponent().GetDiamondCount()
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetEnergyData(ctx context.Context, req *game.CGetEnergyData) (res *game.SGetEnergyData, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetEnergyData on goroutine %v\n", goroutineID)
	res = &game.SGetEnergyData{}
	res.EnergyCount = s.player.ResourceComponent().GetEnergyCount()
	res.NextUpdateTs = s.player.ResourceComponent().GetEnergyNextUpdateTimestamp()
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetGoldCount(ctx context.Context, req *game.CGetGoldCount) (res *game.SGetGoldCount, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetGoldCount on goroutine %v\n", goroutineID)
	res = &game.SGetGoldCount{}
	res.GoldCount = s.player.ResourceComponent().GetGoldCount()
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetLevelExp(ctx context.Context, req *game.CGetLevelExp) (res *game.SGetLevelExp, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetLevelExp on goroutine %v\n", goroutineID)
	res = &game.SGetLevelExp{}
	res.Level, res.Exp = s.player.BasicComponent().GetLevel(), s.player.BasicComponent().GetExp()
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetMaterials(ctx context.Context, req *game.CGetMaterials) (res *game.SGetMaterials, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetMaterials on goroutine %v\n", goroutineID)
	res = &game.SGetMaterials{}
	res.Materials = make(map[int64]int64)
	s.player.ResourceComponent().GetMaterials()
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetOrderList(ctx context.Context, req *game.CGetOrderList) (res *game.SGetOrderList, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetOrderList on goroutine %v\n", goroutineID)
	res = &game.SGetOrderList{}
	res.OrderList = make(map[int64]*models.PBSingleOrder)
	// TODO:
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetPlayerName(c context.Context, req *game.CGetPlayerName) (res *game.SGetPlayerName, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetPlayerName on goroutine %v\n", goroutineID)
	res = &game.SGetPlayerName{}
	res.Name = s.player.IdentifyComponent().GetName()
	return
}

// TODO: Player 层的 handler
func (s *Service) GetProfile(ctx context.Context, req *game.CGetProfile) (res *game.SGetProfile, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetProfile on goroutine %v\n", goroutineID)
	res = &game.SGetProfile{}
	res.ProfilePicture, res.ProfilePictureFrame = s.player.BasicComponent().GetProfilePicture(), s.player.BasicComponent().GetProfilePictureFrame()
	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) GetResults(ctx context.Context, req *game.CGetResults) (res *game.SGetResults, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("GetResults on goroutine %v\n", goroutineID)
	res = &game.SGetResults{}
	res.Results = s.player.ResourceComponent().GetResults()
	return
}

// TODO: Server 层的 handler
// Login
func (s *Service) Login(c context.Context, req *game.CLogin) (*game.SLogin, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("Login on goroutine %v\n", goroutineID)
	ctx := c.(framework.IContext)
	p, err := playerMgr.Instance().GetPlayer(ctx.GetPlayerId())
	if err != nil {
		return nil, err
	}

	// 创建player
	if p == nil {
		p, err = obj.NewPlayer(ctx.GetPlayerId())
		if err != nil {
			return nil, err
		}

		if err = playerMgr.Instance().PutPlayer(p); err != nil {
			return nil, err
		}

		// 保存
		err = p.SavePlayerFull()
		if err != nil {
			return nil, err
		}
	}
	s.player = p
	p.GetModelPlayerInfo().SetLastLoginTime(time.Now().UnixMilli())

	//push 登录事件
	ctx.PublishEventLocal(&event.LoginEvent{Pid: s.playerId})

	// 返回值
	res := &game.SLogin{}
	res.Model = s.player.ToPB()
	logger.Info("player login finished！", "playerId", ctx.GetPlayerId())

	logger.Debugf("player login at Goroutine %v", func() uint64 {
		gID, _ := stp.GoroutineID()
		return gID
	}())

	return res, nil
}

// TODO: Player 层的 handler
func (s *Service) Logout(ctx context.Context, req *game.CLogout) (res *game.SLogout, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("Logout on goroutine %v\n", goroutineID)
	return nil, nil
}

// TODO: Player 层的 handler
func (s *Service) MergeGenerator(ctx context.Context, req *game.CMergeGenerator) (res *game.SMergeGenerator, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("MergeGenerator on goroutine %v\n", goroutineID)
	return nil, nil
}

// TODO: Player 层的 handler
func (s *Service) PlayerOffline(c context.Context, req *game.CPlayerOffline) (*game.SPlayerOffline, error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("PlayerOffline on goroutine %v\n", goroutineID)
	ctx := c.(framework.IContext)
	_, err := playerMgr.Instance().GetPlayer(ctx.GetPlayerId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// TODO: Player 层的 handler
// SyncModel 处理客户端往服务器同步
// 测试数据中 model_player_info，客户端不可写入
// 测试数据中 model_client，服务端不可写入
/*
{
    "items":  [
        {
            "id":  "model_player_info.name",
            "str_value":  "player_101"
        },
        {
            "id":  "model_player_info.register_time",
            "num_value":  1716170400
        },
        {
            "id": "model_client.main_level",
            "num_value": 1
        },
        {
            "id": "model_client.attrs.100",
            "num_value": 100
        }
    ]
}
*/
func (s *Service) SyncModel(c context.Context, req *game.CSyncModel) (res *game.SSyncModel, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("SyncModel on goroutine %v\n", goroutineID)
	// 写入 server model
	for _, syncItem := range req.GetItems() {
		err = s.player.SyncFromClient(strings.Split(syncItem.Id, "."), 0, syncItem)
		if err != nil {
			// TODO: 这里无法保证原子操作
			fmt.Printf("player %v SyncFromClient occurs error %v", s.player, err.Error())
			continue
		}
	}

	// 全量写入 db
	s.player.SavePlayerFull()

	// TODO: 增量写入 db

	// response
	res = &game.SSyncModel{}
	res.TestRes = "this is a test response"

	return
}

// TODO: Player 层的 handler
func (s *Service) UnlockMaterialSlot(ctx context.Context, req *game.CUnlockMaterialSlot) (res *game.SUnlockMaterialSlot, err error) {
	goroutineID, _ := stp.GoroutineID()
	fmt.Printf("UnlockMaterialSlot on goroutine %v\n", goroutineID)
	return nil, nil
}
