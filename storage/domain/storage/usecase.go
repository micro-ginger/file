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
	SaveFile(_ context.Context, request *file.StoreRequest,
		file *f.File[T]) (*SaveResult, errors.Error)

	GetAbsPath(file *f.File[T]) string
	GetRelativeDirPath(key string) string

	GetTempAbsPath(file *f.File[T], extra string) string
	GetTempRelativeDirPath(key string) string
	GetTempFileName(file *f.File[T], extra string) string
}
