package main

import (
	"amarthaloan/config"
	"amarthaloan/db"
	"amarthaloan/routes"
)

func main() {
	conf := config.AppConfig()
	db.Init()
	e := routes.Init()
	e.Start(":" + conf.APP_PORT)
}
