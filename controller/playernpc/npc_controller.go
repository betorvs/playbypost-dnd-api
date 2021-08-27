package playernpc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/betorvs/playbypost-dnd/domain/player"
	playersnpcUsecase "github.com/betorvs/playbypost-dnd/usecase/playersnpc"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllNPC func
func GetAllNPC(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	npcs, err := playersnpcUsecase.GetAllNPCS(queryParams)
	if err != nil {
		errString := "Cannot find any npc"
		return c.JSON(http.StatusBadGateway, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, npcs)
}

// PostNPC func
func PostNPC(c echo.Context) (err error) {
	npc := new(player.NPC)
	if err = c.Bind(npc); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.CreateNPC(npc)
	if err != nil {
		errString := fmt.Sprintf("Cannot create new npc %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// DeleteNPC func
func DeleteNPC(c echo.Context) (err error) {
	npcID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.DeleteNPC(npcID)
	if err != nil {
		errString := fmt.Sprintf("Cannot delete npc %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	res := fmt.Sprintf("Records Deleted: %v", result)
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}

// PostDamageNPC func
func PostDamageNPC(c echo.Context) (err error) {
	// fmt.Println("damage npc")
	allowedAction := []string{"remove", "add"}
	action := c.Param("action")
	if !utils.StringInSlice(action, allowedAction) {
		errString := "Action not allowed. Use remove or add"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	value := c.Param("value")
	hp, err := strconv.Atoi(value)
	if err != nil {
		errString := "Value not allowed. Use valid int."
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	npcID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.DamageNPC(npcID, action, hp)
	if err != nil {
		errString := fmt.Sprintf("Cannot add damage on player: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// ChangeNPCCondition func
func ChangeNPCCondition(c echo.Context) (err error) {
	npcID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	// if !playersnpcUsecase.CheckPlayerExist(npcID) {
	//      errString := fmt.Sprintf("Player not found")
	//      return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	// }
	condition := new(player.NPCCondition)
	if err = c.Bind(condition); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.ChangeNPCCondition(npcID, condition)
	if err != nil {
		errString := fmt.Sprintf("Cannot change NPC condition: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}
