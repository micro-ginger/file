package usecase

import (
	"path"

	f "github.com/micro-ginger/file/file/domain/file"
)

func (uc *useCase[T]) GetAbsDirPath(baseDir string) string {
	return path.Join(uc.config.DirPath, baseDir)
}

func (uc *useCase[T]) GetAbsPath(file *f.File[T]) string {
	return path.Join(uc.config.DirPath, file.Path)
}
