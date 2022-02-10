package milestone

import (
	"context"
	"time"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/repository"
	"gorm.io/gorm"
)

type defaultMilestoneService struct {
	db *gorm.DB
}

func NewDefaultMilestoneService() idl.MilestoneService {
	return &defaultMilestoneService{}
}

func (ms *defaultMilestoneService) NextGroup(ctx context.Context, req idl.NextGroupReq) (*idl.PaginationList, error) {
	var (
		timestamp uint
		list      []*domain.Milestone
		nextKey   *uint
		err       error
	)

	if req.NextKey == 0 {
		timestamp = uint(time.Now().UnixNano() / 1e6)
	} else {
		timestamp = req.NextKey
	}

	milestoneRepo := repository.NewMilestoneRepo(ms.db.WithContext(ctx))

	if _list := ms.findAllFromCacheByTimestampDesc(milestoneRepo, req.NextKey, req.Size); _list != nil {
		return &idl.PaginationList{
			List:    fmtToReply(_list),
			NextKey: _list[len(_list)-1].Timestamp,
		}, nil
	}

	if list, err = milestoneRepo.FindAllByTimestamp(timestamp, req.Size+uint(1), "DESC"); err != nil {
		return nil, err
	}

	if len(list) > int(req.Size) {
		nextKey = &list[len(list)-1].Timestamp
		list = list[0 : len(list)-1]
	}

	resp := &idl.PaginationList{
		List:    fmtToReply(list),
		NextKey: nextKey,
	}

	return resp, nil
}

func (ms *defaultMilestoneService) findAllFromCacheByTimestampDesc(milestoneRepo domain.MilestoneRepo, startTimestamp, limit uint) []*domain.Milestone {
	cacheMilestones := milestoneRepo.GetCache()
	if k := getMilestonesIndexByStartTimestamp(cacheMilestones, startTimestamp); k < 0 {
		return nil
	} else if (len(cacheMilestones) - k) < int(limit) {
		return nil
	} else {
		_list := make([]*domain.Milestone, 0, limit)
		for _, milestone := range cacheMilestones[k:(k + int(limit))] {
			_m := *milestone
			_list = append(_list, &_m)
		}
		return _list
	}
}

func getMilestonesIndexByStartTimestamp(milestones []*domain.Milestone, startTimestamp uint) int {
	for i := 0; i < len(milestones); i++ {
		if milestones[i].Timestamp < startTimestamp {
			return i
		} else if milestones[i].Timestamp > startTimestamp {
			return -1
		}
	}
	return -1
}

func fmtToReply(list []*domain.Milestone) []*idl.NextGroupReply {
	_list := make([]*idl.NextGroupReply, 0, len(list))
	for _, m := range list {
		_list = append(_list, &idl.NextGroupReply{
			Title:     m.Title,
			Subtitled: m.Subtitled,
			Type:      m.Type,
			Content:   m.Content,
			TargetUrl: m.TargetUrl,
			Timestamp: m.Timestamp,
		})
	}
	return _list
}
