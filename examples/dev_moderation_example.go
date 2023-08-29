package examples

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

// 开发者文本审核
func ExampleDevTextModeration() {
	mc, err := moderation.NewModerationClient("appid", "secret", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, err := mc.TextModeration("Hello World", "businessId")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("output=== type:%s,score:%f", output.Type, output.Score)
}
