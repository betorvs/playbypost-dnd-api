package diceroll

import (
	"fmt"

	"github.com/betorvs/dice"
	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/config"
)

// RollInternal struct
type RollInternal struct {
}

func (r RollInternal) DiceRoll(text string) (int, string, error) {
	diceRolled, _, err := dice.Roll(text)
	logLocal := config.GetLogger()
	if err != nil {
		logLocal.Error(err)
		return 0, "No dices to roll", err
	} else {
		message := fmt.Sprintf("Dice Rolled %s and result %v with rolls %s", diceRolled.Description(), diceRolled.Int(), diceRolled.String())
		logLocal.Info(message)
		return diceRolled.Int(), message, nil
	}
}

// DicerInit lazy funcion to init Dice
func DicerInit() appcontext.Component {
	return RollInternal{}
}

//
func init() {
	if config.Values.TestRun {
		return
	}
	logLocal := config.GetLogger()
	logLocal.Info("D20 are ready")
	appcontext.Current.Add(appcontext.Dice, DicerInit)
}
