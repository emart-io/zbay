package service

import (
	"context"
	"strings"

	"github.com/emart.io/zbay/internal/impl/biz"
	pb "github.com/emart.io/zbay/service/go"
	"github.com/jmzwcn/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	p2pQueues    = make(map[string]chan *pb.Message)
	pubsubQueues = make(map[string][]chan *pb.Topic)
)

const (
	messageTable = "message"
)

type MessageImpl struct {
	pb.UnimplementedMessagesServer
}

func (s *MessageImpl) Add(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	in.Id = biz.UUID()
	in.Created = timestamppb.Now()
	if err := db.Insert(messageTable, in); err != nil {
		return nil, err
	}

	return in, nil
}

func (s *MessageImpl) List(in *pb.Message, stream pb.Messages_ListServer) error {
	clause := "WHERE (data->'$.to'='" + in.To + "' AND " + " data->'$.from'='" + in.From + "')" +
		" OR (data->'$.to'='" + in.From + "' AND " + " data->'$.from'='" + in.To + "')"
	messages := []*pb.Message{}
	if err := db.List(messageTable, &messages, clause, "ORDER BY data->'$.created' ASC"); err != nil {
		return err
	}

	for _, v := range messages {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *MessageImpl) GroupBy(in *pb.User, stream pb.Messages_GroupByServer) error {
	db.DB.Exec("CREATE TABLE IF NOT EXISTS " + messageTable + " (data JSON)")
	query := "SELECT data->'$.from' FROM " + messageTable + " WHERE data->'$.to'='" + in.Id + "' GROUP BY data->'$.from'"
	rows, err := db.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		from := ""
		rows.Scan(&from)
		from = strings.Trim(from, "\"")
		if from != "" {
			if err := stream.Send(&pb.Message{From: from}); err != nil {
				log.Errorln(err)
				//return err
			}
		}
	}

	return nil
}

func (s *MessageImpl) Send(ctx context.Context, in *pb.Message) (*emptypb.Empty, error) {
	if p2pQueues[in.To] == nil {
		p2pQueues[in.To] = make(chan *pb.Message, 100)
	}
	log.Println("rec:", *in)
	go func() {
		p2pQueues[in.To] <- in
	}()

	return &emptypb.Empty{}, nil
}

func (s *MessageImpl) Receive(in *pb.User, stream pb.Messages_ReceiveServer) error {
	if p2pQueues[in.Id] == nil {
		p2pQueues[in.Id] = make(chan *pb.Message, 100)
	}

	for msg := range p2pQueues[in.Id] {
		log.Println("send", msg)
		if err := stream.Send(msg); err != nil {
			log.Errorln(err)
		}
	}

	return nil
}

func (s *MessageImpl) Publish(ctx context.Context, in *pb.Topic) (*emptypb.Empty, error) {
	if pubsubQueues[in.GroupId] == nil {
		pubsubQueues[in.GroupId] = []chan *pb.Topic{}
	}

	go func() {
		for k := range pubsubQueues[in.GroupId] {
			pubsubQueues[in.GroupId][k] <- in
		}
	}()

	return &emptypb.Empty{}, nil
}

func (s *MessageImpl) Subscribe(in *pb.Topic, stream pb.Messages_SubscribeServer) error {
	if pubsubQueues[in.GroupId] == nil {
		pubsubQueues[in.GroupId] = []chan *pb.Topic{}
	}

	queue := make(chan *pb.Topic, 100)
	pubsubQueues[in.GroupId] = append(pubsubQueues[in.GroupId], queue)
	for topic := range queue {
		if err := stream.Send(topic); err != nil {
			log.Errorln(err)
		}
	}

	return nil
}
