package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ginger-core/errors"
	"github.com/google/uuid"
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/file/domain/file"
	"github.com/micro-ginger/file/storage/domain/storage"
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

	sr, err := uc.SaveFile(ctx, request, f)
	if err != nil {
		return nil, err.WithTrace("Store.saveFile")
	}
	resp.AbsPath = sr.AbsPath

	if err := uc.file.Create(ctx, f); err != nil {
		return nil, err.WithTrace("Store.file.Create")
	}

	resp.Url = uc.download.GetAbsUrlById(f.Id)
	return resp, nil
}

func (uc *useCase[T]) SaveFile(_ context.Context,
	request *file.StoreRequest, file *f.File[T]) (*storage.SaveResult, errors.Error) {
	// make directory ready
	var baseDir string
	if request.BaseDir != nil {
		baseDir = *request.BaseDir
	}

	if baseDir == "" {
		baseDir = uc.GetRelativeDirPath(file.Key)
	}

	fullDirPath := uc.GetAbsDirPath(baseDir)
	if err := os.MkdirAll(fullDirPath, os.ModePerm); err != nil {
		return nil, errors.New(err).WithTrace("saveFile.MkdirAll")
	}

	file.Path = path.Join(baseDir, file.Key)
	r := &storage.SaveResult{
		AbsPath: uc.GetAbsPath(file),
	}

	fi, err := os.Stat(r.AbsPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, errors.New(err).
			WithTrace("os.Stat")
	}
	if fi != nil {
		r.IsExists = true
		return r, nil
	}

	// save file
	fo, err := os.Create(r.AbsPath)
	if err != nil {
		return nil, errors.New(err).WithTrace("saveFile.os.Create")
	}
	defer fo.Close()
	if request.Reader != nil {
		_, cErr := io.Copy(fo, request.Reader)
		if cErr != nil {
			return nil, errors.New(cErr).
				WithTrace("io.Copy")
		}
	} else {
		// write data
		if _, err := fo.Write(request.Data); err != nil {
			return nil, errors.New(err).WithTrace("saveFile.fo.Write")
		}
	}
	return r, nil
}
