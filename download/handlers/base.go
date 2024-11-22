package handlers

import (
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/file/download/domain/storage"
)

type base[T file.Model] struct {
	storage storage.UseCase[T]
}

func newBase[T file.Model]() *base[T] {
	return &base[T]{}
}

func (h *base[T]) Initialize(storage storage.UseCase[T]) {
	h.storage = storage
}
