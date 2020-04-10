package main

import (
	"encoding/json"
	"icssight/messaging/model"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

var filename string = "../../json/100bin.json"
var client *nats.Conn
var sc stan.Conn
var err error

func init() {
	client, err = nats.Connect(nats.DefaultURL)
	checkErr(err)

	sc, err = stan.Connect("test-cluster", "pubID", stan.NatsConn(client))
	checkErr(err)
}

func main() {
	defer client.Close()
	defer sc.Close()

	start := time.Now()

	jsonFile, err := os.Open(filename)
	checkErr(err)

	decoder := json.NewDecoder(jsonFile)

	// Read opening file
	_, err = decoder.Token()
	checkErr(err)

	var message model.Message
	for decoder.More() {
		err := decoder.Decode(&message)
		checkErr(err)

		messageJSON, err := json.Marshal(message)
		checkErr(err)

		AddNats(messageJSON)
	}

	// Close the file
	_, err = decoder.Token()
	checkErr(err)

	elapsed := time.Since(start)
	log.Println("Asset Nats Parser took %s", elapsed)

}

func AddNats(data []byte) {
	err = sc.Publish("message", data)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
