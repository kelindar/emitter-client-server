package main

import (
	"strings"

	emitter "github.com/emitter-io/go/v2"
)

const key = "iuERBYYtScnbp6YLagyHc8laQ94waUbc"

func main() {
	client, err := emitter.Connect("tcp://127.0.0.1:8080", func(c *emitter.Client, msg emitter.Message) {
		println("unknown message: ", msg.Topic(), string(msg.Payload()))
	})
	if err != nil {
		panic(err)
	}

	client.Subscribe(key, "demo/", onMessage)
	println("server started")

	for {
	}
}

func onMessage(c *emitter.Client, msg emitter.Message) {
	response := strings.ToUpper(string(msg.Payload()))
	c.Publish(key, msg.Topic(), response, emitter.WithoutEcho())
}
