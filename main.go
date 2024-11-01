package main

import (
	"bubble/config"
	"bubble/routers"
)

func main() {
	config.InitConfig()
	r := routers.SetupRouter()
	r.Run(config.AppConfig.APP.Port)
}
