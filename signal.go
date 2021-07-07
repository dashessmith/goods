package util

import (
	"os"
	"os/signal"
	"syscall"
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
