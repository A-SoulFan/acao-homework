package idl

import "context"

type MemberService interface {
	GetAllMembers(ctx context.Context) (*MemberAll, error)
	GetMemberExperience(ctx context.Context, req MemberExperienceReq) (*MemberExperienceResp, error)
	GetMemberVideos(ctx context.Context, req MemberVideoReq) (*MemberExperienceResp, error)
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
