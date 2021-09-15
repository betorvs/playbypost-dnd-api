package playernpc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllNPC(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/npc")

	// Assertions
	if assert.NoError(t, GetAllNPC(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestPostNPC(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/npc")

	// Assertions
	if assert.NoError(t, PostNPC(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example3 := player.NPC{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/npc")

	// Assertions
	if assert.NoError(t, PostNPC(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestDeleteNPC(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/npc/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, DeleteNPC(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/npc/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, DeleteNPC(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}
	//
	req2 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/npc/:id")
	c2.SetParamNames("id")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, DeleteNPC(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}
}

func TestPostDamageNPC(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/npc/:id/hp/:action/:value")
	c.SetParamNames("id", "action", "value")
	c.SetParamValues("BLA", "add", "10")

	// Assertions
	if assert.NoError(t, PostDamageNPC(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodPut, "/", nil)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/npc/:id/hp/:action/:value")
	c1.SetParamNames("id", "action", "value")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8", "add", "10")

	// Assertions
	if assert.NoError(t, PostDamageNPC(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	req2 := httptest.NewRequest(http.MethodPut, "/", nil)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/npc/:id/hp/:action/:value")
	c2.SetParamNames("id", "action", "value")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6", "delete", "10")

	// Assertions
	if assert.NoError(t, PostDamageNPC(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	req3 := httptest.NewRequest(http.MethodPut, "/", nil)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/npc/:id/hp/:action/:value")
	c3.SetParamNames("id", "action", "value")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6", "add", "third")

	// Assertions
	if assert.NoError(t, PostDamageNPC(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	req4 := httptest.NewRequest(http.MethodPut, "/", nil)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/npc/:id/hp/:action/:value")
	c4.SetParamNames("id", "action", "value")
	c4.SetParamValues("5e70e4c5d2f3f777c16b29f6", "add", "10")

	// Assertions
	if assert.NoError(t, PostDamageNPC(c4)) {
		assert.Equal(t, http.StatusOK, rec4.Code)
	}
}

func TestChangeNPCCondition(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	e := echo.New()
	// player.Condition
	example := player.Condition{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/npc/:id/condition")
	c.SetParamNames("id")
	c.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, ChangeNPCCondition(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := player.Condition{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/npc/:id/condition")
	c1.SetParamNames("id")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, ChangeNPCCondition(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := "{\"key\":\"value\"}"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/npc/:id/condition")
	c2.SetParamNames("id")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, ChangeNPCCondition(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := player.Condition{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/npc/:id/condition")
	c3.SetParamNames("id")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, ChangeNPCCondition(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}
