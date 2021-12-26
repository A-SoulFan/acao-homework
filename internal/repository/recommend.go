package repository

import (
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/launch"
)

func (m *defaultRecommendRepo) SetCache(videos []*domain.RecommendVideo) {
	launch.RecommendSliceCache.Set(videos)
}

func (m *defaultRecommendRepo) GetCache() []*domain.RecommendVideo {
	return launch.RecommendSliceCache.Get()
}

type defaultRecommendRepo struct{}

func NewRecommendRepo() domain.RecommendRepo {
	return &defaultRecommendRepo{}
}
