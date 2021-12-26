package domain

type BannerRepo interface {
	defaultRepo
	FindAllByType(t string) ([]*Banner, error)
}

type Banner struct {
	Id        uint   `json:"id"`
	Type      string `json:"type"`
	Sort      uint   `json:"sort"`
	Url       string `json:"url"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Content   string `json:"content"`
	DeletedAt uint   `json:"-"`
}
