package rule

import (
	"fmt"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//CalculateCharacter func
func CalculateCharacter(purpose *rule.NewCharacter) (*rule.Character, error) {
	characterFinal := new(rule.Character)
	characterFinal.Level = purpose.Level
	characterFinal.Race = purpose.Race
	var subrace string
	if purpose.Subrace != "" {
		subrace = purpose.Subrace
		characterFinal.Subrace = subrace
	}
	// save chosen options
	characterFinal.ChosenLanguages = purpose.ChosenLanguages
	characterFinal.ChosenSkills = purpose.ChosenSkills
	characterFinal.ChosenAbility = purpose.ChosenAbility
	// add empty slices
	characterFinal.Language = []string{}
	characterFinal.DamageResistence = []string{}
	characterFinal.DamageImmunities = []string{}
	characterFinal.DamageVulnerabilities = []string{}
	characterFinal.ConditionImmunities = []string{}
	characterFinal.Disvantages = []string{}
	characterFinal.MagicalEffect = []string{}
	// size, speedmeasure, speed, ability, special, language, skills, resistance, advantage, condition, disvantages := RaceStatistics(purpose.Race, subrace)
	raceStats := RaceStatistics(purpose.Race, subrace)
	characterFinal.RaceFeatures = raceStats.Special
	characterFinal.Size = raceStats.Size
	characterFinal.Speed = raceStats.Speed
	characterFinal.SpeedMeasure = raceStats.Speedmeasure
	characterFinal.Ability = raceStats.Ability
	if len(raceStats.Language) != 0 {
		characterFinal.Language = raceStats.Language
	}
	characterFinal.Skills = raceStats.Skills
	// check advantage, disvantage, resistances
	if len(raceStats.Resistance) != 0 {
		characterFinal.DamageResistence = raceStats.Resistance
	}
	characterFinal.Advantages = raceStats.Advantages
	if len(raceStats.Conditions) != 0 {
		characterFinal.ConditionImmunities = raceStats.Conditions
	}
	if len(raceStats.Disvantages) != 0 {
		characterFinal.Disvantages = raceStats.Disvantages
	}
	if len(raceStats.ArmorProficiency) != 0 {
		characterFinal.ArmorProficiency = raceStats.ArmorProficiency
	}
	characterFinal.Ability = purpose.Ability
	characterFinal.AbilityModifier = make(map[string]int)
	for k := range purpose.Ability {
		if purpose.Race == "half-elf" && utils.StringInSlice(k, purpose.ChosenAbility) {
			raceStats.Ability[k]++
		}
		if raceStats.Ability[k] != 0 {
			characterFinal.Ability[k] += raceStats.Ability[k]
		}
	}
	// class options
	class := ClassStatistics(purpose.Class, purpose.Level)
	characterFinal.Class = class.Name
	// hitDice, savings, armorProficiency, skillNumber := ClassDetails(purpose.Class)

	characterFinal.HitDice = fmt.Sprintf("d%v", class.HitDice)
	characterFinal.Savings = class.Savings
	// add more armor and weapon proficiency
	characterFinal.ArmorProficiency = append(characterFinal.ArmorProficiency, class.ArmorProficiency...)
	// if len(armorProficiency) != 0 {
	// 	characterFinal.ArmorProficiency = armorProficiency
	// }
	// for _, s := range characterFinal.RaceFeatures {
	// 	// fmt.Println(s)
	// 	newList := RaceArmorProficiencyExtra(s)
	// 	if len(newList) != 0 {
	// 		armorProficiency = append(armorProficiency, newList...)
	// 	}
	// 	// armorAndWeaponProficiency = RaceArmorProficiencyExtra(s, armorProficiency)
	// }
	// characterFinal.ArmorProficiency = armorProficiency
	// class features options
	characterFinal.ClassFeatures = CalculateClassFeatureList(class.Features, purpose.Level, purpose.ChosenClassFeatures)
	// ability by level adjustment
	var abilityIncrementByLevel int
	for _, v := range characterFinal.ClassFeatures {
		if v == "ability-score-improvement" {
			abilityIncrementByLevel++
		}
	}
	if purpose.Level > 1 && abilityIncrementByLevel != 0 {
		for _, a := range purpose.ChosenAbilityByLevel {
			characterFinal.Ability[a]++
			// characterFinal.AbilityModifier[a] = CalcAbilityModifier(characterFinal.Ability[a])
		}
		characterFinal.ChosenAbilityByLevel = purpose.ChosenAbilityByLevel
	}
	// calc modifiers only after changing abitily itself
	for k, v := range characterFinal.Ability {
		characterFinal.AbilityModifier[k] = CalcAbilityModifier(v)
	}
	// fmt.Println(armorProficiency)

	characterFinal.XPNextLevel = XPNeeded(purpose.Level)
	characterFinal.Proficiency = CalcProficiency(purpose.Level)
	// hit points
	hpModifier := characterFinal.AbilityModifier["constitution"]
	if purpose.Subrace == "hill-dwarf" {
		hpModifier = characterFinal.AbilityModifier["constitution"] + 1
	}
	characterFinal.HPMax = CalcMaxHP(purpose.Level, class.HitDice, hpModifier)

	background := BackgroundStatistics(purpose.Background)
	characterFinal.Background = background.Name
	// languages chosen
	var numberOfLanguages int
	// if purpose.Background == "acolyte" || purpose.Background == "sage" {
	if background.AdditionalLanguages != 0 {
		numberOfLanguages = 2
	}
	// if purpose.Race == "human" || purpose.Race == "half-elf" {
	if raceStats.AdditionalLanguages != 0 {
		numberOfLanguages = numberOfLanguages + raceStats.AdditionalLanguages
	}
	var extraLanguageMessage string
	verified, err := languagesAdded(purpose.ChosenLanguages, characterFinal.Language, numberOfLanguages)
	if err != nil {
		extraLanguageMessage = fmt.Sprintf(" %v", err)
	}
	characterFinal.Language = append(characterFinal.Language, verified...)
	// skills list
	characterFinal.Skills = append(characterFinal.Skills, background.Skills...)

	// Spells List
	if utils.StringInSlice(purpose.Class, ClassWithSpell()) {
		spellListLevel, spellMaxLevel := CalculateSpellList(class.Name, purpose.Level)

		if len(spellListLevel) != 0 && purpose.Class != "warlock" {
			characterFinal.SpellListLevel = spellListLevel
		}
		characterFinal.SpellMaxLevel = spellMaxLevel
	}
	// Spells Cantrips
	if utils.StringInSlice(purpose.Class, ClassWithCantrips()) {
		// characterFinal.CantripsKnown = CantripsKnown(purpose.Class, purpose.Level)
		characterFinal.CantripsKnown = class.CantripsKnown[purpose.Level]
	}
	// Spells Known
	if utils.StringInSlice(purpose.Class, ClassWithSpellKnown()) {
		// characterFinal.SpellKnown = SpellKnown(purpose.Class, purpose.Level)
		characterFinal.SpellKnown = class.SpellKnown[purpose.Level]
	}
	// ranger options to check if was completed or not
	var favoredEnemy string
	var naturalExplorer string
	var favoredEnemyLanguages string
	var enemyLanguages []string

	var extraSkillsMessage string
	// extraSkills
	var skillsByFeatureMessage string
	// extra languages
	var extraLanguageByFeatureMessage string

	// Skills
	verified, err = skillsAdded(purpose.ChosenSkills, characterFinal.Skills, skillListByClass(class.Name), class.SkillNumber)
	if err != nil {
		extraSkillsMessage = fmt.Sprintf(" %v", err)
	}
	characterFinal.Skills = append(characterFinal.Skills, verified...)

	// var archetypeMessage string
	var archetypeMessage string

	// Warlock Special Features
	if class.WarlockSpellSlots != nil {
		characterFinal.WarlockSpellSlots = class.WarlockSpellSlots[purpose.Level]
	}
	if class.WarlockSlotLevel != nil {
		characterFinal.WarlockSlotLevel = class.WarlockSlotLevel[purpose.Level]
	}

	if class.WarlockInvocationsKnown != nil {
		characterFinal.WarlockInvocationsKnown = class.WarlockInvocationsKnown[purpose.Level]
	}

	// class features
	switch purpose.Class {
	case "barbarian":

		rage, damage := BarbarianClass(purpose.Level)
		characterFinal.BarbarianRage = rage
		characterFinal.BarbarianDamage = damage
		var barbarian string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, barbarianPrimalPathList()) {
				barbarian = f
			}
		}
		if barbarian == "" && purpose.Level >= 3 {
			archetypeMessage = fmt.Sprintf("primal path choose %v", barbarianPrimalPathList())
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "primal-path")

	case "bard":
		var bard string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, bardCollege()) {
				bard = f
			}
		}
		if bard == "" && purpose.Level >= 3 {
			archetypeMessage = fmt.Sprintf("bard-college choose %v", bardCollege())
		}
		if utils.StringInSlice("college-of-lore-bonus-proficiencies", characterFinal.ClassFeatures) {
			if len(purpose.ChosenSkillsByFeatures) != 0 {
				verified, err := skillsAdded(purpose.ChosenSkillsByFeatures, characterFinal.Skills, []string{}, 3)
				if err != nil {
					skillsByFeatureMessage = fmt.Sprintf(" %v", err)
				}
				characterFinal.Skills = append(characterFinal.Skills, verified...)
			}
			if len(purpose.ChosenSkillsByFeatures) == 0 {
				skillsByFeatureMessage = fmt.Sprintf(" college-of-lore-bonus-proficiencies choose 3 extra skills %v", SkillList())
			}
		}

		if utils.StringInSlice("college-of-valor-bonus-proficiencies", characterFinal.ClassFeatures) {
			characterFinal.ArmorProficiency = append(characterFinal.ArmorProficiency, []string{"medium-armor", "shields", "martial-weapon"}...)
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "bard-college")

	case "cleric":
		var cleric string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, clericDivineDomainList()) {
				cleric = f
			}
		}
		if cleric == "" {
			archetypeMessage = fmt.Sprintf("cleric-domain choose %v", clericDivineDomainList())
		}
		if utils.StringInSlice("domain-light-bonus-cantrip", characterFinal.ClassFeatures) {
			characterFinal.CantripsKnown++
		}
		// "domain-nature-acolyte-of-nature-and-bonus-proficiency" +1 +druid list
		if utils.StringInSlice("acolyte-of-nature", characterFinal.ClassFeatures) {
			characterFinal.CantripsKnown++
			skillAllowed := []string{"animal-handling", "nature", "survival"}
			if len(purpose.ChosenSkillsByFeatures) != 0 {
				verified, err := skillsAdded(purpose.ChosenSkillsByFeatures, characterFinal.Skills, skillAllowed, 1)
				if err != nil {
					skillsByFeatureMessage = fmt.Sprintf(" %v", err)
				}
				characterFinal.Skills = append(characterFinal.Skills, verified...)
			}
			if len(purpose.ChosenSkillsByFeatures) == 0 {
				skillsByFeatureMessage = fmt.Sprintf(" acolyte-of-nature need to choose 1 extra skills %v", skillAllowed)
			}
		}
		if utils.StringInSlice("blessings-of-knowledge", characterFinal.ClassFeatures) {
			// skills
			skillAllowed := blessingsOfKnowledge()
			if len(purpose.ChosenSkillsByFeatures) != 0 {
				verified, err := skillsAdded(purpose.ChosenSkillsByFeatures, characterFinal.Skills, skillAllowed, 2)
				if err != nil {
					skillsByFeatureMessage = fmt.Sprintf(" %v", err)
				}
				characterFinal.Skills = append(characterFinal.Skills, verified...)
			}
			if len(purpose.ChosenSkillsByFeatures) == 0 {
				skillsByFeatureMessage = fmt.Sprintf(" blessings-of-knowledge need to choose 2 extra skills %v", skillAllowed)
			}
			// languages
			if len(purpose.ChosenLanguagesByFeatures) != 0 {
				verified, err := languagesAdded(purpose.ChosenLanguagesByFeatures, characterFinal.Language, 2)
				if err != nil {
					extraLanguageByFeatureMessage = fmt.Sprintf(" %v", err)
				}
				characterFinal.Language = append(characterFinal.Language, verified...)
			}
			if len(purpose.ChosenLanguagesByFeatures) == 0 {
				extraLanguageByFeatureMessage = fmt.Sprintf(" blessings-of-knowledge choose 2 extra languages %v", LanguageList())
			}
		}
		if utils.StringInSlice("domain-light-bonus-cantrip", characterFinal.ClassFeatures) {
			characterFinal.CantripsKnown++
		}
		if utils.StringInSlice("domain-life-bonus-proficiency", characterFinal.ClassFeatures) {
			characterFinal.ArmorProficiency = append(characterFinal.ArmorProficiency, "heavy-armor")
		}
		if utils.StringInSlice("domain-nature-bonus-proficiency", characterFinal.ClassFeatures) {
			characterFinal.ArmorProficiency = append(characterFinal.ArmorProficiency, "heavy-armor")
		}
		if utils.StringInSlice("domain-tempest-bonus-proficiencies", characterFinal.ClassFeatures) {
			characterFinal.ArmorProficiency = append(characterFinal.ArmorProficiency, []string{"heavy-armor", "martial-weapon"}...)
		}
		if utils.StringInSlice("blessing-of-the-trickster", characterFinal.ClassFeatures) {
			characterFinal.Advantages = append(characterFinal.Advantages, "stealth")
		}
		if utils.StringInSlice("domain-war-bonus-proficiencies", characterFinal.ClassFeatures) {
			characterFinal.ArmorProficiency = append(characterFinal.ArmorProficiency, []string{"heavy-armor", "martial-weapon"}...)
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "divine-domain")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "divine-domain-feature")

	case "druid":
		var druid string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, druidCircleList()) {
				druid = f
			}
		}
		if druid == "" && purpose.Level >= 2 {
			archetypeMessage = fmt.Sprintf("druid-circle choose %v. if land choose one terrain %v", druidCircleList(), terrainList())
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "druid-circle")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "circle-of-the-land-circle-spells")
		// "circle-of-the-land-bonus-cantrip-and-natural-recovery-and-circle-spells-%s" terrainList()
		// druid feature bonus cantrip
		if utils.StringInSlice("circle-of-the-land-bonus-cantrip", characterFinal.ClassFeatures) {
			characterFinal.CantripsKnown++
		}

	case "fighter":
		var fighter string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, fighterArchetypeList()) {
				fighter = f
			}
		}
		if fighter == "" && purpose.Level >= 3 {
			archetypeMessage = fmt.Sprintf("martial-archetype choose %v", fighterArchetypeList())
		}
		// "archetype-eldritch-knight-spellcasting"
		if utils.StringInSlice("archetype-eldritch-knight-spellcasting", characterFinal.ClassFeatures) {
			spellListLevel, spellMaxLevel := CalculateSpellList("eldritch-knight", purpose.Level)

			characterFinal.SpellListLevel = spellListLevel
			characterFinal.SpellMaxLevel = spellMaxLevel
			characterFinal.CantripsKnown = CantripsKnown("eldritch-knight", purpose.Level)
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "fighting-style")
		if utils.StringInSlice("archetype-champion-additional-fighting-style", characterFinal.ClassFeatures) {
			var firstFightStyle string
			for _, v := range characterFinal.ClassFeatures {
				if utils.StringInSlice(v, fightingStyleFullNameList()) {
					firstFightStyle = strings.ReplaceAll(v, "fighting-style-", "")
				}
			}
			var secondFightStyle string
			for _, s := range purpose.ChosenClassFeatures {
				if s != firstFightStyle && secondFightStyle == "" {
					if utils.StringInSlice(s, fightingStyleList()) {
						secondFightStyle = s
					}
				}
			}
			fight := choosenClassFeatures("fighting-style", secondFightStyle, 1)
			characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, fight...)

		}

		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "martial-archetype")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "martial-archetype-feature")

	case "monk":
		martial, ki, movement := MonkClass(purpose.Level)
		characterFinal.MonkMartial = martial
		characterFinal.MonkKi = ki
		characterFinal.MonkMovement = movement
		var monk string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, monkMonasticTraditionList()) {
				monk = f
			}
		}
		if monk == "" && purpose.Level >= 3 {
			archetypeMessage = fmt.Sprintf("monastic-tradition choose %v", monkMonasticTraditionList())
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "monastic-tradition")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "monastic-tradition-feature")

	case "ranger":
		var enemyTypeList []string
		var enemyHumanoidList []string
		var terrainChoosen []string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, humanoidList()) {
				enemyHumanoidList = append(enemyHumanoidList, f)
			}
			if utils.StringInSlice(f, monsterTypeList()) {
				enemyTypeList = append(enemyTypeList, f)
			}
			if utils.StringInSlice(f, terrainList()) {
				terrainChoosen = append(terrainChoosen, f)
			}
		}
		var favoredEnemyNumber int
		var terrainNumber int
		for _, v := range characterFinal.ClassFeatures {
			if v == "favored-enemy" || v == "favored-enemy-improvement" {
				favoredEnemyNumber++
			}
			if v == "natural-explorer" || v == "natural-explorer-improvement" {
				terrainNumber++
			}
		}

		totalEnemies := len(enemyTypeList) + (len(enemyHumanoidList) / 2)
		if totalEnemies != favoredEnemyNumber {
			favoredEnemy = "wrong number of favored enemies "
		}
		if !utils.Even(len(enemyHumanoidList)) {
			favoredEnemy += "wrong number of favored enemies from humonoid list, should be odd "
		}
		// remove old entries for favored and natural
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "favored-enemy")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "favored-enemy-improvement")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "natural-explorer")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "natural-explorer-improvement")

		for _, s := range enemyTypeList {
			if utils.StringInSlice(s, monsterTypeList()) {
				enemy := choosenClassFeatures("favored-enemy", s, 1)
				characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, enemy...)
			}
		}
		for _, s := range enemyHumanoidList {
			if utils.StringInSlice(s, humanoidList()) {
				enemy := choosenClassFeatures("favored-enemy", s, 1)
				characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, enemy...)
			}
		}
		if len(terrainChoosen) != terrainNumber {
			naturalExplorer = "wrong number of natural explorer options "
		}
		for _, s := range terrainChoosen {
			if utils.StringInSlice(s, terrainList()) {
				enemy := choosenClassFeatures("natural-explorer", s, 1)
				characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, enemy...)
			}
		}

		for _, enemy := range enemyTypeList {
			enemies := returnMosterByType(monsterTypeSingular(enemy))
			for _, monster := range enemies {
				for _, lang := range monster.Languages {
					if utils.StringInSlice(lang, LanguageList()) && lang != "common" && !utils.StringInSlice(lang, enemyLanguages) {
						enemyLanguages = append(enemyLanguages, lang)
					}
				}
			}
		}
		for _, enemy := range enemyHumanoidList {
			enemies := MosterByName(enemy)
			for _, lang := range enemies.Languages {
				if utils.StringInSlice(lang, LanguageList()) && lang != "common" && !utils.StringInSlice(lang, enemyLanguages) {
					enemyLanguages = append(enemyLanguages, lang)
				}
			}
		}
		totalLanguages := len(enemyTypeList) + len(enemyHumanoidList)
		verified, err := languagesAdded(purpose.ChosenLanguagesByFeatures, characterFinal.Language, totalLanguages)
		if err != nil {
			favoredEnemyLanguages = fmt.Sprintf(" %v", err)
		}
		characterFinal.Language = append(characterFinal.Language, verified...)
		// archetype hunter
		var ranger string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, rangerArchetypesList()) {
				ranger = f
			}
		}
		if ranger == "" && purpose.Level >= 3 {
			archetypeMessage = fmt.Sprintf("ranger-archetype choose %v", rangerArchetypesList())
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "ranger-archetype")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "ranger-archetype-feature")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "fighting-style")
		if utils.StringInSlice("archetype-hunter-hunters-prey", characterFinal.ClassFeatures) {
			var hunterList []string
			for _, f := range purpose.ChosenClassFeatures {
				if utils.StringInSlice(f, huntersPreyList()) {
					hunterList = append(hunterList, f)
				}
			}
			if len(hunterList) == 0 {
				archetypeMessage = fmt.Sprintf("for ranger hunters prey you need to choose one of %v", huntersPreyList())
			}
			for _, f := range hunterList {
				if utils.StringInSlice(f, huntersPreyList()) {
					feature := choosenClassFeatures("ranger-archetype", f, 3)
					characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, feature...)
				}
			}
			characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "archetype-hunter-hunters-prey")
		}
		if utils.StringInSlice("archetype-hunter-defensive-tactics", characterFinal.ClassFeatures) {
			var hunterList []string
			for _, f := range purpose.ChosenClassFeatures {
				if utils.StringInSlice(f, defensiveTacticsList()) {
					hunterList = append(hunterList, f)
				}
			}
			if len(hunterList) == 0 {
				archetypeMessage = fmt.Sprintf("for ranger hunters defensive tactics you need to choose one of %v", defensiveTacticsList())
			}
			for _, f := range hunterList {
				if utils.StringInSlice(f, defensiveTacticsList()) {
					feature := choosenClassFeatures("ranger-archetype", f, 7)
					characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, feature...)
				}
			}
			characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "archetype-hunter-defensive-tactics")
		}
		if utils.StringInSlice("archetype-hunter-multiattack", characterFinal.ClassFeatures) {
			var hunterList []string
			for _, f := range purpose.ChosenClassFeatures {
				if utils.StringInSlice(f, multiattackList()) {
					hunterList = append(hunterList, f)
				}
			}
			if len(hunterList) == 0 {
				archetypeMessage = fmt.Sprintf("for ranger hunters multiattack you need to choose one of %v", multiattackList())
			}
			for _, f := range hunterList {
				if utils.StringInSlice(f, multiattackList()) {
					feature := choosenClassFeatures("ranger-archetype", f, 11)
					characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, feature...)
				}
			}
			characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "archetype-hunter-multiattack")
		}
		if utils.StringInSlice("archetype-hunter-superior-hunters-defense", characterFinal.ClassFeatures) {
			var hunterList []string
			for _, f := range purpose.ChosenClassFeatures {
				if utils.StringInSlice(f, huntersDefenseList()) {
					hunterList = append(hunterList, f)
				}
			}
			if len(hunterList) == 0 {
				archetypeMessage = fmt.Sprintf("for ranger hunters superior defense you need to choose one of %v", huntersDefenseList())
			}
			for _, f := range hunterList {
				if utils.StringInSlice(f, huntersDefenseList()) {
					feature := choosenClassFeatures("ranger-archetype", f, 15)
					characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, feature...)
				}
			}
			characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "archetype-hunter-superior-hunters-defense")
		}

	case "rogue":
		characterFinal.RogueSneak = RogueClass(purpose.Level)
		//"archetype-arcane-trickster-spellcasting-and-mage-hand-legerdemain"
		var rogue string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, rogueArchetypeList()) {
				rogue = f
			}
		}
		if rogue == "" && purpose.Level >= 3 {
			archetypeMessage = fmt.Sprintf("roguish-archetype choose %v", rogueArchetypeList())
		}
		if utils.StringInSlice("archetype-arcane-trickster-spellcasting-and-mage-hand-legerdemain", characterFinal.ClassFeatures) {
			spellListLevel, spellMaxLevel := CalculateSpellList("arcane-trickster", purpose.Level)

			characterFinal.SpellListLevel = spellListLevel
			characterFinal.SpellMaxLevel = spellMaxLevel
			characterFinal.CantripsKnown = CantripsKnown("arcane-trickster", purpose.Level)
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "roguish-archetype")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "roguish-archetype-feature")

	case "paladin":
		var paladin string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, paladinSacredOaths()) {
				paladin = f
			}
		}
		if paladin == "" && purpose.Level >= 3 {
			archetypeMessage = fmt.Sprintf("sacred-oaths choose %v", paladinSacredOaths())
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "sacred-oath")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "sacred-oath-feature")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "fighting-style")

	case "sorcerer":
		if purpose.Level != 1 {
			characterFinal.SorceryPoints = characterFinal.Level
		}
		var sorcerer string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, sorcererOriginList()) {
				sorcerer = f
			}
		}
		if sorcerer == "" {
			archetypeMessage = fmt.Sprintf("Please choose one %v", sorcererOriginList())
		}

		if utils.StringInSlice("sorcerous-origin-draconic-resistance", characterFinal.ClassFeatures) {
			var dragonType string
			for _, f := range purpose.ChosenClassFeatures {
				if utils.StringInSlice(f, dragonKinds()) {
					dragonType = f
				}
			}
			if dragonType == "" {
				archetypeMessage += fmt.Sprintf("Please choose one dragon color %v", dragonKinds())
			}
			characterFinal.Language = append(characterFinal.Language, "draconic")
			extraHP := purpose.Level
			characterFinal.HPMax = characterFinal.HPMax + extraHP

			// feature := choosenClassFeatures("sorcerous-origin", dragonType, 1)
			// characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, feature...)
			characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "sorcerous-origin-draconic-bloodline-dragon-ancestor")
		}

		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "sorcerous-origin")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "sorcerous-origin-feature")

	case "warlock":

		var warlock string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, warlockOtherworldlyPatronList()) {
				warlock = f
			}
		}
		if warlock == "" {
			archetypeMessage = fmt.Sprintf("Please choose one %v", warlockOtherworldlyPatronList())
		}

		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "otherworldly-patron")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "otherworldly-patron-feature")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "pact-boon")

	case "wizard":
		var wizard string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, wizardArcaneTraditionList()) {
				wizard = f
			}
		}
		if wizard == "" && purpose.Level >= 2 {
			archetypeMessage = fmt.Sprintf("Please choose one %v", wizardArcaneTraditionList())
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "arcane-tradition")
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "arcane-tradition-feature")

	}

	var expertiseChoosen string
	if utils.StringInSlice("expertise", characterFinal.ClassFeatures) {
		var expertiseNumber int
		for _, v := range characterFinal.ClassFeatures {
			if v == "expertise" {
				expertiseNumber = expertiseNumber + 2
			}
		}
		var expertiseList []string
		for _, f := range purpose.ChosenClassFeatures {
			if utils.StringInSlice(f, SkillList()) {
				expertiseList = append(expertiseList, f)
			}
		}
		// fmt.Println(len(expertiseList), expertiseNumber)
		if len(expertiseList) != expertiseNumber {
			expertiseChoosen = "wrong number of expertise options "
		}
		for _, e := range expertiseList {
			if !utils.StringInSlice(e, characterFinal.Skills) {
				expertiseChoosen += fmt.Sprintf(" %s is not in your skills proficiency list", e)
			}
		}
		characterFinal.ClassFeatures = utils.RemoveItemSlice(characterFinal.ClassFeatures, "expertise")
		for _, e := range expertiseList {
			if utils.StringInSlice(e, characterFinal.Skills) {
				expertise := fmt.Sprintf("expertise-%s", e)
				characterFinal.ClassFeatures = append(characterFinal.ClassFeatures, expertise)
			}
		}
	}
	// save options by level
	characterFinal.ChosenAbilityByLevel = purpose.ChosenAbilityByLevel
	characterFinal.ChosenClassFeatures = purpose.ChosenClassFeatures
	characterFinal.ChosenSkillsByFeatures = purpose.ChosenSkillsByFeatures
	characterFinal.ChosenLanguagesByFeatures = purpose.ChosenLanguagesByFeatures

	// final HP Temp
	characterFinal.HPTemp = characterFinal.HPMax

	// verify if all choosen options match with characters options
	incompleteList := make(map[string]string)
	// ChosenLanguages in controller
	if extraLanguageMessage != "" {
		incompleteList["chosen_languages"] = fmt.Sprintf(" %s", extraLanguageMessage)
	}
	// ChosenSkills in controller one check
	if extraSkillsMessage != "" {
		incompleteList["chosen_skills"] = fmt.Sprintf(" %s", extraSkillsMessage)
	}
	// languages and skills verification
	if extraLanguageByFeatureMessage != "" {
		incompleteList["chosen_languages_features"] = fmt.Sprintf(" %s", extraLanguageByFeatureMessage)
	}
	if skillsByFeatureMessage != "" {
		incompleteList["chosen_skills_features"] = fmt.Sprintf(" %s", skillsByFeatureMessage)
	}

	// ChosenAbilityByLevel
	if len(purpose.ChosenAbilityByLevel) != abilityIncrementByLevel {
		if abilityIncrementByLevel > len(purpose.ChosenAbilityByLevel) {
			incompleteList["chosen_ability_level"] += "you choose less abilities to increase that you received by level"
		}
		if abilityIncrementByLevel < len(purpose.ChosenAbilityByLevel) {
			incompleteList["chosen_ability_level"] += "you choose more abilities to increase that you received by level"
		}

	}
	// special class features
	if utils.StringInSlice("fighting-style", characterFinal.ClassFeatures) {
		incompleteList["chosen_class_features"] += fmt.Sprintf("you need to make choice for fighting-style: %v ", fightingStyleList())
	}
	if naturalExplorer != "" {
		incompleteList["chosen_class_features"] += fmt.Sprintf("%s you need to make choice for natural-explorer: %v ", naturalExplorer, terrainList())
	}
	if favoredEnemy != "" {
		incompleteList["chosen_class_features"] += fmt.Sprintf("%s you need to make choice for favored-enemy: %v or two from %v", favoredEnemy, monsterTypeList(), humanoidList())
	}
	if favoredEnemyLanguages != "" {
		incompleteList["chosen_class_features"] += fmt.Sprintf(" %s", favoredEnemyLanguages)
	}
	if expertiseChoosen != "" {
		incompleteList["chosen_class_features"] += fmt.Sprintf("%s you need to make choice for expertise: %v ", expertiseChoosen, characterFinal.Skills)
	}
	// all archetypes messages here
	if archetypeMessage != "" {
		incompleteList["chosen_class_features"] += fmt.Sprintf(" %s", archetypeMessage)
	}

	if utils.StringInSlice("pact-boon", characterFinal.ClassFeatures) {
		incompleteList["chosen_class_features"] += fmt.Sprintf("you need to make choice for pact-boon: %v ", warlockPactBoonList())
	}

	characterFinal.IncompleteOptions = incompleteList
	if len(characterFinal.IncompleteOptions) != 0 {
		errString := utils.PrintMapStringString(characterFinal.IncompleteOptions)
		return characterFinal, fmt.Errorf(errString)
	}
	return characterFinal, nil

}

func languagesAdded(extra, current []string, add int) (verified []string, err error) {
	var wrongly []string
	if len(extra) != add {
		err = fmt.Errorf("wrongly number of languages %v should be %v", extra, add)
	}
	for _, language := range extra {
		if utils.StringInSlice(language, current) {
			wrongly = append(wrongly, language)
			err = fmt.Errorf("you already have this language %s ", language)
		}
		if !utils.StringInSlice(language, LanguageList()) {
			wrongly = append(wrongly, language)
			err = fmt.Errorf("language not found %s ", language)
		}
		if !utils.StringInSlice(language, current) && utils.StringInSlice(language, LanguageList()) {
			verified = append(verified, language)
		}
	}
	if err != nil {
		return wrongly, err
	}
	return verified, nil
}

func skillsAdded(extra, current, allowed []string, add int) (verified []string, err error) {
	// fmt.Println(extra, current, allowed, add)
	var wrongly []string
	if len(extra) != add {
		err = fmt.Errorf("wrongly number of skills %v should be %v", extra, add)
		if len(allowed) != 0 {
			err = fmt.Errorf("wrongly number of skills %v should be %v from %v ", extra, add, allowed)
		}
	}
	for _, skill := range extra {
		if utils.StringInSlice(skill, current) {
			wrongly = append(wrongly, skill)
			err = fmt.Errorf("you already have this skill %s", skill)
		}
		if !utils.StringInSlice(skill, SkillList()) {
			wrongly = append(wrongly, skill)
			err = fmt.Errorf("skill not found %s ", skill)
		}
		if len(allowed) != 0 && !utils.StringInSlice(skill, allowed) {
			wrongly = append(wrongly, skill)
			err = fmt.Errorf("skill not allowed %s use from these list %v", skill, allowed)
		}
		if !utils.StringInSlice(skill, current) && utils.StringInSlice(skill, SkillList()) {
			verified = append(verified, skill)
		}

	}
	if err != nil {
		return wrongly, err
	}
	return verified, nil
}

func blessingsOfKnowledge() []string {
	return []string{"arcana", "history", "nature", "religion"}
}
