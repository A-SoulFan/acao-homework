package stroll

import (
	"errors"
	"fmt"
	"math/rand"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	appErr "github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
	"github.com/A-SoulFan/acao-homework/internal/pkg/utility"
)

type defaultStrollService struct {
	defaultStrollTask
}

func NewDefaultStrollService(stx *svcCtx.ServiceContext, strollRepo domain.StrollRepo) idl.StrollService {
	return &defaultStrollService{
		defaultStrollTask: defaultStrollTask{
			svcCtx:     stx,
			strollRepo: strollRepo,
		},
	}
}

func (r *defaultStrollService) LastUpdateTime() (*idl.StrollLastUpdateReply, error) {
	resp := &idl.StrollLastUpdateReply{LastUpdateTime: r.getLastUpdateTime()}
	return resp, nil
}

func (r *defaultStrollService) RandomGetStroll() (*idl.StrollReply, error) {
	if stroll, err := r.randomStroll(); err != nil {
		// r.svcCtx.Logger.Error(err)
		return nil, appErr.NewServiceError("暂时没有可以溜的数据哦，请稍后再试。").Wrap(err)
	} else {
		if stroll.Cover == "" {
			if err := getBliBilCover(&stroll); err != nil {
				// r.svcCtx.Logger.Error(err)
			} else {
				if err := r.strollRepo.UpdateCover(stroll.BV, stroll.Cover); err != nil {
					// r.svcCtx.Logger.Error(err)
				}
			}
		}

		return &idl.StrollReply{
			Title:     stroll.Title,
			Cover:     stroll.Cover,
			BV:        stroll.BV,
			PlayUrl:   fmt.Sprintf("//player.bilibili.com/player.html?bvid=%s", stroll.BV),
			TargetUrl: stroll.TargetUrl,
			CreatedAt: stroll.CreatedAt,
		}, nil
	}
}

func (r *defaultStrollService) randomStroll() (domain.Stroll, error) {
	candidateList := r.strollRepo.GetCache()
	l := len(candidateList)
	if l != 0 {
		if stroll := candidateList[rand.Intn(l)]; stroll != nil {
			return *stroll, nil
		}
	}

	return domain.Stroll{}, errors.New("candidate list is empty. ")
}

func getBliBilCover(stroll *domain.Stroll) error {
	resp, err := (&utility.BiliBili{}).WebInterfaceView(stroll.BV)
	if err != nil {
		return err
	}

	stroll.Cover = resp.Pic
	return nil
}
