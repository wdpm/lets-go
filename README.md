# lets-go
Learn Go.

## 安装
- [Installing Go](https://golang.org/doc/install)

## 配置
- [Setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
- [Go Proxy](https://goproxy.io/)
```bash
go env -w GO111MODULE=on
go env -w GOPROXY="https://goproxy.io,direct"
```

## 入门
> 参阅 [A Tour of Go Excercises](https://tour.go-zh.org/)
- 基础
  - 函数 func
  - 包 package
  - 变量 var
  - 基本类型 bool string int? uint? byte rune float? complex?
  - 流程控制 if for switch defer
  - 指针地址 &i 和指针读值 *p
  - 结构体 struct
  - 数组 var a [2]string
  - 数组切片 s[1:4]，长度len(s)和容量cap(s)，0值nil，make创建，range遍历
  - 映射 map
- 方法、接口、其他
  - 方法 method，定义中有调用者
  - 接口 interface
  - 错误 error
  - 图像 image
- 并发




