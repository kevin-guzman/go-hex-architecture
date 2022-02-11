package utils

import (
	"encoding/json"
	"log"
)

func JSONParse(data, v interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error parsing data", err)
	}
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		log.Fatal("Error parsing data", err)
	}
}
