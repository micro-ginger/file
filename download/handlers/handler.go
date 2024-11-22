package handlers

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/file/download/domain/storage"
	f "github.com/micro-ginger/file/file/domain/file"
)

type Handler[T file.Model] interface {
	Initialize(storage storage.UseCase[T])

	Resize(ctx context.Context, file *f.File[T],
		request Request) (string, errors.Error)
}
