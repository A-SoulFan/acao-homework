package handler

import (
	"net/http"

	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	strollSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/stroll"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func RandomStrollHandler(svc *svcCtx.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := strollSvc.NewRandomGetLogic(ctx, svc)
		if resp, err := lg.RandomGetStroll(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}

func LastUpdateTimeHandler(svc *svcCtx.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := strollSvc.NewRandomGetLogic(ctx, svc)
		if resp, err := lg.LastUpdateTime(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
