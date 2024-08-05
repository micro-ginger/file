package app

import (
	"os"
	"os/signal"
	"syscall"
)

func (a *App[acc, f]) Start() {
	go func() {
		if err := a.Ginger.Run(); err != nil {
			panic(err)
		}
	}()

	done := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		close(done)
	}()
	<-done
}
