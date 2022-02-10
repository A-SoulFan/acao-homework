package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type Team struct {
	teamService idl.TeamService
}

func NewTeamApi(teamService idl.TeamService) *Team {
	return &Team{
		teamService: teamService,
	}
}

func (h *Team) TeamServiceGetTeamVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.TeamVideosReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		if resp, err := h.teamService.GetTeamVideos(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func (h *Team) TeamServiceGetTeamEvents() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.TeamEventsReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		if resp, err := h.teamService.GetTeamEvents(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}
