package idl

import (
	"context"

	"github.com/A-SoulFan/acao-homework/internal/domain"
)

type RecommendTask interface {
	Task
}

type RecommendService interface {
	RecommendTask
	TopRecommendSlices(ctx context.Context) ([]*domain.RecommendVideo, error)
}
