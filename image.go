// captcha_image.go
package captcha

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// 创建验证码图片
func createCaptchaImage(code string, noiseLevel NoiseLevel) image.Image {
	// 根据验证码的长度动态调整图片宽度
	charWidth := 20 // 假设每个字符需要20像素的宽度
	width := len(code) * charWidth + 20 // 加上左右边距
	height := 40
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 填充背景色
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	// 加载字体
	ttfFont, err := opentype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	// 创建字体面
	face, err := opentype.NewFace(ttfFont, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	// 绘制验证码
	drawCaptcha(img, code, face)

	if noiseLevel == Simple {
		return img
	}
	// 添加噪点
	addNoise(img, noiseLevel)

	// 添加干扰线
	addLines(img)

	return img
}

// 绘制验证码
func drawCaptcha(img *image.RGBA, code string, face font.Face) {
	_, height := img.Bounds().Dx(), img.Bounds().Dy()
	rand.Seed(time.Now().UnixNano())

	for i, char := range code {
		// 随机颜色
		col := color.RGBA{
			uint8(rand.Intn(256)),
			uint8(rand.Intn(256)),
			uint8(rand.Intn(256)),
			255,
		}

		// 随机倾斜角度
		angle := rand.Float64()*20 - 10 // 倾斜角度在 -10 到 10 度之间
		rad := angle * math.Pi / 180

		// 计算字符位置
		x := 10 + i*20
		y := height/2 + 10

		// 旋转字符
		drawRotatedChar(img, char, x, y, rad, col, face)
	}
}

// 绘制旋转字符
func drawRotatedChar(img *image.RGBA, char rune, x, y int, rad float64, col color.RGBA, face font.Face) {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()

	// 创建一个新的图像来绘制旋转后的字符
	rotatedImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// 绘制字符到新的图像
	d := &font.Drawer{
		Dst:  rotatedImg,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  fixed.P(x, y),
	}
	d.DrawString(string(char))

	// 旋转图像
	rotatedImg = rotateImage(rotatedImg, rad)

	// 将旋转后的图像绘制到原始图像
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c := rotatedImg.At(i, j)
			if _, _, _, a := c.RGBA(); a > 0 {
				img.Set(i, j, c)
			}
		}
	}
}

// 旋转图像
func rotateImage(img *image.RGBA, rad float64) *image.RGBA {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	centerX, centerY := width/2, height/2

	rotatedImg := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// 计算旋转后的坐标
			xx := int(float64(x-centerX)*math.Cos(rad)-float64(y-centerY)*math.Sin(rad)) + centerX
			yy := int(float64(x-centerX)*math.Sin(rad)+float64(y-centerY)*math.Cos(rad)) + centerY

			if xx >= 0 && xx < width && yy >= 0 && yy < height {
				rotatedImg.Set(xx, yy, img.At(x, y))
			}
		}
	}

	return rotatedImg
}

// 添加噪点
func addNoise(img *image.RGBA, noiseLevel NoiseLevel) {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	rand.Seed(time.Now().UnixNano())

	var noiseCount int
	switch noiseLevel {
	case Simple:
		// Simple 级别几乎不做噪声处理
		return
	case Mid:
		noiseCount = width * height / 50
	case Hard:
		noiseCount = width * height / 30
	default:
		noiseCount = width * height / 50
	}

	for i := 0; i < noiseCount; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		img.Set(x, y, color.RGBA{0, 0, 0, 255})
	}
}

// 添加干扰线
func addLines(img *image.RGBA) {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		x1 := rand.Intn(width)
		y1 := rand.Intn(height)
		x2 := rand.Intn(width)
		y2 := rand.Intn(height)
		drawLine(img, x1, y1, x2, y2, color.RGBA{0, 0, 0, 255})
	}
}

// 绘制直线
func drawLine(img *image.RGBA, x1, y1, x2, y2 int, col color.RGBA) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx, sy := 1, 1
	if x1 >= x2 {
		sx = -1
	}
	if y1 >= y2 {
		sy = -1
	}
	err := dx - dy

	for {
		img.Set(x1, y1, col)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}