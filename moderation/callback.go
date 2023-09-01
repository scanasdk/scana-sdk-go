package moderation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ParseCallback 解析回调参数，异步审核时，接口接收回调可以使用该方法解析获得审核数据
// r:回调请求的http.Request
func (client *moderationClient) ParseCallback(r *http.Request) (resp *CallbackDataItemV3, result *APIResult, err error) {
	doLog(LEVEL_INFO, "callback parse body Enter...")
	query := r.URL.Query()
	if err = client.ValidAuth(query); err != nil {
		doLog(LEVEL_DEBUG, "callback valid auth fail,err:%v", err)
		return nil, nil, err
	}

	jb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, nil, err
	}

	var v struct {
		Code int    `json:"code"`
		Msg  string `json:"message"`
		callbackDataItemV3
	}
	if err = json.Unmarshal(jb, &v); err != nil {
		return nil, nil, err
	}

	// parse diffirent type result
	switch v.ContentType {
	case CONTENT_TYPE_TEXT:
		var data TextModerationResult
		if err := json.Unmarshal([]byte(v.TextModerationResult_), &data); err != nil {
			doLog(LEVEL_ERROR, "pase different type string to struct failure:%v", err)
			return nil, nil, err
		}
		v.TextModerationResult = &data
	case CONTENT_TYPE_IMAGE:
		var data ImageModerationResult
		if err := json.Unmarshal([]byte(v.ImageModerationResult_), &data); err != nil {
			doLog(LEVEL_ERROR, "pase different type string to struct failure:%v", err)
			return nil, nil, err
		}
		v.ImageModerationResult = &data
	case CONTENT_TYPE_AUDIO:
		var data AudioModerationResult
		if err := json.Unmarshal([]byte(v.AudioModerationResult_), &data); err != nil {
			doLog(LEVEL_ERROR, "pase different type string to struct failure:%v", err)
			return nil, nil, err
		}
		v.AudioModerationResult = &data
	case CONTENT_TYPE_VIDEO:
		var data VideoModerationResult
		if err := json.Unmarshal([]byte(v.VideoModerationResult_), &data); err != nil {
			doLog(LEVEL_ERROR, "pase different type string to struct failure:%v", err)
			return nil, nil, err
		}
		v.VideoModerationResult = &data
	case CONTENT_TYPE_DOC:
		var data DocModerationResult
		if err := json.Unmarshal([]byte(v.DocModerationResult_), &data); err != nil {
			doLog(LEVEL_ERROR, "pase different type string to struct failure:%v", err)
			return nil, nil, err
		}
		v.DocModerationResult = &data
	default:
		return nil, nil, fmt.Errorf("unknown type:%d", v.ContentType)
	}

	result = &APIResult{Code: v.Code, Msg: v.Msg}
	resp = &v.CallbackDataItemV3

	return resp, result, nil
}
