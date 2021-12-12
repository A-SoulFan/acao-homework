package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/types"

	bannerSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/banner"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"

	"github.com/gin-gonic/gin"
)

func GetBannerListHandler(svc *svcCtx.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.BannerListReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		lg := bannerSvc.NewBannerListLogic(ctx, svc)
		if resp, err := lg.GetList(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

type ASFResponse struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func ErrorResponse(err error) *ASFResponse {
	return &ASFResponse{Code: -1, ErrMsg: err.Error()}
}

func SuccessResponse(data interface{}) *ASFResponse {
	return &ASFResponse{Code: 0, ErrMsg: "ok", Data: data}
}
