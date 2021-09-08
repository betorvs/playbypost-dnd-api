package rule

import (
	"fmt"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

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
