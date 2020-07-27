# NeteaseCloudMusicApiWithGo

本项目基于[网易云音乐 API](https://github.com/Binaryify/NeteaseCloudMusicApi) 重新用golang实现  
网易云api Golng 版（开发中已完成基本的请求解码编码模块）  


## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init go-crud
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在3333端口（可以修改，参考gin文档)