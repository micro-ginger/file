package properties

import (
	"context"

	"github.com/micro-blonde/file/proto/file"
)

type GrpcPropertiesGetter interface {
	GetProperties(context.Context,
		*file.PropertiesRequest) (*file.Properties, error)
}
