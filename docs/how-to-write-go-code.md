# 如何写 Go 代码

## 代码组织
- A package is a collection of source files in the same directory that are compiled together.
- A module is a collection of related Go packages that are released together.
- A file named `go.mod` located at the root of the repository declares the module path: 
the import path prefix for all packages within the module.
- the module `github.com/google/go-cmp` contains a package in the directory `cmp/`. 
That package's import path is `github.com/google/go-cmp/cmp`.
- Packages in the standard library do not have a module path prefix.

## First Program
> 请保证全局 GOPATH 已配置

新建文件夹hello/，目录内容
```
│  hello.go
└─ morestrings/
        reverse.go
```
- 进入 morestrings/ 目录，编辑 reverse.go
  ```go
  package morestrings
  
  func ReverseRunes(s string) string {
  	r := []rune(s)
  	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
  		r[i], r[j] = r[j], r[i]
  	}
  	return string(r)
  }
  ```
  执行 `go build`

- 进入 hello/ 目录，编辑 hello.go 内容
  ```go
  package main
  
  import (
  	"example.com/user/hello/morestrings"
  	"fmt"
  	"github.com/google/go-cmp/cmp"
  )
  
  func main() {
  	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
  	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
  }
  ```
  执行 `go install example.com/user/hello`，留意生成的 \hello\go.mod 文件。
  
- 测试 hello
  ```bash
  G:\lets-go\hello>hello
  Hello, Go!
    string(
  -       "Hello World",
  +       "Hello Go",
    )
  ```

## 测试