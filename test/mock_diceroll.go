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
	switch DiceType {
	case "1d100b20":
		return 14, "Rolled 14", nil
	case "1d100e35":
		return 35, "Rolled 35", nil
	case "1d100e65":
		return 65, "Rolled 65", nil
	case "1d100e80":
		return 80, "Rolled 80", nil
	case "1d100e96":
		return 96, "Rolled 96", nil
	case "1d100e0":
		return 101, "Rolled 101", nil
	case "1d100e96h":
		fmt.Println(text)
		switch text {
		case "1d100":
			return 96, "Rolled 96", nil
		case "6d6":
			return 36, "Rolled 36", nil
		case "3d6":
			return 18, "Rolled 18", nil
		case "2d6":
			return 12, "Rolled 12", nil

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

// InitDice100b20Mock func returns a RepositoryDiceMock interface
func InitDice100b20Mock() appcontext.Component {
	DiceType = "1d100b20"
	return MockRollInternal{}
}

// InitDice100e35Mock func returns a RepositoryDiceMock interface
func InitDice100e35Mock() appcontext.Component {
	DiceType = "1d100e35"
	return MockRollInternal{}
}

// InitDice100e65Mock func returns a RepositoryDiceMock interface
func InitDice100e65Mock() appcontext.Component {
	DiceType = "1d100e65"
	return MockRollInternal{}
}

// InitDice100e80Mock func returns a RepositoryDiceMock interface
func InitDice100e80Mock() appcontext.Component {
	DiceType = "1d100e80"
	return MockRollInternal{}
}

// InitDice100e96Mock func returns a RepositoryDiceMock interface
func InitDice100e96Mock() appcontext.Component {
	DiceType = "1d100e96"
	return MockRollInternal{}
}

// InitDice100e0Mock func returns a RepositoryDiceMock interface
func InitDice100e0Mock() appcontext.Component {
	DiceType = "1d100e0"
	return MockRollInternal{}
}

// InitDice100e96HoardMock func returns a RepositoryDiceMock interface
func InitDice100e96HoardMock() appcontext.Component {
	DiceType = "1d100e96h"
	return MockRollInternal{}
}
