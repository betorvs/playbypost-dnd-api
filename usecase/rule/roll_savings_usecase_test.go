package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestCalcSavingsAbility(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	test1 := new(rule.SavingsCheck)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	test1.Ability = ability
	test1.AutoFail = []string{"constitution"}
	test1.Saving = "constitution"
	test1.DifficultyClass = 10
	test.DiceResult = 10
	res1 := CalcSavingsAbility(test1)
	assert.NotEmpty(t, res1)
	test2 := new(rule.SavingsCheck)
	test2.Ability = ability
	test2.Saving = "constitution"
	test2.DifficultyClass = 10
	test2.Savings = []string{"strength", "constitution"}
	test2.Level = 6
	test2.MagicBonus = 1
	res2 := CalcSavingsAbility(test2)
	assert.NotEmpty(t, res2)
	test2.Race = "halfling"
	test.DiceResult = 1
	res3 := CalcSavingsAbility(test2)
	assert.NotEmpty(t, res3)
	test2.Advantages = []string{"constitution"}
	test.DiceResult = 2
	res4 := CalcSavingsAbility(test2)
	assert.NotEmpty(t, res4)
	test3 := new(rule.SavingsCheck)
	test3.Ability = ability
	test3.Saving = "intelligence"
	test3.DifficultyClass = 10
	test3.Savings = []string{"strength", "constitution"}
	test3.Level = 6
	test3.Check = "spell"
	test3.Race = "gnome"
	res5 := CalcSavingsAbility(test3)
	assert.NotEmpty(t, res5)
}
