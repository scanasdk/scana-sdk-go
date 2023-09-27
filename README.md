# scana-sdk-go
scana golang v3 sdk

v3版本接口，包括图片、文本同步接口，图片、文本、音频、视频、文档异步入单接口以及回调校验、解析方法


# install 

`go get github.com/scanasdk/scana-sdk-go/v3`

# usage

```go
import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation/v3"
)

func main() {
	mc, err := moderation.NewModerationClient("appid", "secret", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, err := mc.TextSyncModeration(&moderation.TextSyncModerationInput{
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

	log.Printf("output=== %+#v", *output)
}
```