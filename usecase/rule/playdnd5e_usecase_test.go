package rule

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcProficiency(t *testing.T) {
	test1 := CalcProficiency(5)
	exp1 := 3
	assert.Equal(t, exp1, test1)
}

func TestCalcAbilityModifier(t *testing.T) {
	test1 := CalcAbilityModifier(8)
	exp1 := -1
	assert.Equal(t, exp1, test1)
	test2 := CalcAbilityModifier(18)
	exp2 := 4
	assert.Equal(t, exp2, test2)
}

func TestCalcMaxHP(t *testing.T) {
	test1 := CalcMaxHP(1, 6, 2)
	exp1 := 8
	assert.Equal(t, exp1, test1)
	test2 := CalcMaxHP(2, 6, 2)
	exp2 := 13
	assert.Equal(t, exp2, test2)
}

func TestXPNeeded(t *testing.T) {
	test1 := XPNeeded(0)
	exp1 := 0
	assert.Equal(t, exp1, test1)
	test2 := XPNeeded(19)
	exp2 := 355000
	assert.Equal(t, exp2, test2)
}

func TestCalculateSpellList(t *testing.T) {
	test1map, test1int := CalculateSpellList("wizard", 3)
	assert.NotEmpty(t, test1map)
	assert.GreaterOrEqual(t, test1int, 2)
}

func TestCalculateClassFeatureList(t *testing.T) {
	test1 := "ranger"
	level1 := 1
	chosen1 := []string{"forest", "giants"}
	res1 := CalculateClassFeatureList(test1, level1, chosen1)
	assert.GreaterOrEqual(t, len(res1), 2)
	test2 := "ranger"
	level2 := 2
	chosen2 := []string{"dueling"}
	res2 := CalculateClassFeatureList(test2, level2, chosen2)
	assert.GreaterOrEqual(t, len(res2), 2)
}

func TestSpellSlotsMultiClass(t *testing.T) {
	for _, v := range ClassWithSpell() {
		res := SpellSlotsMultiClass(v, 4)
		exp := 2
		if strings.Contains(v, "-") || v == "warlock" {
			exp = 0
		}
		assert.GreaterOrEqual(t, res, exp)
	}
}

func TestSpellsPerLevel(t *testing.T) {
	for _, v := range ClassWithSpell() {
		sum := 0
		if v == "warlock" {
			continue
		}
		for i := 1; i < 21; i++ {
			res := SpellsPerLevel(v, i)
			assert.GreaterOrEqual(t, len(res), 1)
			sum += i
		}
	}

}

func TestRaceStatistics(t *testing.T) {}

func TestRaceArmorProficiencyExtra(t *testing.T) {}

func TestBackgroundStatistics(t *testing.T) {}

func TestClassStatistics(t *testing.T) {}

func TestClassFeatures(t *testing.T) {}

func TestSpellKnown(t *testing.T) {}

func TestCantripsKnown(t *testing.T) {}

func TestBarbarianClass(t *testing.T) {}

func TestMonkClass(t *testing.T) {}

func TestRogueClass(t *testing.T) {}

func TestWarlockClass(t *testing.T) {}

func TestRaceSpecialTrait(t *testing.T) {}

func TestCoinList(t *testing.T) {}

func TestCoinShortnameList(t *testing.T) {}

func TestCoinShortName(t *testing.T) {}

func TestExchangeRates(t *testing.T) {}
