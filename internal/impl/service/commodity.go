package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/emart.io/zbay/internal/impl/biz"
	pb "github.com/emart.io/zbay/service/go"
)

const (
	commodityTable = "commodities"
)

type CommoditiesImpl struct{}

func (s *CommoditiesImpl) Add(ctx context.Context, in *pb.Commodity) (*pb.Commodity, error) {
	in.Id = types.TimestampNow().String()
	in.Created = types.TimestampNow()
	if err := biz.Insert(commodityTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *CommoditiesImpl) Get(ctx context.Context, in *pb.Commodity) (*pb.Commodity, error) {
	commodity := pb.Commodity{}
	if err := biz.GetById(commodityTable, in.Id, &commodity); err != nil {
		return nil, err
	}
	return &commodity, nil
}

func (s *CommoditiesImpl) Update(ctx context.Context, in *pb.Commodity) (*pb.Commodity, error) {
	commodity, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	commodity.Content = in.Content
	if err := biz.Update(commodityTable, in.Id, commodity); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *CommoditiesImpl) ListByUser(in *pb.Commodity, stream pb.Commodities_ListByUserServer) error {
	commodities := []*pb.Commodity{}
	if err := biz.List(commodityTable, &commodities, " order by data->'$.created.seconds' desc"); err != nil {
		return err
	}

	for _, v := range commodities {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *CommoditiesImpl) Delete(ctx context.Context, in *pb.Commodity) (*types.Empty, error) {
	if err := biz.Delete(commodityTable, in.Id); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
