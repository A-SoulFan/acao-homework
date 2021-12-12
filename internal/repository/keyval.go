package repository

import (
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"gorm.io/gorm"
)

const (
	keyValueTableName = "key_val"
)

type defaultKeyValueRepo struct {
	conn *gorm.DB
}

func NewDefaultKeyValueRepo(conn *gorm.DB) domain.KeyValueRepo {
	return &defaultKeyValueRepo{conn: conn}
}

func (m *defaultKeyValueRepo) FindAllByKey(key string) ([]*domain.KeyValue, error) {
	var list []*domain.KeyValue
	result := m.conn.Table(keyValueTableName).Where("key = ? AND deleted_at = 0", key).
		Order("sort").
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}

func (m *defaultKeyValueRepo) FindOneByKey(key string) (*domain.KeyValue, error) {
	var data *domain.KeyValue
	result := m.conn.Table(keyValueTableName).Where("`key` = ? AND deleted_at = 0", key).
		Order("sort").
		Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
