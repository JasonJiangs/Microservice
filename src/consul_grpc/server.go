package consul_grpc

import (
	"consul_grpc/pb"
	"google.golang.org/grpc"
)

type Children struct {
}

func main() {
	// init grpc
	grpcServer := grpc.NewServer()
	// register service
	pb.RegisterRpcServiceServer(grpcServer, )

}
