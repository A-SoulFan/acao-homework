package idl

import "context"

type defaultDB interface {
	SetDBwithCtx(ctx context.Context)
}
