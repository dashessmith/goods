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

type timedcache struct {
	data sync.Map
	once sync.Once
}

func NewTimedcache() *timedcache {
	tc := &timedcache{}

	tc.once.Do(func() { go tc.cleanExpired() })

	return tc
}

func (tc *timedcache) Get(key string, fetchfunc func() (x interface{}, kt time.Duration, err error)) (x interface{}, err error) {
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

func (tc *timedcache) AsyncGet(key string, asyncDefault interface{}, fetchfunc func() (x interface{}, kt time.Duration, err error)) (x interface{}, err error) {
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

func (tc *timedcache) AsyncGetf(key string, firstcall func() interface{}, fetchfunc func() (x interface{}, kt time.Duration, err error)) (x interface{}, err error) {
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

func (tc *timedcache) cleanExpired() {
	for ; ; time.Sleep(time.Minute) {
		tc.data.Range(func(key, value any) bool {
			if item := value.(*timedcacheitem); item.et.Before(time.Now()) {
				tc.Delete(key.(string))
			}
			return true
		})
	}
}

func (tc *timedcache) Delete(key string) {
	tc.data.Delete(key)
}
