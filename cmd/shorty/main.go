package main

import (
	"log"

	"github.com/sancheschris/shorty/internal/config"
	"github.com/sancheschris/shorty/internal/database"
	"github.com/sancheschris/shorty/internal/router"
)

func main() {

	cfg := config.LoadConfig()
	database.InitDB(cfg.DBUrl)

	r := router.SetupRouter()
	log.Printf("Server running on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
