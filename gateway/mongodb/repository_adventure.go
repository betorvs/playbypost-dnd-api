package mongodb

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddAdventure func
func (repo MongoRepository) AddAdventure(adventures *adventure.Adventure) (primitive.ObjectID, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionAdventure)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, adventures)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id := res.InsertedID
	return id.(primitive.ObjectID), nil
}

// UpdateAdventure func
func (repo MongoRepository) UpdateAdventure(adventureID primitive.ObjectID, adventures *adventure.Adventure) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionAdventure)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": adventureID}
	doc := make(map[string]interface{})
	data, _ := json.Marshal(adventures)
	_ = json.Unmarshal(data, &doc)
	update := bson.M{"$set": doc}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no adventure updated")
	}
	return res.ModifiedCount, nil
}

//GetAdventure return all games created
func (repo MongoRepository) GetAdventure(queryParameters url.Values) ([]*adventure.Adventure, error) {
	adventures := make([]*adventure.Adventure, 0)
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionAdventure)
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
		var result adventure.Adventure
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		adventures = append(adventures, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return adventures, nil
}

// GetOneAdventure func
func (repo MongoRepository) GetOneAdventure(adventureID primitive.ObjectID) (*adventure.Adventure, error) {
	var adventures *adventure.Adventure
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionAdventure)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// search
	filter := bson.M{"_id": adventureID}
	result := collection.FindOne(ctx, filter)
	err := result.Decode(&adventures)
	if err != nil {
		return adventures, err
	}
	return adventures, nil
}

//AddEncounterToAdventure update the config data
func (repo MongoRepository) AddEncounterToAdventure(adventureID primitive.ObjectID, encounter string) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionEncounter)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": adventureID}
	update := bson.M{"$push": bson.M{"encounters": encounter}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no fight scenes added")
	}

	return res.ModifiedCount, nil
}

//ChangeAdventureStatusByID func
func (repo MongoRepository) ChangeAdventureStatusByID(adventureID primitive.ObjectID, status string, playersID *adventure.AddPlayersID) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionAdventure)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": adventureID}
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

//DeleteAdventureByID delete game by ID
func (repo MongoRepository) DeleteAdventureByID(adventureID primitive.ObjectID) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionAdventure)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": adventureID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
