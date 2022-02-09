package repository

import (
	"github.com/A-SoulFan/acao-homework/internal/domain"
	"gorm.io/gorm"
)

const (
	bannerTableName = "banner"
)

type defaultBannerRepo struct {
	conn *gorm.DB
}

func NewBannerRepo(conn *gorm.DB) domain.BannerRepo {
	return &defaultBannerRepo{conn: conn}
}

func (m *defaultBannerRepo) FindAllByType(t string) ([]*domain.Banner, error) {
	var list []*domain.Banner
	result := m.conn.Table(bannerTableName).Where("type = ? AND deleted_at = 0", t).Order("sort").Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
