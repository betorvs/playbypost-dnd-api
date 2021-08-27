package mongodb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetPlayers return all players created
func (repo MongoRepository) GetPlayers(queryParameters url.Values) ([]*player.Players, error) {
	playersResult := make([]*player.Players, 0)
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	numberParameters := len(queryParameters)
	var filter bson.D
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
		var result player.Players
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		playersResult = append(playersResult, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return playersResult, nil
}

// GetOnePlayer func
func (repo MongoRepository) GetOnePlayer(playerID primitive.ObjectID) (*player.Players, error) {
	var player *player.Players
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// search
	filter := bson.M{"_id": playerID}
	result := collection.FindOne(ctx, filter)
	err := result.Decode(&player)
	if err != nil {
		return player, err
	}
	return player, nil
}

//SavePlayer func
func (repo MongoRepository) SavePlayer(player *player.Players) (primitive.ObjectID, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, player)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id := res.InsertedID
	return id.(primitive.ObjectID), nil
}

//UpdatePlayer func
func (repo MongoRepository) UpdatePlayer(playerID primitive.ObjectID, player *player.Players) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// fmt.Printf("%+v\n", player)
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": playerID}
	doc := make(map[string]interface{})
	data, _ := json.Marshal(player)
	_ = json.Unmarshal(data, &doc)
	update := bson.M{"$set": doc}
	res, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

//AddCampaignToPlayer update the config data
func (repo MongoRepository) AddCampaignToPlayer(playerID primitive.ObjectID, campaignID, campaignTitle, slackChannelID string) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": playerID}
	update := bson.M{"$set": bson.M{"game_id": campaignID, "campaign_title": campaignTitle, "slack_channel_id": slackChannelID}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

// ChangePlayerCondition func
func (repo MongoRepository) ChangePlayerCondition(playerID primitive.ObjectID, player *player.Condition) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// fmt.Printf("%+v\n", player)
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": playerID}
	doc := make(map[string]interface{})
	data, _ := json.Marshal(player)
	_ = json.Unmarshal(data, &doc)
	update := bson.M{"$set": doc}
	res, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

//AddPlayerXP update the config data
func (repo MongoRepository) AddPlayerXP(playerID primitive.ObjectID, xp int) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": playerID}
	update := bson.M{"$inc": bson.M{"xp": xp}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no xp was granted to player")
	}
	return res.ModifiedCount, nil
}

//AddOrRemovePlayerHP update the config data
func (repo MongoRepository) AddOrRemovePlayerHP(playerID primitive.ObjectID, hit int) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": playerID}
	update := bson.M{"$inc": bson.M{"hp_temp": hit}}
	// fmt.Println(update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no HP was removed from player")
	}

	return res.ModifiedCount, nil
}

//UsageSpellByLevel func spells_used.level0
func (repo MongoRepository) UsageSpellByLevel(playerID primitive.ObjectID, spellByLevel map[string]int) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": playerID}
	update := bson.M{"$inc": spellByLevel}
	// fmt.Println(update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no spell used for player")
	}

	return res.ModifiedCount, nil
}

//SetHPTempPlayerByID update the config data
func (repo MongoRepository) SetHPTempPlayerByID(playerID primitive.ObjectID, hp int) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": playerID}
	// update := bson.M{"$pull": bson.M{"others_items": bson.M{"$in": item}}}
	update := bson.M{"$set": bson.M{"hp_temp": hp}}
	// fmt.Println(update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no rest for this player")
	}
	return res.ModifiedCount, nil
}

//SetSpellUsedPlayerByID update the config data
func (repo MongoRepository) SetSpellUsedPlayerByID(playerID primitive.ObjectID, spell map[string]int) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": playerID}
	// update := bson.M{"$pull": bson.M{"others_items": bson.M{"$in": item}}}
	update := bson.M{"$set": spell}
	// fmt.Println(update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no rest for this player")
	}
	return res.ModifiedCount, nil
}

//DeletePlayerByID delete Player by ID
func (repo MongoRepository) DeletePlayerByID(playerID primitive.ObjectID) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": playerID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

//CreateInventory func
func (repo MongoRepository) CreateInventory(inventory *player.Inventory) error {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, inventory)
	if err != nil {
		return err
	}
	// id := res.InsertedID
	return nil
}

// GetInventoryByID func
func (repo MongoRepository) GetInventoryByID(inventoryID primitive.ObjectID) (*player.Inventory, error) {
	var inventory *player.Inventory
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// search
	filter := bson.M{"_id": inventoryID}
	result := collection.FindOne(ctx, filter)
	err := result.Decode(&inventory)
	if err != nil {
		return inventory, err
	}
	return inventory, nil
}

//DeleteInventoryByID delete Inventory by ID
func (repo MongoRepository) DeleteInventoryByID(inventoryID primitive.ObjectID) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": inventoryID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

// GetInventoryID func
func (repo MongoRepository) GetInventoryID(playerID string) (primitive.ObjectID, error) {
	var inventory *player.Inventory
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// search
	filter := bson.M{"player_id": playerID}
	result := collection.FindOne(ctx, filter)
	err := result.Decode(&inventory)
	if err != nil {
		return inventory.ID, err
	}
	return inventory.ID, nil
}

//AddTreasurePlayer func
func (repo MongoRepository) AddTreasurePlayer(inventoryID primitive.ObjectID, treasure map[string]int) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": inventoryID}
	update := bson.M{"$inc": treasure}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return res.MatchedCount, nil
}

//AddOthersItems func
func (repo MongoRepository) AddOthersItems(playerID primitive.ObjectID, item []string, magic bool) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": playerID}
	update := bson.M{"$push": bson.M{"others_items": bson.M{"$each": item}}}
	if magic {
		update = bson.M{"$push": bson.M{"magic_items": bson.M{"$each": item}}}
	}
	// fmt.Println(update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	// if res.ModifiedCount == 0 {
	//      return 0, errors.New("No items added")
	// }

	return res.MatchedCount, nil
}

//RemoveOthersItems func
func (repo MongoRepository) RemoveOthersItems(playerID primitive.ObjectID, item []string, magic bool) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": playerID}
	update := bson.M{"$pull": bson.M{"others_items": bson.M{"$in": item}}}
	if magic {
		update = bson.M{"$pull": bson.M{"magic_items": bson.M{"$in": item}}}
	}
	// fmt.Println(update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	// if res.ModifiedCount == 0 {
	//      return 0, errors.New("No items removed")
	// }

	return res.MatchedCount, nil
}

// GetArmory func
func (repo MongoRepository) GetArmory(inventoryID primitive.ObjectID) (*player.Armory, error) {
	var inventory *player.Inventory
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// search
	filter := bson.M{"_id": inventoryID}
	result := collection.FindOne(ctx, filter)
	err := result.Decode(&inventory)
	if err != nil {
		return &inventory.Armory, err
	}
	return &inventory.Armory, nil
}

//SetArmorWeaponPlayerByID update the config data
func (repo MongoRepository) SetArmorWeaponPlayerByID(playerID primitive.ObjectID, armory *player.Armory) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": playerID}
	// update := bson.M{"$pull": bson.M{"others_items": bson.M{"$in": item}}}
	// map
	doc := make(map[string]interface{})
	data, _ := json.Marshal(armory)
	_ = json.Unmarshal(data, &doc)
	update := bson.M{"$set": bson.M{"armory": doc}}
	// update := bson.M{"$set": bson.M{"armor_class_bonus": armorMagic, "armor_name": armor, "weapon_name": weapon, "shield_name": shield, "weapon_bonus": weaponMagic}}
	// fmt.Println(update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no armor or weapon for this player")
	}
	return res.ModifiedCount, nil
}

//SetMagicalEffect func
func (repo MongoRepository) SetMagicalEffect(playerID primitive.ObjectID, item []string, add bool) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionInventory)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": playerID}
	update := bson.M{"$pull": bson.M{"magical_effect": bson.M{"$in": item}}}
	if add {
		update = bson.M{"$push": bson.M{"magical_effect": bson.M{"$each": item}}}
	}
	// fmt.Println(update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	// if res.ModifiedCount == 0 {
	//      return 0, errors.New("No items removed")
	// }

	return res.MatchedCount, nil
}

//SaveNPC func
func (repo MongoRepository) SaveNPC(npc *player.NPC) (primitive.ObjectID, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionNPCS)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, npc)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id := res.InsertedID
	return id.(primitive.ObjectID), nil
}

//GetNPCS return all npcs created
func (repo MongoRepository) GetNPCS(queryParameters url.Values) ([]*player.NPC, error) {
	npcs := make([]*player.NPC, 0)
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionNPCS)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	numberParameters := len(queryParameters)
	var filter bson.D
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
		var result player.NPC
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		npcs = append(npcs, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return npcs, nil
}

//SetDamageNPCByID update the config data
func (repo MongoRepository) SetDamageNPCByID(npcID primitive.ObjectID, hit int) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionNPCS)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// fmt.Println(npcID, hit)
	// opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": npcID}
	update := bson.M{"$inc": bson.M{"hp": hit}}
	// fmt.Println(filter, update)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	if res.ModifiedCount == 0 {
		return 0, errors.New("no damage in NPC")
	}
	return res.ModifiedCount, nil
}

//SetContitionNPCByID update the config data
func (repo MongoRepository) SetContitionNPCByID(npcID primitive.ObjectID, npc *player.NPCCondition) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionNPCS)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fmt.Printf("%+v\n", npc)
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": npcID}
	doc := make(map[string]interface{})
	data, _ := json.Marshal(npc)
	_ = json.Unmarshal(data, &doc)
	update := bson.M{"$set": doc}
	res, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

//DeleteNPCByID delete NPC by ID
func (repo MongoRepository) DeleteNPCByID(npcID primitive.ObjectID) (int64, error) {
	collection := repo.Conn.Database(config.Values.Database).Collection(collectionNPCS)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": npcID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
