package upload

import (
	"context"

	"github.com/micro-blonde/file/proto/file"
)

type GrpcStoreHandler interface {
	Store(context.Context, *file.StoreRequest) (*file.StoreResponse, error)
}
