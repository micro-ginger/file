package file

import (
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	f "github.com/micro-blonde/file"
	"github.com/micro-ginger/file/file/domain"
	repo "github.com/micro-ginger/file/file/repository"
	"github.com/micro-ginger/file/file/usecase"
)

type Module[T f.Model] struct {
	Repository domain.Repository[T]
	UseCase    domain.UseCase[T]
}

func New[T f.Model](logger log.Logger, base repository.Repository) *Module[T] {
	repo := repo.New[T](base)
	uc := usecase.New(logger.WithTrace("uc"), repo)
	m := &Module[T]{
		Repository: repo,
		UseCase:    uc,
	}
	return m
}
