package utils

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func SendEmail(mail string, code string) {
	message := `
		<p>Hello!</p>
		<p style="text-indent:2em">This is a your code.</p>
		<p style="text-indent:2em">%s</p>
	`
	host := "smtp.qq.com"
	port := 465
	userName := "1691401076@qq.com"
	password := "ibplgwgyjcktjfcb"

	m := gomail.NewMessage()
	m.SetHeader("From", userName)
	m.SetHeader("To", mail)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/html", fmt.Sprintf(message, code))

	d := gomail.NewDialer(host, port, userName, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// GenerateDigitCode 生成数字验证码
func GenerateDigitCode(length int) string {
	numbers := []byte("0123456789")
	code := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range code {
		code[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(code)
}
