package file

import (
	"github.com/micro-blonde/auth/account"
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/file/app"
)

func NewApp[acc account.Model, f file.Model](configType string) app.Application {
	return app.New[acc, f](configType)
}
