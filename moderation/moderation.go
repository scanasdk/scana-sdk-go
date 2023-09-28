// v3审核接口
package moderation

import (
	"errors"
	"net/http"
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
	if result, err := client.doPost("TextSyncModeration", MODERATION_DOMAIN+"/kms-open/v3/text/sync", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "text sync moderation output:%+#v", resp.TextSyncModerationOutput)

	output = &resp.TextSyncModerationOutput
	result = &resp.APIResult

	if result.Code != http.StatusOK {
		return output, result, errors.New(resp.Msg)
	}

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
	if result, err := client.doPost("ImageSyncModeration", MODERATION_DOMAIN+"/kms-open/v3/image/sync", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "image sync moderation output:%+#v", resp.ImageSyncModerationOutput)

	output = &resp.ImageSyncModerationOutput
	result = &resp.APIResult

	if result.Code != http.StatusOK {
		return output, result, errors.New(resp.Msg)
	}

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
	if result, err := client.doPost("TextAsyncModeration", MODERATION_DOMAIN+"/kms-open/v3/text/async", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "text async moderation output:%+#v", resp.TextAsyncModerationOutput)

	output = &resp.TextAsyncModerationOutput
	result = &resp.APIResult

	if result.Code != http.StatusOK {
		return output, result, errors.New(resp.Msg)
	}

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
	if result, err := client.doPost("ImageAsyncModeration", MODERATION_DOMAIN+"/kms-open/v3/image/async", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "image async moderation output:%+#v", resp.ImageAsyncModerationOutput)

	output = &resp.ImageAsyncModerationOutput
	result = &resp.APIResult

	if result.Code != http.StatusOK {
		return output, result, errors.New(resp.Msg)
	}

	return output, result, nil
}

/**
 * AudioAsyncModeration 音频异步审核请求
 * @param {*AudioModerationInput} input 请求参数
 * @returns {*AudioAsyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) AudioAsyncModeration(input *AudioModerationInput) (output *AudioAsyncModerationOutput, result *APIResult, err error) {
	if input == nil {
		return nil, nil, errors.New("nil input")
	}
	if input.BusinessId == "" {
		return nil, nil, errors.New("businessId is not set")
	}
	if input.URL == "" {
		return nil, nil, errors.New("url is not set")
	}

	body := map[string]interface{}{
		"appId":      client.appid,
		"secretKey":  client.secretKey,
		"businessId": input.BusinessId,
		"contentId":  input.ContentId,
		"extra":      input.Extra,
		"url":        input.URL,
	}

	doLog(LEVEL_DEBUG, "audio async moderation input:%+#v", body)
	var resp struct {
		APIResult
		AudioAsyncModerationOutput
	}
	if result, err := client.doPost("AudioAsyncModeration", MODERATION_DOMAIN+"/kms-open/v3/audio/async", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "audio async moderation output:%+#v", resp.AudioAsyncModerationOutput)

	output = &resp.AudioAsyncModerationOutput
	result = &resp.APIResult

	if result.Code != http.StatusOK {
		return output, result, errors.New(resp.Msg)
	}

	return output, result, nil
}

/**
 * VideoAsyncModeration 视频异步审核请求
 * @param {*VideoModerationInput} input 请求参数
 * @returns {*VideoAsyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) VideoAsyncModeration(input *VideoModerationInput) (output *VideoAsyncModerationOutput, result *APIResult, err error) {
	if input == nil {
		return nil, nil, errors.New("nil input")
	}
	if input.BusinessId == "" {
		return nil, nil, errors.New("businessId is not set")
	}
	if input.URL == "" {
		return nil, nil, errors.New("url is not set")
	}

	body := map[string]interface{}{
		"appId":      client.appid,
		"secretKey":  client.secretKey,
		"businessId": input.BusinessId,
		"contentId":  input.ContentId,
		"extra":      input.Extra,
		"url":        input.URL,
	}

	doLog(LEVEL_DEBUG, "video async moderation input:%+#v", body)
	var resp struct {
		APIResult
		VideoAsyncModerationOutput
	}
	if result, err := client.doPost("VideoAsyncModeration", MODERATION_DOMAIN+"/kms-open/v3/video/async", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "video async moderation output:%+#v", resp.VideoAsyncModerationOutput)

	output = &resp.VideoAsyncModerationOutput
	result = &resp.APIResult

	if result.Code != http.StatusOK {
		return output, result, errors.New(resp.Msg)
	}
	return output, result, nil
}

/**
 * DocAsyncModeration 文档异步审核请求
 * @param {*DocModerationInput} input 请求参数
 * @returns {*DocAsyncModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) DocAsyncModeration(input *DocModerationInput) (output *DocAsyncModerationOutput, result *APIResult, err error) {
	if input == nil {
		return nil, nil, errors.New("nil input")
	}
	if input.BusinessId == "" {
		return nil, nil, errors.New("businessId is not set")
	}
	if input.URL == "" {
		return nil, nil, errors.New("url is not set")
	}

	body := map[string]interface{}{
		"appId":      client.appid,
		"secretKey":  client.secretKey,
		"businessId": input.BusinessId,
		"contentId":  input.ContentId,
		"extra":      input.Extra,
		"url":        input.URL,
	}

	doLog(LEVEL_DEBUG, "doc async moderation input:%+#v", body)
	var resp struct {
		APIResult
		DocAsyncModerationOutput
	}
	if result, err := client.doPost("DocAsyncModeration", MODERATION_DOMAIN+"/kms-open/v3/doc/async", body, nil /*no param*/, &resp); err != nil {
		return nil, result, err
	}
	doLog(LEVEL_DEBUG, "doc async moderation output:%+#v", resp.DocAsyncModerationOutput)

	output = &resp.DocAsyncModerationOutput
	result = &resp.APIResult

	if result.Code != http.StatusOK {
		return output, result, errors.New(resp.Msg)
	}

	return output, result, nil
}
