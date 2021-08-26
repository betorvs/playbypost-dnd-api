package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	localtest "github.com/betorvs/playbypost-dnd/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRunDiceRoll(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, localtest.InitDiceMock)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/playbypost-dnd/v1/roll/:dice")
	c.SetParamNames("dice")
	c.SetParamValues("1d20")

	// Assertions
	if assert.NoError(t, RunDiceRoll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	// BadRequest test
	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPut, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)
	c1.SetPath("/playbypost-dnd/v1/roll/:dice")
	c1.SetParamNames("dice")
	c1.SetParamValues("1d20h")

	// Assertions
	if assert.NoError(t, RunDiceRoll(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}
}
