// captcha_storage.go
package captcha

import (
	"sync"
	"time"

	log "github.com/yowaimono/captcha/internal/log"
)

// CaptchaInfo 存储验证码信息
type CaptchaInfo struct {
	Code      string
	ExpiresAt time.Time
}

var (
	captchaMap  = make(map[string]*CaptchaInfo)
	captchaLock sync.RWMutex
)

// 存储验证码信息
func storeCaptcha(captchaID, code string) {
	captchaLock.Lock()
	defer captchaLock.Unlock()

	captchaMap[captchaID] = &CaptchaInfo{
		Code:      code,
		ExpiresAt: time.Now().Add(60 * time.Second),
	}
	log.Info("Stored captcha with ID: %s, code: %s", captchaID, code)
}

// 获取验证码信息
func getCaptcha(captchaID string) *CaptchaInfo {
	captchaLock.RLock()
	defer captchaLock.RUnlock()

	info, exists := captchaMap[captchaID]
	if !exists {
		log.Warn("Captcha not found for ID: %s", captchaID)
		return nil
	}

	log.Info("Retrieved captcha with ID: %s, code: %s", captchaID, info.Code)
	return info
}

// 删除验证码信息
func deleteCaptcha(captchaID string) {
	captchaLock.Lock()
	defer captchaLock.Unlock()

	delete(captchaMap, captchaID)
	log.Info("Deleted captcha with ID: %s", captchaID)
}
