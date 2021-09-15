package encounters

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/domain/encounter"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	value := make(url.Values)
	_, err := GetEncounter(value)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetEncounter)
}

func TestGetOneEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := GetOneEncounter(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetOneEncounter)
}

func TestCreateEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(encounter.Encounter)
	_, err := CreateEncounter(example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddEncounter)
}

func TestUpdateEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(encounter.Encounter)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := UpdateEncounter(exampleID, example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsUpdateEncounter)
}

func TestDeleteEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := DeleteEncounter(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsDeleteEncounterByID)
}

func TestCheckNPCExist(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, res := CheckNPCExist(exampleID, "NPC")
	assert.False(t, res)
}

func TestAddNPC(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := AddNPC(exampleID, "testEncounter")
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddNPCTOEncounter)
}

func TestCheckEncounterExist(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	res := CheckEncounterExist(exampleID)
	assert.False(t, res)
}

func TestChangeEncounterStatus(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	player := new(adventure.AddPlayersID)
	_, err := ChangeEncounterStatus(exampleID, "living", player)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsChangeEncounterStatusByID)
}
