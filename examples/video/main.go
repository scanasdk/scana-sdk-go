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
	output, result, err := mc.VideoAsyncModeration(&moderation.VideoModerationInput{
		BusinessId: "1813481786970231031",
		Extra:      "GolangSDK-Test",
		ContentId:  "golang-sdk-video",
		URL:        "https://rs-videos.zsxq.com/a5k24jk3fvxzqxszla24lp07desj5tg2.m3u8?MtsHlsUriToken=ZXQ9MTcwMTA2NzQ2MzA0NSZlbmNyeXB0ZWQ9MCZ1aWQ9MCZ2aWQ9ODQ5Nzcmc2lnbj0zNjg5QTI5QzcxODA4N0FFMkNERUIwQkI0NzA2MkFEQw%3D%3D",
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
