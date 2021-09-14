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

func TestGetAllWeapons(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/weapon")

	// Assertions
	if assert.NoError(t, GetAllWeapons(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllArmors(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/armor")

	// Assertions
	if assert.NoError(t, GetAllArmors(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllGear(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/gear")

	// Assertions
	if assert.NoError(t, GetAllGear(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllPacks(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/packs")

	// Assertions
	if assert.NoError(t, GetAllPacks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllTools(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/tools")

	// Assertions
	if assert.NoError(t, GetAllTools(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllMounts(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/mounts")

	// Assertions
	if assert.NoError(t, GetAllMounts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCalcShop(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	example := rule.SimpleList{
		List: []string{"horse"},
	}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/shops")

	// Assertions
	if assert.NoError(t, CalcShop(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestRandomTreasure(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/randomtreasure/:level")
	c.SetParamNames("level")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, RandomTreasure(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestFastTreasure(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/averagetreasure/:level")
	c.SetParamNames("level")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, RandomTreasure(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestRandomTreasureHoard(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/treasurehoard/:level")
	c.SetParamNames("level")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, RandomTreasure(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetAllServices(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/services")

	// Assertions
	if assert.NoError(t, GetAllServices(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
