package storage

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/file/domain/file"
	"github.com/micro-ginger/file/storage/domain"
	"github.com/micro-ginger/file/storage/domain/download"
	"github.com/micro-ginger/file/storage/usecase"
)

type Module[T file.Model] struct {
	UseCase domain.UseCase[T]
}

func New[T file.Model](logger log.Logger, registry registry.Registry) *Module[T] {
	uc := usecase.New[T](logger.WithTrace("uc"), registry)
	m := &Module[T]{
		UseCase: uc,
	}
	return m
}

func (m *Module[T]) Initialize(file f.UseCase[T], download download.UseCase) {
	m.UseCase.Initialize(file, download)
}
