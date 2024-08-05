package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/file/file/domain/file"
)

func (uc *useCase[T]) Create(ctx context.Context, item *file.File[T]) errors.Error {
	q := query.New(ctx)
	if err := uc.repo.Create(q, item); err != nil {
		return err
	}
	return nil
}
