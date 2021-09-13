package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestCalcFullAttackwithWeapon(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.Attack)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	test1.Ability = ability
	test1.Race = "human"
	test1.Weapon = []string{"longsword"}
	test1.ArmorProficiency = ProficiencyFullList()
	test1.ClassFeatures = ClassFeatures("fighter", 1)
	test1.ClassFeatures = append(test1.ClassFeatures, "fighting-style-great-weapon-fighting")
	test1.Monster = "kobold"
	test1.TwoHands = true
	test1.MagicBonus = 1
	test1.Level = 1
	test.DiceResult = 20
	res1 := CalcFullAttackwithWeapon(test1)
	assert.NotEmpty(t, res1)
	test2 := new(rule.Attack)
	test2.Ability = ability
	test2.Ability["dexterity"] = 18
	test2.Race = "human"
	test2.Weapon = []string{"dagger"}
	test2.ArmorProficiency = ProficiencyFullList()
	test2.ClassFeatures = ClassFeatures("fighter", 1)
	test2.ClassFeatures = append(test2.ClassFeatures, "fighting-style-dueling")
	test2.Monster = "kobold"
	test2.Level = 1
	test2.Advantages = []string{"dexterity"}
	test.DiceResult = 20
	res2 := CalcFullAttackwithWeapon(test2)
	assert.NotEmpty(t, res2)
	test2.Weapon = []string{"crossbow-light"}
	test2.ClassFeatures = append(test2.ClassFeatures, "fighting-style-archery")
	res3 := CalcFullAttackwithWeapon(test2)
	assert.NotEmpty(t, res3)
	test3 := new(rule.Attack)
	test3.Ability = ability
	test3.Race = "half-orc"
	test3.Weapon = []string{"unarmed"}
	test3.ArmorProficiency = ProficiencyFullList()
	test3.ClassFeatures = ClassFeatures("monk", 1)
	test3.Monster = "kobold"
	test3.Level = 1
	test.DiceResult = 20
	res4 := CalcFullAttackwithWeapon(test3)
	assert.NotEmpty(t, res4)
	test4 := new(rule.Attack)
	test4.Ability = ability
	test4.Ability["dexterity"] = 18
	test4.Race = "human"
	test4.Weapon = []string{"dagger"}
	test4.ArmorProficiency = ProficiencyFullList()
	test4.ClassFeatures = ClassFeatures("rogue", 1)
	test4.Monster = "kobold"
	test4.Level = 1
	test4.UsingFeature = "sneak-attack"
	test.DiceResult = 20
	res5 := CalcFullAttackwithWeapon(test4)
	assert.NotEmpty(t, res5)
}

func TestCalcBonusAttackPerWeapon(t *testing.T) {}

func TestCheckCriticalHit(t *testing.T) {}
