package register

import (
	"log"
	"mecs/config"
	"mecs/mqtt"
	"mecs/pkg/db"
)

func CodeRegisterHandler(msg string) {
	log.Println("收到硬件设备注册请求, mecs-code-register: ", msg)
	code := db.Code{Code: msg}
	if err := db.DB().Create(&code).Error; err != nil {
		log.Fatal(err)
	}
	mqtt.Publish(msg, config.HardwareRegisterSucceed)
}
