package main

import (
	"github.com/spf13/viper"
	"study_room_management_backend/router"
	"study_room_management_backend/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router()
	err := r.Run("localhost:" + viper.GetString("server.port"))
	if err != nil {
		return
	}
}
