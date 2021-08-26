package controller

import (
	"net/http"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/labstack/echo/v4"
)

func RunDiceRoll(c echo.Context) (err error) {
	diceRoll := c.Param("dice")
	r := diceroll.GetDice()
	_, output, err := r.DiceRoll(diceRoll)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, output)
}
