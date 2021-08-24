package diceroll

import (
	"github.com/betorvs/dice"
	"github.com/betorvs/playbypost-dnd/appcontext"
)

//Dice interface
type Dice interface {
	appcontext.Component
	// DiceRoll receives text string
	DiceRoll(text string) (dice.RollResult, error)
}

//GetDice gets the Dice current implementation
func GetDice() Dice {
	return appcontext.Current.Get(appcontext.Dice).(Dice)
}
