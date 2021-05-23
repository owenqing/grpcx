package interceptor

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 一元拦截器
func Interceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// 服务端拦截器
	// ===============================
	// 在此处编写拦截类容
	// ===============================
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("拦截元数据: %#v\n", md)
	}
	log.Printf("服务端一元拦截器触发...")
	return handler(ctx, req)
}

// 流拦截器
// 三种流都走这个拦截器
func StreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	// 服务端拦截器
	// ===============================
	// 在此处编写拦截类容
	// ===============================
	log.Println("服务端流拦截器触发...")
	return handler(srv, ss)
}
