package properties

import (
	"github.com/ginger-core/log"
	"github.com/micro-ginger/file/properties/delivery/grpc"
	pd "github.com/micro-ginger/file/properties/domain/delivery/properties"
	"github.com/micro-ginger/file/properties/domain/properties"
	"github.com/micro-ginger/file/properties/usecase"
)

type Module struct {
	UseCase properties.UseCase

	GrpcPropertiesHandler pd.GrpcPropertiesGetter
}

func New(logger log.Logger) *Module {
	uc := usecase.New()
	m := &Module{
		UseCase: uc,
		GrpcPropertiesHandler: grpc.NewGet(
			logger.WithTrace("grpc.get"),
			uc,
		),
	}
	return m
}
