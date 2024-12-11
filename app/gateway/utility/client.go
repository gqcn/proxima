package utility

import (
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

const Timeout = 3 * time.Second

func ClientConn(service string, opts ...grpc.DialOption) *grpc.ClientConn {
	return grpcx.Client.MustNewGrpcClientConn(service, opts...)
}

func UserClientConn(opts ...grpc.DialOption) *grpc.ClientConn {
	return ClientConn("user", opts...)
}

func WordClientConn(opts ...grpc.DialOption) *grpc.ClientConn {
	return ClientConn("word", opts...)
}
