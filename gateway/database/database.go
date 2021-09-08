package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/rule"
)

//Database struct
type Database struct {
	MonsterDatabase          []byte
	SpellDescriptionDatabase []byte
	SpellListByClass         []byte
	MagicItemDatabase        []byte
	ArmorDatabase            []byte
	WeaponDatabase           []byte
	GearDatabase             []byte
	PacksDatabase            []byte
	ToolsDatabase            []byte
	MountsDatabase           []byte
	HoardDatabase            []byte
	ServicesDatabase         []byte
}

func lazyDatabaseInit(file string) []byte {
	DB, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()

	byteValue, err := ioutil.ReadAll(DB)
	if err != nil {
		fmt.Println(err)
	}
	return byteValue
}

//GetMonsterDatabase gets the Database current implementation
func (db Database) GetMonsterDatabase() []rule.MonsterNPC {
	var monsters []rule.MonsterNPC

	err := json.Unmarshal(db.MonsterDatabase, &monsters)
	if err != nil {
		fmt.Println(err)
	}
	return monsters
}

//GetSpellListByClass gets the Database current implementation
func (db Database) GetSpellListByClass() rule.SpellListByClass {
	var spellList rule.SpellListByClass

	err := json.Unmarshal(db.SpellListByClass, &spellList)
	if err != nil {
		fmt.Println("3", err)
	}
	return spellList
}

//GetSpellDescriptionDatabase gets the Database current implementation
func (db Database) GetSpellDescriptionDatabase() []rule.SpellDescription {
	var spellList []rule.SpellDescription

	err := json.Unmarshal(db.SpellDescriptionDatabase, &spellList)
	if err != nil {
		fmt.Println(err)
	}

	return spellList
}

//GetMagicItemDatabase gets the Database current implementation
func (db Database) GetMagicItemDatabase() []rule.MagicItem {
	var magicItems []rule.MagicItem

	err := json.Unmarshal(db.MagicItemDatabase, &magicItems)
	if err != nil {
		fmt.Println(err)
	}

	return magicItems
}

//GetArmorDatabase gets the Database current implementation
func (db Database) GetArmorDatabase() []rule.Armor {
	var armors []rule.Armor

	err := json.Unmarshal(db.ArmorDatabase, &armors)
	if err != nil {
		fmt.Println(err)
	}

	return armors
}

//GetWeaponDatabase gets the Database current implementation
func (db Database) GetWeaponDatabase() []rule.Weapon {
	var weapons []rule.Weapon

	err := json.Unmarshal(db.WeaponDatabase, &weapons)
	if err != nil {
		fmt.Println(err)
	}
	unarmed := rule.Weapon{
		Name:       "unarmed",
		Kind:       "",
		Cost:       0,
		CoinType:   "",
		Damage:     "1",
		DamageType: "bludgeoning",
		Weight:     0,
		Measure:    "",
		Properties: "",
	}
	weapons = append(weapons, unarmed)

	return weapons
}

//GetGearDatabase gets the Database current implementation
func (db Database) GetGearDatabase() []rule.Gear {
	var gears []rule.Gear

	err := json.Unmarshal(db.GearDatabase, &gears)
	if err != nil {
		fmt.Println(err)
	}

	return gears
}

//GetPacksDatabase gets the Database current implementation
func (db Database) GetPacksDatabase() []rule.Packs {
	var packs []rule.Packs

	err := json.Unmarshal(db.PacksDatabase, &packs)
	if err != nil {
		fmt.Println(err)
	}

	return packs
}

//GetToolsDatabase gets the Database current implementation
func (db Database) GetToolsDatabase() []rule.Tools {
	var tools []rule.Tools

	err := json.Unmarshal(db.ToolsDatabase, &tools)
	if err != nil {
		fmt.Println(err)
	}

	return tools
}

//GetMountsDatabase gets the Database current implementation
func (db Database) GetMountsDatabase() []rule.Mounts {
	var mounts []rule.Mounts

	err := json.Unmarshal(db.MountsDatabase, &mounts)
	if err != nil {
		fmt.Println(err)
	}

	return mounts
}

//GetHoardDatabase gets the Database current implementation
func (db Database) GetHoardDatabase() []rule.TreasureHoard {
	var hoards []rule.TreasureHoard

	err := json.Unmarshal(db.HoardDatabase, &hoards)
	if err != nil {
		fmt.Println(err)
	}

	return hoards
}

//GetServicesDatabase gets the Database current implementation
func (db Database) GetServicesDatabase() []rule.Services {
	var hoards []rule.Services

	// dec := json.NewDecoder(bytes.NewReader(db.ServicesDatabase))
	// dec.DisallowUnknownFields()
	// if err := dec.Decode(&hoards); err != nil {
	// 	fmt.Println(err)
	// }
	err := json.Unmarshal(db.ServicesDatabase, &hoards)
	if err != nil {
		fmt.Println(err)
	}

	return hoards
}

func lazyMaster() appcontext.Component {
	return &Database{
		MonsterDatabase:          lazyDatabaseInit(config.Values.MonstersDatabaseJSON),
		SpellDescriptionDatabase: lazyDatabaseInit(config.Values.SpellsDatabaseJSON),
		SpellListByClass:         lazyDatabaseInit(config.Values.SpellListDatabaseJSON),
		MagicItemDatabase:        lazyDatabaseInit(config.Values.MagicItemsDatabaseJSON),
		ArmorDatabase:            lazyDatabaseInit(config.Values.ArmorsDatabaseJSON),
		WeaponDatabase:           lazyDatabaseInit(config.Values.WeaponsDatabaseJSON),
		GearDatabase:             lazyDatabaseInit(config.Values.GearDatabaseJSON),
		PacksDatabase:            lazyDatabaseInit(config.Values.PacksDatabaseJSON),
		ToolsDatabase:            lazyDatabaseInit(config.Values.ToolsDatabaseJSON),
		MountsDatabase:           lazyDatabaseInit(config.Values.MountsDatabaseJSON),
		HoardDatabase:            lazyDatabaseInit(config.Values.HoardDatabaseJSON),
		ServicesDatabase:         lazyDatabaseInit(config.Values.ServicesDatabaseJSON),
	}
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.Database, lazyMaster)
	logLocal := config.GetLogger()
	logLocal.Info("JSONs Database ready")

}
