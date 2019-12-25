package service

import (
	"context"
	"github.com/emart.io/zbay/internal/impl/biz"
	"github.com/emart.io/zbay/internal/impl/db"
	pb "github.com/emart.io/zbay/service/go"
	"github.com/gogo/protobuf/types"
)

const (
	accountTable = "accounts"
)

type AccountImpl struct{}

func (s *AccountImpl) Add(ctx context.Context, in *pb.Account) (*pb.Account, error) {
	in.Id = types.TimestampNow().String()
	in.Created = types.TimestampNow()
	if err := db.Insert(accountTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *AccountImpl) Total(ctx context.Context, in *pb.Account) (*pb.Account, error) {
	//workaround
	accounts := []*pb.Account{}
	if err := db.List(accountTable, &accounts, "WHERE data->'$.userId'='"+in.UserId+"'"); err != nil {
		return nil, err
	}
	var totalAmount int64
	for _, v := range accounts {
		totalAmount = totalAmount + v.Amount
	}
	in.Amount = totalAmount
	return in, nil
}

func (s *AccountImpl) List(in *pb.User, stream pb.Accounts_ListServer) error {
	accounts := []*pb.Account{}
	if err := db.List(accountTable, &accounts, "order by data->'$.created' desc"); err != nil {
		return err
	}

	for _, v := range accounts {
		if err := stream.Send(v); err != nil {
			return err
		}
	}
	return nil
}

func (s *AccountImpl) SignAlipay(ctx context.Context, in *pb.Order) (*types.StringValue, error) {
	v, err := biz.SignAlipay(float32(in.Amount) / 100)
	if err != nil {
		return nil, err
	}
	return &types.StringValue{Value: v}, nil
}

func (s *AccountImpl) PrepayWechat(ctx context.Context, in *pb.Order) (*pb.WechatPayParams, error) {
	return biz.WechatPayParams(in)
}
