package grpc

import (
	"context"

	"github.com/ginger-core/log"
	"github.com/micro-blonde/file/proto/file"
	pd "github.com/micro-ginger/file/properties/domain/delivery/properties"
	"github.com/micro-ginger/file/properties/domain/properties"
)

type propertiesHandler struct {
	logger log.Logger
	uc     properties.UseCase
}

func NewGet(logger log.Logger) pd.GrpcPropertiesGetter {
	h := &propertiesHandler{
		logger: logger,
	}
	return h
}

func (h *propertiesHandler) GetProperties(context.Context,
	*file.PropertiesRequest) (*file.Properties, error) {
	props := h.uc.GetProperties()

	resp := &file.Properties{
		KeyBaseUrl: props.KeyBaseUrl,
	}

	return resp, nil
}
