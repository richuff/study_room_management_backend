package mapper

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Open *gorm.DB
	Rdb  *redis.Client
)

func InitMysql(config string) (err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 彩色打印
		},
	)
	//连接数据库
	Open, err = gorm.Open(mysql.Open(config), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// InitRedis 初始化Redis
func InitRedis(addr string, password string, db int) {
	ctx := context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis服务器地址和端口
		Password: password, // Redis访问密码，如果没有可以为空字符串
		DB:       db,       // 使用的Redis数据库编号，默认为0
	})

	pong, err := Rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("连接Redis失败:", err)
		return
	}
	fmt.Println("成功连接到Redis:", pong)
}
