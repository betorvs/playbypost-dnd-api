package database

import (
	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
)

// DatabaseRepository interface
type DatabaseRepository interface {
	appcontext.Component
	//GetMonsterDatabase returns monster json database
	GetMonsterDatabase() []rule.MonsterNPC
	// GetSpellListByClass return list of spell by class
	GetSpellListByClass() rule.SpellListByClass
	//GetSpellDescriptionDatabase returns spell descriptions json database
	GetSpellDescriptionDatabase() []rule.SpellDescription
	// GetMagicItemDatabase returns magic items json database
	GetMagicItemDatabase() []rule.MagicItem
	// GetArmorDatabase returns armors json database
	GetArmorDatabase() []rule.Armor
	// GetWeaponDatabase returns weapons json database
	GetWeaponDatabase() []rule.Weapon
	// GetGearDatabase returns advantures gear to shop
	GetGearDatabase() []rule.Gear
	// GetPacksDatabase returns adventures gear packs to shop
	GetPacksDatabase() []rule.Packs
	// GetToolsDatabase returns Tools database to shop
	GetToolsDatabase() []rule.Tools
	// GetMountsDatabase returns Mounts database to shop
	GetMountsDatabase() []rule.Mounts
	// GetHoardDatabase returns TreasureHoard database to create random treasure hoards
	GetHoardDatabase() []rule.TreasureHoard
	// GetServicesDatabase returns services database
	GetServicesDatabase() []rule.Services
}

// GetDatabaseRepository func return DatabaseRepository interface
func GetDatabaseRepository() DatabaseRepository {
	return appcontext.Current.Get(appcontext.Database).(DatabaseRepository)
}
