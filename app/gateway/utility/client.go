package utility

import (
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

var Timeout = 2 * time.Second

func ClientConn(service string, opts ...grpc.DialOption) *grpc.ClientConn {
	return grpcx.Client.MustNewGrpcClientConn(service, opts...)
}

func UserClientConn() *grpc.ClientConn {
	return ClientConn("user")
}

func WordClientConn() *grpc.ClientConn {
	return ClientConn("word")
}
