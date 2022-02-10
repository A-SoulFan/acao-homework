package banner

import (
	"context"
	"strings"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	"github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"
	"github.com/A-SoulFan/acao-homework/internal/repository"
	"gorm.io/gorm"
)

const (
	allowedTypes = "homepage"
)

type defaultBannerService struct {
	db *gorm.DB
}

func NewDefaultBannerService(db *gorm.DB) idl.BannerService {
	return &defaultBannerService{
		db: db,
	}
}

func (bs *defaultBannerService) GetBannerList(ctx context.Context, req idl.BannerListReq) (*idl.BannerListReply, error) {
	if !checkType(req.Type) {
		return nil, apperrors.NewValidateError("无效的类型")
	}

	bannerRepo := repository.NewBannerRepo(bs.db.WithContext(ctx))

	list, err := bannerRepo.FindAllByType(req.Type)
	if err != nil {
		return nil, err
	}

	resp := &idl.BannerListReply{TotalCount: len(list), PictureList: make([]idl.BannerPicture, 0, len(list))}
	for _, banner := range list {
		resp.PictureList = append(resp.PictureList, idl.BannerPicture{
			PictureUrl:      banner.Url,
			PictureDescribe: banner.Desc,
			Title:           banner.Title,
			Content:         banner.Content,
		})
	}

	return resp, nil
}

func checkType(t string) bool {
	for _, allowedType := range strings.Split(allowedTypes, ",") {
		if allowedType == strings.ToLower(t) {
			return true
		}
	}

	return false
}
