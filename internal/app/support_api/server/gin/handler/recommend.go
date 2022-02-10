package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"

	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type Recommend struct {
	recommendService idl.RecommendService
}

func NewRecommendApi(recommendService idl.RecommendService) *Recommend {
	return &Recommend{
		recommendService: recommendService,
	}
}

func (h *Recommend) RecommendServiceTopRecommendSlices() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if resp, err := h.recommendService.TopRecommendSlices(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
