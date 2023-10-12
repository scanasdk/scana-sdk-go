package main

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

func main() {
	mc, err := moderation.New("app_id", "secret", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, result, err := mc.TextSyncModeration(&moderation.TextModerationInput{
		BusinessId: "business_id",
		Extra:      "GolangSDK-Test",
		Text: []moderation.Text{
			{ContentId: "golang-sdk", Data: "Hello World!"},
		}})
	if err != nil {
		if result != nil {
			log.Printf("code:%d,message:%s", result.Code, result.Msg)
			return
		}
		log.Println(err)
		return
	}

	log.Printf("output=== type:%v", output)
}
