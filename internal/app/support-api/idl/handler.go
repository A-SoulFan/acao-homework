package idl

import (
	"github.com/gin-gonic/gin"
)

type SupportAPIhandler interface {
	MilestoneServiceNextGroup() gin.HandlerFunc

	StrollServiceRandomStroll() gin.HandlerFunc
	StrollServiceLastUpdateTime() gin.HandlerFunc

	RecommendServiceTopRecommendSlices() gin.HandlerFunc

	MemberServiceGetAllMembers() gin.HandlerFunc
	MemberServiceGetMemberExperience() gin.HandlerFunc
	MemberServiceGetMemberVideos() gin.HandlerFunc

	TeamServiceGetTeamVideos() gin.HandlerFunc
	TeamServiceGetTeamEvents() gin.HandlerFunc

	BannerServiceGetBannerList() gin.HandlerFunc
}
