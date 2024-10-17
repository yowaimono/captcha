# captchp

`captchp` 是一个用于生成和验证验证码的 Go 包。它支持生成不同格式的验证码图片，并提供验证码的验证功能。

## 功能

- 生成随机验证码图片，支持多种格式（大小写字母和数字混合、大写字母和数字混合、小写字母和数字混合、大写和小写字母混合）。
- 将生成的验证码图片编码为 base64 格式。
- 存储验证码信息，并在验证码过期后自动删除。
- 提供验证码验证功能。

## 安装

使用 `go get` 命令安装：

```bash
go get github.com/yowaimono/captcha

```

## 示例

```go
package main

import (
	"fmt"
	"github.com/yowaimono/captcha"
)

func main() {
	// 生成一个长度为 6 的混合验证码
	captchaID, imgBase64, err := captcha.GetOne(6, captcha.Mixed)
	if err != nil {
		fmt.Println("生成验证码失败:", err)
		return
	}

	fmt.Println("验证码ID:", captchaID)
	fmt.Println("验证码图片 (base64):", imgBase64)

	// 验证用户输入的验证码
	userInput := "123456" // 假设用户输入的验证码是 123456
	isValid := captcha.Verify(captchaID, userInput)
	fmt.Println("验证码是否有效:", isValid)
}
```