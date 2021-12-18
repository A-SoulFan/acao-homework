package handler

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	milestoneSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/milestone"
	"github.com/A-SoulFan/acao-homework/internal/repository"
)

type defaultSupportAPIhandler struct {
	stx              *svcCtx.ServiceContext
	milestoneService idl.MilestoneService
}

func NewDefaultSupportAPIhandler(stx *svcCtx.ServiceContext) idl.SupportAPIhandler {
	// repo
	milestoneRepo := repository.NewMilestoneRepo()

	// service
	milestoneService := milestoneSvc.NewDefaultMilestoneService(stx, milestoneRepo)

	return &defaultSupportAPIhandler{
		stx:              stx,
		milestoneService: milestoneService,
	}
}
