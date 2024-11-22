package usecase

import (
	"path"

	f "github.com/micro-ginger/file/file/domain/file"
)

func (uc *useCase[T]) GetAbsDirPath(baseDir string) string {
	return path.Join(uc.config.DirPath, baseDir)
}

func (uc *useCase[T]) GetTempBaseDir() string {
	return uc.config.TempBaseDirPath
}

func (uc *useCase[T]) GetAbsPath(file *f.File[T]) string {
	return path.Join(uc.config.DirPath, file.Path)
}

func (uc *useCase[T]) GetTempAbsPath(file *f.File[T], extra string) string {
	return path.Join(
		uc.config.DirPath,
		uc.GetTempRelativePath(file, extra),
	)
}

func (uc *useCase[T]) GetTempFileName(file *f.File[T], extra string) string {
	return file.Key + "_" + extra
}

func (uc *useCase[T]) GetTempRelativePath(file *f.File[T], extra string) string {
	return path.Join(
		uc.GetTempRelativeDirPath(file.Key),
		uc.GetTempFileName(file, extra),
	)
}

func (uc *useCase[T]) GetTempRelativeDirPath(key string) string {
	return path.Join(
		uc.config.TempBaseDirPath,
		uc.GetRelativeDirPath(key),
	)
}

func (uc *useCase[T]) GetRelativeDirPath(key string) string {
	return path.Join(key[:2], key[2:4])
}
