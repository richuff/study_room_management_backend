package code

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	base64Captcha "github.com/mojocn/base64Captcha"
	"image/color"
	"study_room_management_backend/mapper"
	"study_room_management_backend/model/vo"
	"study_room_management_backend/result"
	"study_room_management_backend/utils"
	"time"
)

var store = base64Captcha.DefaultMemStore

// Captcha
// @Summary 获取验证码
// @Tags 测试模块
// @Description 获取验证码接口
// @Success   200 {object} result.CodeResp "业务代码"
// @Router /test/captcha [get]
func Captcha(c *gin.Context) {
	// 配置验证码驱动
	driver := &base64Captcha.DriverString{
		Height:          60,                                                                   // 高度
		Width:           200,                                                                  // 宽度
		NoiseCount:      5,                                                                    // 干扰线数量
		ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine, // 干扰线样式
		Length:          6,                                                                    // 验证码长度
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM",     // 字符源
		BgColor: &color.RGBA{ // 背景色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"}, // 字体文件
	}

	// 创建验证码实例
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	id, b64s, answer, err := captcha.Generate()
	fmt.Println(answer)
	if err != nil {
		result.Error(c, "验证码生成失败！")
		return
	}
	codeVo := vo.CodeVo{IdKey: id, Image: b64s}
	result.Ok(c, 1, codeVo)
}

// Verify
// @Summary 验证验证码
// @Tags 测试模块
// @Description 验证验验证码接口
// @Param id  query string true "验证码ID"
// @Param code  query string true "用户输入的验证码"
// @Success   200 {object} result.CodeResp "业务代码"
// @Router /test/captcha/verify [get]
func Verify(c *gin.Context) {
	// 获取请求参数
	id := c.Query("id")     // 验证码ID
	code := c.Query("code") // 用户输入的验证码

	// 调用存储验证方法
	if store.Verify(id, code, true) { // 第三个参数表示验证后删除
		result.Ok(c, 1, "验证成功")
	} else {
		result.ErrorWithCode(c, "验证码错误", 0)
	}
}

// SendSmsCode 存储并发送验证码
// @Summary 存储并发送验证码
// @Tags 用户模块
// @Description 存储并发送验证码接口
// @Param email  query string true "邮箱"
// @Success   200 {object} result.CodeResp "业务代码"
// @Router /api/captcha [get]
func SendSmsCode(c *gin.Context) {
	ctx := context.Background()
	email := c.Query("email") // 验证码ID
	code := utils.GenerateDigitCode(6)
	// 存储到Redis，设置5分钟过期
	err := mapper.Rdb.Set(ctx, email, code, 5*time.Minute)
	if err != nil {
		fmt.Println(err)
	}
	// 调用短信服务发送验证码
	utils.SendEmail(email, code)
	result.Ok(c, 1, "已发送验证码")
}

// VerifySmsCode 验证验证码
// @Summary 验证验证码
// @Tags 用户模块
// @Description 验证验验证码接口
// @Param email  query string true "用户邮箱"
// @Param ucode  query string true "用户输入的验证码"
// @Success   200 {object} result.CodeResp "业务代码"
// @Router /api/captcha/verify [get]
func VerifySmsCode(c *gin.Context) {
	ctx := context.Background()

	email := c.Query("email")
	ucode := c.Query("ucode")

	code, err := mapper.Rdb.Get(ctx, email).Result()
	if err != nil {
		fmt.Println(err)
	}
	if code != ucode {
		result.Error(c, "验证码错误")
	} else {
		result.Ok(c, 1, "验证成功")
	}
}
