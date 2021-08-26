package rule

import (
	"net/http"

	usecase "github.com/betorvs/playbypost-dnd/usecase/rule"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
)

// func getAllMonsters(c echo.Context) (err error) {
// 	result := usecase.ReadAllMonsterFromFile()
// 	return c.JSON(http.StatusOK, result)
// }

// func getOneMonster(c echo.Context) (err error) {
// 	name := c.Param("name")
// 	result := usecase.ReadOneMonsterFromFile(name)
// 	return c.JSON(http.StatusOK, result)
// }

// func getMonsterByChallenge(c echo.Context) (err error) {
// 	queryParams := c.QueryParams()
// 	allowedParams := []string{"challenge", "name", "xp"}
// 	for paramName := range queryParams {
// 		if !utils.StringInSlice(paramName, allowedParams) {
// 			errString := fmt.Sprintf("Allowed query params: challenge, name and xp")
// 			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
// 		}
// 	}
// 	result := usecase.MonsterByChallenge(queryParams)
// 	return c.JSON(http.StatusOK, result)
// }

// GetMonsterForNPC controller
func GetMonsterForNPC(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	allowedParams := []string{"name", "xp", "type"}
	for paramName := range queryParams {
		if !utils.StringInSlice(paramName, allowedParams) {
			errString := "Allowed query params: name"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}
	result := usecase.MonsterForNPC(queryParams)
	return c.JSON(http.StatusOK, result)
}
