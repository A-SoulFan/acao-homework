package team

import (
	"context"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	appErr "github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
	"github.com/A-SoulFan/acao-homework/internal/repository"
	"gorm.io/gorm"

	"encoding/json"
	"fmt"
)

const (
	teamVideos       = "team_videos"
	teamEventsPrefix = "team_events_"
)

type defaultTeamService struct {
	db *gorm.DB
}

func NewDefaultTeamService(db *gorm.DB) idl.TeamService {
	return &defaultTeamService{db: db}
}

func (ts *defaultTeamService) GetTeamVideos(ctx context.Context, req idl.TeamVideosReq) (*idl.TeamVideosResp, error) {
	teamRepo := repository.NewKeyValueRepo(ts.db.WithContext(ctx))

	val, err := teamRepo.FindOneByKey(teamVideos)
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

func (ts *defaultTeamService) GetTeamEvents(ctx context.Context, req idl.TeamEventsReq) (*idl.TeamEventsResp, error) {
	teamRepo := repository.NewKeyValueRepo(ts.db.WithContext(ctx))

	queryKey := fmt.Sprintf("%s%s", teamEventsPrefix, req.Year)
	val, err := teamRepo.FindOneByKey(queryKey)
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
