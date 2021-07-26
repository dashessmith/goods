package goods

import "sync"

func WithMutex(mtx *sync.Mutex, f func()) {
	if mtx != nil {
		mtx.Lock()
		defer mtx.Unlock()
	}
	f()
}
