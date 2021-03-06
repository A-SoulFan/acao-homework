package idl

import "context"

type BannerService interface {
	GetBannerList(ctx context.Context, req BannerListReq) (*BannerListReply, error)
}

type BannerListReq struct {
	Type string `form:"type,default=homepage" binding:"omitempty"`
}

type BannerPicture struct {
	PictureUrl      string `json:"pictureUrl"`
	PictureDescribe string `json:"pictureDescribe"`
	Title           string `json:"title"`
	Content         string `json:"content"`
}

type BannerListReply struct {
	TotalCount  int             `json:"totalCount"`
	PictureList []BannerPicture `json:"pictureList"`
}
