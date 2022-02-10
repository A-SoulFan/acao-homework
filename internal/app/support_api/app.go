package support_api

import (
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/server/gin"
	"github.com/A-SoulFan/acao-homework/internal/app/support_api/server/task"
)

type App struct {
	ginServer  *gin.Server
	taskServer *task.Tasks
}

func NewApp(ginServer *gin.Server,
	recommendTask idl.RecommendService,
	milestoneTask idl.MilestoneService,
	strollTask idl.StrollService,
) *App {
	taskServer := task.NewTasks()
	taskServer.Register(recommendTask, milestoneTask, strollTask)

	return &App{
		ginServer:  ginServer,
		taskServer: taskServer,
	}
}

func (app *App) Run() error {
	if err := app.ginServer.Start(); err != nil {
		return err
	}

	app.taskServer.Run()
	return nil
}

func (app *App) AwaitSignal() {
	app.ginServer.AwaitSignal()
}
