package config

import (
	"os"
	"sync/atomic"
)

//Values stores the current configuration values
var Values Config

//Config contains the application's configuration values. Add here your own variables and bind it on init() function
type Config struct {
	//DBConnectionString to connect to Mongo
	DBConnectionString string
	//DBConnectionCertificateFileName defines the TLS Certificate for DB Connections. If not set, no TLS is configured
	DBConnectionCertificateFileName string
	//Port contains the port in which the application listens
	Port string
	//AppName for displaying in Monitoring
	AppName string
	//LogLevel - DEBUG or INFO or WARNING or ERROR or PANIC or FATAL
	LogLevel string
	//TestRun state if the current execution is a test execution
	TestRun bool
	//UsePrometheus to enable prometheus metrics endpoint
	UsePrometheus bool
	// IsReady atomic.Value
	IsReady atomic.Value
	// WaitTime int
	WaitTime int
	// MonstersDatabaseJSON string
	MonstersDatabaseJSON string
	// SpellsDatabaseJSON string
	SpellsDatabaseJSON string
	// SpellListDatabaseJSON string
	SpellListDatabaseJSON string
	// MagicItemsDatabaseJSON string
	MagicItemsDatabaseJSON string
	// ArmorsDatabaseJSON string
	ArmorsDatabaseJSON string
	// WeaponsDatabaseJSON string
	WeaponsDatabaseJSON string
	// GearDatabaseJSON string
	GearDatabaseJSON string
	// PacksDatabaseJSON string
	PacksDatabaseJSON string
	// ToolsDatabaseJSON string
	ToolsDatabaseJSON string
	// MountsDatabaseJSON string
	MountsDatabaseJSON string
	// HoardDatabaseJSON string
	HoardDatabaseJSON string
	// ServicesDatabaseJSON string
	ServicesDatabaseJSON string
}

// GetEnv gets an environment variable content or a default value
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	Values.TestRun = GetEnv("TESTRUN", "false") == "true"
	Values.UsePrometheus = GetEnv("USEPROMETHEUS", "false") == "true"
	Values.Port = GetEnv("PORT", "8080")
	Values.AppName = GetEnv("APP_NAME", "playbypost-dnd")
	Values.LogLevel = GetEnv("LOG_LEVEL", "INFO")
	Values.DBConnectionCertificateFileName = GetEnv("DB_CONNECTION_CERTIFICATE_FILE_NAME", "")
	Values.DBConnectionString = GetEnv("DB_CONNECTION_STRING", "")
	Values.WaitTime = 10

	// Load all json databases files
	Values.MonstersDatabaseJSON = GetEnv("MonstersDatabaseJSON", "./database/new-monster-list.json")
	Values.SpellsDatabaseJSON = GetEnv("SpellsDatabaseJSON", "./database/new-spell-description-list.json")
	Values.SpellListDatabaseJSON = GetEnv("SpellListDatabaseJSON", "./database/spell-list.json")
	Values.MagicItemsDatabaseJSON = GetEnv("MagicItemsDatabaseJSON", "./database/new-magic-itens-list.json")
	Values.ArmorsDatabaseJSON = GetEnv("ArmorsDatabaseJSON", "./database/new-armor-list.json")
	Values.WeaponsDatabaseJSON = GetEnv("WeaponsDatabaseJSON", "./database/new-weapon-list.json")
	Values.GearDatabaseJSON = GetEnv("GearDatabaseJSON", "./database/new-gear-list.json")
	Values.PacksDatabaseJSON = GetEnv("PacksDatabaseJSON", "./database/new-packs-list.json")
	Values.ToolsDatabaseJSON = GetEnv("ToolsDatabaseJSON", "./database/new-tools-list.json")
	Values.MountsDatabaseJSON = GetEnv("MountsDatabaseJSON", "./database/new-mounts-list.json")
	Values.HoardDatabaseJSON = GetEnv("HoardDatabaseJSON", "./database/new-treasure-hoard-list.json")
	Values.ServicesDatabaseJSON = GetEnv("ServicesDatabaseJSON", "./database/new-services-list.json")
}
