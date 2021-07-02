package util

import (
	"sync"
	"time"
)

type timedcacheitem struct {
	x   interface{}
	et  time.Time
	mtx sync.Mutex
}

type Timedcache struct {
	data sync.Map
}

func NewTimedcache() *Timedcache {
	return &Timedcache{}
}

func (tc *Timedcache) Get(key string, fetchfunc func() (x interface{}, kt time.Duration, err error)) (x interface{}, err error) {
	var item *timedcacheitem
	actual, _ := tc.data.Load(key)
	if actual != nil {
		item = actual.(*timedcacheitem)
		if item.et.After(time.Now()) {
			return item.x, nil
		}
	} else {
		actual, _ = tc.data.LoadOrStore(key, &timedcacheitem{})
		item = actual.(*timedcacheitem)
	}
	WithMutex(&item.mtx, func() {
		if item.et.After(time.Now()) {
			x = item.x
			return
		}
		var kt time.Duration
		x, kt, err = fetchfunc()
		if err != nil {
			return
		}
		item.x = x
		item.et = time.Now().Add(kt)
	})
	return
}
