// v3审核接口
package moderation

import (
	"errors"
)

/**
 * TextSyncModeration 文本同步审核请求
 * @param {*TextModerationInput} input 同步请求参数
 * @returns {*TextSyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) TextSyncModeration(input *TextModerationInput) (output *TextSyncModerationOutput, result *APIResult, err error) {
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
 * @param {*ImageModerationInput} input 同步请求参数
 * @returns {*ImageSyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) ImageSyncModeration(input *ImageModerationInput) (output *ImageSyncModerationOutput, result *APIResult, err error) {
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

/**
 * TextAsyncModeration 文本异步审核请求
 * @param {*TextModerationInput} input 请求参数
 * @returns {*TextAsyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 * 异步审核只判定是否成功，需要您通过回调的形式获取处理结果
 */
func (client *moderationClient) TextAsyncModeration(input *TextModerationInput) (output *TextAsyncModerationOutput, result *APIResult, err error) {
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

	doLog(LEVEL_DEBUG, "text async moderation input:%+#v", body)
	var resp struct {
		APIResult
		TextAsyncModerationOutput
	}
	if result, err := client.doPost("TextAsyncModeration", MODERATION_DOMAIN+"/v3/text/async", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "text async moderation output:%+#v", resp.TextAsyncModerationOutput)

	output = &resp.TextAsyncModerationOutput
	result = &resp.APIResult

	return output, result, nil
}

/**
 * ImageAsyncModeration 图片异步审核请求
 * @param {*ImageModerationInput} input 请求参数
 * @returns {*ImageAsyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) ImageAsyncModeration(input *ImageModerationInput) (output *ImageAsyncModerationOutput, result *APIResult, err error) {
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

	doLog(LEVEL_DEBUG, "image async moderation input:%+#v", body)
	var resp struct {
		APIResult
		ImageAsyncModerationOutput
	}
	if result, err := client.doPost("ImageAsyncModeration", MODERATION_DOMAIN+"/v3/image/async", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "image async moderation output:%+#v", resp.ImageAsyncModerationOutput)

	output = &resp.ImageAsyncModerationOutput
	result = &resp.APIResult

	return output, result, nil
}
