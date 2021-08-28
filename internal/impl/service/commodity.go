package service

import (
	"context"
	"os"

	"github.com/emart.io/zbay/internal/impl/biz"
	pb "github.com/emart.io/zbay/service/go"
	"github.com/jmzwcn/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	commodityTable = "commodities"
)

type CommoditiesImpl struct {
	pb.UnimplementedCommoditiesServer
}

func (s *CommoditiesImpl) Add(ctx context.Context, in *pb.Commodity) (*pb.Commodity, error) {
	in.Id = biz.UUID()
	in.Created = timestamppb.Now()
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
	if in.Title != "" {
		os.Rename("/uploads/"+commodity.OwnerId+"/"+commodity.Title, "/uploads/"+commodity.OwnerId+"/"+in.Title)
		commodity.Title = in.Title
	}
	if in.Category != "" {
		commodity.Category = in.Category
	}
	if in.Content != "" {
		commodity.Content = in.Content
	}
	if in.Prices != nil {
		commodity.Prices = in.Prices
	}
	if in.Tags != nil {
		commodity.Tags = in.Tags
	}
	if in.Status != "" {
		commodity.Status = in.Status
	}
	in.Updated = timestamppb.Now()
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
	if err := db.List(commodityTable, &commodities, clause, "ORDER BY data->'$.created' DESC"); err != nil {
		return err
	}

	for _, v := range commodities {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *CommoditiesImpl) Search(in *wrapperspb.StringValue, stream pb.Commodities_SearchServer) error {
	log.Infoln(in.Value)
	clause := "WHERE (data->'$.category' LIKE '%" + in.Value + "%' OR " + " data->'$.title' LIKE '%" + in.Value + "%')" +
		" AND data->'$.status'='已上线'"
	commodities := []*pb.Commodity{}
	if err := db.List(commodityTable, &commodities, clause, "ORDER BY data->'$.created' DESC"); err != nil {
		return err
	}
	if len(commodities) == 0 {
		log.Infoln("commodities==0", "begin no clause query")
		if err := db.List(commodityTable, &commodities, "WHERE data->'$.status'='已上线'", "ORDER BY data->'$.created' DESC"); err != nil {
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

func (s *CommoditiesImpl) Delete(ctx context.Context, in *pb.Commodity) (*emptypb.Empty, error) {
	commodity, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	if err := db.Delete(commodityTable, in.Id); err != nil {
		return nil, err
	}
	os.RemoveAll("/uploads/" + commodity.OwnerId + "/" + commodity.Title)
	return &emptypb.Empty{}, nil
}
