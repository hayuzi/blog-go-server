package v0

import (
	"github.com/gin-gonic/gin"
	"blog-go-server/pkg/e"
	"blog-go-server/pkg/logging"
	"net/http"
	"blog-go-server/pkg/upload"
	"blog-go-server/pkg/app"
)

func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, e.Error, data)
		return
	}

	if image == nil {
		appG.Response(http.StatusOK, e.InvalidParams, data)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if ! upload.CheckImageExt(imageName) || ! upload.CheckImageSize(file) {
		appG.Response(http.StatusOK, e.ErrorUploadCheckImageFormat, data)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, e.ErrorUploadCheckImageFail, data)
		return
	} else if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, e.ErrorUploadSaveImageFail, data)
		return
	} else {
		data["image_url"] = upload.GetImageFullUrl(imageName)
		data["image_save_url"] = savePath + imageName
	}

	appG.Response(http.StatusOK, e.Success, data)
}
