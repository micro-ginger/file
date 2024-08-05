package app

import "github.com/ginger-core/log"

func (a *App[acc, f]) initializeLogger() {
	a.Logger = log.NewLogger(a.Registry.ValueOf("logger"))
	a.Logger.SetSource("file")
	a.Logger.Start()
}
