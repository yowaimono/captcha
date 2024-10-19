// aliyun_sms.go
package captcha

import (
	"bytes"
	"fmt"
	"math/rand"
	"text/template"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/yowaimono/captcha/internal/log"
)

// AliyunConfig 存储阿里云短信服务配置
type AliyunConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
}

var aliyunConfig AliyunConfig

// SetAliyunConfig 设置阿里云短信服务配置
func SetAliyunConfig(config AliyunConfig) {
	aliyunConfig = config
}

// 生成随机验证码
func generatePhoneCaptchaCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = digits[rand.Intn(len(digits))]
	}
	return string(code)
}

// 渲染模板
func renderTemplate(tpl string, data map[string]interface{}) (string, error) {
	t, err := template.New("sms").Parse(tpl)
	if err != nil {
		return "", err
	}

	var renderedTemplate bytes.Buffer
	if err := t.Execute(&renderedTemplate, data); err != nil {
		return "", err
	}

	return renderedTemplate.String(), nil
}

// 发送验证码到手机
func SendCaptchaToPhone(phoneNumber string, templateCode string, templateContent string, captchaLength int) (string, error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", aliyunConfig.AccessKeyID, aliyunConfig.AccessKeySecret)
	if err != nil {
		log.Error("Failed to create Aliyun client: %v", err)
		return "", err
	}

	captchaCode := generatePhoneCaptchaCode(captchaLength)
	data := map[string]interface{}{
		"code1": captchaCode,
	}

	renderedTemplate, err := renderTemplate(templateContent, data)
	if err != nil {
		log.Error("Failed to render template: %v", err)
		return "", err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phoneNumber
	request.SignName = aliyunConfig.SignName
	request.TemplateCode = templateCode
	request.TemplateParam = renderedTemplate

	response, err := client.SendSms(request)
	if err != nil {
		log.Error("Failed to send SMS: %v", err)
		return "", err
	}

	if response.Code != "OK" {
		log.Error("Failed to send SMS: %s", response.Message)
		return "", fmt.Errorf("failed to send SMS: %s", response.Message)
	}

	storeCaptcha(phoneNumber, captchaCode)
	return captchaCode, nil
}