package service

import (
	"context"

	pb "github.com/emart.io/zbay/service/go"
	"github.com/jmzwcn/db"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	userTable = "users"
)

type UsersImpl struct {
	pb.UnimplementedUsersServer
}

func (s *UsersImpl) Add(ctx context.Context, in *pb.User) (*pb.User, error) {
	in.Id = in.Telephone
	in.Created = timestamppb.Now() // timestamppb.Now()
	if err := db.Insert(userTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *UsersImpl) Get(ctx context.Context, in *pb.User) (*pb.User, error) {
	user := pb.User{}
	if err := db.GetById(userTable, in.Id, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UsersImpl) Update(ctx context.Context, in *pb.User) (*pb.User, error) {
	user, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	if in.Icon != "" {
		user.Icon = in.Icon
	}
	if in.Signature != "" {
		user.Signature = in.Signature
	}
	if in.Shops != nil {
		user.Shops = in.Shops
	}
	if err := db.Update(userTable, in.Id, user); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *UsersImpl) List(in *pb.User, stream pb.Users_ListServer) error {
	users := []*pb.User{}
	if err := db.List(userTable, &users, " order by data->'$.created' desc"); err != nil {
		return err
	}

	for _, v := range users {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *UsersImpl) Delete(ctx context.Context, in *pb.User) (*emptypb.Empty, error) {
	if err := db.Delete(userTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UsersImpl) Login(ctx context.Context, in *pb.User) (*pb.User, error) {
	user := pb.User{}
	err := db.Get(userTable, map[string]interface{}{"$.telephone": in.Telephone, "$.password": in.Password}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UsersImpl) Certificate(ctx context.Context, in *pb.User) (*pb.User, error) {
	user, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	user.Cert = in.Cert
	if err := db.Update(userTable, in.Id, user); err != nil {
		return nil, err
	}

	return in, nil
}
