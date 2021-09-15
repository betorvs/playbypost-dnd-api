package rule

import "github.com/betorvs/playbypost-dnd/domain/rule"

//AbilityList func
func AbilityList() []string {
	return []string{"strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"}
}

//RaceList func
func RaceList() []string {
	return []string{"human", "dwarf", "elf", "halfling", "dragonborn", "gnome", "half-elf", "half-orc", "tiefling"}
}

//RaceListWithSubrace func
func RaceListWithSubrace() []string {
	return []string{"dwarf", "elf", "halfling", "gnome"}
}

//SubraceList func
func SubraceList(value string) []string {
	switch value {
	case "dwarf":
		return []string{"hill-dwarf", "mountain-dwarf"}
	case "elf":
		return []string{"high-elf", "wood-elf"}
	case "halfling":
		return []string{"lightfoot", "stout"}
	case "gnome":
		return []string{"rock-gnomes"}
	default:
		return []string{"hill-dwarf", "mountain-dwarf", "high-elf", "wood-elf", "lightfoot", "stout", "rock-gnomes"}
	}
}

//BackgroundList func
func BackgroundList() []string {
	return []string{"acolyte", "criminal", "folk hero", "noble", "sage", "soldier"}
}

//ClassList return a []string with all classes
func ClassList() []string {
	return []string{"barbarian", "bard", "cleric", "druid", "fighter", "monk", "paladin", "ranger", "rogue", "sorcerer", "warlock", "wizard"}
}

//AlignmentList func
func AlignmentList() []string {
	return []string{"lawful-good", "neutral-good", "chaotic-good", "lawful-neutral", "neutral", "chaotic-neutral", "lawful-evil", "neutral-evil", "chaotic-evil"}
}

//SkillList func
func SkillList() []string {
	return []string{"acrobatics", "animal-handling", "arcana", "athletics", "deception", "history", "insight", "intimidation", "investigation", "medicine", "nature", "perception", "performance", "persuasion", "religion", "sleight-of-hand", "stealth", "survival"}
}

//SkillListAbility func
func SkillListAbility() []string {
	return []string{"acrobatics (dex)", "animal-handling (Wis)", "arcana (int)", "athletics (str)", "deception (cha)", "history (int)", "insight (Wis)", "intimidation (cha)", "investigation (int)", "medicine (Wis)", "nature (int)", "perception (Wis)", "performance (cha)", "persuasion (cha)", "religion (int)", "sleight-of-hand (dex)", "stealth (dex)", "survival (Wis)"}
}

//ProficiencyFullList func
func ProficiencyFullList() []string {
	return []string{"light-armor", "medium-armor", "heavy-armor", "shields", "simple-weapon", "martial-weapon"}
}

//LanguageList func
func LanguageList() []string {
	return []string{"common", "dwarvish", "elvish", "giant", "gnomish", "draconic", "goblin", "halfling", "orc", "abyssal", "celestial", "deep-speech", "infernal", "primordial", "sylvan", "undercommon", "gnoll", "auran", "terran", "aquan", "ignan", "sahuagin", "worg", "otyugh", "winter-wolf", "yeti", "druidic", "sphinx"}
}

func barbarianSkillList() []string {
	return []string{"animal-handling", "athletics", "intimidation", "nature", "perception", "survival"}
}
func bardSkillList() []string {
	return SkillList()
}
func clericSkillList() []string {
	return []string{"history", "insight", "medicine", "persuasion", "religion"}
}
func druidSkillList() []string {
	return []string{"animal-handling", "arcana", "insight", "medicine", "nature", "perception", "survival", "religion"}
}
func fighterSkillList() []string {
	return []string{"acrobatics", "animal-handling", "athletics", "history", "insight", "intimidation", "perception", "survival"}
}
func monkSkillList() []string {
	return []string{"acrobatics", "athletics", "history", "insight", "religion", "stealth"}
}
func paladinSkillList() []string {
	return []string{"athletics", "insight", "intimidation", "medicine", "persuasion", "religion"}
}
func rangerSkillList() []string {
	return []string{"animal-handling", "athletics", "insight", "investigation", "nature", "perception", "stealth", "survival"}
}
func rogueSkillList() []string {
	return []string{"acrobatics", "athletics", "deception", "insight", "intimidation", "investigation", "perception", "performance", "persuasion", "sleight-of-hand", "stealth"}
}
func sorcererSkillList() []string {
	return []string{"arcana", "deception", "insight", "intimidation", "persuasion", "religion"}
}
func warlockSkillList() []string {
	return []string{"arcana", "deception", "history", "intimidation", "investigation", "nature", "religion"}
}
func wizardSkillList() []string {
	return []string{"arcana", "history", "insight", "investigation", "medicine", "religion"}
}

func healSpellList() []string {
	return []string{"cure-wounds", "mass-cure-wounds", "healing-word", "prayer-of-healing", "mass-healing-word", "heal", "mass-heal"}
}

// ClassWithSpellKnown func
func ClassWithSpellKnown() []string {
	return []string{"bard", "ranger", "sorcerer", "warlock", "arcane-trickster", "eldritch-knight"}
}

// ClassWithPreparedSpell func
func ClassWithPreparedSpell() []string {
	return []string{"cleric", "druid", "paladin", "wizard"}
}

// ClassWithSpell func
func ClassWithSpell() []string {
	return []string{"bard", "ranger", "sorcerer", "cleric", "warlock", "wizard", "druid", "paladin", "arcane-trickster", "eldritch-knight"}
}

// ClassWithCantrips func
func ClassWithCantrips() []string {
	return []string{"bard", "cleric", "druid", "sorcerer", "warlock", "wizard", "arcane-trickster", "eldritch-knight"}
}

func skillListByClass(class string) []string {
	switch class {
	case "barbarian":
		return barbarianSkillList()
	case "bard":
		return bardSkillList()
	case "cleric":
		return clericSkillList()
	case "druid":
		return druidSkillList()
	case "fighter":
		return fighterSkillList()
	case "monk":
		return monkSkillList()
	case "paladin":
		return paladinSkillList()
	case "ranger":
		return rangerSkillList()
	case "rogue":
		return rogueSkillList()
	case "sorcerer":
		return sorcererSkillList()
	case "warlock":
		return warlockSkillList()
	case "wizard":
		return wizardSkillList()

	default:
		return []string{}
	}
}

//AbilityForSpell func
func AbilityForSpell(class string) string {
	switch class {
	case "bard":
		return "charisma"
	case "cleric":
		return "wisdom"
	case "druid":
		return "wisdom"
	case "paladin":
		return "charisma"
	case "ranger":
		return "wisdom"
	case "sorcerer":
		return "charisma"
	case "warlock":
		return "charisma"
	case "wizard":
		return "intelligence"
	default:
		return "class not found"
	}
}

//AbilitySkill func
func AbilitySkill(skill string) string {
	switch skill {
	case "acrobatics":
		return "dexterity"
	case "animal-handling":
		return "wisdom"
	case "arcana":
		return "intelligence"
	case "athletics":
		return "strength"
	case "deception":
		return "charisma"
	case "history":
		return "intelligence"
	case "insight":
		return "wisdom"
	case "intimidation":
		return "charisma"
	case "investigation":
		return "intelligence"
	case "medicine":
		return "wisdom"
	case "nature":
		return "intelligence"
	case "perception":
		return "wisdom"
	case "performance":
		return "charisma"
	case "persuasion":
		return "charisma"
	case "religion":
		return "intelligence"
	case "sleight-of-hand":
		return "dexterity"
	case "stealth":
		return "dexterity"
	case "survival":
		return "wisdom"
	default:
		return "ask to master"
	}
}

//RaceTraits func
func RaceTraits(race, subrace string) map[string]string {
	var ability string
	var size string
	var speed string
	var languages string
	var subraces string
	var special string
	var link string
	var subraceAbility string
	var subraceSpecial string
	switch race {
	case "Dwarf", "dwarf":
		ability = "Your Constitution score increases by 2."
		size = "Dwarves stand between 4 and 5 feet tall and average about 150 pounds. Your size is Medium."
		speed = "25 feet"
		languages = "You can speak, read, and write Common and Dwarvish"
		subraces = "hill-dwarf and mountain-dwarf"
		special = "Darkvision, Dwarven Resilience, Dwarven Combat Training, Tool Proficiency, Stonecunning"
		link = "https://www.dndbeyond.com/sources/basic-rules/races#Dwarf"
	case "Elf", "elf":
		ability = "Your Dexterity score increases by 2."
		size = "Elves range from under 5 to over 6 feet tall and have slender builds. Your size is Medium."
		speed = "30 feet"
		languages = "You can speak, read, and write Common and Elvish."
		subraces = "high-elf, wood-elf, and dark elves, who are commonly called drow"
		special = "Darkvision, Keen Senses, Fey Ancestry, Trance"
		link = "https://www.dndbeyond.com/sources/basic-rules/races#Elf"
	case "Halfling", "halfling":
		ability = "Your Dexterity score increases by 2."
		size = "Halflings average about 3 feet tall and weigh about 40 pounds. Your size is Small."
		speed = "25 feet"
		languages = "You can speak, read, and write Common and Halfling"
		subraces = "lightfoot and stout"
		special = "Lucky, Brave, Halfling Nimbleness"
		link = "https://www.dndbeyond.com/sources/basic-rules/races#Halfling"
	case "Human", "human":
		ability = "Your ability scores each increase by 1."
		size = "Humans vary widely in height and build, from barely 5 feet to well over 6 feet tall. Regardless of your position in that range, your size is Medium."
		speed = "30 feet"
		languages = "You can speak, read, and write Common and one extra language of your choice"
		subraces = ""
		special = ""
		link = "https://www.dndbeyond.com/sources/basic-rules/races#Human"
	case "Dragonborn", "dragonborn":
		ability = "Your Strength score increases by 2, and your Charisma score increases by 1."
		size = "Dragonborn are taller and heavier than humans, standing well over 6 feet tall and averaging almost 250 pounds. Your size is Medium."
		speed = "30 feet"
		languages = "You can speak, read, and write Common and Draconic"
		subraces = ""
		special = "Draconic Ancestry, Breath Weapon, Damage Resistance"
		link = "https://www.dndbeyond.com/sources/basic-rules/races#Dragonborn"
	case "Gnome", "gnome":
		ability = "Your Intelligence score increases by 2."
		size = "Gnomes are between 3 and 4 feet tall and average about 40 pounds. Your size is Small."
		speed = "25 feet"
		languages = "You can speak, read, and write Common and Gnomish"
		subraces = "rock-gnomes"
		special = "Darkvision, Gnome Cunning"
		link = "https://www.dndbeyond.com/sources/basic-rules/races#Gnome"
	case "Half-Elf", "half elf", "half-elf":
		ability = "Your Charisma score increases by 2, and two other ability scores of your choice increase by 1."
		size = "Half-elves are about the same size as humans, ranging from 5 to 6 feet tall. Your size is Medium."
		speed = "30 feet"
		languages = "You can speak, read, and write Common, Elvish, and one extra language of your choice."
		subraces = ""
		special = "Darkvision, Fey Ancestry, Skill Versatility"
		link = "https://www.dndbeyond.com/sources/basic-rules/races#HalfElf"
	case "Half-Orc", "half orc", "half-orc":
		ability = "Your Strength score increases by 2, and your Constitution score increases by 1"
		size = "Half-orcs are somewhat larger and bulkier than humans, and they range from 5 to well over 6 feet tall. Your size is Medium."
		speed = "30 fee"
		languages = "You can speak, read, and write Common and Orc"
		subraces = ""
		special = "Darkvision, Menacing, Relentless Endurance, Savage Attacks"
		link = "https://www.dndbeyond.com/sources/basic-rules/races#HalfOrc"
	case "Tiefling", "tiefling":
		ability = "Your Intelligence score increases by 1, and your Charisma score increases by 2."
		size = "Tieflings are about the same size and build as humans. Your size is Medium."
		speed = "30 feet"
		languages = " You can speak, read, and write Common and Infernal."
		subraces = ""
		special = "Darkvision, Hellish Resistance, Infernal Legacy"
		link = "https://www.dndbeyond.com/sources/basic-rules/races#Tiefling"
	}
	subraceAbility, subraceSpecial = SubraceTraits(subrace)
	raceMap := make(map[string]string)
	raceMap["ability"] = ability
	raceMap["size"] = size
	raceMap["speed"] = speed
	raceMap["languages"] = languages
	raceMap["special"] = special
	raceMap["link"] = link
	raceMap["subraces"] = subraces
	raceMap["subraceAbility"] = subraceAbility
	raceMap["subraceSpecial"] = subraceSpecial
	return raceMap
}

//SubraceTraits func
func SubraceTraits(subrace string) (string, string) {
	var subraceAbility string
	var subraceSpecial string
	switch subrace {
	case "Hill Dwarf", "hill-dwarf", "hill", "dwarf-hill":
		subraceAbility = "Your Wisdom score increases by 1"
		subraceSpecial = "Dwarven Toughness"
	case "Mountain Dwarf", "mountain-dwarf", "mountain", "dwarf-mountain":
		subraceAbility = "Your Strength score increases by 2."
		subraceSpecial = "Dwarven Armor Training"
	case "High Elf", "high-elf", "high", "highelf":
		subraceAbility = "Your Intelligence score increases by 1."
		subraceSpecial = "Elf Weapon Training, Cantrip, Extra Language"
	case "Wood Elf", "wood-elf", "woodelf", "wood":
		subraceAbility = "Your Wisdom score increases by 1."
		subraceSpecial = "Elf Weapon Training, Fleet of Foot, Mask of the Wild"
	case "Lightfoot", "lightfoot", "foot", "light":
		subraceAbility = "Your Charisma score increases by 1"
		subraceSpecial = "Naturally Stealthy"
	case "Stout", "stout":
		subraceAbility = "Your Constitution score increases by 1."
		subraceSpecial = "Stout Resilience"
	case "Rock Gnomes", "rock-gnomes", "rock-gnome", "rock", "rockgnome":
		subraceAbility = "Your Constitution score increases by 1."
		subraceSpecial = "Artificer’s Lore, Tinker"
	default:
		subraceAbility = ""
		subraceSpecial = ""
	}
	return subraceAbility, subraceSpecial
}

//ClassInfo func return an map[string]string with description, hit dice, primary Ability, saving Throws and armor and Weapon Proficiency
func ClassInfo(class string) map[string]string {
	var description string
	var hitDice string
	var primaryAbility string
	var savingThrows string
	var armorWeaponProficiency string
	var skills string
	var link string
	switch class {
	case "barbarian":
		description = "A fierce warrior of primitive background who can enter a battle rage"
		hitDice = "d12"
		primaryAbility = "Strength"
		savingThrows = "Strength & Constitution"
		armorWeaponProficiency = "Light and medium armor, shields, simple and martial weapons"
		skills = "Choose two from animal-handling, Athletics, Intimidation, Nature, Perception, and Survival"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Barbarian"
	case "bard":
		description = "An inspiring magician whose power echoes the music of creation"
		hitDice = "d8"
		primaryAbility = "Charisma"
		savingThrows = "Dexterity & Charisma"
		armorWeaponProficiency = "Light armor, simple weapons, hand crossbows, longswords, rapiers, shortswords"
		skills = "Choose any three"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Bard"
	case "cleric":
		description = "A priestly champion who wields divine magic in service of a higher power"
		hitDice = "d8"
		primaryAbility = "Wisdom"
		savingThrows = "Wisdom & Charisma"
		armorWeaponProficiency = "Light and medium armor, shields, simple weapons"
		skills = "Choose two from History, Insight, Medicine, Persuasion, and Religion"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Cleric"
	case "druid":
		description = "A priest of the Old Faith, wielding the powers of nature — moonlight and plant growth, fire and lightning — and adopting animal forms"
		hitDice = "d8"
		primaryAbility = "Wisdom"
		savingThrows = "Intelligence & Wisdom"
		armorWeaponProficiency = "Light and medium armor (nonmetal), shields (nonmetal), clubs, daggers, darts, javelins, maces, quarterstaffs, scimitars, sickles, slings, spears"
		skills = "Choose two from Arcana, animal-handling, Insight, Medicine, Nature, Perception, Religion, and Survival"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Druid"
	case "fighter":
		description = "A master of martial combat, skilled with a variety of weapons and armor"
		hitDice = "d10"
		primaryAbility = "Strength or Dexterity"
		savingThrows = "Strength & Constitution"
		armorWeaponProficiency = "All armor, shields, simple and martial weapons"
		skills = "Choose two skills from Acrobatics, animal-handling, Athletics, History, Insight, Intimidation, Perception, and Survival"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Fighter"
	case "monk":
		description = "A master of martial arts, harnessing the power of the body in pursuit of physical and spiritual perfection"
		hitDice = "d8"
		primaryAbility = "Dexterity & Wisdom"
		savingThrows = "Strength & Dexterity"
		armorWeaponProficiency = "Simple weapons, shortswords"
		skills = "Choose two from Acrobatics, Athletics, History, Insight, Religion, and Stealth"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Monk"
	case "paladin":
		description = "A holy warrior bound to a sacred oath"
		hitDice = "d10"
		primaryAbility = "Strength & Charisma"
		savingThrows = "Wisdom & Charisma"
		armorWeaponProficiency = "All armor, shields, simple and martial weapons"
		skills = "Choose two from Athletics, Insight, Intimidation, Medicine, Persuasion, and Religion"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Paladin"
	case "ranger":
		description = "A warrior who uses martial prowess and nature magic to combat threats on the edges of civilization"
		hitDice = "d10"
		primaryAbility = "Dexterity & Wisdom"
		savingThrows = "Strength & Dexterity"
		armorWeaponProficiency = "Light and medium armor, shields, simple and martial weapons"
		skills = "Choose three from animal-handling, Athletics, Insight, Investigation, Nature, Perception, Stealth, and Survival"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Ranger"
	case "rogue":
		description = "A scoundrel who uses stealth and trickery to overcome obstacles and enemies"
		hitDice = "d8"
		primaryAbility = "Dexterity"
		savingThrows = "Dexterity & Intelligence"
		armorWeaponProficiency = "Light armor, simple weapons, hand crossbows, longswords, rapiers, shortswords"
		skills = "Skills: Choose four from Acrobatics, Athletics, Deception, Insight, Intimidation, Investigation, Perception, Performance, Persuasion, sleight-of-hand, and Stealth"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Rogue"
	case "sorcerer":
		description = "A spellcaster who draws on inherent magic from a gift or bloodline"
		hitDice = "d6"
		primaryAbility = "Charisma"
		savingThrows = "Constitution & Charisma"
		armorWeaponProficiency = "Daggers, darts, slings, quarterstaffs, light crossbows"
		skills = "Skills: Choose two from Arcana, Deception, Insight, Intimidation, Persuasion, and Religion"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Sorcerer"
	case "warlock":
		description = "A wielder of magic that is derived from a bargain with an extraplanar entity"
		hitDice = "d8"
		primaryAbility = "Charisma"
		savingThrows = "Wisdom & Charisma"
		armorWeaponProficiency = "Light armor, simple weapons"
		skills = "Choose two skills from Arcana, Deception, History, Intimidation, Investigation, Nature, and Religion"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Warlock"
	case "wizard":
		description = "A scholarly magic-user capable of manipulating the structures of reality"
		hitDice = "d6"
		primaryAbility = "Intelligence"
		savingThrows = "Intelligence & Wisdom"
		armorWeaponProficiency = "Daggers, darts, slings, quarterstaffs, light crossbows"
		skills = "Choose two from Arcana, History, Insight, Investigation, Medicine, and Religion"
		link = "https://www.dndbeyond.com/sources/basic-rules/classes#Wizard"
	}
	classInfo := make(map[string]string)
	classInfo["description"] = description
	classInfo["hitDice"] = hitDice
	classInfo["primaryAbility"] = primaryAbility
	classInfo["savingThrows"] = savingThrows
	classInfo["armorWeaponProficiency"] = armorWeaponProficiency
	classInfo["skiils"] = skills
	classInfo["link"] = link
	return classInfo
}

//ListInformation func
func ListInformation(name, value string) (*rule.SimpleList, error) {
	list := new(rule.SimpleList)
	switch name {
	case "race", "races":
		list.List = RaceList()
		if value == "subrace" {
			list.List = RaceListWithSubrace()
		}

	case "class", "classes":
		list.List = ClassList()

	case "background", "backgrounds":
		list.List = BackgroundList()

	case "ability":
		list.List = AbilityList()

	case "alignment":
		list.List = AlignmentList()

	case "skills", "skill":
		list.List = SkillListAbility()

	case "subrace", "subraces":
		list.List = SubraceList(value)

	case "condition", "conditions":
		list.List = ListConditions()

	case "damage", "damagetype":
		list.List = ListOfDamageTypes()
	}
	return list, nil
}

//FullDescription func
func FullDescription(kind, name, subname string) (*rule.FullDescription, error) {
	fullDescription := new(rule.FullDescription)
	switch kind {
	case "race":
		fullDescription.Description = RaceTraits(name, subname)
	case "class":
		fullDescription.Description = ClassInfo(name)

	case "background":
		m, _ := BackgroundStatistics(name)
		fullDescription.Description = m

	case "condition", "conditions":
		d := conditionsMap(name)
		fullDescription.Description = d

	case "damage", "damagetype":
		d := damageTypeMap(name)
		fullDescription.Description = d
	}

	return fullDescription, nil
}
