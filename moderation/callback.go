package moderation

import (
	"encoding/json"
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
		CallbackDataItemV3
	}
	if err = json.Unmarshal(jb, &v); err != nil {
		return nil, nil, err
	}

	result = &APIResult{Code: v.Code, Msg: v.Msg}
	resp = &v.CallbackDataItemV3

	return resp, result, nil
}
