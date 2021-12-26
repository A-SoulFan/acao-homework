package idl

type TeamService interface {
	defaultDB
	GetTeamVideos(req TeamVideosReq) (*TeamVideosResp, error)
	GetTeamEvents(req TeamEventsReq) (*TeamEventsResp, error)
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
