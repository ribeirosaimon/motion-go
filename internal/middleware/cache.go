package middleware

import (
	"github.com/ribeirosaimon/motion-go/internal/domain/cache"
	"time"
)

type MotionCache struct {
	after  *Store
	before *Store
	Size   uint16
}

type Store struct {
	Key        string
	Info       interface{}
	expiration time.Time
	next       *Store
}

func (m *MotionCache) Get(i string) {
	for _, v := range m.store {
		if v.expiration.After(time.Now()) {
			v = nil
		}
	}
}

func (m *MotionCache) Add(company cache.Company) {
	if m.Size == 0 {
		var store = &Store{
			Key:        company.Name,
			Info:       company.Code,
			expiration: time.Now().Add(time.Minute),
		}
		m.before, m.after = nil, store
	} else {
		var store = Store{
			Key:        company.Name,
			Info:       company.Code,
			expiration: time.Now().Add(time.Minute),
		}
	}
}
