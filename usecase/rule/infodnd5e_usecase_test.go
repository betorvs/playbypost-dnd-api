package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/stretchr/testify/assert"
)

func TestAbilityList(t *testing.T) {
	test1 := "strength"
	assert.Contains(t, AbilityList(), test1)
}

func TestRaceList(t *testing.T) {
	test1 := "human"
	assert.Contains(t, RaceList(), test1)
}

func TestRaceListWithSubrace(t *testing.T) {
	test1 := "dwarf"
	assert.Contains(t, RaceListWithSubrace(), test1)
}

func TestSubraceList(t *testing.T) {
	test1 := []string{"dwarf", "elf", "halfling", "gnome", ""}
	for _, v := range test1 {
		res := SubraceList(v)
		assert.GreaterOrEqual(t, len(res), 1)
	}
}

func TestBackgroundList(t *testing.T) {
	test1 := "acolyte"
	assert.Contains(t, BackgroundList(), test1)
}

func TestClassList(t *testing.T) {
	test1 := "barbarian"
	assert.Contains(t, ClassList(), test1)
}

func TestAlignmentList(t *testing.T) {
	test1 := "lawful-good"
	assert.Contains(t, AlignmentList(), test1)
}

func TestSkillList(t *testing.T) {
	test1 := "acrobatics"
	assert.Contains(t, SkillList(), test1)
}

func TestSkillListAbility(t *testing.T) {
	test1 := "acrobatics (dex)"
	assert.Contains(t, SkillListAbility(), test1)
}

func TestProficiencyFullList(t *testing.T) {
	test1 := "light-armor"
	assert.Contains(t, ProficiencyFullList(), test1)
}

func TestLanguageList(t *testing.T) {
	test1 := "common"
	assert.Contains(t, LanguageList(), test1)
}

func TestBarbarianSkillList(t *testing.T) {
	test1 := "animal-handling"
	assert.Contains(t, barbarianSkillList(), test1)
}

func TestBardSkillList(t *testing.T) {
	test1 := "acrobatics"
	assert.Contains(t, bardSkillList(), test1)
}

func TestClericSkillList(t *testing.T) {
	test1 := "history"
	assert.Contains(t, clericSkillList(), test1)
}

func TestDruidSkillList(t *testing.T) {
	test1 := "animal-handling"
	assert.Contains(t, druidSkillList(), test1)
}

func TestFighterSkillList(t *testing.T) {
	test1 := "acrobatics"
	assert.Contains(t, fighterSkillList(), test1)
}

func TestMonkSkillList(t *testing.T) {
	test1 := "acrobatics"
	assert.Contains(t, monkSkillList(), test1)
}

func TestPaladinSkillList(t *testing.T) {
	test1 := "athletics"
	assert.Contains(t, paladinSkillList(), test1)
}

func TestRangerSkillList(t *testing.T) {
	test1 := "athletics"
	assert.Contains(t, rangerSkillList(), test1)
}

func TestRogueSkillList(t *testing.T) {
	test1 := "athletics"
	assert.Contains(t, rogueSkillList(), test1)
}

func TestSorcererSkillList(t *testing.T) {
	test1 := "arcana"
	assert.Contains(t, sorcererSkillList(), test1)
}

func TestWarlockSkillList(t *testing.T) {
	test1 := "arcana"
	assert.Contains(t, warlockSkillList(), test1)
}

func TestWizardSkillList(t *testing.T) {
	test1 := "arcana"
	assert.Contains(t, wizardSkillList(), test1)
}

func TestHealSpellList(t *testing.T) {
	test1 := "cure-wounds"
	assert.Contains(t, healSpellList(), test1)
}

func TestClassWithSpellKnown(t *testing.T) {
	test1 := "bard"
	assert.Contains(t, ClassWithSpellKnown(), test1)
}

func TestClassWithPreparedSpell(t *testing.T) {
	test1 := "cleric"
	assert.Contains(t, ClassWithPreparedSpell(), test1)
}

func TestClassWithSpell(t *testing.T) {
	test1 := "cleric"
	assert.Contains(t, ClassWithSpell(), test1)
}

func TestClassWithCantrips(t *testing.T) {
	test1 := "bard"
	assert.Contains(t, ClassWithCantrips(), test1)
}

func TestAbilityForSpell(t *testing.T) {
	test1 := []string{"bard", "ranger", "sorcerer", "cleric", "warlock", "wizard", "druid", "paladin"}
	for _, v := range test1 {
		res := AbilityForSpell(v)
		assert.True(t, utils.StringInSlice(res, AbilityList()))
		assert.NotEmpty(t, res)
	}
	res2 := AbilityForSpell("")
	assert.Contains(t, res2, "not found")
}

func TestAbilitySkill(t *testing.T) {
	test1 := SkillList()
	for _, v := range test1 {
		res := AbilityForSpell(v)
		assert.True(t, utils.StringInSlice(res, AbilityList()))
		assert.NotEmpty(t, res)
	}
	res2 := AbilityForSpell("")
	assert.Contains(t, res2, "ask to master")
}

func TestRaceTraits(t *testing.T) {}

func TestSubraceTraits(t *testing.T) {}

func TestClassInfo(t *testing.T) {}
