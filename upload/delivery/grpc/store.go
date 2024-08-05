package grpc

import (
	"context"
	"time"

	errorsgrpc "github.com/ginger-core/errors/grpc"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	f "github.com/micro-blonde/file"
	"github.com/micro-blonde/file/proto/file"
	pf "github.com/micro-ginger/file/file/domain/delivery/file"
	"github.com/micro-ginger/file/upload/domain/delivery/upload"
	"github.com/micro-ginger/file/upload/domain/download"
	"github.com/micro-ginger/file/upload/domain/storage"
)

type StoreHandler[T f.Model] interface {
	upload.GrpcStoreHandler
	Initialize(storage storage.UseCase[T], download download.UseCase)
}

type store[T f.Model] struct {
	logger   log.Logger
	storage  storage.UseCase[T]
	download download.UseCase
}

func NewStore[T f.Model](logger log.Logger) StoreHandler[T] {
	h := &store[T]{
		logger: logger,
	}
	return h
}

func (h *store[T]) Initialize(storage storage.UseCase[T],
	download download.UseCase) {
	h.storage = storage
	h.download = download
}

func (h *store[T]) Store(ctx context.Context,
	request *file.StoreRequest) (*file.StoreResponse, error) {
	req := &f.StoreRequest{
		Type: request.Type,
		Data: request.Data,
		Name: request.Name,
	}
	if request.BaseDir != "" {
		req.BaseDir = &request.BaseDir
	}
	if request.ExpiresInSecs != 0 {
		expIn := time.Duration(request.ExpiresInSecs) * time.Second
		req.ExpiresIn = &expIn
	}
	if request.AccountId != 0 {
		req.AccountId = &request.AccountId
	}

	result, err := h.storage.Store(ctx, req)
	if err == nil {
		h.logger.
			With(logger.Field{
				"request.type": request.Type,
			}).
			Infof("store request")
	} else {
		h.logger.
			With(logger.Field{
				"request.type": request.Type,
				"error":        err.Error(),
			}).
			Errorf("store request")
		return nil, errorsgrpc.Generate(err)
	}

	resp := &file.StoreResponse{
		File: pf.GetFileProto(result.File),
		Url:  result.Url,
	}

	return resp, err
}
