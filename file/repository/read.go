package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/file/file/domain/file"
)

func (repo *repo[T]) List(q query.Query) ([]*file.File[T], errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]*file.File[T])
		})
	r, err := repo.Repository.List(q)
	if err != nil {
		return nil, err
	}
	return *r.(*[]*file.File[T]), nil
}

func (repo *repo[T]) Get(q query.Query) (*file.File[T], errors.Error) {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(file.File[T])
		})
	r, err := repo.Repository.Get(q)
	if err != nil {
		return nil, err
	}
	return r.(*file.File[T]), nil
}
