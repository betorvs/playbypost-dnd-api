package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestCalcSpellcastAttackAndSave(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.SpellcastAbility)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	test1.Ability = ability
	test1.AutoFail = []string{"spellcast"}
	res1 := CalcSpellcastAttackAndSave(test1)
	assert.NotEmpty(t, res1)
	test2 := new(rule.SpellcastAbility)
	test2.Ability = ability
	test2.Level = 4
	test2.Class = "wizard"
	test2.SpellName = "cure-wounds"
	res2 := CalcSpellcastAttackAndSave(test2)
	assert.NotEmpty(t, res2)
	test2.SpellName = "fireball"
	res3 := CalcSpellcastAttackAndSave(test2)
	assert.NotEmpty(t, res3)
	test2.Level = 5
	test2.Monster = []string{"kobold"}
	test2.MagicBonus = 1
	res4 := CalcSpellcastAttackAndSave(test2)
	assert.NotEmpty(t, res4)
	test2.SpellLevel = 4
	res5 := CalcSpellcastAttackAndSave(test2)
	assert.NotEmpty(t, res5)
	test2.SpellName = "chill-touch"
	test2.Level = 15
	test2.SpellLevel = 6
	test2.Advantages = []string{"intelligence"}
	test.DiceResult = 10
	res6 := CalcSpellcastAttackAndSave(test2)
	assert.NotEmpty(t, res6)
	test3 := new(rule.SpellcastAbility)
	test3.Ability = ability
	test3.Level = 5
	test3.Class = "cleric"
	test3.SpellName = "cure-wounds"
	res7 := CalcSpellcastAttackAndSave(test3)
	assert.NotEmpty(t, res7)

}

func TestSpellForHealing(t *testing.T) {}
