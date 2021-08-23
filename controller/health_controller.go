package controller

import (
	"net/http"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain"
	"github.com/labstack/echo/v4"
)

//CheckHealth handles the application Health Check
func CheckHealth(c echo.Context) error {
	health := domain.Health{}
	health.Status = "UP"
	return c.JSON(http.StatusOK, health)
}

//CheckReady handles the application Ready Check
func CheckReady(c echo.Context) error {
	health := domain.Health{}
	if config.Values.IsReady.Load().(bool) {
		health.Status = "UP"
		return c.JSON(http.StatusOK, health)
	}
	health.Status = "DOWN"
	return c.JSON(http.StatusServiceUnavailable, health)

}
