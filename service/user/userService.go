package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"mime/multipart"
	"path/filepath"
	"strconv"
	jwtUtil "study_room_management_backend/jwt"
	"study_room_management_backend/mapper"
	"study_room_management_backend/model"
	"study_room_management_backend/model/dto"
	"study_room_management_backend/model/vo"
	"study_room_management_backend/result"
	"study_room_management_backend/service/filer"
	"study_room_management_backend/utils"
	"time"
)

// Register
// @Summary 用户注册
// @Tags 用户模块
// @Description 用户注册接口
// @Accept    json
// @Produce   json
// @Param     req body dto.UserRegisterDto true "注册信息"
// @Success   200 {object} result.CodeResp "业务代码"
// @Router    /api/user/register [post]
func Register(c *gin.Context) {
	userRegisterDto := &dto.UserRegisterDto{}
	user := &model.User{}
	if err := c.BindJSON(userRegisterDto); err != nil {
		fmt.Println(err)
		return
	}
	user = model.GetUserByUserName(userRegisterDto.Name)
	if user.Name != "" {
		result.ErrorWithCode(c, "该用户名已被使用", 0)
		return
	}
	user.Email = userRegisterDto.Email
	user.Name = userRegisterDto.Name
	ctx := context.Background()

	code, err := mapper.Rdb.Get(ctx, userRegisterDto.Email).Result()
	if err != nil {
		fmt.Println(err)
	}
	if code != userRegisterDto.Code {
		result.ErrorWithCode(c, "验证码错误", 0)
		return
	}

	user.Password = utils.MakePassword(userRegisterDto.Password)
	user.IsDelete = false
	user.Avatar = "/public/default"

	success := model.CreateUser(user)
	if success {
		expirationTime := time.Now().Add(90 * time.Minute)
		claims := &jwtUtil.Claims{
			Username: user.Name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtUtil.JwtKey)
		if err != nil {
			fmt.Println(err)
		}
		result.Ok(c, 1, tokenString)
	} else {
		result.ErrorWithCode(c, "该邮箱已被占用", 0)
	}
}

// Login
// @Summary 用户登录
// @Tags 用户模块
// @Description 用户登录接口
// @Accept    json
// @Produce   json
// @Param     req body dto.UserLoginDto true "登录信息"
// @Success   200 {object} result.CodeResp "业务代码"
// @Router /api/user/login [post]
func Login(c *gin.Context) {
	userLoginDto := &dto.UserLoginDto{}
	if err := c.BindJSON(userLoginDto); err != nil {
		fmt.Println(err)
		return
	}

	userLoginDto.Password = utils.ValidPassword(userLoginDto.Password)
	success, message := model.GetMessageByPassword(userLoginDto.Email, userLoginDto.Password)
	if success {
		user := model.GetUserByPassword(userLoginDto.Email, userLoginDto.Password)

		expirationTime := time.Now().Add(90 * time.Minute)
		claims := &jwtUtil.Claims{
			Username: user.Name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtUtil.JwtKey)
		if err != nil {
			fmt.Println(err)
		}
		result.Ok(c, 1, tokenString)
	} else {
		result.ErrorWithCode(c, message, 0)
	}
}

// Logoff
// @Summary 用户注销
// @Tags 用户模块
// @Description 用户注销接口
// @Param user_id query int true "用户id"
// @Success   200 {object} result.CodeResp "业务代码"
// @Router /api/user/logoff [get]
func Logoff(c *gin.Context) {
	//tokenStr := c.GetHeader("Authorization")
	//claims := &jwtUtil.Claims{}
	//token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
	//	return jwtUtil.JwtKey, nil
	//})
	//
	//if err != nil || !token.Valid {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//	return
	//}

	userId, err := strconv.ParseUint(c.Request.FormValue("user_id"), 10, 64)
	if utils.ErrHandler(c, err) {
		return
	}
	user := model.GetUserByUserID(userId)
	if user.Name != "" {
		user.IsDelete = true
		model.UpdateUser(&user)

		result.Ok(c, 1, nil)
	} else {
		result.ErrorWithCode(c, "用户不存在", 0)
	}
}

// CheckInfo
// @Summary 查看用户信息
// @Tags 用户模块
// @Description 查看用户信息接口
// @Param user_id query int true "用户id"
// @Success   200 {object} result.CodeResp "业务代码"
// @Router /api/user/checkInfo [get]
func CheckInfo(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Request.FormValue("user_id"), 10, 64)
	if utils.ErrHandler(c, err) {
		return
	}

	user := model.GetUserByUserID(userId)
	student := model.GetStudentByUserID(userId)

	studentInfoVo := vo.StudentInfoVo{}
	if user.Name != "" {
		err := utils.SimpleCopyProperties(&studentInfoVo, user)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = utils.SimpleCopyProperties(&studentInfoVo, student)
		if err != nil {
			fmt.Println(err)
			return
		}
		result.Ok(c, 1, studentInfoVo)
	} else {
		result.ErrorWithCode(c, "用户不存在", 0)
	}
}

// SetInfo
// @Summary 设置用户信息
// @Tags 用户模块
// @Description 设置用户信息接口
// @Accept    json
// @Produce   json
// @Param     req body dto.UserInfoDto true "登录信息"
// @Success   200 {object} result.CodeResp "业务代码"
// @Router /api/user/setInfo [post]
func SetInfo(c *gin.Context) {
	userInfoDto := &dto.UserInfoDto{}
	if err := c.BindJSON(userInfoDto); err != nil {
		fmt.Println(err)
		return
	}
	model.InsertStudentInfo(userInfoDto)
	result.Ok(c, 1, "更新成功")
}

type UploadResp struct {
	URL string `json:"url"`
}

// SetAvatar
// @Summary 设置用户头像
// @Tags 用户模块
// @Description 设置用户头像接口
// @Accept  multipart/form-data
// @Param   user_id formData string true "用户id"
// @Param   avatar formData file true "头像"
// @Security Bearer
// @Success 200 {object} UploadResp
// @Router /api/user/setAvatar [post]
func SetAvatar(c *gin.Context) {
	file, header, err := c.Request.FormFile("avatar")
	if err != nil {
		result.ErrorWithCode(c, "缺少文件", 0)
		return
	}

	userId, err := strconv.ParseUint(c.Request.FormValue("user_id"), 10, 64)
	fmt.Println(userId)
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	// 简单校验
	ext := filepath.Ext(header.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(400, gin.H{"msg": "仅支持 jpg/png"})
		return
	}

	// 生成对象 key：avatar/10001_1628051212.png
	uid := c.GetInt64("uid") // 从 jwt 中间件拿
	key := filer.GenObjectKey("avatar", uid, ext)

	url, err := filer.SaveLocal(file, key)
	success := model.UpdateAvatar(userId, url)
	if !success {
		result.ErrorWithCode(c, "该用户不存在或已注销", 0)
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	result.Ok(c, 1, UploadResp{URL: url})
}
