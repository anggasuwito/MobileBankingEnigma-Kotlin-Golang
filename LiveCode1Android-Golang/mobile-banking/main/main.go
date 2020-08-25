package main

import (
	"banking/master"
	"banking/master/config"
)

func main() {
	db := config.Connection()
	router := config.CreateRouter()
	master.InitAll(router, db)
	config.RunServer(router)
}
