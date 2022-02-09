package idl

import "context"

type TeamService interface {
	GetTeamVideos(ctx context.Context, req TeamVideosReq) (*TeamVideosResp, error)
	GetTeamEvents(ctx context.Context, req TeamEventsReq) (*TeamEventsResp, error)
}

type TeamVideosReq struct {
	//MemberName string `form:"memberName" binding:"required,oneof=ava bella carol diana eileen"`
}

type TeamVideosResp struct {
	TotalCount int         `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
	VideoList  interface{} `json:"videoList"`
}

type TeamEventsReq struct {
	Year string `form:"year,default=2021" binding:"omitempty,numeric"`
}

type TeamEventsResp struct {
	TotalCount int         `json:"totalCount"`
	TotalPage  int         `json:"totalPage"`
	EventList  interface{} `json:"eventList"`
}
