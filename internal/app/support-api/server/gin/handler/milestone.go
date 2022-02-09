package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type Milestone struct {
	milestoneService idl.MilestoneService
}

func NewMilestoneApi(milestoneService idl.MilestoneService) *Milestone {
	return &Milestone{
		milestoneService: milestoneService,
	}
}

func (h *Milestone) MilestoneServiceNextGroup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.NextGroupReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		if resp, err := h.milestoneService.NextGroup(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
