package usecase

import (
	"github.com/ginger-core/log"
	f "github.com/micro-blonde/file"
	"github.com/micro-ginger/file/file/domain"
)

type useCase[T f.Model] struct {
	logger log.Logger

	repo domain.Repository[T]
}

func New[T f.Model](logger log.Logger, repo domain.Repository[T]) domain.UseCase[T] {
	uc := &useCase[T]{
		logger: logger,
		repo:   repo,
	}
	return uc
}
