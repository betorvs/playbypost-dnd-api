package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//CalcSavingsAbility func
func CalcSavingsAbility(saving *rule.SavingsCheck) *rule.ReturnCalcMessage {
	returnResult := new(rule.ReturnCalcMessage)
	var rolledMessage string
	if utils.StringInSlice(saving.Saving, saving.AutoFail) || utils.StringInSlice("all", saving.AutoFail) {
		returnResult.Message = "Automatically Faills"
		returnResult.RolledValue = 0
		return returnResult
	}
	// rage feature
	if utils.StringInSlice("rage", saving.ClassFeatures) && saving.Rage {
		saving.Advantages = append(saving.Advantages, "strength")
	}
	prof := CalcProficiency(saving.Level)
	abilityModifier := CalcAbilityModifier(saving.Ability[saving.Saving])
	rollDice := abilityModifier
	if utils.StringInSlice(saving.Saving, saving.Savings) {
		rollDice = prof + abilityModifier
	}
	returnResult.Success = false
	// d20 := "1d20"
	var adv, dis bool
	if utils.StringInSlice(saving.Saving, saving.Disvantages) || utils.StringInSlice(saving.Check, saving.Disvantages) {
		// d20 = "2d20k-1"
		dis = true
	}
	if utils.StringInSlice(saving.Saving, saving.Advantages) || utils.StringInSlice(saving.Check, saving.Advantages) {
		// d20 = "2d20k1"
		adv = true
	}
	if saving.Race == "gnome" && saving.Check == "spell" {
		if saving.Saving == "intelligence" || saving.Saving == "wisdom" || saving.Saving == "charisma" {
			// d20 = "2d20k1"
			adv = true
		}
	}
	r := diceroll.GetDice()
	d20 := utils.GetD20ToRoll(adv, dis)
	res, text, _ := r.DiceRoll(d20)
	var magic int
	if saving.MagicBonus != 0 {
		magic = saving.MagicBonus
	}
	result := res + rollDice + magic + saving.TemporaryBonus
	message := fmt.Sprintf("Saving Check Total: %v using %s", result, saving.Saving)
	rolledMessage = text
	if saving.DifficultClass != 0 && result >= saving.DifficultClass {
		message += " Hey. You got over it!"
		returnResult.Success = true
	}
	if res == 1 && saving.Race == "halfling" {
		message = "Hey, You need to run this again. Have a good Lucky!"
	}

	returnResult.Message = message
	returnResult.RolledMessage = rolledMessage
	returnResult.RolledValue = result
	return returnResult
}
