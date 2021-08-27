package encounter

import (
	"fmt"
	"net/http"

	"github.com/betorvs/playbypost-dnd/domain/adventure"
	"github.com/betorvs/playbypost-dnd/domain/encounter"
	encountersUsecase "github.com/betorvs/playbypost-dnd/usecase/encounters"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllEncounter func
func GetAllEncounter(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	fight, err := encountersUsecase.GetEncounter(queryParams)
	if err != nil {
		errString := "Cannot find any encounter"
		return c.JSON(http.StatusBadGateway, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, fight)
}

// GetOneEncounter func
func GetOneEncounter(c echo.Context) (err error) {
	encounterID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse encounter id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	res, err := encountersUsecase.GetOneEncounter(encounterID)
	if err != nil {
		errString := "Cannot find any encounter "
		return c.JSON(http.StatusBadGateway, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, res)
}

// PostEncounter func
func PostEncounter(c echo.Context) (err error) {
	postBody := new(encounter.Encounter)
	if err = c.Bind(postBody); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !utils.StringInSlice(postBody.Status, utils.AllowedStatus()) {
		errString := fmt.Sprintf("Encounter can only have the following status %v", utils.AllowedStatus())
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := encountersUsecase.CreateEncounter(postBody)
	if err != nil {
		errString := fmt.Sprintf("Cannot create encounter: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// PutEncounter func
func PutEncounter(c echo.Context) (err error) {
	putBody := new(encounter.Encounter)
	if err = c.Bind(putBody); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !utils.StringInSlice(putBody.Status, utils.AllowedStatus()) {
		errString := fmt.Sprintf("Encounter can only have the following status %v", utils.AllowedStatus())
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := encountersUsecase.UpdateEncounter(putBody.ID, putBody)
	if err != nil {
		errString := fmt.Sprintf("Cannot create encounter: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// DeleteEncounter func
func DeleteEncounter(c echo.Context) (err error) {
	fightID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := encountersUsecase.DeleteEncounter(fightID)
	if err != nil {
		errString := fmt.Sprintf("Cannot delete encounter: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	res := fmt.Sprintf("Records Deleted: %v", result)
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}

// AddNPC func
func AddNPC(c echo.Context) (err error) {
	npc := new(encounter.AddNPC)
	if err = c.Bind(npc); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	encounterID, err := primitive.ObjectIDFromHex(npc.EncounterID)
	if err != nil {
		errString := "Cannot parse encounter id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := encountersUsecase.AddNPC(encounterID, npc.NPCS)
	if err != nil {
		errString := fmt.Sprintf("Cannot add NPC to encounter: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	res := fmt.Sprintf("Records Changed: %v", result)
	if result == -1 {
		res = fmt.Sprintf("NPC in encounter already registered: %s", npc.NPCS)
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}

// ChangeEncounterStatus func
func ChangeEncounterStatus(c echo.Context) (err error) {
	encounterID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse encounter id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !encountersUsecase.CheckEncounterExist(encounterID) {
		errString := "Encounter not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	encounterStatus := c.Param("status")
	if !utils.StringInSlice(encounterStatus, utils.AllowedStatus()) {
		errString := fmt.Sprintf("Encounter can only have the following status %v", utils.AllowedStatus())
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	playersID := new(adventure.AddPlayersID)
	if err = c.Bind(playersID); err != nil {
		errString := fmt.Sprintf("Cannot parse json playersID %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := encountersUsecase.ChangeEncounterStatus(encounterID, encounterStatus, playersID)
	if err != nil {
		errString := fmt.Sprintf("Cannot change status for encounter: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	res := fmt.Sprintf("Records Changed: %v", result)
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}
