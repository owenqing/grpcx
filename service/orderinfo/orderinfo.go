package orderinfo

import (
	"io"
	"log"
	"strconv"
	"time"

	pb "github.com/owenqing/grpcx/pb/orderinfo"
	"google.golang.org/grpc/metadata"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type OrderInfoService struct{}

// 获取订单信息
func (s *OrderInfoService) GetOrderInfo(ctx context.Context, request *pb.OrderInfoReq) (*pb.OrderInfoRsp, error) {
	orderId := request.GetOrderId()
	// 返回
	response := &pb.OrderInfoRsp{
		OrderId: proto.Int64(orderId),
		Price:   proto.String("200$"),
		Desc:    proto.String("肉类"),
		UserId:  proto.Int64(2),
	}
	return response, nil
}

// 获取全部订单信息
// 服务端流模式
func (s *OrderInfoService) GetAll(request *pb.GetAllReq, stream pb.OrderInfoService_GetAllServer) error {
	var order = []*pb.OrderInfoRsp{
		{OrderId: proto.Int64(1), Price: proto.String("2$"), Desc: proto.String("蔬菜"), UserId: proto.Int64(1000)},
		{OrderId: proto.Int64(2), Price: proto.String("18$"), Desc: proto.String("水果"), UserId: proto.Int64(1009)},
		{OrderId: proto.Int64(3), Price: proto.String("16$"), Desc: proto.String("肉类"), UserId: proto.Int64(1009)},
	}
	for _, orderItem := range order {
		// 流模式 send
		stream.Send(orderItem)
		time.Sleep(2 * time.Second)
	}
	return nil
}

// 订单图片上传
func (s *OrderInfoService) AddImage(stream pb.OrderInfoService_AddImageServer) error {
	// 订单 id 从 metadata 中取出
	// metadata 是一个 map
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		orderId, _ := strconv.Atoi(md["order_id"][0])
		log.Printf("AddImage metatada: %#v\n", orderId)
	}
	img := []byte{}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Image Size: %#v\n", len(img))
			// 上传完毕给客户端发送请求
			return stream.SendAndClose(&pb.AddImageRsp{Result: proto.Bool(true)})
		}
		if err != nil {
			// 这个 err 会被序列化直接传给客户端
			return err
		}
		log.Printf("received size: %d\n", len(data.Data))
		img = append(img, data.Data...)
	}
}

func (s *OrderInfoService) AddBatchUsers(stream pb.OrderInfoService_AddBatchUsersServer) error {
	var usersRet []*pb.User
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		usersRet = append(usersRet, user)
		// 处理一个数据就返回 user_id
		stream.Send(&pb.AddBatchUsersRsp{UserId: user.UserId})
	}
	// 输出服务端收到的 users
	for _, u := range usersRet {
		log.Printf("user => %+v\n", u)
	}
	return nil
}
