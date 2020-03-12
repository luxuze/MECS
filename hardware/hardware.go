package main

import (
	"fmt"
	"log"
	"mecs/config"
	"mecs/mqtt"
	"reflect"
	"strings"
	"time"
)

type Hardware struct {
	MecsCode   string
	Version    string
	Type       string
	Body       string
	Proxy      string
	Local      string
	Registered string
	Timestamp  string
}

func (hw *Hardware) generateMECS() {
	log.Println("正在编译 MECS 编码信息 ...", hw)
	v := reflect.ValueOf(*hw)
	var values []string
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Name == "MecsCode" {
			continue
		}
		log.Println(v.Type().Field(i).Name, v.Field(i).Interface())
		values = append(values, fmt.Sprintf("%v", v.Field(i).Interface()))
	}
	hw.MecsCode = strings.Join(values, "\\")
	log.Printf("编译 MECS 编码成功: %v", hw.MecsCode)
}

func (hw Hardware) register() {
	log.Println("正在注册 MECS 编码")
	time.Sleep(time.Second * 1)
	log.Println("正在发送编码注册请求")
	mqtt.Publish(config.RegisterTopic, hw.MecsCode)
}

func (hw Hardware) Listen() {
	log.Println("正在启动设备主程序 ...", hw)
	mqtt.Subscribe(hw.MecsCode, func(msg string) {
		log.Printf("[%v]Received: %v", hw.MecsCode, msg)
		switch msg {
		case config.HardwareRegisterSucceed:
			hw.Registered = "1"
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
	log.Println("正在启动第 1 台设备 ...")
	Hardware{
		Version:    "V1",
		Type:       "Ecode",
		Body:       "1.0096.1.16532547.66554231554487",
		Proxy:      "https://luxuze.cn/proxy",
		Local:      "127.0.0.1:8081",
		Registered: "0",
		Timestamp:  fmt.Sprintf("%v", time.Now().Unix()),
	}.PowerOn()

	log.Println("正在启动第 2 台设备 ...")
	Hardware{
		Version:    "V1",
		Type:       "Ecode",
		Body:       "10008A11111E11",
		Proxy:      "https://luxuze.cn/proxy",
		Local:      "http://127.0.0.1:8082",
		Registered: "0",
		Timestamp:  fmt.Sprintf("%v", time.Now().Unix()),
	}.PowerOn()

	log.Println("正在启动第 3 台设备 ...")
	Hardware{
		Version:    "V1",
		Type:       "Ecode",
		Body:       "1.0096.1.16532547.66554231554488",
		Proxy:      "https://luxuze.cn/proxy",
		Local:      "http://127.0.0.1:8083",
		Registered: "0",
		Timestamp:  fmt.Sprintf("%v", time.Now().Unix()),
	}.PowerOn()

	log.Println("正在启动第 4 台设备 ...")
	Hardware{
		Version:    "V1",
		Type:       "Ucode",
		Body:       "0011000000000000100000101000101001110010000011000100000101000110.",
		Proxy:      "https://luxuze.cn/proxy",
		Local:      "http://127.0.0.1:8084",
		Registered: "0",
		Timestamp:  fmt.Sprintf("%v", time.Now().Unix()),
	}.PowerOn()

	log.Println("正在启动第 5 台设备 ...")
	Hardware{
		Version:    "V1",
		Type:       "Ucode",
		Body:       "10008A11111E11",
		Proxy:      "https://luxuze.cn/proxy",
		Local:      "http://127.0.0.1:8085",
		Registered: "0",
		Timestamp:  fmt.Sprintf("%v", time.Now().Unix()),
	}.PowerOn()

	select {}
}
