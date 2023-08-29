// v3审核接口
package moderation

import (
	"errors"
	"fmt"
	"net/http"
)

/**
 * TextSyncModeration 文本同步审核请求
 * @param {TextSyncModerationInput} input 同步请求参数
 * @param {string} businessId 文本业务id，必传
 * @returns {*TextSyncModerationOutput} output 检测结果
 * @returns {error} err 错误消息
 */
func (client *moderationClient) TextSyncModeration(input *TextSyncModerationInput) (output *TextSyncModerationOutput, err error) {
	if input == nil {
		return nil, errors.New("nil input")
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
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		TextSyncModerationOutput
	}
	if err := client.doPost("TextSyncModeration", MODERATION_DOMAIN+"/v3/text/sync", body, nil /*no param*/, &resp); err != nil {
		return nil, err
	}
	doLog(LEVEL_DEBUG, "text sync moderation output:%+#v", resp.TextSyncModerationOutput)

	if resp.Code != http.StatusOK {
		return nil, fmt.Errorf("text moderation request failure,response code:%d,msg:%s", resp.Code, resp.Msg)
	}
	output = &resp.TextSyncModerationOutput

	return output, nil
}
