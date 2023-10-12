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
	output, result, err := mc.ImageSyncModeration(&moderation.ImageModerationInput{
		BusinessId: "business_id",
		Extra:      "GolangSDK-Test",
		Images: []moderation.Image{
			{ContentId: "golang-sdk-image", Data: "https://www.baidu.com/img/1.png", Type: 1},
		}})
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
