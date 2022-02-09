package repository

import (
	"time"

	"gorm.io/gorm"

	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/launch"
)

func (m *defaultStrollRepo) SetCache(strolls []*domain.Stroll) {
	launch.StrollCache.Set(strolls)
}

func (m *defaultStrollRepo) GetCache() []*domain.Stroll {
	return launch.StrollCache.Get()
}

type defaultStrollRepo struct {
	conn *gorm.DB
}

func NewStrollRepo(conn *gorm.DB) domain.StrollRepo {
	return &defaultStrollRepo{conn: conn}
}

func (m *defaultStrollRepo) Insert(data *domain.Stroll) error {
	result := m.conn.Table("stroll").Create(data)
	return result.Error
}

func (m *defaultStrollRepo) Delete(primaryKey uint) error {
	result := m.conn.Exec("UPDATE stroll SET deleted_at = ? WHERE id = ? AND deleted_at = 0", time.Now().UnixNano()/1e6, primaryKey)
	return result.Error
}

func (m *defaultStrollRepo) Update(data *domain.Stroll) error {
	result := m.conn.Table("stroll").Updates(data)
	return result.Error
}

func (m *defaultStrollRepo) UpdateCover(bvId, cover string) error {
	result := m.conn.Table("stroll").Where("bv = ?", bvId).Update("cover", cover)
	return result.Error
}

func (m *defaultStrollRepo) FindOne(primaryKey uint) (*domain.Stroll, error) {
	stroll := &domain.Stroll{}
	result := m.conn.Raw("SELECT * FROM stroll WHERE id = ? AND deleted_at = 0", primaryKey).First(stroll)
	if result.Error != nil {
		return nil, result.Error
	}
	return stroll, nil
}

func (m *defaultStrollRepo) FindAllByIds(primaryKeyList []uint) ([]*domain.Stroll, error) {
	var strollList []*domain.Stroll
	result := m.conn.Raw("SELECT * FROM stroll WHERE id IN (?) AND deleted_at = 0", primaryKeyList).Find(&strollList)
	return strollList, result.Error
}

func (m *defaultStrollRepo) FindMaxId() (uint, error) {
	stroll := &domain.Stroll{}
	result := m.conn.Raw("SELECT id FROM stroll WHERE deleted_at = 0 ORDER BY id DESC LIMIT 0, 1").Scan(&stroll)
	if result.Error != nil {
		return 0, result.Error
	}

	return stroll.Id, nil
}

func (m *defaultStrollRepo) FindLastUpdateTime() (uint, error) {
	var stroll *domain.Stroll
	result := m.conn.Raw("SELECT * FROM stroll WHERE deleted_at = 0 ORDER BY id DESC LIMIT 0, 1").Find(&stroll)
	if result.Error != nil {
		return 0, result.Error
	}

	return stroll.CreatedAt, nil
}
