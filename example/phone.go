// // main.go
package example

// import (
// 	"fmt"
// 	"github.com/yowaimono/captcha"
// )

// func main() {
// 	phoneNumber := "1234567890"
// 	templateCode := "your_template_code"
// 	templateContent := `{"code":"{{.code1}}"}`
// 	captchaLength := 6

// 	// 发送验证码
// 	captchaCode, err := captcha.SendCaptchaToPhone(phoneNumber, templateCode, templateContent, captchaLength)
// 	if err != nil {
// 		fmt.Println("Failed to send captcha:", err)
// 		return
// 	}
// 	fmt.Println("Captcha sent successfully. Code:", captchaCode)

// 	// 模拟用户输入验证码
// 	userInputCode := "123456" // 这里应该是用户输入的验证码

// 	// 验证验证码
// 	if captcha.VerifyCaptcha(phoneNumber, userInputCode) {
// 		fmt.Println("Captcha verification successful")
// 	} else {
// 		fmt.Println("Captcha verification failed")
// 	}
// }
