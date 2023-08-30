package moderation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
 * 发送 HTTP GET请求
 * @param {string} action 请求的 action
 * @param {string} url 请求的 url
 * @param {url.Values} params 请求的参数，可选
 * @returns {interface{}} receiver 解析响应的数据结构，请传递指针
 * @returns {error} err 错误消息
 */
func (client *moderationClient) doGet(action, url string, params url.Values, receiver interface{}) (*APIResult, error) {
	resp, result, err := client.doRequest(action, url, HTTP_GET, nil, params)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(resp, receiver)
	return result, err
}

/**
 * 发送 HTTP POST请求
 * @param {string} action 请求的 action
 * @param {string} url 请求的 url
 * @param {url.Values} params 请求的参数，可选
 * @param {map[string]interface{}} data 请求的数据，可选
 * @returns {interface{}} receiver 解析响应的数据结构，请传递指针
 * @returns {error} err 错误消息
 */
func (client *moderationClient) doPost(action, url string, data map[string]interface{}, params url.Values, receiver interface{}) (*APIResult, error) {
	resp, result, err := client.doRequest(action, url, HTTP_POST, data, params)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(resp, receiver)
	return result, err
}

/**
 * 发送 HTTP 请求
 * @param {string} action 请求的 action
 * @param {string} url 请求的 url
 * @param {string} method 请求的方法，如 'GET' 或 'POST'
 * @param {map[string]interface{}} data 请求的数据，可选
 * @param {url.Values} params 请求的参数，可选
 * @returns {[]byte} respData 响应消息字节数据
 * @returns {error} err 错误消息
 */
func (client *moderationClient) doRequest(action, url, method string, data map[string]interface{}, params url.Values) (respData []byte, result *APIResult, err error) {
	doLog(LEVEL_INFO, "Enter method %s...", action)

	start := getCurrentTimeStamp()

	if params != nil {
		url += ("?" + params.Encode())
	}

	var bodyr bytes.Buffer
	if data != nil {
		jb, err := json.Marshal(data)
		checkAndLogErr(err, LEVEL_ERROR, "json marshal failure")
		bodyr = *bytes.NewBuffer(jb)
	}
	doLog(LEVEL_DEBUG, "%s request url:%s", action, url)
	req, err := http.NewRequest(method, url, &bodyr)
	if err != nil {
		doLog(LEVEL_ERROR, "http new request failure:%s", err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		doLog(LEVEL_ERROR, "http response failure:%s", err.Error())
		return
	}
	defer resp.Body.Close()

	result = &APIResult{
		Code: resp.StatusCode,
		Msg:  resp.Status,
	}
	if resp.StatusCode != http.StatusOK {
		return nil, result, fmt.Errorf("%s response code:%d", action, resp.StatusCode)
	}

	doLog(LEVEL_DEBUG, "method finished %s elapsed:%d", action, (getCurrentTimeStamp() - start))

	respData, err = ioutil.ReadAll(resp.Body)
	return respData, result, err

}
