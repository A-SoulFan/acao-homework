package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *defaultSupportAPIhandler) RecommendServiceTopRecommendSlices() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if resp, err := h.recommendService.TopRecommendSlices(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
