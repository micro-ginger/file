package app

import (
	"context"

	"github.com/ginger-core/compound/registry"
)

type config struct {
	Gateway struct {
		Language struct {
			DefaultLanguage string
			Dir             string
			Languages       []string
		}
	}
	Upload struct {
		Enabled bool
	}
}

func (a *App[acc, f]) loadConfig() {
	registry, err := registry.New(context.Background())
	if err != nil {
		panic(err)
	}
	a.Registry = registry
}
