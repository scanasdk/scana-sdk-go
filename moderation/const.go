package moderation

const (
	MODERATION_DOMAIN = "https://newkmsapi.qixincha.com"
)

const (
	DEFAULT_TIMEOUT           = 10 // default http timeout
	DEFAULT_CONNECT_TIMEOUT   = 60
	DEFAULT_HEADER_TIMEOUT    = 60
	DEFAULT_IDLE_CONN_TIMEOUT = 30
	DEFAULT_MAX_RETRY_COUNT   = 3
	DEFAULT_MAX_CONN_PER_HOST = 1000
)

const (
	HTTP_GET     = "GET"
	HTTP_POST    = "POST"
	HTTP_PUT     = "PUT"
	HTTP_DELETE  = "DELETE"
	HTTP_HEAD    = "HEAD"
	HTTP_OPTIONS = "OPTIONS"
)
