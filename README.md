# NeteaseCloudMusicApiWithGo

本项目基于[网易云音乐 API](https://github.com/Binaryify/NeteaseCloudMusicApi) 重新用golang实现  
网易云api Golng 版(持续更新)

## 功能特性
1. 登录
2. 刷新登录
3. 发送验证码
4. 校验验证码
5. 注册(修改密码)  
6. 。。。等160多个api

## 环境要求

需要 Golang 1.6以上 环境


## 运行

```shell
go run main.go
```

## 使用文档（为binaryify大佬的文档）

[文档地址](https://binaryify.github.io/NeteaseCloudMusicApi) 

[文档地址2](https://neteasecloudmusicapi.vercel.app)

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init singo
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```


项目运行后启动在3333端口（可以修改，参考gin文档)