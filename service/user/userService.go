package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"study_room_management_backend/model"
	"study_room_management_backend/result"
	"study_room_management_backend/utils"
	"time"
)

// Register
// @Summary 用户注册
// @Tags 用户模块
// @Description 用户注册接口
// @Param email body string false "email"
// @Param password body string false "password"
// @Success 200 {string} json{"code"}
// @Router /api/user/register [post]
func Register(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		fmt.Println(err)
		return
	}

	user.Password = utils.MakePassword(user.Password)
	user.CreatedAt = time.Now()
	user.IsDelete = false

	success := model.CreateUser(user)
	if success {
		result.Ok(c, 1, nil)
	} else {
		result.Ok(c, 0, nil)
	}
}

// Login
// @Summary 用户登录
// @Tags 用户模块
// @Description 用户登录接口
// @Param email body string false "email"
// @Param password body string false "password"
// @Success 200 {string} json{"code"}
// @Router /api/user/login [post]
func Login(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		fmt.Println(err)
		return
	}

	user.Password = utils.ValidPassword(user.Password)
	success := model.GetUserByPassword(user.Email, user.Password)
	if success {
		result.Ok(c, 1, nil)
	} else {
		result.Ok(c, 0, nil)
	}
}

// Logoff
// @Summary 用户注销
// @Tags 用户模块
// @Description 用户注销接口
// @Param email body string false "email"
// @Success 200 {string} json{"code"}
// @Router /api/user/logoff [post]
func Logoff(c *gin.Context) {
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		fmt.Println(err)
		return
	}
	success := model.GetUserByEmail(user.Email)
	if success {
		user.IsDelete = true
		model.UpdateUser(user)

		result.Ok(c, 1, nil)
	} else {
		result.Ok(c, 0, nil)
	}
}
