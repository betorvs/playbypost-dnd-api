package rule

import (
	"fmt"
	"math"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

// CalcProficiency func
func CalcProficiency(level int) int {
	result := 2 + ((level - 1) / 4)
	return result
}

// CalcAbilityModifier func
func CalcAbilityModifier(attr int) int {
	result := math.Floor((float64(attr) - 10) / 2)
	return int(result)
}

//CalcMaxHP func
func CalcMaxHP(level int, dice int, mod int) int {
	result := dice + mod
	if level > 1 {
		calc := level - 1
		result += calc*(dice/2) + calc*mod
	}
	return result
}

// XPNeeded func
func XPNeeded(level int) int {
	xp := []int{0, 300, 900, 2700, 6500, 14000, 23000, 34000, 48000, 64000, 85000, 100000, 120000, 140000, 165000, 195000, 225000, 265000, 305000, 355000, 410000}
	if level < 1 {
		return 0
	}
	return xp[level]
}

func SkillStatistics(name string) rule.Skill {
	var skill rule.Skill
	skill.Name = name
	skill.Description = fmt.Sprintf("https://www.dndbeyond.com/sources/basic-rules/using-ability-scores#%s", strings.ReplaceAll(strings.Title(name), "-", ""))
	skill.Ability = AbilitySkill(name)
	skill.Title = fmt.Sprintf("%s (%s)", strings.Title(name), AbilitySkill(name))
	return skill
}

//RaceStatistics func
func RaceStatistics(race, subrace string) *rule.Race {
	ability := make(map[string]int)
	output := new(rule.Race)
	output.Speedmeasure = "ft"
	output.ArmorProficiency = []string{}
	output.Subrace = ""
	switch race {
	case "Dwarf", "dwarf":
		output.Name = "dwarf"
		output.Size = "medium"
		output.Speed = 25
		output.Special = []string{"Darkvision", "Dwarven Resilience", "Dwarven Combat Training", "Tool Proficiency", "Stonecunning"}
		ability["constitution"] = 2
		output.Language = []string{"common", "dwarvish"}
		output.Resistance = []string{"poison"}
		output.Advantages = []string{"poison"}
		dwarvenCombatTraining := []string{"battleaxe", "handaxe", "light hammer", "warhammer"}
		output.ArmorProficiency = dwarvenCombatTraining
		switch subrace {
		case "Hill Dwarf", "hill-dwarf", "hill", "dwarf-hill":
			output.Subrace = "hill-dwarf"
			output.Special = append(output.Special, "Dwarven Toughness")
			ability["wisdom"] = 1

		case "Mountain Dwarf", "mountain-dwarf", "mountain", "dwarf-mountain":
			output.Subrace = "mountain-dwarf"
			output.Special = append(output.Special, "Dwarven Armor Training")
			dwarvenArmorTraining := []string{"light-armor", "medium-armor"}
			output.ArmorProficiency = append(output.ArmorProficiency, dwarvenArmorTraining...)
			ability["strength"] = 2

		}

	case "Elf", "elf":
		output.Name = "elf"
		output.Size = "medium"
		output.Speed = 30
		output.Special = []string{"Darkvision", "Keen Senses", "Fey Ancestry", "Trance"}
		ability["dexterity"] = 2
		output.Language = []string{"common", "elvish"}
		output.Skills = []string{"perception"}
		output.Advantages = []string{"charmed"}
		output.Conditions = []string{"sleep"}
		elfWeaponTraining := []string{"longsword", "shortsword", "shortbow", "longbow"}

		switch subrace {
		case "High Elf", "high-elf", "high", "highelf":
			output.Subrace = "high-elf"
			output.Special = append(output.Special, "Elf Weapon Training")
			output.Special = append(output.Special, "Cantrip")
			output.Special = append(output.Special, "Extra Language")
			ability["intelligence"] = 1
			output.AdditionalLanguages = 1
			output.ArmorProficiency = elfWeaponTraining

		case "Wood Elf", "wood-elf", "woodelf", "wood":
			output.Subrace = "wood-elf"
			output.Special = append(output.Special, "Elf Weapon Training")
			output.Special = append(output.Special, "Fleet of Foot")
			output.Special = append(output.Special, "Mask of the Wild")
			ability["wisdom"] = 1
			output.ArmorProficiency = elfWeaponTraining

		case "drow":
			output.Subrace = "drow"
			ability["charisma"] = 1
			output.Special = append(output.Special, "Darkvision")
			output.Disvantages = append(output.Disvantages, "sunlight-sensitivity")
			output.Special = append(output.Special, "Drow Weapon Training")
			drowWeaponTraining := []string{"rapier", "shortsword", "crossbow-hand", "crossbow-light", "crossbow-heavy"}
			output.ArmorProficiency = drowWeaponTraining
		}

	case "Halfling", "halfling":
		output.Name = "halfling"
		output.Size = "small"
		output.Speed = 25
		output.Special = []string{"Lucky", "Brave", "Halfling Nimbleness"}
		ability["dexterity"] = 2
		output.Language = []string{"common", "halfling"}
		output.Advantages = []string{"frightened"}
		switch subrace {
		case "Lightfoot", "lightfoot", "foot", "light":
			output.Subrace = "lightfoot"
			output.Special = append(output.Special, "Naturally Stealthy")
			ability["charisma"] = 1

		case "Stout", "stout":
			output.Subrace = "stout"
			output.Special = append(output.Special, "Stout Resilience")
			ability["constitution"] = 1
			output.Advantages = append(output.Advantages, "poison")
			output.Resistance = []string{"poison"}
		}

	case "Human", "human":
		output.Name = "human"
		output.Size = "medium"
		output.Speed = 30
		output.Special = []string{"one extra language"}
		ability["strength"] = 1
		ability["dexterity"] = 1
		ability["constitution"] = 1
		ability["intelligence"] = 1
		ability["wisdom"] = 1
		ability["charisma"] = 1
		output.Language = []string{"common"}
		output.AdditionalLanguages = 1

	case "Dragonborn", "dragonborn":
		output.Name = "dragonborn"
		output.Size = "medium"
		output.Speed = 30
		var breathWeapon string
		ability["strength"] = 2
		ability["charisma"] = 1
		output.Language = []string{"common", "draconic"}
		output.Subrace = subrace
		switch subrace {
		case "black":
			output.Resistance = []string{"acid"}
			breathWeapon = "acid 5 by 30 ft. line (Dex. save)"
		case "blue":
			output.Resistance = []string{"cold"}
			breathWeapon = "cold 5 by 30 ft. line (Dex. save)"
		case "brass":
			output.Resistance = []string{"fire"}
			breathWeapon = "fire 5 by 30 ft. line (Dex. save)"
		case "bronze":
			output.Resistance = []string{"lightning"}
			breathWeapon = "lightning 5 by 30 ft. line (Dex. save)"
		case "copper":
			output.Resistance = []string{"acid"}
			breathWeapon = "acid 5 by 30 ft. line (Dex. save)"
		case "gold":
			output.Resistance = []string{"fire"}
			breathWeapon = "fire 1 5 ft. cone (Dex. save)"
		case "green":
			output.Resistance = []string{"poison"}
			breathWeapon = "poison 15 ft. cone (Con. save)"
		case "red":
			output.Resistance = []string{"fire"}
			breathWeapon = "fire 15 ft. cone ( Dex. save)"
		case "silver":
			output.Resistance = []string{"cold"}
			breathWeapon = "cold 15 ft. cone (Con. save)"
		case "white":
			output.Resistance = []string{"cold"}
			breathWeapon = "cold 15 ft. cone (Con. save)"
		}
		output.Special = []string{"Draconic Ancestry", breathWeapon, "Damage Resistance"}

	case "Gnome", "gnome":
		output.Name = "gnome"
		output.Size = "small"
		output.Speed = 25
		output.Special = []string{"Darkvision", "Gnome Cunning"}
		ability["intelligence"] = 2
		output.Advantages = []string{"spell"}
		output.Language = []string{"common", "gnomish"}

		switch subrace {
		case "Rock Gnomes", "rock-gnomes", "rock-gnome", "rock", "rockgnome":
			output.Subrace = "rock-gnome"
			output.Special = append(output.Special, "Artificer’s Lore")
			output.Special = append(output.Special, "Tinker")
			ability["constitution"] = 1

		case "Forest Gnome", "forest-gnome", "forest":
			output.Subrace = "forest-gnome"
			output.Special = append(output.Special, "Natural Illusionist")
			output.Special = append(output.Special, "Speak with Small Beasts")
			ability["dexterity"] = 1

		}

	case "Half-Elf", "half elf", "half-elf":
		output.Name = "half-elf"
		output.Size = "medium"
		output.Speed = 30
		output.Special = []string{"Darkvision", "Fey Ancestry", "Skill Versatility", "two other ability increase by 1", "one extra language"}
		ability["charisma"] = 2
		output.Language = []string{"common", "elvish"}
		output.Advantages = []string{"charmed"}
		output.Conditions = []string{"sleep"}
		output.AdditionalLanguages = 1

	case "Half-Orc", "half orc", "half-orc":
		output.Name = "half-orc"
		output.Size = "medium"
		output.Speed = 30
		output.Special = []string{"Darkvision", "Menacing", "Relentless Endurance", "Savage Attacks"}
		ability["strength"] = 2
		ability["constitution"] = 1
		output.Language = []string{"common", "orc"}
		output.Skills = []string{"intimidation"}

	case "Tiefling", "tiefling":
		output.Name = "tiefling"
		output.Size = "medium"
		output.Speed = 30
		output.Special = []string{"Darkvision", "Hellish Resistance", "Infernal Legacy"}
		ability["intelligence"] = 1
		ability["charisma"] = 2
		output.Language = []string{"common", "infernal"}
		output.Resistance = []string{"fire"}
	}
	output.Ability = ability
	description := RaceTraits(output.Name, output.Subrace)
	output.Description = description["link"]
	return output
}

//RaceSpecialTrait func
func RaceSpecialTrait(race, subrace string, level int, ability map[string]int) (name string, spellList []string, spellcastAbility, damageDice, damageType, savingThrow, description string, difficultClass int) {
	switch race {
	case "dragonborn":
		damageDice = "2d6"
		if level > 5 {
			damageDice = "3d6"
		}
		if level > 10 {
			damageDice = "4d6"
		}
		if level > 15 {
			damageDice = "5d6"
		}
		var breathWeapon string
		prof := CalcProficiency(level)
		modifier := CalcAbilityModifier(ability["constitution"])
		difficultClass = 8 + prof + modifier
		switch subrace {
		case "black":
			damageType = "acid"
			breathWeapon = "acid 5 by 30 ft. line (Dex. save)"
			savingThrow = "dexterity"
		case "blue":
			damageType = "cold"
			breathWeapon = "cold 5 by 30 ft. line (Dex. save)"
			savingThrow = "dexterity"
		case "brass":
			damageType = "fire"
			breathWeapon = "fire 5 by 30 ft. line (Dex. save)"
			savingThrow = "dexterity"
		case "bronze":
			damageType = "lightning"
			breathWeapon = "lightning 5 by 30 ft. line (Dex. save)"
			savingThrow = "dexterity"
		case "copper":
			damageType = "acid"
			breathWeapon = "acid 5 by 30 ft. line (Dex. save)"
			savingThrow = "dexterity"
		case "gold":
			damageType = "fire"
			breathWeapon = "fire 1 5 ft. cone (Dex. save)"
			savingThrow = "dexterity"
		case "green":
			damageType = "poison"
			breathWeapon = "poison 15 ft. cone (Con. save)"
			savingThrow = "constitution"
		case "red":
			damageType = "fire"
			breathWeapon = "fire 15 ft. cone (Dex. save)"
			savingThrow = "dexterity"
		case "silver":
			damageType = "cold"
			breathWeapon = "cold 15 ft. cone (Con. save)"
			savingThrow = "constitution"
		case "white":
			damageType = "cold"
			breathWeapon = "cold 15 ft. cone (Con. save)"
			savingThrow = "constitution"
		}
		description = breathWeapon
		name = "breath-weapon"

	case "tiefling":
		spellList = []string{"thaumaturgy"}
		if level > 2 {
			spellList = []string{"thaumaturgy", "hellish-rebuke"}
		}
		if level > 4 {
			spellList = []string{"thaumaturgy", "hellish-rebuke", "darkness"}
		}
		description = `You know the thaumaturgy cantrip. When you reach 3rd level, you can cast the hellish rebuke spell as a 2nd-level spell once with this trait and regain the ability to do so when you finish a long rest. When you reach 5th level, you can cast the darkness spell once with this trait and regain the ability to
		do so when you finish a long rest. Charisma is your spellcasting ability for these spells.`
		spellcastAbility = "charisma"
		prof := CalcProficiency(level)
		modifier := CalcAbilityModifier(ability["charisma"])
		difficultClass = 8 + prof + modifier
		name = "infernal-legacy"

	case "gnome":
		if subrace == "forest-gnome" {
			spellList = []string{"minor-illusion"}
			description = "You know the minor illusion cantrip. Intelligence is your spellcasting ability for it."
			spellcastAbility = "intelligence"
			prof := CalcProficiency(level)
			modifier := CalcAbilityModifier(ability["intelligence"])
			difficultClass = 8 + prof + modifier
			name = "natural-illusionist"
		}

	case "elf":
		if subrace == "drow" {
			spellList = []string{"dancing-lights"}
			if level > 2 {
				spellList = []string{"dancing-lights", "faerie-fire"}
			}
			if level > 4 {
				spellList = []string{"dancing-lights", "faerie-fire", "darkness"}
			}
			description = "You know the dancing lights cantrip. When you reach 3rd level, you can cast the faerie fire spell once with this trait and regain the ability to do so when you finish a long rest. When you reach 5th level, you can cast the darkness spell once with this trait and regain the ability to do so when you finish a long rest. Charisma is your spellcasting ability for these spells."
			spellcastAbility = "charisma"
			prof := CalcProficiency(level)
			modifier := CalcAbilityModifier(ability["charisma"])
			difficultClass = 8 + prof + modifier
			name = "drow-magic"
		}
	}

	return name, spellList, spellcastAbility, damageDice, damageType, savingThrow, description, difficultClass
}

//BackgroundStatistics func
func BackgroundStatistics(name string) *rule.Background {
	back := new(rule.Background)
	switch name {
	case "Acolyte", "acolyte":
		back.Name = "acolyte"
		back.Title = "Acolyte"
		back.Skills = []string{"insight", "religion"}
		back.AdditionalLanguages = 2
		back.Description = "https://www.dndbeyond.com/sources/basic-rules/personality-and-background#Acolyte"
	case "Criminal", "criminal":
		back.Name = "criminal"
		back.Title = "Criminal"
		back.Skills = []string{"deception", "stealth"}
		back.Extra = "One type of gaming set, thieves’ tools"
		back.Description = "https://www.dndbeyond.com/sources/basic-rules/personality-and-background#Criminal"
	case "Folk Hero", "folk hero", "folk-hero":
		back.Name = "folk-hero"
		back.Title = "Folk Hero"
		back.Extra = "One type of artisan’s tools, vehicles (land)"
		back.Skills = []string{"animal handling", "survival"}
		back.Description = "https://www.dndbeyond.com/sources/basic-rules/personality-and-background#FolkHero"
	case "Noble", "noble":
		back.Name = "noble"
		back.Title = "Noble"
		back.Skills = []string{"history", "persuasion"}
		back.Extra = "One type of gaming set"
		back.Description = "https://www.dndbeyond.com/sources/basic-rules/personality-and-background#Noble"
	case "Sage", "sage":
		back.Name = "sage"
		back.Title = "Sage"
		back.Skills = []string{"arcana", "history"}
		back.AdditionalLanguages = 2
		back.Description = "https://www.dndbeyond.com/sources/basic-rules/personality-and-background#Sage"
	case "Soldier", "soldier":
		back.Name = "soldier"
		back.Title = "Soldier"
		back.Extra = "One type of gaming set, vehicles (land)"
		back.Skills = []string{"athletics", "intimidation"}
		back.Description = "https://www.dndbeyond.com/sources/basic-rules/personality-and-background#Soldier"
	}
	return back
}

//CalculateSpellList func
func CalculateSpellList(class string, level int) (map[string][]int, int) {
	spellList := make(map[string][]int)
	// for i := level; i > 0; i-- {
	list := SpellsPerLevel(class, level)
	levelName := fmt.Sprintf("Spell Level 1 to %v", len(list))
	// spellList[levelName] = SpellsPerLevel(class, i)
	// }
	spellList[levelName] = list
	max := len(list)
	return spellList, max
}

//CalculateClassFeatureList func
func CalculateClassFeatureList(classFeatures map[int][]string, level int, choosen []string) (features []string) {
	for i := level; i > 0; i-- {
		for _, v := range classFeatures[i] {
			if v != "" {
				value := v
				var values []string
				if utils.StringInSlice(value, featuresListRename()) {
					for k, f := range choosen {
						if k != 0 && value == "fighting-style" {
							continue
						}
						feats := choosenClassFeatures(value, f, i)
						if len(feats) != 0 {
							values = append(values, feats...)
						}
					}

				}
				features = append(features, value)
				features = append(features, values...)
			}
		}
	}
	// more details

	return features
}

//SpellSlotsMultiClass func
func SpellSlotsMultiClass(class string, level int) int {
	switch class {
	case "bard":
		return level
	case "cleric":
		return level
	case "druid":
		return level
	case "paladin":
		return int(math.Floor(float64(level) / 2))
	case "ranger":
		return int(math.Floor(float64(level) / 2))
	case "sorcerer":
		return level
	case "warlock":
		return 0
	case "wizard":
		return level
	default:
		return 0
	}
}

// ClassStatistics func
func ClassStatistics(name string, level int) rule.Class {
	var class rule.Class
	class.Name = name
	class.Title = strings.Title(name)
	class.HitDice, class.Savings, class.ArmorProficiency, class.SkillNumber = ClassDetails(name)
	class.SkillList = skillListByClass(name)

	class.Features = make(map[int][]string)
	class.SpellKnown = make(map[int]int)
	class.CantripsKnown = make(map[int]int)
	class.SpellsPerLevel = make(map[int][]int)
	class.BarbarianRage = make(map[int]int)
	class.BarbarianRageDamage = make(map[int]int)
	class.MonkMartial = make(map[int]string)
	class.MonkKi = make(map[int]int)
	class.MonkMovement = make(map[int]string)
	class.RogueSneakAttack = make(map[int]string)
	class.WarlockSpellSlots = make(map[int]int)
	class.WarlockSlotLevel = make(map[int]string)
	class.WarlockInvocationsKnown = make(map[int]int)
	for i := 1; i <= level; i++ {

		class.Features[i] = ClassFeatures(name, i)
		if utils.StringInSlice(name, ClassWithCantrips()) {
			class.CantripsKnown[i] = CantripsKnown(name, i)
		}
		if utils.StringInSlice(name, ClassWithSpellKnown()) {
			class.SpellKnown[i] = SpellKnown(name, i)
		}
		if utils.StringInSlice(name, ClassWithSpell()) {
			if class.Name != "warlock" {
				class.SpellsPerLevel[i] = SpellsPerLevel(name, i)
			}
			class.AbilityForSpell = AbilityForSpell(name)
		}
		if name == "barbarian" {
			class.BarbarianRage[i], class.BarbarianRageDamage[i] = BarbarianClass(i)
		}
		if name == "monk" {
			class.MonkMartial[i], class.MonkKi[i], class.MonkMovement[i] = MonkClass(i)
		}
		if name == "rogue" {
			class.RogueSneakAttack[i] = RogueClass(i)
		}
		if name == "warlock" {
			class.WarlockSpellSlots[i], class.WarlockSlotLevel[i], class.WarlockInvocationsKnown[i] = WarlockClass(i)
		}
	}
	description := ClassInfo(name)
	class.Description = description["link"]

	return class
}

// SpellsPerLevel func
func SpellsPerLevel(class string, level int) []int {
	var value []int
	switch level {
	case 1:
		switch class {
		case "wizard", "cleric":
			value = []int{2}
		case "sorcerer":
			value = []int{2}
		case "druid", "bard":
			value = []int{2}
		case "eldritch-knight":
			value = []int{0}
		case "arcane-trickster":
			value = []int{0}
		case "paladin", "ranger":
			value = []int{0}
		}
	case 2:
		switch class {
		case "wizard", "cleric":
			value = []int{3}
		case "sorcerer":
			value = []int{3}
		case "druid", "bard":
			value = []int{3}
		case "eldritch-knight":
			value = []int{0}
		case "arcane-trickster":
			value = []int{0}
		case "paladin", "ranger":
			value = []int{2}
		}
	case 3:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 2}
		case "sorcerer":
			value = []int{4, 2}
		case "druid", "bard":
			value = []int{4, 2}
		case "eldritch-knight":
			value = []int{2}
		case "arcane-trickster":
			value = []int{2}
		case "paladin", "ranger":
			value = []int{3}
		}
	case 4:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3}
		case "sorcerer":
			value = []int{4, 3}
		case "druid", "bard":
			value = []int{4, 3}
		case "eldritch-knight":
			value = []int{3}
		case "arcane-trickster":
			value = []int{3}
		case "paladin", "ranger":
			value = []int{3}
		}
	case 5:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 2}
		case "sorcerer":
			value = []int{4, 3, 2}
		case "druid", "bard":
			value = []int{4, 3, 2}
		case "eldritch-knight":
			value = []int{3}
		case "arcane-trickster":
			value = []int{3}
		case "paladin", "ranger":
			value = []int{4, 2}
		}
	case 6:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3}
		case "sorcerer":
			value = []int{4, 3, 3}
		case "druid", "bard":
			value = []int{4, 3, 3}
		case "eldritch-knight":
			value = []int{3}
		case "arcane-trickster":
			value = []int{3}
		case "paladin", "ranger":
			value = []int{4, 2}
		}
	case 7:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 1}
		case "eldritch-knight":
			value = []int{4, 2}
		case "arcane-trickster":
			value = []int{4, 2}
		case "paladin", "ranger":
			value = []int{4, 3}
		}
	case 8:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 2}
		case "sorcerer":
			value = []int{4, 3, 3, 2}
		case "druid", "bard":
			value = []int{4, 3, 3, 2}
		case "eldritch-knight":
			value = []int{4, 2}
		case "arcane-trickster":
			value = []int{4, 2}
		case "paladin", "ranger":
			value = []int{4, 3}
		}
	case 9:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 1}
		case "eldritch-knight":
			value = []int{4, 2}
		case "arcane-trickster":
			value = []int{4, 2}
		case "paladin", "ranger":
			value = []int{4, 3, 2}
		}
	case 10:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 2}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 2}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 2}
		case "eldritch-knight":
			value = []int{4, 3}
		case "arcane-trickster":
			value = []int{4, 2}
		case "paladin", "ranger":
			value = []int{4, 3, 2}
		}
	case 11:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 2, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 2, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 2, 1}
		case "eldritch-knight":
			value = []int{4, 3}
		case "arcane-trickster":
			value = []int{4, 2}
		case "paladin", "ranger":
			value = []int{4, 3, 3}
		}
	case 12:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 2, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 2, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 2, 1}
		case "eldritch-knight":
			value = []int{4, 3}
		case "arcane-trickster":
			value = []int{4, 2}
		case "paladin", "ranger":
			value = []int{4, 3, 3}
		}
	case 13:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 2, 1, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 2, 1, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 2, 1, 1}
		case "eldritch-knight":
			value = []int{4, 3, 2}
		case "arcane-trickster":
			value = []int{4, 3, 2}
		case "paladin", "ranger":
			value = []int{4, 3, 3, 1}
		}
	case 14:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 2, 1, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 2, 1, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 2, 1, 1}
		case "eldritch-knight":
			value = []int{4, 3, 2}
		case "arcane-trickster":
			value = []int{4, 3, 2}
		case "paladin", "ranger":
			value = []int{4, 3, 3, 1}
		}

	case 15:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1}
		case "eldritch-knight":
			value = []int{4, 3, 2}
		case "arcane-trickster":
			value = []int{4, 3, 2}
		case "paladin", "ranger":
			value = []int{4, 3, 3, 2}
		}
	case 16:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1}
		case "eldritch-knight":
			value = []int{4, 3, 3}
		case "arcane-trickster":
			value = []int{4, 3, 3}
		case "paladin", "ranger":
			value = []int{4, 3, 3, 2}
		}
	case 17:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 2, 1, 1, 1, 1}
		case "eldritch-knight":
			value = []int{4, 3, 3}
		case "arcane-trickster":
			value = []int{4, 3, 3}
		case "paladin", "ranger":
			value = []int{4, 3, 3, 3, 1}
		}
	case 18:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 3, 1, 1, 1, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 3, 1, 1, 1, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 3, 1, 1, 1, 1}
		case "eldritch-knight":
			value = []int{4, 3, 3}
		case "arcane-trickster":
			value = []int{4, 3, 3}
		case "paladin", "ranger":
			value = []int{4, 3, 3, 3, 1}
		}
	case 19:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 3, 2, 1, 1, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 3, 2, 1, 1, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 3, 2, 1, 1, 1}
		case "eldritch-knight":
			value = []int{4, 3, 3, 1}
		case "arcane-trickster":
			value = []int{4, 3, 3, 1}
		case "paladin", "ranger":
			value = []int{4, 3, 3, 3, 2}
		}
	case 20:
		switch class {
		case "wizard", "cleric":
			value = []int{4, 3, 3, 3, 3, 2, 2, 1, 1}
		case "sorcerer":
			value = []int{4, 3, 3, 3, 3, 2, 2, 1, 1}
		case "druid", "bard":
			value = []int{4, 3, 3, 3, 3, 2, 2, 1, 1}
		case "eldritch-knight":
			value = []int{4, 3, 3, 1}
		case "arcane-trickster":
			value = []int{4, 3, 3, 1}
		case "paladin", "ranger":
			value = []int{4, 3, 3, 3, 2}
		}
	}
	return value
}

// ClassDetails func
func ClassDetails(class string) (hitDice int, savings []string, armorProficiency []string, skillNumber int) {
	switch class {
	case "barbarian":
		hitDice = 12
		savings = []string{"strength", "constitution"}
		armorProficiency = []string{"light-armor", "medium-armor", "shields", "simple-weapon", "martial-weapon"}
		skillNumber = 2
	case "bard":
		hitDice = 8
		savings = []string{"dexterity", "charisma"}
		armorProficiency = []string{"light-armor", "simple-weapon", "crossbow hand", "longsword", "rapier", "shortsword"}
		skillNumber = 3
	case "cleric":
		hitDice = 8
		savings = []string{"wisdom", "charisma"}
		armorProficiency = []string{"light-armor", "medium-armor", "shields", "simple-weapon"}
		skillNumber = 2

	case "druid":
		hitDice = 8
		savings = []string{"intelligence", "wisdom"}
		armorProficiency = []string{"light-armor", "medium-armor", "shield", "club", "dagger", "dart", "javelin", "mace", "quarterstaff", "scimitar", "sickles", "sling", "spear"}
		skillNumber = 2

	case "fighter":
		hitDice = 10
		savings = []string{"strength", "constitution"}
		armorProficiency = []string{"light-armor", "medium-armor", "heavy-armor", "shields", "simple-weapon", "martial-weapon"}
		skillNumber = 2

	case "monk":
		hitDice = 8
		savings = []string{"dexterity", "wisdom"}
		armorProficiency = []string{"shortswords", "simple-weapon"}
		skillNumber = 2

	case "paladin":
		hitDice = 10
		savings = []string{"strength", "charisma"}
		armorProficiency = []string{"light-armor", "medium-armor", "heavy-armor", "shields", "simple-weapon", "martial-weapon"}
		skillNumber = 2

	case "ranger":
		hitDice = 10
		savings = []string{"dexterity", "wisdom"}
		armorProficiency = []string{"light-armor", "medium-armor", "shield", "simple-weapon", "martial-weapon"}
		skillNumber = 3

	case "rogue":
		hitDice = 8
		savings = []string{"dexterity", "intelligence"}
		armorProficiency = []string{"light-armor", "simple-weapon", "crossbow hand", "longsword", "rapier", "shortsword"}
		skillNumber = 4

	case "sorcerer":
		hitDice = 6
		savings = []string{"charisma", "constitution"}
		armorProficiency = []string{"daggers", "darts", "slings", "quarterstaff", "crossbow light"}
		skillNumber = 2

	case "warlock":
		hitDice = 8
		savings = []string{"wisdom", "charisma"}
		armorProficiency = []string{"light-armor", "simple-weapons"}
		skillNumber = 2

	case "wizard":
		hitDice = 6
		savings = []string{"intelligence", "wisdom"}
		armorProficiency = []string{"dagger", "dart", "sling", "quarterstaff", "crossbow light"}
		skillNumber = 2

	}
	return hitDice, savings, armorProficiency, skillNumber
}

//ClassFeatures func
func ClassFeatures(class string, level int) (features []string) {
	switch level {
	case 1:
		switch class {
		case "barbarian":
			features = []string{"rage", "unarmored-defense-barbarian"}
		case "bard":
			features = []string{"spellcasting", "bardic-inspiration"}
		case "cleric":
			features = []string{"spellcasting", "divine-domain"}

		case "druid":
			features = []string{"druidic", "spellcasting"}

		case "fighter":
			features = []string{"fighting-style", "second-wind"}

		case "monk":
			features = []string{"unarmored-defense-monk", "martial-arts"}

		case "paladin":
			features = []string{"divine-sense", "lay-on-hands"}

		case "ranger":
			features = []string{"favored-enemy", "natural-explorer"}

		case "rogue":
			features = []string{"expertise", "sneak-attack", "thieves-cant"}

		case "sorcerer":
			features = []string{"spellcasting", "sorcerous-origin"}

		case "warlock":
			features = []string{"otherworldly-patron", "pact-magic"}

		case "wizard":
			features = []string{"spellcasting", "arcane-recovery"}

		}
	case 2:
		switch class {
		case "barbarian":
			features = []string{"reckless-attack", "danger-sense"}
		case "bard":
			features = []string{"jack-of-all-trades", "song-of-rest"}
		case "cleric":
			features = []string{"channel-divinity", "divine-domain-feature"}

		case "druid":
			features = []string{"wild-shape", "druid-circle"}

		case "fighter":
			features = []string{"action-surge-1"}

		case "monk":
			features = []string{"ki", "unarmored-movement"}

		case "paladin":
			features = []string{"fighting-style", "spellcasting", "divine-smite"}

		case "ranger":
			features = []string{"fighting-style", "spellcasting"}

		case "rogue":
			features = []string{"cunning-action"}

		case "sorcerer":
			features = []string{"font-of-magic"}

		case "warlock":
			features = []string{"eldritch-invocations"}

		case "wizard":
			features = []string{"arcane-tradition"}

		}
	case 3:
		switch class {
		case "barbarian":
			features = []string{"primal-path"}
		case "bard":
			features = []string{"bard-college", "expertise"}
		case "cleric":
			features = []string{}

		case "druid":
			features = []string{}

		case "fighter":
			features = []string{"martial-archetype"}

		case "monk":
			features = []string{"monastic-tradition", "deflect-missiles"}

		case "paladin":
			features = []string{"divine-health", "sacred-oath"}

		case "ranger":
			features = []string{"ranger-archetype", "primeval-awareness"}

		case "rogue":
			features = []string{"roguish-archetype"}

		case "sorcerer":
			features = []string{"metamagic"}

		case "warlock":
			features = []string{"pact-boon"}

		case "wizard":
			features = []string{}

		}
	case 4:
		switch class {
		case "barbarian":
			features = []string{"ability-score-improvement"}
		case "bard":
			features = []string{"ability-score-improvement"}
		case "cleric":
			features = []string{"ability-score-improvement"}

		case "druid":
			features = []string{"wild-shape-improvement", "ability-score-improvement"}

		case "fighter":
			features = []string{"ability-score-improvement"}

		case "monk":
			features = []string{"ability-score-improvement", "slow-fall"}

		case "paladin":
			features = []string{"ability-score-improvement"}

		case "ranger":
			features = []string{"ability-score-improvement"}

		case "rogue":
			features = []string{"ability-score-improvement"}

		case "sorcerer":
			features = []string{"ability-score-improvement"}

		case "warlock":
			features = []string{"ability-score-improvement"}

		case "wizard":
			features = []string{"ability-score-improvement"}

		}
	case 5:
		switch class {
		case "barbarian":
			features = []string{"extra-attack-1", "fast-movement"}
		case "bard":
			features = []string{"bardic-inspiration", "font-of-inspiration"}
		case "cleric":
			features = []string{"destroy-undead"}

		case "druid":
			features = []string{}

		case "fighter":
			features = []string{"extra-attack-1"}

		case "monk":
			features = []string{"extra-attack-1", "stunning-strike"}

		case "paladin":
			features = []string{"extra-attack-1"}

		case "ranger":
			features = []string{"extra-attack-1"}

		case "rogue":
			features = []string{"uncanny-dodge"}

		case "sorcerer":
			features = []string{}

		case "warlock":
			features = []string{}

		case "wizard":
			features = []string{}

		}
	case 6:
		switch class {
		case "barbarian":
			features = []string{"path-feature"}
		case "bard":
			features = []string{"countercharm", "bard-college-feature"}
		case "cleric":
			features = []string{"channel-divinity", "divine-domain-feature"}

		case "druid":
			features = []string{"druid-circle-feature"}

		case "fighter":
			features = []string{"ability-score-improvement"}

		case "monk":
			features = []string{"ki-empowered-strikes", "monastic-tradition-feature"}

		case "paladin":
			features = []string{"aura-of-protection"}

		case "ranger":
			features = []string{"favored-enemy-improvement", "natural-explorer-improvement"}

		case "rogue":
			features = []string{"expertise"}

		case "sorcerer":
			features = []string{"sorcerous-origin-feature"}

		case "warlock":
			features = []string{"otherworldly-patron-feature"}

		case "wizard":
			features = []string{"arcane-tradition-feature"}

		}
	case 7:
		switch class {
		case "barbarian":
			features = []string{"feral-instinct"}
		case "bard":
			features = []string{}
		case "cleric":
			features = []string{}

		case "druid":
			features = []string{}

		case "fighter":
			features = []string{"martial-archetype-feature"}

		case "monk":
			features = []string{"evasion", "stillness-of-mind"}

		case "paladin":
			features = []string{"sacred-oath-feature"}

		case "ranger":
			features = []string{"ranger-archetype-feature"}

		case "rogue":
			features = []string{"evasion"}

		case "sorcerer":
			features = []string{}

		case "warlock":
			features = []string{}

		case "wizard":
			features = []string{}

		}
	case 8:
		switch class {
		case "barbarian":
			features = []string{"ability-score-improvement"}
		case "bard":
			features = []string{"ability-score-improvement"}
		case "cleric":
			features = []string{"ability-score-improvement", "destroy-undead", "divine-domain-feature"}

		case "druid":
			features = []string{"wild-shape-improvement", "ability-score-improvement"}

		case "fighter":
			features = []string{"ability-score-improvement"}

		case "monk":
			features = []string{"ability-score-improvement"}

		case "paladin":
			features = []string{"ability-score-improvement"}

		case "ranger":
			features = []string{"ability-score-improvement", "lands-stride"}

		case "rogue":
			features = []string{"ability-score-improvement"}

		case "sorcerer":
			features = []string{"ability-score-improvement"}

		case "warlock":
			features = []string{"ability-score-improvement"}

		case "wizard":
			features = []string{"ability-score-improvement"}

		}
	case 9:
		switch class {
		case "barbarian":
			features = []string{"brutal-critical-1"}
		case "bard":
			features = []string{"song-of-rest"}
		case "cleric":
			features = []string{}

		case "druid":
			features = []string{}

		case "fighter":
			features = []string{"indomitable-1"}

		case "monk":
			features = []string{"unarmored-movement-improvement"}

		case "paladin":
			features = []string{}

		case "ranger":
			features = []string{}

		case "rogue":
			features = []string{"roguish-archetype-feature"}

		case "sorcerer":
			features = []string{}

		case "warlock":
			features = []string{}

		case "wizard":
			features = []string{}

		}
	case 10:
		switch class {
		case "barbarian":
			features = []string{"path-feature"}
		case "bard":
			features = []string{"bardic-inspiration", "expertise", "magical-secrets"}
		case "cleric":
			features = []string{"divine-intervention"}

		case "druid":
			features = []string{"druid-circle-feature"}

		case "fighter":
			features = []string{"martial-archetype-feature"}

		case "monk":
			features = []string{"purity-of-body"}

		case "paladin":
			features = []string{"aura-of-courage"}

		case "ranger":
			features = []string{"natural-explorer-improvement", "hide-in-plain-sight"}

		case "rogue":
			features = []string{"ability-score-improvement"}

		case "sorcerer":
			features = []string{"metamagic"}

		case "warlock":
			features = []string{"otherworldly-patron-feature"}

		case "wizard":
			features = []string{"arcane-tradition-feature"}

		}
	case 11:
		switch class {
		case "barbarian":
			features = []string{"relentless-rage"}
		case "bard":
			features = []string{}
		case "cleric":
			features = []string{"destroy-undead"}

		case "druid":
			features = []string{}

		case "fighter":
			features = []string{"extra-attack-2"}

		case "monk":
			features = []string{"monastic-tradition-feature"}

		case "paladin":
			features = []string{"improved-divine-smite"}

		case "ranger":
			features = []string{"ranger-archetype-feature"}

		case "rogue":
			features = []string{"reliable-talent"}

		case "sorcerer":
			features = []string{}

		case "warlock":
			features = []string{"mystic-arcanum-6"}

		case "wizard":
			features = []string{}

		}
	case 12:
		switch class {
		case "barbarian":
			features = []string{"ability-score-improvement"}
		case "bard":
			features = []string{"ability-score-improvement"}
		case "cleric":
			features = []string{"ability-score-improvement"}

		case "druid":
			features = []string{"ability-score-improvement"}

		case "fighter":
			features = []string{"ability-score-improvement"}

		case "monk":
			features = []string{"ability-score-improvement"}

		case "paladin":
			features = []string{"ability-score-improvement"}

		case "ranger":
			features = []string{"ability-score-improvement"}

		case "rogue":
			features = []string{"ability-score-improvement"}

		case "sorcerer":
			features = []string{"ability-score-improvement"}

		case "warlock":
			features = []string{"ability-score-improvement"}

		case "wizard":
			features = []string{"ability-score-improvement"}

		}
	case 13:
		switch class {
		case "barbarian":
			features = []string{"brutal-critical-2"}
		case "bard":
			features = []string{"song-of-rest"}
		case "cleric":
			features = []string{}

		case "druid":
			features = []string{}

		case "fighter":
			features = []string{"indomitable-2"}

		case "monk":
			features = []string{"tongue-of-the-sun-and-moon"}

		case "paladin":
			features = []string{}

		case "ranger":
			features = []string{}

		case "rogue":
			features = []string{"roguish-archetype-feature"}

		case "sorcerer":
			features = []string{}

		case "warlock":
			features = []string{"mystic-arcanum-7"}

		case "wizard":
			features = []string{}

		}
	case 14:
		switch class {
		case "barbarian":
			features = []string{"path-feature"}
		case "bard":
			features = []string{"magical-secrets", "bard-college-feature"}
		case "cleric":
			features = []string{"destroy-undead"}

		case "druid":
			features = []string{"druid-circle-feature"}

		case "fighter":
			features = []string{"ability-score-improvement"}

		case "monk":
			features = []string{"diamond-soul"}

		case "paladin":
			features = []string{"cleansing-touch"}

		case "ranger":
			features = []string{"favored-enemy-improvement", "vanish"}

		case "rogue":
			features = []string{"blindsense"}

		case "sorcerer":
			features = []string{"sorcerous-origin-feature"}

		case "warlock":
			features = []string{"otherworldly-patron-feature"}

		case "wizard":
			features = []string{"arcane-tradition-feature"}

		}
	case 15:
		switch class {
		case "barbarian":
			features = []string{"persistent-rage"}
		case "bard":
			features = []string{"bardic-inspiration"}
		case "cleric":
			features = []string{}

		case "druid":
			features = []string{}

		case "fighter":
			features = []string{"martial-archetype-feature"}

		case "monk":
			features = []string{"timeless-body"}

		case "paladin":
			features = []string{"sacred-oath-feature"}

		case "ranger":
			features = []string{"ranger-archetype-feature"}

		case "rogue":
			features = []string{"slippery-mind"}

		case "sorcerer":
			features = []string{}

		case "warlock":
			features = []string{"mystic-arcanum-8"}

		case "wizard":
			features = []string{}

		}
	case 16:
		switch class {
		case "barbarian":
			features = []string{"ability-score-improvement"}
		case "bard":
			features = []string{"ability-score-improvement"}
		case "cleric":
			features = []string{"ability-score-improvement"}

		case "druid":
			features = []string{"ability-score-improvement"}

		case "fighter":
			features = []string{"ability-score-improvement"}

		case "monk":
			features = []string{"ability-score-improvement"}

		case "paladin":
			features = []string{"ability-score-improvement"}

		case "ranger":
			features = []string{"ability-score-improvement"}

		case "rogue":
			features = []string{"ability-score-improvement"}

		case "sorcerer":
			features = []string{"ability-score-improvement"}

		case "warlock":
			features = []string{"ability-score-improvement"}

		case "wizard":
			features = []string{"ability-score-improvement"}

		}
	case 17:
		switch class {
		case "barbarian":
			features = []string{"brutal-critical-3"}
		case "bard":
			features = []string{"song-of-rest"}
		case "cleric":
			features = []string{"destroy-undead", "divine-domain-feature"}

		case "druid":
			features = []string{}

		case "fighter":
			features = []string{"action-surge-2", "indomitable-3"}

		case "monk":
			features = []string{"monastic-tradition-feature"}

		case "paladin":
			features = []string{}

		case "ranger":
			features = []string{}

		case "rogue":
			features = []string{"roguish-archetype-feature"}

		case "sorcerer":
			features = []string{"metamagic"}

		case "warlock":
			features = []string{"mystic-arcanum-9"}

		case "wizard":
			features = []string{}

		}
	case 18:
		switch class {
		case "barbarian":
			features = []string{"indomitable-might"}
		case "bard":
			features = []string{"magical-secrets"}
		case "cleric":
			features = []string{"channel-divinity"}

		case "druid":
			features = []string{"timeless-body", "beast-spells"}

		case "fighter":
			features = []string{"martial-archetype-feature"}

		case "monk":
			features = []string{"empty-body"}

		case "paladin":
			features = []string{"aura-improvements"}

		case "ranger":
			features = []string{"feral-senses"}

		case "rogue":
			features = []string{"elusive"}

		case "sorcerer":
			features = []string{"sorcerous-origin-feature"}

		case "warlock":
			features = []string{}

		case "wizard":
			features = []string{"spell-mastery"}

		}
	case 19:
		switch class {
		case "barbarian":
			features = []string{"ability-score-improvement"}
		case "bard":
			features = []string{"ability-score-improvement"}
		case "cleric":
			features = []string{"ability-score-improvement"}

		case "druid":
			features = []string{"ability-score-improvement"}

		case "fighter":
			features = []string{"ability-score-improvement"}

		case "monk":
			features = []string{"ability-score-improvement"}

		case "paladin":
			features = []string{"ability-score-improvement"}

		case "ranger":
			features = []string{"ability-score-improvement"}

		case "rogue":
			features = []string{"ability-score-improvement"}

		case "sorcerer":
			features = []string{"ability-score-improvement"}

		case "warlock":
			features = []string{"ability-score-improvement"}

		case "wizard":
			features = []string{"ability-score-improvement"}

		}
	case 20:
		switch class {
		case "barbarian":
			features = []string{"primal-champion"}
		case "bard":
			features = []string{"superior-inspiration"}
		case "cleric":
			features = []string{"divine-intervention-improvement"}

		case "druid":
			features = []string{"archdruid"}

		case "fighter":
			features = []string{"extra-attack-3"}

		case "monk":
			features = []string{"perfect-self"}

		case "paladin":
			features = []string{"sacred-oath-feature"}

		case "ranger":
			features = []string{"foe-slayer"}

		case "rogue":
			features = []string{"stroke-of-luck"}

		case "sorcerer":
			features = []string{"sorcerous-restoration"}

		case "warlock":
			features = []string{"eldritch-master"}

		case "wizard":
			features = []string{"signature-spells"}

		}
	}

	return features
}

//SpellKnown func return a int
func SpellKnown(class string, level int) int {
	var list []int
	switch class {
	case "bard":
		list = []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 15, 16, 18, 19, 19, 20, 22, 22, 22}
	case "ranger":
		list = []int{0, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11}

	case "sorcerer":
		list = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12, 13, 13, 14, 14, 15, 15, 15, 15}

	case "warlock":
		list = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15}

	case "arcane-trickster", "eldritch-knight":
		list = []int{0, 0, 3, 4, 4, 4, 5, 6, 6, 7, 8, 8, 9, 10, 10, 11, 11, 11, 12, 13}

	}
	return list[level-1]
}

//CantripsKnown func return a int
func CantripsKnown(class string, level int) int {
	var list []int
	switch class {
	case "bard":
		list = []int{2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	case "cleric":
		list = []int{3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	case "druid":
		list = []int{2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}

	case "sorcerer":
		list = []int{4, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}

	case "warlock":
		list = []int{2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}

	case "wizard":
		list = []int{3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	case "eldritch-knight":
		list = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}

	case "arcane-trickster":
		list = []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}

	}
	return list[level-1]
}

//BarbarianClass func
func BarbarianClass(level int) (int, int) {
	rage := []int{2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 99}
	rageDamage := []int{2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4}
	if level < 1 {
		return 0, 0
	}
	return rage[level-1], rageDamage[level-1]
}

//MonkClass func
func MonkClass(level int) (string, int, string) {
	martial := []string{"1d4", "1d4", "1d4", "1d4", "1d6", "1d6", "1d6", "1d6", "1d6", "1d6", "1d8", "1d8", "1d8", "1d8", "1d8", "1d8", "1d10", "1d10", "1d10", "1d10"}
	ki := []int{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	movement := []string{"0 ft.", "10 ft.", "10 ft.", "10 ft.", "10 ft.", "15 ft.", "15 ft.", "15 ft.", "15 ft.", "20 ft.", "20 ft.", "20 ft.", "20 ft.", "25 ft.", "25 ft.", "25 ft.", "25 ft.", "30 ft.", "30 ft.", "30 ft."}
	if level < 1 {
		return "", 0, ""
	}
	return martial[level-1], ki[level-1], movement[level-1]
}

//RogueClass func
func RogueClass(level int) string {
	sneakAttack := []string{"1d6", "1d6", "2d6", "2d6", "3d6", "3d6", "4d6", "4d6", "5d6", "5d6", "6d6", "6d6", "7d6", "7d6", "8d6", "8d6", "9d6", "9d6", "10d6", "10d6"}
	if level < 1 {
		return ""
	}
	return sneakAttack[level-1]
}

//WarlockClass func
func WarlockClass(level int) (int, string, int) {
	spellSlots := []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4}
	slotLevel := []string{"1st", "1st", "2nd", "2nd", "3rd", "3rd", "4th", "4th", "5th", "5th", "5th", "5th", "5th", "5th", "5th", "5th", "5th", "5th", "5th", "5th"}
	invocationsKnown := []int{0, 2, 2, 2, 3, 3, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8}
	if level < 1 {
		return 0, "", 0
	}
	return spellSlots[level-1], slotLevel[level-1], invocationsKnown[level-1]
}

// CoinList return a list of coins in []string
func CoinList() []string {
	return []string{"copper", "silver", "electrum", "gold", "platinum"}
}

// CoinShortnameList return all coins shortname in []string
func CoinShortnameList() []string {
	return []string{"cp", "sp", "ep", "gp", "pp"}
}

// CoinShortName return a coin name based on it shortname
func CoinShortName(value string) string {
	switch value {
	case "cp":
		return "copper"
	case "sp":
		return "silver"
	case "ep":
		return "electrum"
	case "gp":
		return "gold"
	case "pp":
		return "platinum"
	default:
		return "unknown"
	}
}

// ExchangeRates returns a int from a exchange rate from coin input to coin output
func ExchangeRates(i, o string, value int) (int, error) {
	switch i {
	case "platinum":
		switch o {
		case "platinum":
			return value, nil
		case "gold":
			return value * 10, nil
		case "electrum":
			return value * 20, nil
		case "silver":
			return value * 100, nil
		case "copper":
			return value * 1000, nil
		}
	case "gold":
		switch o {
		case "platinum":
			return 0, fmt.Errorf("cannot exchange for a more valued currency")
		case "gold":
			return value, nil
		case "electrum":
			return value * 2, nil
		case "silver":
			return value * 10, nil
		case "copper":
			return value * 100, nil
		}
	case "electrum":
		switch o {
		case "platinum":
			return 0, fmt.Errorf("cannot exchange for a more valued currency")
		case "gold":
			return 0, fmt.Errorf("cannot exchange for a more valued currency")
		case "electrum":
			return value, nil
		case "silver":
			return value * 5, nil
		case "copper":
			return value * 50, nil
		}
	case "silver":
		switch o {
		case "platinum":
			return 0, fmt.Errorf("cannot exchange for a more valued currency")
		case "gold":
			return 0, fmt.Errorf("cannot exchange for a more valued currency")
		case "electrum":
			return 0, fmt.Errorf("cannot exchange for a more valued currency")
		case "silver":
			return value, nil
		case "copper":
			return value * 10, nil
		}
	}
	return 0, fmt.Errorf("unknown exchange transaction")
}
