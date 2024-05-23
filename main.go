package main

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/examples"
	"github.com/scanasdk/scana-sdk-go/moderation"
)

func main() {
	err := examples.ExampleModeration("appId", "secret", examples.CALL_TYPE_DOCUMENT, true, &moderation.VideoModerationInput{
		ContentId:  "aaa",
		URL:        "https://scanapi.qixincha.com/api/yobs/view/videos/kuaishouVAC.mp4",
		BusinessId: "businessId",
	})
	if err != nil {
		log.Println("moderation error:", err)
	}
}
