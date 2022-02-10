package idl

import "context"

type Task interface {
	InitTask(ctx context.Context)
}
