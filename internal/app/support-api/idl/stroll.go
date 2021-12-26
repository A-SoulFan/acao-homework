package idl

type StrollTask interface {
	defaultDB
	InitTask()
}

type StrollService interface {
	StrollTask
	RandomGetStroll() (*StrollReply, error)
	LastUpdateTime() (*StrollLastUpdateReply, error)
}

type StrollReply struct {
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	BV        string `json:"bv"`
	PlayUrl   string `json:"play_url"`
	TargetUrl string `json:"target_url"`
	CreatedAt uint   `json:"created_at"`
}

type StrollLastUpdateReply struct {
	LastUpdateTime uint `json:"last_update_time"`
}
