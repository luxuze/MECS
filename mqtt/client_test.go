package mqtt

import (
	"fmt"
	"testing"
)

func TestSubscribe(t *testing.T) {
	type args struct {
		topic    string
		callback func(string)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name:"TestSubscribe",
			args:args{
				topic:"test",callback: func(msg string) {
					fmt.Println("this is callback function", msg)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Subscribe(tt.args.topic, tt.args.callback)
		})
	}
}

func TestPublish(t *testing.T) {
	type args struct {
		topic string
		msg   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestPublish",
			args:args{
				topic: "test",
				msg:   "hahaha",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Publish(tt.args.topic, tt.args.msg)
		})
	}
}