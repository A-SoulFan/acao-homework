package domain

import (
	"github.com/A-SoulFan/support-api/types"
)

type BannerModel interface {
	FindAllByType(t string) ([]*Banner, error)
}

type BannerLogic interface {
	GetList(req types.BannerListReq) (*types.BannerListReply, error)
}

type Banner struct {
	Id        uint   `json:"id"`
	Type      string `json:"type"`
	Sort      uint   `json:"sort"`
	Url       string `json:"url"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Content   string `json:"content"`
	DeletedAt uint   `json:"-"`
}
