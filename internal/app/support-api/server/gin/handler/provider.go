package handler

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewBannerApi,
	NewMemberApi,
	NewMilestoneApi,
	NewRecommendApi,
	NewStrollApi,
	NewTeamApi,
)
