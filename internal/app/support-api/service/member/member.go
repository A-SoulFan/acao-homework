package member

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	appErr "github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
)

const (
	memberListKey          = "member_list"
	memberExperiencePrefix = "member_experience_"
	memberVideoPrefix      = "member_videos_"
)

type defaultMemberService struct {
	memberRepo domain.KeyValueRepo
}

func NewDefaultMemberService() idl.MemberService {
	return &defaultMemberService{}
}

func (ms *defaultMemberService) GetAllMembers(ctx context.Context) (*idl.MemberAll, error) {
	val, err := ms.memberRepo.FindOneByKey(memberListKey)
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

	return &idl.MemberAll{MemberList: list}, nil
}

func (ms *defaultMemberService) GetMemberExperience(ctx context.Context, req idl.MemberExperienceReq) (*idl.MemberExperienceResp, error) {
	queryKey := fmt.Sprintf("%s%s", memberExperiencePrefix, strings.ToLower(req.MemberName))

	val, err := ms.memberRepo.FindOneByKey(queryKey)
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

	return &idl.MemberExperienceResp{
		MemberName: req.MemberName,
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}

func (ms *defaultMemberService) GetMemberVideos(ctx context.Context, req idl.MemberVideoReq) (*idl.MemberExperienceResp, error) {
	queryKey := fmt.Sprintf("%s%s", memberVideoPrefix, strings.ToLower(req.MemberName))

	val, err := ms.memberRepo.FindOneByKey(queryKey)
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

	return &idl.MemberExperienceResp{
		MemberName: req.MemberName,
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}
