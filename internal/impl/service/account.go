package service

import (
	"context"
	"github.com/emart.io/zbay/internal/impl/biz"
	"github.com/emart.io/zbay/internal/impl/biz/db"
	pb "github.com/emart.io/zbay/service/go"
	"github.com/gogo/protobuf/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
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

func (s *AccountImpl) Alipay(ctx context.Context, in *pb.PayMap) (*pb.PayMap, error) {
	return biz.AlipayRequest(in)
}

func (s *AccountImpl) WechatPay(ctx context.Context, in *pb.PayMap) (*pb.PayMap, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	ip := md.Get("x-forwarded-for")
	log.Infoln(md, ip)
	return biz.PayRequest(in, ip[0])
}
