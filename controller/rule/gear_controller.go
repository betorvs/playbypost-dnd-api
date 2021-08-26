package rule

import (
	"net/http"
	"strconv"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	usecase "github.com/betorvs/playbypost-dnd/usecase/rule"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
)

// GetAllWeapons controller
func GetAllWeapons(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name", "kind", "type", "damage_type"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name, kind, damage_type"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetAllWeapons(queryParams)
	return c.JSON(http.StatusOK, result)
}

// GetAllArmors controller
func GetAllArmors(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name", "kind"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name, kind"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetAllArmor(queryParams)
	return c.JSON(http.StatusOK, result)
}

// GetAllGear controller
func GetAllGear(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name", "kind"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name, kind"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetAllGears(queryParams)
	return c.JSON(http.StatusOK, result)
}

// GetAllPacks controller
func GetAllPacks(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetAllPacks(queryParams)
	return c.JSON(http.StatusOK, result)
}

// GetAllTools controller
func GetAllTools(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name", "kind"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name, kind"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetAllTools(queryParams)
	return c.JSON(http.StatusOK, result)
}

// GetAllMounts controller
func GetAllMounts(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetAllMounts(queryParams)
	return c.JSON(http.StatusOK, result)
}

// CalcShop controller
func CalcShop(c echo.Context) (err error) {
	list := new(rule.SimpleList)
	if err = c.Bind(list); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcShoppingCart(list)
	return c.JSON(http.StatusOK, result)
}

// RandomTreasure controller
func RandomTreasure(c echo.Context) (err error) {
	initialLevel := c.Param("level")
	level, err := strconv.Atoi(initialLevel)
	if err != nil {
		errString := "Value not allowed. Use valid int."
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if level < 0 || level > 30 {
		errString := "Value not allowed. Use valid int: 0 to 30"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := usecase.CalcRandomTreasureByChallengeLevel(level, true)
	if err != nil {
		errString := "Cannot calculate"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}

	return c.JSON(http.StatusOK, result)
}

// FastTreasure controller
func FastTreasure(c echo.Context) (err error) {
	initialLevel := c.Param("level")
	level, err := strconv.Atoi(initialLevel)
	if err != nil {
		errString := "Value not allowed. Use valid int."
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if level < 0 || level > 30 {
		errString := "Value not allowed. Use valid int: 0 to 30"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := usecase.CalcRandomTreasureByChallengeLevel(level, false)
	if err != nil {
		errString := "Cannot calculate"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}

	return c.JSON(http.StatusOK, result)
}

// RandomTreasureHoard controller
func RandomTreasureHoard(c echo.Context) (err error) {
	initialLevel := c.Param("level")
	level, err := strconv.Atoi(initialLevel)
	if err != nil {
		errString := "Value not allowed. Use valid int."
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if level < 0 || level > 30 {
		errString := "Value not allowed. Use valid int: 0 to 30"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcHoardPercentageByLevel(level)

	return c.JSON(http.StatusOK, result)
}

// GetAllServices controller
func GetAllServices(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name", "source"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name, source"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.GetAllServices(queryParams)
	return c.JSON(http.StatusOK, result)
}
