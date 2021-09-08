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

func TestSpellDamageIncrease(t *testing.T) {}

func TestSpellDamageMax(t *testing.T) {}

func TestSpellLevelRegex(t *testing.T) {}

func TestSpellHealIncreases(t *testing.T) {}

func TestHealSpell(t *testing.T) {}

func TestGetFullSpellList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	list := []string{"paladin", "bard"}
	for _, v := range list {
		res := GetFullSpellList(v)
		assert.NotNil(t, res)
	}
}

func TestGetFullListWithFeature(t *testing.T) {}

func TestGetSpellListByClass(t *testing.T) {

}
