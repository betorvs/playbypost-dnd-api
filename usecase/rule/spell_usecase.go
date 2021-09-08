package rule

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/database"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//GetSpellListDescription func
func GetSpellListDescription(queryParameters url.Values) []rule.SpellDescription {
	db := database.GetDatabaseRepository()
	spellList := db.GetSpellDescriptionDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.SpellDescription
		for _, v := range spellList {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						// fmt.Println(param)
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				case "level":
					for _, n := range param {
						number := utils.ExtractWholeInt(n)
						if v.Level == number {
							filtered = append(filtered, v)
						}
					}
				case "title":
					for _, n := range param {
						// fmt.Println(param)
						if strings.Contains(v.Title, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		// fmt.Println(filtered)
		// fmt.Println(len(filtered))
		return filtered
	}
	return spellList
}

func getSpellByName(name string) (spell rule.SpellDescription) {
	db := database.GetDatabaseRepository()
	spellList := db.GetSpellDescriptionDatabase()
	for _, v := range spellList {
		if name == v.Name {
			spell = v
		}
	}
	return spell
}

func spellDamageIncrease(spell rule.SpellDescription, newLevel int) string {
	damageNew := spell.DamageDice

	if spell.AtHigherLevels != "" {
		// fmt.Println(spell.AtHigherLevels)
		// "spiritual-weapon" every 2 slot above 1d8 and "flame-blade" 1d6 every 2 slots
		// "arcane-hand" and "wall-of-ice" more than one effect and damage
		var reIncreaseDamage = regexp.MustCompile(`(?m)( increase.* by (\d+)?d(\d+)([\+\-]\d+)? for each slot level above *.+?\.)`)
		newDamage := reIncreaseDamage.FindString(spell.AtHigherLevels)
		if newDamage != "" && !strings.Contains(newDamage, "damage from") {
			minimumLevel := spellLevelRegex(newDamage)
			// fmt.Println(newLevel, minimumLevel)
			if newLevel > minimumLevel {
				var dice = regexp.MustCompile(`(?m)d(\d+)`)
				increaseDice := dice.FindString(newDamage)
				var reFirstNumber = regexp.MustCompile(`(?m)[0-9](\d+)?d`)
				initialNumber := reFirstNumber.FindString(spell.DamageDice)
				// fmt.Println(spell.DamageDice)
				// fmt.Println(increaseDice)
				// fmt.Println(initialNumber)
				if increaseDice != "" && initialNumber != "" {
					initial := utils.ExtractWholeInt(initialNumber)
					usedLevel := newLevel - minimumLevel + initial
					// fmt.Printf("%v%s\n", usedLevel, increaseDice)
					damageNew = fmt.Sprintf("%v%s\n", usedLevel, increaseDice)
				}

			}
		}
	}
	if spell.DamageIncrease != "" {
		switch spell.Name {
		case "produce-flame":
			if newLevel > 4 {
				damageNew = "2d8"
			}
			if newLevel > 10 {
				damageNew = "3d8"
			}
			if newLevel > 16 {
				damageNew = "4d8"
			}
		case "acid-splash":
			if newLevel > 4 {
				damageNew = "2d6"
			}
			if newLevel > 10 {
				damageNew = "3d6"
			}
			if newLevel > 16 {
				damageNew = "4d6"
			}

		case "chill-touch":
			if newLevel > 4 {
				damageNew = "2d8"
			}
			if newLevel > 10 {
				damageNew = "3d8"
			}
			if newLevel > 16 {
				damageNew = "4d8"
			}
		}
	}

	return damageNew
}

func spellDamageMax(damageNew string) int {
	var damageValue int
	var dice = regexp.MustCompile(`(?m)d(\d+)`)
	increaseDice := dice.FindString(damageNew)
	var reFirstNumber = regexp.MustCompile(`(?m)[0-9](\d+)?d`)
	initialNumber := reFirstNumber.FindString(damageNew)
	if increaseDice != "" && initialNumber != "" {
		initial := utils.ExtractWholeInt(initialNumber)
		diceValue := utils.ExtractWholeInt(increaseDice)
		damageValue = initial * diceValue
	}
	return damageValue
}

func spellLevelRegex(value string) int {
	var level int
	var reSpellLevel = regexp.MustCompile(`(?m)(1st|2nd|3rd|4th|5th|6th|7th|8th|9th)`)
	levelString := reSpellLevel.FindString(value)
	fmt.Println(levelString)
	if levelString != "" {
		level = utils.ExtractWholeInt(levelString)
	}
	return level
}

func spellHealIncreases(healDice, healingIncreases string, newLevel int) string {
	healNew := healDice
	//  each slot level above *.+?\.
	// var reDices = regexp.MustCompile(`(?m)(healing increases by (\d+)?d(\d+)([\+\-]\d+)? for )`)
	// var reLevel = regexp.MustCompile(`(?m)(each slot level above *.+?\.)`)
	minimumLevel := spellLevelRegex(healingIncreases)
	var reDices = regexp.MustCompile(`(?m)(\d+)?d(\d+)([\+\-]\d+)?`)
	newHeal := reDices.FindString(healingIncreases)
	if newHeal != "" && newLevel > minimumLevel {
		// fmt.Println(newLevel, minimumLevel)
		var dice = regexp.MustCompile(`(?m)d(\d+)`)
		increaseDice := dice.FindString(newHeal)
		var reFirstNumber = regexp.MustCompile(`(?m)[0-9](\d+)?d`)
		initialNumber := reFirstNumber.FindString(healDice)
		// fmt.Println(spell.DamageDice)
		// fmt.Println(increaseDice)
		// fmt.Println(initialNumber)
		if increaseDice != "" && initialNumber != "" {
			initial := utils.ExtractWholeInt(initialNumber)
			usedLevel := newLevel - minimumLevel + initial
			// fmt.Printf("%v%s\n", usedLevel, increaseDice)
			healNew = fmt.Sprintf("%v%s\n", usedLevel, increaseDice)
		}

	}
	return healNew
}

func healSpell(level int) int {
	heal := 70
	if level > 6 {
		multi := level - 6
		heal = (multi * 10) + 70
	}

	return heal
}

//GetFullSpellList func
func GetFullSpellList(class string) []string {
	var result []string
	if class == "ranger" || class == "paladin" {
		for i := 1; i < 6; i++ {
			partial := GetSpellListByClass(class, i)
			result = append(result, partial.List...)
		}
	}
	if class != "ranger" && class != "paladin" {
		for i := 0; i < 10; i++ {
			partial := GetSpellListByClass(class, i)
			result = append(result, partial.List...)
		}
	}
	return result
}

func getFullListWithFeature(class string, classFeatures []string, level int) []string {
	fullList := GetFullSpellList(class)
	for _, f := range classFeatures {
		if utils.StringInSlice(f, featuresWithExtraSpellList()) {
			for i := 0; i < level; i++ {
				list := extraSpellList(f, i)
				if len(list) != 0 {
					fullList = append(fullList, list...)
				}
			}

		}
	}

	return fullList
}

//GetSpellListByClass func
func GetSpellListByClass(class string, level int) (list rule.SimpleList) {
	//
	db := database.GetDatabaseRepository()
	var spellList []string
	switch class {
	case "bard":
		switch level {
		case 0:

			spellList = db.GetSpellListByClass().Bard.Level0
		case 1:
			spellList = db.GetSpellListByClass().Bard.Level1
		case 2:
			spellList = db.GetSpellListByClass().Bard.Level2
		case 3:
			spellList = db.GetSpellListByClass().Bard.Level3
		case 4:
			spellList = db.GetSpellListByClass().Bard.Level4
		case 5:
			spellList = db.GetSpellListByClass().Bard.Level5
		case 6:
			spellList = db.GetSpellListByClass().Bard.Level6
		case 7:
			spellList = db.GetSpellListByClass().Bard.Level7
		case 8:
			spellList = db.GetSpellListByClass().Bard.Level8
		case 9:
			spellList = db.GetSpellListByClass().Bard.Level9
		}

	case "ranger":
		switch level {
		case 1:
			spellList = db.GetSpellListByClass().Ranger.Level1
		case 2:
			spellList = db.GetSpellListByClass().Ranger.Level2
		case 3:
			spellList = db.GetSpellListByClass().Ranger.Level3
		case 4:
			spellList = db.GetSpellListByClass().Ranger.Level4
		case 5:
			spellList = db.GetSpellListByClass().Ranger.Level5
		default:
			spellList = db.GetSpellListByClass().Ranger.Level1
		}

	case "sorcerer":
		switch level {
		case 0:
			spellList = db.GetSpellListByClass().Sorcerer.Level0
		case 1:
			spellList = db.GetSpellListByClass().Sorcerer.Level1
		case 2:
			spellList = db.GetSpellListByClass().Sorcerer.Level2
		case 3:
			spellList = db.GetSpellListByClass().Sorcerer.Level3
		case 4:
			spellList = db.GetSpellListByClass().Sorcerer.Level4
		case 5:
			spellList = db.GetSpellListByClass().Sorcerer.Level5
		case 6:
			spellList = db.GetSpellListByClass().Sorcerer.Level6
		case 7:
			spellList = db.GetSpellListByClass().Sorcerer.Level7
		case 8:
			spellList = db.GetSpellListByClass().Sorcerer.Level8
		case 9:
			spellList = db.GetSpellListByClass().Sorcerer.Level9
		}

	case "cleric":
		switch level {
		case 0:
			spellList = db.GetSpellListByClass().Cleric.Level0
		case 1:
			spellList = db.GetSpellListByClass().Cleric.Level1
		case 2:
			spellList = db.GetSpellListByClass().Cleric.Level2
		case 3:
			spellList = db.GetSpellListByClass().Cleric.Level3
		case 4:
			spellList = db.GetSpellListByClass().Cleric.Level4
		case 5:
			spellList = db.GetSpellListByClass().Cleric.Level5
		case 6:
			spellList = db.GetSpellListByClass().Cleric.Level6
		case 7:
			spellList = db.GetSpellListByClass().Cleric.Level7
		case 8:
			spellList = db.GetSpellListByClass().Cleric.Level8
		case 9:
			spellList = db.GetSpellListByClass().Cleric.Level9
		}

	case "wizard":
		switch level {
		case 0:
			spellList = db.GetSpellListByClass().Wizard.Level0
		case 1:
			spellList = db.GetSpellListByClass().Wizard.Level1
		case 2:
			spellList = db.GetSpellListByClass().Wizard.Level2
		case 3:
			spellList = db.GetSpellListByClass().Wizard.Level3
		case 4:
			spellList = db.GetSpellListByClass().Wizard.Level4
		case 5:
			spellList = db.GetSpellListByClass().Wizard.Level5
		case 6:
			spellList = db.GetSpellListByClass().Wizard.Level6
		case 7:
			spellList = db.GetSpellListByClass().Wizard.Level7
		case 8:
			spellList = db.GetSpellListByClass().Wizard.Level8
		case 9:
			spellList = db.GetSpellListByClass().Wizard.Level9
		}

	case "arcane-trickster", "eldritch-knight":
		switch level {
		case 0:
			spellList = db.GetSpellListByClass().Wizard.Level0
		case 1:
			spellList = db.GetSpellListByClass().Wizard.Level1
		case 2:
			spellList = db.GetSpellListByClass().Wizard.Level2
		case 3:
			spellList = db.GetSpellListByClass().Wizard.Level3
		case 4:
			spellList = db.GetSpellListByClass().Wizard.Level4
			// case 5:
			// 	spellList = db.GetSpellListByClass().Wizard.Level5
			// case 6:
			// 	spellList = db.GetSpellListByClass().Wizard.Level6
			// case 7:
			// 	spellList = db.GetSpellListByClass().Wizard.Level7
			// case 8:
			// 	spellList = db.GetSpellListByClass().Wizard.Level8
			// case 9:
			// 	spellList = db.GetSpellListByClass().Wizard.Level9
		}

	case "druid":
		switch level {
		case 0:
			spellList = db.GetSpellListByClass().Druid.Level0
		case 1:
			spellList = db.GetSpellListByClass().Druid.Level1
		case 2:
			spellList = db.GetSpellListByClass().Druid.Level2
		case 3:
			spellList = db.GetSpellListByClass().Druid.Level3
		case 4:
			spellList = db.GetSpellListByClass().Druid.Level4
		case 5:
			spellList = db.GetSpellListByClass().Druid.Level5
		case 6:
			spellList = db.GetSpellListByClass().Druid.Level6
		case 7:
			spellList = db.GetSpellListByClass().Druid.Level7
		case 8:
			spellList = db.GetSpellListByClass().Druid.Level8
		case 9:
			spellList = db.GetSpellListByClass().Druid.Level9
		}

	case "paladin":
		switch level {
		case 1:
			spellList = db.GetSpellListByClass().Paladin.Level1
		case 2:
			spellList = db.GetSpellListByClass().Paladin.Level2
		case 3:
			spellList = db.GetSpellListByClass().Paladin.Level3
		case 4:
			spellList = db.GetSpellListByClass().Paladin.Level4
		case 5:
			spellList = db.GetSpellListByClass().Paladin.Level5
		default:
			spellList = db.GetSpellListByClass().Paladin.Level1
		}

	case "warlock":
		switch level {
		case 0:
			spellList = db.GetSpellListByClass().Warlock.Level0
		case 1:
			spellList = db.GetSpellListByClass().Warlock.Level1
		case 2:
			spellList = db.GetSpellListByClass().Warlock.Level2
		case 3:
			spellList = db.GetSpellListByClass().Warlock.Level3
		case 4:
			spellList = db.GetSpellListByClass().Warlock.Level4
		case 5:
			spellList = db.GetSpellListByClass().Warlock.Level5
		case 6:
			spellList = db.GetSpellListByClass().Warlock.Level6
		case 7:
			spellList = db.GetSpellListByClass().Warlock.Level7
		case 8:
			spellList = db.GetSpellListByClass().Warlock.Level8
		case 9:
			spellList = db.GetSpellListByClass().Warlock.Level9
		}

	}
	for _, v := range spellList {
		list.List = append(list.List, strings.ToLower(strings.ReplaceAll(v, " ", "-")))
	}

	return list
}
