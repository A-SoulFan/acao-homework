package milestone

import (
	"context"
	"time"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	milestoneTask "github.com/A-SoulFan/acao-homework/internal/app/support-api/task/milestone"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/types"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/repository"
	"gorm.io/gorm"
)

type NextGroupLogic struct {
	ctx    context.Context
	svcCtx *svcCtx.ServiceContext
	dbCtx  *gorm.DB
}

func NewGroupLogic(ctx context.Context, svcCtx *svcCtx.ServiceContext) NextGroupLogic {
	return NextGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.WithDatabaseContext(ctx),
	}
}

func (ng *NextGroupLogic) NextGroup(req types.NextGroupReq) (*types.PaginationList, error) {
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

	if _list := milestoneTask.FindCacheAllByTimestampDesc(req.NextKey, req.Size); _list != nil {
		return &types.PaginationList{
			List:    toReply(_list),
			NextKey: _list[len(_list)-1].Timestamp,
		}, nil
	}

	if list, err = repository.NewMilestoneRepo(ng.dbCtx).FindAllByTimestamp(timestamp, req.Size+uint(1), "DESC"); err != nil {
		return nil, err
	}

	if len(list) > int(req.Size) {
		nextKey = &list[len(list)-1].Timestamp
		list = list[0 : len(list)-1]
	}

	resp := &types.PaginationList{
		List:    toReply(list),
		NextKey: nextKey,
	}

	return resp, nil
}

func toReply(list []*domain.Milestone) []*types.NextGroupReply {
	_list := make([]*types.NextGroupReply, 0, len(list))
	for _, m := range list {
		_list = append(_list, &types.NextGroupReply{
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
