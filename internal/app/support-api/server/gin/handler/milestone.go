package handler

import (
	"net/http"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	milestoneSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/milestone"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/types"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func MilestoneNextGroup(svc *svcCtx.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.NextGroupReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		lg := milestoneSvc.NewGroupLogic(ctx, svc)
		if resp, err := lg.NextGroup(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
