package app

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func (a *App[acc, f]) Start() {
	go func() {
		if err := a.Ginger.Run(); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := a.Grpc.Run(); err != nil {
			panic(err)
		}
	}()

	done := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		wg := new(sync.WaitGroup)
		// stop
		wg.Add(1)
		go func() {
			a.Logger.WithTrace("exit.ginger").Debugf("stopping...")
			a.Ginger.Shutdown(time.Minute)
			wg.Done()
			a.Logger.WithTrace("exit.ginger").Debugf("stopped")
		}()
		wg.Add(1)
		go func() {
			a.Logger.WithTrace("exit.grpc").Debugf("stopping...")
			a.Grpc.Stop()
			wg.Done()
			a.Logger.WithTrace("exit.grpc").Debugf("stopped")
		}()
		wg.Wait()
		close(done)
	}()
	<-done
}
