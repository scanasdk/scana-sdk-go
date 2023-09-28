# scana-sdk-go
scana golang v3 sdk

v3版本接口，包括图片、文本同步接口，图片、文本、音频、视频、文档异步入单接口以及回调校验、解析方法


# install 

`go get github.com/scanasdk/scana-sdk-go/v3`

# usage

```go
package main

import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

func main() {
	mc, err := moderation.New("appid", "secret", moderation.WithTimeout(10))
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
		log.Println(err)
		return
	}
	if result != nil {
		log.Printf("moderation failure,code:%d,msg:%s\n", result.Code, result.Msg)
	}

	log.Printf("output=== %+#v", *output)
}

```