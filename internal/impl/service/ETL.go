package service

// import (
// 	"github.com/emart.io/zbay/internal/impl/biz/db"
// 	olddb "github.com/emart.io/zbay/internal/impl/biz/db/old"
// 	pb "github.com/emart.io/zbay/service/go"
// 	log "github.com/sirupsen/logrus"
// )

// func init() {
// 	commodities := []*pb.Commodity{}
// 	if err := olddb.List(commodityTable, &commodities); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range commodities {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(commodityTable, v.Id, v)
// 	}
// 	// list account
// 	accounts := []*pb.Account{}
// 	if err := olddb.List(accountTable, &accounts); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range accounts {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(accountTable, v.Id, v)
// 	}

// 	// list address
// 	addresses := []*pb.Address{}
// 	if err := olddb.List(addressTable, &addresses); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range addresses {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(addressTable, v.Id, v)
// 	}

// 	// list comments
// 	comments := []*pb.Comment{}
// 	if err := olddb.List(commentTable, &comments); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range comments {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(commentTable, v.Id, v)
// 	}

// 	// list coupons
// 	coupons := []*pb.Coupon{}
// 	if err := olddb.List(couponTable, &coupons); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range coupons {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(couponTable, v.Id, v)
// 	}

// 	// list Memo
// 	memos := []*pb.Memo{}
// 	if err := olddb.List(memoTable, &memos); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range memos {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(memoTable, v.Id, v)
// 	}

// 	// list Message
// 	messages := []*pb.Message{}
// 	if err := olddb.List(messageTable, &messages); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range messages {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(messageTable, v.Id, v)
// 	}

// 	// list Order
// 	orders := []*pb.Order{}
// 	if err := olddb.List(orderTable, &orders); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range orders {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(orderTable, v.Id, v)
// 	}

// 	// list User
// 	users := []*pb.User{}
// 	if err := olddb.List(userTable, &users); err != nil {
// 		log.Errorln(err)
// 	}
// 	for _, v := range users {
// 		if v.Created == nil || v.Created.Seconds == 0 {
// 			continue
// 		}
// 		db.Update(userTable, v.Id, v)
// 	}
// }
