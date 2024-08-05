package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/file/file/domain/file"
)

func (uc *useCase[T]) List(ctx context.Context,
	query query.Query) ([]*file.File[T], errors.Error) {
	return uc.repo.List(query)
}

func (uc *useCase[T]) Get(ctx context.Context,
	query query.Query) (*file.File[T], errors.Error) {
	return uc.repo.Get(query)
}

func (uc *useCase[T]) GetById(ctx context.Context,
	key string) (*file.File[T], errors.Error) {
	q := query.New(ctx)
	return uc.getById(ctx, q, key)
}

func (uc *useCase[T]) getById(ctx context.Context,
	q query.Query, key string) (*file.File[T], errors.Error) {
	q = query.NewFilter(q).
		WithOr(query.NewFilter(nil).
			WithMatch(&query.Match{
				Key:      "`id`",
				Operator: query.Equal,
				Value:    key,
			}),
		)

	return uc.get(ctx, q)
}

func (uc *useCase[T]) get(ctx context.Context,
	q query.Query) (*file.File[T], errors.Error) {
	return uc.repo.Get(q)
}
