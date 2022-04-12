# BloodPressureMeasurementRecordApplet

## 高血压测量记录小程序后端
采用Golang开发，使用了Gin, Gorm, Viper, zap, jwt, go.uuid, sys, fsnotify, lumberjack.v2, goini 等系列库进行开发

项目分层结构参考：https://github.com/xmgtony/apiserver-gin（一个基于gin的go生产及服务端）

### 代码结构介绍：
example/
Gin的标准案例

internal/
网路接口层，大量调用pkg代码，在这一层中分为model, repo, service, handler, router 逐级向上提供接口，也存在middleware 中间层

pkg/
包代码，包含了整个项目的逻辑架构分层

tools/
工具类，均是简单封装直接使用的方法，不需要实例创建

resource/
资源类，包含资源文件，图片，sql等

config/
配置文件，采用.ini创建，由于采用了viper库，也可以修改成各种常见的配置文件样式，比如json，yaml等

server/
http的路由加载内容

### 业务记录





## 小程序前端
见：XXX 