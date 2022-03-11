package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	App         string
	Environment string // dev,test, prod

	ServiceHost string
	HTTPPort    string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	DefaultOffset string
	DefaultLimit  string
}

// Load ...
func Load() Config{
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("PROJECT_NAME","go_boilerplate"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT","dev"))

	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST","localhsot"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT",":8080"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST","localhost"))
	config.PostgresPort = cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT",5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER","postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD","your_db_password"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE","your_db_database"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET","0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT","10"))

	return config
}

func getOrReturnDefaultValue(key string,defaultValue interface{})  interface{}{
	val,exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}