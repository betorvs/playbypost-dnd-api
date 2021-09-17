package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//CheckPreparedSpellsList func
func CheckPreparedSpellsList(spellList *rule.PreparedSpellsList) (*rule.PreparedSpellsList, error) {
	ability := AbilityForSpell(spellList.Class)
	abilityModifier := CalcAbilityModifier(spellList.Ability[ability])
	numberPreparedSpell := spellList.Level + abilityModifier
	if len(spellList.PreparedSpells) != numberPreparedSpell {
		err := fmt.Errorf("invalid number of prepared spells")
		return spellList, err
	}
	for _, s := range spellList.PreparedSpells {
		if !utils.StringInSlice(s, GetFullSpellList(spellList.Class)) {
			err := fmt.Errorf("prepared spell %s is not in spell list for %s", s, spellList.Class)
			return spellList, err
		}
	}
	spellList.Verified = true
	return spellList, nil
}

//CheckCantripsKnownList func
func CheckCantripsKnownList(spellList *rule.KnownCantripList) (*rule.KnownCantripList, error) {
	numberCantrips := spellList.CantripsKnown
	if len(spellList.CantripsList) != numberCantrips {
		err := fmt.Errorf("invalid number of cantrips")
		return spellList, err
	}
	class := spellList.Class
	if spellList.Class == "fighter" || spellList.Class == "rogue" {
		if utils.StringInSlice("archetype-eldritch-knight-spellcasting", spellList.ClassFeatures) {
			class = "eldritch-knight"
		}
		if utils.StringInSlice("archetype-arcane-trickster-spellcasting", spellList.ClassFeatures) {
			class = "arcane-trickster"
		}
	}
	cantrips := GetSpellListByClass(class, 0)
	if utils.StringInSlice("domain-nature", spellList.ClassFeatures) {
		druidCantrips := GetSpellListByClass(class, 0)
		cantrips.List = append(cantrips.List, druidCantrips.List...)
	}

	for _, s := range spellList.CantripsList {
		if !utils.StringInSlice(s, cantrips.List) {
			err := fmt.Errorf("cantrip %s is not in your cantrip list for %s", s, class)
			return spellList, err
		}
	}
	spellList.Verified = true
	return spellList, nil
}

//CheckKnownList func
func CheckKnownList(spellList *rule.KnownSpellsList) (*rule.KnownSpellsList, error) {
	numberSpell := spellList.KnownSpells
	if len(spellList.SpellList) != numberSpell {
		err := fmt.Errorf("invalid number of known spells")
		return spellList, err
	}
	class := spellList.Class
	spells := GetFullSpellList(class)

	for _, s := range spellList.SpellList {
		if !utils.StringInSlice(s, spells) {
			err := fmt.Errorf("spell added %s is not in your spell list for %s", s, class)
			return spellList, err
		}
		spell := getSpellByName(s)
		if spell.Level > spellList.SpellMaxLevel {
			err := fmt.Errorf("spell added %s is higher than your spell maximum of %v level of spell for %s", s, spellList.SpellMaxLevel, class)
			return spellList, err
		}
	}
	spellList.Verified = true
	return spellList, nil
}

//CheckArmorClass func
func CheckArmorClass(ac *rule.ArmorClass) *rule.ReturnACMessage {
	// var disvantage string
	result := new(rule.ReturnACMessage)
	result.ArmorClassDisvantages = []string{}
	result.ArmorClassAutomaticallyFails = []string{}
	result.ArmorClassSpeedReduced = false
	result.ArmorClassStealthDisvantage = false
	// armorType, _, armorClassInitial, dexterityMax, strengthMin, stealth, _ := ArmorByName(ac.Armor)
	armor := ArmorByName(ac.Armor)
	armorClassInitial := armor.ArmorClass
	if armor.ArmorClass == 0 {
		armorClassInitial = 10
	}
	if !utils.StringInSlice(armor.Kind, ac.ArmorProficiency) && ac.Armor != "" {
		result.ArmorClassDisvantages = []string{"strength", "dexterity"}
		result.ArmorClassAutomaticallyFails = []string{"spellcast"}
	}
	if ac.Ability["strength"] < armor.Strength {
		result.ArmorClassSpeedReduced = true
	}
	var unarmoredDefense int

	if utils.StringInSlice("sorcerous-origin-draconic-resistance", ac.ClassFeatures) && ac.Armor == "" {
		armorClassInitial = 13
	}
	if utils.StringInSlice("unarmored-defense-monk", ac.ClassFeatures) && ac.Armor == "" {
		unarmoredDefense = CalcAbilityModifier(ac.Ability["wisdom"])
	}

	if utils.StringInSlice("unarmored-defense-barbarian", ac.ClassFeatures) && ac.Armor == "" {
		unarmoredDefense = CalcAbilityModifier(ac.Ability["constitution"])
	}
	var fightingStyle int
	if utils.StringInSlice("fighting-style-defense", ac.ClassFeatures) {
		fightingStyle = 1
	}

	abilityModifier := CalcAbilityModifier(ac.Ability["dexterity"])
	dexterityModifier := abilityModifier
	if ac.Armor != "" && abilityModifier > armor.DexterityModifier {
		dexterityModifier = armor.DexterityModifier
	}
	var shield int
	if ac.Shield != "" {
		shield = 2

	}
	var magic int
	if ac.ArmorMagicBonus != 0 {
		magic = ac.ArmorMagicBonus
		if ac.ShieldMagicBonus != 0 {
			magic = ac.ArmorMagicBonus + ac.ShieldMagicBonus
		}
	}
	armorClass := armorClassInitial + dexterityModifier + shield + magic + unarmoredDefense + fightingStyle
	fmt.Println(armorClassInitial, dexterityModifier, shield, magic, unarmoredDefense, fightingStyle)
	// if utils.StringInSlice("unarmored-defense-monk", ac.ClassFeatures) {
	// 	armorClass = armorClassInitial + dexterityModifier + magic + unarmoredDefense + fightingStyle
	// }
	// if utils.StringInSlice("unarmored-defense-barbarian", ac.ClassFeatures) {
	// 	armorClass = armorClassInitial + dexterityModifier + magic + unarmoredDefense + fightingStyle
	// }
	if armor.Stealth {
		result.ArmorClassDisvantages = append(result.ArmorClassDisvantages, "stealth")
		result.ArmorClassStealthDisvantage = true
	}
	result.ArmorClassMaxDexterity = dexterityModifier
	result.ArmorClass = armorClass
	return result
}
