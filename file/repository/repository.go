package repository

import (
	"github.com/ginger-core/repository"
	f "github.com/micro-blonde/file"
	"github.com/micro-ginger/file/file/domain"
)

type repo[T f.Model] struct {
	repository.Repository
}

func New[T f.Model](base repository.Repository) domain.Repository[T] {
	repo := &repo[T]{
		Repository: base,
	}
	return repo
}
