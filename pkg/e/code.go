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
	ErrorTagExists        = 200001
	ErrorTagNotExists     = 200002
	ErrorTagCreateFailed  = 200003
	ErrorTagUpdateFailed  = 200004
	ErrorArticleNotExists = 200005
	ErrorArticleAddFailed = 200005

	// 用户与授权相关
	ErrorAuthCheckTokenFail    = 210001
	ErrorAuthCheckTokenTimeout = 210002
	ErrorAuthToken             = 210003
	ErrorAuth                  = 210004
)
