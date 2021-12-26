package recommend

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/domain"
)

type defaultRecommendService struct {
	defaultRecommendTask
}

func NewDefaultRecommendService(stx *svcCtx.ServiceContext, recommendRepo domain.RecommendRepo) idl.RecommendService {
	return &defaultRecommendService{
		defaultRecommendTask: defaultRecommendTask{
			svcCtx:        stx,
			recommendRepo: recommendRepo,
		},
	}
}

func (rs *defaultRecommendService) TopRecommendSlices() ([]*domain.RecommendVideo, error) {
	return rs.top(20), nil
}

func (rs *defaultRecommendService) top(n int) []*domain.RecommendVideo {
	videoList := make([]*domain.RecommendVideo, 0, n)

	cache := rs.recommendRepo.GetCache()

	for i, video := range cache {
		if i == n {
			break
		}

		videoList = append(videoList, video)
	}

	return videoList
}
