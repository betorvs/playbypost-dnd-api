package mongodb

import (
	"net/url"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/domain/campaign"
	"github.com/betorvs/playbypost-dnd/domain/encounter"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MongoRepository interface
type MongoRepository interface {
	appcontext.Component
	// Campaign interfaces

	// SaveCampaign creates a campaing
	SaveCampaign(campaign *campaign.Campaign) (primitive.ObjectID, error)
	// UpdateCampaign updates a campaign
	UpdateCampaign(campaignID primitive.ObjectID, campaign *campaign.Campaign) (int64, error)
	// AddPlayerToCampaign add player to campaign
	AddPlayerToCampaign(campaignID primitive.ObjectID, playerID, playerName string) (int64, error)
	// GetCampaign Get all campaigns
	GetCampaign(queryParameters url.Values) ([]*campaign.Campaign, error)
	// GetOneCampaign get campaign by id
	GetOneCampaign(campaignID primitive.ObjectID) (*campaign.Campaign, error)
	// DeleteCampaignByID deletes campaign by id
	DeleteCampaignByID(campaignID primitive.ObjectID) (int64, error)

	// Adventure interfaces

	// AddAdventure creates an adventure
	AddAdventure(adventures *adventure.Adventure) (primitive.ObjectID, error)
	// UpdateAdventure updates an adventure
	UpdateAdventure(adventureID primitive.ObjectID, adventure *adventure.Adventure) (int64, error)
	// GetAdventure get all adventures
	GetAdventure(queryParameters url.Values) ([]*adventure.Adventure, error)
	// GetOneAdventure get adventure by id
	GetOneAdventure(adventureID primitive.ObjectID) (*adventure.Adventure, error)
	// AddEncounterToAdventure add encounter to adventure
	AddEncounterToAdventure(adventureID primitive.ObjectID, encounter string) (int64, error)
	// ChangeAdventureStatusByID change adventure status
	ChangeAdventureStatusByID(adventureID primitive.ObjectID, status string, playersID *adventure.AddPlayersID) (int64, error)
	// DeleteAdventureByID delete an adventure by id
	DeleteAdventureByID(adventureID primitive.ObjectID) (int64, error)

	// Encounters interfaces

	// AddEncounter creates an encounter
	AddEncounter(encounters *encounter.Encounter) (primitive.ObjectID, error)
	// UpdateEncounter updates an encounter
	UpdateEncounter(encounterID primitive.ObjectID, encounters *encounter.Encounter) (int64, error)
	// GetEncounter get all encounters
	GetEncounter(queryParameters url.Values) ([]*encounter.Encounter, error)
	// GetOneEncounter returns an encounter by id
	GetOneEncounter(encounterID primitive.ObjectID) (*encounter.Encounter, error)
	// AddNPCTOEncounter add npc to be used in an encounter
	AddNPCTOEncounter(encounterID primitive.ObjectID, npc string) (int64, error)
	// ChangeEncounterStatusByID change encounter status
	ChangeEncounterStatusByID(encounterID primitive.ObjectID, status string, playersID *adventure.AddPlayersID) (int64, error)
	// DeleteEncounterByID delete encounter by id
	DeleteEncounterByID(encounterID primitive.ObjectID) (int64, error)

	//Players Interfaces

	// Get all players
	GetPlayers(queryParameters url.Values) ([]*player.Players, error)
	// GetOnePlayer get one player by id
	GetOnePlayer(playerID primitive.ObjectID) (*player.Players, error)
	// SavePlayer creates a player
	SavePlayer(player *player.Players) (primitive.ObjectID, error)
	// UpdatePlayer update player
	UpdatePlayer(playerID primitive.ObjectID, player *player.Players) (int64, error)
	// AddCampaignToPlayer add campaign to a player
	AddCampaignToPlayer(playerID primitive.ObjectID, campaignID, campaignTitle, slackChannelID string) (int64, error)
	// ChangePlayerCondition change player condition
	ChangePlayerCondition(playerID primitive.ObjectID, player *player.Condition) (int64, error)
	// AddPlayerXP add xp to a player
	AddPlayerXP(playerID primitive.ObjectID, xp int) (int64, error)
	// AddOrRemovePlayerHP add or remove players HP
	AddOrRemovePlayerHP(playerID primitive.ObjectID, hit int) (int64, error)
	// UsageSpellByLevel return spell used by a player
	UsageSpellByLevel(playerID primitive.ObjectID, spellByLevel map[string]int) (int64, error)
	// SetHPTempPlayerByID change players HP
	SetHPTempPlayerByID(playerID primitive.ObjectID, hp int) (int64, error)
	// SetSpellUsedPlayerByID configure players spells
	SetSpellUsedPlayerByID(playerID primitive.ObjectID, spell map[string]int) (int64, error)
	// DeletePlayerByID delete a player by id
	DeletePlayerByID(playerID primitive.ObjectID) (int64, error)
	// CreateInventory creates a player inventory
	CreateInventory(inventory *player.Inventory) error
	// GetInventoryByID gets an inventory by ID
	GetInventoryByID(inventoryID primitive.ObjectID) (*player.Inventory, error)
	// DeleteInventoryByID delete an inventory by ID
	DeleteInventoryByID(inventoryID primitive.ObjectID) (int64, error)
	// GetInventoryID gets player inventory
	GetInventoryID(playerID string) (primitive.ObjectID, error)
	// AddTreasurePlayer add treasure into player inventory
	AddTreasurePlayer(inventoryID primitive.ObjectID, treasure map[string]int) (int64, error)
	// AddOthersItems Add others itens
	AddOthersItems(playerID primitive.ObjectID, item []string, magic bool) (int64, error)
	// RemoveOthersItems remove items
	RemoveOthersItems(playerID primitive.ObjectID, item []string, magic bool) (int64, error)
	// GetArmory gets players armory
	GetArmory(inventoryID primitive.ObjectID) (*player.Armory, error)
	// SetArmorWeaponPlayerByID update the config data
	SetArmorWeaponPlayerByID(playerID primitive.ObjectID, armory *player.Armory) (int64, error)
	// SetMagicalEffect apply a magical effect in a player
	SetMagicalEffect(playerID primitive.ObjectID, item []string, add bool) (int64, error)
	// SaveNPC creates a npc
	SaveNPC(npc *player.NPC) (primitive.ObjectID, error)
	// GetNPCS return all npcs created
	GetNPCS(queryParameters url.Values) ([]*player.NPC, error)
	// SetDamageNPCByID update the config data
	SetDamageNPCByID(npcID primitive.ObjectID, hit int) (int64, error)
	// SetContitionNPCByID update the config data
	SetContitionNPCByID(npcID primitive.ObjectID, npc *player.NPCCondition) (int64, error)
	// DeleteNPCByID delete NPC by ID
	DeleteNPCByID(npcID primitive.ObjectID) (int64, error)
}

//GetMongoRepository gets the MongoRepository current implementation
func GetMongoRepository() MongoRepository {
	return appcontext.Current.Get(appcontext.MongoRepository).(MongoRepository)
}
