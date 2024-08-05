package domain

import (
	f "github.com/micro-blonde/file"
	"github.com/micro-ginger/file/file/domain/file"
)

type UseCase[T f.Model] interface {
	file.UseCase[T]
}
