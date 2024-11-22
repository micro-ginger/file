package delivery

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/file"
	d "github.com/micro-ginger/file/download/domain/download"
	f "github.com/micro-ginger/file/download/domain/file"
	"github.com/micro-ginger/file/download/domain/storage"
	"github.com/micro-ginger/file/download/handlers"
)

type DownloadHandler[T file.Model] interface {
	gateway.Handler
	Initialize(file f.UseCase[T], storage storage.UseCase[T])
}

type download[T file.Model] struct {
	gateway.Responder
	logger log.Logger
	config config

	uc             d.UseCase
	file           f.UseCase[T]
	storage        storage.UseCase[T]
	resizeHandlers map[string]handlers.Handler[T]
}

func NewDownload[T file.Model](logger log.Logger, registry registry.Registry,
	uc d.UseCase, responder gateway.Responder) DownloadHandler[T] {
	h := &download[T]{
		logger:         logger,
		uc:             uc,
		Responder:      responder,
		resizeHandlers: make(map[string]handlers.Handler[T]),
	}
	if err := registry.Unmarshal(&h.config); err != nil {
		panic(err)
	}
	imageHandler := handlers.NewImage[T](registry.ValueOf("handlers.image"))
	for _, t := range h.config.ResizableImageTypes {
		h.resizeHandlers[t] = imageHandler
	}
	return h
}

func (h *download[T]) Initialize(file f.UseCase[T], storage storage.UseCase[T]) {
	h.file = file
	h.storage = storage
	for _, h := range h.resizeHandlers {
		h.Initialize(storage)
	}
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

	maxQ, ok := request.GetQuery("max")
	if ok {
		max, pErr := strconv.ParseUint(maxQ, 10, 64)
		if pErr != nil {
			return nil, errors.Validation(pErr)
		}
		if max == 0 {
			return nil, errors.Validation().
				WithTrace("max=0")
		}
		handler, ok := h.resizeHandlers[f.Type]
		if ok {
			newPath, err := handler.Resize(ctx, f,
				handlers.NewRequest().
					WithMaxSize(max),
			)
			if err != nil {
				return nil, err.WithTrace("handler.Resize")
			}
			if newPath != "" {
				path = newPath
			}
		}
	}

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
