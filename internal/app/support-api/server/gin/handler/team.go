package handler

import (
	"net/http"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	keyvalSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/keyval"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/types"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func GetTeamVideoListHandler(svc *svcCtx.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TeamVideosReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		lg := keyvalSvc.NewTeamLogic(ctx, svc)
		if resp, err := lg.GetVideos(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func GetTeamEventListHandler(svc *svcCtx.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TeamEventsReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		lg := keyvalSvc.NewTeamLogic(ctx, svc)
		if resp, err := lg.GetEvents(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}
