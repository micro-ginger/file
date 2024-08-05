package domain

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	f "github.com/micro-blonde/file"
	"github.com/micro-ginger/file/file/domain/file"
)

type Repository[T f.Model] interface {
	Create(q query.Query, item *file.File[T]) errors.Error
	List(q query.Query) ([]*file.File[T], errors.Error)
	Get(q query.Query) (*file.File[T], errors.Error)
	Update(q query.Query, update *file.File[T]) errors.Error
}
