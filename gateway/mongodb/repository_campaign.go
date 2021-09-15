package mongodb

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/campaign"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveCampaign func
func (repo MongoRepository) SaveCampaign(campaigns *campaign.Campaign) (primitive.ObjectID, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionCampaign)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, campaigns)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id := res.InsertedID
	return id.(primitive.ObjectID), nil
}

//UpdateCampaign update the config data
func (repo MongoRepository) UpdateCampaign(campaignID primitive.ObjectID, campaigns *campaign.Campaign) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionCampaign)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": campaignID}
	doc := make(map[string]interface{})
	data, _ := json.Marshal(campaigns)
	_ = json.Unmarshal(data, &doc)
	update := bson.M{"$set": doc}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no campaign updated")
	}
	return res.ModifiedCount, nil
}

//AddPlayerToCampaign update the config data
func (repo MongoRepository) AddPlayerToCampaign(campaignID primitive.ObjectID, playerID, playerName string) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionCampaign)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": campaignID}
	update := bson.M{"$push": bson.M{"players_id": playerID, "players_username": playerName}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no players added")
	}

	return res.ModifiedCount, nil
}

//GetCampaign return all campaigns created
func (repo MongoRepository) GetCampaign(queryParameters url.Values) ([]*campaign.Campaign, error) {
	campaignsResult := make([]*campaign.Campaign, 0)
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionCampaign)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var filter bson.D
	numberParameters := len(queryParameters)
	if numberParameters == 0 {
		filter = bson.D{{}}
	} else {
		var allDataFilter = bson.A{}
		for paramName, param := range queryParameters {
			dataFilter := bson.D{primitive.E{Key: paramName, Value: bson.D{primitive.E{Key: "$in", Value: param}}}}
			allDataFilter = append(allDataFilter, dataFilter)
		}
		filter = bson.D{primitive.E{Key: "$or", Value: allDataFilter}}
	}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result campaign.Campaign
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		campaignsResult = append(campaignsResult, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return campaignsResult, nil
}

// GetOneCampaign func
func (repo MongoRepository) GetOneCampaign(campaignID primitive.ObjectID) (*campaign.Campaign, error) {
	var campaign *campaign.Campaign
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionCampaign)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// search
	filter := bson.M{"_id": campaignID}
	result := collection.FindOne(ctx, filter)
	err := result.Decode(&campaign)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

//DeleteCampaignByID delete game by ID
func (repo MongoRepository) DeleteCampaignByID(gameID primitive.ObjectID) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionCampaign)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": gameID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
