package test

import (
	"errors"
	"strings"

	"github.com/betorvs/playbypost-dnd/appcontext"
)

var (
	// DiceRollCall int
	DiceRollCall int
)

// MockRollInternal struct
type MockRollInternal struct {
}

//Dice interface
type Dice interface {
	appcontext.Component
	// DiceRoll receives text string
	DiceRoll(text string) (int, string, error)
}

// DiceRoll
func (r MockRollInternal) DiceRoll(text string) (int, string, error) {
	if strings.Contains(text, "h") {
		DiceRollCall = 0
		return 1, "Rolled 1", errors.New("invalid dice")
	}
	DiceRollCall++
	return 20, "Rolled 20", nil
}

//GetDice gets the Dice current implementation
func GetDice() Dice {
	return appcontext.Current.Get(appcontext.Dice).(Dice)
}

// InitDiceMock func returns a RepositoryDiceMock interface
func InitDiceMock() appcontext.Component {

	return MockRollInternal{}
}
