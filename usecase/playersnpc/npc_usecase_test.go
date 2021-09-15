package playersnpc

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllNPCS(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	value := make(url.Values)
	_, err := GetAllNPCS(value)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetNPCS)
}

func TestCreateNPC(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	example := new(player.NPC)
	_, err := CreateNPC(example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsSaveNPC)
}

func TestDeleteNPC(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := DeleteNPC(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsDeleteNPCByID)
}

func TestDamageNPC(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := DamageNPC(exampleID, "remove", 10)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsSetDamageNPCByID)
	_, err1 := DamageNPC(exampleID, "add", 10)
	assert.NoError(t, err1)
}

func TestChangeNPCCondition(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	player := new(player.NPCCondition)
	_, err := ChangeNPCCondition(exampleID, player)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsSetContitionNPCByID)
}
