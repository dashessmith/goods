package goods

import (
	"sync"
	"time"
)

type timedcacheitem struct {
	x        interface{}
	fetching bool
	et       time.Time
	mtx      sync.Mutex
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

func (tc *Timedcache) AsyncGet(key string, asyncDefault interface{}, fetchfunc func() (x interface{}, kt time.Duration, err error)) (x interface{}, err error) {
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
		WithMutex(&item.mtx, func() {
			if item.x == nil {
				item.x = asyncDefault
			}
		})
	}
	x = item.x
	if item.fetching {
		return
	}
	WithMutex(&item.mtx, func() {
		if item.et.After(time.Now()) {
			return
		}
		if item.fetching {
			return
		}
		item.fetching = true
		go func() {
			defer func() {
				item.fetching = false
			}()
			var kt time.Duration
			x, kt, err = fetchfunc()
			if err != nil {
				return
			}
			item.x = x
			item.et = time.Now().Add(kt)
		}()
	})
	return
}

func (tc *Timedcache) AsyncGetf(key string, firstcall func() interface{}, fetchfunc func() (x interface{}, kt time.Duration, err error)) (x interface{}, err error) {
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
		WithMutex(&item.mtx, func() {
			if item.x == nil {
				if firstcall != nil {
					item.x = firstcall()
				} else {
					item.x, _, _ = fetchfunc()
				}
			}
		})
	}
	x = item.x
	if item.fetching {
		return
	}
	WithMutex(&item.mtx, func() {
		if item.et.After(time.Now()) {
			return
		}
		if item.fetching {
			return
		}
		item.fetching = true
		go func() {
			defer func() {
				item.fetching = false
			}()
			var kt time.Duration
			x, kt, err = fetchfunc()
			if err != nil {
				return
			}
			item.x = x
			item.et = time.Now().Add(kt)
		}()
	})
	return
}
