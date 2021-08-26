package rule

import (
	"net/http"

	usecase "github.com/betorvs/playbypost-dnd/usecase/rule"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
)

// GetMagicItem controller
func GetMagicItem(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name", "title"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name, title"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetMagicItem(queryParams)
	return c.JSON(http.StatusOK, result)
}
