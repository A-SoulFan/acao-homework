package recommend

import (
	"context"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	recommendTask "github.com/A-SoulFan/acao-homework/internal/app/support-api/task/recommend"
	"github.com/A-SoulFan/acao-homework/internal/domain"
)

type RecommendSliceLogic struct {
	ctx    context.Context
	svcCtx *svcCtx.ServiceContext
}

func NewRecommendSliceLogic(svc *svcCtx.ServiceContext) *RecommendSliceLogic {
	return &RecommendSliceLogic{svcCtx: svc}
}

func (rm *RecommendSliceLogic) Handle() ([]domain.RecommendVideo, error) {
	return recommendTask.Hot(20), nil
}
