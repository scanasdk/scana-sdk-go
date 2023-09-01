package moderation

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"time"
)

var (
	ErrSignError = errors.New("signature error")
	signTimeout  = time.Minute * 5
)

const (
	signKey = "signature"
)

// 通过Auth方法可以获得加签后的query参数(url.Values)
func (client *moderationClient) Auth(businessId string) url.Values {
	values := url.Values{}
	values.Add("nonce", randomString(10))
	values.Add("timestamp", strconv.Itoa(int(getCurrentTimeStamp())))
	values.Add("appId", client.appid)
	values.Add("businessId", businessId)

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
	values.Add(signKey, hex.EncodeToString(md5Reader.Sum(nil)))

	return values
}

// 校验签名正确性
// 接收回调后可以使用该方法校验正确性
// values：query参数
func (client *moderationClient) ValidAuth(values url.Values) error {
	if !values.Has("appId") ||
		!values.Has("businessId") ||
		!values.Has("timestamp") ||
		!values.Has(signKey) {
		return errors.New("参数缺失")
	}

	timestamp, _ := strconv.Atoi(values.Get("timestamp"))
	thatTime := time.Unix(int64(timestamp), 0)
	if time.Since(thatTime) > signTimeout {
		return fmt.Errorf("TIMESTAMP_TIMEOUT:<%d>", timestamp)
	}

	doLog(LEVEL_DEBUG, "validauth:client sign:%s", values.Get(signKey))

	keys := make([]string, 0, len(values))
	for k := range values {
		if k != signKey {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var paramStr string
	for _, key := range keys {
		paramStr += key + values[key][0]
	}
	paramStr += client.secretKey

	md5Reader := md5.New()
	md5Reader.Write([]byte(paramStr))

	clientSign := values.Get(signKey)
	serverSign := hex.EncodeToString(md5Reader.Sum(nil))
	doLog(LEVEL_DEBUG, "validauth:server sign:%s", serverSign)

	if clientSign != serverSign {
		return ErrSignError
	}

	return nil
}
