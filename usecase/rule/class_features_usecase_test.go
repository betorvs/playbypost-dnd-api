package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanoidList(t *testing.T) {
	test1 := "bugbear"
	assert.Contains(t, humanoidList(), test1)
}

func TestMonsterTypeList(t *testing.T) {
	test1 := "aberrations"
	assert.Contains(t, monsterTypeList(), test1)
}

func TestFeaturesListRename(t *testing.T) {
	test1 := "fighting-style"
	assert.Contains(t, featuresListRename(), test1)
}

func TestTerrainList(t *testing.T) {
	test1 := "arctic"
	assert.Contains(t, terrainList(), test1)
}

func TestDragonKinds(t *testing.T) {
	test1 := "black"
	assert.Contains(t, dragonKinds(), test1)
}

func TestFightingStyleList(t *testing.T) {
	test1 := "archery"
	assert.Contains(t, fightingStyleList(), test1)
}

func TestFightingStyleFullNameList(t *testing.T) {
	test1 := "fighting-style-defense"
	assert.Contains(t, fightingStyleFullNameList(), test1)
}

func TestBarbarianPrimalPathList(t *testing.T) {
	test1 := "berseker"
	assert.Contains(t, barbarianPrimalPathList(), test1)
}

func TestBardCollege(t *testing.T) {
	test1 := "lore"
	assert.Contains(t, bardCollege(), test1)
}

func TestClericDivineDomainList(t *testing.T) {
	test1 := "knowledge"
	assert.Contains(t, clericDivineDomainList(), test1)
}

func TestDruidCircleList(t *testing.T) {
	test1 := "land"
	assert.Contains(t, druidCircleList(), test1)
}

func TestFighterArchetypeList(t *testing.T) {
	test1 := "champion"
	assert.Contains(t, fighterArchetypeList(), test1)
}

func TestMonkMonasticTraditionList(t *testing.T) {
	test1 := "open-hand"
	assert.Contains(t, monkMonasticTraditionList(), test1)
}

func TestPaladinSacredOaths(t *testing.T) {
	test1 := "devotion"
	assert.Contains(t, paladinSacredOaths(), test1)
}

func TestRangerArchetypesList(t *testing.T) {
	test1 := "hunter"
	assert.Contains(t, rangerArchetypesList(), test1)
}

func TestRogueArchetypeList(t *testing.T) {
	test1 := "thief"
	assert.Contains(t, rogueArchetypeList(), test1)
}

func TestSorcererOriginList(t *testing.T) {
	test1 := "draconic-bloodline"
	assert.Contains(t, sorcererOriginList(), test1)
}

func TestWarlockOtherworldlyPatronList(t *testing.T) {
	test1 := "archfey"
	assert.Contains(t, warlockOtherworldlyPatronList(), test1)
}

func TestWarlockPactBoonList(t *testing.T) {
	test1 := "chain"
	assert.Contains(t, warlockPactBoonList(), test1)
}

func TestWizardArcaneTraditionList(t *testing.T) {
	test1 := "abjuration"
	assert.Contains(t, wizardArcaneTraditionList(), test1)
}

func TestHuntersPreyList(t *testing.T) {
	test1 := "colossus-slayer"
	assert.Contains(t, huntersPreyList(), test1)
}

func TestDefensiveTacticsList(t *testing.T) {
	test1 := "escape-the-horde"
	assert.Contains(t, defensiveTacticsList(), test1)
}

func TestMultiattackList(t *testing.T) {
	test1 := "volley"
	assert.Contains(t, multiattackList(), test1)
}

func TestHuntersDefenseList(t *testing.T) {
	test1 := "evasion"
	assert.Contains(t, huntersDefenseList(), test1)
}

func TestTurnMonstersFeatureList(t *testing.T) {
	test1 := "sacred-oath-of-devotion-channel-divinity"
	assert.Contains(t, turnMonstersFeatureList(), test1)
}

func TestChoosenClassFeatures(t *testing.T) {
	test := map[string][]string{
		"fighting-style": {
			"archery", "defense", "dueling", "great-weapon-fighting", "protection", "two-weapon-fighting",
		},
		"primal-path": {
			"berseker", "totem",
		},
		"bard-college": {
			"lore", "valor",
		},
		"divine-domain": {
			"knowledge", "life", "light", "nature", "tempest", "trickery", "war",
		},
		"druid-circle": {
			"arctic", "land", "moon",
		},
		"martial-archetype": {
			"champion", "battle-master", "eldritch-knight",
		},
		"monastic-tradition": {
			"open-hand", "shadow", "four-elements",
		},
		"sacred-oath": {
			"devotion", "ancients", "vengeance",
		},
		"ranger-archetype": {
			"hunter", "beast-master", "colossus-slayer", "escape-the-horde", "volley", "evasion",
		},
		"natural-explorer": {
			"arctic", "coast", "desert", "forest", "grassland", "mountain", "swamp", "underdark",
		},
		"favored-enemy": {
			"aberrations", "beasts", "celestials", "constructs", "dragons", "elementals", "fey", "fiends", "giants", "monstrosities", "oozes", "plants", "undead", "bugbear", "gnoll", "goblin", "grimlock", "hobgoblin", "kobold", "lizardfolk", "merfolk", "orc", "sahuagin", "thug", "werebear", "wereboar", "wererat", "weretiger", "werewolf", "gnome", "human", "elf", "dwarf",
		},
		"roguish-archetype": {
			"thief", "assassin", "arcane-trickster",
		},
		"sorcerous-origin": {
			"black", "blue", "brass", "bronze", "copper", "gold", "green", "red", "silver", "white", "draconic", "wild",
		},
		"otherworldly-patron": {
			"archfey", "fiend", "old-one",
		},
		"pact-boon": {
			"chain", "blade", "tome",
		},
		"arcane-tradition": {
			"abjuration", "conjuration", "divination", "enchantment", "evocation", "illusion", "necromancy", "transmutation",
		},
	}
	levels := map[string][]int{
		"fighting-style": {
			1,
		},
		"primal-path": {
			3, 6, 10, 14,
		},
		"bard-college": {
			3, 6, 14,
		},
		"divine-domain": {
			1, 2, 6, 8, 17,
		},
		"druid-circle": {
			2, 6, 10, 14,
		},
		"martial-archetype": {
			3, 7, 10, 15, 18,
		},
		"monastic-tradition": {
			3, 6, 11, 17,
		},
		"sacred-oath": {
			3, 7, 15, 20,
		},
		"ranger-archetype": {
			3, 7, 11, 15,
		},
		"natural-explorer": {
			1,
		},
		"favored-enemy": {
			1,
		},
		"roguish-archetype": {
			3, 9, 13, 17,
		},
		"sorcerous-origin": {
			1, 6, 14, 18,
		},
		"otherworldly-patron": {
			1, 6, 10, 14,
		},
		"pact-boon": {
			1,
		},
		"arcane-tradition": {
			2, 6, 10, 14,
		},
	}
	for k, v := range test {
		for _, value := range v {
			for _, level := range levels[k] {
				tmp := choosenClassFeatures(k, value, level)
				assert.GreaterOrEqual(t, len(tmp), 1)
			}
		}
	}
	test2 := choosenClassFeatures("", "", 1)
	assert.Empty(t, test2)

}

func TestFeatureImprovedByLevel(t *testing.T) {
	test1 := "bardic-inspiration"
	res1 := "1d10"
	assert.Contains(t, featureImprovedByLevel(test1, 11), res1)
	res2 := "1d6"
	assert.Contains(t, featureImprovedByLevel(test1, 4), res2)
	res3 := "1d12"
	assert.Contains(t, featureImprovedByLevel(test1, 15), res3)
	res4 := "1d8"
	assert.Contains(t, featureImprovedByLevel(test1, 9), res4)

	test2 := "song-of-rest"
	assert.Contains(t, featureImprovedByLevel(test2, 14), res1)
	assert.Contains(t, featureImprovedByLevel(test2, 4), res2)
	assert.Contains(t, featureImprovedByLevel(test2, 18), res3)
	assert.Contains(t, featureImprovedByLevel(test2, 9), res4)
}

func TestClericDestroyUndead(t *testing.T) {
	res1 := float64(4)
	assert.Equal(t, clericDestroyUndead(17), res1)
	res2 := float64(0.5)
	assert.Equal(t, clericDestroyUndead(1), res2)
}

func TestFeaturesWithExtraSpellList(t *testing.T) {
	res1 := "circle-of-the-land-spells-arctic"
	assert.Contains(t, featuresWithExtraSpellList(), res1)
}

func TestExtraSpellList(t *testing.T) {
	test1 := []string{"domain-knowledge", "domain-life", "domain-light", "domain-nature", "domain-tempest", "domain-trickery", "domain-war"}
	levelsDomains := []int{1, 3, 5, 7, 9}
	for _, v := range test1 {
		for _, value := range levelsDomains {
			tmp := extraSpellList(v, value)
			assert.GreaterOrEqual(t, len(tmp), 1)
		}
	}
	test2 := []string{"sacred-oath-of-devotion", "sacred-oath-of-ancients", "sacred-oath-of-vengeance"}
	levelsSacred := []int{3, 5, 9, 13, 17}
	for _, v := range test2 {
		for _, value := range levelsSacred {
			tmp := extraSpellList(v, value)
			assert.GreaterOrEqual(t, len(tmp), 1)
		}
	}
	test3 := []string{"circle-of-the-land-spells-arctic", "circle-of-the-land-spells-coast", "circle-of-the-land-spells-desert", "circle-of-the-land-spells-forest", "circle-of-the-land-spells-grassland", "circle-of-the-land-spells-mountain", "circle-of-the-land-spells-swamp", "circle-of-the-land-spells-underdark"}
	levelsCircle := []int{3, 5, 7, 9}
	for _, v := range test3 {
		for _, value := range levelsCircle {
			tmp := extraSpellList(v, value)
			assert.GreaterOrEqual(t, len(tmp), 1)
		}
	}
	// assert.Contains(t, extraSpellList(test1, 3), res1)
	// assert.GreaterOrEqual(t, len(tmp), 1)
}

func TestExtraDamageMeleeAttackFeature(t *testing.T) {
	test1 := []string{"domain-life-divine-strike", "sneak-attack", "archetype-hunter-hunters-prey-colossus-slayer", "domain-life-divine-strike", "domain-nature-divine-strike", "domain-tempest-divine-strike", "domain-trickery-divine-strike", "domain-war-divine-strike", "divine-smite"}
	for _, v := range test1 {
		chosen := "piercing"
		if v == "domain-nature-divine-strike" {
			chosen = "cold"
		}
		resDice, resType := extraDamageMeleeAttackFeature(v, chosen, 14, 0, false)
		assert.NotEmpty(t, resType)
		assert.Contains(t, resDice, "d")
		if v == "divine-smite" {
			res2Dice, res2Type := extraDamageMeleeAttackFeature(v, chosen, 14, 6, false)
			assert.NotEmpty(t, res2Type)
			assert.Contains(t, res2Dice, "d")
			res3Dice, res3Type := extraDamageMeleeAttackFeature(v, chosen, 14, 0, true)
			assert.NotEmpty(t, res3Type)
			assert.Contains(t, res3Dice, "d")
		}
	}
	test2Dice, test2Type := extraDamageMeleeAttackFeature("", "", 1, 0, false)
	assert.Empty(t, test2Dice)
	assert.Empty(t, test2Type)

}
