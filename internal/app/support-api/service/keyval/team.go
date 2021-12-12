package logic

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/types"
	appErr "github.com/A-SoulFan/acao-homework/internal/pkg/err"
	"github.com/A-SoulFan/acao-homework/internal/repository"

	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

const (
	teamVideos       = "team_videos"
	teamEventsPrefix = "team_events_"
)

type TeamLogic struct {
	ctx    context.Context
	svcCtx *svcCtx.ServiceContext
	dbCtx  *gorm.DB
}

func NewTeamLogic(ctx context.Context, svcCtx *svcCtx.ServiceContext) TeamLogic {
	return TeamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.Db.WithContext(ctx),
	}
}

func (t *TeamLogic) GetVideos(req types.TeamVideosReq) (*types.TeamVideosResp, error) {
	val, err := repository.NewDefaultKeyValueRepo(t.dbCtx).FindOneByKey(teamVideos)
	if err != nil {
		t.svcCtx.Logger.Error(err)
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewError("获取数据失败")
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		t.svcCtx.Logger.Error(err)
		return nil, err
	}

	return &types.TeamVideosResp{
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}

func (t *TeamLogic) GetEvents(req types.TeamEventsReq) (*types.TeamEventsResp, error) {
	queryKey := fmt.Sprintf("%s%s", teamEventsPrefix, req.Year)
	val, err := repository.NewDefaultKeyValueRepo(t.dbCtx).FindOneByKey(queryKey)
	if err != nil {
		t.svcCtx.Logger.Error(err)
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewError("获取数据失败")
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		t.svcCtx.Logger.Error(err)
		return nil, err
	}

	return &types.TeamEventsResp{
		TotalCount: len(list),
		TotalPage:  1,
		EventList:  list,
	}, nil
}
