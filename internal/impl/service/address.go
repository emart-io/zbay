package service

import (
	"context"

	"github.com/emart.io/zbay/internal/impl/biz"
	pb "github.com/emart.io/zbay/service/go"
	"github.com/jmzwcn/db"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	addressTable = "addresses"
)

type AddressImpl struct {
	pb.UnimplementedAddressesServer
}

func (s *AddressImpl) Add(ctx context.Context, in *pb.Address) (*pb.Address, error) {
	in.Id = biz.UUID()
	in.Created = timestamppb.Now()
	if err := db.Insert(addressTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *AddressImpl) Get(ctx context.Context, in *pb.Address) (*pb.Address, error) {
	address := pb.Address{}
	if err := db.GetById(addressTable, in.Id, &address); err != nil {
		return nil, err
	}
	return &address, nil
}

func (s *AddressImpl) Update(ctx context.Context, in *pb.Address) (*pb.Address, error) {
	address, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	address.Default = in.Default
	address.Contact = in.Contact
	address.Telephone = in.Telephone
	address.Location = in.Location
	if err := db.Update(addressTable, in.Id, address); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *AddressImpl) List(in *pb.User, stream pb.Addresses_ListServer) error {
	clause := "WHERE data->'$.userId'='" + in.Id + "'"
	addresses := []*pb.Address{}
	if err := db.List(addressTable, &addresses, clause, "ORDER BY data->'$.created' DESC"); err != nil {
		return err
	}

	for _, v := range addresses {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *AddressImpl) Delete(ctx context.Context, in *pb.Address) (*emptypb.Empty, error) {
	if err := db.Delete(addressTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
