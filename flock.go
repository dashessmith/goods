package goods

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofrs/flock"
)

func WithFlock(path string, f func()) {
	if len(path) <= 0 {
		path = BinNameExt + ".lock"
	}
	lock := flock.New(path)
	ok, err := lock.TryLock()
	if !ok || err != nil {
		log.Printf("%v is held by another process\n", path)
		return
	}
	defer func() {
		lock.Close()
		os.Remove(path)
	}()
	f()
}

func WithFlockContext(ctx context.Context, delay time.Duration, path string, f func()) {
	if len(path) <= 0 {
		path = BinNameExt + ".lock"
	}
	lock := flock.New(path)
	ok, err := lock.TryLockContext(ctx, delay)
	if !ok || err != nil {
		log.Printf("%v is held by another process\n", path)
		return
	}
	defer func() {
		lock.Close()
		os.Remove(path)
	}()
	f()
}
