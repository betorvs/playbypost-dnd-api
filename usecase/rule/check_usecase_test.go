package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestCheckPreparedSpellsList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test := new(rule.PreparedSpellsList)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 14
	}
	test.Ability = ability
	test.Level = 1
	test.Class = "wizard"
	test.PreparedSpells = []string{"Acid Splash", "Chill Touch", "Sacred Flame", "Acid Splash", "Chill Touch", "Acid Splash", "Chill Touch"}
	res, err := CheckPreparedSpellsList(test)
	assert.NotEmpty(t, res)
	assert.Error(t, err)
	test.PreparedSpells = []string{"Acid Splash", "Chill Touch", "Sacred Flame"}
	res1, err1 := CheckPreparedSpellsList(test)
	assert.NotEmpty(t, res1)
	assert.Error(t, err1)
	test.PreparedSpells = []string{"chill-touch", "chill-touch", "chill-touch"}
	res2, err2 := CheckPreparedSpellsList(test)
	assert.NotEmpty(t, res2)
	assert.NoError(t, err2)
}

func TestCheckCantripsKnownList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test := new(rule.KnownCantripList)
	test.Class = "cleric"
	test.ClassFeatures = ClassFeatures("cleric", 1)
	test.ClassFeatures = append(test.ClassFeatures, "domain-nature")
	test.CantripsKnown = 3
	test.CantripsList = []string{"guidance",
		"light",
		"mending",
		"resistance"}
	res1, err1 := CheckCantripsKnownList(test)
	assert.NotEmpty(t, res1)
	assert.Error(t, err1)
	test.CantripsList = []string{"guidance",
		"light",
		"resistance"}
	res2, err2 := CheckCantripsKnownList(test)
	assert.NotEmpty(t, res2)
	assert.NoError(t, err2)
	test.CantripsList = []string{"guidance",
		"light",
		"shillelagh"}
	res3, err3 := CheckCantripsKnownList(test)
	assert.NotEmpty(t, res3)
	assert.Error(t, err3)
	test2 := new(rule.KnownCantripList)
	test2.Class = "fighter"
	test2.ClassFeatures = ClassFeatures("fighter", 3)
	test2.ClassFeatures = append(test2.ClassFeatures, "archetype-eldritch-knight-spellcasting")
	test2.CantripsKnown = 3
	test2.CantripsList = []string{"chill-touch",
		"chill-touch",
		"chill-touch"}
	res4, err4 := CheckCantripsKnownList(test2)
	assert.NotEmpty(t, res4)
	assert.NoError(t, err4)
	test3 := new(rule.KnownCantripList)
	test3.Class = "rogue"
	test3.ClassFeatures = ClassFeatures("rogue", 3)
	test3.ClassFeatures = append(test3.ClassFeatures, "archetype-arcane-trickster-spellcasting")
	test3.CantripsKnown = 3
	test3.CantripsList = []string{"chill-touch",
		"chill-touch",
		"chill-touch"}
	res5, err5 := CheckCantripsKnownList(test3)
	assert.NotEmpty(t, res5)
	assert.NoError(t, err5)
}

func TestCheckKnownList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	test1 := new(rule.KnownSpellsList)
	test1.Class = "sorcerer"
	test1.Level = 1
	test1.KnownSpells = 2
	test1.SpellMaxLevel = 1
	test1.SpellList = []string{"chill-touch",
		"chill-touch",
		"chill-touch"}
	res1, err1 := CheckKnownList(test1)
	assert.NotEmpty(t, res1)
	assert.Error(t, err1)
	test1.SpellList = []string{"chill-touch",
		"chill-touch"}
	res2, err2 := CheckKnownList(test1)
	assert.NotEmpty(t, res2)
	assert.NoError(t, err2)
	test1.SpellList = []string{"chill-touch",
		"shillelagh"}
	res3, err3 := CheckKnownList(test1)
	assert.NotEmpty(t, res3)
	assert.Error(t, err3)
	test1.SpellList = []string{"chill-touch",
		"alter-self"}
	res4, err4 := CheckKnownList(test1)
	assert.NotEmpty(t, res4)
	assert.Error(t, err4)
}

func TestCheckArmorClass(t *testing.T) {
	test1 := new(rule.ArmorClass)
	ability := make(map[string]int)
	for _, v := range AbilityList() {
		ability[v] = 12
	}
	test1.Ability = ability
	test1.ClassFeatures = ClassFeatures("fighter", 1)
	test1.ClassFeatures = append(test1.ClassFeatures, "fighting-style-defense")
	test1.ArmorProficiency = []string{"light-armor"}
	test1.Armor = "padded"
	test1.Shield = "shield"
	test1.ArmorMagicBonus = 1
	test1.ShieldMagicBonus = 1
	res1 := CheckArmorClass(test1)
	assert.NotEmpty(t, res1)
	test2 := new(rule.ArmorClass)
	test2.Ability = ability
	test2.ClassFeatures = ClassFeatures("monk", 1)
	test2.ArmorProficiency = []string{"light-armor"}
	test2.Armor = ""
	res2 := CheckArmorClass(test2)
	assert.NotEmpty(t, res2)
	test3 := new(rule.ArmorClass)
	test3.Ability = ability
	test3.ClassFeatures = ClassFeatures("barbarian", 1)
	test3.ArmorProficiency = []string{"light-armor"}
	test3.Armor = "plate"
	res3 := CheckArmorClass(test3)
	assert.NotEmpty(t, res3)
	test3.Armor = ""
	res4 := CheckArmorClass(test3)
	assert.NotEmpty(t, res4)
	test5 := new(rule.ArmorClass)
	test5.Ability = ability
	test5.ClassFeatures = ClassFeatures("sorcerer", 1)
	test5.ClassFeatures = append(test5.ClassFeatures, "sorcerous-origin-draconic-resistance")
	test5.ArmorProficiency = []string{"light-armor"}
	test5.Armor = ""
	res5 := CheckArmorClass(test5)
	assert.NotEmpty(t, res5)
}
