// 开发者文本审核moderation
package moderation

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/url"
	"sort"
	"strconv"
)

func (client *moderationClient) generateSignature(businessId string) url.Values {
	values := url.Values{}
	values.Add("nonce", randomString(10))
	values.Add("timestamp", strconv.Itoa(int(getCurrentTimeStamp())))
	values.Add("appId", client.appid)
	values.Add("businessId", businessId)

	var paramStr string
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		paramStr += key + values[key][0]
	}
	paramStr += client.secretKey

	md5Reader := md5.New()
	md5Reader.Write([]byte(paramStr))
	values.Add("signature", hex.EncodeToString(md5Reader.Sum(nil)))

	return values
}

/**
 * TextModeration 发送 HTTP GET请求
 * @param {string} text 待检测的文本，必传
 * @param {string} businessId 文本业务id，必传
 * @returns {*TextModerationOutput} output 检测结果
 * @returns {*APIResult} result 请求失败可以通过该返回判定状态码
 * @returns {error} err 错误消息
 */
func (client *moderationClient) TextModeration(text string, businessId string) (output *TextModerationOutput, result *APIResult, err error) {
	if text == "" {
		return nil, nil, errors.New("待检测文本不能为空")
	}

	body := map[string]interface{}{
		"text":        text,
		"business_id": businessId,
	}

	var resp struct {
		APIResult
		Data TextModerationOutput `json:"data"`
	}
	if result, err := client.doPost("TextModeration", MODERATION_DOMAIN+"/kms-open/v2/openapi/synctextmoderation", body, client.generateSignature(businessId), &resp); err != nil {
		return nil, result, err
	}

	output = &resp.Data
	result = &resp.APIResult

	return output, result, nil
}
