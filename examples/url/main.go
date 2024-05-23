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
	output, result, err := mc.URLAsyncModeration(&moderation.URLModerationInput{
		BusinessId: "1813481786970231031",
		Extra:      "GolangSDK-Test",
		ContentId:  "golang-sdk-video",
		URL:        "https://www.baidu.com/2.m3u8",
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
