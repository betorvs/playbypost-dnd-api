package controller

import (
	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/controller/adventure"
	"github.com/betorvs/playbypost-dnd/controller/campaign"
	"github.com/betorvs/playbypost-dnd/controller/encounter"
	"github.com/betorvs/playbypost-dnd/controller/playernpc"
	"github.com/betorvs/playbypost-dnd/controller/rule"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//MapRoutes for the endpoints which the API listens for
func MapRoutes(e *echo.Echo) {
	g := e.Group("/playbypost-dnd/v1")
	if config.Values.UsePrometheus {
		p := prometheus.NewPrometheus("echo", nil)
		p.Use(e)
	}
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	g.GET("/health", CheckHealth)
	g.GET("/ready", CheckReady)
	g.GET("/info", GetInfo)
	// roll dice
	g.PUT("/roll/:dice", RunDiceRoll)

	// rule controller
	// rule items and shop
	g.GET("/rule/monster", rule.GetMonsterForNPC)
	g.GET("/rule/spell", rule.GetSpellListDescription)
	g.GET("/rule/magicitem", rule.GetMagicItem)
	g.GET("/rule/weapon", rule.GetAllWeapons)
	g.GET("/rule/armor", rule.GetAllArmors)
	g.GET("/rule/gear", rule.GetAllGear)
	g.GET("/rule/packs", rule.GetAllPacks)
	g.GET("/rule/tools", rule.GetAllTools)
	g.GET("/rule/mounts", rule.GetAllMounts)
	g.GET("/rule/services", rule.GetAllServices)
	g.POST("/rule/shops", rule.CalcShop)
	// spell list
	g.GET("/rule/spelllist/:class/:level", rule.ListSpellByClass)
	g.GET("/rule/list/:kind", rule.ListContent)
	g.GET("/rule/list/:kind/:value", rule.ListContent)
	// description
	g.GET("/rule/description/:kind/:name/:subname", rule.GetDescription)
	g.GET("/rule/description/:kind/:name", rule.GetDescription)
	// return condition data
	g.GET("/rule/condition/:name", rule.GetCondition)
	g.GET("/rule/condition/:name/:level", rule.GetCondition)
	// treasure
	g.GET("/rule/randomtreasure/:level", rule.RandomTreasure)
	g.GET("/rule/averagetreasure/:level", rule.FastTreasure)
	g.GET("/rule/treasurehoard/:level", rule.RandomTreasureHoard)

	// campaign
	g.POST("/campaign", campaign.PostCampaign)
	g.GET("/campaign", campaign.GetAllCampaign)
	g.GET("/campaign/:id", campaign.GetOneCampaign)
	g.GET("/campaign/:id/:playerid", campaign.CheckPlayerAllowed)
	g.POST("/campaign/:id/player", campaign.AddPlayerCampaign)
	g.PUT("/campaign", campaign.PutCampaign)
	g.DELETE("/campaign/:id", campaign.DeleteCampaign)

	// adventure
	g.GET("/adventure", adventure.GetAllAdventure)
	g.GET("/adventure/:id", adventure.GetOneAdventure)
	g.POST("/adventure", adventure.PostAdventure)
	g.PUT("/adventure", adventure.PutAdventure)
	g.POST("/adventure/:id/encounter", adventure.AddEncounter)
	g.PUT("/adventure/:id/:status", adventure.ChangeAdventureStatus)
	g.DELETE("/adventure/:id", adventure.DeleteAdventure)

	// encounter
	g.GET("/encounter", encounter.GetAllEncounter)
	g.GET("/encounter/:id", encounter.GetOneEncounter)
	g.POST("/encounter", encounter.PostEncounter)
	g.PUT("/encounter", encounter.PutEncounter)
	g.POST("/encounter/npc", encounter.AddNPC)
	g.PUT("/encounter/:id/:status", encounter.ChangeEncounterStatus)
	g.DELETE("/encounter/:id", encounter.DeleteEncounter)

	// player
	g.GET("/player", playernpc.GetAllPlayers)
	g.GET("/player/:id", playernpc.GetOnePlayer)
	g.POST("/player", playernpc.PostPlayer)
	g.PUT("/player", playernpc.UpdateOnePlayer)
	g.POST("/player/:playerid/campaign", playernpc.AddCampaignToPlayer)
	g.PUT("/player/:playerid/hp/:action/:value", playernpc.AddOrRemoveHP)
	g.PUT("/player/:playerid/xp/:value", playernpc.AddPlayerXP)
	g.PUT("/player/:playerid/spell/:level/:value", playernpc.UseSpellByLevel)
	g.POST("/player/:playerid/condition", playernpc.ChangeCondition)
	g.PUT("/player/:playerid/fullrest", playernpc.FullRestPlayer)
	g.POST("/player/:playerid/armory", playernpc.AddArmorWeaponPlayerByID)
	g.POST("/player/:playerid/treasure", playernpc.AddTreasure)
	g.POST("/player/:playerid/items/:action", playernpc.AddOrRemoveOtherItems)
	g.POST("/player/:playerid/magicitems/:action", playernpc.AddOrRemoveMagicItems)
	g.DELETE("/player/:id", playernpc.DeletePlayer)

	// npc
	g.GET("/npc", playernpc.GetAllNPC)
	g.POST("/npc", playernpc.PostNPC) // not in postman collection
	g.PUT("/npc/:id/hp/:action/:value", playernpc.PostDamageNPC)
	g.POST("/npc/:id/condition", playernpc.ChangeNPCCondition)
	g.DELETE("/npc/:id", playernpc.DeleteNPC)

}
