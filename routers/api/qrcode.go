package api

import (
	"github.com/gin-gonic/gin"
	"blog-go-server/pkg/qrcode"
	"github.com/boombuler/barcode/qr"
)

func QrCode(c *gin.Context) {
	content := c.Query("content")
	level := qr.Q
	mode := qr.Auto
	code := qrcode.NewQrCode(content, 200, 200, level, mode)
	code.EncodeForDirect(c)
}