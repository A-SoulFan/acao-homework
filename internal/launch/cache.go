package launch

import (
	"github.com/A-SoulFan/acao-homework/internal/types"
)

var (
	MilestoneCache      *types.MilestoneCache
	StrollCache         *types.StrollCache
	RecommendSliceCache *types.RecommendSliceCache
)

func launchCaches() {
	MilestoneCache = types.NewMilestoneCache()
	StrollCache = types.NewStrollCache()
	RecommendSliceCache = types.NewRecommendSliceCache()
}
