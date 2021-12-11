package mysql

import (
	"github.com/A-SoulFan/support-api/domain"
	"gorm.io/gorm"
)

const (
	bannerTableName = "banner"
)

type (
	defaultBannerModel struct {
		conn *gorm.DB
	}
)

func NewDefaultBannerModel(conn *gorm.DB) domain.BannerModel {
	return &defaultBannerModel{conn: conn}
}

func (m *defaultBannerModel) FindAllByType(t string) ([]*domain.Banner, error) {
	var list []*domain.Banner
	result := m.conn.Table(bannerTableName).Where("type = ? AND deleted_at = 0", t).Order("sort").Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
