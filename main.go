package main

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

func main() {
	mc, err := moderation.NewModerationClient("appid", "secret", "businessId", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, err := mc.TextModeration("Hello World")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("output=== type:%s,score:%f", output.Type, output.Score)
}
