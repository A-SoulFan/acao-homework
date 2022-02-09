package service

import (
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/service/banner"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/service/member"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/service/milestone"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/service/recommend"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/service/stroll"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/service/team"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	banner.NewDefaultBannerService,
	member.NewDefaultMemberService,
	milestone.NewDefaultMilestoneService,
	recommend.NewDefaultRecommendService,
	stroll.NewDefaultStrollService,
	team.NewDefaultTeamService,
)
