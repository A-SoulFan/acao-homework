package task

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	milestoneSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/milestone"
	"github.com/A-SoulFan/acao-homework/internal/repository"
)

type defaultSupportAPItaskController struct {
	milestoneService idl.MilestoneService
}

func NewDefaultSupportAPItaskController(stx *svcCtx.ServiceContext) idl.SupportAPItaskController {
	// repo
	milestoneRepo := repository.NewMilestoneRepo()

	// service
	milestoneService := milestoneSvc.NewDefaultMilestoneService(stx, milestoneRepo)

	return &defaultSupportAPItaskController{
		milestoneService: milestoneService,
	}
}
