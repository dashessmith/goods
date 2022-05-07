package goods

import (
	"context"
	"fmt"
	"reflect"
)

func SafeClose(x interface{}) {
	defer func() {
		_ = recover()
	}()
	rv := reflect.ValueOf(x)
	if rv.Kind() == reflect.Chan {
		rv.Close()
	}
}

func ChanClose[T any](ch chan T) {
	defer func() {
		_ = recover()
	}()
	close(ch)
}

func ChanRead[T any](ctx context.Context, ch chan T) (ret T, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	var ok bool
	select {
	case ret, ok = <-ch:
		if !ok {
			err = fmt.Errorf("chan closed")
		}
	case <-ctx.Done():
		err = ctx.Err()
	}
	return
}

func ChanWrite[T any](ctx context.Context, ch chan T, x T) (err error) {
	defer func() {
		if r := recover(); r != nil && err == nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	if ctx == nil {
		ctx = context.Background()
	}
	select {
	case ch <- x:
	case <-ctx.Done():
		err = ctx.Err()
	}
	return
}
