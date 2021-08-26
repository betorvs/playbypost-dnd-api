package rule

import (
	"net/url"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/rule"
)

// GetAllHoard returns a []TreasureHoard
func GetAllHoard(queryParameters url.Values) []rule.TreasureHoard {
	db := rule.GetDatabaseRepository()
	hoards := db.GetHoardDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.TreasureHoard
		for _, v := range hoards {
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
	return hoards
}

// HoardByName returns a TreasureHoard by name
func HoardByName(name string) rule.TreasureHoard {
	db := rule.GetDatabaseRepository()
	hoards := db.GetHoardDatabase()
	var hoard rule.TreasureHoard
	for _, v := range hoards {
		if v.Name == name {
			hoard = v
		}
	}
	return hoard
}

// HoardNameList return a list of treasure hoards in []string by value
func HoardNameList(kind string, value int) []string {
	db := rule.GetDatabaseRepository()
	hoards := db.GetHoardDatabase()
	hoardsList := []string{}
	for _, v := range hoards {
		if v.Value == value {
			if v.Kind == kind {
				hoardsList = append(hoardsList, v.Name)
			}
		}
	}
	return hoardsList
}
