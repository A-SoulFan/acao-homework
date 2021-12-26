package handler

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	bannerSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/banner"
	memberSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/member"
	milestoneSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/milestone"
	recommendSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/recommend"
	strollSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/stroll"
	teamSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/team"
	"github.com/A-SoulFan/acao-homework/internal/repository"
)

type defaultSupportAPIhandler struct {
	milestoneService idl.MilestoneService
	strollService    idl.StrollService
	recommendService idl.RecommendService
	teamService      idl.TeamService
	memberService    idl.MemberService
	bannerService    idl.BannerService
}

func NewDefaultSupportAPIhandler(stx *svcCtx.ServiceContext) idl.SupportAPIhandler {
	// repo
	milestoneRepo := repository.NewMilestoneRepo()
	strollRepo := repository.NewStrollRepo()
	recommendRepo := repository.NewRecommendRepo()
	kvRepo := repository.NewKeyValueRepo()
	bannerRepo := repository.NewBannerRepo()

	// service
	milestoneService := milestoneSvc.NewDefaultMilestoneService(stx, milestoneRepo)
	strollService := strollSvc.NewDefaultStrollService(stx, strollRepo)
	recommendService := recommendSvc.NewDefaultRecommendService(stx, recommendRepo)
	teamService := teamSvc.NewDefaultTeamService(stx, kvRepo)
	memberService := memberSvc.NewDefaultMemberService(stx, kvRepo)
	bannerService := bannerSvc.NewDefaultBannerService(stx, bannerRepo)

	return &defaultSupportAPIhandler{
		milestoneService: milestoneService,
		strollService:    strollService,
		recommendService: recommendService,
		teamService:      teamService,
		memberService:    memberService,
		bannerService:    bannerService,
	}
}
