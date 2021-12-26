package logic

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	appErr "github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"

	"context"
	"encoding/json"
	"fmt"
)

const (
	teamVideos       = "team_videos"
	teamEventsPrefix = "team_events_"
)

type defaultTeamService struct {
	svcCtx   *svcCtx.ServiceContext
	teamRepo domain.KeyValueRepo
}

func NewDefaultTeamService(svcCtx *svcCtx.ServiceContext, teamRepo domain.KeyValueRepo) idl.TeamService {
	return &defaultTeamService{
		svcCtx:   svcCtx,
		teamRepo: teamRepo,
	}
}

func (ts *defaultTeamService) SetDBwithCtx(ctx context.Context) {
	db := ts.svcCtx.WithDatabaseContext(ctx)
	ts.teamRepo.SetDB(db)
}

func (ts *defaultTeamService) GetTeamVideos(req idl.TeamVideosReq) (*idl.TeamVideosResp, error) {
	val, err := ts.teamRepo.FindOneByKey(teamVideos)
	if err != nil {
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewServiceError("获取数据失败").Wrap(err)
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		return nil, err
	}

	return &idl.TeamVideosResp{
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}

func (ts *defaultTeamService) GetTeamEvents(req idl.TeamEventsReq) (*idl.TeamEventsResp, error) {
	queryKey := fmt.Sprintf("%s%s", teamEventsPrefix, req.Year)
	val, err := ts.teamRepo.FindOneByKey(queryKey)
	if err != nil {
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewServiceError("获取数据失败").Wrap(err)
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		return nil, err
	}

	return &idl.TeamEventsResp{
		TotalCount: len(list),
		TotalPage:  1,
		EventList:  list,
	}, nil
}
