package rule

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	usecase "github.com/betorvs/playbypost-dnd/usecase/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

// ListContent controller
func ListContent(c echo.Context) (err error) {
	kind := c.Param("kind")
	value := c.Param("value")
	allowedKindList := []string{"class", "race", "subrace", "background", "alignment", "ability", "skills", "conditions", "condition", "damage", "damagetype"}
	if !utils.StringInSlice(kind, allowedKindList) {
		errString := "Kind not implemented"
		return c.JSON(http.StatusNotImplemented, utils.FormatMessage(errString))
	}
	result, err := usecase.ListInformation(kind, value)
	if err != nil {
		errString := "Cannot list it"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}

// GetDescription controller
func GetDescription(c echo.Context) (err error) {
	kind := c.Param("kind")
	name := c.Param("name")
	subname := c.Param("subname")
	allowedKindDescription := []string{"class", "race", "background", "conditions", "condition", "damage", "damagetype"}
	if !utils.StringInSlice(kind, allowedKindDescription) {
		errString := "Kind not implemented"
		return c.JSON(http.StatusNotImplemented, utils.FormatMessage(errString))
	}

	result, err := usecase.FullDescription(kind, name, subname)
	if err != nil {
		errString := "Cannot describe it"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}

// CheckRoll controller
func CheckRoll(c echo.Context) (err error) {
	kind := c.Param("kind")
	allowedKind := []string{"ability"}
	if !utils.StringInSlice(kind, allowedKind) {
		errString := "Kind not implemented"
		return c.JSON(http.StatusNotImplemented, utils.FormatMessage(errString))
	}
	switch kind {
	case "ability":
		// 	result, err := usecase.()
		return c.JSON(http.StatusOK, "OK")

	default:
		err := fmt.Errorf("not implemented")
		return c.JSON(http.StatusNotImplemented, err)
	}
}

// Character controller
func Character(c echo.Context) (err error) {
	character := new(rule.NewCharacter)
	if err = c.Bind(character); err != nil {
		// fmt.Printf("%+v\n", character)
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	validMessage, validBool := validateNames(character)
	if !validBool {
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(validMessage))
	}

	result, err := usecase.CalculateCharacter(character)
	if err != nil {
		errString := "Cannot calculate character post"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusCreated, result)
}

func validateNames(req *rule.NewCharacter) (string, bool) {
	// fmt.Println(req)
	if !utils.StringInSlice(req.Race, usecase.RaceList()) {
		return "Race invalid", false
	}
	if utils.StringInSlice(req.Race, usecase.RaceListWithSubrace()) && !utils.StringInSlice(req.Subrace, usecase.SubraceList("")) {
		return "Subrace invalid", false
	}
	if !utils.StringInSlice(req.Class, usecase.ClassList()) {
		return "Class invalid", false
	}
	if !utils.StringInSlice(req.Background, usecase.BackgroundList()) {
		return "Background invalid", false
	}
	if len(req.Ability) != 6 {
		return "Ability number invalid", false
	}
	for k := range req.Ability {
		if !utils.StringInSlice(k, usecase.AbilityList()) {
			return "Ability name error", false
		}
	}
	// languages check
	_, _, _, _, _, languages, _, _, _, _, _ := usecase.RaceStatistics(req.Race, "")
	var numberOfLanguages int
	if req.Background == "acolyte" || req.Background == "sage" {
		numberOfLanguages = 2
	}
	if req.Race == "human" || req.Race == "half-elf" {
		numberOfLanguages++
	}
	if req.Subrace == "high-elf" {
		numberOfLanguages++
	}
	if len(req.ChosenLanguages) != numberOfLanguages {
		return fmt.Sprintf("Number of Languages wrongly %v must be %v", len(req.ChosenLanguages), numberOfLanguages), false
	}
	for _, v := range req.ChosenLanguages {
		if utils.StringInSlice(v, languages) {
			return fmt.Sprintf("Language %s already present in your Race", v), false
		}
	}
	// skills check
	_, skills := usecase.BackgroundStatistics(req.Background)
	for _, v := range req.ChosenSkills {
		if utils.StringInSlice(v, skills) {
			return fmt.Sprintf("Skill %s already present in your Background", v), false
		}
	}
	// abilities check
	if req.Race == "half-elf" {
		if len(req.ChosenAbility) != 2 {
			return "Please choose 2 abilities to be increased by one", false
		}
		if utils.StringInSlice("charisma", req.ChosenAbility) {
			return "Please choose another ability. Cannot choose Charisma twice.", false
		}
	}
	return "", true
}

// CheckFullAttack func
func CheckFullAttack(c echo.Context) (err error) {
	attack := new(rule.Attack)
	if err = c.Bind(attack); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcFullAttackwithWeapon(attack)
	return c.JSON(http.StatusOK, result)
}

// CheckSpellAttack func
func CheckSpellAttack(c echo.Context) (err error) {
	attack := new(rule.SpellcastAbility)
	if err = c.Bind(attack); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	denyList := []string{"barbarian", "fighter", "monk", "rogue"}
	if utils.StringInSlice(attack.Class, denyList) {
		errString := "Class without spellcast ability"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcSpellcastAttackAndSave(attack)
	return c.JSON(http.StatusOK, result)
}

// GetCondition controller
func GetCondition(c echo.Context) (err error) {
	condition := c.Param("name")
	if !utils.StringInSlice(condition, usecase.ListConditions()) {
		errString := "Condition doenst exist"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	var level int
	initialLevel := c.Param("level")
	if initialLevel != "" {
		level, err = strconv.Atoi(initialLevel)
		if err != nil {
			errString := "Value not allowed. Use valid int"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
		if level < 1 || level > 6 {
			errString := "Value not allowed. Use valid int: 1 to 6"
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
	}

	result := usecase.GetConditions(condition, level)
	return c.JSON(http.StatusOK, result)
}

// CheckSkillOrAbility controller
func CheckSkillOrAbility(c echo.Context) (err error) {
	check := new(rule.SkillOrAbilityCheck)
	if err = c.Bind(check); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcSkillOrAbility(check)
	return c.JSON(http.StatusOK, result)
}

// CheckSavingsAbility controller
func CheckSavingsAbility(c echo.Context) (err error) {
	attack := new(rule.SavingsCheck)
	if err = c.Bind(attack); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcSavingsAbility(attack)
	return c.JSON(http.StatusOK, result)
}

// CalcArmorClass controller
func CalcArmorClass(c echo.Context) (err error) {
	ac := new(rule.ArmorClass)
	if err = c.Bind(ac); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcArmorClass(ac)
	return c.JSON(http.StatusOK, result)
}

// UseClassFeature controller
func UseClassFeature(c echo.Context) (err error) {
	feature := new(rule.Feature)
	if err = c.Bind(feature); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}

	result, err := usecase.ClassFeatureRoll(feature)
	if err != nil {
		errString := fmt.Sprintf("error %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}

// UseRaceFeature controller
func UseRaceFeature(c echo.Context) (err error) {
	feature := new(rule.SpecialRaceFeature)
	if err = c.Bind(feature); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}

	result, err := usecase.SpecialRaceFeature(feature)
	if err != nil {
		errString := fmt.Sprintf("You cannot use %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}

// CheckPreparedSpellList controller
func CheckPreparedSpellList(c echo.Context) (err error) {
	spellList := new(rule.PreparedSpellsList)
	if err = c.Bind(spellList); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !utils.StringInSlice(spellList.Class, usecase.ClassWithPreparedSpell()) {
		errString := "class dont need to prepare spell list"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := usecase.CheckPreparedSpellsList(spellList)
	if err != nil {
		errString := fmt.Sprintf("You cannot use %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}

// CheckKnownSpellList controller
func CheckKnownSpellList(c echo.Context) (err error) {
	spellList := new(rule.KnownSpellsList)
	if err = c.Bind(spellList); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	class := spellList.Class
	if !utils.StringInSlice(class, usecase.ClassWithCantrips()) {
		errString := "your class dont have cantrips"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := usecase.CheckKnownList(spellList)
	if err != nil {
		errString := fmt.Sprintf("You cannot use %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}

// CheckCantripList controller
func CheckCantripList(c echo.Context) (err error) {
	spellList := new(rule.KnownCantripList)
	if err = c.Bind(spellList); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	class := spellList.Class
	if spellList.Class == "fighter" || spellList.Class == "rogue" {
		if utils.StringInSlice("archetype-eldritch-knight-spellcasting", spellList.ClassFeatures) {
			class = "eldritch-knight"
		}
		if utils.StringInSlice("archetype-arcane-trickster-spellcasting", spellList.ClassFeatures) {
			class = "arcane-trickster"
		}
	}
	if !utils.StringInSlice(class, usecase.ClassWithCantrips()) {
		errString := "your class dont have cantrips"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := usecase.CheckCantripsKnownList(spellList)
	if err != nil {
		errString := fmt.Sprintf("You cannot use %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}

// CheckMonsterAttack controller
func CheckMonsterAttack(c echo.Context) (err error) {
	attack := new(rule.MonsterRoll)
	if err = c.Bind(attack); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	allowedChecks := []string{"attack", "strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"}
	if !utils.StringInSlice(attack.Check, allowedChecks) {
		errString := "Check not allowed: attack, strength, dexterity, constitution, intelligence, wisdom, charisma"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcFullMonsterAttackwithWeapon(attack)
	return c.JSON(http.StatusOK, result)
}

// CheckMonsterSavings controller
func CheckMonsterSavings(c echo.Context) (err error) {
	saving := new(rule.MonsterRoll)
	if err = c.Bind(saving); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcMonsterSavingsAbility(saving)
	return c.JSON(http.StatusOK, result)
}

// CheckMonsterChecks controller
func CheckMonsterChecks(c echo.Context) (err error) {
	saving := new(rule.MonsterRoll)
	if err = c.Bind(saving); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcMonsterChecks(saving)
	return c.JSON(http.StatusOK, result)
}

// CheckMonstersInitiative controller
func CheckMonstersInitiative(c echo.Context) (err error) {
	list := new(rule.SimpleList)
	if err = c.Bind(list); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result := usecase.CalcMonstersInitiative(list)
	return c.JSON(http.StatusOK, result)
}

// CheckMonsterTurn controller
func CheckMonsterTurn(c echo.Context) (err error) {
	monster := new(rule.MonsterTurn)
	if err = c.Bind(monster); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := usecase.TurnUndeadRolls(monster)
	if err != nil {
		errString := fmt.Sprintf("error %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}

// UsePotionPlayer controller
func UsePotionPlayer(c echo.Context) (err error) {
	potion := new(rule.Potion)
	if err = c.Bind(potion); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := usecase.UsePotion(potion)
	if err != nil {
		errString := "Potion not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, result)
}
