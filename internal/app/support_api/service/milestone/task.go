package milestone

import (
	"context"
	"time"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	"github.com/A-SoulFan/acao-homework/internal/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	defaultCacheNumber = 1000
)

type defaultMilestoneTask struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewDefaultMilestoneTask(logger *zap.Logger, db *gorm.DB) idl.MilestoneTask {
	return newDefaultMilestoneTask(logger, db)
}

func newDefaultMilestoneTask(logger *zap.Logger, db *gorm.DB) *defaultMilestoneTask {
	return &defaultMilestoneTask{
		logger: logger,
		db:     db,
	}
}

func (mt *defaultMilestoneTask) InitTask(ctx context.Context) {
	if err := mt.rebuildCache(ctx); err != nil {
		panic(err)
	}

	mt.startTick()
}

func (mt *defaultMilestoneTask) rebuildCache(ctx context.Context) error {
	startTimestamp := uint(time.Now().UnixNano() / 1e6)

	milestoneRepo := repository.NewMilestoneRepo(mt.db.WithContext(ctx))
	if list, err := milestoneRepo.FindAllByTimestamp(startTimestamp, defaultCacheNumber, "DESC"); err != nil {
		return err
	} else {
		milestoneRepo.SetCache(list)
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
				if err := mt.rebuildCache(context.Background()); err != nil {
					mt.logger.Error("task.rebuildCache error", zap.Error(err))
				}
				mt.logger.Info("task.rebuildCache: successfully.")
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(tk)
}
