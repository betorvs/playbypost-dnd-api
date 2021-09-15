package playernpc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	playersnpcUsecase "github.com/betorvs/playbypost-dnd/usecase/playersnpc"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllPlayers func
func GetAllPlayers(c echo.Context) (err error) {
	queryParams := c.QueryParams()
	players, err := playersnpcUsecase.GetAllPlayers(queryParams)
	if err != nil {
		errString := "Cannot find any player"
		return c.JSON(http.StatusUnprocessableEntity, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, players)
}

// GetOnePlayer func
func GetOnePlayer(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	players, err := playersnpcUsecase.GetOnePlayer(playerID)
	if err != nil {
		errString := "Cannot find any player"
		return c.JSON(http.StatusUnprocessableEntity, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, players)
}

// PostPlayer func
func PostPlayer(c echo.Context) (err error) {
	player := new(player.Players)
	// fmt.Println("post player")

	if err = c.Bind(player); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	// fmt.Printf("%#v\n", player)
	result, err := playersnpcUsecase.CreatePlayer(player)
	if err != nil {
		errString := fmt.Sprintf("Cannot create new player %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// UpdateOnePlayer func
func UpdateOnePlayer(c echo.Context) (err error) {
	// fmt.Println("update one player")
	// playerID, err := primitive.ObjectIDFromHex(c.Param("id"))
	// if err != nil {
	//      errString := "Cannot parse id"
	//      return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	// }
	player := new(player.Players)
	// fmt.Println("update player")
	if err = c.Bind(player); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if config.Values.RuleBuiltin {
		result, err := playersnpcUsecase.UpdatePlayerRecalc(player.ID, player)
		if err != nil {
			errString := fmt.Sprintf("Cannot update player with recalculation: %v", err)
			return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
		}
		return c.JSON(http.StatusOK, utils.FormatMessage(result))
	}
	result, err := playersnpcUsecase.UpdatePlayer(player.ID, player)
	if err != nil {
		errString := fmt.Sprintf("Cannot update player: %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// DeletePlayer func
func DeletePlayer(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	inventoryID, err := playersnpcUsecase.InventoryIDByPlayerID(playerID)
	if err != nil {
		errString := "Player without Inventory"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.DeletePlayer(playerID, inventoryID)
	if err != nil {
		errString := fmt.Sprintf("Cannot delete player %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	res := fmt.Sprintf("Records Deleted (Players and Inventories): %v", result)
	return c.JSON(http.StatusOK, utils.FormatMessage(res))
}

// AddCampaignToPlayer func
func AddCampaignToPlayer(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	campaigns := new(player.AddCampaign)
	if err = c.Bind(campaigns); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.AddCampaignToPlayer(playerID, campaigns)
	if err != nil {
		errString := fmt.Sprintf("Cannot add campaign %v", err)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// AddOrRemoveHP func
func AddOrRemoveHP(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
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
	result, err := playersnpcUsecase.AddOrRemoveHP(playerID, action, hp)
	if err != nil {
		errString := "Cannot remove HP"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// AddPlayerXP func
func AddPlayerXP(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	value := c.Param("value")
	xp, err := strconv.Atoi(value)
	if err != nil {
		errString := "Value not allowed. Use valid int."
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.AddPlayerXP(playerID, xp)
	if err != nil {
		errString := "Cannot add xp to player"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// UseSpellByLevel func
func UseSpellByLevel(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	level := c.Param("level")
	spellLevel, err := strconv.Atoi(level)
	if err != nil {
		errString := "Value not allowed. Use valid int."
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if spellLevel > 9 || spellLevel < 0 {
		errString := "Spell Level needs to be lower than 9 and higher than 0"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	value := c.Param("value")
	number, err := strconv.Atoi(value)
	if err != nil {
		errString := "Value not allowed. Use valid int."
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.UseSpellByLevel(playerID, spellLevel, number)
	if err != nil {
		errString := "Cannot use spell by level"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// FullRestPlayer func
func FullRestPlayer(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.FullRestPlayer(playerID)
	if err != nil {
		errString := "Cannot fullRestPlayer"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// ChangeCondition func
func ChangeCondition(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	condition := new(player.Condition)
	if err = c.Bind(condition); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.ChangeCondition(playerID, condition)
	if err != nil {
		errString := "Cannot change player condition"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// AddTreasure func
func AddTreasure(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	treasure := new(player.Treasure)
	if err = c.Bind(treasure); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	inventoryID, err := playersnpcUsecase.InventoryIDByPlayerID(playerID)
	if err != nil {
		errString := "Player without Inventory"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.AddTreasure(inventoryID, treasure)
	if err != nil {
		errString := "Cannot add treasure"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// AddOrRemoveOtherItems func
func AddOrRemoveOtherItems(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	allowedAction := []string{"remove", "add"}
	action := c.Param("action")
	if !utils.StringInSlice(action, allowedAction) {
		errString := "Action not allowed. Use remove or add"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	items := new(rule.SimpleList)
	if err = c.Bind(items); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	inventoryID, err := playersnpcUsecase.InventoryIDByPlayerID(playerID)
	if err != nil {
		errString := "Player without Inventory"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.AddOrRemoveOtherItems(inventoryID, action, items.List)
	if err != nil {
		errString := "Cannot add/remove any items"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// AddArmorWeaponPlayerByID func
func AddArmorWeaponPlayerByID(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	ac := new(player.Armory)
	if err = c.Bind(ac); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	inventoryID, err := playersnpcUsecase.InventoryIDByPlayerID(playerID)
	if err != nil {
		errString := "Player without Inventory"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.AddArmorWeaponPlayerByID(playerID, inventoryID, ac)
	if err != nil {
		errString := fmt.Sprintf("Cannot add armor or weapon to a player error %s ", result)
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}

// AddOrRemoveMagicItems func
func AddOrRemoveMagicItems(c echo.Context) (err error) {
	playerID, err := primitive.ObjectIDFromHex(c.Param("playerid"))
	if err != nil {
		errString := "Cannot parse id"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	if !playersnpcUsecase.CheckPlayerExist(playerID) {
		errString := "Player not found"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	allowedAction := []string{"remove", "add", "attune"}
	action := c.Param("action")
	if !utils.StringInSlice(action, allowedAction) {
		errString := "Action not allowed. Use remove or add"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	items := new(rule.SimpleList)
	if err = c.Bind(items); err != nil {
		errString := "Cannot parse json"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	inventoryID, err := playersnpcUsecase.InventoryIDByPlayerID(playerID)
	if err != nil {
		errString := "Player without Inventory"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	result, err := playersnpcUsecase.AddOrRemoveMagicItems(inventoryID, action, items.List)
	if err != nil {
		errString := "Cannot add/remove any items"
		return c.JSON(http.StatusBadRequest, utils.FormatMessage(errString))
	}
	return c.JSON(http.StatusOK, utils.FormatMessage(result))
}
