package corn

import (
	"github.com/robfig/cron/v3"
	"os"
	"strconv"
	"time"
	"weather-push/api"
	"weather-push/mail"
	"weather-push/model"
	"weather-push/util"
)

// 开启定时任务
func Start() {
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	cJob := cron.New(cron.WithLocation(nyc))
	cronCfg := os.Getenv("Cron")

	// 添加定时任务
	_, err := cJob.AddFunc(cronCfg, func() {
		dispatch()
	})
	if err != nil {
		util.Log().Error("启动定时任务出错", err)
	}
	util.Log().Info("定时任务已开启成功: ", cronCfg)
	cJob.Start()
}

func dispatch() {
	users := model.GetUsers()
	// 遍历每一个用户
	for _, user := range users {
		weather := api.QueryWithCityName(user.Address)
		sendMail(user, weather)
	}
}

func sendMail(user model.User,  weather model.Weather) {
	port, _ := strconv.Atoi(os.Getenv("Port"))
	emailConfig := model.EmailConfig{
		Host:     os.Getenv("Host"),
		Port:     port,
		Username: os.Getenv("Email_Username"),
		Password: os.Getenv("Email_Password"),
	}

	mail.SendEmail(emailConfig, user, weather)
}