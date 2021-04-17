## Grpc 学习记录

#### 如何运行
* `source env.sh` source 环境变量，x509 标准证书在 go 中有点问题，需要调整下 GODEBUG 环境变量，详情见 env.sh
* `make ssl` 在本地生成证书
* `make build` 编译
* `./bin/grpcx` 运行 server

### 教程
已订单管理为例，实现 grpc 四种模式。普通模式，服务端流，客户端流，双向流

* service 实现放在 ./service 目录
* ./test 目录下为 client 端的实现，通过 `go test -v` 运行
* `make update` 更新改动的 pb 数据