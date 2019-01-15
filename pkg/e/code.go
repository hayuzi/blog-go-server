package e

const (
	// 通用
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// 文章相关
	ErrorTagExists        = 10001
	ErrorTagNotExists     = 10002
	ErrorArticleNotExists = 10003

	// 用户与授权相关
	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004
)
