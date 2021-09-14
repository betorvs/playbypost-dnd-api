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

func TestGetSpellListDescription(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/spell")

	// Assertions
	if assert.NoError(t, GetSpellListDescription(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestListSpellByClass(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/spelllist/:class/:level")
	c.SetParamNames("class", "level")
	c.SetParamValues("bard", "1")

	// Assertions
	if assert.NoError(t, ListSpellByClass(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
