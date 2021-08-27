package adventure

import "go.mongodb.org/mongo-driver/bson/primitive"

// Adventure struct
type Adventure struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CampaignID  string             `bson:"campaign_id,omitempty" json:"campaign_id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	ChannelID   string             `bson:"channel_id" json:"channel_id"`
	Encounters  []string           `bson:"encounters,omitempty" json:"encounters,omitempty"`
	PlayersID   []string           `bson:"players_id,omitempty" json:"players_id,omitempty"`
	Status      string             `bson:"status" json:"status"`
}

//AddEncounters struct
type AddEncounters struct {
	Encounter   string `json:"encounter"`
	AdventureID string `bson:"adventure_id,omitempty" json:"adventure_id,omitempty"`
}

//AddPlayersID struct
type AddPlayersID struct {
	PlayersID []string `json:"players_id"`
}
