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
	output, result, err := mc.DocAsyncModeration(&moderation.DocModerationInput{
		BusinessId: "business_id",
		Extra:      "GolangSDK-Test",
		ContentId:  "golang-sdk-doc",
		URL:        "https://www.baidu.com/a.pdf",
	})
	if err != nil {
		if result != nil {
			log.Printf("code:%d,message:%s", result.Code, result.Msg)
			return
		}
		log.Println(err)
		return
	}

	log.Printf("output:%v", output)
}
