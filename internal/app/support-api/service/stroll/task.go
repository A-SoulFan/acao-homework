package stroll

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/domain"
)

const (
	defaultCacheNumber = 1000
)

type defaultStrollTask struct {
	svcCtx         *svcCtx.ServiceContext
	strollRepo     domain.StrollRepo
	lastUpdateTime uint
}

func (st *defaultStrollTask) SetDBwithCtx(ctx context.Context) {
	db := st.svcCtx.WithDatabaseContext(ctx)
	st.strollRepo.SetDB(db)
}

func (st *defaultStrollTask) InitTask() {
	if err := st.generateCandidateList(); err != nil {
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
				if err := st.generateCandidateList(); err != nil {
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

func (st *defaultStrollTask) getLastUpdateTime() uint {
	return st.lastUpdateTime
}

func (st *defaultStrollTask) generateCandidateList() error {
	var (
		maxId        uint
		err          error
		randomIdList = make([]uint, 0, 10)
		strollList   = make([]*domain.Stroll, 0, 20)
	)

	if maxId, err = st.strollRepo.FindMaxId(); err != nil {
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

	if _list, err := st.strollRepo.FindAllByIds(randomIdList); err != nil {
		return err
	} else {
		strollList = append(strollList, _list...)
	}

	st.strollRepo.SetCache(strollList)

	if st.lastUpdateTime, err = st.strollRepo.FindLastUpdateTime(); err != nil {
		return err
	}

	return nil
}
