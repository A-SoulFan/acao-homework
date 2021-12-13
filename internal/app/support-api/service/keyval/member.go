package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/types"
	appErr "github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
	"github.com/A-SoulFan/acao-homework/internal/repository"

	"gorm.io/gorm"
)

const (
	memberListKey          = "member_list"
	memberExperiencePrefix = "member_experience_"
	memberVideoPrefix      = "member_videos_"
)

type MemberLogic struct {
	ctx    context.Context
	svcCtx *svcCtx.ServiceContext
	dbCtx  *gorm.DB
}

func NewMemberLogic(ctx context.Context, svcCtx *svcCtx.ServiceContext) MemberLogic {
	return MemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.WithDatabaseContext(ctx),
	}
}

func (m *MemberLogic) GetAll() (*types.MemberAll, error) {
	val, err := repository.NewDefaultKeyValueRepo(m.dbCtx).FindOneByKey(memberListKey)
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

	return &types.MemberAll{MemberList: list}, nil
}

func (m *MemberLogic) GetExperience(req types.MemberExperienceReq) (*types.MemberExperienceResp, error) {
	queryKey := fmt.Sprintf("%s%s", memberExperiencePrefix, strings.ToLower(req.MemberName))

	val, err := repository.NewDefaultKeyValueRepo(m.dbCtx).FindOneByKey(queryKey)
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

	return &types.MemberExperienceResp{
		MemberName: req.MemberName,
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}

func (m *MemberLogic) GetVideos(req types.MemberVideoReq) (*types.MemberExperienceResp, error) {
	queryKey := fmt.Sprintf("%s%s", memberVideoPrefix, strings.ToLower(req.MemberName))

	val, err := repository.NewDefaultKeyValueRepo(m.dbCtx).FindOneByKey(queryKey)
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

	return &types.MemberExperienceResp{
		MemberName: req.MemberName,
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}
