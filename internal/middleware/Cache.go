package middleware

import "time"

type MotionCache struct {
	store []Store
}

type Store struct {
	Key        string
	Info       interface{}
	expiration time.Time
}

func (m *MotionCache) Get(i string) {
	for _, v := range m.store {
		if v.expiration.After(time.Now()) {

		}

	}
}
