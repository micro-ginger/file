package file

import (
	"context"

	"github.com/ginger-core/errors"
	f "github.com/micro-blonde/file"
	"github.com/micro-ginger/file/file/domain/file"
)

type UseCase[T f.Model] interface {
	GetById(ctx context.Context, id string) (*file.File[T], errors.Error)
	Create(ctx context.Context, item *file.File[T]) errors.Error
}
