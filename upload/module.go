package upload

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/file/upload/delivery"
	"github.com/micro-ginger/file/upload/domain/download"
	"github.com/micro-ginger/file/upload/domain/storage"
)

type Module[T file.Model] struct {
	UploadHandler delivery.UploadHandler[T]
}

func New[T file.Model](logger log.Logger, registry registry.Registry,
	responder gateway.Responder) *Module[T] {
	m := &Module[T]{
		UploadHandler: delivery.NewUpload[T](
			logger.WithTrace("handler"),
			registry, responder,
		),
	}
	return m
}

func (m *Module[T]) Initialize(storage storage.UseCase[T],
	download download.UseCase) {
	m.UploadHandler.Initialize(storage, download)
}
