// captcha_format.go
package captcha

import (
	"math/rand"
	"time"

	log "github.com/yowaimono/captcha/internal/log"
)

// CaptchaFormat 定义验证码格式
type CaptchaFormat string

const (
	Mixed  CaptchaFormat = "mixed" // 大小写字母和数字混合
	AplusN CaptchaFormat = "A+N"   // 大写字母和数字混合
	aPlusN CaptchaFormat = "a+N"   // 小写字母和数字混合
	aPlusA CaptchaFormat = "a+A"   // 大写和小写字母混合
)

// 生成指定长度和格式的随机验证码
func generateCaptchaCode(length int, format CaptchaFormat) string {
	log.Info("generateCaptchaCode called with length: %d, format: %v", length, format)

	rand.Seed(time.Now().UnixNano())
	var charset string

	switch format {
	case Mixed:
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
		log.Info("Using charset for Mixed format")
	case AplusN:
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		log.Info("Using charset for AplusN format")
	case aPlusN:
		charset = "abcdefghijklmnopqrstuvwxyz0123456789"
		log.Info("Using charset for aPlusN format")
	case aPlusA:
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		log.Info("Using charset for aPlusA format")
	default:
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		log.Info("Using default charset")
	}

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	log.Info("Generated captcha code: %s", string(code))
	return string(code)
}
