package idl

import "github.com/A-SoulFan/acao-homework/internal/domain"

type RecommendTask interface {
	InitTask()
}

type RecommendService interface {
	RecommendTask
	TopRecommendSlices() ([]*domain.RecommendVideo, error)
}
