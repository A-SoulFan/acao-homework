package domain

type StrollRepo interface {
	// mysql gorm
	Insert(data *Stroll) error
	Delete(primaryKey uint) error
	Update(data *Stroll) error
	UpdateCover(bvId, cover string) error
	FindOne(primaryKey uint) (*Stroll, error)
	FindAllByIds(primaryKeyList []uint) ([]*Stroll, error)
	FindMaxId() (uint, error)
	FindLastUpdateTime() (uint, error)

	// cache
	SetCache([]*Stroll)
	GetCache() []*Stroll
}

type Stroll struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	BV        string `json:"bv" gorm:"uniqueIndex:uq_bv"`
	TargetUrl string `json:"target_url"`
	Play      string `json:"play"`
	CreatedAt uint   `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt uint   `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt uint   `json:"deleted_at" gorm:"index:idx_deleted,uniqueIndex:uq_bv"`
}
