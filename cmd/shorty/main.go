package main

import (
	"github.com/sancheschris/shorty/internal/config"
	"github.com/sancheschris/shorty/internal/database"
)

func main() {

	cfg := config.LoadConfig()

	database.InitDB(cfg.DBUrl)
}
