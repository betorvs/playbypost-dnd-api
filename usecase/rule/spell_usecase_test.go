package rule

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestGetSpellListDescription(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetSpellListDescription(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"magic-missile"}
	res1 := GetSpellListDescription(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["level"] = []string{"1"}
	res2 := GetSpellListDescription(value2)
	assert.NotEmpty(t, res2)
	value3 := make(url.Values)
	value3["title"] = []string{"Magic Missile"}
	res3 := GetSpellListDescription(value3)
	assert.NotEmpty(t, res3)
}

func TestGetSpellByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := getSpellByName("magic-missile")
	assert.NotEmpty(t, res)
}

func TestSpellDamageIncrease(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test := getSpellByName("fireball")
	res := spellDamageIncrease(test, 9)
	assert.NotEmpty(t, res)
	list := []string{"produce-flame", "acid-splash", "chill-touch"}
	levels := []int{5, 11, 17}
	for _, v := range list {
		test1 := getSpellByName(v)
		for _, value := range levels {
			res1 := spellDamageIncrease(test1, value)
			assert.NotEmpty(t, res1)
		}
	}
}

func TestSpellDamageMax(t *testing.T) {
	res := spellDamageMax("6d6")
	assert.NotEmpty(t, res)
}

func TestSpellLevelRegex(t *testing.T) {
	res := spellLevelRegex("1st")
	assert.NotEmpty(t, res)
}

func TestSpellHealIncreases(t *testing.T) {
	res := spellHealIncreases("3d8", "healing increases by 1d8 for each slot level above 5th.", 10)
	assert.NotEmpty(t, res)
}

func TestHealSpell(t *testing.T) {
	res1 := healSpell(5)
	assert.NotEmpty(t, res1)
	res2 := healSpell(10)
	assert.NotEmpty(t, res2)
}

func TestGetFullSpellList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	// list := []string{"paladin", "bard"}
	for _, v := range ClassWithSpell() {
		res := GetFullSpellList(v)
		assert.NotNil(t, res)
	}
}

func TestGetFullListWithFeature(t *testing.T) {
	testFeatures := ClassFeatures("cleric", 3)
	testFeatures = append(testFeatures, "domain-war")
	res := getFullListWithFeature("cleric", testFeatures, 3)
	assert.NotEmpty(t, res)
}
