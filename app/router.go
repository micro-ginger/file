package app

func (a *App[acc, f]) registerRoutes() {
	rg := a.Ginger.NewRouterGroup("/")

	if a.Config.Upload.Enabled {
		rg.Create("/upload", a.Upload.UploadHandler)
	}
	rg.Read("/download/k/:id", a.Download.DownloadHandler)
}
