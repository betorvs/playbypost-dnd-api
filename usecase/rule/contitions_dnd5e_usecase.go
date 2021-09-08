package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/rule"
)

// ListConditions func
func ListConditions() []string {
	return []string{"blinded", "charmed", "deafened", "exhaustion", "frightened", "grappled", "incapacitated", "invisible", "paralyzed", "petrified", "poisoned", "prone", "pestrained", "stunned", "unconscious"}
}

//ListOfDamageTypes func
func ListOfDamageTypes() []string {
	return []string{"acid", "bludgeoning", "cold", "fire", "force", "lightning", "necrotic", "piercing", "poison", "psychic", "radiant", "slashing", "thunder"}
}

func listOfWeaponDamageType() []string {
	return []string{"bludgeoning", "piercing", "slashing"}
}

// func fiveSenses() []string {
// 	return []string{"vision", "hearing", "smell", "taste", "touch"}
// }

func conditionsCheck(condition string, level int) (description []string, disvantages []string, automaticallyFail []string) {
	// always check if creature with that condition has advantage or disvantage
	switch condition {
	case "Blinded", "blinded":
		description = []string{
			"A blinded creature can’t see and automatically fails any ability check that requires sight.",
			"Attack rolls against the creature have advantage, and the creature’s attack rolls have disadvantage."}
		disvantages = []string{"attack"}
		automaticallyFail = []string{"vision"}

	case "Charmed", "charmed":
		description = []string{
			"A charmed creature can’t attack the charmer or target the charmer with harmful abilities or magical effects.",
			"The charmer has advantage on any ability check to interact socially with the creature."}

	case "Deafened", "deafened":
		description = []string{"A deafened creature can’t hear and automatically fails any ability check that requires hearing."}
		automaticallyFail = []string{"hearing"}
	case "Exhaustion", "exhaustion":
		switch level {
		case 1:
			description = []string{"Disadvantage on ability checks"}
			disvantages = []string{"ability"}
		case 2:
			description = []string{"Speed halved"}
			disvantages = []string{"ability"}
		case 3:
			description = []string{"Disadvantage on attack rolls and saving throws"}
			disvantages = []string{"ability", "attack", "saving"}
		case 4:
			description = []string{"Hit point maximum halved"}
			disvantages = []string{"ability", "attack", "saving"}
		case 5:
			description = []string{"Speed reduced to 0"}
			disvantages = []string{"ability", "attack", "saving"}
		case 6:
			description = []string{"Death"}
			disvantages = []string{"ability", "attack", "saving"}
			automaticallyFail = []string{"all"}
		default:
			description = []string{
				"Disadvantage on ability checks",
				"Speed halved",
				"Disadvantage on attack rolls and saving throws",
				"Hit point maximum halved",
				"Speed reduced to 0",
				"Death"}
			disvantages = []string{"ability", "attack", "saving"}
			automaticallyFail = []string{"all"}
		}

	case "Frightened", "frightened":
		description = []string{
			"A frightened creature has disadvantage on ability checks and attack rolls while the source of its fear is within line of sight.",
			"The creature can’t willingly move closer to the source of its fear."}
		disvantages = []string{"ability", "attack"}
	case "Grappled", "grappled":
		description = []string{
			"A grappled creature’s speed becomes 0, and it can’t benefit from any bonus to its speed.",
			"The condition ends if the grappler is incapacitated (see the condition).",
			"The condition also ends if an effect removes the grappled creature from the reach of the grappler or grappling effect, such as when a creature is hurled away by the *thunder-wave* spell."}

	case "Incapacitated", "incapacitated":
		description = []string{"An incapacitated creature can’t take actions or reactions."}
		automaticallyFail = []string{"all"}

	case "Invisible", "invisible":
		description = []string{
			"An invisible creature is impossible to see without the aid of magic or a special sense. For the purpose of hiding, the creature is heavily obscured. The creature’s location can be detected by any noise it makes or any tracks it leaves.",
			"Attack rolls against the creature have disadvantage, and the creature’s attack rolls have advantage."}
		disvantages = []string{"attack"}
	case "Paralyzed", "paralyzed":
		description = []string{
			"A paralyzed creature is incapacitated (see the condition) and can’t move or speak.",
			"The creature automatically fails Strength and Dexterity saving throws.",
			"Attack rolls against the creature have advantage.",
			"Any attack that hits the creature is a critical hit if the attacker is within 5 feet of the creature."}
		disvantages = []string{"attack"}
		automaticallyFail = []string{"strength", "dexterity"}
	case "Petrified", "petrified":
		description = []string{
			"A petrified creature is transformed, along with any nonmagical object it is wearing or carrying, into a solid inanimate substance (usually stone). Its weight increases by a factor of ten, and it ceases aging.",
			"The creature is incapacitated (see the condition), can’t move or speak, and is unaware of its surroundings.",
			"Attack rolls against the creature have advantage.",
			"The creature automatically fails Strength and Dexterity saving throws.",
			"The creature has resistance to all damage.",
			"The creature is immune to poison and disease, although a poison or disease already in its system is suspended, not neutralized."}
		automaticallyFail = []string{"strength", "dexterity"}
	case "Poisoned", "poisoned":
		description = []string{"A poisoned creature has disadvantage on attack rolls and ability checks."}
		disvantages = []string{"ability", "attack"}
	case "Prone", "prone":
		description = []string{
			"A prone creature’s only movement option is to crawl, unless it stands up and thereby ends the condition.",
			"The creature has disadvantage on attack rolls.",
			"An attack roll against the creature has advantage if the attacker is within 5 feet of the creature. Otherwise, the attack roll has disadvantage."}
		disvantages = []string{"attack"}
	case "Restrained", "pestrained":
		description = []string{
			"A restrained creature’s speed becomes 0, and it can’t benefit from any bonus to its speed.",
			"Attack rolls against the creature have advantage, and the creature’s attack rolls have disadvantage.",
			"The creature has disadvantage on Dexterity saving throws."}
		disvantages = []string{"attack", "dexterity"}
	case "Stunned", "stunned":
		description = []string{
			"A stunned creature is incapacitated (see the condition), can’t move, and can speak only falteringly.",
			"The creature automatically fails Strength and Dexterity saving throws.",
			"Attack rolls against the creature have advantage."}
		disvantages = []string{"attack"}
		automaticallyFail = []string{"strength", "dexterity"}
	case "Unconscious", "unconscious":
		description = []string{
			"An unconscious creature is incapacitated (see the condition), can’t move or speak, and is unaware of its surroundings",
			"The creature drops whatever it’s holding and falls prone.",
			"The creature automatically fails Strength and Dexterity saving throws.",
			"Attack rolls against the creature have advantage.",
			"Any attack that hits the creature is a critical hit if the attacker is within 5 feet of the creature."}
		disvantages = []string{"attack"}
		automaticallyFail = []string{"strength", "dexterity"}
	}
	return description, disvantages, automaticallyFail
}

func conditionsMap(name string) map[string]string {
	cond := make(map[string]string)
	desc, _, _ := conditionsCheck(name, 0)
	cond["name"] = name
	var description string
	for _, v := range desc {
		description += fmt.Sprintf("%s \n", v)
	}
	cond["description"] = description

	return cond
}

func damageTypes(name string) string {
	// damage := make(map[string]string, 0)
	switch name {
	case "acid":
		return "The corrosive spray of a black dragon’s breath and the dissolving enzymes secreted by a black pudding deal acid damage."
	case "bludgeoning":
		return "Blunt force attacks—hammers, falling, constriction, and the like—deal bludgeoning damage."
	case "cold":
		return "The infernal chill radiating from an ice devil’s spear and the frigid blast of a white dragon’s breath deal cold damage."
	case "fire":
		return "Red dragons breathe fire, and many spells conjure flames to deal fire damage."
	case "force":
		return "Force is pure magical energy focused into a damaging form. Most effects that deal force damage are spells, including *magic missile* and *spiritual weapon*."
	case "lightning":
		return "A *lightning bolt* spell and a blue dragon’s breath deal lightning damage."
	case "necrotic":
		return "Necrotic damage, dealt by certain undead and a spell such as *chill touch*, withers matter and even the soul."
	case "piercing":
		return "Puncturing and impaling attacks, including spears and monsters’ bites, deal piercing damage."
	case "poison":
		return "Venomous stings and the toxic gas of a green dragon’s breath deal poison damage."
	case "psychic":
		return "Mental abilities such as a mind flayer’s psionic blast deal psychic damage."
	case "radiant":
		return "Radiant damage, dealt by a cleric’s *flame strike* spell or an angel’s smiting weapon, sears the flesh like fire and overloads the spirit with power."
	case "slashing":
		return "Swords, axes, and monsters’ claws deal slashing damage."
	case "thunder":
		return "A concussive burst of sound, such as the effect of the *thunderwave* spell, deals thunder damage."
	default:
		return ""
	}

}

func damageTypeMap(name string) map[string]string {
	damage := make(map[string]string)
	desc := damageTypes(name)
	damage["name"] = name
	damage["description"] = desc

	return damage
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
