package app

import (
	"context"
	"os"

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

func (a *App[acc, f]) loadConfig(configType string) {
	switch configType {
	case "FILE", "":
		filePath := os.Getenv("CONFIG_PATH")
		if filePath == "" {
			filePath = "./config.yaml"
		}
		registry, err := registry.New(context.Background(),
			registry.TypeFile, "yaml", filePath)
		if err != nil {
			panic(err)
		}
		a.Registry = registry
	case "GIT":
		baseUrl := os.Getenv("CONFIG_BASE_URL")
		if baseUrl == "" {
			panic("invalid config remote base url")
		}
		filePath := os.Getenv("CONFIG_PATH")
		if filePath == "" {
			panic("invalid config remote path")
		}
		ref := os.Getenv("CONFIG_REF")
		if ref == "" {
			ref = "master"
		}
		token := os.Getenv("CONFIG_TOKEN")
		if filePath == "" {
			panic("invalid config remote token")
		}
		registry, err := registry.New(context.Background(),
			registry.TypeGitAPI, "yaml", baseUrl, filePath, ref, token)
		if err != nil {
			panic(err)
		}
		a.Registry = registry
	default:
		panic("invalid config type")
	}
}
