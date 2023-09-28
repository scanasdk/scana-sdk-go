package main

import (
	"github.com/scanasdk/scana-sdk-go/moderation"
	"log"
)

func main() {
	mc, err := moderation.New("62ac786311b92177337a933b", "a2f175a8-5373-11ed-9949-0242ac12000e", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, result, err := mc.TextSyncModeration(&moderation.TextModerationInput{
		BusinessId: "1720056633192620049",
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
