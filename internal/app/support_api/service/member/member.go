package member

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"sort"
	"strings"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	appErr "github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
	"github.com/A-SoulFan/acao-homework/internal/repository"
	"gorm.io/gorm"
)

const (
	memberListKey          = "member_list"
	memberExperiencePrefix = "member_experience_"
	memberVideoPrefix      = "member_videos_"
	memberDebtListKey      = "member_debt_list"
)

type defaultMemberService struct {
	db *gorm.DB
}

func NewDefaultMemberService(db *gorm.DB) idl.MemberService {
	return &defaultMemberService{db: db}
}

func (ms *defaultMemberService) GetAllMembers(ctx context.Context) (*idl.MemberAll, error) {
	memberRepo := repository.NewKeyValueRepo(ms.db.WithContext(ctx))
	val, err := memberRepo.FindOneByKey(memberListKey)
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

	memberRepo := repository.NewKeyValueRepo(ms.db.WithContext(ctx))
	val, err := memberRepo.FindOneByKey(queryKey)
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

	memberRepo := repository.NewKeyValueRepo(ms.db.WithContext(ctx))
	val, err := memberRepo.FindOneByKey(queryKey)
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

func (ms defaultMemberService) GetMemberDebts(ctx context.Context, req idl.MemberDebtReq) []*domain.Debt {
	isAll := true
	var memberNames []string
	cacheKey := "all"
	if len(req.MemberName) != 0 {
		isAll = false
		queryKey := strings.ToLower(req.MemberName)
		memberNames = strings.Split(queryKey, ",")
		// 排序
		sort.Strings(memberNames)
		cacheKey = strings.Join(memberNames, "-")
	}

	// 先查询缓存看是否有值
	debtRepo := repository.NewDebtRepo()
	debtRes := debtRepo.GetCache(cacheKey)
	if debtRes != nil {
		return debtRes
	}

	// 根据所有找到符合要去的
	var allDebts []*domain.Debt
	if !isAll {
		allDebts = debtRepo.GetCache("all")
	}

	if allDebts == nil || isAll {
		memberRepo := repository.NewKeyValueRepo(ms.db.WithContext(ctx))
		val, err := memberRepo.FindOneByKey(memberDebtListKey)
		if err != nil {
			return []*domain.Debt{}
		}
		var list []*domain.Debt
		if err := json.Unmarshal(val.Value, &list); err != nil {
			return []*domain.Debt{}
		}

		allDebts = list
		debtRepo.SetCache("all", allDebts)

	}

	var debtVal []*domain.Debt
	// 遍历取出符合条件的
	OuterLoop:
		for _, debt := range allDebts {
			debt := debt
			for _, tag := range debt.Tags {
				memberName := tag.Key
				for _, val := range memberNames {
					val := val
					if val == memberName {
						debtVal = append(debtVal, debt)
						continue OuterLoop
					}
				}
			}
		}



	debtRepo.SetCache(cacheKey, debtVal)
	return debtVal
}
