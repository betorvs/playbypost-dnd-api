package rule

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestListContent(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/list/:kind")
	c.SetParamNames("kind")
	c.SetParamValues("background")

	// Assertions
	if assert.NoError(t, ListContent(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/rule/list/:kind")
	c1.SetParamNames("kindy")
	c1.SetParamValues("background")

	// Assertions
	if assert.NoError(t, ListContent(c1)) {
		assert.Equal(t, http.StatusNotImplemented, rec1.Code)
	}
}

func TestGetDescription(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/description/:kind/:name")
	c.SetParamNames("kind", "name")
	c.SetParamValues("background", "soldier")

	// Assertions
	if assert.NoError(t, GetDescription(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/rule/description/:kind/:name")
	c1.SetParamNames("kindy", "name")
	c1.SetParamValues("background", "soldier")

	// Assertions
	if assert.NoError(t, GetDescription(c1)) {
		assert.Equal(t, http.StatusNotImplemented, rec1.Code)
	}
}

func TestGetCondition(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/condition/:name")
	c.SetParamNames("name")
	c.SetParamValues("blinded")

	// Assertions
	if assert.NoError(t, GetCondition(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/rule/condition/:name")
	c1.SetParamNames("name")
	c1.SetParamValues("crazy")

	// Assertions
	if assert.NoError(t, GetCondition(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}
	//
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/rule/condition/:name/:level")
	c2.SetParamNames("name", "level")
	c2.SetParamValues("exhaustion", "severe")

	// Assertions
	if assert.NoError(t, GetCondition(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	req3 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/rule/condition/:name/:level")
	c3.SetParamNames("name", "level")
	c3.SetParamValues("exhaustion", "7")

	// Assertions
	if assert.NoError(t, GetCondition(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}
}

// func TestCheckRoll(t *testing.T) {}

func TestCharacter(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/character")

	// Assertions
	if assert.NoError(t, Character(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.NewCharacter{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/character")

	// Assertions
	if assert.NoError(t, Character(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.NewCharacter{
		Level:      1,
		Class:      "cleric",
		Race:       "human",
		Subrace:    "",
		Background: "soldier",
		Ability: map[string]int{
			"strength":     14,
			"dexterity":    10,
			"constitution": 14,
			"intelligence": 10,
			"wisdom":       16,
			"charisma":     14,
		},
		ChosenLanguages:     []string{"celestial"},
		ChosenSkills:        []string{"religion", "perception"},
		ChosenClassFeatures: []string{"war"},
	}

	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/character")

	// Assertions
	if assert.NoError(t, Character(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := rule.NewCharacter{
		Level:      1,
		Class:      "cleric",
		Race:       "human",
		Subrace:    "",
		Background: "soldier",
		Ability: map[string]int{
			"strength":     14,
			"dexterity":    10,
			"constitution": 14,
			"intelligence": 10,
			"wisdom":       16,
			"charisma":     14,
		},
		ChosenLanguages:     []string{"celestial"},
		ChosenSkills:        []string{"religion", "medicine"},
		ChosenClassFeatures: []string{"war"},
	}

	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/roll/character")

	// Assertions
	if assert.NoError(t, Character(c3)) {
		assert.Equal(t, http.StatusCreated, rec3.Code)
	}
}

func TestValidateNames(t *testing.T) {
	example1 := rule.NewCharacter{}
	example1.Race = "superhuman"
	_, res1 := validateNames(&example1)
	assert.False(t, res1)
	example2 := rule.NewCharacter{}
	example2.Race = "elf"
	example2.Subrace = "superelf"
	_, res2 := validateNames(&example2)
	assert.False(t, res2)
	example3 := rule.NewCharacter{}
	example3.Race = "elf"
	example3.Subrace = "high-elf"
	example3.Class = "superhero"
	_, res3 := validateNames(&example3)
	assert.False(t, res3)
	example4 := rule.NewCharacter{}
	example4.Race = "elf"
	example4.Subrace = "high-elf"
	example4.Class = "fighter"
	example4.Background = "superhero"
	_, res4 := validateNames(&example4)
	assert.False(t, res4)
	example5 := rule.NewCharacter{}
	example5.Race = "elf"
	example5.Subrace = "high-elf"
	example5.Background = "acolyte"
	example5.Class = "fighter"
	_, res5 := validateNames(&example5)
	assert.False(t, res5)

	example6 := rule.NewCharacter{}
	example6.Race = "elf"
	example6.Subrace = "high-elf"
	example6.Background = "acolyte"
	example6.Class = "fighter"
	example6.Ability = map[string]int{
		"strength":     14,
		"dexterity":    10,
		"constitution": 14,
		"intelligence": 10,
		"wisdom":       16,
		"charisma":     16,
		"super":        14,
	}
	example6.Ability["super"] = 10
	_, res6 := validateNames(&example6)
	assert.False(t, res6)

	example7 := rule.NewCharacter{}
	example7.Race = "elf"
	example7.Subrace = "high-elf"
	example7.Background = "acolyte"
	example7.Class = "fighter"
	example7.Ability = map[string]int{
		"strength":     14,
		"dexterity":    10,
		"constitution": 14,
		"intelligence": 10,
		"wisdom":       16,
		"super":        14,
	}
	_, res7 := validateNames(&example7)
	assert.False(t, res7)

	example8 := rule.NewCharacter{}
	example8.Race = "human"
	example8.Subrace = ""
	example8.Background = "acolyte"
	example8.Class = "fighter"
	example8.Ability = map[string]int{
		"strength":     14,
		"dexterity":    10,
		"constitution": 14,
		"intelligence": 10,
		"wisdom":       16,
		"charisma":     14,
	}
	example8.ChosenLanguages = []string{"elvish", "giant", "gnomish", "draconic", "goblin", "halfling", "orc"}
	_, res8 := validateNames(&example8)
	assert.False(t, res8)

	example9 := rule.NewCharacter{}
	example9.Race = "elf"
	example9.Subrace = "high-elf"
	example9.Background = "acolyte"
	example9.Class = "fighter"
	example9.Ability = map[string]int{
		"strength":     14,
		"dexterity":    10,
		"constitution": 14,
		"intelligence": 10,
		"wisdom":       16,
		"charisma":     14,
	}
	example9.ChosenLanguages = []string{"elvish", "giant", "gnomish"}
	_, res9 := validateNames(&example9)
	assert.False(t, res9)

	example10 := rule.NewCharacter{}
	example10.Race = "elf"
	example10.Subrace = "high-elf"
	example10.Background = "acolyte"
	example10.Class = "fighter"
	example10.Ability = map[string]int{
		"strength":     14,
		"dexterity":    10,
		"constitution": 14,
		"intelligence": 10,
		"wisdom":       16,
		"charisma":     14,
	}
	example10.ChosenLanguages = []string{"giant", "gnomish", "draconic"}
	example10.ChosenSkills = []string{"insight", "intimidation"}
	_, res10 := validateNames(&example10)
	assert.False(t, res10)

	example11 := rule.NewCharacter{}
	example11.Race = "half-elf"
	example11.Subrace = ""
	example11.Background = "acolyte"
	example11.Class = "fighter"
	example11.Ability = map[string]int{
		"strength":     16,
		"dexterity":    10,
		"constitution": 14,
		"intelligence": 12,
		"wisdom":       12,
		"charisma":     12,
	}
	example11.ChosenLanguages = []string{"giant", "gnomish", "draconic"}
	example11.ChosenSkills = []string{"athletics", "intimidation"}
	example11.ChosenAbility = []string{"constitution", "strength", "constitution"}
	_, res11 := validateNames(&example11)
	assert.False(t, res11)

	example12 := rule.NewCharacter{}
	example12.Race = "half-elf"
	example12.Subrace = ""
	example12.Background = "acolyte"
	example12.Class = "fighter"
	example12.Ability = map[string]int{
		"strength":     16,
		"dexterity":    10,
		"constitution": 14,
		"intelligence": 12,
		"wisdom":       12,
		"charisma":     12,
	}
	example12.ChosenLanguages = []string{"giant", "gnomish", "draconic"}
	example12.ChosenSkills = []string{"athletics", "intimidation"}
	example12.ChosenAbility = []string{"constitution", "charisma"}
	_, res12 := validateNames(&example12)
	assert.False(t, res12)
}

func TestCheckFullAttack(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/attack")

	// Assertions
	if assert.NoError(t, CheckFullAttack(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.Attack{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/attack")

	// Assertions
	if assert.NoError(t, CheckFullAttack(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}
}

func TestCheckSpellAttack(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/attack")

	// Assertions
	if assert.NoError(t, CheckSpellAttack(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.SpellcastAbility{}
	example1.Class = "barbarian"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/attack")

	// Assertions
	if assert.NoError(t, CheckSpellAttack(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.SpellcastAbility{}
	example2.Class = "sorcerer"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/player/attack")

	// Assertions
	if assert.NoError(t, CheckSpellAttack(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}
}

func TestCheckSkillOrAbility(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/check")

	// Assertions
	if assert.NoError(t, CheckSkillOrAbility(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.SkillOrAbilityCheck{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/check")

	// Assertions
	if assert.NoError(t, CheckSkillOrAbility(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}
}

func TestCheckSavingsAbility(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/savings")

	// Assertions
	if assert.NoError(t, CheckSavingsAbility(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.SavingsCheck{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/savings")

	// Assertions
	if assert.NoError(t, CheckSavingsAbility(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}
}

func TestCalcArmorClass(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/armorclass")

	// Assertions
	if assert.NoError(t, CalcArmorClass(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.ArmorClass{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/armorclass")

	// Assertions
	if assert.NoError(t, CalcArmorClass(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}
}

func TestUseClassFeature(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/classfeature")

	// Assertions
	if assert.NoError(t, UseClassFeature(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.Feature{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/classfeature")

	// Assertions
	if assert.NoError(t, UseClassFeature(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.Feature{}
	example2.Name = "bless"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/player/classfeature")

	// Assertions
	if assert.NoError(t, UseClassFeature(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}
}

func TestUseRaceFeature(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/racefeature")

	// Assertions
	if assert.NoError(t, UseRaceFeature(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.SpecialRaceFeature{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/racefeature")

	// Assertions
	if assert.NoError(t, UseRaceFeature(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.SpecialRaceFeature{}
	example2.Name = "hellish-rebuke"
	example2.Race = "tiefling"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/player/racefeature")

	// Assertions
	if assert.NoError(t, UseRaceFeature(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}
}

func TestCheckPreparedSpellList(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/preparedspells")

	// Assertions
	if assert.NoError(t, CheckPreparedSpellList(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.PreparedSpellsList{}
	example1.Class = "sorcerer"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/preparedspells")

	// Assertions
	if assert.NoError(t, CheckPreparedSpellList(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.PreparedSpellsList{}
	example2.Class = "wizard"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/player/preparedspells")

	// Assertions
	if assert.NoError(t, CheckPreparedSpellList(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := rule.PreparedSpellsList{}
	example3.Class = "wizard"
	example3.Level = 1
	example3.PreparedSpells = []string{"chill-touch", "chill-touch", "chill-touch"}
	example3.Ability = map[string]int{
		"strength":     14,
		"dexterity":    10,
		"constitution": 14,
		"intelligence": 14,
		"wisdom":       16,
		"charisma":     14,
	}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/roll/player/preparedspells")

	// Assertions
	if assert.NoError(t, CheckPreparedSpellList(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestCheckKnownSpellList(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/knownspells")

	// Assertions
	if assert.NoError(t, CheckKnownSpellList(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.KnownSpellsList{}
	example1.Class = "wizard"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/knownspells")

	// Assertions
	if assert.NoError(t, CheckKnownSpellList(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.KnownSpellsList{}
	example2.Level = 1
	example2.Class = "sorcerer"
	example2.SpellList = []string{"chill-touch", "chill-touch", "chill-touch", "chill-touch", "chill-touch", "chill-touch"}
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/player/knownspells")

	// Assertions
	if assert.NoError(t, CheckKnownSpellList(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := rule.KnownSpellsList{}
	example3.Class = "sorcerer"
	example3.Level = 1
	example3.KnownSpells = 2
	example3.SpellList = []string{"chill-touch", "chill-touch"}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/roll/player/knownspells")

	// Assertions
	if assert.NoError(t, CheckKnownSpellList(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestCheckCantripList(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/cantrips")

	// Assertions
	if assert.NoError(t, CheckCantripList(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.KnownCantripList{}
	example1.Class = "paladin"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/cantrips")

	// Assertions
	if assert.NoError(t, CheckCantripList(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.KnownCantripList{}
	example2.Class = "fighter"
	example2.ClassFeatures = []string{"archetype-eldritch-knight-spellcasting"}
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/player/cantrips")

	// Assertions
	if assert.NoError(t, CheckCantripList(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}

	example3 := rule.KnownCantripList{}
	example3.Class = "rogue"
	example3.ClassFeatures = []string{"archetype-arcane-trickster-spellcasting"}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/roll/player/cantrips")

	// Assertions
	if assert.NoError(t, CheckCantripList(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
	//acid-splash
	example4 := rule.KnownCantripList{}
	example4.Class = "sorcerer"
	example4.CantripsKnown = 2
	example4.CantripsList = []string{"acid-splash", "chill-touch", "produce-flame"}
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPost, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/roll/player/cantrips")

	// Assertions
	if assert.NoError(t, CheckCantripList(c4)) {
		assert.Equal(t, http.StatusBadRequest, rec4.Code)
	}
}

func TestCheckMonsterAttack(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/npc/attack")

	// Assertions
	if assert.NoError(t, CheckMonsterAttack(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.MonsterRoll{}
	example1.Check = "attack"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/npc/attack")

	// Assertions
	if assert.NoError(t, CheckMonsterAttack(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}

	example2 := rule.MonsterRoll{}
	example2.Check = "attacking"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/npc/attack")

	// Assertions
	if assert.NoError(t, CheckMonsterAttack(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}
}

func TestCheckMonsterSavings(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/npc/savings")

	// Assertions
	if assert.NoError(t, CheckMonsterSavings(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.MonsterRoll{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/npc/savings")

	// Assertions
	if assert.NoError(t, CheckMonsterSavings(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}
}

func TestCheckMonsterChecks(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/npc/checks")

	// Assertions
	if assert.NoError(t, CheckMonsterChecks(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.MonsterRoll{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/npc/checks")

	// Assertions
	if assert.NoError(t, CheckMonsterChecks(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}
}

func TestCheckMonstersInitiative(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/npc/initiative")

	// Assertions
	if assert.NoError(t, CheckMonstersInitiative(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.SimpleList{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/npc/initiative")

	// Assertions
	if assert.NoError(t, CheckMonstersInitiative(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}
}

func TestCheckMonsterTurn(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/npc/initiative")

	// Assertions
	if assert.NoError(t, CheckMonsterTurn(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.MonsterTurn{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/npc/turnundead")

	// Assertions
	if assert.NoError(t, CheckMonsterTurn(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.MonsterTurn{}
	example2.ClassFeatures = []string{"sacred-oath-of-devotion-channel-divinity"}
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/npc/turnundead")

	// Assertions
	if assert.NoError(t, CheckMonsterTurn(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}
}

func TestUsePotionPlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/roll/player/potion")

	// Assertions
	if assert.NoError(t, UsePotionPlayer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.Potion{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/roll/player/potion")

	// Assertions
	if assert.NoError(t, UsePotionPlayer(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := rule.Potion{}
	example2.Name = "potion-of-climbing"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/roll/player/potion")

	// Assertions
	if assert.NoError(t, UsePotionPlayer(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}
}
