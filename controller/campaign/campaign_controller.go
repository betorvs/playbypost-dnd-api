package campaign

import (
	"fmt"
	"net/http"

	"github.com/betorvs/playbypost-dnd/domain/campaign"
	campaignsUsecase "github.com/betorvs/playbypost-dnd/usecase/campaigns"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllCampaign func
func GetAllCampaign(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	res, err := campaignsUsecase.GetAllCampaigns(queryParams)
	if err != nil {
		errString := "Cannot find any campaigns"
		return c.JSON(http.StatusBadGateway, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, res)
}

// GetOneCampaign func
func GetOneCampaign(c echo.Context) (err error) {
	campaignID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse campaign id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	res, err := campaignsUsecase.GetOneCampaign(campaignID)
	if err != nil {
		errString := "Cannot find any campaign"
		return c.JSON(http.StatusBadGateway, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, res)
}

// PostCampaign func
func PostCampaign(c echo.Context) (err error) {
	postBody := new(campaign.Campaign)
	if err = c.Bind(postBody); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !utils.StringInSlice(postBody.Status, utils.AllowedStatus()) {
		errString := fmt.Sprintf("Campaign can only have the following status %v", utils.AllowedStatus())
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := campaignsUsecase.CreateCampaign(postBody)
	if err != nil {
		errString := fmt.Sprintf("Cannot parse id %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// PutCampaign func
func PutCampaign(c echo.Context) (err error) {
	putBody := new(campaign.Campaign)
	if err = c.Bind(putBody); err != nil {
		errString := "Cannot parse json 2"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !utils.StringInSlice(putBody.Status, utils.AllowedStatus()) {
		errString := fmt.Sprintf("Campaign can only have the following status %v", utils.AllowedStatus())
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := campaignsUsecase.UpdateCampaign(putBody.ID, putBody)
	if err != nil {
		errString := fmt.Sprintf("Cannot parse id %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// CheckPlayerAllowed func
func CheckPlayerAllowed(c echo.Context) (err error) {
	campaignID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	playerid := c.Param("playerid")
	result, check := campaignsUsecase.CheckListPlayers(campaignID, playerid)
	if !check {
		return c.JSON(http.StatusForbidden, utils.FormatMessage(result))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// AddPlayerCampaign func
func AddPlayerCampaign(c echo.Context) (err error) {
	campaignID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	player := new(campaign.AddPlayer)
	if err = c.Bind(player); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := campaignsUsecase.AddPlayer(campaignID, player.PlayerID, player.PlayerUsername)
	if err != nil {
		errString := fmt.Sprintf("Cannot add player: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}

	res := fmt.Sprintf("Records Changed: %v", result)
	if result == -1 {
		res = fmt.Sprintf("Player already registered: %s", player.PlayerUsername)
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}

// DeleteCampaign func
func DeleteCampaign(c echo.Context) (err error) {
	gameID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := campaignsUsecase.DeleteCampaign(gameID)
	if err != nil {
		errString := fmt.Sprintf("Cannot delete campaign: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	res := fmt.Sprintf("Records Deleted: %v", result)
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}
