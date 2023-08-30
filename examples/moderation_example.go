package examples

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

// 同步文本审核接口
func ExampleTextSyncModeration() {
	mc, err := moderation.NewModerationClient("appId", "secret", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, result, err := mc.TextSyncModeration(&moderation.TextModerationInput{
		Text: []moderation.Text{
			{ContentId: "contentId", Data: "Hello World!"},
		},
		BusinessId: "businessId",
		Extra:      "extra",
	})
	if err != nil {
		if result != nil {
			log.Printf("code:%d,message:%s", result.Code, result.Msg)
		}
		log.Println(err)
		return
	}

	log.Printf("output=== %+#v", *output)
}
