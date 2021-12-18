package launch

import (
	"github.com/A-SoulFan/acao-homework/internal/types"
)

var (
	MilestoneCache *types.MilestoneCache
)

func launchCaches() {
	MilestoneCache = types.NewMilestoneCache()
}
