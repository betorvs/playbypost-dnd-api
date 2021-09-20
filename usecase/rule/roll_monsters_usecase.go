package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//CalcFullMonsterAttackwithWeapon func
func CalcFullMonsterAttackwithWeapon(check *rule.MonsterRoll) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)

	monster := MosterByName(check.Name)
	var attack int
	var damage int
	var damageDice string
	var damageType string
	var criticalHit string
	var rolledMessage string
	var successMessage string
	for _, v := range monster.WeaponAttack {
		if v.Name == check.Weapon {
			attack = v.Attack
			damage = v.AverageDamage
			damageDice = v.Damage
			damageType = v.DamageType
		}
	}
	ability := "strength"
	if utils.StringInSlice(ability, check.AutoFail) || utils.StringInSlice("all", check.AutoFail) {
		returnResult.Message = "Automatically Faills"
		returnResult.RolledValue = 0
		returnResult.DamageValue = 0
		return returnResult
	}
	// enemyRage
	if check.EnemyRage {
		rageResistances := []string{"bludgeoning", "piercing", "slashing"}
		check.EnemyDamageResistances = append(check.EnemyDamageResistances, rageResistances...)
	}
	r := diceroll.GetDice()
	var adv, dis bool
	if utils.StringInSlice(ability, check.Disvantages) {
		dis = true
	}
	if utils.StringInSlice(ability, check.Advantages) {
		adv = true
	}
	d20 := utils.GetD20ToRoll(adv, dis)
	res, text, _ := r.DiceRoll(d20)
	var totalDamage int
	var damageRoll int
	var damageText string
	if damage != 1 {
		damageRoll, damageText, _ = r.DiceRoll(damageDice)
	}
	if res == 20 && !utils.StringInSlice("criticalhit", check.EnemyDamageImmunities) {
		extraDamage, extraDamageText, _ := r.DiceRoll(damageDice)
		damageText += extraDamageText
		totalDamage = damageRoll + extraDamage
		criticalHit = "it has CRITICAL HIT!!!"
	}
	if len(check.EnemyDamageImmunities) > 0 && utils.StringInSlice(damageType, check.EnemyDamageImmunities) {
		damage = 0
	}
	if len(check.EnemyDamageVulnerabilities) > 0 && utils.StringInSlice(damageType, check.EnemyDamageVulnerabilities) {
		damage = damage * 2
	}
	if len(check.EnemyDamageResistances) > 0 && utils.StringInSlice(damageType, check.EnemyDamageResistances) {
		damage = damage / 2
	}
	result := res + attack
	message := fmt.Sprintf("Attack Total: %v and Damage: %v using weapon %s", result, totalDamage, check.Weapon)
	rolledMessage = text + damageText
	returnResult.Success = false
	if check.DifficultClass != 0 && result >= check.DifficultClass {
		successMessage = fmt.Sprintf("Bad news! Monster %s hits you! %s", monster.Name, criticalHit)
		returnResult.Success = true
	}

	returnResult.Message = message
	returnResult.RolledMessage = rolledMessage
	returnResult.SuccessMessage = successMessage
	returnResult.RolledValue = result
	returnResult.DamageValue = damage
	if !returnResult.Success {
		returnResult.DamageValue = 0
	}
	returnResult.DamageType = damageType
	return returnResult
}

//CalcMonsterSavingsAbility func
func CalcMonsterSavingsAbility(check *rule.MonsterRoll) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)
	monster := MosterByName(check.Name)
	var saving int
	var successMessage string
	for k, v := range monster.Savings {
		if k == check.Check {
			saving = v
		}
	}
	if utils.StringInSlice(check.Check, check.AutoFail) || utils.StringInSlice("all", check.AutoFail) {
		returnResult.Message = "Automatically Faills"
		returnResult.RolledValue = 0
		return returnResult
	}
	r := diceroll.GetDice()
	var adv, dis bool
	if utils.StringInSlice(check.Check, check.Disvantages) {
		dis = true
	}
	if utils.StringInSlice(check.Check, check.Advantages) {
		adv = true
	}
	d20 := utils.GetD20ToRoll(adv, dis)
	res, text, _ := r.DiceRoll(d20)

	result := res + saving
	message := fmt.Sprintf("Saving Check Total: %v ( %s ) using %s", result, text, check.Check)
	if check.DifficultClass != 0 && result >= check.DifficultClass {
		successMessage = fmt.Sprintf(" Hey. Monster %s avoid your attack!", monster.Name)
		returnResult.Success = true
	}
	returnResult.Message = message
	returnResult.SuccessMessage = successMessage
	returnResult.RolledValue = result
	return returnResult
}

//CalcMonsterChecks func
func CalcMonsterChecks(check *rule.MonsterRoll) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)
	monster := MosterByName(check.Name)
	if utils.StringInSlice(check.Check, check.AutoFail) || utils.StringInSlice("all", check.AutoFail) {
		returnResult.Message = "Automatically Faills"
		returnResult.RolledValue = 0
		return returnResult
	}
	var successMessage string
	var checkLocal string
	var checkLocalValue int
	// fmt.Println(check)
	if utils.StringInSlice(check.Check, AbilityList()) {
		// fmt.Println("using ability")
		checkLocal = check.Check
		checkLocalValue = CalcAbilityModifier(monster.Ability[checkLocal])
	}
	if utils.StringInSlice(check.Check, SkillList()) {
		// fmt.Println("using skill")
		checkLocal = check.Check
		skillAbility := AbilitySkill(check.Check)
		checkLocalValue = CalcAbilityModifier(monster.Ability[skillAbility])
		// fmt.Println(checkLocal, skillAbility)
		if monster.Skills[checkLocal] != 0 {
			checkLocalValue = monster.Skills[checkLocal]
		}
	}
	r := diceroll.GetDice()
	var adv, dis bool
	if utils.StringInSlice(check.Check, check.Disvantages) {
		dis = true
	}
	if utils.StringInSlice(check.Check, check.Advantages) {
		adv = true
	}
	d20 := utils.GetD20ToRoll(adv, dis)
	res, text, _ := r.DiceRoll(d20)

	result := res + checkLocalValue
	message := fmt.Sprintf("Check Total: %v ( %s ) using %s.", result, text, check.Check)
	if check.DifficultClass != 0 && result >= check.DifficultClass {
		successMessage = fmt.Sprintf(" Hey. Monster %s check was successfully!", monster.Name)
		returnResult.Success = true
	}
	returnResult.Message = message
	returnResult.SuccessMessage = successMessage
	returnResult.RolledValue = result
	return returnResult
}

//CalcMonstersInitiative func
func CalcMonstersInitiative(monsters *rule.SimpleList) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)
	var leader string
	var highest int
	for _, m := range monsters.List {
		monster := MosterByName(m)
		dex := CalcAbilityModifier(monster.Ability["dexterity"])
		if dex > highest {
			highest = dex
			leader = monster.Name
		}
	}
	r := diceroll.GetDice()
	d20 := utils.GetD20ToRoll(false, false)
	res, text, _ := r.DiceRoll(d20)
	result := res + highest
	message := fmt.Sprintf("Initiative: %v from leader %s", result, leader)
	rolledMessage := text
	returnResult.Message = message
	returnResult.RolledMessage = rolledMessage
	returnResult.RolledValue = result

	return returnResult
}

//TurnUndeadRolls func
func TurnUndeadRolls(monsters *rule.MonsterTurn) (*rule.ReturnCalcMessage, error) {
	returnResult := new(rule.ReturnCalcMessage)
	var classFeature string
	for _, f := range monsters.ClassFeatures {
		if utils.StringInSlice(f, turnMonstersFeatureList()) {
			classFeature = f
		}
	}
	if classFeature == "" {
		err := fmt.Errorf("missing turn monster class feature %v", turnMonstersFeatureList())
		return returnResult, err
	}
	monsterTurnAffected := []string{"undead"}
	clericTurnUndead := true

	// "sacred-oath-of-devotion-oauth-spells-and-channel-divinity" fiend and undead
	// "sacred-oath-of-ancients-oauth-spells-and-channel-divinity" fey and fiends
	if classFeature == "sacred-oath-of-devotion-channel-divinity" {
		monsterTurnAffected = []string{"undead", "fiend"}
		clericTurnUndead = false
	}
	if classFeature == "sacred-oath-of-ancients-channel-divinity" {
		monsterTurnAffected = []string{"fey", "fiend"}
		clericTurnUndead = false
	}
	prof := CalcProficiency(monsters.Level)
	abilityModifier := CalcAbilityModifier(monsters.Ability["wisdom"])
	clericDC := prof + 8 + abilityModifier
	if clericTurnUndead {
		for _, u := range monsters.MonsterList {
			tmpMessage := fmt.Sprintf("Your turn undead have DC %v", clericDC)
			rollMonster := savingRollMonster(u, "turn-undead", "wisdom", "", tmpMessage, clericDC, 0, monsters.Level, monsterTurnAffected)
			returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
		}
		returnResult.Message = fmt.Sprintf("Using turn undead with level %v in %v", monsters.Level, monsters.MonsterList)
	}
	if !clericTurnUndead {
		for _, u := range monsters.MonsterList {
			tmpMessage := fmt.Sprintf("Your %s have DC %v", classFeature, clericDC)
			rollMonster := savingRollMonster(u, "turn-monster", "wisdom", "", tmpMessage, clericDC, 0, monsters.Level, monsterTurnAffected)
			returnResult.MonstersResult = append(returnResult.MonstersResult, *rollMonster)
		}
		returnResult.Message = fmt.Sprintf("Using %s with level %v in %v", classFeature, monsters.Level, monsters.MonsterList)
	}

	return returnResult, nil
}

func attackMonster(name, weapon, damageType, attackMessage, rolledMessage string, attackValue, damage int, criticalHit bool) *rule.MonsterResult {
	result := new(rule.MonsterResult)
	monster := MosterByName(name)
	// check if attack hits monster AC
	result.Name = monster.Name
	result.RolledMessage = rolledMessage
	result.Success = false
	if attackValue >= monster.ArmorClass || criticalHit {
		result.SuccessMessage = fmt.Sprintf("GOOD NEWS! You hit with %s %s!", weapon, damageType)
		tempMessage, tempDamage, tempSuccess := validateDamage(monster.Name, damageType, monster.DamageImmunities, monster.DamageResistances, monster.DamageVulnerabilities, damage)
		result.DamageValue = tempDamage
		result.DamageType = damageType
		result.DamageMessage = tempMessage
		result.Message = attackMessage
		result.Success = tempSuccess
	}

	return result
}

func savingRollMonster(name, kind, saving, damageType, message string, difficult, damage, level int, monstersAffected []string) *rule.MonsterResult {
	result := new(rule.MonsterResult)
	// difficult: saving DC
	// level: effect level, like spell level or turn-undead cleric level
	monster := MosterByName(name)
	// check if saving was needed
	r := diceroll.GetDice()
	d20 := utils.GetD20ToRoll(false, false)
	res, text, _ := r.DiceRoll(d20)
	roll := res + monster.Savings[saving]
	if roll < difficult {
		switch kind {
		case "turn-undead":
			result.Name = monster.Name
			result.Success = true
			result.Message = message
			result.DamageMessage = fmt.Sprintf("Undead %s failed", monster.Name)
			result.RolledMessage = fmt.Sprintf("Undead %s rolled %s", monster.Name, text)
			clericDestroy := clericDestroyUndead(level)
			if level >= 5 && clericDestroy > monster.Challenge {
				result.Message = fmt.Sprintf("Undead %s destroyed", monster.Name)
				result.RolledMessage = fmt.Sprintf("Undead %s rolled %s", monster.Name, text)
			}
		case "turn-monster":
			if utils.StringInSlice(monster.Type, monstersAffected) {
				result.Name = monster.Name
				result.Success = true
				result.Message = message
				result.DamageMessage = fmt.Sprintf("Monster %s failed", monster.Name)
				result.RolledMessage = fmt.Sprintf("Monster %s rolled %s", monster.Name, text)
			}
		case "spell":
			result.Name = monster.Name
			tempMessage, tempDamage, tempSuccess := validateDamage(monster.Name, damageType, monster.DamageImmunities, monster.DamageResistances, monster.DamageVulnerabilities, damage)
			result.DamageValue = tempDamage
			result.DamageType = damageType
			result.Message = message
			result.DamageMessage = tempMessage
			result.Success = tempSuccess
			result.RolledMessage = fmt.Sprintf("Monster Rolled %s in 1d20", text)
			result.RolledValue = res
			result.Success = true
		}
	}
	if roll >= difficult {
		switch kind {
		case "turn-undead":
			result.Name = monster.Name
			result.Message = fmt.Sprintf("Undead %s PASS", monster.Name)
			result.RolledMessage = fmt.Sprintf("Undead %s rolled %s", monster.Name, text)
			result.Success = false
		case "turn-monster":
			result.Name = monster.Name
			result.Message = fmt.Sprintf("Monster %s PASS", monster.Name)
			result.RolledMessage = fmt.Sprintf("Monster %s rolled %s", monster.Name, text)
			result.Success = false
		case "spell":
			result.Name = monster.Name
			// result.Message = fmt.Sprintf("Monster %s PASS and receive half of damage", monster.Name)
			// result.RolledMessage = fmt.Sprintf("Monster %s rolled %s", monster.Name, text)
			tempMessage, tempDamage, tempSuccess := validateDamage(monster.Name, damageType, monster.DamageImmunities, monster.DamageResistances, monster.DamageVulnerabilities, damage)
			result.DamageValue = tempDamage / 2
			result.Message = message
			result.DamageMessage = tempMessage
			result.DamageType = damageType
			result.Success = tempSuccess
			result.Success = false
		}
	}
	if kind == "turn-undead" || kind == "turn-monster" {
		if !utils.StringInSlice(monster.Type, monstersAffected) {
			result.Name = monster.Name
			result.Message = fmt.Sprintf("Monster %s NOT AFFECTED", monster.Name)
			result.RolledMessage = ""
			result.Success = false
		}
	}

	return result
}

func touchRollMonster(name, damageType, rolledMessage, message string, attack, damage int) *rule.MonsterResult {
	result := new(rule.MonsterResult)
	monster := MosterByName(name)
	if attack >= monster.ArmorClass {
		result.Name = monster.Name
		tempMessage, tempDamage, tempSuccess := validateDamage(monster.Name, damageType, monster.DamageImmunities, monster.DamageResistances, monster.DamageVulnerabilities, damage)
		result.DamageValue = tempDamage
		result.DamageType = damageType
		result.RolledMessage = rolledMessage
		result.Message = message
		result.DamageMessage = tempMessage
		result.Success = tempSuccess
		result.Success = true
	}
	if attack < monster.ArmorClass {
		result.Name = monster.Name
		result.Message = fmt.Sprintf("Monster %s PASS", monster.Name)
		result.RolledMessage = rolledMessage
		result.Success = false
	}
	return result
}

func validateDamage(name, damageType string, damageImmunities, damageResistances, damageVulnerabilities []string, damageValue int) (message string, damage int, success bool) {
	message = fmt.Sprintf("Monster %s fells your damage", name)
	damage = damageValue
	success = true
	if len(damageImmunities) > 0 && utils.StringInSlice(damageType, damageImmunities) {
		damage = 0
		message = fmt.Sprintf("Monster %s resist all your damage", name)
		success = false
	}
	if len(damageResistances) > 0 && utils.StringInSlice(damageType, damageResistances) {
		damage = damageValue / 2
		message = fmt.Sprintf("Monster %s resist all your damage", name)
	}
	if len(damageVulnerabilities) > 0 && utils.StringInSlice(damageType, damageVulnerabilities) {
		damage = damage * 2
		message = fmt.Sprintf("Monster %s felt a lot more than you expected", name)
	}
	return message, damage, success
}
