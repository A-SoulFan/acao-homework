package stroll

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/repository"

	"context"
	"errors"
	"math/rand"
	"sync"
	"time"
)

type timerQuery struct {
	svcCtx         *svcCtx.ServiceContext
	candidateList  []*domain.Stroll
	lastUpdateTime uint
	strollModel    domain.StrollRepo
}

var (
	once sync.Once
	_tq  *timerQuery
)

func Register(svc *svcCtx.ServiceContext) {
	once.Do(func() {
		_tq = &timerQuery{
			svcCtx:        svc,
			candidateList: make([]*domain.Stroll, 0, 10),
			strollModel:   repository.NewStrollRepo(svc.WithDatabaseContext(context.TODO())),
		}

		_tq.init()
	})
}

func RandomStroll() (domain.Stroll, error) {
	if _tq == nil {
		panic("not registered RandomStroll task.")
	}

	return _tq.randomStroll()
}

func LastUpdateTime() uint {
	return _tq.lastUpdateTime
}

func (tq *timerQuery) init() {
	if err := tq.generateCandidateList(); err != nil {
		panic(err)
	}

	ticker(tq, tq.svcCtx)
}

func (tq *timerQuery) randomStroll() (domain.Stroll, error) {
	l := len(tq.candidateList)
	if l != 0 {
		if stroll := tq.candidateList[rand.Intn(l)]; stroll != nil {
			return *stroll, nil
		}
	}
	return domain.Stroll{}, errors.New("candidate list is empty. ")
}

func (tq *timerQuery) generateCandidateList() error {
	var (
		maxId        uint
		err          error
		randomIdList = make([]uint, 0, 10)
		strollList   = make([]*domain.Stroll, 0, 20)
	)

	if maxId, err = tq.strollModel.FindMaxId(); err != nil {
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

	if _list, err := tq.strollModel.FindAllByIds(randomIdList); err != nil {
		return err
	} else {
		strollList = append(strollList, _list...)
	}

	tq.candidateList = strollList

	if tq.lastUpdateTime, err = tq.strollModel.FindLastUpdateTime(); err != nil {
		return err
	}

	return nil
}

func ticker(tq *timerQuery, svc *svcCtx.ServiceContext) {
	ticker := time.NewTicker(5 * time.Second)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := tq.generateCandidateList(); err != nil {
					// svc.Logger.Error(err)
				}
				//svc.Logger.Info("stroll.task.timerQuery: successfully generated candidate list.")
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(ticker)
}
