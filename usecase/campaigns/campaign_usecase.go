package campaigns

import (
	"fmt"
	"net/url"

	"github.com/betorvs/playbypost-dnd/domain/campaign"
	"github.com/betorvs/playbypost-dnd/domain/mongodb"
	"github.com/betorvs/playbypost-dnd/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetAllCampaigns func
func GetAllCampaigns(queryParameters url.Values) ([]*campaign.Campaign, error) {
	repo := mongodb.GetMongoRepository()
	return repo.GetCampaign(queryParameters)
}

//GetOneCampaign func
func GetOneCampaign(campaignID primitive.ObjectID) (*campaign.Campaign, error) {
	repo := mongodb.GetMongoRepository()
	campaigns, err := repo.GetOneCampaign(campaignID)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

//CheckCampaignExist func
func CheckCampaignExist(campaignID primitive.ObjectID) bool {
	repo := mongodb.GetMongoRepository()
	campaign, err := repo.GetOneCampaign(campaignID)
	if err != nil {
		return false
	}
	if campaignID == campaign.ID {
		return true
	}
	return false
}

//CheckListPlayers func
func CheckListPlayers(campaignID primitive.ObjectID, playerID string) (string, bool) {
	repo := mongodb.GetMongoRepository()
	campaigns, err := repo.GetOneCampaign(campaignID)
	if err != nil {
		return "cannot get campaign from database", false
	}
	if utils.StringInSlice(playerID, campaigns.PlayersID) {
		return "allowed", true
	}
	return "blocked", false
}

//CreateCampaign func
func CreateCampaign(campaigns *campaign.Campaign) (string, error) {
	if campaigns.StartLevel == 0 {
		campaigns.StartLevel = 1
	}
	param := make(url.Values)
	param["channel_id"] = []string{campaigns.ChannelID}
	repo := mongodb.GetMongoRepository()
	checkCampaign, err := repo.GetCampaign(param)
	if err != nil {
		return "", err
	}
	for _, v := range checkCampaign {
		if v.ChannelID == campaigns.ChannelID {
			err = fmt.Errorf("channel in use")
			return fmt.Sprintf("Channel was registered with Campaign: %s", v.Title), err
		}
	}
	if len(campaigns.PlayersID) == 0 {
		campaigns.PlayersID = []string{}
	}
	if len(campaigns.PlayersUsername) == 0 {
		campaigns.PlayersUsername = []string{}
	}

	id, err := repo.SaveCampaign(campaigns)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}

//UpdateCampaign func
func UpdateCampaign(campaignID primitive.ObjectID, campaigns *campaign.Campaign) (string, error) {
	if campaigns.StartLevel == 0 {
		campaigns.StartLevel = 1
	}
	repo := mongodb.GetMongoRepository()
	id, err := repo.UpdateCampaign(campaignID, campaigns)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}

//AddPlayer func
func AddPlayer(campaignID primitive.ObjectID, playerID, playerUsername string) (int64, error) {
	_, checkPlayerID := CheckListPlayers(campaignID, playerID)
	if checkPlayerID {
		return -1, nil
	}
	repo := mongodb.GetMongoRepository()
	result, err := repo.AddPlayerToCampaign(campaignID, playerID, playerUsername)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// DeleteCampaign func
func DeleteCampaign(campaignID primitive.ObjectID) (int64, error) {
	repo := mongodb.GetMongoRepository()
	result, err := repo.DeleteCampaignByID(campaignID)
	if err != nil {
		return 0, err
	}

	return result, nil
}
