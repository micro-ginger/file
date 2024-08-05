package domain

import (
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/file/storage/domain/download"
	sf "github.com/micro-ginger/file/storage/domain/file"
	"github.com/micro-ginger/file/storage/domain/storage"
)

type UseCase[T file.Model] interface {
	Initialize(file sf.UseCase[T], download download.UseCase)
	storage.UseCase[T]
}
