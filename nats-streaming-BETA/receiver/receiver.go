package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

var wg sync.WaitGroup // 1
var nc *nats.Conn
var sc stan.Conn
var err error

// var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {

	// sc, err = stan.Connect("test-cluster", "pubID", stan.NatsConn(client))
	sc, err = stan.Connect("test-cluster", "subID")
	checkErr(err)

}

func main() {
	// Simple Synchronous Publisher
	sc.Publish("message", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	sub, err := sc.Subscribe("message", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	checkErr(err)

	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	sc.Close()
}

func DBWorker() {
	// defer wg.Done()

	// Simple Synchronous Publisher
	sc.Publish("message", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("message", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	sc.Close()

	// sub, err := sc.Subscribe("message", func(m *stan.Msg) {
	// 	fmt.Println("geldi")
	// 	message := model.Message{}
	// 	json.Unmarshal([]byte(m.Data), &message)
	// 	fmt.Println(message)
	// })
	// checkErr(err)

	// // Unsubscribe
	// sub.Unsubscribe()

	// // Close connection
	// sc.Close()

}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
