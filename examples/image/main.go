package main

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

func main() {
	mc, err := moderation.New("62ac786311b92177337a933b", "a2f175a8-5373-11ed-9949-0242ac12000e", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, result, err := mc.ImageSyncModeration(&moderation.ImageModerationInput{
		BusinessId: "1800097935845781514",
		Extra:      "GolangSDK-Test",
		Images: []moderation.Image{
			{ContentId: "golang-sdk-image", Data: "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png", Type: 1},
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
