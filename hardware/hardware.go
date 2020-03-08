package main

import (
	"log"
	"mecs/config"
	"mecs/mqtt"
	"time"
)

type Hardware struct {
	MecsCode     string
	IsRegistered bool
}

func (hw Hardware) generateMECS() {
	log.Println("正在编译 MECS 编码信息 ...")
	time.Sleep(time.Second * 3)
	log.Printf("编译 MECS 编码成功: %v", hw.MecsCode)
}

func (hw Hardware) register() {
	log.Println("正在注册 MECS 编码")
	time.Sleep(time.Second * 1)
	log.Println("正在发送编码注册请求")
	mqtt.Publish(config.RegisterTopic, hw.MecsCode)
}

func (hw Hardware) Listen() {
	log.Println("正在启动设备主程序 ...")
	mqtt.Subscribe(hw.MecsCode, func(msg string) {
		log.Printf("[%v]Received: %v", hw.MecsCode, msg)
		switch msg {
		case config.HardwareRegisterSucceed:
			hw.IsRegistered = true
			log.Println("MECS 编码注册成功 !")
		default:
			log.Printf("无效指令: %v", msg)
		}
	})
}

func (hw Hardware) PowerOn() {
	hw.generateMECS()
	hw.register()
	hw.Listen()
}

func main() {
	hw := Hardware{
		MecsCode:     "mecs-test-1",
		IsRegistered: false,
	}
	hw.PowerOn()
	select {}
}
