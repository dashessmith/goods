package goods

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, d time.Duration) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if d <= 0 {
		return ctx.Err()
	}

	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
