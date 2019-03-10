package e

const (
	// 通用
	Success                     = 200
	Error                       = 500
	InvalidParams               = 400
	ErrorUploadCheckImageFormat = 100001
	ErrorUploadCheckImageFail   = 100002
	ErrorUploadSaveImageFail    = 100003

	// 文章相关
	ErrorTagExists           = 200001
	ErrorTagNotExists        = 200002
	ErrorTagCreateFailed     = 200003
	ErrorTagUpdateFailed     = 200004
	ErrorArticleNotExists    = 200005
	ErrorArticleCreateFailed = 200006
	ErrorArticleUpdateFailed = 200007
	ErrorCommentNotExists    = 200008
	ErrorCommentCreateFailed = 200009
	ErrorCommentUpdateFailed = 200010

	// 用户与授权相关
	ErrorAuthCheckTokenFail    = 210001
	ErrorAuthCheckTokenTimeout = 210002
	ErrorAuthToken             = 210003
	ErrorAuth                  = 210004
	ErrorUserNotExists         = 210005
	ErrorUsernameExists        = 210006
	ErrorUserCreateFailed      = 210007
	ErrorUserNotAdmin      	   = 210008
)
