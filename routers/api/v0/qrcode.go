package v0

import (
	"blog-go-server/pkg/qrcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
)

// QrCode方法渲染生成二维码并且直接以png格式输出图片到客户端
func QrCode(c *gin.Context) {
	content := c.Query("content")
	level := qr.Q
	mode := qr.Auto
	code := qrcode.NewQrCode(content, 200, 200, level, mode)
	code.EncodeForGinResponse(c)
}
