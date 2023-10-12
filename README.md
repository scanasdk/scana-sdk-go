# scana-sdk-go

scana golang v3 sdk

v3版本接口，包括图片、文本同步接口，图片、文本、音频、视频、文档异步入单接口以及回调校验、解析方法

# install

`go get github.com/scanasdk/scana-sdk-go`

# usage

## [文本同步检测](./examples/text/main.go)

```go
package main

import (
	"github.com/scanasdk/scana-sdk-go/moderation"
	"log"
)

func main() {
	mc, err := moderation.New("appid", "secret", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, result, err := mc.TextSyncModeration(&moderation.TextModerationInput{
		BusinessId: "business_id",
		Extra:      "GolangSDK-Test",
		Text: []moderation.Text{
			{ContentId: "golang-sdk", Data: "Hello World!"},
		}})
	if err != nil {
		if result != nil {
			log.Printf("code:%d,message:%s", result.Code, result.Msg)
			return
		}
		log.Println(err)
		return
	}

	log.Printf("output=== type:%v", output)
}
```

输出:

```shell
2023/09/28 15:28:13 output=== type:&{1 [{200 OK 1 正常  golang-sdk 65152b0c08602c000179e182 [] []}] 94f7d244-5dd0-11ee-a2a5-fa163e917534 GolangSDK-Test}
```

## [图片同步检测](./examples/image/main.go)

```go
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
```

输出:

```shell
2023/09/28 15:34:02 output:&{1 [{200 OK 1 正常 正常 golang-sdk-image 65152c6a08602c000179e569 {[] [] [] 1} []}] 6d87b145-3627-4af6-b922-ee5b8d2602a8 GolangSDK-Test}
```

## [音频异步检测](./examples/audio/main.go)

```go
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
	output, result, err := mc.AudioAsyncModeration(&moderation.AudioModerationInput{
		BusinessId: "business_id",
		Extra:      "GolangSDK-Test",
		ContentId:  "golang-sdk-audio",
		URL:        "https://www.baidu.com/audio?attname=test.mp3",
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

```

输出:

```shell
2023/09/28 15:36:34 output:&{1 bff38f65-5dd1-11ee-a613-fa163e91756e GolangSDK-Test}
```

## [视频异步检测](./examples/video/main.go)

```go
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
	output, result, err := mc.VideoAsyncModeration(&moderation.VideoModerationInput{
		BusinessId: "business_id",
		Extra:      "GolangSDK-Test",
		ContentId:  "golang-sdk-video",
		URL:        "https://www.baidu.com/video?attname=test.mp4",
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

```

输出:

```shell
2023/09/28 15:36:34 output:&{1 bff38f65-5dd3-11ee-d413-fa163e91756e GolangSDK-Test}
```

## [文档异步检测](./examples/doc/main.go)

```go
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
	output, result, err := mc.DocAsyncModeration(&moderation.DocModerationInput{
		BusinessId: "business_id",
		Extra:      "GolangSDK-Test",
		ContentId:  "golang-sdk-doc",
		URL:        "https://www.baidu.com/doc?attname=test.doc",
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

```

输出:

```shell
2023/09/28 15:36:34 output:&{1 bff58f65-fdd3-11ee-d413-fad23e91756e GolangSDK-Test}
```