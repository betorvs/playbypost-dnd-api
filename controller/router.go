package controller

import (
	"github.com/betorvs/playbypost-dnd/config"
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
}
