package rule

import (
	"net/http"
	"strconv"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	usecase "github.com/betorvs/playbypost-dnd/usecase/rule"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
)

// GetSpellListDescription controller
func GetSpellListDescription(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name", "level", "title"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetSpellListDescription(queryParams)
	// fmt.Println(result)
	return c.JSON(http.StatusOK, result)
}

// ListSpellByClass controller
func ListSpellByClass(c echo.Context) (err error) {
	class := c.Param("class")
	if !utils.StringInSlice(class, usecase.ClassWithSpell()) {
		errString := "Class cannot spell"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if c.Param("level") == "all" {
		// special case to list all spells by class
		var result rule.SimpleList
		result.List = usecase.GetFullSpellList(class)
		return c.JSON(http.StatusOK, result)
	}
	initialLevel := c.Param("level")
	level, err := strconv.Atoi(initialLevel)
	if err != nil {
		errString := "Value not allowed. Use valid int."
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if level < 0 || level > 9 {
		errString := "Value not allowed. Use valid int: 0 to 9"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.GetSpellListByClass(class, level)
	return c.JSON(http.StatusOK, result)
}
