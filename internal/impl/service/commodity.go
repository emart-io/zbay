package service

import (
	"context"
	"github.com/emart.io/zbay/internal/impl/biz"
	"github.com/emart.io/zbay/internal/impl/db"
	pb "github.com/emart.io/zbay/service/go"
	"github.com/gogo/protobuf/types"
	log "github.com/sirupsen/logrus"
)

const (
	commodityTable = "commodities"
)

type CommoditiesImpl struct{}

func (s *CommoditiesImpl) Add(ctx context.Context, in *pb.Commodity) (*pb.Commodity, error) {
	in.Id = types.TimestampNow().String()
	in.Created = types.TimestampNow()
	if err := db.Insert(commodityTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *CommoditiesImpl) Get(ctx context.Context, in *pb.Commodity) (*pb.Commodity, error) {
	commodity := pb.Commodity{}
	if err := db.GetById(commodityTable, in.Id, &commodity); err != nil {
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
	if err := db.Update(commodityTable, in.Id, commodity); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *CommoditiesImpl) List(in *pb.User, stream pb.Commodities_ListServer) error {
	clause := ""
	if !biz.IsAdmin(in) {
		clause = "WHERE data->'$.ownerId'='" + in.Id + "'"
	}
	commodities := []*pb.Commodity{}
	if err := db.List(commodityTable, &commodities, clause, "ORDER BY data->'$.created.seconds' DESC"); err != nil {
		return err
	}

	for _, v := range commodities {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *CommoditiesImpl) Search(in *types.StringValue, stream pb.Commodities_SearchServer) error {
	log.Infoln(in.Value)
	clause := "WHERE data->'$.category' LIKE '%" + in.Value + "%' OR " + " data->'$.title' LIKE '%" + in.Value + "%'"
	commodities := []*pb.Commodity{}
	if err := db.List(commodityTable, &commodities, clause, "ORDER BY data->'$.created.seconds' DESC"); err != nil {
		return err
	}
	if len(commodities) == 0 {
		log.Infoln("commodities==0", "begin no clause query")
		if err := db.List(commodityTable, &commodities, "ORDER BY data->'$.created.seconds' DESC"); err != nil {
			return err
		}
	}

	for _, v := range commodities {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *CommoditiesImpl) Delete(ctx context.Context, in *pb.Commodity) (*types.Empty, error) {
	if err := db.Delete(commodityTable, in.Id); err != nil {
		return nil, err
	}
	return &types.Empty{}, nil
}
