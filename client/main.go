package main

import (
	"fmt"
	"strings"

	emitter "github.com/emitter-io/go/v2"
)

func main() {
	client, err := emitter.Connect("tcp://127.0.0.1:8080", func(c *emitter.Client, msg emitter.Message) {
		println("received message: ", msg.Topic(), string(msg.Payload()))
	})
	if err != nil {
		panic(err)
	}

	// Create a private link
	if _, err := client.CreatePrivateLink("NcH3EKOpWP4l3eDL5I4nqnPhbVQRymei", "demo/", "s", true); err != nil {
		panic(err)
	}

	println("enter some text or 'q' to exit:")
	for {
		var text string
		if _, err := fmt.Scanln(&text); err != nil {
			println(err.Error())
			return
		}

		if strings.ToLower(text) == "q" {
			return
		}

		client.PublishWithLink("s", text)
	}
}
