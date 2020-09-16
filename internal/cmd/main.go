package main

import (
	"net"
	"net/http"

	impl "github.com/emart.io/zbay/internal/impl/service"
	pb "github.com/emart.io/zbay/service/go"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	go serveFile()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCommoditiesServer(s, &impl.CommoditiesImpl{})
	pb.RegisterUsersServer(s, &impl.UsersImpl{})
	pb.RegisterMessagesServer(s, &impl.MessageImpl{})
	pb.RegisterOrdersServer(s, &impl.OrdersImpl{})
	pb.RegisterAddressesServer(s, &impl.AddressImpl{})
	pb.RegisterCouponsServer(s, &impl.CouponImpl{})
	pb.RegisterAccountsServer(s, &impl.AccountImpl{})
	pb.RegisterCommentsServer(s, &impl.CommentImpl{})
	pb.RegisterMemosServer(s, &impl.MemoImpl{})

	wrappedServer := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor != 2 {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web,authorization-token")
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")
			}
			grpcweb.WrapServer(s).ServeHTTP(w, r)
		}),
	}

	// wrappedServer := &http.Server{
	// 	Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		if r.ProtoMajor != 2 {
	// 			w.Header().Set("Access-Control-Allow-Origin", "*")
	// 			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web,authorization-token")
	// 			w.Header().Set("grpc-status", "")
	// 			w.Header().Set("grpc-message", "")
	// 		}
	// 		grpcweb.WrapServer(s).ServeHTTP(w, r)
	// 	}), &http2.Server{}),
	// }

	log.Infoln("listen:" + port)
	if err := wrappedServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
