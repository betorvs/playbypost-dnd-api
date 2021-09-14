package adventure

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, GetAllAdventure(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetOneAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/adventure/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, GetOneAdventure(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/adventure/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	if assert.NoError(t, GetOneAdventure(c1)) {
		assert.Equal(t, http.StatusBadGateway, rec1.Code)
	}
}

func TestPostAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := adventure.Adventure{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, PostAdventure(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := adventure.Adventure{}
	example1.CampaignID = "BLA"
	example1.Status = "onhold"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, PostAdventure(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := adventure.Adventure{}
	example2.CampaignID = "5e70e4c5d2f3f777c16b29f7"
	example2.Status = "onhold"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, PostAdventure(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := adventure.Adventure{}
	example3.CampaignID = "5e70e4c5d2f3f777c16b29f6"
	example3.Status = "onhold"
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, PostAdventure(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestPutAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := adventure.Adventure{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPut, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, PutAdventure(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := adventure.Adventure{}
	example1.CampaignID = "BLA"
	example1.Status = "onhold"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPut, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, PutAdventure(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := adventure.Adventure{}
	example2.CampaignID = "5e70e4c5d2f3f777c16b29f7"
	example2.Status = "onhold"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPut, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, PutAdventure(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := adventure.Adventure{}
	example3.CampaignID = "5e70e4c5d2f3f777c16b29f6"
	example3.Status = "onhold"
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPut, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/adventure")

	// Assertions
	if assert.NoError(t, PutAdventure(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestDeleteAdventure(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/adventure/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, DeleteAdventure(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/adventure/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, DeleteAdventure(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}
	//
	req2 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/adventure/:id")
	c2.SetParamNames("id")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, DeleteAdventure(c2)) {
		assert.Equal(t, http.StatusBadGateway, rec2.Code)
	}
}

func TestAddEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := adventure.Adventure{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/adventure/:id/encounter")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, AddEncounter(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
	example1 := adventure.Adventure{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/adventure/:id/encounter")
	c1.SetParamNames("id")
	c1.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, AddEncounter(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := adventure.Adventure{}
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/adventure/:id/encounter")
	c2.SetParamNames("id")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, AddEncounter(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := adventure.Adventure{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/adventure/:id/encounter")
	c3.SetParamNames("id")
	c3.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, AddEncounter(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestChangeAdventureStatus(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := adventure.Adventure{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPut, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/adventure/:id/:status")
	c.SetParamNames("id", "status")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7", "living")

	// Assertions
	if assert.NoError(t, ChangeAdventureStatus(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := adventure.Adventure{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPut, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/adventure/:id/:status")
	c1.SetParamNames("id", "status")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f6", "BLA")

	// Assertions
	if assert.NoError(t, ChangeAdventureStatus(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := adventure.Adventure{}
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPut, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/adventure/:id/:status")
	c2.SetParamNames("id", "status")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6", "living")

	// Assertions
	if assert.NoError(t, ChangeAdventureStatus(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}
}
