package rule

import (
	"net/url"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/database"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//MonsterForNPC func
func MonsterForNPC(queryParameters url.Values) []rule.MonsterNPC {
	db := database.GetDatabaseRepository()
	monsters := db.GetMonsterDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.MonsterNPC
		for _, v := range monsters {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				case "xp":
					for _, n := range param {
						number := utils.ExtractWholeInt(n)
						if v.XP <= number {
							filtered = append(filtered, v)
						}
					}
				case "type":
					for _, n := range param {
						if strings.Contains(v.Type, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		// fmt.Println(len(filtered))
		return filtered
	}
	return monsters
}

// MosterByName returns a monster by name
func MosterByName(name string) rule.MonsterNPC {
	db := database.GetDatabaseRepository()
	monsters := db.GetMonsterDatabase()
	var monster rule.MonsterNPC
	for _, v := range monsters {
		if v.Name == name {
			monster = v
		}
	}
	return monster
}
func returnMosterByType(name string) []rule.MonsterNPC {
	db := database.GetDatabaseRepository()
	monsters := db.GetMonsterDatabase()
	var monster []rule.MonsterNPC
	for _, v := range monsters {
		if v.Type == name {
			monster = append(monster, v)
		}
	}
	return monster
}

func returnMonsterPaladinEnemy(name string) bool {
	db := database.GetDatabaseRepository()
	monsters := db.GetMonsterDatabase()
	for _, v := range monsters {
		if v.Name == name && v.Type == "fiend" {
			return true
		}
		if v.Name == name && v.Type == "undead" {
			return true
		}
	}
	return false
}

func monsterTypeSingular(name string) string {
	switch name {
	case "aberrations":
		return "aberration"
	case "beasts":
		return "beast"
	case "celestials":
		return "celestial"
	case "constructs":
		return "construct"
	case "dragons":
		return "dragon"
	case "elementals":
		return "elemental"
	case "fey":
		return "fey"
	case "fiends":
		return "fiend"
	case "giants":
		return "giant"
	case "monstrosities":
		return "monstrosity"
	case "oozes":
		return "ooze"
	case "plants":
		return "plant"
	case "undead":
		return "undead"
	}
	return ""
}
