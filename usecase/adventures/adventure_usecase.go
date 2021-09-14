package adventures

import (
	"fmt"
	"net/url"

	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/domain/mongodb"
	"github.com/betorvs/playbypost-dnd/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetAllAdventure func
func GetAllAdventure(queryParameters url.Values) ([]*adventure.Adventure, error) {
	repo := mongodb.GetMongoRepository()
	return repo.GetAdventure(queryParameters)
}

//GetOneAdventure func
func GetOneAdventure(adventureID primitive.ObjectID) (*adventure.Adventure, error) {
	repo := mongodb.GetMongoRepository()
	adventures, err := repo.GetOneAdventure(adventureID)
	if err != nil {
		return adventures, err
	}
	return adventures, nil
}

//CreateAdventure func
func CreateAdventure(adventures *adventure.Adventure) (string, error) {
	if len(adventures.Encounters) == 0 {
		adventures.Encounters = []string{}
	}
	if len(adventures.PlayersID) == 0 {
		adventures.PlayersID = []string{}
	}
	repo := mongodb.GetMongoRepository()
	id, err := repo.AddAdventure(adventures)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}

//UpdateAdventure func
func UpdateAdventure(adventureID primitive.ObjectID, adventures *adventure.Adventure) (string, error) {
	repo := mongodb.GetMongoRepository()
	id, err := repo.UpdateAdventure(adventureID, adventures)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}

// DeleteAdventure func
func DeleteAdventure(adventureID primitive.ObjectID) (int64, error) {
	repo := mongodb.GetMongoRepository()
	result, err := repo.DeleteAdventureByID(adventureID)
	if err != nil {
		return 0, err
	}
	return result, nil
}

//CheckAdventureExist func
func CheckAdventureExist(adventureID primitive.ObjectID) bool {
	repo := mongodb.GetMongoRepository()
	adventures, err := repo.GetOneAdventure(adventureID)
	if err != nil {
		return false
	}
	if adventureID == adventures.ID {
		return true
	}
	return false
}

//CheckEncounterWasAdded func
func CheckEncounterWasAdded(adventureID primitive.ObjectID, encounter string) (string, bool) {
	repo := mongodb.GetMongoRepository()
	adventures, err := repo.GetOneAdventure(adventureID)
	if err != nil {
		return "cannot get adventure from database", false
	}
	if utils.StringInSlice(encounter, adventures.Encounters) {
		return "dont exist", true
	}
	return "already exist", false
}

//AddEncounter func
func AddEncounter(adventureID primitive.ObjectID, encounters string) (int64, error) {
	_, checkFightName := CheckEncounterWasAdded(adventureID, encounters)
	if checkFightName {
		return -1, nil
	}
	repo := mongodb.GetMongoRepository()
	result, err := repo.AddEncounterToAdventure(adventureID, encounters)
	if err != nil {
		return 0, err
	}
	return result, nil
}

//ChangeAdventureStatus func
func ChangeAdventureStatus(adventureID primitive.ObjectID, status string, playersID *adventure.AddPlayersID) (int64, error) {
	repo := mongodb.GetMongoRepository()
	result, err := repo.ChangeAdventureStatusByID(adventureID, status, playersID)
	if err != nil {
		return 0, err
	}
	return result, nil
}
