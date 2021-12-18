package stroll

import (
	"context"
	"fmt"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	strollTask "github.com/A-SoulFan/acao-homework/internal/app/support-api/task/stroll"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/types"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	appErr "github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
	"github.com/A-SoulFan/acao-homework/internal/pkg/utility"
	"github.com/A-SoulFan/acao-homework/internal/repository"

	"gorm.io/gorm"
)

type RandomGetLogic struct {
	ctx    context.Context
	svcCtx *svcCtx.ServiceContext
	dbCtx  *gorm.DB
}

func NewRandomGetLogic(ctx context.Context, svcCtx *svcCtx.ServiceContext) RandomGetLogic {
	return RandomGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.WithDatabaseContext(ctx),
	}
}

func (r *RandomGetLogic) LastUpdateTime() (*types.StrollLastUpdateReply, error) {
	resp := &types.StrollLastUpdateReply{LastUpdateTime: strollTask.LastUpdateTime()}
	return resp, nil
}

func (r *RandomGetLogic) RandomGetStroll() (*types.StrollReply, error) {
	if stroll, err := strollTask.RandomStroll(); err != nil {
		// r.svcCtx.Logger.Error(err)
		return nil, appErr.NewServiceError("暂时没有可以溜的数据哦，请稍后再试。").Wrap(err)
	} else {
		if stroll.Cover == "" {
			if err := getBliBilCover(&stroll); err != nil {
				// r.svcCtx.Logger.Error(err)
			} else {
				if err := repository.NewStrollRepo(r.dbCtx).UpdateCover(stroll.BV, stroll.Cover); err != nil {
					// r.svcCtx.Logger.Error(err)
				}
			}
		}

		return &types.StrollReply{
			Title:     stroll.Title,
			Cover:     stroll.Cover,
			BV:        stroll.BV,
			PlayUrl:   fmt.Sprintf("//player.bilibili.com/player.html?bvid=%s", stroll.BV),
			TargetUrl: stroll.TargetUrl,
			CreatedAt: stroll.CreatedAt,
		}, nil
	}
}

func getBliBilCover(stroll *domain.Stroll) error {
	resp, err := (&utility.BiliBili{}).WebInterfaceView(stroll.BV)
	if err != nil {
		return err
	}

	stroll.Cover = resp.Pic
	return nil
}
