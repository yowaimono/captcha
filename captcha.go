package captcha

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"math/rand"
	"os"
	"time"

	log "github.com/yowaimono/captcha/internal/log"
)

// GetOne 生成一张验证码图片，并返回验证码ID和base64编码的图片
func GetOne(length int, format CaptchaFormat) (string, string, error) {
	log.Info("GetOne called with length: %d, format: %v", length, format)
	return GetBase64(length, format)
}

// GetBase64 生成一张验证码图片，并返回验证码ID和base64编码的图片
func GetBase64(length int, format CaptchaFormat) (string, string, error) {
	log.Info("GetBase64 called with length: %d, format: %v", length, format)

	// 生成指定长度和格式的随机验证码
	code := generateCaptchaCode(length, format)
	log.Info("Generated captcha code: %s", code)

	// 创建验证码图片
	img := createCaptchaImage(code)
	log.Info("Created captcha image")

	// 将图片编码为PNG格式
	var imgBuf bytes.Buffer
	err := png.Encode(&imgBuf, img)
	if err != nil {
		log.Error("Failed to encode image to PNG: %v", err)
		return "", "", err
	}
	log.Info("Encoded image to PNG")

	// 将图片编码为base64
	imgBase64 := base64.StdEncoding.EncodeToString(imgBuf.Bytes())
	log.Info("Encoded image to base64")

	// 生成验证码ID
	captchaID := generateCaptchaID()
	log.Info("Generated captcha ID: %s", captchaID)

	// 存储验证码信息
	storeCaptcha(captchaID, code)
	log.Info("Stored captcha information")

	// 启动一个goroutine来删除过期的验证码
	go func() {
		time.Sleep(60 * time.Second)
		deleteCaptcha(captchaID)
		log.Info("Deleted expired captcha with ID: %s", captchaID)
	}()

	return captchaID, imgBase64, nil
}

// GetImage 生成一张验证码图片，并返回验证码ID和image.Image对象
func GetImage(length int, format CaptchaFormat) (string, image.Image, error) {
	log.Info("GetImage called with length: %d, format: %v", length, format)

	// 生成指定长度和格式的随机验证码
	code := generateCaptchaCode(length, format)
	log.Info("Generated captcha code: %s", code)

	// 创建验证码图片
	img := createCaptchaImage(code)
	log.Info("Created captcha image")

	// 生成验证码ID
	captchaID := generateCaptchaID()
	log.Info("Generated captcha ID: %s", captchaID)

	// 存储验证码信息
	storeCaptcha(captchaID, code)
	log.Info("Stored captcha information")

	// 启动一个goroutine来删除过期的验证码
	go func() {
		time.Sleep(60 * time.Second)
		deleteCaptcha(captchaID)
		log.Info("Deleted expired captcha with ID: %s", captchaID)
	}()

	return captchaID, img, nil
}

// GetAndSave 生成一张验证码图片，并将其保存到指定路径，返回验证码ID和验证码内容
func GetAndSave(length int, format CaptchaFormat, savePath string) (string, string, error) {
	log.Info("GetAndSave called with length: %d, format: %v, savePath: %s", length, format, savePath)

	// 生成指定长度和格式的随机验证码
	code := generateCaptchaCode(length, format)
	log.Info("Generated captcha code: %s", code)

	// 创建验证码图片
	img := createCaptchaImage(code)
	log.Info("Created captcha image")

	// 将图片保存到指定路径
	file, err := os.Create(savePath)
	if err != nil {
		log.Error("Failed to create file: %v", err)
		return "", "", err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Error("Failed to encode image to PNG: %v", err)
		return "", "", err
	}
	log.Info("Saved image to path: %s", savePath)

	// 生成验证码ID
	captchaID := generateCaptchaID()
	log.Info("Generated captcha ID: %s", captchaID)

	// 存储验证码信息
	storeCaptcha(captchaID, code)
	log.Info("Stored captcha information")

	// 启动一个goroutine来删除过期的验证码
	go func() {
		time.Sleep(60 * time.Second)
		deleteCaptcha(captchaID)
		log.Info("Deleted expired captcha with ID: %s", captchaID)
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
	log.Info("Generated captcha ID: %s", string(id))
	return string(id)
}

// Verify 验证用户输入的验证码是否正确
func Verify(captchaID, userInput string) bool {
	log.Info("Verify called with captchaID: %s, userInput: %s", captchaID, userInput)

	info := getCaptcha(captchaID)
	if info == nil {
		log.Warn("Captcha not found for ID: %s", captchaID)
		return false
	}

	// 检查验证码是否过期
	if time.Now().After(info.ExpiresAt) {
		log.Warn("Captcha expired for ID: %s", captchaID)
		return false
	}

	// 验证用户输入的验证码
	isValid := info.Code == userInput
	if isValid {
		log.Info("Captcha verified successfully for ID: %s", captchaID)
	} else {
		log.Warn("Captcha verification failed for ID: %s", captchaID)
	}
	return isValid
}
