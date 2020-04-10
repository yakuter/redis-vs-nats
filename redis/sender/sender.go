package main

import (
	"encoding/json"
	"icssight/messaging/model"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var filename string = "../../json/100k.json"
var client *redis.Client
var err error

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
}

func main() {

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

		AddRedis(messageJSON)
	}

	// Close the file
	_, err = decoder.Token()
	checkErr(err)

	elapsed := time.Since(start)
	log.Println("Redis Sender took %s", elapsed)

}

/*AddRedis publish incident data to insight channel*/
func AddRedis(data []byte) {
	_, err = client.Do("PUBLISH", "message", string(data)).Result()
	if err != nil {
		log.Println(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
