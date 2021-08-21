package goods

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WithSignalInterface interface {
	Run() (err error)
	Stop() (err error)
}

func WithSignal(x WithSignalInterface) (err error) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		err = x.Stop()
	}()
	err = x.Run()
	return
}

func UniqueRunWithSignal(svc WithSignalInterface, tag string) (err error) {
	WithFlock(tag, func() {
		err = WithSignal(svc)
	})
	return
}

func UniqueRunWithSignalContext(ctx context.Context, delay time.Duration, svc WithSignalInterface, tag string) (err error) {
	WithFlockContext(ctx, delay, tag, func() {
		err = WithSignal(svc)
	})
	return
}
