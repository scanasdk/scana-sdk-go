# scana-sdk-go
scana golang sdk


# install 

`go get github.com/scanasdk/scana-sdk-go`

# usage

```go
import (
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

func main() {
	mc, err := moderation.NewModerationClient("appid", "secretKey", "businessId", moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return
	}
	output, err := mc.TextModeration("Hello world!")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("output=== type:%s,score:%f", output.Type, output.Score)
}
```