package idl

type MemberService interface {
	defaultDB
	GetAllMembers() (*MemberAll, error)
	GetMemberExperience(req MemberExperienceReq) (*MemberExperienceResp, error)
	GetMemberVideos(req MemberVideoReq) (*MemberExperienceResp, error)
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
