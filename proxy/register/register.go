package register

import (
	"log"
	"mecs/config"
	"mecs/mqtt"
)

func CodeRegisterHandler(msg string) {
	log.Println("mecs-code-register: ", msg)
	mqtt.Publish(msg, config.HardwareRegisterSucceed)
}
