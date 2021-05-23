package test

import (
	"io"
	"log"
	"os"
	"testing"
	"time"

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
	getOrderInfo(client)
	// getAll(client)
	// addImage(client)
	// AddBatchUsers(client)
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
		chunk := make([]byte, 10*1024)
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
	log.Printf("AddImage Response: %#v\n", res.GetResult())
}

// 双向流批量上传 user 信息
// 服务端每处理一个请求就返回一个 user_id
func AddBatchUsers(client pb.OrderInfoServiceClient) {
	var users = []*pb.User{
		{UserId: proto.Int64(1), Name: proto.String("Tom"), OrderNum: proto.Int32(10), ConsumptionAmount: proto.Int32(1000)},
		{UserId: proto.Int64(2), Name: proto.String("Jack"), OrderNum: proto.Int32(9), ConsumptionAmount: proto.Int32(899)},
		{UserId: proto.Int64(3), Name: proto.String("Bob"), OrderNum: proto.Int32(6), ConsumptionAmount: proto.Int32(998)},
	}
	stream, err := client.AddBatchUsers(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	// goruntine 去处理服务端返回的结果
	finishSignal := make(chan struct{}, 1)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				finishSignal <- struct{}{}
				break
			}
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Printf("user rsp => %#v\n", res.GetUserId())
		}
	}()

	for _, u := range users {
		err = stream.Send(u)
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(500 * time.Millisecond)
	}
	stream.CloseSend()
	<-finishSignal
}
