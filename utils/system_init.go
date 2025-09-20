package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"study_room_management_backend/mapper"
)

// InitConfig 初始化Config
func InitConfig() {
	viper.SetConfigName("application")
	viper.AddConfigPath("resource")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("初始化错误")
	}
}

// InitMysql 初始化Mysql
func InitMysql() {
	user := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")

	dns := user + ":" + password + "@(" + host + ":" + port + ")/" +
		database + "?charset=utf8&parseTime=True&loc=Local"
	err := mapper.InitMysql(dns)
	if err != nil {
		fmt.Println(err)
		return
	}
}
