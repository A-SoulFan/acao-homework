package launch

import (
	"github.com/A-SoulFan/acao-homework/internal/launch/cache"
)

var (
	MilestoneCache      *cache.MilestoneCache
	StrollCache         *cache.StrollCache
	RecommendSliceCache *cache.RecommendSliceCache
)

func launchCaches() {
	MilestoneCache = cache.NewMilestoneCache()
	StrollCache = cache.NewStrollCache()
	RecommendSliceCache = cache.NewRecommendSliceCache()
}
