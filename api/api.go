package api

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/poolPlatform/permission-door/message/npool"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedAnubisServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterAnubisServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterAnubisHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
