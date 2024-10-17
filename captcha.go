// captcha.go
package captchp

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"math/rand"
	"time"
)

// GetOne 生成一张验证码图片，并返回验证码ID和base64编码的图片
func GetOne(length int, format CaptchaFormat) (string, string, error) {
	// 生成指定长度和格式的随机验证码
	code := generateCaptchaCode(length, format)

	// 创建验证码图片
	img := createCaptchaImage(code)

	// 将图片编码为PNG格式
	var imgBuf bytes.Buffer
	err := png.Encode(&imgBuf, img)
	if err != nil {
		return "", "", err
	}

	// 将图片编码为base64
	imgBase64 := base64.StdEncoding.EncodeToString(imgBuf.Bytes())

	// 生成验证码ID
	captchaID := generateCaptchaID()

	// 存储验证码信息
	storeCaptcha(captchaID, code)

	// 启动一个goroutine来删除过期的验证码
	go func() {
		time.Sleep(60 * time.Second)
		deleteCaptcha(captchaID)
	}()

	return captchaID, imgBase64, nil
}

// 生成验证码ID
func generateCaptchaID() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := make([]byte, 10)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}
	return string(id)
}

// Verify 验证用户输入的验证码是否正确
func Verify(captchaID, userInput string) bool {
	info := getCaptcha(captchaID)
	if info == nil {
		return false
	}

	// 检查验证码是否过期
	if time.Now().After(info.ExpiresAt) {
		return false
	}

	// 验证用户输入的验证码
	return info.Code == userInput
}
