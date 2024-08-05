package properties

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/file/properties/domain/properties"
	"github.com/micro-ginger/file/properties/usecase"
)

type Module struct {
	UseCase properties.UseCase
}

func New(logger log.Logger) *Module {
	uc := usecase.New()
	m := &Module{
		UseCase: uc,
	}
	return m
}
