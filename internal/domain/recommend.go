package domain

type RecommendRepo interface {
	// cache
	SetCache([]*RecommendVideo)
	GetCache() []*RecommendVideo
}

type RecommendVideo struct {
	Bid        string `json:"bid"`         // BV 号
	Title      string `json:"title"`       // 标题
	Desc       string `json:"description"` // 简介
	ImageUrl   string `json:"image_url"`   // 封面链接
	Auth       string `json:"auth"`        // 作者
	PlayVal    string `json:"play_val"`    // 播放量
	Time       string `json:"time"`        // 时长
	TimeSecond string `json:"time_second"` // 时长（单位秒）
	Url        string `json:"url"`         // 原视频的跳转链接
}
