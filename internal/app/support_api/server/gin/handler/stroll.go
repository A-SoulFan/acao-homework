package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"

	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type Stroll struct {
	strollService idl.StrollService
}

func NewStrollApi(strollService idl.StrollService) *Stroll {
	return &Stroll{
		strollService: strollService,
	}
}

func (h *Stroll) StrollServiceRandomStroll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if resp, err := h.strollService.RandomGetStroll(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}

func (h *Stroll) StrollServiceLastUpdateTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if resp, err := h.strollService.LastUpdateTime(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
