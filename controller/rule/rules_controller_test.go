package rule

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
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
}

func TestCheckRoll(t *testing.T) {}

func TestCharacter(t *testing.T) {}

func TestValidateNames(t *testing.T) {}

func TestCheckFullAttack(t *testing.T) {}

func TestCheckSpellAttack(t *testing.T) {}

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
}

func TestCheckSkillOrAbility(t *testing.T) {}

func TestCheckSavingsAbility(t *testing.T) {}

func TestCalcArmorClass(t *testing.T) {}

func TestUseClassFeature(t *testing.T) {}

func TestUseRaceFeature(t *testing.T) {}

func TestCheckPreparedSpellList(t *testing.T) {}

func TestCheckKnownSpellList(t *testing.T) {}

func TestCheckCantripList(t *testing.T) {}

func TestCheckMonsterAttack(t *testing.T) {}

func TestCheckMonsterSavings(t *testing.T) {}

func TestCheckMonsterChecks(t *testing.T) {}

func TestCheckMonstersInitiative(t *testing.T) {}

func TestCheckMonsterTurn(t *testing.T) {}

func TestUsePotionPlayer(t *testing.T) {}
