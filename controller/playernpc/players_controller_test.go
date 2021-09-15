package playernpc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPlayers(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player")

	// Assertions
	if assert.NoError(t, GetAllPlayers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetOnePlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, GetOnePlayer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	if assert.NoError(t, GetOnePlayer(c1)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec1.Code)
	}

	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:id")
	c2.SetParamNames("id")
	c2.SetParamValues("BLA")

	if assert.NoError(t, GetOnePlayer(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}
}

func TestPostPlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player")

	// Assertions
	if assert.NoError(t, PostPlayer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	// example1 := player.Players{}
	// example1.AdventureID = "BLA"
	// requestByte1, _ := json.Marshal(example1)
	// requestReader1 := bytes.NewReader(requestByte1)
	// req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	// req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec1 := httptest.NewRecorder()
	// c1 := e.NewContext(req1, rec1)
	// c1.SetPath("/player")

	// // Assertions
	// if assert.NoError(t, PostPlayer(c1)) {
	// 	assert.Equal(t, http.StatusBadRequest, rec1.Code)
	// }

	// example2 := player.Players{}
	// requestByte2, _ := json.Marshal(example2)
	// requestReader2 := bytes.NewReader(requestByte2)
	// req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	// req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec2 := httptest.NewRecorder()
	// c2 := e.NewContext(req2, rec2)
	// c2.SetPath("/player")

	// // Assertions
	// if assert.NoError(t, PostPlayer(c2)) {
	// 	assert.Equal(t, http.StatusBadRequest, rec2.Code)
	// }

	example3 := player.Players{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player")

	// Assertions
	if assert.NoError(t, PostPlayer(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}

	// example4 := player.Players{}
	// requestByte4, _ := json.Marshal(example4)
	// requestReader4 := bytes.NewReader(requestByte4)
	// req4 := httptest.NewRequest(http.MethodPost, "/", requestReader4)
	// req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec4 := httptest.NewRecorder()
	// c4 := e.NewContext(req4, rec4)
	// c4.SetPath("/player")

	// // Assertions
	// if assert.NoError(t, PostPlayer(c4)) {
	// 	assert.Equal(t, http.StatusBadRequest, rec4.Code)
	// }
}

func TestUpdateOnePlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := "{\"key\":\"value\"}"
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPut, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player")

	// Assertions
	if assert.NoError(t, UpdateOnePlayer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	// example1 := player.Players{}
	// requestByte1, _ := json.Marshal(example1)
	// requestReader1 := bytes.NewReader(requestByte1)
	// req1 := httptest.NewRequest(http.MethodPut, "/", requestReader1)
	// req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec1 := httptest.NewRecorder()
	// c1 := e.NewContext(req1, rec1)
	// c1.SetPath("/player")

	// // Assertions
	// if assert.NoError(t, UpdateOnePlayer(c1)) {
	// 	assert.Equal(t, http.StatusBadRequest, rec1.Code)
	// }

	// example2 := player.Players{}
	// requestByte2, _ := json.Marshal(example2)
	// requestReader2 := bytes.NewReader(requestByte2)
	// req2 := httptest.NewRequest(http.MethodPut, "/", requestReader2)
	// req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec2 := httptest.NewRecorder()
	// c2 := e.NewContext(req2, rec2)
	// c2.SetPath("/player")

	// // Assertions
	// if assert.NoError(t, UpdateOnePlayer(c2)) {
	// 	assert.Equal(t, http.StatusBadRequest, rec2.Code)
	// }

	example3 := player.Players{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPut, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player")

	// Assertions
	if assert.NoError(t, UpdateOnePlayer(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
	config.Values.RuleBuiltin = false
	example4 := player.Players{}
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPut, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/player")

	// Assertions
	if assert.NoError(t, UpdateOnePlayer(c4)) {
		assert.Equal(t, http.StatusOK, rec4.Code)
	}
}

func TestDeletePlayer(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, DeletePlayer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, DeletePlayer(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}
	//
	req2 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:id")
	c2.SetParamNames("id")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, DeletePlayer(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}
}

func TestAddCampaignToPlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := player.AddCampaign{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/campaign")
	c.SetParamNames("playerid")
	c.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, AddCampaignToPlayer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := player.AddCampaign{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/campaign")
	c1.SetParamNames("playerid")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, AddCampaignToPlayer(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := "{\"key\":\"value\"}"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/campaign")
	c2.SetParamNames("playerid")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, AddCampaignToPlayer(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := player.AddCampaign{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/campaign")
	c3.SetParamNames("playerid")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, AddCampaignToPlayer(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestAddOrRemoveHP(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/hp/:action/:value")
	c.SetParamNames("playerid", "action", "value")
	c.SetParamValues("BLA", "add", "10")

	// Assertions
	if assert.NoError(t, AddOrRemoveHP(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodPut, "/", nil)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/hp/:action/:value")
	c1.SetParamNames("playerid", "action", "value")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8", "add", "10")

	// Assertions
	if assert.NoError(t, AddOrRemoveHP(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	req2 := httptest.NewRequest(http.MethodPut, "/", nil)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/hp/:action/:value")
	c2.SetParamNames("playerid", "action", "value")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6", "delete", "10")

	// Assertions
	if assert.NoError(t, AddOrRemoveHP(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	req3 := httptest.NewRequest(http.MethodPut, "/", nil)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/hp/:action/:value")
	c3.SetParamNames("playerid", "action", "value")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6", "add", "third")

	// Assertions
	if assert.NoError(t, AddOrRemoveHP(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	req4 := httptest.NewRequest(http.MethodPut, "/", nil)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/player/:playerid/hp/:action/:value")
	c4.SetParamNames("playerid", "action", "value")
	c4.SetParamValues("5e70e4c5d2f3f777c16b29f6", "add", "10")

	// Assertions
	if assert.NoError(t, AddOrRemoveHP(c4)) {
		assert.Equal(t, http.StatusOK, rec4.Code)
	}
}

func TestAddPlayerXP(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/xp/:value")
	c.SetParamNames("playerid", "value")
	c.SetParamValues("BLA", "10")

	// Assertions
	if assert.NoError(t, AddPlayerXP(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodPut, "/", nil)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/xp/:value")
	c1.SetParamNames("playerid", "value")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8", "10")

	// Assertions
	if assert.NoError(t, AddPlayerXP(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	req2 := httptest.NewRequest(http.MethodPut, "/", nil)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/xp/:value")
	c2.SetParamNames("playerid", "value")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6", "third")

	// Assertions
	if assert.NoError(t, AddPlayerXP(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	req3 := httptest.NewRequest(http.MethodPut, "/", nil)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/xp/:value")
	c3.SetParamNames("playerid", "value")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6", "100")

	// Assertions
	if assert.NoError(t, AddPlayerXP(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}

	// req4 := httptest.NewRequest(http.MethodPut, "/", nil)
	// req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec4 := httptest.NewRecorder()
	// c4 := e.NewContext(req4, rec4)
	// c4.SetPath("/player/:playerid/xp/:value")
	// c4.SetParamNames("playerid", "value")
	// c4.SetParamValues("5e70e4c5d2f3f777c16b29f6", "10")

	// // Assertions
	// if assert.NoError(t, AddPlayerXP(c4)) {
	// 	assert.Equal(t, http.StatusOK, rec4.Code)
	// }
}

func TestUseSpellByLevel(t *testing.T) {
	//
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/spell/:level/:value")
	c.SetParamNames("playerid", "level", "value")
	c.SetParamValues("BLA", "1", "1")

	// Assertions
	if assert.NoError(t, UseSpellByLevel(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodPut, "/", nil)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/spell/:level/:value")
	c1.SetParamNames("playerid", "level", "value")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8", "1", "1")

	// Assertions
	if assert.NoError(t, UseSpellByLevel(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	req2 := httptest.NewRequest(http.MethodPut, "/", nil)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/spell/:level/:value")
	c2.SetParamNames("playerid", "level", "value")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6", "10", "1")

	// Assertions
	if assert.NoError(t, UseSpellByLevel(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	req3 := httptest.NewRequest(http.MethodPut, "/", nil)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/spell/:level/:value")
	c3.SetParamNames("playerid", "level", "value")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6", "1", "one")

	// Assertions
	if assert.NoError(t, UseSpellByLevel(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	req4 := httptest.NewRequest(http.MethodPut, "/", nil)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/player/:playerid/spell/:level/:value")
	c4.SetParamNames("playerid", "level", "value")
	c4.SetParamValues("5e70e4c5d2f3f777c16b29f6", "1", "1")

	// Assertions
	if assert.NoError(t, UseSpellByLevel(c4)) {
		assert.Equal(t, http.StatusOK, rec4.Code)
	}
}

func TestFullRestPlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/fullrest")
	c.SetParamNames("playerid")
	c.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, FullRestPlayer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodPut, "/", nil)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/fullrest")
	c1.SetParamNames("playerid")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, FullRestPlayer(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	req2 := httptest.NewRequest(http.MethodPut, "/", nil)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/fullrest")
	c2.SetParamNames("playerid")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, FullRestPlayer(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}

	// req3 := httptest.NewRequest(http.MethodPut, "/", nil)
	// req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec3 := httptest.NewRecorder()
	// c3 := e.NewContext(req3, rec3)
	// c3.SetPath("/player/:playerid/fullrest")
	// c3.SetParamNames("playerid", "value")
	// c3.SetParamValues("5e70e4c5d2f3f777c16b29f6", "100")

	// // Assertions
	// if assert.NoError(t, FullRestPlayer(c3)) {
	// 	assert.Equal(t, http.StatusOK, rec3.Code)
	// }
}

func TestChangeCondition(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	// player.Condition
	example := player.Condition{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/condition")
	c.SetParamNames("playerid")
	c.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, ChangeCondition(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := player.Condition{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/condition")
	c1.SetParamNames("playerid")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, ChangeCondition(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := "{\"key\":\"value\"}"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/condition")
	c2.SetParamNames("playerid")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, ChangeCondition(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := player.Condition{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/condition")
	c3.SetParamNames("playerid")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, ChangeCondition(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestAddTreasure(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := player.Treasure{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/treasure")
	c.SetParamNames("playerid")
	c.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, AddTreasure(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := player.Treasure{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/treasure")
	c1.SetParamNames("playerid")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, AddTreasure(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := "{\"key\":\"value\"}"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/treasure")
	c2.SetParamNames("playerid")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, AddTreasure(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := player.Treasure{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/treasure")
	c3.SetParamNames("playerid")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f9")

	// Assertions
	if assert.NoError(t, AddTreasure(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	example4 := player.Treasure{}
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPost, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/player/:playerid/treasure")
	c4.SetParamNames("playerid")
	c4.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, AddTreasure(c4)) {
		assert.Equal(t, http.StatusOK, rec4.Code)
	}
}

func TestAddOrRemoveOtherItems(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := rule.SimpleList{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/items/:action")
	c.SetParamNames("playerid", "action")
	c.SetParamValues("BLA", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveOtherItems(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.SimpleList{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/items/:action")
	c1.SetParamNames("playerid", "action")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveOtherItems(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := "{\"key\":\"value\"}"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/items/:action")
	c2.SetParamNames("playerid", "action")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveOtherItems(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := rule.SimpleList{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/items/:action")
	c3.SetParamNames("playerid", "action")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6", "delete")

	// Assertions
	if assert.NoError(t, AddOrRemoveOtherItems(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	example4 := rule.SimpleList{}
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPost, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/player/:playerid/items/:action")
	c4.SetParamNames("playerid", "action")
	c4.SetParamValues("5e70e4c5d2f3f777c16b29f9", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveOtherItems(c4)) {
		assert.Equal(t, http.StatusBadRequest, rec4.Code)
	}

	example5 := rule.SimpleList{}
	requestByte5, _ := json.Marshal(example5)
	requestReader5 := bytes.NewReader(requestByte5)
	req5 := httptest.NewRequest(http.MethodPost, "/", requestReader5)
	req5.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec5 := httptest.NewRecorder()
	c5 := e.NewContext(req5, rec5)
	c5.SetPath("/player/:playerid/items/:action")
	c5.SetParamNames("playerid", "action")
	c5.SetParamValues("5e70e4c5d2f3f777c16b29f6", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveOtherItems(c5)) {
		assert.Equal(t, http.StatusOK, rec5.Code)
	}
}

func TestAddArmorWeaponPlayerByID(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := player.Armory{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/armory")
	c.SetParamNames("playerid")
	c.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, AddArmorWeaponPlayerByID(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := player.Armory{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/armory")
	c1.SetParamNames("playerid")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, AddArmorWeaponPlayerByID(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := "{\"key\":\"value\"}"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/armory")
	c2.SetParamNames("playerid")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, AddArmorWeaponPlayerByID(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := player.Armory{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/armory")
	c3.SetParamNames("playerid")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f9")

	// Assertions
	if assert.NoError(t, AddArmorWeaponPlayerByID(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	example4 := player.Armory{}
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPost, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/player/:playerid/armory")
	c4.SetParamNames("playerid")
	c4.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, AddArmorWeaponPlayerByID(c4)) {
		assert.Equal(t, http.StatusOK, rec4.Code)
	}
}

func TestAddOrRemoveMagicItems(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := rule.SimpleList{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/player/:playerid/magicitems/:action")
	c.SetParamNames("playerid", "action")
	c.SetParamValues("BLA", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveMagicItems(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := rule.SimpleList{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/player/:playerid/magicitems/:action")
	c1.SetParamNames("playerid", "action")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveMagicItems(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := "{\"key\":\"value\"}"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/player/:playerid/magicitems/:action")
	c2.SetParamNames("playerid", "action")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveMagicItems(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := rule.SimpleList{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/player/:playerid/magicitems/:action")
	c3.SetParamNames("playerid", "action")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6", "delete")

	// Assertions
	if assert.NoError(t, AddOrRemoveMagicItems(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	example4 := rule.SimpleList{}
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPost, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/player/:playerid/magicitems/:action")
	c4.SetParamNames("playerid", "action")
	c4.SetParamValues("5e70e4c5d2f3f777c16b29f9", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveMagicItems(c4)) {
		assert.Equal(t, http.StatusBadRequest, rec4.Code)
	}

	example5 := rule.SimpleList{}
	requestByte5, _ := json.Marshal(example5)
	requestReader5 := bytes.NewReader(requestByte5)
	req5 := httptest.NewRequest(http.MethodPost, "/", requestReader5)
	req5.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec5 := httptest.NewRecorder()
	c5 := e.NewContext(req5, rec5)
	c5.SetPath("/player/:playerid/magicitems/:action")
	c5.SetParamNames("playerid", "action")
	c5.SetParamValues("5e70e4c5d2f3f777c16b29f6", "add")

	// Assertions
	if assert.NoError(t, AddOrRemoveMagicItems(c5)) {
		assert.Equal(t, http.StatusOK, rec5.Code)
	}
}
