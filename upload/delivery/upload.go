package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/h2non/filetype"
	f "github.com/micro-blonde/file"
	"github.com/micro-ginger/file/upload/domain/delivery"
	"github.com/micro-ginger/file/upload/domain/download"
	"github.com/micro-ginger/file/upload/domain/storage"
)

type UploadHandler[T f.Model] interface {
	gateway.Handler
	Initialize(storage storage.UseCase[T], download download.UseCase)
}

type upload[T f.Model] struct {
	gateway.Responder
	logger log.Logger

	config config

	storage  storage.UseCase[T]
	download download.UseCase
}

func NewUpload[T f.Model](logger log.Logger, registry registry.Registry,
	responder gateway.Responder) UploadHandler[T] {
	h := &upload[T]{
		Responder: responder,
		logger:    logger,
	}
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	h.config.initialize()
	return h
}

func (h *upload[T]) Initialize(storage storage.UseCase[T],
	download download.UseCase) {
	h.storage = storage
	h.download = download
}

func (h *upload[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	ginCtx := request.GetConn().(*gin.Context)

	file, fErr := ginCtx.FormFile("file")
	if fErr != nil {
		if fErr == http.ErrMissingFile {
			return nil, errors.Validation(fErr).
				WithId("FileRequiredError").
				WithMessage("file is missing to process your request")
		}
		return nil, errors.Validation(fErr)
	}
	uploadedFile, oErr := file.Open()
	if oErr != nil {
		return nil, errors.Validation(oErr).
			WithTrace("file.Open")
	}
	data := make([]byte, file.Size)
	_, rErr := uploadedFile.Read(data)
	if rErr != nil {
		return nil, errors.New(rErr).
			WithTrace("uploadedFile.Read")
	}
	// detect file
	kind, mErr := filetype.Match(data)
	if mErr != nil {
		return nil, errors.New(mErr).
			WithTrace("filetype.Match")
	}
	req := &f.StoreRequest{
		Data: data,
		Name: file.Filename,
		Type: kind.MIME.Type,
	}

	f, err := h.storage.Store(ctx, req)
	if err != nil {
		return nil, err.
			WithTrace("storage.Store")
	}
	if f == nil {
		return nil, errors.NotFound().
			WithTrace("f=nil")
	}
	return &delivery.UploadResponse{
		Url: h.download.GetAbsUrlById(f.Id),
	}, nil
}
