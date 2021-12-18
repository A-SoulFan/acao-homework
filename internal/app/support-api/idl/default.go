package idl

import (
	"gorm.io/gorm"
)

type defaultDB interface {
	SetDB(db *gorm.DB)
}
