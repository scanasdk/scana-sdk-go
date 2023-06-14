package moderation

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

func (client *moderationClient) generateSignature() url.Values {
	values := url.Values{}
	values.Add("nonce", randomString(10))
	values.Add("timestamp", strconv.Itoa(int(getCurrentTimeStamp())))
	values.Add("appId", client.appid)
	values.Add("businessId", client.textBusinessId)

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
 * @returns {*TextModerationOutput} output 检测结果
 * @returns {error} err 错误消息
 */
func (client *moderationClient) TextModeration(text string) (output *TextModerationOutput, err error) {
	if text == "" {
		return nil, errors.New("待检测文本不能为空")
	}

	body := map[string]interface{}{
		"text":        text,
		"business_id": client.textBusinessId,
	}

	var resp struct {
		Code int                  `json:"code"`
		Msg  string               `json:"msg"`
		Data TextModerationOutput `json:"data"`
	}
	if err := client.doPost("TextModeration", MODERATION_DOMAIN+"/kms-open/v2/openapi/synctextmoderation", body, client.generateSignature(), &resp); err != nil {
		return nil, err
	}

	if resp.Code != http.StatusOK {
		return nil, fmt.Errorf("text moderation request failure,response code:%d,msg:%s", resp.Code, resp.Msg)
	}
	output = &resp.Data

	return output, nil
}
