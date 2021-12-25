package router

import (
	"github.com/A-SoulFan/acao-homework/internal/app/admin-api/context"
	"github.com/A-SoulFan/acao-homework/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(svc *context.ServiceContext, logger *zap.Logger) http.InitRouters {
	return func(r *gin.Engine) {

	}
}
