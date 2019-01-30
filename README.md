blog-go-server
===

> 学习流程参考: 
https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md 

#### 在go的版本 1.11 以后，加入了  modules 管理
我们可以使用 go mod 很方便的维护依赖


#### 包版本的依赖下载， 在项目目录下下载依赖切换到master分支
```
# gin (http服务框架)
go get -u github.com/gin-gonic/gin

# go-int (配置管理)
go get -u github.com/go-ini/ini

# com 依赖包   工具类
go get -u github.com/Unknwon/com

# Gorm 与 数据库管理依赖
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm

# validation 数据验证器
go get -u github.com/astaxie/beego/validation

# jwt 验证包
go get -u github.com/dgrijalva/jwt-go

# 进程管理包（ 服务平滑启动关闭管控， 目前不使用 ）
# go get -u github.com/fvbock/endless

# swaggo 文档管理（ 如果没有科学上网下载不了的化，可考虑 gopm ）
go get -u github.com/swaggo/swag/cmd/swag
# 若 $GOPATH/bin 没有加入$PATH中，你需要执行将其可执行文件移动到$GOBIN下
# mv $GOPATH/bin/swag /usr/local/go/bin

# 或者使用gopm下载swaggo
gopm get -g -v github.com/swaggo/swag/cmd/swag
cd $GOPATH/src/github.com/swaggo/swag/cmd/swag
go install

# gin-swaggo
gopm get -g -v  github.com/swaggo/gin-swagger
go get -u github.com/alecthomas/template 

# 如果需要定时任务包的话，可以使用下面这个包
go get -u github.com/robfig/cron



```




#### 目录结构
```
blog-go-server
    |-- conf                    配置
    |   |-- app.ini                 由于配置不便于暴露, 请复制sample文件并替换配置值
    |   |-- app.sample.ini 
    |
    |-- middleware              中间件
    |
    |-- models                  数据库模型
    |   |-- article.go              文章表model
    |   |-- models.go               模型基础
    |   |-- tag.go                  标签表model
    |  
    |-- pkg                     项目中的第三方包处理
    |   |-- constmap            自定义常量
    |   |   |-- common.go           通用常量    
    |   |      
    |   |-- e                   自定义错误
    |   |   |-- code.go             错误码常量
    |   |   |-- msg.go              错误信息
    |   |
    |   |-- setting             配置加载
    |   |   |-- setting.go
    |   |
    |   |
    |   |-- util                工具类
    |   |   |-- jsontime.go         Gorm中需要用到的自定义时间格式  
    |   |   |-- pagination.go       jwt验证类
    |   |   |-- pagination.go       分页默认参数  
    |
    |
    |-- routers                 路由
    |   |-- api
    |   |   |-- v1              版本一
    |   |   |   |-- article.go      文章控制器        
    |   |   |   |-- tag.go          标签控制器        
    |   |   |               
    |   |   |-- v2              版本二
    |   |
    |   |-- router.go           路由基础文件
    |
    |-- runtime
    |
    |
    |
    |
    |
    
    
    
           
```

---
### 部署

##### 使用 Dockerfile 容器打包镜像
```
# 打包linux环境下的可执行文件
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o blog-go-server .

# 构建docker镜像 
## hayuzi/blog-go-server-scratch:1.0.0 表示的是 <仓库/名称:版本tag> 
## hayuzi是本人仓库, 个人构建的时候只需要 <[仓库名/]名称[:版本tag]>

docker build -t hayuzi/blog-go-server-scratch:1.0.0 .

# 运行docker镜像  ( 需要将项目实际配置所在目录挂载进去, 并将日志所在文件夹暴露出来 ) 日志目录挂载好像暂时有问题 
docker run --name=mygoblog -p 8000:8000 -v $GOPATH/src/blog-go-server/conf:/data/blog/conf -v /$GOPATH/src/blog-go-server/runtime:/data/blog/runtime  hayuzi/blog-go-server-scratch:1.0.0

# 停止并删除容器
docker stop mygoblog && docker rm mygoblog

# 删除镜像
docker rmi hayuzi/blog-go-server-scratch:1.0.0



```

