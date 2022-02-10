package task

import (
	"context"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewTasks)

type Tasks struct {
	list []idl.Task
}

func NewTasks() *Tasks {
	return &Tasks{
		list: make([]idl.Task, 0, 6),
	}
}

func (t *Tasks) Register(task ...idl.Task) {
	t.list = append(t.list, task...)
}

func (t *Tasks) Run() {
	ctx := context.Background()
	for i := range t.list {
		go t.list[i].InitTask(ctx)
	}
}
