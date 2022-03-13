package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"gitlab.com/Udevs/Article/api"
	_ "gitlab.com/Udevs/Article/api/docs"
	"gitlab.com/Udevs/Article/api/handlers"
	"gitlab.com/Udevs/Article/config"
	"gitlab.com/Udevs/Article/storage/postgres"
)

func main() {
	cfg := config.Load()
	pgStore := postgres.NewPostgres(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	))

	h := handlers.NewHandler(pgStore)

	switch cfg.Environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api.SetUpAPI(r, h, cfg)

	r.Run(":8080")
}
