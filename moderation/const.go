package moderation

const (
	MODERATION_DOMAIN = "http://dev-newkmsapi.qixincha.com"

	DEFAULT_LOG_LEVEL = LEVEL_DEBUG //default log level
)

const (
	CONTENT_TYPE_TEXT  = 1
	CONTENT_TYPE_IMAGE = 2
	CONTENT_TYPE_AUDIO = 3
	CONTENT_TYPE_VIDEO = 4
	CONTENT_TYPE_DOC   = 5
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
