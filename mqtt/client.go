package main

import (
	"fmt"
	"log"
	"time"

	mq "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId string, uri string) mq.Client {
	opts := createClientOptions(clientId, uri)
	client := mq.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri string) *mq.ClientOptions {
	opts := mq.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri))
	//opts.SetUsername(uri.User.Username())
	//password, _ := uri.User.Password()
	//opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func listen(uri string, topic string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 0, func(client mq.Client, msg mq.Message) {
		fmt.Printf("Received [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {
	uri := "127.0.0.1:1883"
	topic := "test"

	go listen(uri, topic)

	client := connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, t.String())
	}
}
