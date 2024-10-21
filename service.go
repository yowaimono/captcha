package captcha

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Service 处理生成验证码的请求
func GinService(c *gin.Context) {
	// 从请求中获取参数
	lengthStr := c.DefaultQuery("length", "6")
	formatStr := c.DefaultQuery("format", "numeric")
	savePath := c.DefaultQuery("savePath", "")

	// 解析验证码长度
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid length parameter"})
		return
	}

	// 解析验证码格式
	var format CaptchaFormat
	switch formatStr {
	case "numeric":
		format = Numeric
	case "alphanumeric":
		format = AlphaNumeric
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format parameter"})
		return
	}

	// 根据 savePath 参数决定调用哪个函数
	var captchaID, captchaData string
	var img image.Image
	if savePath != "" {
		// 生成验证码并保存到指定路径
		captchaID, captchaData, err = GetAndSave(length, format, savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		// 生成验证码并返回 base64 编码的图片
		captchaID, captchaData, err = GetBase64(length, format)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// 返回验证码ID和验证码数据
	c.JSON(http.StatusOK, gin.H{
		"captchaID": captchaID,
		"captchaData": captchaData,
	})
}





// Service 处理生成验证码的请求
func FiberService(c *fiber.Ctx) error {
	// 从请求中获取参数
	lengthStr := c.Query("length", "6")
	formatStr := c.Query("format", "numeric")
	savePath := c.Query("savePath", "")

	// 解析验证码长度
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid length parameter"})
	}

	// 解析验证码格式
	var format CaptchaFormat
	switch formatStr {
	case "numeric":
		format = Numeric
	case "alphanumeric":
		format = AlphaNumeric
	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid format parameter"})
	}

	// 根据 savePath 参数决定调用哪个函数
	var captchaID, captchaData string
	var img image.Image
	if savePath != "" {
		// 生成验证码并保存到指定路径
		captchaID, captchaData, err = GetAndSave(length, format, savePath)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	} else {
		// 生成验证码并返回 base64 编码的图片
		captchaID, captchaData, err = GetBase64(length, format)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	// 返回验证码ID和验证码数据
	return c.JSON(fiber.Map{
		"captchaID":   captchaID,
		"captchaData": captchaData,
	})
}