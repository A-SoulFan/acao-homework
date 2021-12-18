package repository

import (
	"gorm.io/gorm"
)

type defaultRepo struct {
	conn *gorm.DB
}

func (m *defaultMilestoneRepo) SetDB(db *gorm.DB) {
	m.conn = db
}
