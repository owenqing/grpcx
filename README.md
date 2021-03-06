## Grpc 学习

### 安装环境
* 安装 protoc 编译器
* 安装 protoc-gen-go 并为该目录配置环境变量

### 如何运行
* `source env.sh` source 环境变量，x509 标准证书在 go 中有点问题，需要调整下 GODEBUG 环境变量，详情见 env.sh
* `make ssl` 在本地生成证书
* `make build` 编译
* `./bin/grpcx` 运行 server
* 执行客户端测试程序 ./test `go test -v` (可以修改 TestClient 函数中测试的接口)

### 教程
已订单管理为例，实现 grpc 四种模式。一元模式，服务端流，客户端流，双向流
1. 在 proto 目录下编写 pb 文件 ([什么是 pb 文件](https://colobu.com/2015/01/07/Protobuf-language-guide/))

2. 使用 `make update` 命令编译 pb 文件，生成 go 语言代码 => pb 目录 
    * 将编译 proto 文件的命令集成到 Makefile 中 
    * `protoc --go_out=plugins=grpc:./pb ./proto/*.proto`
    * 编译 grpc 的 proto 需要指明 grpc 插件
    * go_out 设置编译生成的文件路径，这里文件生成到 pb 目录中
    * 最后指明需要编译的 proto 文件，这里可以使用通配符

3. 编写 service 接口, (实现放在 ./service 目录)
    * 需要建立一个 pb 文件中写好的 service struct 并实现它
    * 写好后到 main 函数中注册
    * service 可以有多个

4. test 文件中放的是 service 接口的客户端请求
    * cd 到 test 目录下，执行 `go test -v`（跑相关用例）

