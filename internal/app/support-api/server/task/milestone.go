package task

func (ts *defaultSupportAPItaskController) RegisterMilestoneTask() {
	ts.milestoneService.InitTask()
}
