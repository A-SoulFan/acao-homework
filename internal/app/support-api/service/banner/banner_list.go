package banner

import (
	"context"
	"strings"

	"github.com/A-SoulFan/acao-homework/internal/domain"
	"github.com/A-SoulFan/acao-homework/internal/pkg/apperrors"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
)

const (
	allowedTypes = "homepage"
)

type defaultBannerService struct {
	svcCtx     *svcCtx.ServiceContext
	bannerRepo domain.BannerRepo
}

func NewDefaultBannerService(svcCtx *svcCtx.ServiceContext, bannerRepo domain.BannerRepo) idl.BannerService {
	return &defaultBannerService{
		svcCtx:     svcCtx,
		bannerRepo: bannerRepo,
	}
}

func (bs *defaultBannerService) SetDBwithCtx(ctx context.Context) {
	db := bs.svcCtx.WithDatabaseContext(ctx)
	bs.bannerRepo.SetDB(db)
}

func (bs *defaultBannerService) GetBannerList(req idl.BannerListReq) (*idl.BannerListReply, error) {
	if !checkType(req.Type) {
		return nil, apperrors.NewValidateError("无效的类型")
	}

	list, err := bs.bannerRepo.FindAllByType(req.Type)
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
