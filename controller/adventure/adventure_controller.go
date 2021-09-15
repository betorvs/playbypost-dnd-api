package adventure

import (
	"fmt"
	"net/http"

	"github.com/betorvs/playbypost-dnd/domain/adventure"
	adventuresUsecase "github.com/betorvs/playbypost-dnd/usecase/adventures"
	campaignsUsecase "github.com/betorvs/playbypost-dnd/usecase/campaigns"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllAdventure func
func GetAllAdventure(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	res, err := adventuresUsecase.GetAllAdventure(queryParams)
	if err != nil {
		errString := "Cannot find any adventure"
		return c.JSON(http.StatusUnprocessableEntity, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, res)
}

// GetOneAdventure func
func GetOneAdventure(c echo.Context) (err error) {
	adventureID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse adventure id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	fmt.Println(adventureID)
	res, err := adventuresUsecase.GetOneAdventure(adventureID)
	if err != nil {
		errString := "Cannot find any adventure"
		return c.JSON(http.StatusUnprocessableEntity, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, res)
}

// PostAdventure func
func PostAdventure(c echo.Context) (err error) {
	postBody := new(adventure.Adventure)
	if err = c.Bind(postBody); err != nil {
		errString := fmt.Sprintf("Cannot parse json adventure %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !utils.StringInSlice(postBody.Status, utils.AllowedStatus()) {
		errString := fmt.Sprintf("Adventure can only have the following status %v", utils.AllowedStatus())
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	campaignID, err := primitive.ObjectIDFromHex(postBody.CampaignID)
	if err != nil {
		errString := "Cannot parse campaign id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !campaignsUsecase.CheckCampaignExist(campaignID) {
		errString := "campaign not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := adventuresUsecase.CreateAdventure(postBody)
	if err != nil {
		errString := "Cannot create adventure"
		return c.JSON(http.StatusUnprocessableEntity, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// PutAdventure func
func PutAdventure(c echo.Context) (err error) {
	putBody := new(adventure.Adventure)
	if err = c.Bind(putBody); err != nil {
		errString := fmt.Sprintf("Cannot parse json adventure %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !utils.StringInSlice(putBody.Status, utils.AllowedStatus()) {
		errString := fmt.Sprintf("Adventure can only have the following status %v", utils.AllowedStatus())
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	advID, err := primitive.ObjectIDFromHex(putBody.CampaignID)
	if err != nil {
		errString := "Cannot parse campaign id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !campaignsUsecase.CheckCampaignExist(advID) {
		errString := "Campaign not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := adventuresUsecase.UpdateAdventure(putBody.ID, putBody)
	if err != nil {
		errString := "Cannot create adventure"
		return c.JSON(http.StatusUnprocessableEntity, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// DeleteAdventure func
func DeleteAdventure(c echo.Context) (err error) {
	adventureID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := adventuresUsecase.DeleteAdventure(adventureID)
	if err != nil {
		errString := "Cannot delete adventure"
		return c.JSON(http.StatusUnprocessableEntity, utils.FormatMessage(errString))
	}
	res := fmt.Sprintf("Records Deleted: %v", result)
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}

// AddEncounter func
func AddEncounter(c echo.Context) (err error) {
	encounter := new(adventure.AddEncounters)
	if err = c.Bind(encounter); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	adventureID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse adventure id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !adventuresUsecase.CheckAdventureExist(adventureID) {
		errString := "Adventure not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := adventuresUsecase.AddEncounter(adventureID, encounter.Encounter)
	if err != nil {
		errString := fmt.Sprintf("Cannot add encounter: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}

	res := fmt.Sprintf("Records Changed: %v", result)
	if result == -1 {
		res = fmt.Sprintf("encounter already registered: %s", encounter.Encounter)
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}

// ChangeAdventureStatus func
func ChangeAdventureStatus(c echo.Context) (err error) {
	adventureID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse adventure id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !adventuresUsecase.CheckAdventureExist(adventureID) {
		errString := "adventure not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	adventureStatus := c.Param("status")
	if !utils.StringInSlice(adventureStatus, utils.AllowedStatus()) {
		errString := fmt.Sprintf("Adventure can only have the following status %v", utils.AllowedStatus())
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	playersID := new(adventure.AddPlayersID)
	if err = c.Bind(playersID); err != nil {
		errString := fmt.Sprintf("Cannot parse json playersID %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := adventuresUsecase.ChangeAdventureStatus(adventureID, adventureStatus, playersID)
	if err != nil {
		errString := fmt.Sprintf("Cannot change adventure status: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}

	res := fmt.Sprintf("Records Changed: %v", result)
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}
