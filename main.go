package main

import (
	"log"
	"net"

	"github.com/owenqing/grpcx/interceptor"
	pb "github.com/owenqing/grpcx/pb/orderinfo"
	service "github.com/owenqing/grpcx/service/orderinfo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const port = ":9000"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// ssl 证书
	creds, err := credentials.NewServerTLSFromFile("./cert/cert.pem", "./cert/key.pem")
	if err != nil {
		log.Fatalln(err.Error())
	}
	options := []grpc.ServerOption{grpc.Creds(creds)}

	// 服务端拦截器
	// 一元添加拦截器
	options = append(options, grpc.UnaryInterceptor(interceptor.Interceptor))
	// 流拦截器
	options = append(options, grpc.StreamInterceptor(interceptor.StreamInterceptor))

	server := grpc.NewServer(options...)
	// rpc 服务注册
	pb.RegisterOrderInfoServiceServer(server, new(service.OrderInfoService))
	log.Printf("GRPC SERVER STARTED %s\n", port)
	server.Serve(listen)
}
