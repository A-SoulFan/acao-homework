package logic

import (
	"context"
	"strings"

	"github.com/A-SoulFan/support-api/app/service"
	"github.com/A-SoulFan/support-api/app/service/banner/model"
	"github.com/A-SoulFan/support-api/domain"
	appErr "github.com/A-SoulFan/support-api/lib/err"
	"github.com/A-SoulFan/support-api/types"

	"gorm.io/gorm"
)

const (
	allowedTypes = "homepage"
)

type BannerLogic struct {
	ctx    context.Context
	svcCtx *service.Context
	dbCtx  *gorm.DB
}

func NewBannerListLogic(ctx context.Context, svcCtx *service.Context) domain.BannerLogic {
	return BannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.Db.WithContext(ctx),
	}
}

func (b *BannerLogic) GetList(req types.BannerListReq) (*types.BannerListReply, error) {
	if !checkType(req.Type) {
		return nil, appErr.NewError("无效的类型")
	}

	list, err := model.NewDefaultBannerModel(b.dbCtx).FindAllByType(req.Type)
	if err != nil {
		b.svcCtx.Logger.Error(err)
		return nil, appErr.NewError("服务器异常，请稍后再试")
	}

	resp := &types.BannerListReply{TotalCount: len(list), PictureList: make([]types.BannerPicture, 0, len(list))}
	for _, banner := range list {
		resp.PictureList = append(resp.PictureList, types.BannerPicture{
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
