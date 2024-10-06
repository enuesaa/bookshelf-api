package router

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/middleware"
	apiTeapods "github.com/enuesaa/teatime/pkg/routes/teapods"
	apiTeapodsInfo "github.com/enuesaa/teatime/pkg/routes/teapods/info"
	apiTeapodsTeas "github.com/enuesaa/teatime/pkg/routes/teapods/teas"
	"github.com/enuesaa/teatime/ui"
	"github.com/labstack/echo/v4"
)

func New(repos repository.Repos) *echo.Echo {
	app := echo.New()
	app.Use(middleware.BindCtx(repos))
	app.Use(middleware.Logger)
	app.Use(middleware.Cors)

	api := app.Group("/api")
	api.Use(middleware.HandleData)
	api.Use(middleware.HandleError)

	api.GET("/teapods", apiTeapods.List)
	api.GET("/teapods/:teapod/info", apiTeapodsInfo.View)
	api.GET("/teapods/:teapod/teas", apiTeapodsTeas.List)
	api.GET("/teapods/:teapod/teas/:teaId", apiTeapodsTeas.View)

	app.Any("/*", ui.Serve)

	return app
}
