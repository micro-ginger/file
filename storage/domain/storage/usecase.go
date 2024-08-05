package storage

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/file/domain/file"
)

type UseCase[T file.Model] interface {
	Store(ctx context.Context,
		request *file.StoreRequest) (*file.StoreResponse[T], errors.Error)

	GetAbsPath(file *f.File[T]) string
}
