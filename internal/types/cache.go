package types

import (
	"sync"

	"github.com/A-SoulFan/acao-homework/internal/domain"
)

func NewMilestoneCache() *MilestoneCache {
	return &MilestoneCache{}
}

type MilestoneCache struct {
	data []*domain.Milestone
	lock sync.Mutex
}

func (mc *MilestoneCache) Set(milestones []*domain.Milestone) {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	mc.data = milestones
}

func (mc *MilestoneCache) Get() []*domain.Milestone {
	return mc.data
}
