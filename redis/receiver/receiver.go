package main

import (
	"encoding/json"
	"fmt"
	"icssight/messaging/model"
	"sync"

	"github.com/go-redis/redis"
)

var wg sync.WaitGroup // 1
var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
}

func main() {
	wg.Add(1)
	go worker()
	wg.Wait()
}

func worker() {
	defer wg.Done()

	pubsub := client.Subscribe("message")
	defer pubsub.Close()

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	channel := pubsub.ChannelSize(1000000)

	message := model.Message{}
	for packet := range channel {
		json.Unmarshal([]byte(packet.Payload), &message)
		fmt.Println(message.Name)
	}

}
