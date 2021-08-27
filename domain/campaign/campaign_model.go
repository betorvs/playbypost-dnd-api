package campaign

import "go.mongodb.org/mongo-driver/bson/primitive"

// Campaign struct
type Campaign struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title           string             `bson:"title" json:"title"`
	ChannelID       string             `bson:"channel_id" json:"channel_id"`
	ChannelName     string             `bson:"channel_name" json:"channel_name"`
	TeamDomain      string             `bson:"team_domain" json:"team_domain"`
	TeamID          string             `bson:"team_id" json:"team_id"`
	MasterID        string             `bson:"master_id" json:"master_id"`
	MasterName      string             `bson:"master_name" json:"master_name"`
	Description     string             `bson:"description" json:"description"`
	PlayersID       []string           `bson:"players_id,omitempty" json:"players_id,omitempty"`
	PlayersUsername []string           `bson:"players_username,omitempty" json:"players_username,omitempty"`
	StartLevel      int                `bson:"start_level" json:"start_level"`
	Status          string             `bson:"status" json:"status"`
}

//AddPlayer struct
type AddPlayer struct {
	PlayerID       string `json:"player_id"`
	PlayerUsername string `json:"player_username"`
}
