package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/file/file/domain/file"
)

func (repo *repo[T]) Update(q query.Query, update *file.File[T]) errors.Error {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(file.File[T])
		})
	if err := repo.Repository.Update(q, update); err != nil {
		return err
	}
	return nil
}
