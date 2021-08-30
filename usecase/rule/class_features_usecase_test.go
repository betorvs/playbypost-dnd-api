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
	test := map[string]map[string]int{
		"fighting-style": {
			"archery": 1,
		},
		"primal-path": {
			"berseker": 10,
		},
		"bard-college": {
			"lore": 3,
		},
		"divine-domain": {
			"knowledge": 2,
		},
		"druid-circle": {
			"arctic": 2,
		},
		"martial-archetype": {
			"champion": 7,
		},
		"monastic-tradition": {
			"open-hand": 6,
		},
		"sacred-oath": {
			"devotion": 7,
		},
		"ranger-archetype": {
			"colossus-slayer": 3,
		},
		"roguish-archetype": {
			"thief": 3,
		},
		"sorcerous-origin": {
			"black": 0,
		},
		"otherworldly-patron": {
			"archfey": 1,
		},
		"arcane-tradition": {
			"abjuration": 6,
		},
	}
	res := map[string]map[string]string{
		"fighting-style": {
			"archery": "fighting-style-archery",
		},
		"primal-path": {
			"berseker": "path-of-berseker-mindless-rage",
		},
		"bard-college": {
			"lore": "college-of-lore-bonus-proficiencies",
		},
		"divine-domain": {
			"knowledge": "channel-divinity-knowledge-of-the-ages",
		},
		"druid-circle": {
			"arctic": "circle-of-the-land-spells-arctic",
		},
		"martial-archetype": {
			"champion": "archetype-champion-remarkable-athlete",
		},
		"monastic-tradition": {
			"open-hand": "monastic-tradition-way-of-the-open-wholeness-of-body",
		},
		"sacred-oath": {
			"devotion": "sacred-oath-of-devotion-aura-of-devotion",
		},
		"ranger-archetype": {
			"colossus-slayer": "archetype-hunter-hunters-prey-colossus-slayer",
		},
		"roguish-archetype": {
			"thief": "archetype-thief-fast-hands-and-second-storywork",
		},
		"sorcerous-origin": {
			"black": "sorcerous-origin-draconic-bloodline-black-dragon-ancestor",
		},
		"otherworldly-patron": {
			"archfey": "otherworldly-patron-the-archfey-fey-presence",
		},
		"arcane-tradition": {
			"abjuration": "arcane-tradition-school-of-abjuration-projected-ward",
		},
	}
	for k, v := range test {
		for key, value := range v {
			tmp := choosenClassFeatures(k, key, value)
			assert.Contains(t, tmp, res[k][key])
		}
	}

}

func TestFeatureImprovedByLevel(t *testing.T) {
	test1 := "bardic-inspiration"
	res1 := "1d10"
	assert.Contains(t, featureImprovedByLevel(test1, 10), res1)
}

func TestClericDestroyUndead(t *testing.T) {
	res1 := float64(4)
	assert.Equal(t, clericDestroyUndead(17), res1)
}

func TestFeaturesWithExtraSpellList(t *testing.T) {
	res1 := "circle-of-the-land-spells-arctic"
	assert.Contains(t, featuresWithExtraSpellList(), res1)
}

func TestExtraSpellList(t *testing.T) {
	test1 := "domain-knowledge"
	res1 := "suggestion"
	assert.Contains(t, extraSpellList(test1, 3), res1)
}

func TestExtraDamageMeleeAttackFeature(t *testing.T) {
	test1 := "domain-life-divine-strike"
	exp1 := "radiant"
	exp1dice := "2d8"
	res1dice, res1 := extraDamageMeleeAttackFeature(test1, "", 14, 0, false)
	assert.Contains(t, res1, exp1)
	assert.Contains(t, res1dice, exp1dice)
}
