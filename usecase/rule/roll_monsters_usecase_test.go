package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestCalcFullMonsterAttackwithWeapon(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.MonsterRoll)
	test1.Name = "orc"
	test1.DifficultClass = 10
	test1.Weapon = "greataxe"
	test.DiceResult = 20
	res1 := CalcFullMonsterAttackwithWeapon(test1)
	assert.NotEmpty(t, res1)
	test1.Name = "kobold"
	test1.Weapon = "dagger"
	test1.Advantages = []string{"strength"}
	test1.EnemyRage = true
	test.DiceResult = 1
	res2 := CalcFullMonsterAttackwithWeapon(test1)
	assert.NotEmpty(t, res2)
	test1.AutoFail = []string{"all"}
	res3 := CalcFullMonsterAttackwithWeapon(test1)
	assert.NotEmpty(t, res3)
}

func TestCalcMonsterSavingsAbility(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.MonsterRoll)
	test1.Name = "orc"
	test1.DifficultClass = 10
	test1.Check = "constitution"
	test.DiceResult = 8
	res1 := CalcMonsterSavingsAbility(test1)
	assert.NotEmpty(t, res1)
	test1.Name = "kobold"
	test1.AutoFail = []string{"all"}
	res2 := CalcMonsterSavingsAbility(test1)
	assert.NotEmpty(t, res2)
}

func TestCalcMonsterChecks(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.MonsterRoll)
	test1.AutoFail = []string{"intelligence"}
	test1.Check = "intelligence"
	test1.Name = "skeleton"
	res1 := CalcMonsterChecks(test1)
	assert.NotEmpty(t, res1)
	test2 := new(rule.MonsterRoll)
	test2.Check = "intelligence"
	test2.Name = "kobold"
	test2.DifficultClass = 10
	test.DiceResult = 8
	res2 := CalcMonsterChecks(test2)
	assert.NotEmpty(t, res2)
	test.DiceResult = 20
	res3 := CalcMonsterChecks(test2)
	assert.NotEmpty(t, res3)
	test2.Name = "orc"
	test2.Check = "intimidation"
	res4 := CalcMonsterChecks(test2)
	assert.NotEmpty(t, res4)

}

func TestCalcMonstersInitiative(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.SimpleList)
	test1.List = []string{"kobold", "skeleton"}
	res1 := CalcMonstersInitiative(test1)
	assert.NotEmpty(t, res1)
}

func TestTurnUndeadRolls(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.MonsterTurn)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	test1.Ability = ability
	res1, err1 := TurnUndeadRolls(test1)
	assert.Empty(t, res1)
	assert.Error(t, err1)
	test1.Level = 5
	test1.MonsterList = []string{"skeleton"}
	test1.ClassFeatures = append(test1.ClassFeatures, "channel-divinity")
	test.DiceResult = 10
	res2, err2 := TurnUndeadRolls(test1)
	assert.NotEmpty(t, res2)
	assert.NoError(t, err2)
	test.DiceResult = 20
	res3, err3 := TurnUndeadRolls(test1)
	assert.NotEmpty(t, res3)
	assert.NoError(t, err3)
	test.DiceResult = 10
	test1.ClassFeatures = []string{"sacred-oath-of-devotion-oauth-spells-and-channel-divinity", "sacred-oath-of-devotion-channel-divinity"}
	res4, err4 := TurnUndeadRolls(test1)
	assert.NotEmpty(t, res4)
	assert.NoError(t, err4)
	test.DiceResult = 20
	test1.ClassFeatures = []string{"sacred-oath-of-ancients-oauth-spells-and-channel-divinity", "sacred-oath-of-ancients-channel-divinity"}
	res5, err5 := TurnUndeadRolls(test1)
	assert.NotEmpty(t, res5)
	assert.NoError(t, err5)

}

func TestAttackMonster(t *testing.T) {}

func TestSavingRollMonster(t *testing.T) {}

func TestTouchRollMonster(t *testing.T) {}

func TestValidateDamage(t *testing.T) {}
