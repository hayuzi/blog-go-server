package v1

import (
	"blog-go-server/pkg/app"
	"blog-go-server/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"fmt"
)

func Dashboard(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]string)
	data["os"] = runtime.GOOS
	data["arch"] = runtime.GOARCH
	data["version"] = runtime.Version()
	data["cups"] = fmt.Sprintf("%d", runtime.NumCPU())
	appG.Response(http.StatusOK, e.Success, data)
}
