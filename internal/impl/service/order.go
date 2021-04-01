package service

import (
	"context"
	"database/sql"

	"github.com/emart.io/zbay/internal/impl/biz"
	"github.com/emart.io/zbay/internal/impl/biz/db"
	pb "github.com/emart.io/zbay/service/go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	orderTable = "orders"
)

type OrdersImpl struct {
	pb.UnimplementedOrdersServer
}

func (s *OrdersImpl) Add(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	in.Id = timestamppb.Now().String()
	in.Created = timestamppb.Now()
	if err := db.Insert(orderTable, in); err != nil {
		return nil, err
	}
	// send message to seller
	if in.Status == "待发货" {
		pb.NewMessagesClient(biz.Connection()).Add(ctx, &pb.Message{
			From:    in.UserId,
			To:      in.Snapshot.OwnerId,
			Content: "[系统自动发送]:用户" + in.UserId + "，已经购买" + in.Snapshot.Title + "，请注意安排发货哦~",
		})
	}
	biz.SendSMS(in.Snapshot.OwnerId, in.Snapshot.Title, in.Destination.Location)
	return in, nil
}

func (s *OrdersImpl) Get(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	order := pb.Order{}
	if err := db.GetById(orderTable, in.Id, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *OrdersImpl) Update(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	order, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	if in.Status != "" {
		if order.Status == "待收货" && in.Status == "待评价" {
			pb.NewAccountsClient(biz.Connection()).Add(ctx, &pb.Account{
				UserId:  order.Snapshot.OwnerId,
				Amount:  int64(order.Amount),
				OrderId: order.Id,
			})
		}
		if order.Status == "待退款" && in.Status == "已退款" {
			pb.NewAccountsClient(biz.Connection()).Add(ctx, &pb.Account{
				UserId:  order.Snapshot.OwnerId,
				Amount:  int64(-order.Amount),
				OrderId: order.Id,
			})
			pb.NewAccountsClient(biz.Connection()).Add(ctx, &pb.Account{
				UserId:  order.UserId,
				Amount:  int64(order.Amount),
				OrderId: order.Id,
			})
		}
		// send message to seller
		if in.Status == "待发货" {
			conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
			pb.NewMessagesClient(conn).Add(ctx, &pb.Message{
				From:    order.UserId,
				To:      order.Snapshot.OwnerId,
				Content: "[系统自动发送]:用户" + order.UserId + "，已经购买" + order.Snapshot.Title + "，请注意安排发货哦~",
			})
		}
		order.Status = in.Status
	}
	if in.Express != nil {
		order.Express = in.Express
	}

	if in.Groupon != nil {
		order.Groupon = in.Groupon
	}

	if in.Comment != "" {
		order.Comment = in.Comment
	}

	if err := db.Update(orderTable, in.Id, order); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *OrdersImpl) ListByOrder(in *pb.Order, stream pb.Orders_ListByOrderServer) error {
	orders := []*pb.Order{}
	clause := "WHERE data->'$.snapshot.id'='" + in.Snapshot.Id + "' AND data->'$.status'='" + in.Status + "'"
	if err := db.List(orderTable, &orders, clause, "ORDER BY data->'$.created' DESC"); err != nil && err != sql.ErrNoRows {
		return err
	}

	for _, v := range orders {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrdersImpl) ListForBuyer(in *pb.ListQuery, stream pb.Orders_ListForBuyerServer) error {
	clause := "WHERE data->'$.status' LIKE '%" + in.Status + "%'"
	if !biz.IsAdmin(in.User) {
		clause = clause + " AND data->'$.userId'='" + in.User.Id + "'"
	}
	orders := []*pb.Order{}
	if err := db.List(orderTable, &orders, clause, "ORDER BY data->'$.created' DESC"); err != nil {
		return err
	}

	for _, v := range orders {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrdersImpl) ListForSeller(in *pb.ListQuery, stream pb.Orders_ListForSellerServer) error {
	clause := "WHERE data->'$.status' LIKE '%" + in.Status + "%'"
	if !biz.IsAdmin(in.User) {
		clause = clause + " AND data->'$.snapshot.ownerId'='" + in.User.Id + "'"
	}
	orders := []*pb.Order{}
	if err := db.List(orderTable, &orders, clause, "ORDER BY data->'$.created' DESC"); err != nil {
		return err
	}

	for _, v := range orders {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrdersImpl) Delete(ctx context.Context, in *pb.Order) (*emptypb.Empty, error) {
	if err := db.Delete(orderTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
