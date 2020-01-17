package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"strings"
	"weather-push/model"
)

func SendEmail(config model.EmailConfig, user model.User, weather model.Weather) {
	m := gomail.NewMessage()

	// 发件人
	m.SetHeader("From", config.Username)

	// 收件人
	m.SetHeader("To", user.Email)

	// 邮件标题
	// 如果有雨，则在标题上进行改动
	title := fmt.Sprintf("%s，今日您所在地的天气情况：", user.Nickname)
	if strings.Contains(weather.Info, "雨") {
		title = fmt.Sprintf("有雨！有雨！有雨！%s，今日您所在地的天气情况：", user.Nickname)
	}
	m.SetHeader("Subject", title)

	// 邮件内容
	content := `
		%s，
		%s。
		最高温：%s℃，
		最低温：%s℃。
	`
	content = fmt.Sprintf(content, user.Address, weather.Info, weather.MaxTemp, weather.MinTemp)

	m.SetBody("text/plain", content)

	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatalln("发送邮件失败: ", err, user)
	} else {
		log.Printf("已为%s发出了天气提醒\n", user.Username)
	}
}