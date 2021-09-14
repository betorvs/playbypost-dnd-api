package campaigns

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/campaign"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllCampaigns(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	value := make(url.Values)
	_, err := GetAllCampaigns(value)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetCampaign)
}

func TestGetOneCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := GetOneCampaign(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsGetOneCampaign)
}

func TestCheckCampaignExist(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	res := CheckCampaignExist(exampleID)
	assert.False(t, res)
	// expected := 1
	// assert.Equal(t, expected, test.MongoDBCallsGetOneCampaign)
}

func TestCheckListPlayers(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, res := CheckListPlayers(exampleID, "playerOne")
	assert.False(t, res)
}

func TestCreateCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(campaign.Campaign)
	example.ChannelID = "nochannel"
	example.StartLevel = 1
	_, err := CreateCampaign(example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsSaveCampaign)
}

func TestUpdateCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	example := new(campaign.Campaign)
	example.ChannelID = "nochannel"
	example.StartLevel = 1
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := UpdateCampaign(exampleID, example)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsUpdateCampaign)
}

func TestAddPlayer(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := AddPlayer(exampleID, "playerID", "playerOne")
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsAddPlayerToCampaign)

}

func TestDeleteCampaign(t *testing.T) {
	appcontext.Current.Add(appcontext.MongoRepository, test.InitMongoMock)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f7")
	_, err := DeleteCampaign(exampleID)
	assert.NoError(t, err)
	expected := 1
	assert.Equal(t, expected, test.MongoDBCallsDeleteCampaignByID)
}
