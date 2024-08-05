package app

import (
	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-repository/sql"
	"github.com/micro-blonde/auth/account"
	"github.com/micro-blonde/auth/authorization"
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/file/download"
	fm "github.com/micro-ginger/file/file"
	"github.com/micro-ginger/file/properties"
	"github.com/micro-ginger/file/storage"
	"github.com/micro-ginger/file/upload"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Application interface {
	Initialize()
	Start()
}

type App[acc account.Model, f file.Model] struct {
	Registry registry.Registry
	Config   config
	Logger   log.Handler
	Language *i18n.Bundle
	/* database */
	Sql sql.Repository
	/* services */
	/* modules */
	Properties *properties.Module
	File       *fm.Module[f]
	Storage    *storage.Module[f]
	Upload     *upload.Module[f]
	Download   *download.Module[f]
	/* server */
	Authenticator authorization.Authenticator[acc]
	Ginger        gateway.Server
	Grpc          GrpcServer
}

func New[acc account.Model, f file.Model](configType string) *App[acc, f] {
	a := &App[acc, f]{
		Language: i18n.NewBundle(language.English),
	}
	a.loadConfig(configType)

	if err := a.Registry.Unmarshal(&a.Config); err != nil {
		panic(err)
	}
	return a
}

func (a *App[acc, f]) Initialize() {
	a.initializeLogger()
	a.initializeLanguage()
	a.initializeServer()
	a.initializeServices()
	a.initializeDatabases()
	a.initializeModules()
	a.initializeGrpc()
	a.registerRoutes()
}
