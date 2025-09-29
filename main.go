package main

import (
	"github.com/spf13/viper"
	"study_room_management_backend/config"
	"study_room_management_backend/router"
	"study_room_management_backend/utils"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	r := router.Router()

	r.Static("/static", config.C.Storage.Local.Path)

	err := r.Run(viper.GetString("server.port"))
	if err != nil {
		return
	}
}
