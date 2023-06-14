package moderation

import (
	"net"
	"net/http"
	"time"
)

type config struct {
	connectTimeout  int
	headerTimeout   int
	maxConnsPerHost int
	idleConnTimeout int
	timeout         int
	maxRetryCount   int
	httpClient      *http.Client
	transport       *http.Transport
}

func (conf *config) initConfigWithDefault() error {
	if conf.timeout <= 0 {
		conf.timeout = DEFAULT_TIMEOUT
	}
	if conf.connectTimeout <= 0 {
		conf.connectTimeout = DEFAULT_CONNECT_TIMEOUT
	}
	if conf.headerTimeout <= 0 {
		conf.headerTimeout = DEFAULT_HEADER_TIMEOUT
	}
	if conf.idleConnTimeout <= 0 {
		conf.idleConnTimeout = DEFAULT_IDLE_CONN_TIMEOUT
	}
	if conf.maxRetryCount < 0 {
		conf.maxRetryCount = DEFAULT_MAX_RETRY_COUNT
	}
	if conf.maxConnsPerHost <= 0 {
		conf.maxConnsPerHost = DEFAULT_MAX_CONN_PER_HOST
	}

	return nil
}

func (conf *config) getTransport() error {
	if conf.transport == nil {
		conf.transport = &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Second*time.Duration(conf.connectTimeout))
			},
			MaxIdleConns:          conf.maxConnsPerHost,
			MaxIdleConnsPerHost:   conf.maxConnsPerHost,
			ResponseHeaderTimeout: time.Second * time.Duration(conf.headerTimeout),
			IdleConnTimeout:       time.Second * time.Duration(conf.idleConnTimeout),
		}
	}

	return nil
}

type configurer func(conf *config)

func WithTimeout(timeout int) configurer {
	return func(conf *config) {
		conf.timeout = timeout
	}
}

func WithMaxRetryCount(maxRetryCount int) configurer {
	return func(conf *config) {
		conf.maxRetryCount = maxRetryCount
	}
}

func WithHttpClient(httpClient *http.Client) configurer {
	return func(conf *config) {
		conf.httpClient = httpClient
	}
}

func WithConnectTimeout(connectTimeout int) configurer {
	return func(conf *config) {
		conf.connectTimeout = connectTimeout
	}
}

func WithHeaderTimeout(headerTimeout int) configurer {
	return func(conf *config) {
		conf.headerTimeout = headerTimeout
	}
}

func WithMaxConnections(maxConnsPerHost int) configurer {
	return func(conf *config) {
		conf.maxConnsPerHost = maxConnsPerHost
	}
}

func WithIdleConnTimeout(idleConnTimeout int) configurer {
	return func(conf *config) {
		conf.idleConnTimeout = idleConnTimeout
	}
}
