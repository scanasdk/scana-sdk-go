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
	output, result, err := mc.DocAsyncModeration(&moderation.DocModerationInput{
		BusinessId: "1818208965452919186",
		Extra:      "GolangSDK-Test",
		ContentId:  "golang-sdk-doc",
		URL:        "https://files-rs.zsxq.com/Fmn3zw-CFYuZElHHnKjmu_iEOBun?attname=A6+A+Real+Room-Temperature+Superconductor.pdf&e=1711435534&token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:iJhA0gLoL6Z0JMacz4SNTDAhLtQ=",
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
