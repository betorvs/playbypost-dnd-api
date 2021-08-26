package rule

import (
	"net/url"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//ReadOneMonsterFromFile func
// func ReadOneMonsterFromFile(name string) rulesdnd.Monster {

// 	monsters := database.GetMonsterDatabase()
// 	var result rulesdnd.Monster
// 	for _, v := range monsters {
// 		sanitizedName := strings.ToLower(strings.ReplaceAll(v.Name, " ", "-"))
// 		if sanitizedName == strings.ToLower(name) {
// prettyJSON, err := json.MarshalIndent(v, "", "  ")
// if err != nil {
// 	log.Fatal("Failed to generate json", err)
// }
// fmt.Printf("%s\n", string(prettyJSON))
// 			result = v
// 		}
// 	}
// 	return result
// }

//ReadAllMonsterFromFile func
// func ReadAllMonsterFromFile() []rulesdnd.Monster {
// 	monsters := database.GetMonsterDatabase()
// 	return monsters
// }

//MonsterByChallenge func
// func MonsterByChallenge(queryParameters url.Values) []rulesdnd.MonsterByChallenge {
// 	monsters := database.GetMonsterDatabase()
// 	var result []rulesdnd.MonsterByChallenge
// 	for _, v := range monsters {
// 		var partialResult rulesdnd.MonsterByChallenge
// 		partialResult.Name = strings.ToLower(strings.ReplaceAll(v.Name, " ", "-"))
// 		data := strings.Split(v.Challenge, " ")
// 		challenge := data[0]
// 		reg, err := regexp.Compile("[^0-9]+")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		xp := reg.ReplaceAllString(data[1], "")
// 		partialResult.Challenge = challenge
// 		partialResult.XP = xp
// 		result = append(result, partialResult)
// 	}
// if len(queryParameters) != 0 {
// 	var filtered []rulesdnd.MonsterByChallenge
// 	for _, v := range result {
// 		for paramName, param := range queryParameters {
// 			switch paramName {
// 			case "challenge":
// 				for _, c := range param {
// 					if v.Challenge == c {
// 						filtered = append(filtered, v)
// 					}
// 				}
// 			case "name":
// 				for _, n := range param {
// 					if strings.Contains(v.Name, n) {
// 				filtered = append(filtered, v)
// 			}
// 		}

// 	case "xp":
// 		for _, x := range param {
// 			if v.XP == x {
// 				filtered = append(filtered, v)
// 			}
// 		}
// 	}
// }
// 		}
// 		return filtered
// 	}
// 	return result
// }

//MonsterForNPC func
func MonsterForNPC(queryParameters url.Values) []rule.MonsterNPC {
	db := rule.GetDatabaseRepository()
	monsters := db.GetMonsterDatabase()
	// var result []rulesdnd.ReturnMonsterNPC
	// for _, v := range monsters {
	// 	var partialResult rulesdnd.ReturnMonsterNPC
	// 	partialResult.Name = strings.ToLower(strings.ReplaceAll(v.Name, " ", "-"))
	// 	armor := strings.Split(v.ArmorClass, "(")
	// 	// fmt.Println(partialResult.Name, armor[0])
	// 	reg, err := regexp.Compile("[^0-9]+")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	armorClass, err := strconv.Atoi(reg.ReplaceAllString(armor[0], ""))
	// 	if err != nil {
	// 		fmt.Printf("error convert armor string to int %v", err)
	// 	}
	// 	hit := strings.Split(v.HitPoints, "(")
	// 	hitpoints, err := strconv.Atoi(strings.TrimSpace(hit[0]))
	// 	if err != nil {
	// 		fmt.Printf("error convert hitpoints string to int %v", err)
	// 	}
	// 	data := strings.Split(v.Challenge, " ")
	// 	xpString := reg.ReplaceAllString(data[1], "")
	// 	xp, err := strconv.Atoi(strings.TrimSpace(xpString))
	// 	if err != nil {
	// 		fmt.Printf("error convert xp string to int %v", err)
	// 	}
	// 	partialResult.ArmorClass = armorClass
	// 	partialResult.HitPoints = hitpoints
	// 	partialResult.XP = xp
	// 	partialResult.ConditionImmunities = []string{}
	// 	partialResult.DamageImmunities = []string{}
	// 	partialResult.DamageVulnerabilities = []string{}
	// 	condImu := strings.Split(v.ConditionImmunities, ",")
	// 	for _, v := range condImu {
	// 		if v != "" {
	// 			partialResult.ConditionImmunities = append(partialResult.ConditionImmunities, strings.TrimSpace(strings.ToLower(v)))
	// 		}
	// 	}
	// 	damaImu := strings.Split(v.DamageImmunities, ",")
	// 	for _, v := range damaImu {
	// 		if v != "" && !strings.Contains(v, "Slashing from Nonmagical") {
	// 			partialResult.DamageImmunities = append(partialResult.DamageImmunities, strings.TrimSpace(strings.ToLower(v)))
	// 		}
	// 		if strings.Contains(v, "Slashing from Nonmagical") {
	// 			partialResult.DamageImmunities = append(partialResult.DamageImmunities, "slashing")
	// 		}
	// 	}
	// 	damaVul := strings.Split(v.DamageVulnerabilities, ",")
	// 	for _, v := range damaVul {
	// 		if v != "" {
	// 			partialResult.DamageVulnerabilities = append(partialResult.DamageVulnerabilities, strings.TrimSpace(strings.ToLower(v)))
	// 		}
	// 	}
	// 	traits := parseAction(v.Traits)
	// 	traitSlice := strings.Split(traits, ";")
	// 	for _, v := range traitSlice {
	// 		if v != "" {
	// 			partialResult.Traits = append(partialResult.Traits, v)
	// 		}
	// 	}
	// 	legAct := parseAction(v.LegendaryActions)
	// 	legActSlice := strings.Split(legAct, ";")
	// 	for _, v := range legActSlice {
	// 		if v != "" {
	// 			partialResult.LegendaryActions = append(partialResult.LegendaryActions, v)
	// 		}
	// 	}

	// 	partialResult.ImgURL = v.ImgURL

	// 	weapons := make([]rulesdnd.WeaponAttack, 0)
	// 	act := parseAction(v.Actions)
	// 	actions := make([]string, 0)
	// 	actSlice := strings.Split(act, ";")
	// 	for _, v := range actSlice {
	// 		if v != "" {
	// 			actions = append(actions, v)
	// 			if strings.Contains(v, "Weapon Attack") {
	// 				weapon := parseAttack(strings.Split(v, ":"))
	// 				weapons = append(weapons, weapon)
	// 			}
	// 		}
	// 	}
	// 	savings := parseSavings(v)
	// 	if strings.Contains(v.Senses, "Passive Perception") {
	// 		pas := strings.Split(v.Senses, "Perception")
	// 		passive, err := strconv.Atoi(strings.TrimSpace(pas[1]))
	// 		if err != nil {
	// 			fmt.Printf("error convert xp string to int %v", err)
	// 		}
	// 		partialResult.PassivePerception = passive
	// 	}
	// 	partialResult.Actions = actions
	// 	partialResult.WeaponAttack = weapons
	// 	partialResult.Savings = savings

	// 	result = append(result, partialResult)
	// }
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
	db := rule.GetDatabaseRepository()
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
	db := rule.GetDatabaseRepository()
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
	db := rule.GetDatabaseRepository()
	monsters := db.GetMonsterDatabase()
	for _, v := range monsters {
		if v.Name == name && v.Type == "fiends" {
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

// func returnMosterByName(name string) rulesdnd.MonsterChecks {
// 	monsters := database.GetMonsterDatabase()
// 	var monster rulesdnd.Monster
// 	for _, v := range monsters {
// 		tempName := strings.ToLower(strings.ReplaceAll(v.Name, " ", "-"))
// 		if tempName == name {
// 			monster = v
// 		}
// 	}
// 	var result rulesdnd.MonsterChecks
// 	weapons := make([]rulesdnd.WeaponAttack, 0)
// 	act := parseAction(monster.Actions)
// 	actions := make([]string, 0)
// 	actSlice := strings.Split(act, ";")
// 	for _, v := range actSlice {
// 		if v != "" {
// 			actions = append(actions, v)
// 			if strings.Contains(v, "Weapon Attack") {
// 				weapon := parseAttack(strings.Split(v, ":"))
// 				weapons = append(weapons, weapon)
// 			}
// 		}
// 	}
// 	if strings.Contains(monster.Senses, "Passive Perception") {
// 		pas := strings.Split(monster.Senses, "Perception")
// 		passive, err := strconv.Atoi(strings.TrimSpace(pas[1]))
// 		if err != nil {
// 			fmt.Printf("error convert xp string to int %v", err)
// 		}
// 		result.PassivePerception = passive
// 	}
// 	armor := strings.Split(monster.ArmorClass, "(")
// 	// fmt.Println(partialResult.Name, armor[0])
// 	reg, err := regexp.Compile("[^0-9]+")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	armorClass, err := strconv.Atoi(reg.ReplaceAllString(armor[0], ""))
// 	if err != nil {
// 		fmt.Printf("error convert armor string to int %v", err)
// 	}
// 	savings := parseSavings(monster)
// 	result.ConditionImmunities = []string{}
// 	result.DamageImmunities = []string{}
// 	result.DamageVulnerabilities = []string{}
// 	condImu := strings.Split(monster.ConditionImmunities, ",")
// 	for _, v := range condImu {
// 		if v != "" {
// 			result.ConditionImmunities = append(result.ConditionImmunities, strings.TrimSpace(strings.ToLower(v)))
// 		}
// 	}
// 	damaImu := strings.Split(monster.DamageImmunities, ",")
// 	for _, v := range damaImu {
// 		if v != "" && !strings.Contains(v, "Slashing from Nonmagical") {
// 			result.DamageImmunities = append(result.DamageImmunities, strings.TrimSpace(strings.ToLower(v)))
// 		}
// 		if strings.Contains(v, "Slashing from Nonmagical") {
// 			result.DamageImmunities = append(result.DamageImmunities, "slashing")
// 		}
// 	}
// 	damaVul := strings.Split(monster.DamageVulnerabilities, ",")
// 	for _, v := range damaVul {
// 		if v != "" {
// 			result.DamageVulnerabilities = append(result.DamageVulnerabilities, strings.TrimSpace(strings.ToLower(v)))
// 		}
// 	}
// 	result.Name = monster.Name
// 	result.Actions = actions
// 	result.WeaponAttack = weapons
// 	result.Savings = savings
// 	result.ArmorClass = armorClass
// 	return result
// }

// func parseAction(s string) string {
// 	var result string
// 	doc, err := html.Parse(strings.NewReader(s))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var f func(*html.Node)
// 	f = func(n *html.Node) {
// 		if n.Data == "strong" {
// 			result += fmt.Sprintf(";")
// 		}
// 		if n.Type == 1 {
// 			result += fmt.Sprintf("%s ", n.Data)
// 		}

// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			f(c)
// 		}
// 	}
// 	f(doc)
// 	return result
// }

// func parseAttack(v []string) rulesdnd.WeaponAttack {
// 	var weapon rulesdnd.WeaponAttack
// 	// fmt.Println(v[2])
// 	name := strings.Split(v[0], ".")
// 	att := strings.Split(v[1], ",")
// 	hitString := strings.Split(v[2], "(")
// 	reg, err := regexp.Compile("[^0-9]+")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	atta := reg.ReplaceAllString(att[0], "")
// 	var attack int
// 	if atta != "" {
// 		attack, err = strconv.Atoi(strings.TrimSpace(atta))
// 		if err != nil {
// 			fmt.Printf("error convert string to int %v for %s", err, name[0])
// 		}
// 	}

// 	hitClean := reg.ReplaceAllString(hitString[0], "")
// 	var hit int
// 	if hitClean != "" {
// 		hit, err = strconv.Atoi(strings.TrimSpace(hitClean))
// 		if err != nil {
// 			fmt.Printf("error convert string to int %v for %s", err, name[0])
// 		}
// 	}
// 	if hit != 1 {
// 		var re1 = regexp.MustCompile(`(?m)\(.*?\)`)
// 		dices := re1.FindString(v[2])
// 		// fmt.Println(strings.ReplaceAll(dices, " ", ""))
// 		var re2 = regexp.MustCompile(`(?m)(\d+)?d(\d+)([\+\-]\d+)?`)
// 		damage := re2.FindString(strings.ReplaceAll(dices, " ", ""))
// 		// fmt.Println(damage)
// 		weapon.Damage = damage
// 	}
// 	var damageType string
// 	if strings.Contains(v[2], "piercing") {
// 		damageType = "piercing"
// 	}
// 	if strings.Contains(v[2], "bludgeoning") {
// 		damageType = "bludgeoning"
// 	}
// 	if strings.Contains(v[2], "slashing") {
// 		damageType = "slashing"
// 	}

// 	weapon.Name = strings.ToLower(name[0])
// 	weapon.AverageDamage = hit
// 	weapon.Attack = attack
// 	weapon.DamageType = damageType
// 	return weapon
// }

// func parseSavings(monster rulesdnd.Monster) map[string]int {
// 	ability := make(map[string]int, 0)

// 	reg, err := regexp.Compile("[^0-9]+")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	str := reg.ReplaceAllString(monster.STRMod, "")
// 	modStr, err := strconv.Atoi(strings.TrimSpace(str))
// 	if err != nil {
// 		fmt.Printf("error convert string to int %v for %s", err, monster.Name)
// 	}
// 	dex := reg.ReplaceAllString(monster.DEXMod, "")
// 	modDex, err := strconv.Atoi(strings.TrimSpace(dex))
// 	if err != nil {
// 		fmt.Printf("error convert string to int %v for %s", err, monster.Name)
// 	}
// 	con := reg.ReplaceAllString(monster.CONMod, "")
// 	modCon, err := strconv.Atoi(strings.TrimSpace(con))
// 	if err != nil {
// 		fmt.Printf("error convert string to int %v for %s", err, monster.Name)
// 	}
// 	inte := reg.ReplaceAllString(monster.INTMod, "")
// 	modInt, err := strconv.Atoi(strings.TrimSpace(inte))
// 	if err != nil {
// 		fmt.Printf("error convert string to int %v for %s", err, monster.Name)
// 	}
// 	wis := reg.ReplaceAllString(monster.WISMod, "")
// 	modWis, err := strconv.Atoi(strings.TrimSpace(wis))
// 	if err != nil {
// 		fmt.Printf("error convert string to int %v for %s", err, monster.Name)
// 	}
// 	car := reg.ReplaceAllString(monster.CHAMod, "")
// 	modCar, err := strconv.Atoi(strings.TrimSpace(car))
// 	if err != nil {
// 		fmt.Printf("error convert string to int %v for %s", err, monster.Name)
// 	}
// 	// "strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"
// 	ability["strength"] = modStr
// 	ability["dexterity"] = modDex
// 	ability["constitution"] = modCon
// 	ability["intelligence"] = modInt
// 	ability["wisdom"] = modWis
// 	ability["charisma"] = modCar
// 	if monster.SavingThrows != "" {
// 		savingsSlice := strings.Split(monster.SavingThrows, ",")
// 		for _, v := range savingsSlice {
// 			h := strings.Split(v, "+")
// 			value, err := strconv.Atoi(strings.TrimSpace(h[1]))
// 			if err != nil {
// 				fmt.Printf("error convert string to int %v for %s", err, monster.Name)
// 			}
// 			switch strings.TrimSpace(h[0]) {
// 			case "STR":
// 				if value > modStr {
// 					ability["strength"] = value
// 				}
// 			case "DEX":
// 				if value > modDex {
// 					ability["dexterity"] = value
// 				}
// 			case "CON":
// 				if value > modCon {
// 					ability["constitution"] = value
// 				}
// 			case "INT":
// 				if value > modInt {
// 					ability["intelligence"] = value
// 				}
// 			case "WIS":
// 				if value > modWis {
// 					ability["wisdom"] = value
// 				}
// 			case "CHA":
// 				if value > modCar {
// 					ability["charisma"] = value
// 				}

// 			}
// 		}
// 	}
// 	return ability
// }
