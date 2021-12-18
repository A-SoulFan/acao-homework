package handler

import (
	"net/http"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	recommendSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/recommend"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func RecommendHandler(svc *svcCtx.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := recommendSvc.NewRecommendSliceLogic(svc)
		if resp, err := lg.Handle(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
