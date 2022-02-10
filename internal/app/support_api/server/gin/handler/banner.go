package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"

	"github.com/gin-gonic/gin"
)

type Banner struct {
	bannerService idl.BannerService
}

func NewBannerApi(bannerService idl.BannerService) *Banner {
	return &Banner{
		bannerService: bannerService,
	}
}

func (h *Banner) BannerServiceGetBannerList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.BannerListReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		if resp, err := h.bannerService.GetBannerList(ctx, req); err != nil {
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
