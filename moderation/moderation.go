// v3审核接口
package moderation

import (
	"errors"
)

/**
 * TextSyncModeration 文本同步审核请求
 * @param {*TextSyncModerationInput} input 同步请求参数
 * @returns {*TextSyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) TextSyncModeration(input *TextSyncModerationInput) (output *TextSyncModerationOutput, result *APIResult, err error) {
	if input == nil {
		return nil, nil, errors.New("nil input")
	}
	if input.BusinessId == "" {
		return nil, nil, errors.New("businessId is not set")
	}

	body := map[string]interface{}{
		"appId":      client.appid,
		"secretKey":  client.secretKey,
		"businessId": input.BusinessId,
		"text":       input.Text,
		"extra":      input.Extra,
	}

	doLog(LEVEL_DEBUG, "text sync moderation input:%+#v", body)
	var resp struct {
		APIResult
		TextSyncModerationOutput
	}
	if result, err := client.doPost("TextSyncModeration", MODERATION_DOMAIN+"/v3/text/sync", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "text sync moderation output:%+#v", resp.TextSyncModerationOutput)

	output = &resp.TextSyncModerationOutput
	result = &resp.APIResult

	return output, result, nil
}

/**
 * ImageSyncModeration 图片同步审核请求
 * @param {*ImageSyncModerationInput} input 同步请求参数
 * @returns {*ImageSyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) ImageSyncModeration(input *ImageSyncModerationInput) (output *ImageSyncModerationOutput, result *APIResult, err error) {
	if input == nil {
		return nil, nil, errors.New("nil input")
	}
	if input.BusinessId == "" {
		return nil, nil, errors.New("businessId is not set")
	}

	body := map[string]interface{}{
		"appId":      client.appid,
		"secretKey":  client.secretKey,
		"businessId": input.BusinessId,
		"images":     input.Images,
		"extra":      input.Extra,
	}

	doLog(LEVEL_DEBUG, "image sync moderation input:%+#v", body)
	var resp struct {
		APIResult
		ImageSyncModerationOutput
	}
	if result, err := client.doPost("ImageSyncModeration", MODERATION_DOMAIN+"/v3/image/sync", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "image sync moderation output:%+#v", resp.ImageSyncModerationOutput)

	output = &resp.ImageSyncModerationOutput
	result = &resp.APIResult

	return output, result, nil
}
