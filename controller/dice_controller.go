package controller

import (
	"fmt"
	"net/http"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/labstack/echo/v4"
)

func RunDiceRoll(c echo.Context) (err error) {
	diceRoll := c.Param("dice")
	r := diceroll.GetDice()
	diceRolled, err := r.DiceRoll(diceRoll)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	output := fmt.Sprintf("Dice Rolled %s and result %v with rolls %s", diceRolled.Description(), diceRolled.Int(), diceRolled.String())
	return c.JSON(http.StatusOK, output)
}
