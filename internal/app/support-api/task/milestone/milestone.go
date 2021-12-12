package milestone

import (
	"context"
	"sync"
	"time"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/repository"
)

const (
	defaultCacheNumber = 1000
)

type taskRebuildCache struct {
	svcCtx        *svcCtx.ServiceContext
	cacheData     []*domain.Milestone
	milestoneRepo domain.MilestoneRepo
	lock          sync.Mutex
}

var (
	once sync.Once
	_tr  *taskRebuildCache
)

func Register(svc *svcCtx.ServiceContext) {
	once.Do(func() {
		_tr = &taskRebuildCache{
			svcCtx:        svc,
			cacheData:     make([]*domain.Milestone, 0, defaultCacheNumber),
			milestoneRepo: repository.NewMilestoneRepo(svc.Db.WithContext(context.TODO())),
		}

		_tr.init()
	})
}

func FindCacheAllByTimestampDesc(startTimestamp, limit uint) []*domain.Milestone {
	return _tr.findCacheAllByTimestampDesc(startTimestamp, limit)
}

func (tr *taskRebuildCache) init() {
	if err := tr.RebuildCache(); err != nil {
		panic(err)
	}

	ticker(tr, tr.svcCtx)
}

func (tr *taskRebuildCache) findCacheAllByTimestampDesc(startTimestamp, limit uint) []*domain.Milestone {
	if k := search(tr.cacheData, startTimestamp); k < 0 {
		return nil
	} else if (len(tr.cacheData) - k) < int(limit) {
		return nil
	} else {
		_list := make([]*domain.Milestone, 0, limit)
		for _, milestone := range tr.cacheData[k:(k + int(limit))] {
			_m := *milestone
			_list = append(_list, &_m)
		}
		return _list
	}
}

func (tr *taskRebuildCache) RebuildCache() error {
	tr.lock.Lock()
	defer tr.lock.Unlock()
	if list, err := tr.milestoneRepo.FindAllByTimestamp(uint(time.Now().UnixNano()/1e6), defaultCacheNumber, "DESC"); err != nil {
		return err
	} else {
		tr.cacheData = list
	}
	return nil
}

func search(milestones []*domain.Milestone, startTimestamp uint) int {
	for i := 0; i < len(milestones); i++ {
		if milestones[i].Timestamp < startTimestamp {
			return i
		} else if milestones[i].Timestamp > startTimestamp {
			return -1
		}
	}
	return -1
}

func ticker(tr *taskRebuildCache, svc *svcCtx.ServiceContext) {
	ticker := time.NewTicker(5 * time.Minute)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := tr.RebuildCache(); err != nil {
					svc.Logger.Error(err)
				}
				//svc.Logger.Info("milestone.task.rebuildCache: successfully.")
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(ticker)
}
