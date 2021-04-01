package service

import (
	"context"

	"github.com/emart.io/zbay/internal/impl/biz/db"
	pb "github.com/emart.io/zbay/service/go"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	commentTable = "comments"
)

type CommentImpl struct {
	pb.UnimplementedCommentsServer
}

func (s *CommentImpl) Add(ctx context.Context, in *pb.Comment) (*pb.Comment, error) {
	in.Id = timestamppb.Now().String()
	in.Created = timestamppb.Now()
	if err := db.Insert(commentTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *CommentImpl) Get(ctx context.Context, in *pb.Comment) (*pb.Comment, error) {
	comment := pb.Comment{}
	if err := db.GetById(commentTable, in.Id, &comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

func (s *CommentImpl) Update(ctx context.Context, in *pb.Comment) (*pb.Comment, error) {
	comment, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	if in.Content != "" {
		comment.Content = in.Content
	}

	if in.Star != 0 {
		comment.Star = in.Star
	}
	if in.Keywords != nil {
		comment.Keywords = in.Keywords
	}

	if err := db.Update(commentTable, in.Id, comment); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *CommentImpl) List(in *pb.Commodity, stream pb.Comments_ListServer) error {
	comments := []*pb.Comment{}
	clause := "WHERE data->'$.commodityId'='" + in.Id + "'"
	if err := db.List(commentTable, &comments, clause, "ORDER BY data->'$.created' DESC"); err != nil {
		return err
	}

	for _, v := range comments {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *CommentImpl) Delete(ctx context.Context, in *pb.Comment) (*emptypb.Empty, error) {
	if err := db.Delete(commentTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
