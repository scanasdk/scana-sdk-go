package examples

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

// 开发者文本审核
func ExampleDevTextModeration() {
	mc, err := moderation.NewModerationClient("appid", "secret", "textBusinessId", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, result, err := mc.TextModeration("Hello World")
	if err != nil {
		if result != nil {
			log.Printf("code:%d,message:%s", result.Code, result.Msg)
			return
		}
		log.Println(err)
		return
	}

	log.Printf("output=== type:%s,score:%f", output.Type, output.Score)
}
