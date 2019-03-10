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
	ErrorArticleUpdateFailed: "文章更新失败",
	ErrorCommentNotExists:    "评论不存在",
	ErrorCommentCreateFailed: "评论失败",
	ErrorCommentUpdateFailed: "评论更新失败",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "登陆失败：用户名或密码错误！",
	ErrorUserNotExists:         "用户不存在",
	ErrorUsernameExists:        "用户名已经被使用",
	ErrorUserCreateFailed:      "用户注册失败",
	ErrorUserNotAdmin:          "登陆失败：您没有管理权限",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
