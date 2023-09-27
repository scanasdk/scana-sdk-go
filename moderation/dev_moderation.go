// 开发者文本审核moderation
package moderation

import (
	"errors"
)

/**
 * TextModeration 发送 HTTP GET请求
 * @param {string} text 待检测的文本，必传
 * @param {string} businessId 文本业务id，必传
 * @returns {*TextModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) TextModeration(text string) (output *TextModerationOutput, result *APIResult, err error) {
	if text == "" {
		return nil, nil, errors.New("待检测文本不能为空")
	}

	body := map[string]interface{}{
		"text":        text,
		"business_id": client.textBusinessId,
	}

	var resp struct {
		APIResult
		Data TextModerationOutput `json:"data"`
	}
	if result, err := client.doPost("TextModeration", MODERATION_DOMAIN+"/kms-open/v2/openapi/synctextmoderation", body, client.Auth(client.textBusinessId), &resp); err != nil {
		return nil, result, err
	}

	output = &resp.Data
	result = &resp.APIResult

	return output, result, nil
}
