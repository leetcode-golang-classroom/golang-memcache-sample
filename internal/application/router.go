package application

import (
	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/service/photo"
)

func (app *App) SetupRoutes() {
	api := app.fiberApp.Group("api")
	handler := NewHandler()
	api.Get("/", handler.HealthCheck)
	photoHdr := photo.NewHandler(app.cfg.JsonServerURL, app.appCache)
	router := photo.NewRouter(photoHdr)
	router.SetupRoutes(api)
}
