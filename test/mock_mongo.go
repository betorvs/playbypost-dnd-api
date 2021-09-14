package test

import (
	"fmt"
	"net/url"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/domain/campaign"
	"github.com/betorvs/playbypost-dnd/domain/encounter"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// MongoDBCalls int
	MongoDBCalls int
	// MongoDBCallsSaveCampaign int
	MongoDBCallsSaveCampaign int
	// MongoDBCallsUpdateCampaign int
	MongoDBCallsUpdateCampaign int
	// MongoDBCallsAddPlayerToCampaign int
	MongoDBCallsAddPlayerToCampaign int
	// MongoDBCallsGetCampaign int
	MongoDBCallsGetCampaign int
	// MongoDBCallsGetOneCampaign int
	MongoDBCallsGetOneCampaign int
	// MongoDBCallsDeleteCampaignByID int
	MongoDBCallsDeleteCampaignByID int
	// MongoDBCallsAddAdventure int
	MongoDBCallsAddAdventure int
	// MongoDBCallsUpdateAdventure int
	MongoDBCallsUpdateAdventure int
	// MongoDBCallsGetAdventure int
	MongoDBCallsGetAdventure int
	// MongoDBCallsGetOneAdventure int
	MongoDBCallsGetOneAdventure int
	// MongoDBCallsAddEncounterToAdventure int
	MongoDBCallsAddEncounterToAdventure int
	// MongoDBCallsChangeAdventureStatusByID int
	MongoDBCallsChangeAdventureStatusByID int
	// MongoDBCallsDeleteAdventureByID int
	MongoDBCallsDeleteAdventureByID int
	// MongoDBCallsAddEncounter int
	MongoDBCallsAddEncounter int
	// MongoDBCallsUpdateEncounter int
	MongoDBCallsUpdateEncounter int
	// MongoDBCallsGetEncounter int
	MongoDBCallsGetEncounter int
	// MongoDBCallsGetOneEncounter int
	MongoDBCallsGetOneEncounter int
	// MongoDBCallsAddNPCTOEncounter int
	MongoDBCallsAddNPCTOEncounter int
	// MongoDBCallsChangeEncounterStatusByID int
	MongoDBCallsChangeEncounterStatusByID int
	// MongoDBCallsDeleteEncounterByID int
	MongoDBCallsDeleteEncounterByID int
	// MongoDBCallsGetPlayers int
	MongoDBCallsGetPlayers int
	// MongoDBCallsGetOnePlayer int
	MongoDBCallsGetOnePlayer int
	// MongoDBCallsSavePlayer int
	MongoDBCallsSavePlayer int
	// MongoDBCallsUpdatePlayer int
	MongoDBCallsUpdatePlayer int
	// MongoDBCallsAddCampaignToPlayer int
	MongoDBCallsAddCampaignToPlayer int
	// MongoDBCallsChangePlayerCondition int
	MongoDBCallsChangePlayerCondition int
	// MongoDBCallsAddPlayerXP int
	MongoDBCallsAddPlayerXP int
	// MongoDBCallsAddOrRemovePlayerHP int
	MongoDBCallsAddOrRemovePlayerHP int
	// MongoDBCallsCreateInventory int
	MongoDBCallsCreateInventory int
	// MongoDBCallsDeleteInventoryByID int
	MongoDBCallsDeleteInventoryByID int
	// MongoDBCallsGetInventoryByID int
	MongoDBCallsGetInventoryByID int
	// MongoDBCallsGetInventoryID int
	MongoDBCallsGetInventoryID int
	// MongoDBCallsAddTreasurePlayer int
	MongoDBCallsAddTreasurePlayer int
	// MongoDBCallsAddOthersItems int
	MongoDBCallsAddOthersItems int
	// MongoDBCallsRemoveOthersItems int
	MongoDBCallsRemoveOthersItems int
	// MongoDBCallsGetArmory int
	MongoDBCallsGetArmory int
	// MongoDBCallsUsageSpellByLevel int
	MongoDBCallsUsageSpellByLevel int
	// MongoDBCallsSetHPTempPlayerByID int
	MongoDBCallsSetHPTempPlayerByID int
	// MongoDBCallsSetSpellUsedPlayerByID int
	MongoDBCallsSetSpellUsedPlayerByID int
	// MongoDBCallsSetArmorWeaponPlayerByID int
	MongoDBCallsSetArmorWeaponPlayerByID int
	// MongoDBCallsDeletePlayerByID int
	MongoDBCallsDeletePlayerByID int
	// MongoDBCallsSetMagicalEffect int
	MongoDBCallsSetMagicalEffect int
	// MongoDBCallsSaveNPC int
	MongoDBCallsSaveNPC int
	// MongoDBCallsGetNPCS int
	MongoDBCallsGetNPCS int
	// MongoDBCallsSetDamageNPCByID int
	MongoDBCallsSetDamageNPCByID int
	// MongoDBCallsSetContitionNPCByID int
	MongoDBCallsSetContitionNPCByID int
	// MongoDBCallsDeleteNPCByID int
	MongoDBCallsDeleteNPCByID int
)

// MongoRepositoryMock struct is used for Mock MongoRepositoryMock requests
type MongoRepositoryMock struct {
}

//SaveCampaign func
func (repo MongoRepositoryMock) SaveCampaign(campaigns *campaign.Campaign) (primitive.ObjectID, error) {
	MongoDBCallsSaveCampaign++
	return primitive.ObjectID{}, nil
}

// UpdateCampaign func
func (repo MongoRepositoryMock) UpdateCampaign(campaignID primitive.ObjectID, campaigns *campaign.Campaign) (int64, error) {
	MongoDBCallsUpdateCampaign++
	return 1, nil
}

// AddPlayerToCampaign func
func (repo MongoRepositoryMock) AddPlayerToCampaign(campaignID primitive.ObjectID, playerID, playerName string) (int64, error) {
	exampleID1, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f8")
	if campaignID == exampleID1 {
		return -1, fmt.Errorf("simulated error")
	}
	MongoDBCallsAddPlayerToCampaign++
	return 1, nil
}

// GetCampaign func
func (repo MongoRepositoryMock) GetCampaign(queryParameters url.Values) ([]*campaign.Campaign, error) {
	campaings := new(campaign.Campaign)
	campaings.MasterID = "testmaster"
	var sliceCampaigns []*campaign.Campaign
	sliceCampaigns = append(sliceCampaigns, campaings)
	MongoDBCallsGetCampaign++
	if len(queryParameters) != 0 {
		// if queryParameters["master_id"][0] == "testerror" {
		// 	return sliceCampaigns, fmt.Errorf("error")
		// }
		// if len(queryParameters["channel_id"]) >= 1 {
		var sliceCampaignsLocal []*campaign.Campaign
		return sliceCampaignsLocal, nil
		// }
	}
	return sliceCampaigns, nil
}

// GetOneCampaign func
func (repo MongoRepositoryMock) GetOneCampaign(campaignID primitive.ObjectID) (*campaign.Campaign, error) {
	campaigns := new(campaign.Campaign)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f6")
	if campaignID == exampleID {
		campaigns.ID = exampleID
		campaigns.PlayersID = []string{"playerID"}
		return campaigns, nil
	}
	exampleID1, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f8")
	if campaignID == exampleID1 {
		return campaigns, fmt.Errorf("simulated error")
	}
	MongoDBCallsGetOneCampaign++
	return campaigns, nil
}

// DeleteCampaignByID func
func (repo MongoRepositoryMock) DeleteCampaignByID(campaignID primitive.ObjectID) (int64, error) {
	MongoDBCallsDeleteCampaignByID++
	return 1, nil
}

//AddAdventure func
func (repo MongoRepositoryMock) AddAdventure(adventures *adventure.Adventure) (primitive.ObjectID, error) {
	MongoDBCallsAddAdventure++
	return primitive.ObjectID{}, nil
}

// UpdateAdventure func
func (repo MongoRepositoryMock) UpdateAdventure(adventureID primitive.ObjectID, adventures *adventure.Adventure) (int64, error) {
	MongoDBCallsUpdateAdventure++
	return 1, nil
}

// GetAdventure func
func (repo MongoRepositoryMock) GetAdventure(queryParameters url.Values) ([]*adventure.Adventure, error) {
	day := new(adventure.Adventure)
	var sliceAdventure []*adventure.Adventure
	sliceAdventure = append(sliceAdventure, day)
	MongoDBCallsGetAdventure++
	return sliceAdventure, nil
}

// GetOneAdventure func
func (repo MongoRepositoryMock) GetOneAdventure(adventureID primitive.ObjectID) (*adventure.Adventure, error) {
	adv := new(adventure.Adventure)
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f8")
	if adventureID == exampleID {
		return adv, fmt.Errorf("simulated error")
	}
	exampleID1, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f6")
	if adventureID == exampleID1 {
		adv.ID = exampleID1
		return adv, nil
	}
	MongoDBCallsGetOneAdventure++
	return adv, nil
}

// AddEncounterToAdventure func
func (repo MongoRepositoryMock) AddEncounterToAdventure(adventureID primitive.ObjectID, encounter string) (int64, error) {
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f8")
	if adventureID == exampleID {
		return 0, fmt.Errorf("sismulated error")
	}
	MongoDBCallsAddEncounterToAdventure++
	return 1, nil
}

//ChangeAdventureStatusByID func
func (repo MongoRepositoryMock) ChangeAdventureStatusByID(adventureID primitive.ObjectID, status string, playersID *adventure.AddPlayersID) (int64, error) {
	MongoDBCallsChangeAdventureStatusByID++
	return 1, nil
}

// DeleteAdventureByID func
func (repo MongoRepositoryMock) DeleteAdventureByID(adventureID primitive.ObjectID) (int64, error) {
	exampleID, _ := primitive.ObjectIDFromHex("5e70e4c5d2f3f777c16b29f8")
	if adventureID == exampleID {
		return 0, fmt.Errorf("simulated error")
	}
	MongoDBCallsDeleteAdventureByID++
	return 1, nil
}

//AddEncounter func
func (repo MongoRepositoryMock) AddEncounter(encounters *encounter.Encounter) (primitive.ObjectID, error) {
	MongoDBCallsAddEncounter++
	return primitive.ObjectID{}, nil
}

//UpdateEncounter func
func (repo MongoRepositoryMock) UpdateEncounter(encounterID primitive.ObjectID, encounters *encounter.Encounter) (int64, error) {
	MongoDBCallsUpdateEncounter++
	return 1, nil
}

// GetEncounter func
func (repo MongoRepositoryMock) GetEncounter(queryParameters url.Values) ([]*encounter.Encounter, error) {
	encounters := new(encounter.Encounter)
	var sliceEncounter []*encounter.Encounter
	sliceEncounter = append(sliceEncounter, encounters)
	MongoDBCallsGetEncounter++
	return sliceEncounter, nil
}

// GetOneEncounter func
func (repo MongoRepositoryMock) GetOneEncounter(encounterID primitive.ObjectID) (*encounter.Encounter, error) {
	encounters := new(encounter.Encounter)
	MongoDBCallsGetOneEncounter++
	return encounters, nil
}

// AddNPCTOEncounter func
func (repo MongoRepositoryMock) AddNPCTOEncounter(encounterID primitive.ObjectID, npc string) (int64, error) {
	MongoDBCallsAddNPCTOEncounter++
	return 1, nil
}

//ChangeEncounterStatusByID func
func (repo MongoRepositoryMock) ChangeEncounterStatusByID(encounterID primitive.ObjectID, status string, playersID *adventure.AddPlayersID) (int64, error) {
	MongoDBCallsChangeEncounterStatusByID++
	return 1, nil
}

// DeleteEncounterByID func
func (repo MongoRepositoryMock) DeleteEncounterByID(encounterID primitive.ObjectID) (int64, error) {
	MongoDBCallsDeleteEncounterByID++
	return 1, nil
}

// GetPlayers func
func (repo MongoRepositoryMock) GetPlayers(queryParameters url.Values) ([]*player.Players, error) {
	play := new(player.Players)
	play.SlackID = "testplayer"
	var slicePlayers []*player.Players
	slicePlayers = append(slicePlayers, play)
	MongoDBCallsGetPlayers++
	if len(queryParameters) != 0 {
		if queryParameters["slack_id"][0] == "testerror" {
			return slicePlayers, fmt.Errorf("error")
		}
	}
	return slicePlayers, nil
}

// GetOnePlayer func
func (repo MongoRepositoryMock) GetOnePlayer(playerID primitive.ObjectID) (*player.Players, error) {
	play := new(player.Players)
	MongoDBCallsGetOnePlayer++
	return play, nil
}

//SavePlayer func
func (repo MongoRepositoryMock) SavePlayer(player *player.Players) (primitive.ObjectID, error) {
	MongoDBCallsSavePlayer++
	return primitive.ObjectID{}, nil
}

//UpdatePlayer func
func (repo MongoRepositoryMock) UpdatePlayer(playerID primitive.ObjectID, player *player.Players) (int64, error) {
	MongoDBCallsUpdatePlayer++
	return 1, nil
}

// AddCampaignToPlayer func
func (repo MongoRepositoryMock) AddCampaignToPlayer(playerID primitive.ObjectID, campaignID, campaignTitle, slackChannelID string) (int64, error) {
	MongoDBCallsAddCampaignToPlayer++
	return 1, nil
}

// ChangePlayerCondition func
func (repo MongoRepositoryMock) ChangePlayerCondition(playerID primitive.ObjectID, player *player.Condition) (int64, error) {
	MongoDBCallsChangePlayerCondition++
	return 1, nil
}

// AddPlayerXP func
func (repo MongoRepositoryMock) AddPlayerXP(playerID primitive.ObjectID, xp int) (int64, error) {
	MongoDBCallsAddPlayerXP++
	return 1, nil
}

// AddOrRemovePlayerHP func
func (repo MongoRepositoryMock) AddOrRemovePlayerHP(playerID primitive.ObjectID, hit int) (int64, error) {
	MongoDBCallsAddOrRemovePlayerHP++
	return 1, nil
}

//CreateInventory func
func (repo MongoRepositoryMock) CreateInventory(inventory *player.Inventory) error {
	MongoDBCallsCreateInventory++
	return nil
}

//DeleteInventoryByID delete Inventory by ID
func (repo MongoRepositoryMock) DeleteInventoryByID(inventoryID primitive.ObjectID) (int64, error) {
	MongoDBCallsDeleteInventoryByID++
	return 1, nil
}

// GetInventoryByID func
func (repo MongoRepositoryMock) GetInventoryByID(inventoryID primitive.ObjectID) (*player.Inventory, error) {
	MongoDBCallsGetInventoryByID++
	inventory := new(player.Inventory)
	return inventory, nil
}

// GetInventoryID func
func (repo MongoRepositoryMock) GetInventoryID(playerID string) (primitive.ObjectID, error) {
	MongoDBCallsGetInventoryID++
	return primitive.ObjectID{}, nil
}

//AddTreasurePlayer func
func (repo MongoRepositoryMock) AddTreasurePlayer(inventoryID primitive.ObjectID, treasure map[string]int) (int64, error) {
	MongoDBCallsAddTreasurePlayer++
	return 1, nil
}

//AddOthersItems func
func (repo MongoRepositoryMock) AddOthersItems(playerID primitive.ObjectID, item []string, magic bool) (int64, error) {
	MongoDBCallsAddOthersItems++
	return 1, nil
}

//RemoveOthersItems func
func (repo MongoRepositoryMock) RemoveOthersItems(playerID primitive.ObjectID, item []string, magic bool) (int64, error) {
	MongoDBCallsRemoveOthersItems++
	return 1, nil
}

// GetArmory func
func (repo MongoRepositoryMock) GetArmory(inventoryID primitive.ObjectID) (*player.Armory, error) {
	MongoDBCallsGetArmory++
	armory := new(player.Armory)
	return armory, nil
}

//UsageSpellByLevel func spells_used.level0
func (repo MongoRepositoryMock) UsageSpellByLevel(playerID primitive.ObjectID, spellByLevel map[string]int) (int64, error) {
	MongoDBCallsUsageSpellByLevel++
	return 1, nil
}

// SetHPTempPlayerByID func
func (repo MongoRepositoryMock) SetHPTempPlayerByID(playerID primitive.ObjectID, hp int) (int64, error) {
	MongoDBCallsSetHPTempPlayerByID++
	return 1, nil
}

// SetSpellUsedPlayerByID func
func (repo MongoRepositoryMock) SetSpellUsedPlayerByID(playerID primitive.ObjectID, spell map[string]int) (int64, error) {
	MongoDBCallsSetSpellUsedPlayerByID++
	return 1, nil
}

// SetArmorWeaponPlayerByID func
func (repo MongoRepositoryMock) SetArmorWeaponPlayerByID(playerID primitive.ObjectID, armory *player.Armory) (int64, error) {
	MongoDBCallsSetArmorWeaponPlayerByID++
	return 1, nil
}

// DeletePlayerByID func
func (repo MongoRepositoryMock) DeletePlayerByID(playerID primitive.ObjectID) (int64, error) {
	MongoDBCallsDeletePlayerByID++
	return 1, nil
}

// SetMagicalEffect
func (repo MongoRepositoryMock) SetMagicalEffect(playerID primitive.ObjectID, item []string, add bool) (int64, error) {
	MongoDBCallsSetMagicalEffect++
	return 1, nil
}

//SaveNPC func
func (repo MongoRepositoryMock) SaveNPC(npc *player.NPC) (primitive.ObjectID, error) {
	MongoDBCallsSaveNPC++
	return primitive.ObjectID{}, nil
}

// GetNPCS func
func (repo MongoRepositoryMock) GetNPCS(queryParameters url.Values) ([]*player.NPC, error) {
	npc := new(player.NPC)
	var sliceNPC []*player.NPC
	sliceNPC = append(sliceNPC, npc)
	MongoDBCallsGetNPCS++
	return sliceNPC, nil
}

// SetDamageNPCByID func
func (repo MongoRepositoryMock) SetDamageNPCByID(npcID primitive.ObjectID, hit int) (int64, error) {
	MongoDBCallsSetDamageNPCByID++
	return 1, nil
}

// SetContitionNPCByID func
func (repo MongoRepositoryMock) SetContitionNPCByID(npcID primitive.ObjectID, npc *player.NPCCondition) (int64, error) {
	MongoDBCallsSetContitionNPCByID++
	return 1, nil
}

// DeleteNPCByID func
func (repo MongoRepositoryMock) DeleteNPCByID(npcID primitive.ObjectID) (int64, error) {
	MongoDBCallsDeleteNPCByID++
	return 1, nil
}

// InitMongoMock func returns a RepositoryMongoMock interface
func InitMongoMock() appcontext.Component {
	return MongoRepositoryMock{}
}
