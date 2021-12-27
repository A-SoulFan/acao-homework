package stroll

import (
	"context"
	"fmt"
	svc "github.com/A-SoulFan/acao-homework/internal/app/admin-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/admin-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
	"github.com/A-SoulFan/acao-homework/internal/pkg/utility/bilibili"
)

type defaultStrollService struct {
	svcCtx     *svc.ServiceContext
	strollRepo domain.StrollRepo
}

func NewDefaultStrollService(svcCtx *svc.ServiceContext, strollRepo domain.StrollRepo) idl.StrollService {
	return &defaultStrollService{
		svcCtx:     svcCtx,
		strollRepo: strollRepo,
	}
}

func (s *defaultStrollService) Create(ctx context.Context, req idl.StrollCreateReq) error {
	view, err := (&bilibili.BiliBili{}).WebInterfaceView(req.Bv)
	if err != nil {
		return apperrors.NewServiceError("调用B站API错误").Wrap(err)
	}

	if req.Pages != 0 {
		flag := false
		for _, page := range view.Pages {
			if page.Page == int(req.Pages) {
				flag = true
				break
			}
		}

		if !flag {
			return apperrors.NewValidateError("无效的 pages ")
		}
	}

	err = s.strollRepo.Insert(&domain.Stroll{
		Title:     view.Title,
		Cover:     view.Pic,
		BV:        view.BvId,
		TargetUrl: fmt.Sprintf("https://www.bilibili.com/video/%s", view.BvId),
		Play:      fmt.Sprintf("//player.bilibili.com/player.html?bvid=%s", view.BvId),
	})

	if err != nil {
		return err
	}

	return nil
}
