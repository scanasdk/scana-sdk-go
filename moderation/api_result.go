// APIResult 数据结构
// 可以通过APIResult 判断响应状态码、错误消息等
package moderation

type APIResult struct {
	// 错误码
	Code int `json:"code"`
	// 错误码对应的消息
	Msg string `json:"message"`
}
