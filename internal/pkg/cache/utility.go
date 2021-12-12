package cache

import "github.com/google/uuid"

func GeneratedRandomKey(prefix string) (string, error) {
	if _uuid, err := uuid.NewRandom(); err != nil {
		return "", nil
	} else {
		return prefix + _uuid.String(), nil
	}
}
