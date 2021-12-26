package repository

import (
	"github.com/A-SoulFan/acao-homework/internal/domain"
)

const (
	keyValueTableName = "key_val"
)

type defaultKeyValueRepo struct {
	defaultRepo
}

func NewKeyValueRepo() domain.KeyValueRepo {
	return &defaultKeyValueRepo{}
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
