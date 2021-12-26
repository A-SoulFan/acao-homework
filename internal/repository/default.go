package repository

import (
	"gorm.io/gorm"
)

type defaultRepo struct {
	conn *gorm.DB
}

func (m *defaultRepo) SetDB(db *gorm.DB) {
	m.conn = db
}
