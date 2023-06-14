package main

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

func main() {
	mc, err := moderation.NewModerationClient("63632c7750689300010fe9aa", "5fe4d89d-5b22-11ed-b868-0242ac130015", "1733792053931827217", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, err := mc.TextModeration("我草泥马")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("output=== type:%s,score:%f", output.Type, output.Score)
}
