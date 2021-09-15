package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestClassFeatureRoll(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.Feature)
	test1.Name = "unkown"
	res1, err1 := ClassFeatureRoll(test1)
	assert.NotEmpty(t, res1)
	assert.Error(t, err1)
	test2 := new(rule.Feature)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	test2.Ability = ability
	featuresList := []string{
		"song-of-rest",
		"bardic-inspiration",
		"second-wind",
		"lay-on-hands",
		"bless",
		"combat-wild-shape",
		"combat-superiority",
		"radiance-of-the-dawn",
		"grim-harvest",
	}
	test2.Level = 1
	for _, v := range featuresList {
		test2.ClassFeatures = append(test2.ClassFeatures, v)
		switch v {
		case "song-of-rest":
			test2.Name = v
			res3, err3 := ClassFeatureRoll(test2)
			assert.Empty(t, res3)
			assert.Error(t, err3)
		case "combat-wild-shape":
			test2.ClassFeatures = append(test2.ClassFeatures, "circle-of-the-moon-combat-wild-shape")
			test2.UsingFeatureSlot = 1
		case "combat-superiority":
			test2.ClassFeatures = append(test2.ClassFeatures, "archetype-battle-master-combat-superiority")
		case "radiance-of-the-dawn":
			test2.ClassFeatures = append(test2.ClassFeatures, "channel-divinity-radiance-of-the-dawn")
			test2.MonsterList = []string{"kobold"}
		case "grim-harvest":
			test2.ClassFeatures = append(test2.ClassFeatures, "arcane-tradition-school-of-necromancy-grim-harvest")
			test2.GenericList = []string{"chill-touch"}
			test2.Name = v
			res3, err3 := ClassFeatureRoll(test2)
			assert.NotEmpty(t, res3)
			assert.NoError(t, err3)
			test2.GenericList = []string{"magic-missile"}
		}
		test2.Name = v
		test2.Level = 3
		res2, err2 := ClassFeatureRoll(test2)
		assert.NotEmpty(t, res2)
		assert.NoError(t, err2)

	}
}

func TestSpecialRaceFeature(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.SpecialRaceFeature)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	test1.Ability = ability
	test1.Race = "human"
	res1, err1 := SpecialRaceFeature(test1)
	assert.Empty(t, res1)
	assert.Error(t, err1)
	test1.Race = "tiefling"
	test1.Name = "hellish-rebuke"
	res2, err2 := SpecialRaceFeature(test1)
	assert.NotEmpty(t, res2)
	assert.NoError(t, err2)
	test1.Monster = []string{"kobold"}
	res3, err3 := SpecialRaceFeature(test1)
	assert.NotEmpty(t, res3)
	assert.NoError(t, err3)
}
