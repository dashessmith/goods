package goods

import (
	"sync"
)

func WithMutex(mtx *sync.Mutex, f func()) {
	if mtx != nil {
		mtx.Lock()
		defer mtx.Unlock()
	}
	f()
}

func WithLock(mtx *sync.RWMutex, f func()) {
	if mtx != nil {
		mtx.Lock()
		defer mtx.Unlock()
	}
	f()
}

func WithRLock(mtx *sync.RWMutex, f func()) {
	if mtx != nil {
		mtx.RLock()
		defer mtx.RUnlock()
	}
	f()
}

var (
	mutexesGuard sync.RWMutex
	mutexes      = map[string]*sync.RWMutex{}
)

func getmtx(tag string) (mtx *sync.RWMutex) {
	mutexesGuard.RLock()
	mtx = mutexes[tag]
	mutexesGuard.RUnlock()
	if mtx != nil {
		return mtx
	}
	mutexesGuard.Lock()
	defer mutexesGuard.Unlock()
	mtx = mutexes[tag]
	if mtx == nil {
		mtx = &sync.RWMutex{}
		mutexes[tag] = mtx
	}
	return
}

func LockTag(tag string) *sync.RWMutex {
	mtx := getmtx(tag)
	mtx.Lock()
	return mtx
}

func WithLockTag(tag string, f func()) {
	mtx := getmtx(tag)
	mtx.Lock()
	defer mtx.Unlock()
	f()
}

func RLockTag(tag string) *sync.RWMutex {
	mtx := getmtx(tag)
	mtx.RLock()
	return mtx
}

func WithRLockTag(tag string, f func()) {
	mtx := getmtx(tag)
	mtx.RLock()
	defer mtx.RUnlock()
	f()
}
