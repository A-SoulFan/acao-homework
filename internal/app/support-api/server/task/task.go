package task

import (
	svcCtx "github.com/A-SoulFan/acao-homework/internal/app/support-api/context"
	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	milestoneSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/milestone"
	recommendSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/recommend"
	strollSvc "github.com/A-SoulFan/acao-homework/internal/app/support-api/service/stroll"
	"github.com/A-SoulFan/acao-homework/internal/repository"
)

type defaultSupportAPItaskController struct {
	milestoneService idl.MilestoneService
	strollService    idl.StrollService
	recommendService idl.RecommendService
}

func NewDefaultSupportAPItaskController(stx *svcCtx.ServiceContext) idl.SupportAPItaskController {
	// repo
	milestoneRepo := repository.NewMilestoneRepo()
	strollRepo := repository.NewStrollRepo()
	recommendRepo := repository.NewRecommendRepo()

	// service
	milestoneService := milestoneSvc.NewDefaultMilestoneService(stx, milestoneRepo)
	strollService := strollSvc.NewDefaultStrollService(stx, strollRepo)
	recommendService := recommendSvc.NewDefaultRecommendService(stx, recommendRepo)

	return &defaultSupportAPItaskController{
		milestoneService: milestoneService,
		strollService:    strollService,
		recommendService: recommendService,
	}
}

func (ts *defaultSupportAPItaskController) RegisterMilestoneTask() {
	ts.milestoneService.InitTask()
}

func (ts *defaultSupportAPItaskController) RegisterStrollTask() {
	ts.strollService.InitTask()
}

func (ts *defaultSupportAPItaskController) RegisterRecommendTask() {
	ts.recommendService.InitTask()
}
