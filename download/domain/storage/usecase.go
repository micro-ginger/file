package storage

import (
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/file/domain/file"
)

type UseCase[T file.Model] interface {
	GetAbsPath(file *f.File[T]) string
}
