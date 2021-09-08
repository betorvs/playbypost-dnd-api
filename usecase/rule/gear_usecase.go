package rule

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/database"
	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

// GetAllWeapons returns a weapon by name
func GetAllWeapons(queryParameters url.Values) []rule.Weapon {
	db := database.GetDatabaseRepository()
	weapons := db.GetWeaponDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.Weapon
		for _, v := range weapons {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				case "kind":
					for _, n := range param {
						if strings.Contains(v.Kind, n) {
							filtered = append(filtered, v)
						}
					}
				case "damage_type", "type":
					for _, n := range param {
						if strings.Contains(v.DamageType, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		return filtered
	}
	return weapons
}

// WeaponsList returns a list weapons in []string
func WeaponsList() []string {
	db := database.GetDatabaseRepository()
	weapons := db.GetWeaponDatabase()
	weaponList := []string{}
	for _, v := range weapons {
		weaponList = append(weaponList, v.Name)
	}
	return weaponList
}

// WeaponsByName returns a monster by name
func WeaponsByName(name string) rule.Weapon {
	db := database.GetDatabaseRepository()
	weapons := db.GetWeaponDatabase()
	var weapon rule.Weapon
	for _, v := range weapons {
		if v.Name == name {
			weapon = v
		}
	}
	return weapon
}

// GetAllArmor returns a monster by name
func GetAllArmor(queryParameters url.Values) []rule.Armor {
	db := database.GetDatabaseRepository()
	armors := db.GetArmorDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.Armor
		for _, v := range armors {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				case "kind":
					for _, n := range param {
						if strings.Contains(v.Kind, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		return filtered
	}
	return armors
}

// ArmorList return a list of armors in []string
func ArmorList() []string {
	db := database.GetDatabaseRepository()
	armors := db.GetArmorDatabase()
	armorList := []string{}
	for _, v := range armors {
		armorList = append(armorList, v.Name)
	}
	return armorList
}

// ArmorByName returns a monster by name
func ArmorByName(name string) rule.Armor {
	db := database.GetDatabaseRepository()
	armors := db.GetArmorDatabase()
	var armor rule.Armor
	for _, v := range armors {
		if v.Name == name {
			armor = v
		}
	}
	return armor
}

// GetAllGears returns a monster by name
func GetAllGears(queryParameters url.Values) []rule.Gear {
	db := database.GetDatabaseRepository()
	gears := db.GetGearDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.Gear
		for _, v := range gears {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				case "kind":
					for _, n := range param {
						if strings.Contains(v.Kind, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		return filtered
	}
	return gears
}

// GearByName returns a monster by name
func GearByName(name string) rule.Gear {
	db := database.GetDatabaseRepository()
	gears := db.GetGearDatabase()
	var gear rule.Gear
	for _, v := range gears {
		if v.Name == name {
			gear = v
		}
	}
	return gear
}

// GetAllPacks returns a monster by name
func GetAllPacks(queryParameters url.Values) []rule.Packs {
	db := database.GetDatabaseRepository()
	packs := db.GetPacksDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.Packs
		for _, v := range packs {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		return filtered
	}
	return packs
}

// PacksByName returns a monster by name
func PacksByName(name string) rule.Packs {
	db := database.GetDatabaseRepository()
	packs := db.GetPacksDatabase()
	var pack rule.Packs
	for _, v := range packs {
		if v.Name == name {
			pack = v
		}
	}
	return pack
}

// PacksList return a list of armors in []string
func PacksList() []string {
	db := database.GetDatabaseRepository()
	packs := db.GetPacksDatabase()
	packsList := []string{}
	for _, v := range packs {
		packsList = append(packsList, v.Name)
	}
	return packsList
}

// GetAllTools returns a monster by name
func GetAllTools(queryParameters url.Values) []rule.Tools {
	db := database.GetDatabaseRepository()
	tools := db.GetToolsDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.Tools
		for _, v := range tools {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				case "kind":
					for _, n := range param {
						if strings.Contains(v.Kind, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		return filtered
	}
	return tools
}

// ToolsByName returns a monster by name
func ToolsByName(name string) rule.Tools {
	db := database.GetDatabaseRepository()
	tools := db.GetToolsDatabase()
	var tool rule.Tools
	for _, v := range tools {
		if v.Name == name {
			tool = v
		}
	}
	return tool
}

// ToolsList return a list of armors in []string
func ToolsList() []string {
	db := database.GetDatabaseRepository()
	tools := db.GetToolsDatabase()
	toolsList := []string{}
	for _, v := range tools {
		toolsList = append(toolsList, v.Name)
	}
	return toolsList
}

// GetAllMounts returns a monster by name
func GetAllMounts(queryParameters url.Values) []rule.Mounts {
	db := database.GetDatabaseRepository()
	mounts := db.GetMountsDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.Mounts
		for _, v := range mounts {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		return filtered
	}
	return mounts
}

// MountsByName returns a monster by name
func MountsByName(name string) rule.Mounts {
	db := database.GetDatabaseRepository()
	mounts := db.GetMountsDatabase()
	var mount rule.Mounts
	for _, v := range mounts {
		if v.Name == name {
			mount = v
		}
	}
	return mount
}

// MountsList return a list of armors in []string
func MountsList() []string {
	db := database.GetDatabaseRepository()
	mounts := db.GetMountsDatabase()
	mountsList := []string{}
	for _, v := range mounts {
		mountsList = append(mountsList, v.Name)
	}
	return mountsList
}

// CalcShoppingCart func
func CalcShoppingCart(cart *rule.SimpleList) *rule.ShoppingCart {
	shoppingCart := new(rule.ShoppingCart)
	shoppingCart.Items = []string{}
	shoppingCart.ServicesItems = []string{}
	shoppingCart.UnavailableItems = []string{}
	for _, v := range cart.List {
		if utils.StringInSlice(v, WeaponsList()) {
			weapon := WeaponsByName(v)
			shoppingCart.Items = append(shoppingCart.Items, v)
			shoppingCart = calcCoinExpenses(shoppingCart, weapon.Cost, weapon.CoinType)
			continue
		}
		if utils.StringInSlice(v, ArmorList()) {
			armor := ArmorByName(v)
			shoppingCart.Items = append(shoppingCart.Items, v)
			shoppingCart = calcCoinExpenses(shoppingCart, armor.Cost, armor.CoinType)
			continue
		}
		if utils.StringInSlice(v, PacksList()) {
			packs := PacksByName(v)
			shoppingCart.Items = append(shoppingCart.Items, v)
			shoppingCart = calcCoinExpenses(shoppingCart, packs.Cost, packs.CoinType)
			continue
		}
		if utils.StringInSlice(v, ToolsList()) {
			tools := ToolsByName(v)
			shoppingCart.Items = append(shoppingCart.Items, v)
			shoppingCart = calcCoinExpenses(shoppingCart, tools.Cost, tools.CoinType)
			continue
		}
		if utils.StringInSlice(v, MountsList()) {
			mounts := MountsByName(v)
			shoppingCart.Items = append(shoppingCart.Items, v)
			shoppingCart = calcCoinExpenses(shoppingCart, mounts.Cost, mounts.CoinType)
			continue
		}
		if utils.StringInSlice(v, ServicesNameList()) {
			services := ServiceByName(v)
			shoppingCart.ServicesItems = append(shoppingCart.ServicesItems, v)
			shoppingCart = calcCoinExpenses(shoppingCart, services.Cost, services.CoinType)
			continue
		}
		gear := GearByName(v)
		if gear.Name != "" {
			shoppingCart.Items = append(shoppingCart.Items, v)
			shoppingCart = calcCoinExpenses(shoppingCart, gear.Cost, gear.CoinType)
		} else {
			shoppingCart.UnavailableItems = append(shoppingCart.UnavailableItems, v)
		}

	}

	return shoppingCart
}

func calcCoinExpenses(shoppingCart *rule.ShoppingCart, value int, kind string) *rule.ShoppingCart {
	switch kind {
	case "platinum":
		shoppingCart.Cost.Platinum = shoppingCart.Cost.Platinum + value
	case "gold":
		shoppingCart.Cost.Gold = shoppingCart.Cost.Gold + value
	case "electrum":
		shoppingCart.Cost.Electrum = shoppingCart.Cost.Electrum + value
	case "silver":
		shoppingCart.Cost.Silver = shoppingCart.Cost.Silver + value
	case "copper":
		shoppingCart.Cost.Copper = shoppingCart.Cost.Copper + value
	}
	return shoppingCart
}

// CalcRandomTreasureByChallengeLevel returns in coins by challenge rate a monsters hoard
func CalcRandomTreasureByChallengeLevel(level int, random bool) (*rule.RandomTreasure, error) {
	treasure := new(rule.RandomTreasure)
	message, coins := individualPercentageByLevel(level, random)
	if message == "notFound" {
		return treasure, fmt.Errorf("error generating random treasure")
	}
	treasure.Challenge = level
	treasure.Message = message
	for k, v := range coins {
		switch k {
		case "platinum":
			treasure.Treasure.Platinum = v
		case "gold":
			treasure.Treasure.Gold = v
		case "electrum":
			treasure.Treasure.Electrum = v
		case "silver":
			treasure.Treasure.Silver = v
		case "copper":
			treasure.Treasure.Copper = v
		}
	}
	return treasure, nil
}

func individualPercentageByLevel(level int, random bool) (string, map[string]int) {
	treasure := make(map[string]int)
	r := diceroll.GetDice()

	res, t, _ := r.DiceRoll("1d100")
	msg := fmt.Sprintf("d100: %s ", t)
	switch {
	case level >= 1 && level <= 4:
		coin1 := checkTableBelowFive(res)
		if random {
			res1, t1, _ := r.DiceRoll(coin1.Dice)
			msg += fmt.Sprintf("%s: %s", coin1.Dice, t1)
			treasure[coin1.Type] = res1 * coin1.Multiple
			return msg, treasure
		}
		treasure[coin1.Type] = coin1.AverageValue
		return msg, treasure

	case level > 4 && level <= 10:
		// coin1, randomDice1, average1, coin2, randomDice2, average2 := checkTableBelowTen(res)
		coins := checkTableBelowTen(res)
		if random {
			for _, v := range coins {
				res1, t1, _ := r.DiceRoll(v.Dice)
				treasure[v.Type] = res1 * v.Multiple
				msg += fmt.Sprintf("%s: %s, ", v.Dice, t1)
			}
			return msg, treasure
		}
		for _, v := range coins {
			treasure[v.Type] = v.AverageValue
		}
		return msg, treasure

	case level > 10 && level <= 16:
		// coin1, randomDice1, average1, coin2, randomDice2, average2 := checkTableBelowSeventeen(res)
		coins := checkTableBelowSeventeen(res)
		if random {
			for _, v := range coins {
				res1, t1, _ := r.DiceRoll(v.Dice)
				treasure[v.Type] = res1 * v.Multiple
				msg += fmt.Sprintf("%s: %s, ", v.Dice, t1)
			}
			return msg, treasure
		}
		for _, v := range coins {
			treasure[v.Type] = v.AverageValue
		}
		return msg, treasure

	case level > 16 && level <= 20:
		// coin1, randomDice1, average1, coin2, randomDice2, average2 := checkTableAboveSeventeen(res)
		coins := checkTableAboveSeventeen(res)
		if random {
			for _, v := range coins {
				res1, t1, _ := r.DiceRoll(v.Dice)
				treasure[v.Type] = res1 * v.Multiple
				msg += fmt.Sprintf("%s: %s, ", v.Dice, t1)
			}
			return msg, treasure
		}
		for _, v := range coins {
			treasure[v.Type] = v.AverageValue
		}
		return msg, treasure
	}
	return "notFound", treasure
}

func checkTableBelowFive(value int) rule.CoinTable {
	var coinRes rule.CoinTable
	switch {
	case value <= 30:
		coinRes.Type = "copper"
		coinRes.AverageValue = 17
		coinRes.Dice = "5d6"
		coinRes.Multiple = 1
		return coinRes
	case value > 30 && value <= 60:
		coinRes.Type = "silver"
		coinRes.AverageValue = 14
		coinRes.Dice = "4d6"
		coinRes.Multiple = 1
		return coinRes
	case value > 60 && value <= 70:
		coinRes.Type = "electrum"
		coinRes.AverageValue = 10
		coinRes.Dice = "3d6"
		coinRes.Multiple = 1
		return coinRes
	case value > 70 && value <= 95:
		coinRes.Type = "gold"
		coinRes.AverageValue = 10
		coinRes.Dice = "3d6"
		coinRes.Multiple = 1
		return coinRes
	case value > 95 && value <= 100:
		coinRes.Type = "platinum"
		coinRes.AverageValue = 3
		coinRes.Dice = "1d6"
		coinRes.Multiple = 1
		return coinRes
	default:
		coinRes.Type = "copper"
		coinRes.AverageValue = 17
		coinRes.Dice = "5d6"
		coinRes.Multiple = 1
		return coinRes
	}
}

func checkTableBelowTen(value int) []rule.CoinTable {
	var sliceCoinRes []rule.CoinTable
	var coinRes1, coinRes2 rule.CoinTable
	switch {
	case value <= 30:
		coinRes1.Type = "copper"
		coinRes1.AverageValue = 1400
		coinRes1.Dice = "4d6"
		coinRes1.Multiple = 100
		coinRes2.Type = "electrum"
		coinRes2.AverageValue = 35
		coinRes2.Dice = "1d6"
		coinRes2.Multiple = 10
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 30 && value <= 60:
		coinRes1.Type = "silver"
		coinRes1.AverageValue = 210
		coinRes1.Dice = "6d6"
		coinRes1.Multiple = 10
		coinRes2.Type = "gold"
		coinRes2.AverageValue = 70
		coinRes2.Dice = "2d6"
		coinRes2.Multiple = 10
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 60 && value <= 70:
		coinRes1.Type = "electrum"
		coinRes1.AverageValue = 105
		coinRes1.Dice = "3d6"
		coinRes1.Multiple = 10
		coinRes2.Type = "gold"
		coinRes2.AverageValue = 70
		coinRes2.Dice = "2d6"
		coinRes2.Multiple = 10
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 70 && value <= 95:
		coinRes1.Type = "gold"
		coinRes1.AverageValue = 70
		coinRes1.Dice = "2d6"
		coinRes1.Multiple = 10
		coinRes2.Type = "platinum"
		coinRes2.AverageValue = 7
		coinRes2.Dice = "2d6"
		coinRes2.Multiple = 1
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 95 && value <= 100:
		coinRes1.Type = "gold"
		coinRes1.AverageValue = 70
		coinRes1.Dice = "2d6"
		coinRes1.Multiple = 10
		coinRes2.Type = "platinum"
		coinRes2.AverageValue = 10
		coinRes2.Dice = "3d6"
		coinRes2.Multiple = 1
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	default:
		coinRes1.Type = "copper"
		coinRes1.AverageValue = 1400
		coinRes1.Dice = "4d6"
		coinRes1.Multiple = 100
		coinRes2.Type = "electrum"
		coinRes2.AverageValue = 35
		coinRes2.Dice = "1d6"
		coinRes2.Multiple = 10
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	}
	// return "", "", 0, "", "", 0
}

func checkTableBelowSeventeen(value int) []rule.CoinTable {
	var sliceCoinRes []rule.CoinTable
	var coinRes1, coinRes2 rule.CoinTable
	switch {
	case value <= 20:
		coinRes1.Type = "silver"
		coinRes1.AverageValue = 1400
		coinRes1.Dice = "4d6"
		coinRes1.Multiple = 100
		coinRes2.Type = "gold"
		coinRes2.AverageValue = 350
		coinRes2.Dice = "1d6"
		coinRes2.Multiple = 100
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 21 && value <= 35:
		coinRes1.Type = "electrum"
		coinRes1.AverageValue = 350
		coinRes1.Dice = "1d6"
		coinRes1.Multiple = 100
		coinRes2.Type = "gold"
		coinRes2.AverageValue = 350
		coinRes2.Dice = "1d6"
		coinRes2.Multiple = 100
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 36 && value <= 75:
		coinRes1.Type = "gold"
		coinRes1.AverageValue = 700
		coinRes1.Dice = "2d6"
		coinRes1.Multiple = 100
		coinRes2.Type = "platinum"
		coinRes2.AverageValue = 35
		coinRes2.Dice = "1d6"
		coinRes2.Multiple = 10
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 76 && value <= 100:
		coinRes1.Type = "gold"
		coinRes1.AverageValue = 700
		coinRes1.Dice = "2d6"
		coinRes1.Multiple = 100
		coinRes2.Type = "platinum"
		coinRes2.AverageValue = 70
		coinRes2.Dice = "2d6"
		coinRes2.Multiple = 10
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	default:
		coinRes1.Type = "silver"
		coinRes1.AverageValue = 1400
		coinRes1.Dice = "4d6"
		coinRes1.Multiple = 100
		coinRes2.Type = "gold"
		coinRes2.AverageValue = 350
		coinRes2.Dice = "1d6"
		coinRes2.Multiple = 100
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	}
	// return "", "", 0, "", "", 0
}

func checkTableAboveSeventeen(value int) []rule.CoinTable {
	var sliceCoinRes []rule.CoinTable
	var coinRes1, coinRes2 rule.CoinTable
	switch {
	case value <= 15:
		coinRes1.Type = "electrum"
		coinRes1.AverageValue = 7000
		coinRes1.Dice = "2d6"
		coinRes1.Multiple = 1000
		coinRes2.Type = "gold"
		coinRes2.AverageValue = 2800
		coinRes2.Dice = "8d6"
		coinRes2.Multiple = 100
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 15 && value <= 55:
		coinRes1.Type = "gold"
		coinRes1.AverageValue = 3500
		coinRes1.Dice = "1d6"
		coinRes1.Multiple = 1000
		coinRes2.Type = "platinum"
		coinRes2.AverageValue = 350
		coinRes2.Dice = "1d6"
		coinRes2.Multiple = 100
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	case value > 56 && value <= 100:
		coinRes1.Type = "gold"
		coinRes1.AverageValue = 3500
		coinRes1.Dice = "1d6"
		coinRes1.Multiple = 1000
		coinRes2.Type = "platinum"
		coinRes2.AverageValue = 700
		coinRes2.Dice = "2d6"
		coinRes2.Multiple = 100
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	default:
		coinRes1.Type = "electrum"
		coinRes1.AverageValue = 7000
		coinRes1.Dice = "2d6"
		coinRes1.Multiple = 1000
		coinRes2.Type = "gold"
		coinRes2.AverageValue = 2800
		coinRes2.Dice = "8d6"
		coinRes2.Multiple = 100
		sliceCoinRes = append(sliceCoinRes, coinRes1, coinRes2)
		return sliceCoinRes
	}
	// return "", "", 0, "", "", 0
}

// CalcHoardPercentageByLevel returns HOARD for groups of monsters
func CalcHoardPercentageByLevel(level int) *rule.RandomTreasureHoard {
	treasure := new(rule.RandomTreasureHoard)
	treasure.Challenge = level
	treasure.ArtObjects = []string{}
	treasure.Gemstones = []string{}
	treasure.MagicItems = []string{}
	treasure.ArmoryItems = make(map[string]string)
	treasure.ChargeItems = make(map[string]int)
	r := diceroll.GetDice()

	res, t, _ := r.DiceRoll("1d100")
	msg := fmt.Sprintf("d100: %s ", t)
	switch {
	case level <= 4:
		copperDice := "6d6"
		res1, t1, _ := r.DiceRoll(copperDice)
		msg += fmt.Sprintf("%s: %s", copperDice, t1)
		treasure.Treasure.Copper = res1 * 100
		silverDice := "3d6"
		res2, t2, _ := r.DiceRoll(silverDice)
		msg += fmt.Sprintf("%s: %s", silverDice, t2)
		treasure.Treasure.Silver = res2 * 100
		goldDice := "2d6"
		res3, t3, _ := r.DiceRoll(goldDice)
		msg += fmt.Sprintf("%s: %s", goldDice, t3)
		treasure.Treasure.Gold = res3 * 10
		dice, kind, price, magicItemsRoll := checkHoardBelowFive(res)
		res4, t4, _ := r.DiceRoll(dice)
		msg += fmt.Sprintf("%s: %s", dice, t4)
		switch kind {
		case "gemstone":
			list := HoardNameList(kind, price)
			treasure.Gemstones = selectHoardFromSlice(list, res4)

		case "art-object":
			list := HoardNameList(kind, price)
			treasure.ArtObjects = selectHoardFromSlice(list, res4)
		}
		if len(magicItemsRoll) != 0 {
			for k, v := range magicItemsRoll {
				res5, t5, _ := r.DiceRoll(k)
				msg += fmt.Sprintf("%s: %s", k, t5)
				for i := 0; i <= res5; i++ {
					var itemName, itemType, itemShape string
					var itemValue int
					switch v {
					case "A":
						itemName, itemType, itemShape, itemValue = randomMagicItemTable("A")

					case "B":
						itemName, itemType, itemShape, itemValue = randomMagicItemTable("B")
					case "C":
						itemName, itemType, itemShape, itemValue = randomMagicItemTable("C")
					case "F":
						itemName, itemType, itemShape, itemValue = randomMagicItemTable("F")
					case "G":
						itemName, itemType, itemShape, itemValue = randomMagicItemTable("G")
					}
					switch itemType {
					case "item-with-charge":
						treasure.ChargeItems[itemName] = itemValue

					case "armory-item":
						treasure.ArmoryItems[itemName] = itemShape

					case "single-item":
						treasure.MagicItems = append(treasure.MagicItems, itemName)
					}

				}

			}
		}

		// case level > 4 && level <= 10:
		// case level > 10 && level <= 16:
		// case level > 16:
	}
	treasure.Message = msg

	return treasure
}

func selectHoardFromSlice(in []string, q int) (list []string) {
	for i := 0; i <= q; i++ {
		randonIndex := rand.Intn(len(in))
		list = append(list, in[randonIndex])
	}
	return list
}

func checkHoardBelowFive(value int) (string, string, int, map[string]string) {
	magicItemsRools := make(map[string]string)
	var dice, kind string
	var price int
	switch {
	case value >= 1 && value <= 6:
		dice = ""
		kind = ""
		price = 0
	case value >= 7 && value <= 16:
		dice = "2d6"
		price = 10
		kind = "gemstone"
	case value >= 17 && value <= 26:
		dice = "2d4"
		price = 25
		kind = "art-object"
	case value >= 27 && value <= 36:
		dice = "2d6"
		price = 50
		kind = "gemstone"
	case value >= 37 && value <= 44:
		dice = "2d6"
		price = 10
		kind = "gemstone"
		magicItemsRools["1d6"] = "A"
	case value >= 45 && value <= 52:
		dice = "2d4"
		price = 25
		kind = "art-object"
		magicItemsRools["1d6"] = "A"
	case value >= 53 && value <= 60:
		dice = "2d6"
		price = 50
		kind = "gemstone"
		magicItemsRools["1d6"] = "A"
	case value >= 61 && value <= 65:
		dice = "2d6"
		price = 10
		kind = "gemstone"
		magicItemsRools["1d4"] = "B"
	case value >= 66 && value <= 70:
		dice = "2d4"
		price = 25
		kind = "art-object"
		magicItemsRools["1d4"] = "B"
	case value >= 71 && value <= 75:
		dice = "2d6"
		price = 50
		kind = "gemstone"
		magicItemsRools["1d4"] = "B"
	case value >= 76 && value <= 78:
		dice = "2d6"
		price = 10
		kind = "gemstone"
		magicItemsRools["1d4"] = "C"
	case value >= 79 && value <= 80:
		dice = "2d4"
		price = 25
		kind = "art-object"
		magicItemsRools["1d4"] = "C"
	case value >= 81 && value <= 85:
		dice = "2d6"
		price = 50
		kind = "gemstone"
		magicItemsRools["1d4"] = "C"
	case value >= 86 && value <= 92:
		dice = "2d4"
		price = 25
		kind = "art-object"
		magicItemsRools["1d4"] = "F"
	case value >= 93 && value <= 97:
		dice = "2d6"
		price = 50
		kind = "gemstone"
		magicItemsRools["1d4"] = "F"
	case value >= 98 && value <= 99:
		dice = "2d4"
		price = 25
		kind = "art-object"
		magicItemsRools["1"] = "G"
	case value == 100:
		dice = "2d6"
		price = 50
		kind = "gemstone"
		magicItemsRools["1"] = "G"

	}
	return dice, kind, price, magicItemsRools
}
