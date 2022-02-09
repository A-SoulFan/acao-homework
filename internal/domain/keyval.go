package domain

type KeyValueRepo interface {
	FindAllByKey(key string) ([]*KeyValue, error)
	FindOneByKey(key string) (*KeyValue, error)
}

type KeyValue struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Value     []byte `json:"value"`
	Sort      uint   `json:"sort"`
	DeletedAt uint   `json:"-"`
}
