package types

type StrollLastUpdateReply struct {
	LastUpdateTime uint `json:"last_update_time"`
}

type StrollReply struct {
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	BV        string `json:"bv"`
	PlayUrl   string `json:"play_url"`
	TargetUrl string `json:"target_url"`
	CreatedAt uint   `json:"created_at"`
}

type BannerListReq struct {
	Type string `form:"type,default=homepage" binding:"omitempty"`
}

type BannerPicture struct {
	PictureUrl      string `json:"pictureUrl"`
	PictureDescribe string `json:"pictureDescribe"`
	Title           string `json:"title"`
	Content         string `json:"content"`
}

type BannerListReply struct {
	TotalCount  int             `json:"totalCount"`
	PictureList []BannerPicture `json:"pictureList"`
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
