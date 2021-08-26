package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//CalcSpellcastAttackAndSave func
func CalcSpellcastAttackAndSave(spellcast *rule.SpellcastAbility) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)
	if utils.StringInSlice("spellcast", spellcast.AutoFail) || utils.StringInSlice("all", spellcast.AutoFail) || spellcast.Rage {
		returnResult.Message = "Automatically Faills"
		returnResult.RolledValue = 0
		returnResult.Success = false
		return returnResult
	}
	var message string
	// var successMessage string
	damageDice := spellcast.SpellDamage
	damageType := spellcast.SpellDamageType
	spellSaving := spellcast.SpellSaving
	var spellAttack bool
	// var limitRange string
	if spellcast.SpellName != "" && !spellcast.RacialTrait {
		spell := getSpellByName(spellcast.SpellName)
		listByClass := getFullListWithFeature(spellcast.Class, spellcast.ClassFeatures, spellcast.Level)
		// listByClass = checkForExtraSpellLists(listByClass, spellcast.ClassFeatures, spellcast.Level)
		if !utils.StringInSlice(spell.Name, listByClass) {
			returnResult.Message = "Automatically Faills: This Spell is not in your list"
			returnResult.RolledValue = 0
			return returnResult
		}
		listSpellSlots := SpellsPerLevel(spellcast.Class, spellcast.Level)
		// fmt.Println(len(listSpellSlots))
		if len(listSpellSlots) < spell.Level {
			returnResult.Message = "Automatically Faills: You cannot spell with these level"
			returnResult.RolledValue = 0
			return returnResult
		}
		message = fmt.Sprintf("Casting Spell %s with level %v", spell.Name, spell.Level)

		if utils.StringInSlice(spell.Name, healSpellList()) {
			return spellForHealing(spell.Name, spellcast.Class, spell.HealDice, spell.HealingIncreases, spell.AtHigherLevels, spellcast.ClassFeatures, spellcast.Ability, spellcast.Level, spellcast.SpellLevel)
		}
		if spell.DamageDice != "" {
			damageDice = spell.DamageDice
			// fmt.Println(spellcast.SpellLevel, len(listSpellSlots)-1)
			if spellcast.SpellLevel != 0 {
				if spellcast.SpellLevel > len(listSpellSlots)-1 {
					returnResult.Message = fmt.Sprintf("Automatically Faills: You cannot spell with these level changed. Maximum level %v", len(listSpellSlots)-1)
					returnResult.RolledValue = 0
					return returnResult
				}
				damageDice = spellDamageIncrease(spell, spellcast.SpellLevel)
				message = fmt.Sprintf("Casting Spell %s with level %v", spell.Name, spellcast.SpellLevel)
			}
			if spell.DamageIncrease != "" {
				damageDice = spellDamageIncrease(spell, spellcast.Level)
			}
		}
		if spell.DamageType != "" {
			damageType = spell.DamageType
		}
		if spell.SavingThrow != "" {
			spellSaving = spell.SavingThrow
		}
		// limitRange = spell.Range
		if spell.AttackRolls {
			spellAttack = true
		}
		if spell.Range == "Touch" {
			spellAttack = true
		}
	}
	ability := AbilityForSpell(spellcast.Class)
	prof := CalcProficiency(spellcast.Level)
	abilityModifier := CalcAbilityModifier(spellcast.Ability[ability])
	attack := prof + abilityModifier
	save := prof + 8 + abilityModifier
	// d20 := "1d20"
	var adv, dis bool
	if utils.StringInSlice(ability, spellcast.Disvantages) {
		// d20 = "2d20k-1"
		dis = true
	}
	if utils.StringInSlice(ability, spellcast.Advantages) {
		// d20 = "2d20k1"
		adv = true
	}
	d20 := utils.GetD20ToRoll(adv, dis)
	r := diceroll.GetDice()
	if damageDice != "" {
		final, textDamage, _ := r.DiceRoll(damageDice)
		var magic int
		if spellcast.MagicBonus != 0 {
			magic = spellcast.MagicBonus
		}

		resultDamage := final
		if utils.StringInSlice("channel-divinity-destructive-wrath", spellcast.ClassFeatures) {
			if damageType == "lightning" || damageType == "thunder" {
				resultDamage = spellDamageMax(damageDice)
				// fmt.Println(damageDice, resultDamage)
			}
		}
		// var monsterSavingResult int
		if len(spellcast.Monster) != 0 {
			// if spellcast.RacialTrait {
			// 	save = racialDifficult
			// }

			for i, m := range spellcast.Monster {
				// if limitRange == "Touch" && i == 0 {
				// 	text, res := rolledDicesInternal(d20)
				// 	result := res + attack + magic
				// 	tempMessage := fmt.Sprintf("Spell Attack Total: %v using %s with bonus: +%v", result, ability, attack)
				// 	rolledMessage := fmt.Sprintf("Rolled %s in 1d20, Rolled Damage: %s in %s", text, textDamage, damageDice)
				// 	rollMonster := touchRollMonster(m, damageType, rolledMessage, tempMessage, result, resultDamage)
				// 	returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
				// }
				if spellAttack && i == 0 {
					res, text, _ := r.DiceRoll(d20)
					result := res + attack + magic
					tempMessage := fmt.Sprintf("Spell Attack Total: %v using %s with bonus: +%v", result, ability, attack)
					rolledMessage := text
					rollMonster := touchRollMonster(m, damageType, rolledMessage, tempMessage, result, resultDamage)
					returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
				}
				// if limitRange == "Touch" && i > 0 {
				// 	continue
				// }
				if spellAttack && i > 0 {
					continue
				}
				// if limitRange != "Touch" {
				if !spellAttack {
					rolledMessage := fmt.Sprintf("Rolled Damage: %s in %s", textDamage, damageDice)
					tmpMessage := fmt.Sprintf("Your Spell Save CD: %v. Spell Damage: %v", save, resultDamage)
					rollMonster := savingRollMonster(m, "spell", spellSaving, damageType, tmpMessage, save, resultDamage, spellcast.SpellLevel, []string{})
					returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
					returnResult.RolledMessage = rolledMessage
				}

			}

		}
		// message := fmt.Sprintf("Spell Attack Total: %v using %s with bonus: +%v. Your Spell Save CD: %v. Spell Damage: %v", result, ability, attack, save, resultDamage)
		// rolledMessage = fmt.Sprintf("Rolled %s in 1d20, Rolled Damage: %s in %s", text, textDamage, damageDice)
		// if spellcast.RacialTrait {
		// 	middle := fmt.Sprintf("Damage: %v", resultDamage)
		// 	message = racialTraitMessage + middle + racialTraitMessageSuffix
		// }
		returnResult.Message = message
		// returnResult.RolledMessage = rolledMessage
		// returnResult.RolledValue = result
		// returnResult.DamageValue = resultDamage
		// returnResult.DamageType = damageType
	}
	if damageDice == "" {
		returnResult.Message = fmt.Sprintf("Spell %s doesnt have any dice damage related", spellcast.SpellName)
	}

	return returnResult
}

func spellForHealing(name, class, healDice, healingIncreases, higherLevels string, classFeatures []string, ability map[string]int, level, spellLevel int) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)
	abilitySpell := AbilityForSpell(class)
	abilityModifier := CalcAbilityModifier(ability[abilitySpell])
	diceHeal := healDice
	if healingIncreases != "" {
		diceHeal = spellHealIncreases(healDice, healingIncreases, spellLevel)
	}
	var extraHeal int
	// "disciple-of-life" healing spells add 2 + spell level to healing value
	if utils.StringInSlice("disciple-of-life", classFeatures) {
		extraHeal = level + 2
	}
	if utils.StringInSlice("blessed-healer", classFeatures) {
		returnResult.RolledValue = extraHeal
	}
	r := diceroll.GetDice()
	healValue, text, _ := r.DiceRoll(diceHeal)
	heal := healValue + abilityModifier
	message := fmt.Sprintf("Healing: %v", heal)
	rolledMessage := fmt.Sprintf("Rolled: %s in %s", text, diceHeal)
	returnResult.HealingValue = heal + extraHeal
	returnResult.Message = message
	returnResult.RolledMessage = rolledMessage
	returnResult.Success = true
	if name == "heal" {
		if spellLevel == 0 {
			spellLevel = 6
		}
		basicHeal := healSpell(spellLevel)
		returnResult.HealingValue = basicHeal
		returnResult.Message = ""
		returnResult.RolledMessage = ""
	}
	if name == "mass-heal" {
		returnResult.HealingValue = 700
		returnResult.Message = ""
		returnResult.RolledMessage = ""
	}
	return returnResult
}
