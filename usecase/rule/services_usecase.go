package rule

import (
	"net/url"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/rule"
)

// GetAllServices returns all services
func GetAllServices(queryParameters url.Values) []rule.Services {
	db := rule.GetDatabaseRepository()
	services := db.GetServicesDatabase()
	if len(queryParameters) != 0 {
		var filtered []rule.Services
		for _, v := range services {
			for paramName, param := range queryParameters {
				switch paramName {
				case "name":
					for _, n := range param {
						if strings.Contains(v.Name, n) {
							filtered = append(filtered, v)
						}
					}
				case "source":
					for _, n := range param {
						if strings.Contains(v.Source, n) {
							filtered = append(filtered, v)
						}
					}
				}
			}
		}
		return filtered
	}
	return services
}

// ServiceByName returns a service by name
func ServiceByName(name string) rule.Services {
	db := rule.GetDatabaseRepository()
	services := db.GetServicesDatabase()
	var service rule.Services
	for _, v := range services {
		if strings.Contains(v.Name, name) {
			service = v
		}
	}
	return service
}

// ServicesNameList return a list of services based on word
func ServicesNameList() []string {
	db := rule.GetDatabaseRepository()
	services := db.GetServicesDatabase()
	servicesList := []string{}
	for _, v := range services {
		servicesList = append(servicesList, v.Name)
	}
	return servicesList
}
