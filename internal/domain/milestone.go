package domain

type MilestoneRepo interface {
	// mysql gorm
	Insert(data *Milestone) error
	Delete(primaryKey uint) error
	Update(data *Milestone) error
	SearchTitles(keyword string, limit uint) ([]*Milestone, error)
	FindOne(primaryKey uint) (*Milestone, error)
	FindAllByTimestamp(startTimestamp, limit uint, order string) ([]*Milestone, error)

	// cache
	SetCache([]*Milestone)
	GetCache() []*Milestone
}

type Milestone struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Subtitled string `json:"subtitled"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	TargetUrl string `json:"target_url"`
	Timestamp uint   `json:"timestamp"`
	CreatedAt uint   `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt uint   `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt uint   `json:"deleted_at" gorm:"index:idx_deleted"`
}
