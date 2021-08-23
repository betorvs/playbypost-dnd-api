package controller

import (
	"net/http"
	"time"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain"
	"github.com/labstack/echo/v4"
)

//GetInfo of the application like version
func GetInfo(c echo.Context) error {
	info := domain.Info{Version: config.Version}
	time.Sleep(8 * time.Second)
	return c.JSON(http.StatusOK, info)
}
