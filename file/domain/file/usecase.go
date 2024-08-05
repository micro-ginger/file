package file

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/file"
)

type UseCase[T file.Model] interface {
	Create(ctx context.Context, item *File[T]) errors.Error

	List(ctx context.Context, q query.Query) ([]*File[T], errors.Error)

	Get(ctx context.Context, query query.Query) (*File[T], errors.Error)
	GetById(ctx context.Context, key string) (*File[T], errors.Error)
}
