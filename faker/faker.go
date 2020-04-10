package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"redis-vs-nats/model"

	"github.com/bxcodec/faker"
)

func main() {

	datasets := map[string]int{
		"1k.json":   1000,
		"10k.json":  10000,
		"100k.json": 100000,
	}

	for filename, value := range datasets {
		message := model.Message{}
		messages := []model.Message{}

		// Added +1 to start with 1 not 0
		for i := 1; i < value+1; i++ {
			faker.FakeData(&message)
			message.ID = uint(i)
			messages = append(messages, message)
		}

		f, err := os.Create("../json/" + filename)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		json, _ := json.MarshalIndent(messages, "", "    ")
		err = ioutil.WriteFile("../json/"+filename, json, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

}
