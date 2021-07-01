package util

import (
	"sync"
	"time"
)

type timedcacheitem struct {
	x  interface{}
	et time.Time
}

type Timedcache struct {
	data sync.Map
	mtx  sync.Mutex
}

func NewTimedcache() *Timedcache {
	return &Timedcache{}
}

func (tc *Timedcache) Get(key string, fetchfunc func() (x interface{}, kt time.Duration, err error)) (x interface{}, err error) {
	actual, _ := tc.data.Load(key)
	if actual != nil {
		item := actual.(*timedcacheitem)
		if item.et.After(time.Now()) {
			return item.x, nil
		}
	}

	tc.mtx.Lock()
	defer tc.mtx.Unlock()

	actual, _ = tc.data.Load(key)
	if actual != nil {
		item := actual.(*timedcacheitem)
		if item.et.After(time.Now()) {
			return item.x, nil
		}
	}

	x, kt, err := fetchfunc()
	if err != nil {
		return nil, err
	}

	tc.data.Store(key, &timedcacheitem{
		x:  x,
		et: time.Now().Add(kt),
	})
	return x, nil
}
