// captcha_storage.go
package captchp

import (
	"sync"
	"time"
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
	captchaMap[captchaID] = &CaptchaInfo{
		Code:      code,
		ExpiresAt: time.Now().Add(60 * time.Second),
	}
	captchaLock.Unlock()
}

// 获取验证码信息
func getCaptcha(captchaID string) *CaptchaInfo {
	captchaLock.RLock()
	info, exists := captchaMap[captchaID]
	captchaLock.RUnlock()

	if !exists {
		return nil
	}

	return info
}

// 删除验证码信息
func deleteCaptcha(captchaID string) {
	captchaLock.Lock()
	delete(captchaMap, captchaID)
	captchaLock.Unlock()
}
