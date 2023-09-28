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
	output, result, err := mc.AudioAsyncModeration(&moderation.AudioModerationInput{
		BusinessId: "1720063483019526161",
		Extra:      "GolangSDK-Test",
		ContentId:  "golang-sdk-audio",
		URL:        "https://files-rs.zsxq.com/lt1UJY5sdy4GH4vgz1nNW0We392F?attname=%E7%AE%A1%E4%B8%AD%E7%AA%A5%E8%B1%B9%EF%BC%8C%E6%94%B6%E5%85%A5%E7%BB%93%E6%9E%84%E7%9C%8B%E5%8C%BB%E6%94%B920230927.mp3&e=1711435288&token=kIxbL07-8jAj8w1n4s9zv64FuZZNEATmlU_Vm6zD:z48Bjhm0K1n0Fj78GYjQ0tPa9fk=",
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
