package stroll

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	"github.com/A-SoulFan/acao-homework/internal/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/A-SoulFan/acao-homework/internal/domain"
)

const (
	defaultCacheNumber = 1000
)

type defaultStrollTask struct {
	logger         *zap.Logger
	db             *gorm.DB
	lastUpdateTime uint
}

func NewDefaultStrollTask(logger *zap.Logger, db *gorm.DB) idl.StrollTask {
	return newDefaultStrollTask(logger, db)
}

func newDefaultStrollTask(logger *zap.Logger, db *gorm.DB) *defaultStrollTask {
	return &defaultStrollTask{
		logger:         logger,
		db:             db,
		lastUpdateTime: 0,
	}
}

func (st *defaultStrollTask) InitTask(ctx context.Context) {
	if err := st.generateCandidateList(ctx); err != nil {
		panic(err)
	}

	st.startTick()
}

func (st *defaultStrollTask) startTick() {
	tk := time.NewTicker(5 * time.Minute)

	stopChan := make(chan bool)
	go func(tk *time.Ticker) {
		defer tk.Stop()

		for {
			select {
			case <-tk.C:
				if err := st.generateCandidateList(context.Background()); err != nil {
					st.logger.Error("task.generateCandidateList error", zap.Error(err))
				}
				st.logger.Info("task.generateCandidateList: successfully.")
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(tk)
}

func (st *defaultStrollTask) getLastUpdateTime() uint {
	return st.lastUpdateTime
}

func (st *defaultStrollTask) generateCandidateList(ctx context.Context) error {
	var (
		maxId        uint
		err          error
		randomIdList = make([]uint, 0, 10)
		strollList   = make([]*domain.Stroll, 0, 20)
	)

	strollRepo := repository.NewStrollRepo(st.db.WithContext(ctx))
	if maxId, err = strollRepo.FindMaxId(); err != nil {
		return err
	} else {
		// on database clear
		if maxId < 1 {
			return errors.New("candidate list database is empty. ")
		}
	}

	for i := 0; i < 20; i++ {
		randId := rand.Intn(int(maxId)) + 1
		randomIdList = append(randomIdList, uint(randId))
	}

	if _list, err := strollRepo.FindAllByIds(randomIdList); err != nil {
		return err
	} else {
		strollList = append(strollList, _list...)
	}

	strollRepo.SetCache(strollList)

	if st.lastUpdateTime, err = strollRepo.FindLastUpdateTime(); err != nil {
		return err
	}

	return nil
}
