package moderation

import (
	"math/rand"
	"time"
)

var DEFAULT_TIMELOC = time.FixedZone("CST", 8*3600)

func init() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
}

func getCurrentTime() time.Time {
	return time.Now().In(DEFAULT_TIMELOC)
}

func getCurrentTimeStamp() int64 {
	return time.Now().Unix()
}

// 生成指定长度的随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
