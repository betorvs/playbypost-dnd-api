package encounters

import (
	"fmt"
	"net/url"

	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/domain/encounter"
	"github.com/betorvs/playbypost-dnd/domain/mongodb"
	"github.com/betorvs/playbypost-dnd/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetEncounter func
func GetEncounter(queryParameters url.Values) ([]*encounter.Encounter, error) {
	repo := mongodb.GetMongoRepository()
	return repo.GetEncounter(queryParameters)
}

//GetOneEncounter func
func GetOneEncounter(encounterID primitive.ObjectID) (*encounter.Encounter, error) {
	repo := mongodb.GetMongoRepository()
	encounters, err := repo.GetOneEncounter(encounterID)
	if err != nil {
		return encounters, err
	}
	return encounters, nil
}

//CreateEncounter func
func CreateEncounter(encounters *encounter.Encounter) (string, error) {
	if len(encounters.NPCS) == 0 {
		encounters.NPCS = []string{}
	}
	if len(encounters.PlayersID) == 0 {
		encounters.PlayersID = []string{}
	}
	repo := mongodb.GetMongoRepository()
	id, err := repo.AddEncounter(encounters)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}

//UpdateEncounter func
func UpdateEncounter(encounterID primitive.ObjectID, encounters *encounter.Encounter) (string, error) {
	if len(encounters.NPCS) == 0 {
		encounters.NPCS = []string{}
	}
	if len(encounters.PlayersID) == 0 {
		encounters.PlayersID = []string{}
	}
	repo := mongodb.GetMongoRepository()
	id, err := repo.UpdateEncounter(encounterID, encounters)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}

// DeleteEncounter func
func DeleteEncounter(encounterID primitive.ObjectID) (int64, error) {
	repo := mongodb.GetMongoRepository()
	result, err := repo.DeleteEncounterByID(encounterID)
	if err != nil {
		return 0, err
	}
	return result, nil
}

//CheckNPCExist func
func CheckNPCExist(encounterID primitive.ObjectID, npc string) (string, bool) {
	repo := mongodb.GetMongoRepository()
	encounters, err := repo.GetOneEncounter(encounterID)
	if err != nil {
		return "cannot get Encounter from database", false
	}
	if utils.StringInSlice(npc, encounters.NPCS) {
		return "dont exist", true
	}
	return "already exist", false
}

//AddNPC func
func AddNPC(encounterID primitive.ObjectID, npc string) (int64, error) {
	_, checkEncounter := CheckNPCExist(encounterID, npc)
	if checkEncounter {
		return -1, nil
	}
	repo := mongodb.GetMongoRepository()
	result, err := repo.AddNPCTOEncounter(encounterID, npc)
	if err != nil {
		return 0, err
	}

	return result, nil
}

//CheckEncounterExist func
func CheckEncounterExist(encounterID primitive.ObjectID) bool {
	repo := mongodb.GetMongoRepository()
	encounters, err := repo.GetOneEncounter(encounterID)
	if err != nil {
		return false
	}
	if encounterID == encounters.ID {
		return true
	}
	return false
}

//ChangeEncounterStatus func
func ChangeEncounterStatus(encounterID primitive.ObjectID, status string, playersID *adventure.AddPlayersID) (int64, error) {
	repo := mongodb.GetMongoRepository()
	result, err := repo.ChangeEncounterStatusByID(encounterID, status, playersID)
	if err != nil {
		return 0, err
	}
	return result, nil
}
