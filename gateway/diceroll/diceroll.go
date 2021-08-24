package diceroll

import (
	"github.com/betorvs/dice"
	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/config"
)

// RollInternal struct
type RollInternal struct {
}

func (r RollInternal) DiceRoll(text string) (dice.RollResult, error) {
	diceRolled, _, err := dice.Roll(text)
	logLocal := config.GetLogger()
	if err != nil {
		logLocal.Error(err)
	} else {
		logLocal.Info("Dice Rolled ", diceRolled.Description(), " and result ", diceRolled.Int(), " with rolls ", diceRolled.String())
	}
	return diceRolled, err
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
