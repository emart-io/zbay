package service

import (
	"context"
	"database/sql"
	"github.com/emart.io/zbay/internal/impl/biz"
	"github.com/emart.io/zbay/internal/impl/db"
	pb "github.com/emart.io/zbay/service/go"
	"github.com/gogo/protobuf/types"
)

const (
	orderTable = "orders"
)

type OrdersImpl struct{}

func (s *OrdersImpl) Add(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	in.Id = types.TimestampNow().String()
	in.Created = types.TimestampNow()
	if err := db.Insert(orderTable, in); err != nil {
		return nil, err
	}
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
	if err := db.List(orderTable, &orders, clause, "ORDER BY data->'$.created.seconds' DESC"); err != nil && err != sql.ErrNoRows {
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
	if err := db.List(orderTable, &orders, clause, "ORDER BY data->'$.created.seconds' DESC"); err != nil {
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
	if err := db.List(orderTable, &orders, clause, "ORDER BY data->'$.created.seconds' DESC"); err != nil {
		return err
	}

	for _, v := range orders {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrdersImpl) Delete(ctx context.Context, in *pb.Order) (*types.Empty, error) {
	if err := db.Delete(orderTable, in.Id); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}

func (s *OrdersImpl) SignAlipay(ctx context.Context, in *pb.Order) (*types.StringValue, error) {
	v, err := biz.SignAlipay(float32(in.Amount) / 100)
	if err != nil {
		return nil, err
	}
	return &types.StringValue{Value: v}, nil
}
