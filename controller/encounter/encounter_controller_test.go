package encounter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/domain/encounter"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, GetAllEncounter(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetOneEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/encounter/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, GetOneEncounter(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/encounter/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	if assert.NoError(t, GetOneEncounter(c1)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec1.Code)
	}

	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/encounter/:id")
	c2.SetParamNames("id")
	c2.SetParamValues("BLA")

	if assert.NoError(t, GetOneEncounter(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}
}

func TestPostEncounter(t *testing.T) {
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
	c.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PostEncounter(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := encounter.Encounter{}
	example1.AdventureID = "BLA"
	example1.Status = "onhold"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PostEncounter(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := encounter.Encounter{}
	example2.AdventureID = "5e70e4c5d2f3f777c16b29f7"
	example2.Status = "onhold"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PostEncounter(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := encounter.Encounter{}
	example3.AdventureID = "5e70e4c5d2f3f777c16b29f6"
	example3.Status = "onhold"
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PostEncounter(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}

	example4 := encounter.Encounter{}
	example4.AdventureID = "5e70e4c5d2f3f777c16b29f6"
	example4.Status = "BLA"
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPost, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PostEncounter(c4)) {
		assert.Equal(t, http.StatusBadRequest, rec4.Code)
	}
}

func TestPutEncounter(t *testing.T) {
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
	c.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PutEncounter(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := encounter.Encounter{}
	example1.AdventureID = "BLA"
	example1.Status = "BLA"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPut, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PutEncounter(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := encounter.Encounter{}
	example2.AdventureID = "BLA"
	example2.Status = "onhold"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPut, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PutEncounter(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := encounter.Encounter{}
	example3.AdventureID = "5e70e4c5d2f3f777c16b29f8"
	example3.Status = "onhold"
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPut, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PutEncounter(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	example4 := encounter.Encounter{}
	example4.AdventureID = "5e70e4c5d2f3f777c16b29f6"
	example4.Status = "onhold"
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPut, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/encounter")

	// Assertions
	if assert.NoError(t, PutEncounter(c4)) {
		assert.Equal(t, http.StatusOK, rec4.Code)
	}
}

func TestDeleteEncounter(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/encounter/:id")
	c.SetParamNames("id")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7")

	// Assertions
	if assert.NoError(t, DeleteEncounter(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req1 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/encounter/:id")
	c1.SetParamNames("id")
	c1.SetParamValues("BLA")

	// Assertions
	if assert.NoError(t, DeleteEncounter(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}
	//
	req2 := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/encounter/:id")
	c2.SetParamNames("id")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f8")

	// Assertions
	if assert.NoError(t, DeleteEncounter(c2)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec2.Code)
	}
}

func TestAddNPC(t *testing.T) {
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
	c.SetPath("/encounter/npc")

	// Assertions
	if assert.NoError(t, AddNPC(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
	example1 := encounter.AddNPC{}
	example1.EncounterID = "BLA"
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPost, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/encounter/npc")

	// Assertions
	if assert.NoError(t, AddNPC(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := encounter.AddNPC{}
	example2.EncounterID = "5e70e4c5d2f3f777c16b29f8"
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPost, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/encounter/npc")

	// Assertions
	if assert.NoError(t, AddNPC(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	example3 := encounter.AddNPC{}
	example3.EncounterID = "5e70e4c5d2f3f777c16b29f6"
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPost, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/encounter/npc")

	// Assertions
	if assert.NoError(t, AddNPC(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestChangeEncounterStatus(t *testing.T) {
	appcontext.Current.Add(appcontext.Logger, test.InitMockLogger)
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	e := echo.New()
	example := adventure.AddPlayersID{}
	requestByte, _ := json.Marshal(example)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPut, "/", requestReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/encounter/:id/:status")
	c.SetParamNames("id", "status")
	c.SetParamValues("5e70e4c5d2f3f777c16b29f7", "living")

	// Assertions
	if assert.NoError(t, ChangeEncounterStatus(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	example1 := adventure.AddPlayersID{}
	requestByte1, _ := json.Marshal(example1)
	requestReader1 := bytes.NewReader(requestByte1)
	req1 := httptest.NewRequest(http.MethodPut, "/", requestReader1)
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)
	c1.SetPath("/encounter/:id/:status")
	c1.SetParamNames("id", "status")
	c1.SetParamValues("5e70e4c5d2f3f777c16b29f6", "BLA")

	// Assertions
	if assert.NoError(t, ChangeEncounterStatus(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	example2 := adventure.AddPlayersID{}
	requestByte2, _ := json.Marshal(example2)
	requestReader2 := bytes.NewReader(requestByte2)
	req2 := httptest.NewRequest(http.MethodPut, "/", requestReader2)
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	c2.SetPath("/encounter/:id/:status")
	c2.SetParamNames("id", "status")
	c2.SetParamValues("5e70e4c5d2f3f777c16b29f6", "living")

	// Assertions
	if assert.NoError(t, ChangeEncounterStatus(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}

	example3 := adventure.AddPlayersID{}
	requestByte3, _ := json.Marshal(example3)
	requestReader3 := bytes.NewReader(requestByte3)
	req3 := httptest.NewRequest(http.MethodPut, "/", requestReader3)
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	c3.SetPath("/encounter/:id/:status")
	c3.SetParamNames("id", "status")
	c3.SetParamValues("BLA", "onhold")

	// Assertions
	if assert.NoError(t, ChangeEncounterStatus(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	example4 := "{\"key\":\"value\"}"
	requestByte4, _ := json.Marshal(example4)
	requestReader4 := bytes.NewReader(requestByte4)
	req4 := httptest.NewRequest(http.MethodPut, "/", requestReader4)
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	c4.SetPath("/encounter/:id/:status")
	c4.SetParamNames("id", "status")
	c4.SetParamValues("5e70e4c5d2f3f777c16b29f6", "onhold")

	// Assertions
	if assert.NoError(t, ChangeEncounterStatus(c4)) {
		assert.Equal(t, http.StatusBadRequest, rec4.Code)
	}
}
