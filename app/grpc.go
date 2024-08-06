package app

import (
	"net"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/log"
	"github.com/micro-blonde/file/proto/file"
	"github.com/micro-ginger/file/properties/domain/delivery/properties"
	"github.com/micro-ginger/file/upload/domain/delivery/upload"
	"google.golang.org/grpc"
)

type GrpcServer interface {
	file.ServiceServer
	Run() errors.Error
	Stop()
}

type grpcServer struct {
	file.UnsafeServiceServer
	properties.GrpcPropertiesGetter
	upload.GrpcStoreHandler
	logger log.Logger
	config struct {
		ListenAddr string
	}
	// server
	gRpcServer *grpc.Server
}

func (a *App[acc, f]) newGrpc(registry registry.Registry) GrpcServer {
	s := &grpcServer{
		logger:               a.Logger.WithTrace("grpc"),
		GrpcPropertiesGetter: a.Properties.GrpcPropertiesHandler,
		GrpcStoreHandler:     a.Upload.GrpcStoreHandler,
	}
	if err := registry.Unmarshal(&s.config); err != nil {
		panic(err)
	}
	return s
}

func (s *grpcServer) Run() errors.Error {
	var sererOptions []grpc.ServerOption
	s.gRpcServer = grpc.NewServer(sererOptions...)
	file.RegisterServiceServer(s.gRpcServer, s)
	l, err := net.Listen("tcp", s.config.ListenAddr)
	if err != nil {
		return errors.New(err)
	}
	s.logger.Infof("grpc server listening to %s", s.config.ListenAddr)
	if err = s.gRpcServer.Serve(l); err != nil {
		return errors.New(err)
	}
	return nil
}

func (s *grpcServer) Stop() {
	s.gRpcServer.Stop()
}
