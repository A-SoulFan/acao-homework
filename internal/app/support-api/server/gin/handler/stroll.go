package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *defaultSupportAPIhandler) StrollServiceRandomStroll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.strollService.SetDBwithCtx(ctx)

		if resp, err := h.strollService.RandomGetStroll(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}

func (h *defaultSupportAPIhandler) StrollServiceLastUpdateTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.strollService.SetDBwithCtx(ctx)

		if resp, err := h.strollService.LastUpdateTime(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
