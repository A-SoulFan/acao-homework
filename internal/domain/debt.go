package domain

type DebtRepo interface {
	// cache
	SetCache(key string, debts []*Debt)
	GetCache(key string) []*Debt
}

type Debt struct {
	StartDate        string `json:"start_date"`         // 开始日期
	Subject      string `json:"subject"`       // 主体
	StartUrl       string `json:"start_url"` // 开始地址
	CompletionUrl   string `json:"completion_url"`   // 完成地址
	CompletionDate       string `json:"completion_date"`        // 完成日期
	Tags    []Tag `json:"tags"`    // 播放量
}

type Tag struct {
	Key string `json:"key"` // 名字
	Name string `json:"name"` // 名字
	IsLeader bool `json:"is_leader"` // 是否是发起人
	HomePage string `json:"home_page"`// 个人主页
}
