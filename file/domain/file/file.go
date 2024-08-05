package file

import (
	f "github.com/micro-blonde/file"
)

type File[T f.Model] struct {
	f.File[T]
}

func (m *File[T]) TableName() string {
	return "files"
}
