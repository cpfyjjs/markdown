## go语言常用命令

直接在终端输入`go help`即可显示所有的go命令以及相应命令功能简介

### 1.相关命令简介

`go build `:编译包和依赖

`go clean`:移除对象文件

`go doc`:显示包或者符号的文档

`go env`：显示go 的环境信息

`go bug`:启动错误报告

`go fix`：运行go tool fix

`go fmt`:运行gofmt进行格式化

`go generate`:从processing source 生成go文件

`go get`：下载并安装包和依赖

`go install`:编译并安装包和依赖

`go list`：列出包

`go run`:编译并运行go程序

`go test`:运行测试

`go tool`：运行go提供的工具

`go version`：显示go的版本

`go vet`：运行go tool vet

命令的使用方式为：`go command [args]`,除此之外，可以使用go help <command>来显示指定命令的更多帮助信息。在运行go help时，不仅仅打印了这些命令的基本信息，还给出了一些概念的版主信息：

`c`：Go和C的相互调用

`bulidmode`：构建模式的描述

`filetype`：文件类型

`gopath`：GPOPATH环境变量

`enviroment`：环境变量

`importpath`：导入路径变量

`packages`：包列表描述

`testflag`：测试符号描述

`testfunc`：测试函数描述

同样可以使用`go help <topic>`来查看这些概念信息

### 2.build和run命令

就像其他静态类型的语言一样，要执行go程序，需要先编译，然后在执行产生的可执行文件。`go bulid`命令就是用来编译go程序生成可执行文件的。但不是所有的go程序够可以编译生成可执行程序，要生成可执行程序，go程序必须满足两个条件：

* 该go程序需要属于main包
* 在main包中必须还得包含main函数

而`go run`将会直接运行main程序，而不产生中间文件

此外`go clean`命令，可以用于清除产生的可执行程序

### 3.fmt和doc命令

go语言对格式要求严格，这样可以保持代码的清晰一致，编译组合开发，并且go还提供以个非常强大的工具来格式化代码，他就是`go fmt sourcefile.go`，不过通常不需要手调，编译器自动调整。

`go doc`名命可以方便我们快速产看文档，`go doc package`命令将会在终端打印指定的package文档

`godoc -http :8080`：启动文档服务器

### 4.get命令



