package service

import (
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/service/banner"
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/service/member"
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/service/milestone"
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/service/recommend"
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/service/stroll"
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/service/team"
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

var TaskProviderSet = wire.NewSet(
	milestone.NewDefaultMilestoneTask,
	recommend.NewDefaultRecommendTask,
	stroll.NewDefaultStrollTask,
)
