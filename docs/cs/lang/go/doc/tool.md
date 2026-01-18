# tool

## Go中的环境变量

>- GOROOT 默认为/usr/local/go
>- GOPATH 默认为~/go
>- GOBIN 默认为~/go/bin

## IDE

VSCode

> 1. 安装vscode-go插件
> 2. 安装go tool工具
  > 自动分析
  > 手动安装

## go tool

>- godef                    自动跳转
>- gocode                   自动补全
>- golint                   语法检查
>- gorename                 重命名
>- go-outline               文件大纲
>- go-symbols               工作区符号搜索
>- go tool vet              语法拼写建议
>- gotests                  生成自动化测试代码

### 格式化

>- gofmt                    格式化
>- goformat                 格式化
>- goreturns                格式化, 自动导包  
>- goimports                格式化, 自动导包

### 文档

>- godoc                    代码文档
>- guru                     代码文档
>- gogetdoc                 代码文档

### 包

>- gopkgs                   列出引用的包
>- go-find-references       列出引用的包

### Debug

>- goreportcard             质量检查
>- dlv                      调试
>- goreplay                 真实数据压测

### 包管理工具

go mod(vgo)

v(major).(minor).(patch)

1. 在GOPATH外的话直接用, 或者GOMODULE=auto, unset
2. 在GOPATH需要手动设置, GOMODULE=on
