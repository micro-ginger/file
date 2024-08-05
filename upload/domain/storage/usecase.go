package storage

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/file"
)

type UseCase[T file.Model] interface {
	Store(ctx context.Context,
		request *file.StoreRequest) (*file.StoreResponse[T], errors.Error)
}
