package main

import (
	"fmt"

	"github.com/Arismonx/van-booking-system/config"
	"github.com/Arismonx/van-booking-system/database"
	"github.com/Arismonx/van-booking-system/server"
)

func main() {
	conf := config.LoadConfig() // test config
	db := database.NewPostgresDatabase(conf.Database)
	server := server.NewServer(conf, db)
	server.Start()

	fmt.Println(conf)
}
