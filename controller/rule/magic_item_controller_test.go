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

func TestGetMagicItem(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rule/magicitem")

	// Assertions
	if assert.NoError(t, GetMagicItem(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
