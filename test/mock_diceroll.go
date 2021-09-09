package test

import (
	"errors"
	"fmt"
	"strings"

	"github.com/betorvs/playbypost-dnd/appcontext"
)

var (
	// DiceRollCall int
	DiceRollCall int
	// DiceType string
	DiceType string
	// DiceResult int
	DiceResult int
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
	if DiceResult != 0 {
		switch text {
		case "6d6":
			return 36, "Rolled 36", nil
		case "3d6":
			return 18, "Rolled 18", nil
		case "2d6":
			return 12, "Rolled 12", nil

		default:
			return DiceResult, fmt.Sprintf("Rolled %v", DiceResult), nil
		}
	}
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
