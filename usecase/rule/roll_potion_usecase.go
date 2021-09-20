package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
)

//UsePotion func
func UsePotion(potion *rule.Potion) (*rule.Potion, error) {
	magicItem := GetMagicItemByName(potion.Name)
	if magicItem.Content == "" {
		err := fmt.Errorf("potion not found")
		return potion, err
	}
	if magicItem.RolePlay {
		potion.Message = "This potion doens't have any automatic calculation"
		return potion, nil
	}
	// add any improvements in a return struct and potion name in MagicalEffect in Player struct
	if magicItem.Feature != nil {
		switch {
		case len(magicItem.Feature.Advantages) != 0:
			potion.Advantages = magicItem.Feature.Advantages

		case len(magicItem.Feature.Disvantages) != 0:
			potion.Disvantages = magicItem.Feature.Disvantages

		case len(magicItem.Feature.AutoFail) != 0:
			potion.AutoFail = magicItem.Feature.AutoFail

		case len(magicItem.Feature.DamageResistence) != 0:
			potion.DamageResistence = magicItem.Feature.DamageResistence

		case len(magicItem.Feature.DamageVulnerabilities) != 0:
			potion.DamageVulnerabilities = magicItem.Feature.DamageVulnerabilities

		case len(magicItem.Feature.DamageImmunities) != 0:
			potion.DamageImmunities = magicItem.Feature.DamageImmunities

		case len(magicItem.Feature.ConditionImmunities) != 0:
			potion.ConditionImmunities = magicItem.Feature.ConditionImmunities

		case magicItem.Feature.AttackBonus != 0:
			potion.AttackMagicBonus = magicItem.Feature.AttackBonus

		case magicItem.Feature.NewAbility != nil:
			potion.NewAbility = magicItem.Feature.NewAbility
		}
	}
	r := diceroll.GetDice()
	if magicItem.Power != nil {
		switch {
		case magicItem.Power.Purpose == "poison":
			potion.DamageType = magicItem.Power.DamageType
			damage, _, err := r.DiceRoll(magicItem.Power.DamageDice)
			if err != nil {
				potion.Message = fmt.Sprintf("error to roll poisoning dices %v ", err)
			}
			potion.DamageValue = damage
			potion.DifficultClass = magicItem.Power.DifficultClass
			potion.Conditions = magicItem.Power.Condition
			potion.SavingThrow = magicItem.Power.SavingThrow
		case magicItem.Power.Purpose == "spell":
			potion.MagicalEffect = []string{magicItem.Power.SpellName}
		case magicItem.Power.Purpose == "spell-attack":
			potion.DifficultClass = magicItem.Power.DifficultClass
			potion.SavingThrow = magicItem.Power.SavingThrow
			damage, _, err := r.DiceRoll(magicItem.Power.DamageDice)
			if err != nil {
				potion.Message = fmt.Sprintf("error to roll poisoning dices %v ", err)
			}
			potion.DamageValue = damage
			potion.DamageType = magicItem.Power.DamageType

		case magicItem.Power.Purpose == "heal":
			heal, _, err := r.DiceRoll(magicItem.Power.Dice)
			if err != nil {
				potion.Message = fmt.Sprintf("error to roll healing dices %v ", err)
			}
			potion.HealingValue = heal
		case magicItem.Power.Purpose == "potion":
			potion.MagicalEffect = []string{magicItem.Power.SpellName}
			potion.HealingValue = magicItem.Power.ExtraHitPoints

		}
	}

	// if strings.Contains(potion.Name, "healing") {
	// 	r := dice.GetDice()
	// 	text, heal := r.RolledDicesInternal(magicItem.HealingDice, 0, false, false)
	// 	potion.Message = fmt.Sprintf("Rolled: %s in %s", text, magicItem.HealingDice)
	// 	potion.HealingValue = heal
	// }

	// if strings.Contains(potion.Name, "resistance") {
	// 	potion.DamageResistence = append(potion.DamageResistence, magicItem.DamageResistance...)

	// }

	// if strings.Contains(potion.Name, "strength") {
	// 	if potion.Ability["strength"] >= magicItem.NewAbility["strength"] {
	// 		potion.Message = "no effect"
	// 	}
	// 	if potion.Ability["strength"] < magicItem.NewAbility["strength"] {
	// 		potion.Ability["strength"] = magicItem.NewAbility["strength"]
	// 		potion.Message = "you feel much more stronger"
	// 	}

	// }

	// if strings.Contains(potion.Name, "poison") {
	// 	r := dice.GetDice()
	// 	text, damage := r.RolledDicesInternal(magicItem.PowerDamage, 0, false, false)
	// 	potion.Message = fmt.Sprintf("Rolled: %s in %s", text, magicItem.PowerDamage)
	// 	potion.DamageValue = damage
	// 	potion.DamageType = magicItem.PowerDamageType
	// }
	// potion.MagicalEffect = append(potion.MagicalEffect, magicItem.Name)

	// if potion.Name == "potion-of-heroism" {
	// 	potion.HealingValue = 10
	// 	potion.Message = "You are under the effect of the *bless* spell"
	// 	potion.MagicalEffect = append(potion.MagicalEffect, "bless")
	// }

	// if magicItem.MagicalEffect != "" {
	// 	potion.MagicalEffect = append(potion.MagicalEffect, magicItem.MagicalEffect)
	// }

	// if magicItem.AttackBonus != 0 {
	// 	potion.AttackMagicBonus = magicItem.AttackBonus
	// }

	return potion, nil
}
