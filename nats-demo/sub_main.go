package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.QueueSubscribe("test-pub", "a", func(msg *nats.Msg) {
		log.Println("Queue[A]: ", string(msg.Data))
	})
	nc.QueueSubscribe("test-pub", "a", func(msg *nats.Msg) {
		log.Println("Queue[B]: ", string(msg.Data))
	})

	fmt.Println("listen")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// nc, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer nc.Close()

	// // 订阅主题
	// subject := "updates"
	// nc.Subscribe(subject, func(msg *nats.Msg) {
	// 	log.Printf("Received a message on subject %s: %s", msg.Subject, string(msg.Data))
	// })

	// // 保持连接，等待消息
	// select {}
}
