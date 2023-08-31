package examples

import (
	"errors"
	"log"

	"github.com/scanasdk/scana-sdk-go/moderation"
)

var (
	ErrInput = errors.New("invalid input:审核类型与input不一致")
)

// 示例方法
// ty:text/image/audio/video/doc
// sync:true/false for image/text
// input:输入参数 example:&TextModerationInput{}
func ExampleModeration(appId string, secret string, ty string, sync bool, input interface{}) error {
	mc, err := moderation.NewModerationClient(appId, secret, moderation.WithTimeout(10))
	if err != nil {
		log.Println("new moderation client failure", err)
		return err
	}

	var (
		output interface{}
		result *moderation.APIResult
	)
	switch ty {
	case "text":
		i, ok := input.(*moderation.TextModerationInput)
		if !ok {
			return ErrInput
		}
		if sync {
			output, result, err = mc.TextSyncModeration(i)
		} else {
			output, result, err = mc.TextAsyncModeration(i)
		}
	case "image":
		i, ok := input.(*moderation.ImageModerationInput)
		if !ok {
			return ErrInput
		}
		if sync {
			output, result, err = mc.ImageSyncModeration(i)
		} else {
			output, result, err = mc.ImageAsyncModeration(i)
		}
	case "audio":
		i, ok := input.(*moderation.AudioModerationInput)
		if !ok {
			return ErrInput
		}
		output, result, err = mc.AudioAsyncModeration(i)
	case "video":
		i, ok := input.(*moderation.VideoModerationInput)
		if !ok {
			return ErrInput
		}
		output, result, err = mc.VideoAsyncModeration(i)
	case "doc":
		i, ok := input.(*moderation.DocModerationInput)
		if !ok {
			return ErrInput
		}
		output, result, err = mc.DocAsyncModeration(i)
	}

	if err != nil {
		if result != nil {
			log.Printf("code:%d,message:%s", result.Code, result.Msg)
		}
		log.Println(err)
		return err
	}

	log.Printf("output=== %+#v", output)

	return nil
}
