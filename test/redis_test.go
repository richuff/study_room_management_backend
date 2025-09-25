package test

import (
	"context"
	"fmt"
	"study_room_management_backend/mapper"
)

func testRedis() {
	ctx := context.Background()
	err := mapper.Rdb.Set(ctx, "name", "John Doe", 0).Err()
	if err != nil {
		fmt.Println("设置键失败:", err)
		return
	}
	val, err := mapper.Rdb.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println("获取值失败:", err)
		return
	}
	fmt.Println("Name:", val)
}
