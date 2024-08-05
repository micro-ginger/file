package app

import (
	"github.com/micro-ginger/file/download"
	"github.com/micro-ginger/file/file"
	"github.com/micro-ginger/file/properties"
	"github.com/micro-ginger/file/storage"
	"github.com/micro-ginger/file/upload"
)

func (a *App[acc, f]) initializeModules() {
	a.initializeProperties()
	a.initializeFile()
	a.initializeStore()
	if a.Config.Upload.Enabled {
		a.initializeUpload()
	}
	a.initializeDownload()
	// initialize
	a.Storage.Initialize(a.File.UseCase, a.Download.UseCase)
	a.Download.Initialize(a.File.UseCase, a.Storage.UseCase)
}

func (a *App[acc, f]) initializeProperties() {
	a.Properties = properties.New(
		a.Logger.WithTrace("properties"))
}

func (a *App[acc, f]) initializeFile() {
	a.File = file.New[f](a.Logger.WithTrace("file"), a.Sql)
}

func (a *App[acc, f]) initializeStore() {
	a.Storage = storage.New[f](a.Logger.WithTrace("storage"),
		a.Registry.ValueOf("storage"))
}

func (a *App[acc, f]) initializeUpload() {
	a.Upload = upload.New[f](
		a.Logger.WithTrace("upload"),
		a.Registry.ValueOf("upload"),
		a.Ginger.GetController())
}

func (a *App[acc, f]) initializeDownload() {
	a.Download = download.New[f](
		a.Logger.WithTrace("download"),
		a.Registry.ValueOf("download"),
		a.Properties.UseCase, a.Ginger.GetController())
}
