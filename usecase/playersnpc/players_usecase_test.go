package playersnpc

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllPlayers(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	value := make(url.Values)
	_, err := GetAllPlayers(value)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetPlayers)
}

func TestGetOnePlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := GetOnePlayer(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetOnePlayer)
}

func TestCreatePlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(player.Players)
	_, err := CreatePlayer(example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsSavePlayer)
}

func TestUpdatePlayerRecalc(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(player.Players)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := UpdatePlayerRecalc(exampleID, example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsUpdatePlayer)
}

func TestUpdatePlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(player.Players)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := UpdatePlayer(exampleID, example)
	assert.NoError(t, err)
	expected := 2
	assert.Equal(t, expected, test.MongoDBCallsUpdatePlayer)

}

func TestCheckPlayerExist(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	res := CheckPlayerExist(exampleID)
	assert.False(t, res)
}

func TestDeletePlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := DeletePlayer(exampleID, exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsDeletePlayerByID)
}

func TestAddCampaignToPlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	player := new(player.AddCampaign)
	_, err := AddCampaignToPlayer(exampleID, player)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddCampaignToPlayer)
}

func TestAddOrRemoveHP(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := AddOrRemoveHP(exampleID, "remove", 10)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddOrRemovePlayerHP)
	_, err1 := AddOrRemoveHP(exampleID, "add", 10)
	assert.NoError(t, err1)
}

func TestAddPlayerXP(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := AddPlayerXP(exampleID, 10)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddPlayerXP)
}

func TestUseSpellByLevel(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := UseSpellByLevel(exampleID, 1, 10)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsUsageSpellByLevel)
}

func TestFullRestPlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := FullRestPlayer(exampleID)
	assert.NoError(t, err)
}

func TestChangeCondition(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	player := new(player.Condition)
	_, err := ChangeCondition(exampleID, player)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsChangePlayerCondition)
}

func TestInventoryIDByPlayerID(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	// player := new(player.Condition)
	_, err := InventoryIDByPlayerID(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetInventoryID)
}

func TestGetInventoryByPlayerID(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	// player := new(player.Condition)
	_, err := GetInventoryByPlayerID(exampleID)
	assert.NoError(t, err)
}

func TestAddTreasure(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	player := new(player.Treasure)
	_, err := AddTreasure(exampleID, player)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddTreasurePlayer)
}

func TestAddOrRemoveOtherItems(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	// player := new(player.Treasure)
	_, err := AddOrRemoveOtherItems(exampleID, "add", []string{"horse"})
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddOthersItems)
	_, err1 := AddOrRemoveOtherItems(exampleID, "remove", []string{"horse"})
	assert.NoError(t, err1)
	assert.Equal(t, expected, test.MongoDBCallsRemoveOthersItems)
}

func TestAddArmorWeaponPlayerByID(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	player := new(player.Armory)
	_, err := AddArmorWeaponPlayerByID(exampleID, exampleID, player)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsSetArmorWeaponPlayerByID)
}

func TestAddOrRemoveMagicItems(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	// player := new(player.Treasure)
	_, err := AddOrRemoveMagicItems(exampleID, "add", []string{"ring"})
	assert.NoError(t, err)
	// expected := 1
	// assert.Equal(t, expected, test.MongoDBCallsAddOthersItems)
	_, err1 := AddOrRemoveMagicItems(exampleID, "remove", []string{"ring"})
	assert.NoError(t, err1)
	// assert.Equal(t, expected, test.MongoDBCallsRemoveOthersItems)
	_, err2 := AddOrRemoveMagicItems(exampleID, "attune", []string{"ring"})
	assert.NoError(t, err2)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsSetMagicalEffect)
}

func TestAddMagicalEffects(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	// exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	player := new(player.Players)
	players := rule.MagicItem{}
	players.Feature = nil
	_, err := addMagicalEffects(player, players)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsSetArmorWeaponPlayerByID)
}

func TestNewInventory(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	// exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	// player := new(player.Condition)
	err := newInventory("exampleID")
	assert.NoError(t, err)
	expected := 2
	assert.GreaterOrEqual(t, expected, test.MongoDBCallsCreateInventory)
}

func TestNewPlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	player := new(player.Players)
	_, err := newPlayer(player)
	assert.NoError(t, err)
}
