package main

import (
	"log"
	"mecs/config"
	"mecs/mqtt"
	"mecs/proxy/command"
	"mecs/proxy/register"
)

func main() {
	log.Println("启动 MECS 编码注册服务端程序 ...")
	go mqtt.Subscribe(config.RegisterTopic, register.CodeRegisterHandler)

	log.Println("启动用户指令处理服务 ...")
	go command.WaitingUserCommand()
	select {}
}
