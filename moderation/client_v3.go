package moderation

import "net/http"

/**
 * NewModerationClient 生成moderation client
 * @param {string} appid 调用凭证
 * @param {string} secretKey 调用凭证
 * @returns {*TextModerationOutput} output 检测结果
 * @returns {error} err 错误消息
 */
func New(appid, secretKey string, configurers ...configurer) (*moderationClient, error) {
	client := &moderationClient{
		appid:     appid,
		secretKey: secretKey,
	}

	conf := &config{}
	conf.maxRetryCount = -1
	for _, configurer := range configurers {
		configurer(conf)
	}

	if err := conf.initConfigWithDefault(); err != nil {
		return nil, err
	}

	err := conf.getTransport()
	if err != nil {
		return nil, err
	}

	if conf.httpClient != nil {
		client.httpClient = conf.httpClient
		client.conf = conf
		return client, nil
	}

	client.conf = conf
	client.httpClient = &http.Client{
		Transport: conf.transport,
	}

	return client, nil
}
