package orderinfo

import (
	pb "gitee.com/owenqing/grpcx/pb/orderinfo"

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
