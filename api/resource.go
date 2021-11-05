package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/poolPlatform/permission-door/message/npool"
	"github.com/poolPlatform/permission-door/pkg/casbin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteResource(ctx context.Context, in *npool.DeleteResourceRequest) (*emptypb.Empty, error) {
	for _, resource := range in.ResourceIds {
		ok, err := casbin.Enforcer().RemoveFilteredPolicy(1, in.AppId, resource)
		if err != nil {
			logger.Sugar().Errorf("delete resource error: %v", err)
			return new(emptypb.Empty), status.Error(codes.Internal, fmt.Sprintf("delete resource error: %v", err))
		}
		if !ok {
			return new(emptypb.Empty), status.Error(codes.Aborted, fmt.Sprintf("can not delete resource %v", resource))
		}
	}
	return new(emptypb.Empty), status.Error(codes.OK, "delete resource successfully!")
}
