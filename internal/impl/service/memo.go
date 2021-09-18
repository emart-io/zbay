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
	memoTable = "memoes"
)

type MemoImpl struct {
	pb.UnimplementedMemosServer
}

func (s *MemoImpl) Add(ctx context.Context, in *pb.Memo) (*pb.Memo, error) {
	in.Id = biz.UUID()
	in.Created = timestamppb.Now()
	if err := db.Insert(memoTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *MemoImpl) Get(ctx context.Context, in *pb.Memo) (*pb.Memo, error) {
	memo := pb.Memo{}
	if err := db.GetById(memoTable, in.Id, &memo); err != nil {
		return nil, err
	}
	return &memo, nil
}

func (s *MemoImpl) Update(ctx context.Context, in *pb.Memo) (*pb.Memo, error) {
	memo, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	if in.Title != "" {
		memo.Title = in.Title
	}
	if in.Content != "" {
		memo.Content = in.Content
	}
	memo.Updated = timestamppb.Now()
	if err := db.Update(memoTable, in.Id, memo); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *MemoImpl) List(in *pb.User, stream pb.Memos_ListServer) error {
	clause := "WHERE data->'$.userId'='" + in.Id + "'"
	memoes := []*pb.Memo{}
	if err := db.List(memoTable, &memoes, clause, "ORDER BY data->'$.created' DESC"); err != nil {
		return err
	}

	for _, v := range memoes {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *MemoImpl) Delete(ctx context.Context, in *pb.Memo) (*emptypb.Empty, error) {
	if err := db.Delete(memoTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
