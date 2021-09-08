package rule

import (
	"fmt"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//CalcFullAttackwithWeapon func
func CalcFullAttackwithWeapon(attack *rule.Attack) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)
	var message string
	var rolledMessage string
	for i, w := range attack.Weapon {
		if i == 1 && attack.TwoHands {
			continue
		}
		damage, damageType, usedAbility, attackBonus, damageBonus := calcBonusAttackPerWeapon(w, attack.ArmorProficiency, attack.ClassFeatures, attack.Level, i, attack.Ability, attack.Rage, attack.TwoHands)
		var adv, dis bool
		if utils.StringInSlice(usedAbility, attack.Disvantages) {
			dis = true
		}
		if utils.StringInSlice(usedAbility, attack.Advantages) {
			adv = true
		}
		r := diceroll.GetDice()
		d20 := utils.GetD20ToRoll(adv, dis)
		res, text, _ := r.DiceRoll(d20)
		damageRollDice, textDamage, _ := r.DiceRoll(damage)
		var magic int
		if attack.MagicBonus != 0 {
			magic = attack.MagicBonus
		}
		var tempBonus int
		if attack.TemporaryBonus != 0 && i == 0 {
			tempBonus = attack.TemporaryBonus
		}
		result := res + attackBonus + magic + tempBonus
		// fighting-style-great-weapon-fighting only after roll
		var firstRollGreatFighting string
		var twoHands string
		if attack.TwoHands && utils.StringInSlice("fighting-style-great-weapon-fighting", attack.ClassFeatures) && attack.TwoHands {
			if damageRollDice == 1 || damageRollDice == 2 {
				firstRollGreatFighting = fmt.Sprintf("First Damage Rolled: %s", textDamage)
				damageRollDice, textDamage, _ = r.DiceRoll(damage)
				twoHands = "with two hands"
			}
		}
		var extraDamageType string
		var textExtraDamage string
		var damageExtraRollDice int
		if attack.UsingFeature != "" && i == 0 && utils.StringInSlice(attack.UsingFeature, attack.ClassFeatures) {
			choosen := attack.UsingFeatureType
			paladinEnemy := returnMonsterPaladinEnemy(attack.Monster)
			// using the same as weapon
			if attack.UsingFeature == "domain-war-divine-strike" || attack.UsingFeature == "archetype-hunter-hunters-prey-colossus-slayer" || attack.UsingFeature == "sneak-attack" {
				choosen = damageType
			}
			var damageExtra string
			damageExtra, extraDamageType = extraDamageMeleeAttackFeature(attack.UsingFeature, choosen, attack.Level, attack.UsingFeatureSlot, paladinEnemy)
			damageExtraRollDice, textExtraDamage, _ = r.DiceRoll(damageExtra)
			textExtraDamage = fmt.Sprintf("Rolled Extra damage %s in %s", textExtraDamage, damageExtra)
		}

		resultDamage := damageRollDice + damageBonus + magic
		// if utils.StringInSlice(extraDamageType, listOfWeaponDamageType()) {
		// 	rolledMessage += textExtraDamage
		// 	resultDamage = damageRollDice + damageBonus + magic + damageExtraRollDice
		// }
		if attack.UsingFeature == "" {
			var classFeature string
			for _, v := range fightingStyleFullNameList() {
				classFeature = utils.ExtractInSlice(v, attack.ClassFeatures)
			}
			message = fmt.Sprintf("Attack using %s with %v %s", classFeature, attack.Weapon, twoHands)
		}

		if attack.Monster != "" {
			attackMessage := fmt.Sprintf("Attack Total: %v and Damage: %v of %s using weapon %s, with attack Bonus +%v using %s", result, resultDamage, damageType, w, attackBonus, usedAbility)
			rolledMessage = fmt.Sprintf("Rolled %s in 1d20 and damage rolled %s in %s %s", text, textDamage, damage, firstRollGreatFighting)
			// if utils.StringInSlice(extraDamageType, listOfWeaponDamageType()) {
			// 	rolledMessage += textExtraDamage
			// }
			if res >= 18 {
				damageBonus, damageText := checkCriticalHit(damage, attack.Race, attack.ClassFeatures, res)
				resultDamage = damageRollDice + damageBonus + magic + damageBonus
				rolledMessage += damageText
			}
			rollMonster := attackMonster(attack.Monster, w, damageType, attackMessage, rolledMessage, result, resultDamage, false)
			returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
			if utils.StringInSlice(attack.UsingFeature, attack.ClassFeatures) {
				attackMessage := fmt.Sprintf("Using %s", attack.UsingFeature)
				rollMonster := attackMonster(attack.Monster, w, extraDamageType, attackMessage, textExtraDamage, result, damageExtraRollDice, false)
				returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
				message = fmt.Sprintf("Attack using %s and %s", attack.UsingFeature, w)
			}
		}

	}

	returnResult.Message = message
	// returnResult.RolledMessage = rolledMessage
	return returnResult
}

func calcBonusAttackPerWeapon(weapon string, armorProficiency, classFeatures []string, level, attackNumber int, ability map[string]int, usingRage, twoHands bool) (damage, damageType, usedAbility string, attackBonus, damageBonus int) {
	// kind, _, damage, damageType, _, properties := WeaponsByName(weapon)
	weapons := WeaponsByName(weapon)
	damage = weapons.Damage
	damageType = weapons.DamageType
	var prof int
	if utils.StringInSlice(weapons.Kind, armorProficiency) {
		prof = CalcProficiency(level)
	}
	if utils.StringInSlice(weapon, armorProficiency) {
		prof = CalcProficiency(level)
	}
	if strings.Contains(weapon, "unarmed") {
		prof = CalcProficiency(level)
	}
	// "martial-arts"
	if utils.StringInSlice("martial-arts", classFeatures) {
		if strings.Contains(weapon, "unarmed") {
			weapons.Properties = "Finesse"
			unarmed, _, _ := MonkClass(level)
			damage = unarmed
		}

		if weapon == "shortsword" || weapons.Kind == "simple-weapon" {
			weapons.Properties += "Finesse"
		}
	}
	usedAbility = "strength"
	modifier := CalcAbilityModifier(ability["strength"])
	if strings.Contains(weapons.Properties, "Finesse") && ability["dexterity"] > ability["strength"] {
		modifier = CalcAbilityModifier(ability["dexterity"])
		usedAbility = "dexterity"
	}
	var archeryBonus int
	if strings.Contains(weapons.Properties, "Ammunition") {
		modifier = CalcAbilityModifier(ability["dexterity"])
		usedAbility = "dexterity"
		if utils.StringInSlice("fighting-style-archery", classFeatures) {
			archeryBonus = 2
		}
	}
	var twoHandWeapon bool
	if twoHands && strings.Contains(weapons.Properties, "Versatile") {
		versatile := strings.Split(weapons.Properties, ":")
		damage = versatile[1]
		twoHandWeapon = true
	}
	if strings.Contains(weapons.Properties, "Two-handed") {
		twoHandWeapon = true
	}
	// rage feature
	var rageBonus int
	if usedAbility == "strength" && usingRage && utils.StringInSlice("rage", classFeatures) {
		_, rage := BarbarianClass(level)
		rageBonus = rage
	}
	// fighting-style-dueling
	var duelingBonus int
	if utils.StringInSlice("fighting-style-dueling", classFeatures) && !twoHandWeapon {
		duelingBonus = 2
	}

	attackBonus = prof + modifier + archeryBonus
	if !utils.StringInSlice("fighting-style-two-weapon-fighting", classFeatures) {
		modifier = 0
	}
	damageBonus = modifier + duelingBonus + rageBonus

	return damage, damageType, usedAbility, attackBonus, damageBonus
}

func checkCriticalHit(damage, race string, classFeatures []string, roll int) (damageValue int, damageText string) {
	var critical bool
	// archetype-champion-improved-critical 19 e 20
	if utils.StringInSlice("archetype-champion-improved-critical", classFeatures) && roll >= 19 {
		critical = true
	}
	if utils.StringInSlice("archetype-champion-superior-critical", classFeatures) && roll >= 18 {
		critical = true
	}
	if roll == 20 {
		critical = true
	}
	r := diceroll.GetDice()
	// archetype-champion-superior-critical
	if critical {
		var halfOrcExtraDamage int
		var halfOrcExtraTextDamage string
		if race == "half-orc" {
			halfOrcExtraDamage, halfOrcExtraTextDamage, _ = r.DiceRoll(damage)
		}
		extraDamage, extraTextDamage, _ := r.DiceRoll(damage)
		damageText = fmt.Sprintf("Critical Hit Rolled: %s %s", extraTextDamage, halfOrcExtraTextDamage)
		damageValue = extraDamage + halfOrcExtraDamage
	}

	return damageValue, damageText
}
