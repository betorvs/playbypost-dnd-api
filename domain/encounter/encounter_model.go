package encounter

import "go.mongodb.org/mongo-driver/bson/primitive"

// Encounter struct
type Encounter struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CampaignID  string             `bson:"campaign_id,omitempty" json:"campaign_id,omitempty"`
	AdventureID string             `bson:"adventure_id,omitempty" json:"adventure_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	ChannelID   string             `bson:"channel_id" json:"channel_id"`
	NPCS        []string           `bson:"npcs" json:"npcs"`
	PlayersID   []string           `bson:"players_id,omitempty" json:"players_id,omitempty"`
	Status      string             `bson:"status" json:"status"`
}

//AddNPC struct
type AddNPC struct {
	NPCS        string `json:"npcs"`
	EncounterID string `bson:"encounter_id,omitempty" json:"encounter_id,omitempty"`
}
