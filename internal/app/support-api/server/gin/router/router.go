package router

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/server/gin/handler"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/server/gin/middleware"
	"github.com/A-SoulFan/acao-homework/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 根据业务可以拆分不同 router
func InitRouter(svc *svcCtx.ServiceContext, logger *zap.Logger) http.InitRouters {
	return func(r *gin.Engine) {

		defaultHandler := handler.NewDefaultSupportAPIhandler(svc)

		r.Use(middleware.NewErrorInterceptor(logger).Handler())
		r.Use(middleware.Cors())

		// 随机溜
		r.GET("/api/stroll/random", defaultHandler.StrollServiceRandomStroll())
		r.GET("/api/stroll/last-update-time", defaultHandler.StrollServiceLastUpdateTime())

		// 大事件
		r.GET("/api/milestone/next-group", defaultHandler.MilestoneServiceNextGroup())

		// 下列为新人指南相关API 暂时实现在这里
		// 注意 response request 风格均不相同

		// 推荐切片
		r.GET("/api/recommend-slice", defaultHandler.RecommendServiceTopRecommendSlices())

		// 头部图片
		r.GET("/asf/mobile/headpicture", defaultHandler.BannerServiceGetBannerList())

		// 团队成员
		r.GET("/asf/mobile/member/all", defaultHandler.MemberServiceGetAllMembers())
		// 团队个人经历
		r.GET("/asf/mobile/member/experience", defaultHandler.MemberServiceGetMemberExperience())
		// 个人作品
		r.GET("/asf/mobile/member/videos", defaultHandler.MemberServiceGetMemberVideos())
		// 团队作品
		r.GET("/asf/mobile/team/videos", defaultHandler.TeamServiceGetTeamVideos())
		// 团队事件
		r.GET("/asf/mobile/team/events", defaultHandler.TeamServiceGetTeamEvents())
	}
}
