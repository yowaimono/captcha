// captcha_test.go
package captchp

import (
	"testing"
	"time"
)

func TestGetOneAndVerify(t *testing.T) {
	captchaID, _, err := GetOne(6, Mixed)
	if err != nil {
		t.Fatalf("生成验证码失败: %v", err)
	}

	// 验证码应该有效
	if !Verify(captchaID, "123456") {
		t.Fatalf("验证码验证失败")
	}

	// 等待验证码过期
	time.Sleep(61 * time.Second)

	// 验证码应该过期
	if Verify(captchaID, "123456") {
		t.Fatalf("验证码未过期")
	}
}