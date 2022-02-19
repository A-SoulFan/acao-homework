package repository

import (
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/launch"
)

func (m *defaultDebtRepo) SetCache(key string, debts []*domain.Debt) {
	launch.DebtCache.Set(key, debts)
}

func (m *defaultDebtRepo) GetCache(key string) []*domain.Debt {
	return launch.DebtCache.Get(key)
}

type defaultDebtRepo struct{}

func NewDebtRepo() domain.DebtRepo {
	return &defaultDebtRepo{}
}
