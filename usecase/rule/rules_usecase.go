package rule

import (
	"fmt"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//ListInformation func
func ListInformation(name, value string) (*rule.SimpleList, error) {
	list := new(rule.SimpleList)
	switch name {
	case "race", "races":
		list.List = RaceList()
		if value == "subrace" {
			list.List = RaceListWithSubrace()
		}

	case "class", "classes":
		list.List = ClassList()

	case "background", "backgrounds":
		list.List = BackgroundList()

	case "ability":
		list.List = AbilityList()

	case "alignment":
		list.List = AlignmentList()

	case "skills", "skill":
		list.List = SkillListAbility()

	case "subrace", "subraces":
		list.List = SubraceList(value)

	case "condition", "conditions":
		list.List = ListConditions()

	case "damage", "damagetype":
		list.List = ListOfDamageTypes()
	}
	return list, nil
}

//FullDescription func
func FullDescription(kind, name, subname string) (*rule.FullDescription, error) {
	fullDescription := new(rule.FullDescription)
	switch kind {
	case "race":
		fullDescription.Description = RaceTraits(name, subname)
	case "class":
		fullDescription.Description = ClassInfo(name)

	case "background":
		m, _ := BackgroundStatistics(name)
		fullDescription.Description = m

	case "condition", "conditions":
		d := conditionsMap(name)
		fullDescription.Description = d

	case "damage", "damagetype":
		d := damageTypeMap(name)
		fullDescription.Description = d
	}

	return fullDescription, nil
}

//GetConditions func
func GetConditions(condition string, level int) *rule.ReturnCondition {
	message := new(rule.ReturnCondition)
	desc, dis, auto := conditionsCheck(condition, level)
	message.Name = condition
	message.Description = desc
	message.Disvantages = dis
	message.AutoFail = auto
	return message
}

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

//CalcArmorClass func
func CalcArmorClass(ac *rule.ArmorClass) *rule.ReturnACMessage {
	// var disvantage string
	result := new(rule.ReturnACMessage)
	result.ArmorClassDisvantages = []string{}
	result.ArmorClassAutomaticallyFails = []string{}
	result.ArmorClassSpeedReduced = false
	result.ArmorClassStealthDisvantage = false
	// armorType, _, armorClassInitial, dexterityMax, strengthMin, stealth, _ := ArmorByName(ac.Armor)
	armor := ArmorByName(ac.Armor)
	armorClassInitial := armor.ArmorClass
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
	if utils.StringInSlice("unarmored-defense-monk", ac.ClassFeatures) {
		unarmoredDefense = CalcAbilityModifier(ac.Ability["wisdom"])
	}

	if utils.StringInSlice("unarmored-defense-barbarian", ac.ClassFeatures) {
		unarmoredDefense = CalcAbilityModifier(ac.Ability["constitution"])
	}
	var fightingStyle int
	if utils.StringInSlice("fighting-style-defense", ac.ClassFeatures) {
		fightingStyle = 1
	}

	abilityModifier := CalcAbilityModifier(ac.Ability["dexterity"])
	dexterityModifier := abilityModifier
	if abilityModifier > armor.DexterityModifier {
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
	if utils.StringInSlice("unarmored-defense-monk", ac.ClassFeatures) {
		armorClass = armorClassInitial + dexterityModifier + magic + unarmoredDefense + fightingStyle
	}
	if armor.Stealth {
		result.ArmorClassDisvantages = append(result.ArmorClassDisvantages, "stealth")
		result.ArmorClassStealthDisvantage = true
	}
	result.ArmorClassMaxDexterity = dexterityModifier
	result.ArmorClass = armorClass
	return result
}

//ClassFeatureRoll func
func ClassFeatureRoll(feature *rule.Feature) (*rule.ReturnCalcMessage, error) {
	returnResult := new(rule.ReturnCalcMessage)
	returnResult.Success = false
	r := diceroll.GetDice()
	if feature.Name == "song-of-rest" && utils.StringInSlice("song-of-rest", feature.ClassFeatures) {
		if feature.Level == 1 {
			err := fmt.Errorf("bard level invalid")
			return returnResult, err
		}
		dice := featureImprovedByLevel("song-of-rest", feature.Level)
		res, text, _ := r.DiceRoll(dice)
		returnResult.Success = true
		returnResult.HealingValue = res
		returnResult.RolledMessage = fmt.Sprintf("Rolled %s in %s", text, dice)
		returnResult.Message = fmt.Sprintf("After Performing %s you fell recovered", feature.Name)
	}
	if feature.Name == "bardic-inspiration" && utils.StringInSlice("bardic-inspiration", feature.ClassFeatures) {
		dice := featureImprovedByLevel("bardic-inspiration", feature.Level)
		res, text, _ := r.DiceRoll(dice)
		returnResult.Success = true
		returnResult.RolledValue = res
		returnResult.RolledMessage = fmt.Sprintf("Rolled %s in %s", text, dice)
		returnResult.Message = fmt.Sprintf("You can add %v in your test because of %s", res, feature.Name)
	}

	if feature.Name == "second-wind" && utils.StringInSlice("second-wind", feature.ClassFeatures) {
		dice := "1d10"
		res, text, _ := r.DiceRoll(dice)
		result := res + feature.Level
		returnResult.Success = true
		returnResult.HealingValue = result
		returnResult.RolledMessage = fmt.Sprintf("Rolled %s in %s", text, dice)
		returnResult.Message = fmt.Sprintf("%s makes you regain %v hit points", feature.Name, result)
	}

	if feature.Name == "lay-on-hands" && utils.StringInSlice("lay-on-hands", feature.ClassFeatures) {
		result := 5 * feature.Level
		returnResult.HealingValue = result
		returnResult.Success = true
		returnResult.Message = fmt.Sprintf("%s you can restore a %v total number of hit points or you can expend 5 hit points from your pool of healing to cure the target of one disease or neutralize one poison affecting it", feature.Name, result)
	}

	if feature.Name == "bless" || feature.Name == "guidance" || feature.Name == "resistance" {
		spell := getSpellByName(feature.Name)
		res, text, _ := r.DiceRoll(spell.ExtraDice)
		returnResult.Success = true
		returnResult.RolledValue = res
		returnResult.RolledMessage = fmt.Sprintf("Rolled %s in %s", text, spell.ExtraDice)
		returnResult.Message = fmt.Sprintf("You can add %v in your test of %v because of %s", res, spell.ExtraDiceUsage, feature.Name)
	}

	if feature.Name == "combat-wild-shape" && feature.UsingFeatureSlot != 0 && utils.StringInSlice("circle-of-the-moon-combat-wild-shape", feature.ClassFeatures) {
		dice := fmt.Sprintf("%vd8", feature.UsingFeatureSlot)
		res, text, _ := r.DiceRoll(dice)
		returnResult.Success = true
		returnResult.RolledValue = res
		returnResult.RolledMessage = fmt.Sprintf("Rolled %s in %s", text, dice)
		returnResult.Message = fmt.Sprintf("You can heal %v in your hit points of %v because of %s", res, dice, feature.Name)
	}

	if feature.Name == "combat-superiority" && utils.StringInSlice("archetype-battle-master-combat-superiority", feature.ClassFeatures) {
		dice := "1d8"
		res, text, _ := r.DiceRoll(dice)
		returnResult.Success = true
		returnResult.RolledValue = res
		returnResult.RolledMessage = fmt.Sprintf("Rolled %s in %s", text, dice)
		returnResult.Message = fmt.Sprintf("You can add %v in your test of %v because of %s", res, dice, feature.Name)
	}

	// "channel-divinity-radiance-of-the-dawn"
	// damage radiant 2d10 + cleric level to all creatures saving constitution on save half damage
	if feature.Name == "radiance-of-the-dawn" && utils.StringInSlice("channel-divinity-radiance-of-the-dawn", feature.ClassFeatures) {
		dice := "2d10"
		damageType := "radiant"
		res, text, _ := r.DiceRoll(dice)
		result := res + feature.Level
		prof := CalcProficiency(feature.Level)
		abilityModifier := CalcAbilityModifier(feature.Ability["wisdom"])
		abilitySaving := "constitution"
		difficult := prof + 8 + abilityModifier
		if len(feature.MonsterList) != 0 {
			for _, m := range feature.MonsterList {
				tmpMessage := fmt.Sprintf("Your %s Feature have saving with %s and DC: %v", feature.Name, abilitySaving, difficult)
				rollMonster := savingRollMonster(m, "spell", abilitySaving, damageType, tmpMessage, difficult, result, feature.Level, []string{})
				returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
			}
		}
		returnResult.RolledMessage = fmt.Sprintf("Rolled: %s", text)
	}

	if feature.Name == "grim-harvest" && utils.StringInSlice("arcane-tradition-school-of-necromancy-grim-harvest", feature.ClassFeatures) {
		// if necromancy spell 3 x
		// if not spell 2 x
		var recover int
		for _, s := range feature.GenericList {
			spell := getSpellByName(s)
			if !strings.Contains(spell.Title, "necromancy") {
				recover = spell.Level * 2
			}
			if strings.Contains(spell.Title, "necromancy") {
				recover = spell.Level * 3
			}

		}
		returnResult.HealingValue = recover
		returnResult.Success = true
		returnResult.Message = fmt.Sprintf("%s you can restore a %v total number of hit points", feature.Name, recover)

	}

	// if dont find any feature
	if !returnResult.Success {
		returnResult.Message = fmt.Sprintf("Cannot find any feature with name %s", feature.Name)
	}
	return returnResult, nil
}

//SpecialRaceFeature func
func SpecialRaceFeature(feature *rule.SpecialRaceFeature) (*rule.ReturnCalcMessage, error) {
	returnResult := new(rule.ReturnCalcMessage)
	var abilitySaving string
	var damageDice string
	var damageType string
	var difficult int
	name, _, _, RacialDamageDice, RacialDamageType, savingThrow, _, difficultClass := RaceSpecialTrait(feature.Race, feature.Subrace, feature.Level, feature.Ability)
	if name == "" {
		err := fmt.Errorf("special Race feature %s its not listed in your race %s", feature.Name, feature.Race)
		return returnResult, err
	}
	damageDice = RacialDamageDice
	damageType = RacialDamageType
	abilitySaving = savingThrow
	difficult = difficultClass
	returnResult.Message = fmt.Sprintf("Racial Trait %s ", name)
	// racialTraitMessageSuffix = fmt.Sprintf(" Saving with %s and DC: %v", savingThrow, difficultClass)
	if feature.Name == "hellish-rebuke" && feature.Race == "tiefling" {
		spell := getSpellByName(feature.Name)
		if spell.DamageDice != "" {
			damageDice = spell.DamageDice
		}
		if spell.DamageType != "" {
			damageType = spell.DamageType
		}
		if spell.SavingThrow != "" {
			abilitySaving = spell.SavingThrow
		}
		// returnResult.Message = fmt.Sprintf("Racial Trait %s ", feature.Name)
		// racialTraitMessageSuffix = fmt.Sprintf(" Saving with %s and DC: %v", spellSaving, difficultClass)

		difficult = difficultClass
	}
	r := diceroll.GetDice()
	if damageDice != "" {
		resultDamage, text, _ := r.DiceRoll(damageDice)
		if len(feature.Monster) != 0 {
			for _, m := range feature.Monster {
				tmpMessage := fmt.Sprintf("Your Race Feature have saving with %s and DC: %v", abilitySaving, difficultClass)
				rollMonster := savingRollMonster(m, "spell", abilitySaving, damageType, tmpMessage, difficult, resultDamage, feature.Level, []string{})
				returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
			}
		}
		returnResult.RolledMessage = fmt.Sprintf("Rolled: %s", text)
	}

	return returnResult, nil
}
