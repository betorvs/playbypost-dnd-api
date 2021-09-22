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
	test1 := ClassStatistics("ranger", 1)
	level1 := 1
	chosen1 := []string{"forest", "giants"}
	res1 := CalculateClassFeatureList(test1.Features, level1, chosen1)
	assert.GreaterOrEqual(t, len(res1), 2)
	test2 := ClassStatistics("ranger", 2)
	level2 := 2
	chosen2 := []string{"dueling", "fighting-style"}
	res2 := CalculateClassFeatureList(test2.Features, level2, chosen2)
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

func TestRaceStatistics(t *testing.T) {
	for _, value := range RaceList() {
		subrace := []string{""}
		switch value {
		case "dwarf":
			subrace = append(subrace, "hill-dwarf", "mountain-dwarf")
		case "elf":
			subrace = append(subrace, "high-elf", "wood-elf", "drow")
		case "halfling":
			subrace = append(subrace, "lightfoot", "stout")
		case "dragonborn":
			subrace = append(subrace, dragonKinds()...)
		case "gnome":
			subrace = append(subrace, "rock-gnome", "forest-gnome")
		}

		for _, v := range subrace {
			//size, speedMeasure, speed, ability, _, language, _, _, _, _, _ := RaceStatistics(value, v)
			res := RaceStatistics(value, v)
			assert.NotEmpty(t, res.Size)
			assert.NotEmpty(t, res.Speedmeasure)
			assert.NotEmpty(t, res.Speed)
			assert.NotEmpty(t, res.Ability)
			assert.NotEmpty(t, res.Language)
		}

	}
}

func TestRaceArmorProficiencyExtra(t *testing.T) {
	// test1 := []string{"Elf Weapon Training", "Dwarven Combat Training", "Dwarven Armor Training"}
	// for _, v := range test1 {
	res1 := RaceStatistics("elf", "high-elf")
	assert.GreaterOrEqual(t, len(res1.ArmorProficiency), 2)
	assert.NotEmpty(t, res1)
	res2 := RaceStatistics("dwarf", "mountain-dwarf")
	assert.GreaterOrEqual(t, len(res2.ArmorProficiency), 2)
	assert.NotEmpty(t, res2)
	// }
	test2 := RaceStatistics("human", "")
	assert.Empty(t, test2.ArmorProficiency)
}

func TestBackgroundStatistics(t *testing.T) {
	for _, v := range BackgroundList() {
		res1 := BackgroundStatistics(v)
		assert.NotEmpty(t, res1)
	}
}

func TestClassDetails(t *testing.T) {
	for _, v := range ClassList() {
		hitDice, savings, armor, skill := ClassDetails(v)
		assert.NotEmpty(t, hitDice)
		assert.NotEmpty(t, savings)
		assert.NotEmpty(t, armor)
		assert.NotEmpty(t, skill)
	}
}

func TestClassFeatures(t *testing.T) {
	for _, value := range ClassList() {
		for i := 1; i < 21; i++ {
			res := ClassFeatures(value, i)
			assert.NotNil(t, res)
		}
	}
}

func TestSpellKnown(t *testing.T) {
	for _, v := range ClassWithSpellKnown() {
		for i := 1; i < 21; i++ {
			res := SpellKnown(v, i)
			assert.NotNil(t, res)
			if i >= 3 {
				assert.GreaterOrEqual(t, res, 3)
			}
		}
	}
}

func TestCantripsKnown(t *testing.T) {
	for _, v := range ClassWithCantrips() {
		for i := 1; i < 21; i++ {
			res := CantripsKnown(v, i)
			assert.NotNil(t, res)
			assert.GreaterOrEqual(t, res, 2)
		}
	}
}

func TestBarbarianClass(t *testing.T) {
	for i := 1; i < 21; i++ {
		res1, res2 := BarbarianClass(i)
		assert.NotNil(t, res1)
		assert.GreaterOrEqual(t, res1, 2)
		assert.NotNil(t, res2)
		assert.GreaterOrEqual(t, res2, 2)
	}
	test1, _ := BarbarianClass(0)
	assert.Empty(t, test1)
}

func TestMonkClass(t *testing.T) {
	for i := 1; i < 21; i++ {
		res1, res2, res3 := MonkClass(i)
		assert.NotEmpty(t, res1)
		assert.NotNil(t, res2)
		if i >= 2 {
			assert.GreaterOrEqual(t, res2, 2)
		}
		assert.NotEmpty(t, res3)
	}
	test1, _, _ := MonkClass(0)
	assert.Empty(t, test1)
}

func TestRogueClass(t *testing.T) {
	for i := 1; i < 21; i++ {
		res1 := RogueClass(i)
		assert.NotEmpty(t, res1)
	}
	test1 := RogueClass(0)
	assert.Empty(t, test1)
}

func TestWarlockClass(t *testing.T) {
	for i := 1; i < 21; i++ {
		res1, res2, res3 := WarlockClass(i)
		assert.NotEmpty(t, res2)
		assert.NotEmpty(t, res1)
		assert.GreaterOrEqual(t, res1, 1)
		if i >= 2 {
			assert.NotEmpty(t, res3)
			assert.GreaterOrEqual(t, res3, 2)
		}
	}
	test1, _, _ := WarlockClass(0)
	assert.Empty(t, test1)
}

func TestRaceSpecialTrait(t *testing.T) {
	test := []string{"dragonborn", "tiefling", "gnome", "elf"}
	ability := map[string]int{}
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	for _, value := range test {
		switch value {
		case "dragonborn":
			levels := []int{1, 6, 11, 16}
			for _, l := range levels {
				for _, v := range dragonKinds() {
					name, _, spellcast, damageDice, damageType, savingThrow, description, difficult := RaceSpecialTrait(value, v, l, ability)
					assert.NotEmpty(t, name)
					assert.NotNil(t, spellcast)
					assert.NotEmpty(t, damageDice)
					assert.NotEmpty(t, damageType)
					assert.NotEmpty(t, savingThrow)
					assert.NotEmpty(t, description)
					assert.NotEmpty(t, difficult)
				}
			}
		case "tiefling":
			levels := []int{1, 3, 5}
			for _, l := range levels {
				name, spell, spellcast, _, _, _, description, difficult := RaceSpecialTrait(value, "", l, ability)
				assert.NotEmpty(t, name)
				assert.NotNil(t, spell)
				assert.NotNil(t, spellcast)
				assert.NotEmpty(t, description)
				assert.NotEmpty(t, difficult)
			}
		case "gnome":
			name, spell, spellcast, _, _, _, description, difficult := RaceSpecialTrait(value, "forest-gnome", 1, ability)
			assert.NotEmpty(t, name)
			assert.NotNil(t, spell)
			assert.NotNil(t, spellcast)
			assert.NotEmpty(t, description)
			assert.NotEmpty(t, difficult)
		case "elf":
			levels := []int{1, 3, 5}
			for _, l := range levels {
				name, spell, spellcast, _, _, _, description, difficult := RaceSpecialTrait(value, "drow", l, ability)
				assert.NotEmpty(t, name)
				assert.NotNil(t, spell)
				assert.NotNil(t, spellcast)
				assert.NotEmpty(t, description)
				assert.NotEmpty(t, difficult)
			}
		}
	}
}

func TestCoinList(t *testing.T) {
	test1 := "electrum"
	assert.Contains(t, CoinList(), test1)
}

func TestCoinShortnameList(t *testing.T) {
	test1 := "sp"
	assert.Contains(t, CoinShortnameList(), test1)
}

func TestCoinShortName(t *testing.T) {
	for _, v := range CoinShortnameList() {
		res := CoinShortName(v)
		assert.NotEmpty(t, res)
		assert.Contains(t, CoinList(), res)
	}
	test := CoinShortName("diamond")
	assert.Equal(t, "unknown", test)
}

func TestExchangeRates(t *testing.T) {
	for _, i := range CoinList() {
		for _, o := range CoinList() {
			res, err := ExchangeRates(i, o, 1)
			assert.NotNil(t, res)
			switch i {
			case "platinum":
				assert.NoError(t, err)
			case "gold":
				if o == "platinum" {
					assert.Error(t, err)
				}
			case "electrum":
				if o == "platinum" || o == "gold" {
					assert.Error(t, err)
				}
			case "silver":
				if o == "platinum" || o == "gold" || o == "electrum" {
					assert.Error(t, err)
				}
			case "copper":
				assert.Error(t, err)
			}
		}
	}
}
