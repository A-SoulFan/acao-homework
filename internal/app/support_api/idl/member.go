package idl

import (
	"context"

	"github.com/A-SoulFan/acao-homework/internal/domain"
)

type MemberService interface {
	GetAllMembers(ctx context.Context) (*MemberAll, error)
	GetMemberExperience(ctx context.Context, req MemberExperienceReq) (*MemberExperienceResp, error)
	GetMemberVideos(ctx context.Context, req MemberVideoReq) (*MemberExperienceResp, error)
	GetMemberDebts(ctx context.Context, req MemberDebtReq) ([]*domain.Debt, error)
}

type MemberAll struct {
	MemberList interface{} `json:"memberList"`
}

type MemberExperienceReq struct {
	MemberName string `form:"memberName" binding:"required,oneof=ava bella carol diana eileen"`
}

type MemberExperienceResp struct {
	MemberName string      `json:"memberName"`
	TotalCount int         `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
	VideoList  interface{} `json:"videoList"`
}

type MemberVideoReq struct {
	MemberName string `form:"memberName" binding:"required,oneof=ava bella carol diana eileen"`
}

type MemberVideoResp struct {
	MemberName string      `json:"memberName"`
	TotalCount int         `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
	VideoList  interface{} `json:"videoList"`
}

type MemberDebtReq struct {
	MemberName string `form:"memberName" binding:"omitempty"`
}
