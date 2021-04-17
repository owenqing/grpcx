package test

import (
	"log"
	"testing"

	pb "github.com/owenqing/grpcx/pb/orderinfo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

func TestGetOrderInfo(t *testing.T) {

	// 客户端证书
	creds, err := credentials.NewClientTLSFromFile("../cert/cert.pem", "")
	if err != nil {
		log.Fatal(err.Error())
	}
	// 注意与 server 端区别
	options := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	const port = ":9000"
	conn, err := grpc.Dial("localhost"+port, options...)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	// rpc 调用
	client := pb.NewOrderInfoServiceClient(conn)
	response, err := client.GetOrderInfo(context.Background(), &pb.OrderInfoReq{OrderId: proto.Int64(1)})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("response: order_id[%v]\n", response.GetOrderId())
	log.Printf("response: price[%v]\n", response.GetPrice())
	log.Printf("response: desc[%v]\n", response.GetDesc())
	log.Printf("response: user_id[%v]\n", response.GetUserId())
}
