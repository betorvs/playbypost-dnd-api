package rule

import (
	"math/rand"
	"net/url"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/diceroll"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/utils"
)

//GetMagicItem func
func GetMagicItem(queryParameters url.Values) []rule.MagicItem {
	db := rule.GetDatabaseRepository()
	magicItems := db.GetMagicItemDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.MagicItem
		for _, v := range magicItems {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				case "title":
					for _, n := range param {
						if strings.Contains(v.Title, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		// fmt.Println(len(filtered))
		return filtered
	}
	return magicItems
}

// GetMagicItemByName return a magic item by name
func GetMagicItemByName(name string) (item rule.MagicItem) {
	db := rule.GetDatabaseRepository()
	magicItems := db.GetMagicItemDatabase()
	for _, v := range magicItems {
		if name == v.Name {
			item = v
		}
	}
	return item
}

func getMagicItemByHoardTable(table string) (items []rule.MagicItem) {
	db := rule.GetDatabaseRepository()
	magicItems := db.GetMagicItemDatabase()
	for _, v := range magicItems {
		if utils.StringInSlice(table, v.HoardTable) {
			items = append(items, v)
		}
	}
	return items
}

func randomMagicItemTable(value string) (string, string, string, int) {
	items := getMagicItemByHoardTable(value)
	randonIndex := rand.Intn(len(items))
	item := items[randonIndex]
	r := diceroll.GetDice()
	if item.Power != nil && item.Power.ChargeType {
		if item.Power.Charges == 0 && item.Power.DiceCharges != "" {
			res1, _, _ := r.DiceRoll(item.Power.DiceCharges)
			return item.Name, "item-with-charge", "", res1
		}
		return item.Name, "item-with-charge", "", item.Power.Charges
	}
	if item.Category == "weapons" || item.Category == "armor" {
		return item.Name, "armory-item", item.Shape, 0
	}
	return item.Name, "single-item", "", 0
}
