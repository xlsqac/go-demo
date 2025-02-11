package main

import (
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	// defer nc.Close()
	nc.Publish("test-pub", []byte("hello1"))
	nc.Publish("test-pub", []byte("hello2"))
	nc.Publish("test-pub", []byte("hello3"))
	time.Sleep(1 * time.Second)

	// nc.Flush()
	//nc, _ := nats.Connect(nats.DefaultURL) // 默认地址：nats://127.0.0.1:4222
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer nc.Close()

	// 发布消息
	// subject := "testpub"
	// message := "Hello, NATS!"
	// err := nc.Publish(subject, []byte(message))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// message = "Hello, NATS! Group"
	// err = nc.Publish(subject, []byte(message))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Message published to subject %s: %s", subject, message)
}
