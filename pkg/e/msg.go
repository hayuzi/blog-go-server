package e

var MsgFlags = map[int]string{
	Success:                     "ok",
	Error:                       "fail",
	InvalidParams:               "请求参数错误",
	ErrorUploadCheckImageFormat: "校验图片错误，图片格式错误或大小超出限制",
	ErrorUploadCheckImageFail:   "检查图片失败",
	ErrorUploadSaveImageFail:    "图片保存失败",

	ErrorTagExists:           "已存在该标签名称",
	ErrorTagNotExists:        "该标签不存在",
	ErrorTagCreateFailed:     "标签创建失败",
	ErrorTagUpdateFailed:     "标签更新失败",
	ErrorArticleNotExists:    "该文章不存在",
	ErrorArticleCreateFailed: "文章新增失败",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
