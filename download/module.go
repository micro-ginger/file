package download

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/file/download/delivery"
	"github.com/micro-ginger/file/download/domain"
	f "github.com/micro-ginger/file/download/domain/file"
	"github.com/micro-ginger/file/download/domain/storage"
	"github.com/micro-ginger/file/download/usecase"
	"github.com/micro-ginger/file/properties/domain/properties"
)

type Module[T file.Model] struct {
	UseCase         domain.UseCase
	DownloadHandler delivery.DownloadHandler[T]
}

func New[T file.Model](logger log.Logger, registry registry.Registry,
	properties properties.UseCase, responder gateway.Responder) *Module[T] {
	uc := usecase.New(logger.WithTrace("uc"), registry, properties)
	m := &Module[T]{
		UseCase: uc,
		DownloadHandler: delivery.NewDownload[T](
			logger.WithTrace("delivery.download"),
			responder,
		),
	}
	return m
}

func (m *Module[T]) Initialize(file f.UseCase[T], storage storage.UseCase[T]) {
	m.DownloadHandler.Initialize(file, storage)
}
