package adventures

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	value := make(url.Values)
	_, err := GetAllAdventure(value)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetAdventure)
}

func TestGetOneAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := GetOneAdventure(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetOneAdventure)
}

func TestCreateAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(adventure.Adventure)
	_, err := CreateAdventure(example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddAdventure)
}

func TestUpdateAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(adventure.Adventure)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := UpdateAdventure(exampleID, example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsUpdateAdventure)
}

func TestDeleteAdventure(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := DeleteAdventure(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsDeleteAdventureByID)
}

func TestCheckAdventureExist(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	res := CheckAdventureExist(exampleID)
	assert.False(t, res)
}

func TestCheckEncounterWasAdded(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, res := CheckEncounterWasAdded(exampleID, "testEncounter")
	assert.False(t, res)
}

func TestAddEncounter(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := AddEncounter(exampleID, "testEncounter")
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddEncounterToAdventure)
}

func TestChangeAdventureStatus(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	player := new(adventure.AddPlayersID)
	_, err := ChangeAdventureStatus(exampleID, "living", player)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsChangeAdventureStatusByID)
}
