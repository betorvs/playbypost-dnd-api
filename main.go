package main

import (
	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/controller"
	_ "github.com/betorvs/playbypost-dnd/gateway/customlog"
	_ "github.com/betorvs/playbypost-dnd/gateway/mongodb"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
