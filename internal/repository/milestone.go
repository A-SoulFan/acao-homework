package repository

import (
	"fmt"
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"gorm.io/gorm"
	"time"
)

type defaultMilestoneRepo struct {
	conn *gorm.DB
}

func NewMilestoneRepo(conn *gorm.DB) domain.MilestoneRepo {
	return &defaultMilestoneRepo{conn: conn}
}

func (m *defaultMilestoneRepo) Insert(data *domain.Milestone) error {
	result := m.conn.Table("milestones").Create(data)
	return result.Error
}

func (m *defaultMilestoneRepo) Delete(primaryKey uint) error {
	result := m.conn.Exec("UPDATE milestones SET deleted_at = ? WHERE id = ? AND deleted_at = 0", time.Now().UnixNano()/1e6, primaryKey)
	return result.Error
}

func (m *defaultMilestoneRepo) Update(data *domain.Milestone) error {
	result := m.conn.Table("milestones").Updates(data)
	return result.Error
}

func (m *defaultMilestoneRepo) SearchTitles(keyword string, limit uint) ([]*domain.Milestone, error) {
	var list []*domain.Milestone
	result := m.conn.Raw("SELECT * FROM milestones WHERE Title LIKE ? AND deleted_at = 0 LIMIT 0, ?", keyword+"%", limit).Find(&list)
	return list, result.Error
}

func (m *defaultMilestoneRepo) FindOne(primaryKey uint) (*domain.Milestone, error) {
	milestone := &domain.Milestone{}
	result := m.conn.Raw("SELECT * FROM milestones WHERE id = ? AND deleted_at = 0", primaryKey).First(milestone)
	if result.Error != nil {
		return nil, result.Error
	}
	return milestone, nil
}

func (m *defaultMilestoneRepo) FindAllByTimestamp(startTimestamp, limit uint, order string) ([]*domain.Milestone, error) {
	var list []*domain.Milestone
	result := m.conn.Raw(fmt.Sprintf("SELECT * FROM milestones WHERE timestamp < ? AND deleted_at = 0 ORDER BY timestamp %s LIMIT 0, ?", order), startTimestamp, limit).Find(&list)
	return list, result.Error
}
