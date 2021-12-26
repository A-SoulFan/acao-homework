package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h *defaultSupportAPIhandler) TeamServiceGetTeamVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.TeamVideosReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		h.teamService.SetDBwithCtx(ctx)

		if resp, err := h.teamService.GetTeamVideos(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func (h *defaultSupportAPIhandler) TeamServiceGetTeamEvents() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.TeamEventsReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		h.teamService.SetDBwithCtx(ctx)

		if resp, err := h.teamService.GetTeamEvents(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}
