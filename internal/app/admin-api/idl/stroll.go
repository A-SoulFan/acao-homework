package idl

import "context"

type StrollService interface {
	Create(ctx context.Context, req StrollCreateReq) error
}

type StrollCreateReq struct {
	Bv    string `json:"bv"`
	Pages uint8  `json:"pages"`
}
