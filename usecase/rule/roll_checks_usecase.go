package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//CalcSkillOrAbility func
func CalcSkillOrAbility(check *rule.SkillOrAbilityCheck) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)
	prof := CalcProficiency(check.Level)
	var ability string
	var rolledMessage string
	if utils.StringInSlice(check.Check, AbilityList()) {
		ability = check.Check
	}
	if utils.StringInSlice(check.Check, SkillList()) {
		ability = AbilitySkill(check.Check)
	}
	// rage feature
	if utils.StringInSlice("rage", check.ClassFeatures) && check.Rage {
		check.Advantages = append(check.Advantages, "strength")
	}

	if utils.StringInSlice(ability, check.AutoFail) || utils.StringInSlice("all", check.AutoFail) {
		returnResult.Message = "Automatically Faills"
		returnResult.RolledValue = 0
		return returnResult
	}
	abilityModifier := CalcAbilityModifier(check.Ability[ability])
	rollDice := abilityModifier
	if utils.StringInSlice(check.Check, check.Skills) {
		rollDice = prof + abilityModifier
	}
	// expertise class feature
	if utils.StringInSlice(check.Check, SkillList()) {
		expertise := fmt.Sprintf("expertise-%s", check.Check)
		if utils.StringInSlice(expertise, check.ClassFeatures) {
			rollDice = rollDice + prof
		}
	}
	// if any features give a double, use it
	if check.DoubleProficiency {
		rollDice = rollDice + prof
	}
	// jack-of-all-trades class feature
	if !utils.StringInSlice(check.Check, check.Skills) && utils.StringInSlice("jack-of-all-trades", check.ClassFeatures) {
		rollDice = prof / 2
	}
	var adv, dis bool
	if utils.StringInSlice(ability, check.Disvantages) || utils.StringInSlice(check.Check, check.Disvantages) {
		dis = true
	}
	if utils.StringInSlice(ability, check.Advantages) || utils.StringInSlice(check.Check, check.Advantages) {
		adv = true
	}
	r := diceroll.GetDice()
	d20 := utils.GetD20ToRoll(adv, dis)
	res, text, _ := r.DiceRoll(d20)
	// if err1 != nil {

	// }
	// fmt.Println(text)
	var magic int
	if check.MagicBonus != 0 {
		magic = check.MagicBonus
	}
	result := res + rollDice + magic + check.TemporaryBonus

	message := fmt.Sprintf("Skill %s Check Total: %v using %s.", check.Check, result, ability)
	rolledMessage = text
	if check.DifficultyClass != 0 && result >= check.DifficultyClass {
		message += " Great! You did it!"
	}
	if res == 1 && check.Race == "halfling" {
		message = "Hey, You need to run this again. Have a good Lucky!"
	}
	returnResult.Message = message
	returnResult.RolledMessage = rolledMessage
	returnResult.RolledValue = result
	return returnResult
}
