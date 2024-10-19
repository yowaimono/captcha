package example

// import (
// 	"fmt"

// 	"github.com/yowaimono/captcha"
// )

// func main() {
// 	// 生成一个长度为 6 的混合验证码
// 	captchaID, code, err := captcha.GetAndSave(6, captcha.AplusN, "./test.jpg")
// 	// captcha.GetOne(6,captcha.AplusN) Base64 Code
// 	// captcha.GetImage(5,captcha.AplusN)
// 	if err != nil {
// 		fmt.Println("生成验证码失败:", err)
// 		return
// 	}

// 	fmt.Println("验证码ID:", captchaID)
// 	fmt.Println("验证码图片 (base64):", code)

// 	// 验证用户输入的验证码
// 	userInput := "123456" // 假设用户输入的验证码是 123456
// 	isValid := captcha.Verify(captchaID, userInput)
// 	fmt.Println("验证码是否有效:", isValid)
// }
