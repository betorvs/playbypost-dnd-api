package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/rule"
)

//UsePotion func
func UsePotion(potion *rule.Potion) (*rule.Potion, error) {
	magicItem := GetMagicItemByName(potion.Name)
	if magicItem.Content == "" {
		err := fmt.Errorf("potion not found")
		return potion, err
	}
	// add any improvements in a return struct and potion name in MagicalEffect in Player struct

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
