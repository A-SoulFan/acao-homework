package router

import (
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/server/gin/handler"
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/server/gin/middleware"
	"github.com/A-SoulFan/acao-homework/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
)

// 根据业务可以拆分不同 router
func InitRouter(
	errorInterceptor *middleware.ErrorInterceptor,
	strollApi *handler.Stroll,
	milestoneApi *handler.Milestone,
	recommendApi *handler.Recommend,
	bannerApi *handler.Banner,
	memberApi *handler.Member,
	TeamApi *handler.Team,
) http.InitRouters {
	return func(r *gin.Engine) {

		r.Use(errorInterceptor.Handler())
		r.Use(middleware.Cors())

		// 随机溜
		r.GET("/api/stroll/random", strollApi.StrollServiceRandomStroll())
		r.GET("/api/stroll/last-update-time", strollApi.StrollServiceLastUpdateTime())

		// 大事件
		r.GET("/api/milestone/next-group", milestoneApi.MilestoneServiceNextGroup())

		// 下列为新人指南相关API 暂时实现在这里
		// 注意 response request 风格均不相同

		// 推荐切片
		r.GET("/api/recommend-slice", recommendApi.RecommendServiceTopRecommendSlices())

		// 头部图片
		r.GET("/asf/mobile/headpicture", bannerApi.BannerServiceGetBannerList())

		// 团队成员
		r.GET("/asf/mobile/member/all", memberApi.MemberServiceGetAllMembers())
		// 团队个人经历
		r.GET("/asf/mobile/member/experience", memberApi.MemberServiceGetMemberExperience())
		// 个人作品
		r.GET("/asf/mobile/member/videos", memberApi.MemberServiceGetMemberVideos())
		// 团队作品
		r.GET("/asf/mobile/team/videos", TeamApi.TeamServiceGetTeamVideos())
		// 团队事件
		r.GET("/asf/mobile/team/events", TeamApi.TeamServiceGetTeamEvents())
	}
}
