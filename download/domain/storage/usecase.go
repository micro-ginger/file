package storage

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/file/domain/file"
	"github.com/micro-ginger/file/storage/domain/storage"
)

type UseCase[T file.Model] interface {
	SaveFile(_ context.Context, request *file.StoreRequest,
		file *f.File[T]) (*storage.SaveResult, errors.Error)

	GetAbsPath(file *f.File[T]) string
	GetTempAbsPath(file *f.File[T], extra string) string
	GetRelativeDirPath(key string) string
	GetTempRelativeDirPath(key string) string
	GetTempFileName(file *f.File[T], extra string) string
}
