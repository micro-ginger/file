package file

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/file/domain/file"
)

type UseCase[T file.Model] interface {
	GetById(ctx context.Context, id string) (*f.File[T], errors.Error)
}
