package mongodb

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/domain/encounter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddEncounter func
func (repo MongoRepository) AddEncounter(encounters *encounter.Encounter) (primitive.ObjectID, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionEncounter)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, encounters)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id := res.InsertedID
	return id.(primitive.ObjectID), nil
}

// UpdateEncounter func
func (repo MongoRepository) UpdateEncounter(encounterID primitive.ObjectID, encounters *encounter.Encounter) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionEncounter)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": encounterID}
	doc := make(map[string]interface{})
	data, _ := json.Marshal(encounters)
	_ = json.Unmarshal(data, &doc)
	update := bson.M{"$set": doc}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no Encounter updated")
	}
	return res.ModifiedCount, nil
}

//GetEncounter return all Encounter created
func (repo MongoRepository) GetEncounter(queryParameters url.Values) ([]*encounter.Encounter, error) {
	encounters := make([]*encounter.Encounter, 0)
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionEncounter)
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
		var result encounter.Encounter
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		encounters = append(encounters, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return encounters, nil
}

// GetOneEncounter func
func (repo MongoRepository) GetOneEncounter(encounterID primitive.ObjectID) (*encounter.Encounter, error) {
	var encounters *encounter.Encounter
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionEncounter)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// search
	filter := bson.M{"_id": encounterID}
	result := collection.FindOne(ctx, filter)
	err := result.Decode(&encounters)
	if err != nil {
		return encounters, err
	}
	return encounters, nil
}

//AddNPCTOEncounter update the config data
func (repo MongoRepository) AddNPCTOEncounter(encounterID primitive.ObjectID, npc string) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionEncounter)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": encounterID}
	update := bson.M{"$push": bson.M{"npcs": npc}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no NPCs added")
	}

	return res.ModifiedCount, nil
}

//ChangeEncounterStatusByID func
func (repo MongoRepository) ChangeEncounterStatusByID(encounterID primitive.ObjectID, status string, playersID *adventure.AddPlayersID) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionEncounter)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": encounterID}
	update := bson.M{"$set": bson.M{"status": status}}
	if len(playersID.PlayersID) != 0 {
		update = bson.M{"$set": bson.M{"status": status, "players_id": playersID.PlayersID}}
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no Status changed")
	}
	return res.ModifiedCount, nil
}

//DeleteEncounterByID delete game by ID
func (repo MongoRepository) DeleteEncounterByID(encounterID primitive.ObjectID) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionEncounter)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": encounterID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
