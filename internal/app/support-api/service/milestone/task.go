package milestone

import (
	"log"
	"time"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"gorm.io/gorm"
)

const (
	defaultCacheNumber = 1000
)

type defaultMilestoneTask struct {
	svcCtx        *svcCtx.ServiceContext
	milestoneRepo domain.MilestoneRepo
}

func (mt *defaultMilestoneTask) SetDB(db *gorm.DB) {
	mt.milestoneRepo.SetDB(db)
}

func (mt *defaultMilestoneTask) InitTask() {
	if err := mt.rebuildCache(); err != nil {
		panic(err)
	}

	mt.startTick()
}

func (mt *defaultMilestoneTask) rebuildCache() error {
	startTimestamp := uint(time.Now().UnixNano() / 1e6)
	if list, err := mt.milestoneRepo.FindAllByTimestamp(startTimestamp, defaultCacheNumber, "DESC"); err != nil {
		return err
	} else {
		mt.milestoneRepo.SetCache(list)
	}
	return nil
}

func (mt *defaultMilestoneTask) startTick() {
	tk := time.NewTicker(5 * time.Minute)

	stopChan := make(chan bool)
	go func(tk *time.Ticker) {
		defer tk.Stop()

		for {
			select {
			case <-tk.C:
				if err := mt.rebuildCache(); err != nil {
					log.Println(err)
				}
				//svc.Logger.Info("milestone.task.rebuildCache: successfully.")
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(tk)
}
