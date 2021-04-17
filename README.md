## Grpc 学习记录

#### 如何运行
`source env.sh` source 环境变量，x509 标准证书在 go 中有点问题，需要调整下 GODEBUG 环境变量，详情见 env.sh
`make ssl` 在本地生成证书
`make build` 编译
`./bin/grpcx` 运行 server

