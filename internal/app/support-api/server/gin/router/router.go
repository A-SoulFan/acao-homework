package router

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/server/gin/handler"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/server/gin/middleware"
	"github.com/A-SoulFan/acao-homework/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
)

// 根据业务可以拆分不同 router
func InitRouter(svc *svcCtx.ServiceContext) http.InitRouters {
	return func(r *gin.Engine) {
		r.Use(middleware.Cors())

		// // 随机溜
		// r.GET("/api/stroll/random", handler.RandomStrollHandler(svc))
		// r.GET("/api/stroll/last-update-time", handler.LastUpdateTimeHandler(svc))

		// 大事件
		r.GET("/api/milestone/next-group", handler.MilestoneNextGroup(svc))

		// 下列为新人指南相关API 暂时实现在这里
		// 注意 response request 风格均不相同

		// // 推荐切片
		// r.GET("/api/recommend-slice", handler.RecommendHandler(svc))

		// 头部图片
		r.GET("/asf/mobile/headpicture", handler.GetBannerListHandler(svc))

		// 团队成员
		r.GET("/asf/mobile/member/all", handler.GetAllHandler(svc))
		// 团队个人经历
		r.GET("/asf/mobile/member/experience", handler.GetExperienceListHandler(svc))
		// 个人作品
		r.GET("/asf/mobile/member/videos", handler.GetVideoListHandler(svc))
		// 团队作品
		r.GET("/asf/mobile/team/videos", handler.GetTeamVideoListHandler(svc))
		// 团队事件
		r.GET("/asf/mobile/team/events", handler.GetTeamEventListHandler(svc))
	}
}
