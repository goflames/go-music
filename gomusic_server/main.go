package main

import (
	"gomusic_server/config"
	"gomusic_server/router"
)

func main() {
	config.InitDB()
	config.InitMinio()
	config.InitRedis()
	r := router.InitRouter()
	r.Run(":8888")
}
