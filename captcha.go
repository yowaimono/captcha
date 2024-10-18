// captcha.go
package captcha

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"math/rand"
	"os"
	"time"
)

// GetOne 生成一张验证码图片，并返回验证码ID和base64编码的图片
func GetOne(length int, format CaptchaFormat) (string, string, error) {
	// 直接调用GetBase64方法
	return GetBase64(length, format)
}

// GetBase64 生成一张验证码图片，并返回验证码ID和base64编码的图片
func GetBase64(length int, format CaptchaFormat) (string, string, error) {
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

// GetImage 生成一张验证码图片，并返回验证码ID和image.Image对象
func GetImage(length int, format CaptchaFormat) (string, image.Image, error) {
	// 生成指定长度和格式的随机验证码
	code := generateCaptchaCode(length, format)

	// 创建验证码图片
	img := createCaptchaImage(code)

	// 生成验证码ID
	captchaID := generateCaptchaID()

	// 存储验证码信息
	storeCaptcha(captchaID, code)

	// 启动一个goroutine来删除过期的验证码
	go func() {
		time.Sleep(60 * time.Second)
		deleteCaptcha(captchaID)
	}()

	return captchaID, img, nil
}

// GetAndSave 生成一张验证码图片，并将其保存到指定路径，返回验证码ID和验证码内容
func GetAndSave(length int, format CaptchaFormat, savePath string) (string, string, error) {
	// 生成指定长度和格式的随机验证码
	code := generateCaptchaCode(length, format)

	// 创建验证码图片
	img := createCaptchaImage(code)

	// 将图片保存到指定路径
	file, err := os.Create(savePath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return "", "", err
	}

	// 生成验证码ID
	captchaID := generateCaptchaID()

	// 存储验证码信息
	storeCaptcha(captchaID, code)

	// 启动一个goroutine来删除过期的验证码
	go func() {
		time.Sleep(60 * time.Second)
		deleteCaptcha(captchaID)
	}()

	return captchaID, code, nil
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
