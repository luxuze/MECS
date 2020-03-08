package mqtt

import (
	"fmt"
	mq "github.com/eclipse/paho.mqtt.golang"
	"log"
	"mecs/config"
	"time"
)

func connect( uri string) mq.Client {
	opts := createClientOptions( uri)
	client := mq.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(uri string) *mq.ClientOptions {
	opts := mq.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri))
	//opts.SetUsername(uri.User.Username())
	//password, _ := uri.User.Password()
	//opts.SetPassword(password)
	//opts.SetClientID(clientId)
	return opts
}

func Subscribe(topic string, callback func()) {
	log.Printf("Mqtt Subscribe: %v", topic)
	client := connect(config.MqttUrl)
	client.Subscribe(topic, 0, func(client mq.Client, msg mq.Message) {
		log.Printf("Mqtt Received: [%s] %s\n", msg.Topic(), string(msg.Payload()))
		callback()
	})
}

func Publish(topic string, msg string) {
	log.Printf("Mqtt Publish: %v, %v", topic, msg)
	client := connect(config.MqttUrl)
	client.Publish(topic, 0, false, msg)
}
