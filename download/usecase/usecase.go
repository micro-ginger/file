package usecase

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/file/download/domain"
	"github.com/micro-ginger/file/properties/domain/properties"
)

type useCase struct {
	logger log.Logger
	config config
}

func New(logger log.Logger, registry registry.Registry,
	properties properties.UseCase) domain.UseCase {
	uc := &useCase{
		logger: logger,
	}
	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()

	properties.SetKeyBaseUrl(uc.getUrlIdPrefix())
	return uc
}
