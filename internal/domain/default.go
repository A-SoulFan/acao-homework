package domain

import (
	"gorm.io/gorm"
)

type defaultRepo interface {
	SetDB(db *gorm.DB)
}
