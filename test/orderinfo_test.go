package test

import (
	"io"
	"log"
	"os"
	"testing"

	pb "github.com/owenqing/grpcx/pb/orderinfo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// 客户端 => Server
func TestClient(t *testing.T) {
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
	// test
	// getOrderInfo(client)
	// getAll(client)
	addImage(client)
}

// 获取订单信息
func getOrderInfo(client pb.OrderInfoServiceClient) {
	response, err := client.GetOrderInfo(context.Background(), &pb.OrderInfoReq{OrderId: proto.Int64(1)})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("response: order_id[%v]\n", response.GetOrderId())
	log.Printf("response: price[%v]\n", response.GetPrice())
	log.Printf("response: desc[%v]\n", response.GetDesc())
	log.Printf("response: user_id[%v]\n", response.GetUserId())
}

// 获取所有订单信息
func getAll(client pb.OrderInfoServiceClient) {
	stream, err := client.GetAll(context.Background(), &pb.GetAllReq{})
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("res => %#v\n", response.GetDesc())
	}

}

func addImage(client pb.OrderInfoServiceClient) {
	imgFile, err := os.Open("Mars.jpg")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer imgFile.Close()
	// metadata 传递 order_id
	md := metadata.New(map[string]string{
		"order_id": "2021",
	})
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := client.AddImage(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		chunk := make([]byte, 1024)
		chunkSize, err := imgFile.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		if chunkSize < len(chunk) {
			chunk = chunk[:chunkSize]
		}
		stream.Send(&pb.AddImageReq{Data: chunk})
	}
	// 通知服务端数据已发完
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("AddImage Response: %#v\n", res.Result)
}
