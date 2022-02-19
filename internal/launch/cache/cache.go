package cache

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

func NewStrollCache() *StrollCache {
	return &StrollCache{}
}

type StrollCache struct {
	data []*domain.Stroll
	lock sync.Mutex
}

func (sc *StrollCache) Set(strolls []*domain.Stroll) {
	sc.lock.Lock()
	defer sc.lock.Unlock()
	sc.data = strolls
}

func (sc *StrollCache) Get() []*domain.Stroll {
	return sc.data
}

func NewRecommendSliceCache() *RecommendSliceCache {
	return &RecommendSliceCache{}
}

type RecommendSliceCache struct {
	data []*domain.RecommendVideo
	lock sync.Mutex
}

func (rsc *RecommendSliceCache) Set(videos []*domain.RecommendVideo) {
	rsc.lock.Lock()
	defer rsc.lock.Unlock()
	rsc.data = videos
}

func (rsc *RecommendSliceCache) Get() []*domain.RecommendVideo {
	return rsc.data
}

func NewDebtCache() *DebtCache {
	return &DebtCache{
		data: make(map [string][]*domain.Debt),
	}
}

type DebtCache struct {
	data map[string][]*domain.Debt
	lock sync.Mutex
}

func (dc *DebtCache) Set(key string, debt []*domain.Debt) {
	dc.lock.Lock()
	defer dc.lock.Unlock()
	dc.data[key] = debt
}

func (dc *DebtCache) Get(key string) []*domain.Debt {
	if val, ok := dc.data[key]; ok {
		return val
	}
	return nil
}
