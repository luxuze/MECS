package root

import (
	"log"

	"github.com/monaco-io/request"
)

func SendUserCmd(msg []byte) {
	client := request.Client{
		URL:         "https://luxuze.cn/mecs/cmd",
		Method:      "POST",
		ContentType: request.ApplicationJSON, // default is "application/json"
		Body:        msg,
	}
	resp, err := client.Do()
	log.Println(string(resp), err)
}
