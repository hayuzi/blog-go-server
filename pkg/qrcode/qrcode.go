package qrcode

import (
	"blog-go-server/pkg/setting"
	"github.com/boombuler/barcode/qr"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	ExtPng = ".png"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Ext:    ExtPng,
		Level:  level,
		Mode:   mode,
	}
}

func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath
}
