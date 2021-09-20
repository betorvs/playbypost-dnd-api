package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestCalcSkillOrAbility(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	test1 := new(rule.SkillOrAbilityCheck)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	test1.Ability = ability
	test1.Race = "human"
	test1.Check = "dexterity"
	test1.Level = 6
	test1.DifficultClass = 10
	test1.MagicBonus = 1
	test.DiceResult = 2
	res1 := CalcSkillOrAbility(test1)
	assert.NotEmpty(t, res1)
	test1.Skills = []string{"acrobatics"}
	test1.Check = "acrobatics"
	test1.DifficultClass = 14
	test.DiceResult = 10
	res2 := CalcSkillOrAbility(test1)
	assert.NotEmpty(t, res2)
	test2 := new(rule.SkillOrAbilityCheck)
	test2.Ability = ability
	test2.Race = "halfling"
	test2.Check = "dexterity"
	test2.Level = 10
	test2.DifficultClass = 10
	test.DiceResult = 1
	res3 := CalcSkillOrAbility(test2)
	assert.NotEmpty(t, res3)
	test2.Advantages = []string{"dexterity"}
	res4 := CalcSkillOrAbility(test2)
	assert.NotEmpty(t, res4)
	test3 := new(rule.SkillOrAbilityCheck)
	test3.Ability = ability
	test3.Race = "human"
	test3.Check = "dexterity"
	test3.Level = 10
	test3.DifficultClass = 10
	test3.AutoFail = []string{"dexterity"}
	test.DiceResult = 10
	res5 := CalcSkillOrAbility(test3)
	assert.NotEmpty(t, res5)
	test3.AutoFail = []string{}
	test3.Disvantages = []string{"dexterity"}
	res67 := CalcSkillOrAbility(test3)
	assert.NotEmpty(t, res67)
}
