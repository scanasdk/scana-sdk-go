package moderation

import "net/http"

type moderationClient struct {
	appid, secretKey string
	textBusinessId   string // text audit business id
	conf             *config
	httpClient       *http.Client
}

/**
 * NewModerationClient 生成moderation client
 * @param {string} appid 调用凭证
 * @param {string} secretKey 调用凭证
 * @param {string} textBusinessId 文本审核的业务id
 * @returns {*TextModerationOutput} output 检测结果
 * @returns {error} err 错误消息
 */
func NewModerationClient(appid, secretKey, textBusinessId string, configurers ...configurer) (*moderationClient, error) {
	client := &moderationClient{
		appid:          appid,
		secretKey:      secretKey,
		textBusinessId: textBusinessId,
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
