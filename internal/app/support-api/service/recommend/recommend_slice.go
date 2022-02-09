package recommend

import (
	"context"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/domain"
)

type defaultRecommendService struct {
	defaultRecommendTask
}

func NewDefaultRecommendService() idl.RecommendService {
	return &defaultRecommendService{
		defaultRecommendTask: defaultRecommendTask{},
	}
}

func (rs *defaultRecommendService) TopRecommendSlices(ctx context.Context) ([]*domain.RecommendVideo, error) {
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
