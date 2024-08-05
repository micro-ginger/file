package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ginger-core/errors"
	"github.com/google/uuid"
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/file/domain/file"
)

func (uc *useCase[T]) Store(ctx context.Context,
	request *file.StoreRequest) (*file.StoreResponse[T], errors.Error) {
	hash := md5.Sum(request.Data)
	f := &f.File[T]{
		File: file.File[T]{
			Id:        strings.Replace(uuid.NewString(), "-", "", 4),
			Key:       hex.EncodeToString(hash[:]),
			AccountId: request.AccountId,
			Name:      request.Name,
			Type:      request.Type,
			// Meta:           &file.Meta{},
		},
	}

	if request.ExpiresIn != nil {
		expAt := time.Now().UTC().Add(*request.ExpiresIn)
		f.ExpirationTime = &expAt
	}

	resp := &file.StoreResponse[T]{
		File: &f.File,
	}

	sr, err := uc.saveFile(ctx, request, f)
	if err != nil {
		return nil, err.WithTrace("Store.saveFile")
	}
	resp.AbsPath = sr.absPath

	if err := uc.file.Create(ctx, f); err != nil {
		return nil, err.WithTrace("Store.file.Create")
	}

	resp.Url = uc.download.GetAbsUrlById(f.Id)
	return resp, nil
}

type saveResult struct {
	isExists bool
	absPath  string
}

func (uc *useCase[T]) saveFile(_ context.Context,
	request *file.StoreRequest, file *f.File[T]) (*saveResult, errors.Error) {
	// make directory ready
	var baseDir string
	if request.BaseDir != nil {
		baseDir = *request.BaseDir
	}

	if baseDir == "" {
		baseDir = path.Join(file.Key[:2], file.Key[2:4])
	}

	fullDirPath := uc.GetAbsDirPath(baseDir)
	if err := os.MkdirAll(fullDirPath, os.ModePerm); err != nil {
		return nil, errors.New(err).WithTrace("saveFile.MkdirAll")
	}

	file.Path = path.Join(baseDir, file.Key)
	r := &saveResult{
		absPath: uc.GetAbsPath(file),
	}

	fi, err := os.Stat(r.absPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, errors.New(err).
			WithTrace("os.Stat")
	}
	if fi != nil {
		r.isExists = true
		return r, nil
	}

	// save file
	fo, err := os.Create(r.absPath)
	if err != nil {
		return nil, errors.New(err).WithTrace("saveFile.os.Create")
	}
	defer fo.Close()
	// write data
	if _, err := fo.Write(request.Data); err != nil {
		return nil, errors.New(err).WithTrace("saveFile.fo.Write")
	}
	return r, nil
}
