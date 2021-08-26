package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/playbypost-dnd/v1/health")

	// Assertions
	if assert.NoError(t, CheckHealth(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetReady(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/playbypost-dnd/v1/ready")
	config.Values.IsReady.Store(true)
	// Assertions
	if assert.NoError(t, CheckReady(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/playbypost-dnd/v1/ready")
	config.Values.IsReady.Store(true)
	config.Values.IsReady.Store(false)
	if assert.NoError(t, CheckReady(c1)) {
		assert.Equal(t, http.StatusServiceUnavailable, rec1.Code)
	}
}
