package usecase

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/file/storage/domain"
	"github.com/micro-ginger/file/storage/domain/download"
	f "github.com/micro-ginger/file/storage/domain/file"
)

type useCase[T file.Model] struct {
	logger log.Logger
	config config

	file     f.UseCase[T]
	download download.UseCase
}

func New[T file.Model](logger log.Logger, registry registry.Registry) domain.UseCase[T] {
	uc := &useCase[T]{
		logger: logger,
	}

	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	return uc
}

func (uc *useCase[T]) Initialize(file f.UseCase[T], download download.UseCase) {
	uc.file = file
	uc.download = download
}
