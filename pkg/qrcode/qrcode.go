package qrcode

import (
	"blog-go-server/pkg/file"
	"blog-go-server/pkg/setting"
	"blog-go-server/pkg/util"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode"
	"image/png"
	"github.com/gin-gonic/gin"
)

type QrCode struct {
	Content string
	Width   int
	Height  int
	Ext     string
	Level   qr.ErrorCorrectionLevel
	Mode    qr.Encoding
}

const (
	ExtPng = ".png"
)

func NewQrCode(content string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		Content: content,
		Width:   width,
		Height:  height,
		Ext:     ExtPng,
		Level:   level,
		Mode:    mode,
	}
}

func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath
}

func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.QrCodeSavePath
}

func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}

func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.Content) + q.GetQrCodeExt()
	if file.CheckNotExist(src) == true {
		return false
	}
	return true
}

func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.Content) + q.GetQrCodeExt()
	src := path + name
	if file.CheckNotExist(src) == true {
		code, err := qr.Encode(q.Content, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = png.Encode(f, code)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}


func (q *QrCode) EncodeForDirect (c *gin.Context) (string, error) {
	code, err := qr.Encode(q.Content, q.Level, q.Mode)
	if err != nil {
		return "", err
	}
	code, err = barcode.Scale(code, q.Width, q.Height)
	if err != nil {
		return "", err
	}
	err = png.Encode(c.Writer, code)
	if err != nil {
		return "",  err
	}
	return "", nil
}
