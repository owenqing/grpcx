package orderinfo

import (
	"io"
	"log"
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
		&pb.OrderInfoRsp{OrderId: proto.Int64(1), Price: proto.String("2$"), Desc: proto.String("蔬菜"), UserId: proto.Int64(1000)},
		&pb.OrderInfoRsp{OrderId: proto.Int64(2), Price: proto.String("18$"), Desc: proto.String("水果"), UserId: proto.Int64(1009)},
		&pb.OrderInfoRsp{OrderId: proto.Int64(3), Price: proto.String("16$"), Desc: proto.String("肉类"), UserId: proto.Int64(1009)},
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
		log.Printf("AddImage metatada: %#v\n", md["order_id"])
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
