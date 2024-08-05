package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/file"
	f "github.com/micro-ginger/file/download/domain/file"
	"github.com/micro-ginger/file/download/domain/storage"
)

type DownloadHandler[T file.Model] interface {
	gateway.Handler
	Initialize(file f.UseCase[T], storage storage.UseCase[T])
}

type download[T file.Model] struct {
	gateway.Responder
	logger log.Logger

	file    f.UseCase[T]
	storage storage.UseCase[T]
}

func NewDownload[T file.Model](logger log.Logger,
	responder gateway.Responder) DownloadHandler[T] {
	h := &download[T]{
		Responder: responder,
		logger:    logger,
	}
	return h
}

func (h *download[T]) Initialize(file f.UseCase[T], storage storage.UseCase[T]) {
	h.file = file
	h.storage = storage
}

func (h *download[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	// auth := request.GetAuthorization().(authorization.Authorization)
	// accountId := auth.GetApplicantId().(uint64)

	id := request.GetParam("id")

	f, err := h.file.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, errors.NotFound()
	}

	path := h.storage.GetAbsPath(f)

	// file, fErr := os.Open(path)
	// if fErr != nil {
	// 	if os.IsNotExist(fErr) {
	// 		return nil, errors.NotFound(err)
	// 	}
	// }
	ginCtx := request.GetConn().(*gin.Context)
	// ginCtx.Header("Content-Description", "File Transfer")
	// ginCtx.Header("Content-Transfer-Encoding", "binary")
	ginCtx.Header("Content-Disposition", "attachment; filename="+f.Name)
	// ginCtx.Header("Content-Type", "application/octet-stream")
	ginCtx.File(path)

	request.SetResponded()
	return nil, nil
}
