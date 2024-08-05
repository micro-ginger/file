package app

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func (a *App[acc, f]) initializeLanguage() {
	a.Language.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, lang := range a.Config.Gateway.Language.Languages {
		a.Language.MustLoadMessageFile(
			filepath.Join(a.Config.Gateway.Language.Dir, lang))
	}
}
