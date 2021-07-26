package goods

import "time"

func Elapse(f func()) (d time.Duration) {
	t := time.Now()
	defer func() {
		d = time.Since(t)
	}()
	f()
	return
}
