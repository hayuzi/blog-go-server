blog-go-server
===

> 学习流程参考: 
https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md 


#### 依赖
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

# 


```


#### 目录结构
```
blog-go-server
    |-- conf                配置
    |   |-- app.ini         由于配置不便于暴露, 请复制sample文件并替换配置值
    |   |-- app.sample.ini 
    |
    |-- middleware          中间件
    |
    |-- models              数据库模型
    |   |-- models.go       模型基础
    |  
    |-- pkg                 项目中的第三方包处理
    |   |-- e               自定义错误
    |       |-- code.go     错误码常量
    |       |-- msg.go      错误信息
    |-- 
    |
    |-- routers             路由
    |   |-- api
    |   |   |-- v1          版本一
    |   |
    |   |-- router.go       路由基础文件
    |
    |-- runtime
    |
    |
    |
    |
    |
    
    
    
           
```


#### 部署


