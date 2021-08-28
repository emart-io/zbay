package service

import (
	"context"

	pb "github.com/emart.io/zbay/service/go"
	"github.com/jmzwcn/db"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	couponTable = "coupons"
)

type CouponImpl struct {
	pb.UnimplementedCouponsServer
}

func (s *CouponImpl) Add(ctx context.Context, in *pb.Coupon) (*pb.Coupon, error) {
	in.Id = timestamppb.Now().String()
	in.Created = timestamppb.Now()
	if err := db.Insert(couponTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *CouponImpl) Get(ctx context.Context, in *pb.Coupon) (*pb.Coupon, error) {
	coupon := pb.Coupon{}
	if err := db.GetById(couponTable, in.Id, &coupon); err != nil {
		return nil, err
	}
	return &coupon, nil
}

func (s *CouponImpl) Update(ctx context.Context, in *pb.Coupon) (*pb.Coupon, error) {
	coupon, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	if in.Name != "" {
		coupon.Name = in.Name
	}

	if in.Denomination != 0 {
		coupon.Denomination = in.Denomination
	}
	if in.Count != 0 {
		coupon.Count = in.Count
	}
	if in.Limit != 0 {
		coupon.Limit = in.Limit
	}

	if err := db.Update(couponTable, in.Id, coupon); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *CouponImpl) List(in *pb.User, stream pb.Coupons_ListServer) error {
	coupons := []*pb.Coupon{}
	clause := "WHERE data->'$.owner'='" + in.Id + "'"
	if err := db.List(couponTable, &coupons, clause, "ORDER BY data->'$.created' DESC"); err != nil {
		return err
	}

	for _, v := range coupons {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *CouponImpl) Delete(ctx context.Context, in *pb.Coupon) (*emptypb.Empty, error) {
	if err := db.Delete(couponTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
