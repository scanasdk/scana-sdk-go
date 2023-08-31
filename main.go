package main

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/examples"
	"github.com/scanasdk/scana-sdk-go/moderation"
)

func main() {
	err := examples.ExampleModeration("appId", "secret", "video", true, &moderation.VideoModerationInput{
		ContentId:  "aaa",
		URL:        "https://scanapi.qixincha.com/api/yobs/view/videos/kuaishou/scana-kuaishou-video-77IwV5uVAC.mp4",
		BusinessId: "businessId",
	})
	if err != nil {
		log.Println("moderation error:", err)
	}
}
