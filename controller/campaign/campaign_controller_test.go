package campaign

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/campaign"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/campaign")

	// Assertions
	if assert.NoError(t, GetAllCampaign(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetOneCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/adventure/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, GetOneCampaign(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/campaign/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	if assert.NoError(t, GetOneCampaign(c1)) {
		assert.Equal(t, http.StatusBadGateway, rec1.Code)
	}
}

func TestPostCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := campaign.Campaign{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/campaign")

	// Assertions
	if assert.NoError(t, PostCampaign(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := campaign.Campaign{}
	// example1.CampaignID = "BLA"
	example1.Status = "onhold"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/campaign")

	// Assertions
	if assert.NoError(t, PostCampaign(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}
}

func TestPutCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := campaign.Campaign{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPut, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/campaign")

	// Assertions
	if assert.NoError(t, PutCampaign(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := campaign.Campaign{}
	// example1.CampaignID = "BLA"
	example1.Status = "onhold"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPut, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/campaign")

	// Assertions
	if assert.NoError(t, PutCampaign(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}

}

func TestCheckPlayerAllowed(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/campaign/:id/:playerid")
	c.SetParamNames("id", "playerid")
	c.SetParamValues("BLA", "playerID")

	// Assertions
	if assert.NoError(t, CheckPlayerAllowed(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/campaign/:id/:playerid")
	c1.SetParamNames("id", "playerid")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f6", "playerID")

	// Assertions
	if assert.NoError(t, CheckPlayerAllowed(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}

	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/campaign/:id/:playerid")
	c2.SetParamNames("id", "playerid")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f8", "playerID")

	// Assertions
	if assert.NoError(t, CheckPlayerAllowed(c2)) {
		assert.Equal(t, http.StatusForbidden, rec2.Code)
	}
}

func TestAddPlayerCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := campaign.AddPlayer{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/campaign/:id/player")
	c.SetParamNames("id")
	c.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, AddPlayerCampaign(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := campaign.AddPlayer{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/campaign/:id/player")
	c1.SetParamNames("id")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f6")

	// Assertions
	if assert.NoError(t, AddPlayerCampaign(c1)) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	}

	example2 := campaign.AddPlayer{}
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/campaign/:id/player")
	c2.SetParamNames("id")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, AddPlayerCampaign(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}
}

func TestDeleteCampaign(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/campaign/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, DeleteCampaign(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/campaign/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, DeleteCampaign(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}
}
