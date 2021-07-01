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
	data map[string]*timedcacheitem
	mtx  sync.RWMutex
}

func NewTimedcache() *Timedcache {
	return &Timedcache{
		data: make(map[string]*timedcacheitem),
	}
}

func (tc *Timedcache) Get(key string, fetchfunc func() (x interface{}, kt time.Duration, err error)) (x interface{}, err error) {
	if exists, x := func() (bool, interface{}) {
		tc.mtx.RLock()
		defer tc.mtx.RUnlock()
		if item := tc.data[key]; item != nil && item.et.After(time.Now()) {
			return true, item.x
		}
		return false, nil
	}(); exists {
		return x, nil
	}

	tc.mtx.Lock()
	defer tc.mtx.Unlock()
	if item := tc.data[key]; item != nil && item.et.After(time.Now()) {
		return item.x, nil
	}

	x, kt, err := fetchfunc()
	if err != nil {
		return
	}
	item := timedcacheitem{
		x:  x,
		et: time.Now().Add(kt),
	}
	tc.data[key] = &item
	return
}
